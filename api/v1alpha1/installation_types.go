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
	ClusterMode   InstallationMode = "Cluster"
	NamespaceMode InstallationMode = "Namespace"
)

type InstallationPlan struct {
	Enabled   bool
	Channel   string
	Mode      InstallationMode
	Namespace string
	Approval  operatorsv1alpha1.Approval
}

// Installation plan
type ThreeScaleInstallationPlan struct {
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
type AMQStreamsInstallationPlan struct {
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
type APIDesignerInstallationPlan struct {
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
type CamelKInstallationPlan struct {
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
type FuseOnlineInstallationPlan struct {
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
type ServiceRegistryInstallationPlan struct {
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

// Installation plan
type SSOInstallationPlan struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:alpha"}
	Channel string `json:"channel"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Installation Mode",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Namespace"}
	Mode InstallationMode `json:"mode"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:fieldDependency:sso-installation.mode:Namespace"}
	Namespace string `json:"namespace"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Approval Strategy",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Automatic","urn:alm:descriptor:com.tectonic.ui:select:Manual"}
	Approval operatorsv1alpha1.Approval `json:"approval"`
}

// InstallationSpec defines the desired state of Installation
type InstallationSpec struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="3scale Operator"
	ThreeScaleInstallationPlan ThreeScaleInstallationPlan `json:"3scale-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="AMQ Streams Operator"
	AMQStreamsInstallationPlan AMQStreamsInstallationPlan `json:"amq-streams-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="API Designer Operator"
	APIDesignerInstallationPlan APIDesignerInstallationPlan `json:"api-designer-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Camel K Operator"
	CamelKInstallationPlan CamelKInstallationPlan `json:"camel-k-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Fuse Online Operator"
	FuseOnlineInstallationPlan FuseOnlineInstallationPlan `json:"fuse-online-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Service Registry Operator"
	ServiceRegistryInstallationPlan ServiceRegistryInstallationPlan `json:"service-registry-installation"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Single Sign-On Operator"
	SSOInstallationPlan SSOInstallationPlan `json:"sso-installation"`
}

// ProductStatusValue contains the status of a product operator installation
type ProductStatusValue struct {
	// +operator-sdk:csv:customresourcedefinitions:type=status,displayName="Status",xDescriptors="urn:alm:descriptor:text"
	Phase operatorsv1alpha1.ClusterServiceVersionPhase `json:"phase"`
	// +operator-sdk:csv:customresourcedefinitions:type=status,displayName="Status Reason",xDescriptors="urn:alm:descriptor:text"
	Message string `json:"message"`
}

// InstallationStatus defines the observed state of Installation
type InstallationStatus struct {
	Phase         operatorsv1alpha1.ClusterServiceVersionPhase `json:"phase"`
	Message       string                                       `json:"message"`
	ProductStatus map[string]ProductStatusValue                `json:"products"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Installation is the Schema for the installations API
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
