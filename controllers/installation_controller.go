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
	operatorsv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	olmv1 "github.com/redhat-integration/integration-operator/olm/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	integrationv1 "github.com/redhat-integration/integration-operator/api/v1"
)

const (
	// Subscriptions' catalog source
	catalogSource = "redhat-operators"
	// Subscriptions' catalog source namespace
	catalogSourceNamespace = "openshift-marketplace"
	// Subscriptions' approval policy
	approvalPolicy = operatorsv1alpha1.ApprovalAutomatic
	// Condition reason when installation is disabled
	conditionReasonDisabled = "Disabled"
)

// InstallationReconciler reconciles a Installation object
type InstallationReconciler struct {
	client.Client
	Scheme    *runtime.Scheme
	APIReader client.Reader
	Config    *InstallationConfig
}

// +kubebuilder:rbac:groups=integration.redhat.com,resources=installations,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=integration.redhat.com,resources=installations/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch;create
// +kubebuilder:rbac:groups=operators.coreos.com,resources=operatorgroups,verbs=get;list;watch;create
// +kubebuilder:rbac:groups=operators.coreos.com,resources=subscriptions,verbs=get;list;watch;create;update;patch
// +kubebuilder:rbac:groups=operators.coreos.com,resources=clusterserviceversions,verbs=get;list;watch
// +kubebuilder:rbac:groups=packages.operators.coreos.com,resources=packagemanifests,verbs=get;list;watch

// Reconcile is called when watch events happen
func (r *InstallationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	installation := &integrationv1.Installation{}
	shouldReturn, result, err := r.getInstallation(ctx, log, req.NamespacedName, installation)
	if shouldReturn {
		return result, err
	}

	installationPlans := CreateInstallationPlans(installation, r.Config)

	shouldReturn, result, err = r.validateInstallationPlans(ctx, log, installationPlans)
	if shouldReturn {
		return result, err
	}

	shouldReturn, result, err = r.initializeStatus(ctx, log, installation, installationPlans)
	if shouldReturn {
		return result, err
	}

	for _, installationPlan := range installationPlans {
		if installationPlan.Enabled {
			shouldReturn, result, err = r.reconcileInstallationPlan(ctx, log, installation, installationPlan)
			if shouldReturn {
				return result, err
			}
			shouldReturn, result, err = r.updateStatus(ctx, log, installation, installationPlan)
			if shouldReturn {
				return result, err
			}
		}
	}

	if r.isInstalling(installation.Status.Conditions) {
		return ctrl.Result{RequeueAfter: 10 * time.Second}, nil
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) validateInstallationPlans(ctx context.Context, log logr.Logger, installationPlans []*InstallationPlan) (bool, ctrl.Result, error) {
	for _, installationPlan := range installationPlans {
		if installationPlan.Enabled {
			packageManifest := &olmv1.PackageManifest{}
			err := r.APIReader.Get(ctx, types.NamespacedName{Namespace: "openshift-marketplace", Name: installationPlan.PackageName}, packageManifest)
			if err != nil {
				if errors.IsNotFound(err) {
					log.Error(err, "PackageManifest resource not found", "Name", installationPlan.PackageName)
					installationPlan.Enabled = false
					installationPlan.StatusMessage = "operator not found"
					continue
				}
				// Error reading the object - requeue the request.
				log.Error(err, "Failed to get PackageManifest", "Name", installationPlan.PackageName)
				return true, ctrl.Result{}, err
			}

			channelFound := false
			for _, channel := range packageManifest.Status.Channels {
				if channel.Name == installationPlan.Channel {
					channelFound = true
					continue
				}
			}
			if !channelFound {
				log.Error(err, "Update channel not found", "Channel", installationPlan.Channel, "PackageName", installationPlan.PackageName)
				installationPlan.Enabled = false
				installationPlan.StatusMessage = "update channel not found: " + installationPlan.Channel
			}
		}
	}

	return false, ctrl.Result{}, nil
}

func (r *InstallationReconciler) getInstallation(ctx context.Context, log logr.Logger, namespaceName types.NamespacedName, installation *integrationv1.Installation) (bool, ctrl.Result, error) {
	err := r.Get(ctx, namespaceName, installation)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return true, ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get Installation")
		return true, ctrl.Result{}, err
	}

	return false, ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileInstallationPlan(ctx context.Context, log logr.Logger, installation *integrationv1.Installation, installationPlan *InstallationPlan) (bool, ctrl.Result, error) {
	if installationPlan.IsNamespaceMode() {
		shouldReturn, result, err := r.reconcileNamespace(ctx, log, installationPlan.Namespace)
		if shouldReturn {
			return shouldReturn, result, err
		}

		shouldReturn, result, err = r.reconcileOperatorGroup(ctx, log, installationPlan)
		if shouldReturn {
			return shouldReturn, result, err
		}
	}

	shouldReturn, result, err := r.reconcileSubscription(ctx, log, installation, installationPlan)
	if shouldReturn {
		return shouldReturn, result, err
	}

	return false, ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileNamespace(ctx context.Context, log logr.Logger, namespace string) (bool, ctrl.Result, error) {
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
			err = r.Create(ctx, ns)
			if err != nil {
				log.Error(err, "Failed to create Namespace", "Name", namespace)
				return true, ctrl.Result{}, err
			}
			// Namespace created successfully - return and requeue
			return true, ctrl.Result{Requeue: true}, nil
		}
		// Error reading the object - requeue the request
		log.Error(err, "Failed to get Namespace", "Name", namespace)
		return true, ctrl.Result{}, err
	}

	return false, ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileOperatorGroup(ctx context.Context, log logr.Logger, installationPlan *InstallationPlan) (bool, ctrl.Result, error) {
	namespace := installationPlan.Namespace

	operatorGroupList := &operatorsv1.OperatorGroupList{}
	err := r.List(ctx, operatorGroupList, &client.ListOptions{Namespace: namespace})
	if err != nil {
		// Error reading objects - requeue the request
		log.Error(err, "Failed to list OperatorGroups", "Namespace", namespace)
		return true, ctrl.Result{}, err
	}

	if len(operatorGroupList.Items) == 0 {
		// Define a new OperatorGroup
		operatorGroup := &operatorsv1.OperatorGroup{
			ObjectMeta: metav1.ObjectMeta{
				Name:      namespace,
				Namespace: namespace,
			},
			Spec: operatorsv1.OperatorGroupSpec{
				TargetNamespaces: []string{
					namespace,
				},
			},
		}
		err = r.Create(ctx, operatorGroup)
		if err != nil {
			log.Error(err, "Failed to create OperatorGroup", "Name", operatorGroup.Name, "Namespace", operatorGroup.Namespace)
			return true, ctrl.Result{}, err
		}
		// OperatorGroup created successfully - return and requeue
		return true, ctrl.Result{Requeue: true}, nil
	}

	return false, ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileSubscription(ctx context.Context, log logr.Logger, installation *integrationv1.Installation, installationPlan *InstallationPlan) (bool, ctrl.Result, error) {
	name := installationPlan.Name
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
					InstallPlanApproval:    approvalPolicy,
					Package:                installationPlan.PackageName,
				},
			}
			err = r.Create(ctx, subscription)
			if err != nil {
				log.Error(err, "Failed to create Subscription", "Name", subscription.Name, "Namespace", subscription.Namespace)
				return true, ctrl.Result{}, err
			}
			// Subscription created successfully - return and requeue
			return true, ctrl.Result{Requeue: true}, nil
		}
		// Error reading the object - requeue the request
		log.Error(err, "Failed to get Subscription", "Name", name, "Namespace", namespace)
		return true, ctrl.Result{}, err
	}

	// Ensure the update channel is the same as the spec
	if subscription.Spec.Channel != installationPlan.Channel {
		subscription.Spec.Channel = installationPlan.Channel
		err = r.Update(ctx, subscription)
		if err != nil {
			log.Error(err, "Failed to update Subscription", "Name", subscription.Name, "Namespace", subscription.Namespace)
			return true, ctrl.Result{}, err
		}
		// Subscription updated - return and requeue
		return true, ctrl.Result{Requeue: true}, nil
	}

	return false, ctrl.Result{}, nil
}

