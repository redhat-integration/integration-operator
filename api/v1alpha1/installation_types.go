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

type BaseInstallationPlan struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Namespace string `json:"namespace,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Approval Strategy",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:Automatic","urn:alm:descriptor:com.tectonic.ui:select:Manual"}
	ApprovalStrategy operatorsv1alpha1.Approval `json:"approval-strategy,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Enabled bool `json:"enabled,omitempty"`
}

// Installation plan
type ThreeScaleInstallationPlan struct {
	BaseInstallationPlan `json:",omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:threescale-2.6","urn:alm:descriptor:com.tectonic.ui:select:threescale-2.7","urn:alm:descriptor:com.tectonic.ui:select:threescale-2.8","urn:alm:descriptor:com.tectonic.ui:select:threescale-2.9"}
	UpdateChannel string `json:"update-channel,omitempty"`
}

// Installation plan
type AMQStreamsInstallationPlan struct {
	BaseInstallationPlan `json:",omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:stable","urn:alm:descriptor:com.tectonic.ui:select:amq-streams-1.5.x","urn:alm:descriptor:com.tectonic.ui:select:amq-streams-1.x"}
	UpdateChannel string `json:"update-channel,omitempty"`
}

// Installation plan
type APIDesignerInstallationPlan struct {
	BaseInstallationPlan `json:",omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:alpha"}
	UpdateChannel string `json:"update-channel,omitempty"`
}

// Installation plan
type CamelKInstallationPlan struct {
	BaseInstallationPlan `json:",omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:techpreview"}
	UpdateChannel string `json:"update-channel,omitempty"`
}

// Installation plan
type FuseOnlineInstallationPlan struct {
	BaseInstallationPlan `json:",omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:alpha"}
	UpdateChannel string `json:"update-channel,omitempty"`
}

// Installation plan
type ServiceRegistryInstallationPlan struct {
	BaseInstallationPlan `json:",omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:serviceregistry-1","urn:alm:descriptor:com.tectonic.ui:select:serviceregistry-1.0"}
	UpdateChannel string `json:"update-channel,omitempty"`
}

// Installation plan
type SSOInstallationPlan struct {
	BaseInstallationPlan `json:",omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Update Channel",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:select:alpha"}
	UpdateChannel string `json:"update-channel,omitempty"`
}

// InstallationSpec defines the desired state of Installation
type InstallationSpec struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="3scale Operator"
	ThreeScaleInstallationPlan ThreeScaleInstallationPlan `json:"3scale-installation-plan,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="AMQ Streams Operator"
	AMQStreamsInstallationPlan AMQStreamsInstallationPlan `json:"amq-streams-installation-plan,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="API Designer Operator"
	APIDesignerInstallationPlan APIDesignerInstallationPlan `json:"api-designer-installation-plan,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Camel K Operator"
	CamelKInstallationPlan CamelKInstallationPlan `json:"camel-k-installation-plan,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Fuse Online Operator"
	FuseOnlineInstallationPlan FuseOnlineInstallationPlan `json:"fuse-online-installation-plan,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Service Registry Operator"
	ServiceRegistryInstallationPlan ServiceRegistryInstallationPlan `json:"service-registry-installation-plan,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Single Sign-On Operator"
	SSOInstallationPlan SSOInstallationPlan `json:"sso-installation-plan,omitempty"`
}

// InstallationStatus defines the observed state of Installation
type InstallationStatus struct {
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
