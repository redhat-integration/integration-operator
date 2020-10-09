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
	catalogSource          = "redhat-operators"
	catalogSourceNamespace = "openshift-marketplace"
	clusterModeNamespace   = "openshift-operators"
	// Package names
	threeScalePackageName      = "3scale-operator"
	amqStreamsPackageName      = "amq-streams"
	apiDesignerPackageName     = "fuse-apicurito"
	camelKPackageName          = "red-hat-camel-k"
	fuseOnlinePackageName      = "fuse-online"
	serviceRegistryPackageName = "service-registry-operator"
	ssoPackageName             = "rhsso-operator"
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

	installationPlans := map[string]integrationv1alpha1.InstallationPlan{
		threeScalePackageName:      integrationv1alpha1.InstallationPlan(installation.Spec.ThreeScaleInstallationPlan),
		amqStreamsPackageName:      integrationv1alpha1.InstallationPlan(installation.Spec.AMQStreamsInstallationPlan),
		apiDesignerPackageName:     integrationv1alpha1.InstallationPlan(installation.Spec.APIDesignerInstallationPlan),
		camelKPackageName:          integrationv1alpha1.InstallationPlan(installation.Spec.CamelKInstallationPlan),
		fuseOnlinePackageName:      integrationv1alpha1.InstallationPlan(installation.Spec.FuseOnlineInstallationPlan),
		serviceRegistryPackageName: integrationv1alpha1.InstallationPlan(installation.Spec.ServiceRegistryInstallationPlan),
		ssoPackageName:             integrationv1alpha1.InstallationPlan(installation.Spec.SSOInstallationPlan),
	}

	shouldReturn = r.initializeStatus(ctx, log, installation, installationPlans)
	if shouldReturn {
		return ctrl.Result{}, nil
	}

	result, err := r.reconcileInstallations(ctx, log, installation, installationPlans)
	if r.shouldReturn(result, err) {
		return result, err
	}

	return r.updateStatus(ctx, log, installation, installationPlans)
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
	spec := *installation.Spec.DeepCopy()

	if spec.ThreeScaleInstallationPlan.Mode == integrationv1alpha1.ClusterMode {
		spec.ThreeScaleInstallationPlan.Namespace = clusterModeNamespace
	}
	if spec.AMQStreamsInstallationPlan.Mode == integrationv1alpha1.ClusterMode {
		spec.AMQStreamsInstallationPlan.Namespace = clusterModeNamespace
	}
	if spec.APIDesignerInstallationPlan.Mode == integrationv1alpha1.ClusterMode {
		spec.APIDesignerInstallationPlan.Namespace = clusterModeNamespace
	}
	if spec.CamelKInstallationPlan.Mode == integrationv1alpha1.ClusterMode {
		spec.CamelKInstallationPlan.Namespace = clusterModeNamespace
	}
	if spec.FuseOnlineInstallationPlan.Mode == integrationv1alpha1.ClusterMode {
		spec.FuseOnlineInstallationPlan.Namespace = clusterModeNamespace
	}
	if spec.ServiceRegistryInstallationPlan.Mode == integrationv1alpha1.ClusterMode {
		spec.ServiceRegistryInstallationPlan.Namespace = clusterModeNamespace
	}
	if spec.SSOInstallationPlan.Mode == integrationv1alpha1.ClusterMode {
		spec.SSOInstallationPlan.Namespace = clusterModeNamespace
	}

	if !reflect.DeepEqual(spec, installation.Spec) {
		installation.Spec = spec
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

func (r *InstallationReconciler) reconcileInstallations(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation, installationPlans map[string]integrationv1alpha1.InstallationPlan) (ctrl.Result, error) {
	for packageName, installationPlan := range installationPlans {
		if installationPlan.Enabled {
			result, err := r.reconcileInstallation(ctx, log, installation, packageName, &installationPlan)
			if r.shouldReturn(result, err) {
				return result, err
			}
		}
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileInstallation(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation, packageName string, installationPlan *integrationv1alpha1.InstallationPlan) (ctrl.Result, error) {
	if installationPlan.Mode == integrationv1alpha1.NamespaceMode {
		result, err := r.reconcileNamespace(ctx, log, installationPlan.Namespace)
		if r.shouldReturn(result, err) {
			return result, err
		}

		result, err = r.reconcileOperatorGroupTargetingOwnNamespace(ctx, log, installation, packageName, installationPlan.Namespace)
		if r.shouldReturn(result, err) {
			return result, err
		}
	}

	result, err := r.reconcileSubscription(ctx, log, installation, packageName, installationPlan)
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

func (r *InstallationReconciler) reconcileOperatorGroupTargetingOwnNamespace(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation, packageName string, namespace string) (ctrl.Result, error) {
	operatorGroup := &operatorsv1.OperatorGroup{}
	err := r.Get(ctx, types.NamespacedName{Name: packageName, Namespace: namespace}, operatorGroup)
	if err != nil {
		if errors.IsNotFound(err) {
			// Define a new OperatorGroup
			operatorGroup = &operatorsv1.OperatorGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      packageName,
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
		log.Error(err, "Failed to get OperatorGroup", "OperatorGroup.Name", packageName, "OperatorGroup.Namespace", namespace)
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileSubscription(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation, packageName string, installationPlan *integrationv1alpha1.InstallationPlan) (ctrl.Result, error) {
	subscription := &operatorsv1alpha1.Subscription{}
	err := r.Get(ctx, types.NamespacedName{Name: packageName, Namespace: installationPlan.Namespace}, subscription)
	if err != nil {
		if errors.IsNotFound(err) {
			// Define a new Subscription
			subscription = &operatorsv1alpha1.Subscription{
				ObjectMeta: metav1.ObjectMeta{
					Name:      packageName,
					Namespace: installationPlan.Namespace,
				},
				Spec: &operatorsv1alpha1.SubscriptionSpec{
					CatalogSource:          catalogSource,
					CatalogSourceNamespace: catalogSourceNamespace,
					Channel:                installationPlan.Channel,
					InstallPlanApproval:    installationPlan.Approval,
					Package:                packageName,
				},
			}
			log.Info("Creating a new Subscription", "Subscription.Name", subscription.Name, "Subscription.Namespace", subscription.Namespace)
			err = r.Create(ctx, subscription)
			if err != nil {
				log.Error(err, "Failed to create new Subscription", "Subscription.Name", subscription.Name, "Subscription.Namespace", subscription.Namespace)
				return ctrl.Result{}, err
			}
			// Subscription created successfully - return and requeue
			return ctrl.Result{RequeueAfter: 10 * time.Second}, nil
		}
		// Error reading the object - requeue the request
		log.Error(err, "Failed to get Subscription", "Subscription.Name", packageName, "Subscription.Namespace", installationPlan.Namespace)
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) initializeStatus(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation, installationPlans map[string]integrationv1alpha1.InstallationPlan) bool {
	if installation.Status.ProductStatus == nil {
		statusMap := make(map[string]integrationv1alpha1.ProductStatusValue)

		for packageName, installationPlan := range installationPlans {
			if installationPlan.Enabled {
				statusMap[packageName] = integrationv1alpha1.ProductStatusValue{Phase: operatorsv1alpha1.CSVPhaseInstalling, Message: ""}
			}
		}

		installation.Status.Phase = operatorsv1alpha1.CSVPhaseInstalling
		installation.Status.Message = ""
		installation.Status.ProductStatus = statusMap

		log.Info("Initializing Installation status")
		err := r.Status().Update(ctx, installation)
		if err != nil {
			log.Error(err, "Failed to initialize Installation status")
		}
		return true
	}

	return false
}

func (r *InstallationReconciler) updateStatus(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation, installationPlans map[string]integrationv1alpha1.InstallationPlan) (ctrl.Result, error) {
	statusMap := make(map[string]integrationv1alpha1.ProductStatusValue)

	for packageName, installationPlan := range installationPlans {
		if installationPlan.Enabled {
			subscription := &operatorsv1alpha1.Subscription{}
			err := r.Get(ctx, types.NamespacedName{Name: packageName, Namespace: installationPlan.Namespace}, subscription)
			if err != nil {
				// Error reading the object - requeue the request
				log.Error(err, "Failed to get Subscription", "Subscription.Name", packageName, "Subscription.Namespace", installationPlan.Namespace)
				return ctrl.Result{}, err
			}

			csvName := subscription.Status.InstalledCSV
			csv := &operatorsv1alpha1.ClusterServiceVersion{}
			err = r.Get(ctx, types.NamespacedName{Name: csvName, Namespace: installationPlan.Namespace}, csv)
			if err != nil && !errors.IsNotFound(err) {
				// Error reading the object - requeue the request
				log.Error(err, "Failed to get ClusterServiceVersion", "ClusterServiceVersion.Name", csvName, "ClusterServiceVersion.Namespace", installationPlan.Namespace)
				return ctrl.Result{}, err
			}

			statusMap[packageName] = integrationv1alpha1.ProductStatusValue{Phase: csv.Status.Phase, Message: csv.Status.Message}
		}
	}

	if !reflect.DeepEqual(statusMap, installation.Status.ProductStatus) {
		installation.Status.Phase, installation.Status.Message = r.getOverallPhaseAndMessage(statusMap)
		installation.Status.ProductStatus = statusMap
		log.Info("Updating Installation status")
		err := r.Status().Update(ctx, installation)
		if err != nil {
			log.Error(err, "Failed to update Installation status")
			return ctrl.Result{}, err
		}
	}

	if installation.Status.Phase == operatorsv1alpha1.CSVPhaseInstalling {
		log.Info("Installation in progress")
		return ctrl.Result{RequeueAfter: 10 * time.Second}, nil
	}

	log.Info("Installation completed")
	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) getOverallPhaseAndMessage(statusMap map[string]integrationv1alpha1.ProductStatusValue) (operatorsv1alpha1.ClusterServiceVersionPhase, string) {
	for _, value := range statusMap {
		if value.Phase == operatorsv1alpha1.CSVPhaseFailed {
			return operatorsv1alpha1.CSVPhaseFailed, "Some installations failed"
		}
		if value.Phase != operatorsv1alpha1.CSVPhaseSucceeded {
			return operatorsv1alpha1.CSVPhaseInstalling, ""
		}
	}
	return operatorsv1alpha1.CSVPhaseSucceeded, "All installations succeeded"
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