func (r *InstallationReconciler) initializeStatus(ctx context.Context, log logr.Logger, installation *integrationv1.Installation, installationPlans []*InstallationPlan) (bool, ctrl.Result, error) {
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
				Message:            installationPlan.StatusMessage,
			})
		}

		installation.Status.Conditions = conditions

		installation.Status.Phase = r.calculatePhase(conditions)

		err := r.Status().Update(ctx, installation)
		if err != nil {
			log.Error(err, "Failed to initialize Installation status")
			return true, ctrl.Result{}, err
		}
		return true, ctrl.Result{}, nil
	}

	return false, ctrl.Result{}, nil
}

func (r *InstallationReconciler) updateStatus(ctx context.Context, log logr.Logger, installation *integrationv1.Installation, installationPlan *InstallationPlan) (bool, ctrl.Result, error) {
	for i, condition := range installation.Status.Conditions {
		if condition.Type == installationPlan.ConditionType {
			namespace := installationPlan.Namespace

			subscriptionName := installationPlan.Name
			subscription := &operatorsv1alpha1.Subscription{}
			err := r.Get(ctx, types.NamespacedName{Name: subscriptionName, Namespace: namespace}, subscription)
			if err != nil && !errors.IsNotFound(err) {
				// Error reading the object - requeue the request
				log.Error(err, "Failed to get Subscription", "Name", subscriptionName, "Namespace", namespace)
				return true, ctrl.Result{}, err
			}

			csv := &operatorsv1alpha1.ClusterServiceVersion{}

			if err == nil {
				csvName := subscription.Status.InstalledCSV
				err = r.Get(ctx, types.NamespacedName{Name: csvName, Namespace: namespace}, csv)
				if err != nil && !errors.IsNotFound(err) {
					// Error reading the object - requeue the request
					log.Error(err, "Failed to get ClusterServiceVersion", "Name", csvName, "Namespace", namespace)
					return true, ctrl.Result{}, err
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

				err := r.Status().Update(ctx, installation)
				if err != nil {
					log.Error(err, "Failed to update Installation status")
					return true, ctrl.Result{}, err
				}
				return true, ctrl.Result{}, nil
			}

			break
		}
	}

	return false, ctrl.Result{}, nil
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

// SetupWithManager configures the controller
func (r *InstallationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&integrationv1.Installation{}).
		Complete(r)
}
