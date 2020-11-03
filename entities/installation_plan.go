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

package entities

import (
	integrationv1 "github.com/redhat-integration/integration-operator/api/v1"
)

const (
	// ClusterMode is a cluster wide installation mode
	clusterMode = "cluster"
	// NamespaceMode is a single namespace installation mode
	namespaceMode = "namespace"
	// Namespace for cluster wide operators
	clusterModeNamespace = "openshift-operators"
)

var (
	installationPlanFor3scale = InstallationPlan{
		channel:       "threescale-2.9",
		conditionType: "3scaleOperatorInstalled",
		enabled:       false,
		mode:          namespaceMode,
		namespace:     "rhi-3scale",
		packageName:   "3scale-operator",
	}

	installationPlanFor3scaleAPIcast = InstallationPlan{
		channel:       "threescale-2.9",
		conditionType: "3scaleAPIcastOperatorInstalled",
		enabled:       false,
		mode:          namespaceMode,
		namespace:     "rhi-3scale-apicast",
		packageName:   "apicast-operator",
	}

	installationPlanForAMQBroker = InstallationPlan{
		channel:       "current",
		conditionType: "AMQBrokerOperatorInstalled",
		enabled:       false,
		mode:          namespaceMode,
		namespace:     "rhi-amq-broker",
		packageName:   "amq-broker",
	}
	installationPlanForAMQInterconnect = InstallationPlan{
		channel:       "1.2.0",
		conditionType: "AMQInterconnectOperatorInstalled",
		enabled:       false,
		mode:          namespaceMode,
		namespace:     "rhi-amq-interconnect",
		packageName:   "amq7-interconnect-operator",
	}

	installationPlanForAMQStreams = InstallationPlan{
		channel:       "stable",
		conditionType: "AMQStreamsOperatorInstalled",
		enabled:       false,
		mode:          namespaceMode,
		namespace:     "rhi-amq-streams",
		packageName:   "amq-streams",
	}
	installationPlanForAPIDesigner = InstallationPlan{
		channel:       "alpha",
		conditionType: "APIDesignerOperatorInstalled",
		enabled:       false,
		mode:          namespaceMode,
		namespace:     "rhi-api-designer",
		packageName:   "fuse-apicurito",
	}
	installationPlanForCamelK = InstallationPlan{
		channel:       "techpreview",
		conditionType: "CamelKOperatorInstalled",
		enabled:       false,
		mode:          namespaceMode,
		namespace:     "rhi-camel-k",
		packageName:   "red-hat-camel-k",
	}
	installationPlanForFuseConsole = InstallationPlan{
		channel:       "alpha",
		conditionType: "FuseConsoleOperatorInstalled",
		enabled:       false,
		mode:          namespaceMode,
		namespace:     "rhi-fuse-console",
		packageName:   "fuse-console",
	}
	installationPlanForFuseOnline = InstallationPlan{
		channel:       "alpha",
		conditionType: "FuseOnlineOperatorInstalled",
		enabled:       false,
		mode:          namespaceMode,
		namespace:     "rhi-fuse-online",
		packageName:   "fuse-online",
	}
	installationPlanForServiceRegistry = InstallationPlan{
		channel:       "serviceregistry-1.0",
		conditionType: "ServiceRegistryOperatorInstalled",
		enabled:       false,
		mode:          namespaceMode,
		namespace:     "rhi-service-registry",
		packageName:   "service-registry-operator",
	}
)

// InstallationPlan defines the information required for the installation of an operator via OLM
type InstallationPlan struct {
	channel       string
	conditionType string
	enabled       bool
	mode          string
	namespace     string
	packageName   string
}

// GetChannel returns the Operator's update channel
func (ip *InstallationPlan) GetChannel() string {
	return ip.channel
}

// GetConditionType returns the Operator's condition type to be used in the CR status conditions
func (ip *InstallationPlan) GetConditionType() string {
	return ip.conditionType
}

// IsEnabled is used to check whether an installation is enabled or disabled
func (ip *InstallationPlan) IsEnabled() bool {
	return ip.enabled
}

// SetEnabled is used to enable or disable an installation
func (ip *InstallationPlan) SetEnabled(enabled bool) {
	ip.enabled = enabled
}

// IsNamespaceMode returns true when the installation mode is 'namespace'
func (ip *InstallationPlan) IsNamespaceMode() bool {
	return ip.mode == namespaceMode
}

