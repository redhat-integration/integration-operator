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
	// Package names
	package3scale          = "3scale-operator"
	package3scaleAPIcast   = "apicast-operator"
	packageAMQBroker       = "amq-broker"
	packageAMQInterconnect = "amq7-interconnect-operator"
	packageAMQStreams      = "amq-streams"
	packageAPIDesigner     = "fuse-apicurito"
	packageCamelK          = "red-hat-camel-k"
	packageFuseConsole     = "fuse-console"
	packageFuseOnline      = "fuse-online"
	packageServiceRegistry = "service-registry-operator"
	// Condition types
	conditionType3scale          = "3scaleOperatorInstalled"
	conditionType3scaleAPIcast   = "3scaleAPIcastOperatorInstalled"
	conditionTypeAMQBroker       = "AMQBrokerOperatorInstalled"
	conditionTypeAMQInterconnect = "AMQInterconnectOperatorInstalled"
	conditionTypeAMQStreams      = "AMQStreamsOperatorInstalled"
	conditionTypeAPIDesigner     = "APIDesignerOperatorInstalled"
	conditionTypeCamelK          = "CamelKOperatorInstalled"
	conditionTypeFuseConsole     = "FuseConsoleOperatorInstalled"
	conditionTypeFuseOnline      = "FuseOnlineOperatorInstalled"
	conditionTypeServiceRegistry = "ServiceRegistryOperatorInstalled"
)

type InstallationInputFields struct {
	Enabled   bool
	Channel   string
	Mode      InstallationMode
	Namespace string
	Approval  operatorsv1alpha1.Approval
}

// Installation plan
type ThreeScaleInstallationInputFields struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:threescale-2.6","urn:alm:descriptor:com.tectonic.ui:select:threescale-2.7","urn:alm:descriptor:com.tectonic.ui:select:threescale-2.8","urn:alm:descriptor:com.tectonic.ui:select:threescale-2.9"}
	Channel string `json:"channel"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:3scale-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Approval Strategy",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Automatic","urn:alm:descriptor:com.tectonic.ui:select:Manual"}
	Approval operatorsv1alpha1.Approval `json:"approval"`
}

// Installation plan
type ThreeScaleAPIcastInstallationInputFields struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:threescale-2.8","urn:alm:descriptor:com.tectonic.ui:select:threescale-2.9"}
	Channel string `json:"channel"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:3scale-apicast-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Approval Strategy",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Automatic","urn:alm:descriptor:com.tectonic.ui:select:Manual"}
	Approval operatorsv1alpha1.Approval `json:"approval"`
}

// Installation plan
type AMQBrokerInstallationInputFields struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:current"}
	Channel string `json:"channel"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:amq-broker-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Approval Strategy",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Automatic","urn:alm:descriptor:com.tectonic.ui:select:Manual"}
	Approval operatorsv1alpha1.Approval `json:"approval"`
}

// Installation plan
type AMQInterconnectInstallationInputFields struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:1.2.0"}
	Channel string `json:"channel"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:amq-interconnect-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Approval Strategy",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Automatic","urn:alm:descriptor:com.tectonic.ui:select:Manual"}
	Approval operatorsv1alpha1.Approval `json:"approval"`
}

// Installation plan
type AMQStreamsInstallationInputFields struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:stable","urn:alm:descriptor:com.tectonic.ui:select:amq-streams-1.5.x","urn:alm:descriptor:com.tectonic.ui:select:amq-streams-1.x"}
	Channel string `json:"channel"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Cluster","urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:amq-streams-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Approval Strategy",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Automatic","urn:alm:descriptor:com.tectonic.ui:select:Manual"}
	Approval operatorsv1alpha1.Approval `json:"approval"`
}

// Installation plan
type APIDesignerInstallationInputFields struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:alpha"}
	Channel string `json:"channel"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:api-designer-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Approval Strategy",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Automatic","urn:alm:descriptor:com.tectonic.ui:select:Manual"}
	Approval operatorsv1alpha1.Approval `json:"approval"`
}

// Installation plan
type CamelKInstallationInputFields struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:techpreview"}
	Channel string `json:"channel"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Cluster","urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:camel-k-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Approval Strategy",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Automatic","urn:alm:descriptor:com.tectonic.ui:select:Manual"}
	Approval operatorsv1alpha1.Approval `json:"approval"`
}

// Installation plan
type FuseConsoleInstallationInputFields struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:alpha"}
	Channel string `json:"channel"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:fuse-console-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Approval Strategy",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Automatic","urn:alm:descriptor:com.tectonic.ui:select:Manual"}
	Approval operatorsv1alpha1.Approval `json:"approval"`
}

// Installation plan
type FuseOnlineInstallationInputFields struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:alpha"}
	Channel string `json:"channel"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:fuse-online-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Approval Strategy",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Automatic","urn:alm:descriptor:com.tectonic.ui:select:Manual"}
	Approval operatorsv1alpha1.Approval `json:"approval"`
}

// Installation plan
type ServiceRegistryInstallationInputFields struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:serviceregistry-1","urn:alm:descriptor:com.tectonic.ui:select:serviceregistry-1.0"}
	Channel string `json:"channel"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:service-registry-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Approval Strategy",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Automatic","urn:alm:descriptor:com.tectonic.ui:select:Manual"}
	Approval operatorsv1alpha1.Approval `json:"approval"`
}

