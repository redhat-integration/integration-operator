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

package v1alpha1

import (
	operatorsv1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Operator available cluster wide or in a specific namespace
type InstallationMode string

const (
	// ClusterMode is the cluster wide installation mode
	ClusterMode InstallationMode = "Cluster"
	// NamespaceMode is the single namespace installation mode
	NamespaceMode InstallationMode = "Namespace"
	// Namespace for cluster wide operators
	clusterModeNamespace = "openshift-operators"
)

var (
	installationPlanBase3scale = InstallationPlanBase{
		Channel:       "threescale-2.9",
		ConditionType: "3scaleOperatorInstalled",
		PackageName:   "3scale-operator",
	}
	installationPlanBase3scaleAPIcast = InstallationPlanBase{
		Channel:       "threescale-2.9",
		ConditionType: "3scaleAPIcastOperatorInstalled",
		PackageName:   "apicast-operator",
	}
	installationPlanBaseAMQBroker = InstallationPlanBase{
		Channel:       "current",
		ConditionType: "AMQBrokerOperatorInstalled",
		PackageName:   "amq-broker",
	}
	installationPlanBaseAMQInterconnect = InstallationPlanBase{
		Channel:       "1.2.0",
		ConditionType: "AMQInterconnectOperatorInstalled",
		PackageName:   "amq7-interconnect-operator",
	}
	installationPlanBaseAMQStreams = InstallationPlanBase{
		Channel:       "stable",
		ConditionType: "AMQStreamsOperatorInstalled",
		PackageName:   "amq-streams",
	}
	installationPlanBaseAPIDesigner = InstallationPlanBase{
		Channel:       "alpha",
		ConditionType: "APIDesignerOperatorInstalled",
		PackageName:   "fuse-apicurito",
	}
	installationPlanBaseCamelK = InstallationPlanBase{
		Channel:       "techpreview",
		ConditionType: "CamelKOperatorInstalled",
		PackageName:   "red-hat-camel-k",
	}
	installationPlanBaseFuseConsole = InstallationPlanBase{
		Channel:       "alpha",
		ConditionType: "FuseConsoleOperatorInstalled",
		PackageName:   "fuse-console",
	}
	installationPlanBaseFuseOnline = InstallationPlanBase{
		Channel:       "alpha",
		ConditionType: "FuseOnlineOperatorInstalled",
		PackageName:   "fuse-online",
	}
	installationPlanBaseServiceRegistry = InstallationPlanBase{
		Channel:       "serviceregistry-1.0",
		ConditionType: "ServiceRegistryOperatorInstalled",
		PackageName:   "service-registry-operator",
	}
)

type InstallationPlanBase struct {
	Channel       string
	ConditionType string
	PackageName   string
}

type InstallationPlanInput struct {
	// Enabled this operator
	Enabled bool
	Mode    InstallationMode
	// Namespace where the operator will be installed
	Namespace string
}

// InstallationPlan defines the information required for the installation of an operator via OLM
type InstallationPlan struct {
	InstallationPlanBase
	InstallationPlanInput
}

// Installation plan
type ThreeScaleInstallationPlanInput struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Install",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:3scale-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
}

// Installation plan
type ThreeScaleAPIcastInstallationPlanInput struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:3scale-apicast-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
}

// Installation plan
type AMQBrokerInstallationPlanInput struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:amq-broker-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
}

// Installation plan
type AMQInterconnectInstallationPlanInput struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:amq-interconnect-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
}

// Installation plan
type AMQStreamsInstallationPlanInput struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Cluster","urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:amq-streams-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
}

// Installation plan
type APIDesignerInstallationPlanInput struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:api-designer-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
}

// Installation plan
type CamelKInstallationPlanInput struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Cluster","urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:camel-k-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
}

// Installation plan
type FuseConsoleInstallationPlanInput struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:fuse-console-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
}

// Installation plan
type FuseOnlineInstallationPlanInput struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:fuse-online-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
}

// Installation plan
type ServiceRegistryInstallationPlanInput struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:service-registry-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
}

