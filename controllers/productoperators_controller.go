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

	integrationv1 "github.com/redhat-integration/integration-operator/api/v1"
)

// ProductOperatorsReconciler reconciles a ProductOperators object
type ProductOperatorsReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

var (
	phaseInstalling = "Installing"
	phaseInstalled  = "Installed"
)

// +kubebuilder:rbac:groups=integration.redhat.com,resources=productoperators,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=integration.redhat.com,resources=productoperators/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operators.coreos.com,resources=operatorgroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operators.coreos.com,resources=subscriptions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operators.coreos.com,resources=subscriptions/status,verbs=get

// Reconcile is called when watch events happen
func (r *ProductOperatorsReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("product-operators", req.NamespacedName)

	productOperators := &integrationv1.ProductOperators{}
	err := r.Get(ctx, req.NamespacedName, productOperators)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Info("ProductOperators resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get ProductOperators")
		return ctrl.Result{}, err
	}

	productOperatorConfigs := GetProductOperatorConfigs(&productOperators.Spec)

	result, err := r.reconcileStatus(ctx, log, productOperators, "", phaseInstalling)
	if shouldReturn(result, err) {
		return result, err
	}

	for _, config := range productOperatorConfigs {
		result, err := r.reconcileProductOperator(ctx, log, config)
		if shouldReturn(result, err) {
			return result, err
		}
	}

	result, err = r.reconcileStatus(ctx, log, productOperators, phaseInstalling, phaseInstalled)
	if shouldReturn(result, err) {
		return result, err
	}

	return ctrl.Result{}, nil
}

func (r *ProductOperatorsReconciler) reconcileStatus(ctx context.Context, log logr.Logger, productOperators *integrationv1.ProductOperators, currentPhase string, nextPhase string) (ctrl.Result, error) {
	if productOperators.Status.Phase == currentPhase {
		productOperators.Status.Phase = nextPhase
		log.Info("Updating Status", "Phase", productOperators.Status.Phase)
		err := r.Status().Update(ctx, productOperators)
		if err != nil {
			return ctrl.Result{}, err
		}
	}
	return ctrl.Result{}, nil
}

func (r *ProductOperatorsReconciler) reconcileProductOperator(ctx context.Context, log logr.Logger, config ProductOperatorConfig) (ctrl.Result, error) {
	result, err := r.reconcileNamespace(ctx, log, config)
	if shouldReturn(result, err) {
		return result, err
	}

	result, err = r.reconcileOperatorGroup(ctx, log, config)
	if shouldReturn(result, err) {
		return result, err
	}

	result, err = r.reconcileSubscription(ctx, log, config)
	if shouldReturn(result, err) {
		return result, err
	}

	return ctrl.Result{}, nil
}

func (r *ProductOperatorsReconciler) reconcileNamespace(ctx context.Context, log logr.Logger, config ProductOperatorConfig) (ctrl.Result, error) {
	namespace := &corev1.Namespace{}
	err := r.Get(ctx, types.NamespacedName{Name: config.Namespace}, namespace)
	if err != nil {
		if errors.IsNotFound(err) {
			// Define a new Namespace
			namespace = &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: config.Namespace,
				},
			}
			log.Info("Creating a new Namespace", "Namespace.Name", namespace.Name)
			err = r.Create(ctx, namespace)
			if err != nil {
				log.Error(err, "Failed to create new Namespace", "Namespace.Name", namespace.Name)
				return ctrl.Result{}, err
			}
			// Namespace created successfully - return and requeue
			return ctrl.Result{Requeue: true}, nil
		}
		// Error reading the object - requeue the request
		log.Error(err, "Failed to get Namespace", "Namespace.Name", config.Namespace)
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ProductOperatorsReconciler) reconcileOperatorGroup(ctx context.Context, log logr.Logger, config ProductOperatorConfig) (ctrl.Result, error) {
	operatorGroup := &operatorsv1.OperatorGroup{}
	err := r.Get(ctx, types.NamespacedName{Name: config.PackageName, Namespace: config.Namespace}, operatorGroup)
	if err != nil {
		if errors.IsNotFound(err) {
			// Define a new OperatorGroup
			operatorGroup = &operatorsv1.OperatorGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      config.PackageName,
					Namespace: config.Namespace,
				},
				Spec: operatorsv1.OperatorGroupSpec{
					TargetNamespaces: []string{
						config.Namespace,
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
		log.Error(err, "Failed to get OperatorGroup", "OperatorGroup.Name", config.PackageName, "OperatorGroup.Namespace", config.Namespace)
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ProductOperatorsReconciler) reconcileSubscription(ctx context.Context, log logr.Logger, config ProductOperatorConfig) (ctrl.Result, error) {
	subscription := &operatorsv1alpha1.Subscription{}
	err := r.Get(ctx, types.NamespacedName{Name: config.PackageName, Namespace: config.Namespace}, subscription)
	if err != nil {
		if errors.IsNotFound(err) {
			// Define a new Subscription
			subscription = &operatorsv1alpha1.Subscription{
				ObjectMeta: metav1.ObjectMeta{
					Name:      config.PackageName,
					Namespace: config.Namespace,
				},
				Spec: &operatorsv1alpha1.SubscriptionSpec{
					CatalogSource:          CatalogSource,
					CatalogSourceNamespace: CatalogSourceNamespace,
					Channel:                config.UpdateChannel,
					InstallPlanApproval:    InstallPlanApproval,
					Package:                config.PackageName,
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
		log.Error(err, "Failed to get Subscription", "Subscription.Name", config.PackageName, "Subscription.Namespace", config.Namespace)
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func shouldReturn(result ctrl.Result, err error) bool {
	return result.Requeue || result.RequeueAfter > 0 || err != nil
}

// SetupWithManager configures the controller
func (r *ProductOperatorsReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&integrationv1.ProductOperators{}).
		Complete(r)
}
