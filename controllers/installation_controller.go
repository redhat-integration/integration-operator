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

var (
	catalogSource              = "redhat-operators"
	catalogSourceNamespace     = "openshift-marketplace"
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

// Reconcile is called when watch events happen
func (r *InstallationReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("installation", req.NamespacedName)

	installation, err := r.getInstallation(ctx, log, req)
	if installation == nil || err != nil {
		return ctrl.Result{}, err
	}

	reconcileInstallationFuncs := []func(context.Context, logr.Logger, *integrationv1alpha1.Installation) (ctrl.Result, error){
		r.reconcileThreeScaleInstallation,
		r.reconcileAMQStreamsInstallation,
		r.reconcileAPIDesignerInstallation,
		r.reconcileCamelKInstallation,
		r.reconcileFuseOnlineInstallation,
		r.reconcileServiceRegistryInstallation,
		r.reconcileSSOInstallation,
	}

	for _, reconcileInstallationFunc := range reconcileInstallationFuncs {
		result, err := reconcileInstallationFunc(ctx, log, installation)
		if r.shouldReturn(result, err) {
			return result, err
		}
	}

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
			return installation, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get Installation")
		return installation, err
	}

	return installation, nil
}

func (r *InstallationReconciler) reconcileThreeScaleInstallation(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation) (ctrl.Result, error) {
	if installation.Spec.ThreeScaleInstallationPlan.Enabled {
		approvalStrategy := installation.Spec.ThreeScaleInstallationPlan.ApprovalStrategy
		namespace := installation.Spec.ThreeScaleInstallationPlan.Namespace
		updateChannel := installation.Spec.ThreeScaleInstallationPlan.UpdateChannel

		result, err := r.reconcileInstallation(ctx, log, namespace, threeScalePackageName, approvalStrategy, updateChannel)
		if r.shouldReturn(result, err) {
			return result, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileAMQStreamsInstallation(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation) (ctrl.Result, error) {
	if installation.Spec.AMQStreamsInstallationPlan.Enabled {
		approvalStrategy := installation.Spec.AMQStreamsInstallationPlan.ApprovalStrategy
		namespace := installation.Spec.AMQStreamsInstallationPlan.Namespace
		updateChannel := installation.Spec.AMQStreamsInstallationPlan.UpdateChannel

		result, err := r.reconcileInstallation(ctx, log, namespace, amqStreamsPackageName, approvalStrategy, updateChannel)
		if r.shouldReturn(result, err) {
			return result, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileAPIDesignerInstallation(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation) (ctrl.Result, error) {
	if installation.Spec.APIDesignerInstallationPlan.Enabled {
		approvalStrategy := installation.Spec.APIDesignerInstallationPlan.ApprovalStrategy
		namespace := installation.Spec.APIDesignerInstallationPlan.Namespace
		updateChannel := installation.Spec.APIDesignerInstallationPlan.UpdateChannel

		result, err := r.reconcileInstallation(ctx, log, namespace, apiDesignerPackageName, approvalStrategy, updateChannel)
		if r.shouldReturn(result, err) {
			return result, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileCamelKInstallation(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation) (ctrl.Result, error) {
	if installation.Spec.CamelKInstallationPlan.Enabled {
		approvalStrategy := installation.Spec.CamelKInstallationPlan.ApprovalStrategy
		namespace := installation.Spec.CamelKInstallationPlan.Namespace
		updateChannel := installation.Spec.CamelKInstallationPlan.UpdateChannel

		result, err := r.reconcileInstallation(ctx, log, namespace, camelKPackageName, approvalStrategy, updateChannel)
		if r.shouldReturn(result, err) {
			return result, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileFuseOnlineInstallation(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation) (ctrl.Result, error) {
	if installation.Spec.FuseOnlineInstallationPlan.Enabled {
		approvalStrategy := installation.Spec.FuseOnlineInstallationPlan.ApprovalStrategy
		namespace := installation.Spec.FuseOnlineInstallationPlan.Namespace
		updateChannel := installation.Spec.FuseOnlineInstallationPlan.UpdateChannel

		result, err := r.reconcileInstallation(ctx, log, namespace, fuseOnlinePackageName, approvalStrategy, updateChannel)
		if r.shouldReturn(result, err) {
			return result, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileServiceRegistryInstallation(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation) (ctrl.Result, error) {
	if installation.Spec.ServiceRegistryInstallationPlan.Enabled {
		approvalStrategy := installation.Spec.ServiceRegistryInstallationPlan.ApprovalStrategy
		namespace := installation.Spec.ServiceRegistryInstallationPlan.Namespace
		updateChannel := installation.Spec.ServiceRegistryInstallationPlan.UpdateChannel

		result, err := r.reconcileInstallation(ctx, log, namespace, serviceRegistryPackageName, approvalStrategy, updateChannel)
		if r.shouldReturn(result, err) {
			return result, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileSSOInstallation(ctx context.Context, log logr.Logger, installation *integrationv1alpha1.Installation) (ctrl.Result, error) {
	if installation.Spec.SSOInstallationPlan.Enabled {
		approvalStrategy := installation.Spec.SSOInstallationPlan.ApprovalStrategy
		namespace := installation.Spec.SSOInstallationPlan.Namespace
		updateChannel := installation.Spec.SSOInstallationPlan.UpdateChannel

		result, err := r.reconcileInstallation(ctx, log, namespace, ssoPackageName, approvalStrategy, updateChannel)
		if r.shouldReturn(result, err) {
			return result, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *InstallationReconciler) reconcileInstallation(ctx context.Context, log logr.Logger, namespace string, packageName string, approvalStrategy operatorsv1alpha1.Approval, updateChannel string) (ctrl.Result, error) {
	result, err := r.reconcileNamespace(ctx, log, namespace)
	if r.shouldReturn(result, err) {
		return result, err
	}

	result, err = r.reconcileOperatorGroup(ctx, log, namespace, packageName)
	if r.shouldReturn(result, err) {
		return result, err
	}

	result, err = r.reconcileSubscription(ctx, log, namespace, packageName, approvalStrategy, updateChannel)
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

func (r *InstallationReconciler) reconcileOperatorGroup(ctx context.Context, log logr.Logger, namespace string, packageName string) (ctrl.Result, error) {
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

func (r *InstallationReconciler) reconcileSubscription(ctx context.Context, log logr.Logger, namespace string, packageName string, approvalStrategy operatorsv1alpha1.Approval, updateChannel string) (ctrl.Result, error) {
	subscription := &operatorsv1alpha1.Subscription{}
	err := r.Get(ctx, types.NamespacedName{Name: packageName, Namespace: namespace}, subscription)
	if err != nil {
		if errors.IsNotFound(err) {
			// Define a new Subscription
			subscription = &operatorsv1alpha1.Subscription{
				ObjectMeta: metav1.ObjectMeta{
					Name:      packageName,
					Namespace: namespace,
				},
				Spec: &operatorsv1alpha1.SubscriptionSpec{
					CatalogSource:          catalogSource,
					CatalogSourceNamespace: catalogSourceNamespace,
					Channel:                updateChannel,
					InstallPlanApproval:    approvalStrategy,
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
		log.Error(err, "Failed to get Subscription", "Subscription.Name", packageName, "Subscription.Namespace", namespace)
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
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