// IsClusterMode returns true when the installation mode is 'cluster'
func (ip *InstallationPlan) IsClusterMode() bool {
	return ip.mode == clusterMode
}

// GetNamespace returns the installation namespace
func (ip *InstallationPlan) GetNamespace() string {
	return ip.namespace
}

// GetPackageName returns the Operator's package name
func (ip *InstallationPlan) GetPackageName() string {
	return ip.packageName
}

// newInstallationPlan creates a new installation plan using installationPlan as the base and overriding some values from installationInput
func newInstallationPlan(installationPlan InstallationPlan, installationInput *integrationv1.InstallationInput) *InstallationPlan {
	installationPlan.enabled = installationInput.Enabled
	installationPlan.mode = installationInput.Mode

	if installationInput.Mode == clusterMode {
		installationPlan.namespace = clusterModeNamespace
	} else if installationInput.Namespace != "" {
		installationPlan.namespace = installationInput.Namespace
	}

	return &installationPlan
}

// CreateInstallationPlans returns installation plans updated with values from the Installation CR
func CreateInstallationPlans(i *integrationv1.Installation) []*InstallationPlan {
	installationPlans := []*InstallationPlan{}

	if i.Spec.ThreeScaleInstallationInput != nil {
		installationInput := integrationv1.InstallationInput(*i.Spec.ThreeScaleInstallationInput)
		installationPlan := newInstallationPlan(installationPlanFor3scale, &installationInput)
		installationPlans = append(installationPlans, installationPlan)
	}
	if i.Spec.ThreeScaleAPIcastInstallationInput != nil {
		installationInput := integrationv1.InstallationInput(*i.Spec.ThreeScaleAPIcastInstallationInput)
		installationPlan := newInstallationPlan(installationPlanFor3scaleAPIcast, &installationInput)
		installationPlans = append(installationPlans, installationPlan)
	}
	if i.Spec.AMQBrokerInstallationInput != nil {
		installationInput := integrationv1.InstallationInput(*i.Spec.AMQBrokerInstallationInput)
		installationPlan := newInstallationPlan(installationPlanForAMQBroker, &installationInput)
		installationPlans = append(installationPlans, installationPlan)
	}
	if i.Spec.AMQInterconnectInstallationInput != nil {
		installationInput := integrationv1.InstallationInput(*i.Spec.AMQInterconnectInstallationInput)
		installationPlan := newInstallationPlan(installationPlanForAMQInterconnect, &installationInput)
		installationPlans = append(installationPlans, installationPlan)
	}
	if i.Spec.AMQStreamsInstallationInput != nil {
		installationInput := integrationv1.InstallationInput(*i.Spec.AMQStreamsInstallationInput)
		installationPlan := newInstallationPlan(installationPlanForAMQStreams, &installationInput)
		installationPlans = append(installationPlans, installationPlan)
	}
	if i.Spec.APIDesignerInstallationInput != nil {
		installationInput := integrationv1.InstallationInput(*i.Spec.APIDesignerInstallationInput)
		installationPlan := newInstallationPlan(installationPlanForAPIDesigner, &installationInput)
		installationPlans = append(installationPlans, installationPlan)
	}
	if i.Spec.CamelKInstallationInput != nil {
		installationInput := integrationv1.InstallationInput(*i.Spec.CamelKInstallationInput)
		installationPlan := newInstallationPlan(installationPlanForCamelK, &installationInput)
		installationPlans = append(installationPlans, installationPlan)
	}
	if i.Spec.FuseConsoleInstallationInput != nil {
		installationInput := integrationv1.InstallationInput(*i.Spec.FuseConsoleInstallationInput)
		installationPlan := newInstallationPlan(installationPlanForFuseConsole, &installationInput)
		installationPlans = append(installationPlans, installationPlan)
	}
	if i.Spec.FuseOnlineInstallationInput != nil {
		installationInput := integrationv1.InstallationInput(*i.Spec.FuseOnlineInstallationInput)
		installationPlan := newInstallationPlan(installationPlanForFuseOnline, &installationInput)
		installationPlans = append(installationPlans, installationPlan)
	}
	if i.Spec.ServiceRegistryInstallationInput != nil {
		installationInput := integrationv1.InstallationInput(*i.Spec.ServiceRegistryInstallationInput)
		installationPlan := newInstallationPlan(installationPlanForServiceRegistry, &installationInput)
		installationPlans = append(installationPlans, installationPlan)
	}

	return installationPlans
}