// InstallationSpec defines the desired state of Installation
type InstallationSpec struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="3scale Operator"
	ThreeScaleInstallationPlanInput *ThreeScaleInstallationPlanInput `json:"3scale-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="3scale APIcast Operator"
	ThreeScaleAPIcastInstallationPlanInput *ThreeScaleAPIcastInstallationPlanInput `json:"3scale-apicast-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="AMQ Broker Operator"
	AMQBrokerInstallationPlanInput *AMQBrokerInstallationPlanInput `json:"amq-broker-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="AMQ Interconnect Operator"
	AMQInterconnectInstallationPlanInput *AMQInterconnectInstallationPlanInput `json:"amq-interconnect-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="AMQ Streams Operator"
	AMQStreamsInstallationPlanInput *AMQStreamsInstallationPlanInput `json:"amq-streams-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="API Designer Operator"
	APIDesignerInstallationPlanInput *APIDesignerInstallationPlanInput `json:"api-designer-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Camel K Operator"
	CamelKInstallationPlanInput *CamelKInstallationPlanInput `json:"camel-k-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Fuse Console Operator"
	FuseConsoleInstallationPlanInput *FuseConsoleInstallationPlanInput `json:"fuse-console-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Fuse Online Operator"
	FuseOnlineInstallationPlanInput *FuseOnlineInstallationPlanInput `json:"fuse-online-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Service Registry Operator"
	ServiceRegistryInstallationPlanInput *ServiceRegistryInstallationPlanInput `json:"service-registry-installation,omitempty"`
}

// InstallationStatus defines the observed state of Installation
type InstallationStatus struct {
	Conditions []metav1.Condition                           `json:"conditions,omitempty"`
	Phase      operatorsv1alpha1.ClusterServiceVersionPhase `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Installation is the Schema for the installations API
// +kubebuilder:resource:path=installations,scope=Cluster
// +operator-sdk:csv:customresourcedefinitions:resources={{Namespace,v1},{OperatorGroup,v1},{Subscription,v1alpha1}}
type Installation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstallationSpec   `json:"spec,omitempty"`
	Status InstallationStatus `json:"status,omitempty"`
}

// GetInstallationPlans returns the product operators' installation plans
func (i *Installation) GetInstallationPlans() []*InstallationPlan {
	installationPlans := []*InstallationPlan{}

	if i.Spec.ThreeScaleInstallationPlanInput != nil {
		installationPlans = append(installationPlans, &InstallationPlan{
			InstallationPlanBase:  installationPlanBase3scale,
			InstallationPlanInput: InstallationPlanInput(*i.Spec.ThreeScaleInstallationPlanInput),
		})
	}
	if i.Spec.ThreeScaleAPIcastInstallationPlanInput != nil {
		installationPlans = append(installationPlans, &InstallationPlan{
			InstallationPlanBase:  installationPlanBase3scaleAPIcast,
			InstallationPlanInput: InstallationPlanInput(*i.Spec.ThreeScaleAPIcastInstallationPlanInput),
		})
	}
	if i.Spec.AMQBrokerInstallationPlanInput != nil {
		installationPlans = append(installationPlans, &InstallationPlan{
			InstallationPlanBase:  installationPlanBaseAMQBroker,
			InstallationPlanInput: InstallationPlanInput(*i.Spec.AMQBrokerInstallationPlanInput),
		})
	}
	if i.Spec.AMQInterconnectInstallationPlanInput != nil {
		installationPlans = append(installationPlans, &InstallationPlan{
			InstallationPlanBase:  installationPlanBaseAMQInterconnect,
			InstallationPlanInput: InstallationPlanInput(*i.Spec.AMQInterconnectInstallationPlanInput),
		})
	}
	if i.Spec.AMQStreamsInstallationPlanInput != nil {
		installationPlans = append(installationPlans, &InstallationPlan{
			InstallationPlanBase:  installationPlanBaseAMQStreams,
			InstallationPlanInput: InstallationPlanInput(*i.Spec.AMQStreamsInstallationPlanInput),
		})
	}
	if i.Spec.APIDesignerInstallationPlanInput != nil {
		installationPlans = append(installationPlans, &InstallationPlan{
			InstallationPlanBase:  installationPlanBaseAPIDesigner,
			InstallationPlanInput: InstallationPlanInput(*i.Spec.APIDesignerInstallationPlanInput),
		})
	}
	if i.Spec.CamelKInstallationPlanInput != nil {
		installationPlans = append(installationPlans, &InstallationPlan{
			InstallationPlanBase:  installationPlanBaseCamelK,
			InstallationPlanInput: InstallationPlanInput(*i.Spec.CamelKInstallationPlanInput),
		})
	}
	if i.Spec.FuseConsoleInstallationPlanInput != nil {
		installationPlans = append(installationPlans, &InstallationPlan{
			InstallationPlanBase:  installationPlanBaseFuseConsole,
			InstallationPlanInput: InstallationPlanInput(*i.Spec.FuseConsoleInstallationPlanInput),
		})
	}
	if i.Spec.FuseOnlineInstallationPlanInput != nil {
		installationPlans = append(installationPlans, &InstallationPlan{
			InstallationPlanBase:  installationPlanBaseFuseOnline,
			InstallationPlanInput: InstallationPlanInput(*i.Spec.FuseOnlineInstallationPlanInput),
		})
	}
	if i.Spec.ServiceRegistryInstallationPlanInput != nil {
		installationPlans = append(installationPlans, &InstallationPlan{
			InstallationPlanBase:  installationPlanBaseServiceRegistry,
			InstallationPlanInput: InstallationPlanInput(*i.Spec.ServiceRegistryInstallationPlanInput),
		})
	}

	return installationPlans
}

