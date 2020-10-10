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

const (
	// Values used when creating a new Subscription
	catalogSource          = "redhat-operators"
	catalogSourceNamespace = "openshift-marketplace"
	// Namespace for cluster wide operators
	clusterModeNamespace = "openshift-operators"
	// Operator IDs
	operator3scale          string = "3scale-operator"
	operatorAMQStreams      string = "amq-streams-operator"
	operatorAPIDesigner     string = "api-designer-operator"
	operatorCamelK          string = "camel-k-operator"
	operatorFuseOnline      string = "fuse-online-operator"
	operatorServiceRegistry string = "service-registry-operator"
	operatorSSO             string = "sso-operator"
	// Package names
	package3scale          string = "3scale-operator"
	packageAMQStreams      string = "amq-streams"
	packageAPIDesigner     string = "fuse-apicurito"
	packageCamelK          string = "red-hat-camel-k"
	packageFuseOnline      string = "fuse-online"
	packageServiceRegistry string = "service-registry-operator"
	packageSSO             string = "rhsso-operator"
	// Condition types
	conditionType3scale          string = "3scaleOperatorInstalled"
	conditionTypeAMQStreams      string = "AMQStreamsOperatorInstalled"
	conditionTypeAPIDesigner     string = "APIDesignerOperatorInstalled"
	conditionTypeCamelK          string = "CamelKOperatorInstalled"
	conditionTypeFuseOnline      string = "FuseOnlineOperatorInstalled"
	conditionTypeServiceRegistry string = "ServiceRegistryOperatorInstalled"
	conditionTypeSSO             string = "SingleSignOnOperatorInstalled"
)

var (
	// Map operator IDs to package names
	packageNames = map[string]string{
		operator3scale:          package3scale,
		operatorAMQStreams:      packageAMQStreams,
		operatorAPIDesigner:     packageAPIDesigner,
		operatorCamelK:          packageCamelK,
		operatorFuseOnline:      packageFuseOnline,
		operatorServiceRegistry: packageServiceRegistry,
		operatorSSO:             packageSSO,
	}
	// Map operator IDs to condition types
	conditionTypes = map[string]string{
		operator3scale:          conditionType3scale,
		operatorAMQStreams:      conditionTypeAMQStreams,
		operatorAPIDesigner:     conditionTypeAPIDesigner,
		operatorCamelK:          conditionTypeCamelK,
		operatorFuseOnline:      conditionTypeFuseOnline,
		operatorServiceRegistry: conditionTypeServiceRegistry,
		operatorSSO:             conditionTypeSSO,
	}
	// Map condition types to operator IDs
	operatorIDs = map[string]string{
		conditionType3scale:          operator3scale,
		conditionTypeAMQStreams:      operatorAMQStreams,
		conditionTypeAPIDesigner:     operatorAPIDesigner,
		conditionTypeCamelK:          operatorCamelK,
		conditionTypeFuseOnline:      operatorFuseOnline,
		conditionTypeServiceRegistry: operatorServiceRegistry,
		conditionTypeSSO:             operatorSSO,
	}
)
