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
	operatorsv1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	integrationv1 "github.com/redhat-integration/integration-operator/api/v1"
)

// ProductOperatorConfig contains the configuration required to install a product operator
type ProductOperatorConfig struct {
	Namespace     string
	PackageName   string
	UpdateChannel string
}

// Constants
var (
	CatalogSource          = "redhat-operators"
	CatalogSourceNamespace = "openshift-marketplace"
	InstallPlanApproval    = operatorsv1alpha1.ApprovalAutomatic

	product3scaleOperatorConfig = ProductOperatorConfig{
		Namespace:     "integration-3scale-operator",
		PackageName:   "3scale-operator",
		UpdateChannel: "threescale-2.9",
	}
	productAMQStreamsOperatorConfig = ProductOperatorConfig{
		Namespace:     "integration-amq-streams-operator",
		PackageName:   "amq-streams",
		UpdateChannel: "stable",
	}
	productApicuritoOperatorConfig = ProductOperatorConfig{
		Namespace:     "integration-apicurito-operator",
		PackageName:   "fuse-apicurito",
		UpdateChannel: "alpha",
	}
	productCamelKOperatorConfig = ProductOperatorConfig{
		Namespace:     "integration-camel-k-operator",
		PackageName:   "red-hat-camel-k",
		UpdateChannel: "techpreview",
	}
	productFuseOnlineOperatorConfig = ProductOperatorConfig{
		Namespace:     "integration-fuse-online-operator",
		PackageName:   "fuse-online",
		UpdateChannel: "alpha",
	}
	productServiceRegistryOperatorConfig = ProductOperatorConfig{
		Namespace:     "integration-service-registry-operator",
		PackageName:   "service-registry-operator",
		UpdateChannel: "serviceregistry-1.0",
	}
	productSSOOperatorConfig = ProductOperatorConfig{
		Namespace:     "integration-sso-operator",
		PackageName:   "rhsso-operator",
		UpdateChannel: "alpha",
	}
)

// GetProductOperatorConfigs returns the configuration for the operators that should be installed
func GetProductOperatorConfigs(productOperatorsSpec *integrationv1.ProductOperatorsSpec) []ProductOperatorConfig {
	productOperatorConfigs := make([]ProductOperatorConfig, 0)

	if productOperatorsSpec.Install3scaleOperator {
		productOperatorConfigs = append(productOperatorConfigs, product3scaleOperatorConfig)
	}
	if productOperatorsSpec.InstallAMQStreamsOperator {
		productOperatorConfigs = append(productOperatorConfigs, productAMQStreamsOperatorConfig)
	}
	if productOperatorsSpec.InstallApicuritoOperator {
		productOperatorConfigs = append(productOperatorConfigs, productApicuritoOperatorConfig)
	}
	if productOperatorsSpec.InstallCamelKOperator {
		productOperatorConfigs = append(productOperatorConfigs, productCamelKOperatorConfig)
	}
	if productOperatorsSpec.InstallFuseOnlineOperator {
		productOperatorConfigs = append(productOperatorConfigs, productFuseOnlineOperatorConfig)
	}
	if productOperatorsSpec.InstallServiceRegistryOperator {
		productOperatorConfigs = append(productOperatorConfigs, productServiceRegistryOperatorConfig)
	}
	if productOperatorsSpec.InstallSSOOperator {
		productOperatorConfigs = append(productOperatorConfigs, productSSOOperatorConfig)
	}

	return productOperatorConfigs
}
