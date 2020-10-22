/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"reflect"
	"time"

	"github.com/go-logr/logr"
	operatorsv1 "github.com/operator-framework/api/pkg/operators/v1"
	operatorsv1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	integrationv1alpha1 "github.com/redhat-integration/integration-operator/api/v1alpha1"
)

const (
	// Values used when creating a new Subscription
	catalogSource          = "redhat-operators"
	catalogSourceNamespace = "openshift-marketplace"
	// Condition reason when installation is disabled
	conditionReasonDisabled = "Disabled"
)

// InstallationReconciler reconciles a Installation object
type InstallationReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=integration.redhat.com,resources=installations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=integration.redhat.com,resources=installations/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operators.coreos.com,resources=operatorgroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operators.coreos.com,resources=subscriptions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operators.coreos.com,resources=clusterserviceversions,verbs=get;list;watch;create;update;patch;delete

// Reconcile is called when watch events happen
func (r *InstallationReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("installation", req.NamespacedName)

	installation, err := r.getInstallation(ctx, log, req)
	if installation == nil || err != nil {
		return ctrl.Result{}, err
	}

	shouldReturn := r.updateNamespaceForClusterInstallations(ctx, log, installation)
	if shouldReturn {
		return ctrl.Result{}, nil
	}

	installationPlans := installation.GetInstallationPlans()

	shouldReturn = r.initializeStatus(ctx, log, installation, installationPlans)
	if shouldReturn {
		return ctrl.Result{}, nil
	}

	for _, installationPlan := range installationPlans {
		if installationPlan.Enabled {
			result, err := r.reconcileInstallationPlan(ctx, log, installation, installationPlan)
			if r.shouldReturn(result, err) {
				return result, err
			}
			shouldReturn = r.updateStatus(ctx, log, installation, installationPlan)
			if shouldReturn {
				return ctrl.Result{}, nil
			}
		}
	}

	if r.isInstalling(installation.Status.Conditions) {
		log.Info("Installation in progress")
		return ctrl.Result{RequeueAfter: 5 * time.Second}, nil
	}

	log.Info("Installation completed")
	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) getInstallation(ctx context.Context, log logr.Logger, req ctrl.Request) (*integrationv1alpha1.Installation, error) {
	installation := &integrationv1alpha1.Installation{}
	err := r.Get(ctx, req.NamespacedName, installation)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Info("Installation resource not found. Ignoring since object must be deleted")
			return nil, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get Installation")
		return nil, err
	}

	return installation, nil
}

func (r *InstallationReconciler) updateNamespaceForClusterInstallations(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation) bool {
	specCopy := *installation.Spec.DeepCopy()
	installation.UpdateNamespaceForClusterInstallations()

	if !reflect.DeepEqual(specCopy, installation.Spec) {
		log.Info("Updating Installation spec")
		err := r.Update(ctx, installation)
		if err != nil {
			log.Error(err, "Failed to update Installation spec")
		}
		// Installation updated successfully - return
		return true
	}

	return false
}

