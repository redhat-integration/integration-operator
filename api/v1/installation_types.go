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

package v1

import (
	operatorsv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InstallationInput struct {
	Enabled   bool
	Mode      string
	Namespace string
}

// Installation plan
type ThreeScaleInstallationInput struct {
	// Enabled this operator
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// Operator available cluster wide or in a specific namespace
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:namespace"}
	// +kubebuilder:validation:Enum=namespace
	Mode string `json:"mode"`
	// Namespace where the operator will be installed
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:3scale-installation.mode:namespace"}
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Namespace string `json:"namespace,omitempty"`
}

// Installation plan
type ThreeScaleAPIcastInstallationInput struct {
	// Enabled this operator
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// Operator available cluster wide or in a specific namespace
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:namespace"}
	// +kubebuilder:validation:Enum=namespace
	Mode string `json:"mode"`
	// Namespace where the operator will be installed
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:3scale-apicast-installation.mode:namespace"}
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Namespace string `json:"namespace,omitempty"`
}

// Installation plan
type AMQBrokerInstallationInput struct {
	// Enabled this operator
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// Operator available cluster wide or in a specific namespace
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:namespace"}
	// +kubebuilder:validation:Enum=namespace
	Mode string `json:"mode"`
	// Namespace where the operator will be installed
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:amq-broker-installation.mode:namespace"}
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Namespace string `json:"namespace,omitempty"`
}

// Installation plan
type AMQInterconnectInstallationInput struct {
	// Enabled this operator
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// Operator available cluster wide or in a specific namespace
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:namespace"}
	// +kubebuilder:validation:Enum=namespace
	Mode string `json:"mode"`
	// Namespace where the operator will be installed
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:amq-interconnect-installation.mode:namespace"}
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Namespace string `json:"namespace,omitempty"`
}

// Installation plan
type AMQStreamsInstallationInput struct {
	// Enabled this operator
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// Operator available cluster wide or in a specific namespace
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:cluster","urn:alm:descriptor:com.tectonic.ui:select:namespace"}
	// +kubebuilder:validation:Enum=namespace;cluster
	Mode string `json:"mode"`
	// Namespace where the operator will be installed
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:amq-streams-installation.mode:namespace"}
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Namespace string `json:"namespace,omitempty"`
}

// Installation plan
type APIDesignerInstallationInput struct {
	// Enabled this operator
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// Operator available cluster wide or in a specific namespace
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:namespace"}
	// +kubebuilder:validation:Enum=namespace
	Mode string `json:"mode"`
	// Namespace where the operator will be installed
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:api-designer-installation.mode:namespace"}
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Namespace string `json:"namespace,omitempty"`
}

// Installation plan
type CamelKInstallationInput struct {
	// Enabled this operator
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// Operator available cluster wide or in a specific namespace
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:cluster","urn:alm:descriptor:com.tectonic.ui:select:namespace"}
	// +kubebuilder:validation:Enum=namespace;cluster
	Mode string `json:"mode"`
	// Namespace where the operator will be installed
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:camel-k-installation.mode:namespace"}
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Namespace string `json:"namespace,omitempty"`
}

// Installation plan
type FuseConsoleInstallationInput struct {
	// Enabled this operator
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// Operator available cluster wide or in a specific namespace
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:namespace"}
	// +kubebuilder:validation:Enum=namespace
	Mode string `json:"mode"`
	// Namespace where the operator will be installed
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:fuse-console-installation.mode:namespace"}
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Namespace string `json:"namespace,omitempty"`
}

// Installation plan
type FuseOnlineInstallationInput struct {
	// Enabled this operator
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// Operator available cluster wide or in a specific namespace
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:namespace"}
	// +kubebuilder:validation:Enum=namespace
	Mode string `json:"mode"`
	// Namespace where the operator will be installed
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:fuse-online-installation.mode:namespace"}
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Namespace string `json:"namespace,omitempty"`
}

// Installation plan
type ServiceRegistryInstallationInput struct {
	// Enabled this operator
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// Operator available cluster wide or in a specific namespace
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:namespace"}
	// +kubebuilder:validation:Enum=namespace
	Mode string `json:"mode"`
	// Namespace where the operator will be installed
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:service-registry-installation.mode:namespace"}
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Namespace string `json:"namespace,omitempty"`
}

// InstallationSpec defines the desired state of Installation
type InstallationSpec struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="3scale Operator"
	ThreeScaleInstallationInput *ThreeScaleInstallationInput `json:"3scale-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="3scale APIcast Operator"
	ThreeScaleAPIcastInstallationInput *ThreeScaleAPIcastInstallationInput `json:"3scale-apicast-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="AMQ Broker Operator"
	AMQBrokerInstallationInput *AMQBrokerInstallationInput `json:"amq-broker-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="AMQ Interconnect Operator"
	AMQInterconnectInstallationInput *AMQInterconnectInstallationInput `json:"amq-interconnect-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="AMQ Streams Operator"
	AMQStreamsInstallationInput *AMQStreamsInstallationInput `json:"amq-streams-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="API Designer Operator"
	APIDesignerInstallationInput *APIDesignerInstallationInput `json:"api-designer-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Camel K Operator"
	CamelKInstallationInput *CamelKInstallationInput `json:"camel-k-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Fuse Console Operator"
	FuseConsoleInstallationInput *FuseConsoleInstallationInput `json:"fuse-console-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Fuse Online Operator"
	FuseOnlineInstallationInput *FuseOnlineInstallationInput `json:"fuse-online-installation,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Service Registry Operator"
	ServiceRegistryInstallationInput *ServiceRegistryInstallationInput `json:"service-registry-installation,omitempty"`
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
