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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ProductOperatorsSpec defines the desired state of ProductOperators
type ProductOperatorsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Install 3scale Operator",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Install3scaleOperator bool `json:"install-3scale-operator,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Install AMQ Streams Operator",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	InstallAMQStreamsOperator bool `json:"install-amq-streams-operator,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Install Apicurito Operator",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	InstallApicuritoOperator bool `json:"install-apicurito-operator,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Install Camel K Operator",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	InstallCamelKOperator bool `json:"install-camel-k-operator,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Install Fuse Online Operator",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	InstallFuseOnlineOperator bool `json:"install-fuse-online-operator,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Install Service Registry Operator",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	InstallServiceRegistryOperator bool `json:"install-service-registry-operator,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Install Single Sign-On Operator",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	InstallSSOOperator bool `json:"install-sso-operator,omitempty"`
}

// ProductOperatorsStatus defines the observed state of ProductOperators
type ProductOperatorsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Phase string `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ProductOperators is the Schema for the productoperators API
// +kubebuilder:resource:path=productoperators,scope=Cluster
type ProductOperators struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProductOperatorsSpec   `json:"spec,omitempty"`
	Status ProductOperatorsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProductOperatorsList contains a list of ProductOperators
type ProductOperatorsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProductOperators `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProductOperators{}, &ProductOperatorsList{})
}