func (r *InstallationReconciler) reconcileInstallationPlan(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation, installationPlan *integrationv1alpha1.InstallationPlan) (ctrl.Result, error) {
	if installationPlan.Mode == integrationv1alpha1.NamespaceMode {
		result, err := r.reconcileNamespace(ctx, log, installationPlan.Namespace)
		if r.shouldReturn(result, err) {
			return result, err
		}

		result, err = r.reconcileOperatorGroup(ctx, log, installationPlan)
		if r.shouldReturn(result, err) {
			return result, err
		}
	}

	result, err := r.reconcileSubscription(ctx, log, installationPlan)
	if r.shouldReturn(result, err) {
		return result, err
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileNamespace(ctx context.Context, log logr.Logger, namespace string) (ctrl.Result, error) {
	ns := &corev1.Namespace{}
	err := r.Get(ctx, types.NamespacedName{Name: namespace}, ns)
	if err != nil {
		if errors.IsNotFound(err) {
			// Define a new Namespace
			ns = &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: namespace,
				},
			}
			log.Info("Creating a new Namespace", "Namespace.Name", ns.Name)
			err = r.Create(ctx, ns)
			if err != nil {
				log.Error(err, "Failed to create new Namespace", "Namespace.Name", ns.Name)
				return ctrl.Result{}, err
			}
			// Namespace created successfully - return and requeue
			return ctrl.Result{Requeue: true}, nil
		}
		// Error reading the object - requeue the request
		log.Error(err, "Failed to get Namespace", "Namespace.Name", namespace)
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileOperatorGroup(ctx context.Context, log logr.Logger, installationPlan *integrationv1alpha1.InstallationPlan) (ctrl.Result, error) {
	name := installationPlan.PackageName
	namespace := installationPlan.Namespace

	operatorGroup := &operatorsv1.OperatorGroup{}
	err := r.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, operatorGroup)
	if err != nil {
		if errors.IsNotFound(err) {
			// Define a new OperatorGroup
			operatorGroup = &operatorsv1.OperatorGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: namespace,
				},
				Spec: operatorsv1.OperatorGroupSpec{
					TargetNamespaces: []string{
						namespace,
					},
				},
			}
			log.Info("Creating a new OperatorGroup", "OperatorGroup.Name", operatorGroup.Name, "OperatorGroup.Namespace", operatorGroup.Namespace)
			err = r.Create(ctx, operatorGroup)
			if err != nil {
				log.Error(err, "Failed to create new OperatorGroup", "OperatorGroup.Name", operatorGroup.Name, "OperatorGroup.Namespace", operatorGroup.Namespace)
				return ctrl.Result{}, err
			}
			// OperatorGroup created successfully - return and requeue
			return ctrl.Result{Requeue: true}, nil
		}
		// Error reading the object - requeue the request
		log.Error(err, "Failed to get OperatorGroup", "OperatorGroup.Name", name, "OperatorGroup.Namespace", namespace)
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileSubscription(ctx context.Context, log logr.Logger, installationPlan *integrationv1alpha1.InstallationPlan) (ctrl.Result, error) {
	name := installationPlan.PackageName
	namespace := installationPlan.Namespace

	subscription := &operatorsv1alpha1.Subscription{}
	err := r.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, subscription)
	if err != nil {
		if errors.IsNotFound(err) {
			// Define a new Subscription
			subscription = &operatorsv1alpha1.Subscription{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: namespace,
				},
				Spec: &operatorsv1alpha1.SubscriptionSpec{
					CatalogSource:          catalogSource,
					CatalogSourceNamespace: catalogSourceNamespace,
					Channel:                installationPlan.Channel,
					InstallPlanApproval:    installationPlan.Approval,
					Package:                installationPlan.PackageName,
				},
			}
			log.Info("Creating a new Subscription", "Subscription.Name", subscription.Name, "Subscription.Namespace", subscription.Namespace)
			err = r.Create(ctx, subscription)
			if err != nil {
				log.Error(err, "Failed to create new Subscription", "Subscription.Name", subscription.Name, "Subscription.Namespace", subscription.Namespace)
				return ctrl.Result{}, err
			}
			// Subscription created successfully - return and requeue
			return ctrl.Result{RequeueAfter: 5 * time.Second}, nil
		}
		// Error reading the object - requeue the request
		log.Error(err, "Failed to get Subscription", "Subscription.Name", name, "Subscription.Namespace", namespace)
		return ctrl.Result{}, err
	}

	// Ensure the update channel is the same as the spec
	if subscription.Spec.Channel != installationPlan.Channel || subscription.Spec.InstallPlanApproval != installationPlan.Approval {
		subscription.Spec.Channel = installationPlan.Channel
		subscription.Spec.InstallPlanApproval = installationPlan.Approval
		err = r.Update(ctx, subscription)
		if err != nil {
			log.Error(err, "Failed to update Subscription", "Subscription.Name", subscription.Name, "Subscription.Namespace", subscription.Namespace)
			return ctrl.Result{}, err
		}
		// Subscription update - return and requeue
		return ctrl.Result{Requeue: true}, nil
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) initializeStatus(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation, installationPlans []*integrationv1alpha1.InstallationPlan) bool {
	if installation.Status.Conditions == nil {
		conditions := []metav1.Condition{}

		for _, installationPlan := range installationPlans {
			var status metav1.ConditionStatus
			var reason string
			if installationPlan.Enabled {
				status = metav1.ConditionUnknown
				reason = string(operatorsv1alpha1.CSVPhaseInstalling)
			} else {
				status = metav1.ConditionFalse
				reason = conditionReasonDisabled
			}
			conditions = append(conditions, metav1.Condition{
				Type:               installationPlan.ConditionType,
				Status:             status,
				LastTransitionTime: metav1.Now(),
				Reason:             reason,
				Message:            "",
			})
		}

		installation.Status.Conditions = conditions
		installation.Status.Phase = r.calculatePhase(conditions)
		log.Info("Initializing Installation status")
		err := r.Status().Update(ctx, installation)
		if err != nil {
			log.Error(err, "Failed to initialize Installation status")
		}
		return true
	}

	return false
}