// InstallationSpec defines the desired state of Installation
type InstallationSpec struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="3scale Operator"
	ThreeScaleInstallationInputFields ThreeScaleInstallationInputFields `json:"3scale-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="3scale APIcast Operator"
	ThreeScaleAPIcastInstallationInputFields ThreeScaleAPIcastInstallationInputFields `json:"3scale-apicast-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="AMQ Broker Operator"
	AMQBrokerInstallationInputFields AMQBrokerInstallationInputFields `json:"amq-broker-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="AMQ Interconnect Operator"
	AMQInterconnectInstallationInputFields AMQInterconnectInstallationInputFields `json:"amq-interconnect-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="AMQ Streams Operator"
	AMQStreamsInstallationInputFields AMQStreamsInstallationInputFields `json:"amq-streams-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="API Designer Operator"
	APIDesignerInstallationInputFields APIDesignerInstallationInputFields `json:"api-designer-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Camel K Operator"
	CamelKInstallationInputFields CamelKInstallationInputFields `json:"camel-k-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Fuse Console Operator"
	FuseConsoleInstallationInputFields FuseConsoleInstallationInputFields `json:"fuse-console-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Fuse Online Operator"
	FuseOnlineInstallationInputFields FuseOnlineInstallationInputFields `json:"fuse-online-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Service Registry Operator"
	ServiceRegistryInstallationInputFields ServiceRegistryInstallationInputFields `json:"service-registry-installation"`
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
	return []*InstallationPlan{
		{
			PackageName:             package3scaleAPIcast,
			ConditionType:           conditionType3scaleAPIcast,
			InstallationInputFields: InstallationInputFields(i.Spec.ThreeScaleAPIcastInstallationInputFields),
		},
		{
			PackageName:             package3scale,
			ConditionType:           conditionType3scale,
			InstallationInputFields: InstallationInputFields(i.Spec.ThreeScaleInstallationInputFields),
		},
		{
			PackageName:             packageAMQBroker,
			ConditionType:           conditionTypeAMQBroker,
			InstallationInputFields: InstallationInputFields(i.Spec.AMQBrokerInstallationInputFields),
		},
		{
			PackageName:             packageAMQInterconnect,
			ConditionType:           conditionTypeAMQInterconnect,
			InstallationInputFields: InstallationInputFields(i.Spec.AMQInterconnectInstallationInputFields),
		},
		{
			PackageName:             packageAMQStreams,
			ConditionType:           conditionTypeAMQStreams,
			InstallationInputFields: InstallationInputFields(i.Spec.AMQStreamsInstallationInputFields),
		},
		{
			PackageName:             packageAPIDesigner,
			ConditionType:           conditionTypeAPIDesigner,
			InstallationInputFields: InstallationInputFields(i.Spec.APIDesignerInstallationInputFields),
		},
		{
			PackageName:             packageCamelK,
			ConditionType:           conditionTypeCamelK,
			InstallationInputFields: InstallationInputFields(i.Spec.CamelKInstallationInputFields),
		},
		{
			PackageName:             packageFuseConsole,
			ConditionType:           conditionTypeFuseConsole,
			InstallationInputFields: InstallationInputFields(i.Spec.FuseConsoleInstallationInputFields),
		},
		{
			PackageName:             packageFuseOnline,
			ConditionType:           conditionTypeFuseOnline,
			InstallationInputFields: InstallationInputFields(i.Spec.FuseOnlineInstallationInputFields),
		},
		{
			PackageName:             packageServiceRegistry,
			ConditionType:           conditionTypeServiceRegistry,
			InstallationInputFields: InstallationInputFields(i.Spec.ServiceRegistryInstallationInputFields),
		},
	}
}

// UpdateNamespaceForClusterInstallations sets the correct namespace for installation where installation mode is cluster wide
func (i *Installation) UpdateNamespaceForClusterInstallations() {
	if i.Spec.ThreeScaleAPIcastInstallationInputFields.Mode == ClusterMode {
		i.Spec.ThreeScaleAPIcastInstallationInputFields.Namespace = clusterModeNamespace
	}
	if i.Spec.ThreeScaleInstallationInputFields.Mode == ClusterMode {
		i.Spec.ThreeScaleInstallationInputFields.Namespace = clusterModeNamespace
	}
	if i.Spec.AMQBrokerInstallationInputFields.Mode == ClusterMode {
		i.Spec.AMQBrokerInstallationInputFields.Namespace = clusterModeNamespace
	}
	if i.Spec.AMQInterconnectInstallationInputFields.Mode == ClusterMode {
		i.Spec.AMQInterconnectInstallationInputFields.Namespace = clusterModeNamespace
	}
	if i.Spec.AMQStreamsInstallationInputFields.Mode == ClusterMode {
		i.Spec.AMQStreamsInstallationInputFields.Namespace = clusterModeNamespace
	}
	if i.Spec.APIDesignerInstallationInputFields.Mode == ClusterMode {
		i.Spec.APIDesignerInstallationInputFields.Namespace = clusterModeNamespace
	}
	if i.Spec.CamelKInstallationInputFields.Mode == ClusterMode {
		i.Spec.CamelKInstallationInputFields.Namespace = clusterModeNamespace
	}
	if i.Spec.FuseConsoleInstallationInputFields.Mode == ClusterMode {
		i.Spec.FuseConsoleInstallationInputFields.Namespace = clusterModeNamespace
	}
	if i.Spec.FuseOnlineInstallationInputFields.Mode == ClusterMode {
		i.Spec.FuseOnlineInstallationInputFields.Namespace = clusterModeNamespace
	}
	if i.Spec.ServiceRegistryInstallationInputFields.Mode == ClusterMode {
		i.Spec.ServiceRegistryInstallationInputFields.Namespace = clusterModeNamespace
	}
}

// +kubebuilder:object:root=true

// InstallationList contains a list of Installation
type InstallationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Installation `json:"items"`
}

// InstallationPlan defines the information required for the installation of an operator via OLM
type InstallationPlan struct {
	PackageName   string
	ConditionType string
	InstallationInputFields
}

func init() {
	SchemeBuilder.Register(&Installation{}, &InstallationList{})
}