// UpdateNamespaceForClusterInstallations sets the correct namespace for installation where installation mode is cluster wide
func (i *Installation) UpdateNamespaceForClusterInstallations() {
	if i.Spec.ThreeScaleInstallationPlanInput != nil && i.Spec.ThreeScaleInstallationPlanInput.Mode == ClusterMode {
		i.Spec.ThreeScaleInstallationPlanInput.Namespace = clusterModeNamespace
	}
	if i.Spec.ThreeScaleAPIcastInstallationPlanInput != nil && i.Spec.ThreeScaleAPIcastInstallationPlanInput.Mode == ClusterMode {
		i.Spec.ThreeScaleAPIcastInstallationPlanInput.Namespace = clusterModeNamespace
	}
	if i.Spec.AMQBrokerInstallationPlanInput != nil && i.Spec.AMQBrokerInstallationPlanInput.Mode == ClusterMode {
		i.Spec.AMQBrokerInstallationPlanInput.Namespace = clusterModeNamespace
	}
	if i.Spec.AMQInterconnectInstallationPlanInput != nil && i.Spec.AMQInterconnectInstallationPlanInput.Mode == ClusterMode {
		i.Spec.AMQInterconnectInstallationPlanInput.Namespace = clusterModeNamespace
	}
	if i.Spec.AMQStreamsInstallationPlanInput != nil && i.Spec.AMQStreamsInstallationPlanInput.Mode == ClusterMode {
		i.Spec.AMQStreamsInstallationPlanInput.Namespace = clusterModeNamespace
	}
	if i.Spec.APIDesignerInstallationPlanInput != nil && i.Spec.APIDesignerInstallationPlanInput.Mode == ClusterMode {
		i.Spec.APIDesignerInstallationPlanInput.Namespace = clusterModeNamespace
	}
	if i.Spec.CamelKInstallationPlanInput != nil && i.Spec.CamelKInstallationPlanInput.Mode == ClusterMode {
		i.Spec.CamelKInstallationPlanInput.Namespace = clusterModeNamespace
	}
	if i.Spec.FuseConsoleInstallationPlanInput != nil && i.Spec.FuseConsoleInstallationPlanInput.Mode == ClusterMode {
		i.Spec.FuseConsoleInstallationPlanInput.Namespace = clusterModeNamespace
	}
	if i.Spec.FuseOnlineInstallationPlanInput != nil && i.Spec.FuseOnlineInstallationPlanInput.Mode == ClusterMode {
		i.Spec.FuseOnlineInstallationPlanInput.Namespace = clusterModeNamespace
	}
	if i.Spec.ServiceRegistryInstallationPlanInput != nil && i.Spec.ServiceRegistryInstallationPlanInput.Mode == ClusterMode {
		i.Spec.ServiceRegistryInstallationPlanInput.Namespace = clusterModeNamespace
	}
}

// +kubebuilder:object:root=true

// InstallationList contains a list of Installation
type InstallationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Installation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Installation{}, &InstallationList{})
}
