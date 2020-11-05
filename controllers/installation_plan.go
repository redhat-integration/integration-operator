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
	integrationv1 "github.com/redhat-integration/integration-operator/api/v1"
)

const (
	clusterMode          = "cluster"
	namespaceMode        = "namespace"
	clusterModeNamespace = "openshift-operators"
)

// InstallationPlan defines the information required for the installation of an operator via OLM
type InstallationPlan struct {
	Channel       string
	ConditionType string
	Enabled       bool
	Mode          string
	Name          string
	Namespace     string
	PackageName   string
}

// IsNamespaceMode returns true when the installation mode is 'namespace'
func (ip *InstallationPlan) IsNamespaceMode() bool {
	return ip.Mode == namespaceMode
}

// IsClusterMode returns true when the installation mode is 'cluster'
func (ip *InstallationPlan) IsClusterMode() bool {
	return ip.Mode == clusterMode
}

// CreateInstallationPlans returns installation plans updated with values from the Installation CR
func CreateInstallationPlans(installation *integrationv1.Installation, config *InstallationConfig) []*InstallationPlan {
	installationPlans := []*InstallationPlan{}

	if installation.Spec.ThreeScaleInstallationInput != nil {
		input := installation.Spec.ThreeScaleInstallationInput
		installationPlans = append(installationPlans, &InstallationPlan{
			Channel:       config.Channel3scale,
			ConditionType: "3scaleOperatorInstalled",
			Enabled:       input.Enabled,
			Mode:          input.Mode,
			Name:          "3scale-operator",
			Namespace:     calculateNamespace(input.Mode, input.Namespace, "rhi-3scale"),
			PackageName:   "3scale-operator",
		})
	}
	if installation.Spec.ThreeScaleAPIcastInstallationInput != nil {
		input := installation.Spec.ThreeScaleAPIcastInstallationInput
		installationPlans = append(installationPlans, &InstallationPlan{
			Channel:       config.Channel3scaleAPIcast,
			ConditionType: "3scaleAPIcastOperatorInstalled",
			Enabled:       input.Enabled,
			Mode:          input.Mode,
			Name:          "3scale-apicast-operator",
			Namespace:     calculateNamespace(input.Mode, input.Namespace, "rhi-3scale-apicast"),
			PackageName:   "apicast-operator",
		})
	}
	if installation.Spec.AMQBrokerInstallationInput != nil {
		input := installation.Spec.AMQBrokerInstallationInput
		installationPlans = append(installationPlans, &InstallationPlan{
			Channel:       config.ChannelAMQBroker,
			ConditionType: "AMQBrokerOperatorInstalled",
			Enabled:       input.Enabled,
			Mode:          input.Mode,
			Name:          "amq-broker-operator",
			Namespace:     calculateNamespace(input.Mode, input.Namespace, "rhi-amq-broker"),
			PackageName:   "amq-broker",
		})
	}
	if installation.Spec.AMQInterconnectInstallationInput != nil {
		input := installation.Spec.AMQInterconnectInstallationInput
		installationPlans = append(installationPlans, &InstallationPlan{
			Channel:       config.ChannelAMQInterconnect,
			ConditionType: "AMQInterconnectOperatorInstalled",
			Enabled:       input.Enabled,
			Mode:          input.Mode,
			Name:          "amq-interconnect-operator",
			Namespace:     calculateNamespace(input.Mode, input.Namespace, "rhi-amq-interconnect"),
			PackageName:   "amq7-interconnect-operator",
		})
	}
	if installation.Spec.AMQStreamsInstallationInput != nil {
		input := installation.Spec.AMQStreamsInstallationInput
		installationPlans = append(installationPlans, &InstallationPlan{
			Channel:       config.ChannelAMQStreams,
			ConditionType: "AMQStreamsOperatorInstalled",
			Enabled:       input.Enabled,
			Mode:          input.Mode,
			Name:          "amq-streams-operator",
			Namespace:     calculateNamespace(input.Mode, input.Namespace, "rhi-amq-streams"),
			PackageName:   "amq-streams",
		})
	}
	if installation.Spec.APIDesignerInstallationInput != nil {
		input := installation.Spec.APIDesignerInstallationInput
		installationPlans = append(installationPlans, &InstallationPlan{
			Channel:       config.ChannelAPIDesigner,
			ConditionType: "APIDesignerOperatorInstalled",
			Enabled:       input.Enabled,
			Mode:          input.Mode,
			Name:          "api-designer-operator",
			Namespace:     calculateNamespace(input.Mode, input.Namespace, "rhi-api-designer"),
			PackageName:   "fuse-apicurito",
		})
	}
	if installation.Spec.CamelKInstallationInput != nil {
		input := installation.Spec.CamelKInstallationInput
		installationPlans = append(installationPlans, &InstallationPlan{
			Channel:       config.ChannelCamelK,
			ConditionType: "CamelKOperatorInstalled",
			Enabled:       input.Enabled,
			Mode:          input.Mode,
			Name:          "camel-k-operator",
			Namespace:     calculateNamespace(input.Mode, input.Namespace, "rhi-camel-k"),
			PackageName:   "red-hat-camel-k",
		})
	}
	if installation.Spec.FuseConsoleInstallationInput != nil {
		input := installation.Spec.FuseConsoleInstallationInput
		installationPlans = append(installationPlans, &InstallationPlan{
			Channel:       config.ChannelFuseConsole,
			ConditionType: "FuseConsoleOperatorInstalled",
			Enabled:       input.Enabled,
			Mode:          input.Mode,
			Name:          "fuse-console-operator",
			Namespace:     calculateNamespace(input.Mode, input.Namespace, "rhi-fuse-console"),
			PackageName:   "fuse-console",
		})
	}
	if installation.Spec.FuseOnlineInstallationInput != nil {
		input := installation.Spec.FuseOnlineInstallationInput
		installationPlans = append(installationPlans, &InstallationPlan{
			Channel:       config.ChannelFuseOnline,
			ConditionType: "FuseOnlineOperatorInstalled",
			Enabled:       input.Enabled,
			Mode:          input.Mode,
			Name:          "fuse-online-operator",
			Namespace:     calculateNamespace(input.Mode, input.Namespace, "rhi-fuse-online"),
			PackageName:   "fuse-online",
		})
	}
	if installation.Spec.ServiceRegistryInstallationInput != nil {
		input := installation.Spec.ServiceRegistryInstallationInput
		installationPlans = append(installationPlans, &InstallationPlan{
			Channel:       config.ChannelServiceRegistry,
			ConditionType: "ServiceRegistryOperatorInstalled",
			Enabled:       input.Enabled,
			Mode:          input.Mode,
			Name:          "service-registry-operator",
			Namespace:     calculateNamespace(input.Mode, input.Namespace, "rhi-service-registry"),
			PackageName:   "service-registry-operator",
		})
	}

	return installationPlans
}

func calculateNamespace(mode string, namespace string, defaultNamespace string) string {
	if mode == clusterMode {
		return clusterModeNamespace
	} else if namespace == "" {
		return defaultNamespace
	}
	return namespace
}