func (r *InstallationReconciler) updateStatus(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation, installationPlan *integrationv1alpha1.InstallationPlan) bool {
	for i, condition := range installation.Status.Conditions {
		if condition.Type == installationPlan.ConditionType {
			namespace := installationPlan.Namespace

			subscriptionName := installationPlan.PackageName
			subscription := &operatorsv1alpha1.Subscription{}
			err := r.Get(ctx, types.NamespacedName{Name: subscriptionName, Namespace: namespace}, subscription)
			if err != nil && !errors.IsNotFound(err) {
				// Error reading the object - requeue the request
				log.Error(err, "Failed to get Subscription", "Subscription.Name", subscriptionName, "Subscription.Namespace", namespace)
				return true
			}

			csv := &operatorsv1alpha1.ClusterServiceVersion{}

			if err == nil {
				csvName := subscription.Status.InstalledCSV
				err = r.Get(ctx, types.NamespacedName{Name: csvName, Namespace: namespace}, csv)
				if err != nil && !errors.IsNotFound(err) {
					// Error reading the object - requeue the request
					log.Error(err, "Failed to get ClusterServiceVersion", "ClusterServiceVersion.Name", csvName, "ClusterServiceVersion.Namespace", namespace)
					return true
				}
			}

			newCondition := metav1.Condition(condition)

			if newCondition.Reason == conditionReasonDisabled {
				newCondition.Status = metav1.ConditionUnknown
				newCondition.Reason = string(operatorsv1alpha1.CSVPhaseInstalling)
			}

			if err == nil {
				switch csv.Status.Phase {
				case operatorsv1alpha1.CSVPhaseSucceeded:
					newCondition.Status = metav1.ConditionTrue
				case operatorsv1alpha1.CSVPhaseFailed:
					newCondition.Status = metav1.ConditionFalse
				}

				if newCondition.Status != condition.Status {
					newCondition.LastTransitionTime = metav1.Now()
				}

				if csv.Status.Phase != operatorsv1alpha1.CSVPhaseNone {
					newCondition.Reason = string(csv.Status.Phase)
				}

				newCondition.Message = csv.Status.Message
			}

			if !reflect.DeepEqual(newCondition, condition) {
				installation.Status.Conditions[i] = newCondition
				installation.Status.Phase = r.calculatePhase(installation.Status.Conditions)
				log.Info("Updating Installation status")
				err := r.Status().Update(ctx, installation)
				if err != nil {
					log.Error(err, "Failed to update Installation status")
				}
				return true
			}

			break
		}
	}

	return false
}

func (r *InstallationReconciler) calculatePhase(conditions []metav1.Condition) operatorsv1alpha1.ClusterServiceVersionPhase {
	for _, condition := range conditions {
		if condition.Status == metav1.ConditionUnknown {
			return operatorsv1alpha1.CSVPhaseInstalling
		}
		if condition.Status == metav1.ConditionFalse && condition.Reason != conditionReasonDisabled {
			return operatorsv1alpha1.CSVPhaseFailed
		}
	}
	return operatorsv1alpha1.CSVPhaseSucceeded
}

func (r *InstallationReconciler) isInstalling(conditions []metav1.Condition) bool {
	for _, condition := range conditions {
		if condition.Status == metav1.ConditionUnknown {
			return true
		}
	}
	return false
}

func (r *InstallationReconciler) shouldReturn(result ctrl.Result, err error) bool {
	return result.Requeue || result.RequeueAfter > 0 || err != nil
}

// SetupWithManager configures the controller
func (r *InstallationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&integrationv1alpha1.Installation{}).
		Complete(r)
}
