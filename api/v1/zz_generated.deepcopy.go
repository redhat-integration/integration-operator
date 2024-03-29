// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMQBrokerInstallationInput) DeepCopyInto(out *AMQBrokerInstallationInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMQBrokerInstallationInput.
func (in *AMQBrokerInstallationInput) DeepCopy() *AMQBrokerInstallationInput {
	if in == nil {
		return nil
	}
	out := new(AMQBrokerInstallationInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMQInterconnectInstallationInput) DeepCopyInto(out *AMQInterconnectInstallationInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMQInterconnectInstallationInput.
func (in *AMQInterconnectInstallationInput) DeepCopy() *AMQInterconnectInstallationInput {
	if in == nil {
		return nil
	}
	out := new(AMQInterconnectInstallationInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMQStreamsInstallationInput) DeepCopyInto(out *AMQStreamsInstallationInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMQStreamsInstallationInput.
func (in *AMQStreamsInstallationInput) DeepCopy() *AMQStreamsInstallationInput {
	if in == nil {
		return nil
	}
	out := new(AMQStreamsInstallationInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIDesignerInstallationInput) DeepCopyInto(out *APIDesignerInstallationInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIDesignerInstallationInput.
func (in *APIDesignerInstallationInput) DeepCopy() *APIDesignerInstallationInput {
	if in == nil {
		return nil
	}
	out := new(APIDesignerInstallationInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CamelKInstallationInput) DeepCopyInto(out *CamelKInstallationInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CamelKInstallationInput.
func (in *CamelKInstallationInput) DeepCopy() *CamelKInstallationInput {
	if in == nil {
		return nil
	}
	out := new(CamelKInstallationInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FuseConsoleInstallationInput) DeepCopyInto(out *FuseConsoleInstallationInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FuseConsoleInstallationInput.
func (in *FuseConsoleInstallationInput) DeepCopy() *FuseConsoleInstallationInput {
	if in == nil {
		return nil
	}
	out := new(FuseConsoleInstallationInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FuseOnlineInstallationInput) DeepCopyInto(out *FuseOnlineInstallationInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FuseOnlineInstallationInput.
func (in *FuseOnlineInstallationInput) DeepCopy() *FuseOnlineInstallationInput {
	if in == nil {
		return nil
	}
	out := new(FuseOnlineInstallationInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Installation) DeepCopyInto(out *Installation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Installation.
func (in *Installation) DeepCopy() *Installation {
	if in == nil {
		return nil
	}
	out := new(Installation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Installation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationInput) DeepCopyInto(out *InstallationInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationInput.
func (in *InstallationInput) DeepCopy() *InstallationInput {
	if in == nil {
		return nil
	}
	out := new(InstallationInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationList) DeepCopyInto(out *InstallationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Installation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationList.
func (in *InstallationList) DeepCopy() *InstallationList {
	if in == nil {
		return nil
	}
	out := new(InstallationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstallationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationSpec) DeepCopyInto(out *InstallationSpec) {
	*out = *in
	if in.ThreeScaleInstallationInput != nil {
		in, out := &in.ThreeScaleInstallationInput, &out.ThreeScaleInstallationInput
		*out = new(ThreeScaleInstallationInput)
		**out = **in
	}
	if in.ThreeScaleAPIcastInstallationInput != nil {
		in, out := &in.ThreeScaleAPIcastInstallationInput, &out.ThreeScaleAPIcastInstallationInput
		*out = new(ThreeScaleAPIcastInstallationInput)
		**out = **in
	}
	if in.AMQBrokerInstallationInput != nil {
		in, out := &in.AMQBrokerInstallationInput, &out.AMQBrokerInstallationInput
		*out = new(AMQBrokerInstallationInput)
		**out = **in
	}
	if in.AMQInterconnectInstallationInput != nil {
		in, out := &in.AMQInterconnectInstallationInput, &out.AMQInterconnectInstallationInput
		*out = new(AMQInterconnectInstallationInput)
		**out = **in
	}
	if in.AMQStreamsInstallationInput != nil {
		in, out := &in.AMQStreamsInstallationInput, &out.AMQStreamsInstallationInput
		*out = new(AMQStreamsInstallationInput)
		**out = **in
	}
	if in.APIDesignerInstallationInput != nil {
		in, out := &in.APIDesignerInstallationInput, &out.APIDesignerInstallationInput
		*out = new(APIDesignerInstallationInput)
		**out = **in
	}
	if in.CamelKInstallationInput != nil {
		in, out := &in.CamelKInstallationInput, &out.CamelKInstallationInput
		*out = new(CamelKInstallationInput)
		**out = **in
	}
	if in.FuseConsoleInstallationInput != nil {
		in, out := &in.FuseConsoleInstallationInput, &out.FuseConsoleInstallationInput
		*out = new(FuseConsoleInstallationInput)
		**out = **in
	}
	if in.FuseOnlineInstallationInput != nil {
		in, out := &in.FuseOnlineInstallationInput, &out.FuseOnlineInstallationInput
		*out = new(FuseOnlineInstallationInput)
		**out = **in
	}
	if in.ServiceRegistryInstallationInput != nil {
		in, out := &in.ServiceRegistryInstallationInput, &out.ServiceRegistryInstallationInput
		*out = new(ServiceRegistryInstallationInput)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationSpec.
func (in *InstallationSpec) DeepCopy() *InstallationSpec {
	if in == nil {
		return nil
	}
	out := new(InstallationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationStatus) DeepCopyInto(out *InstallationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationStatus.
func (in *InstallationStatus) DeepCopy() *InstallationStatus {
	if in == nil {
		return nil
	}
	out := new(InstallationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceRegistryInstallationInput) DeepCopyInto(out *ServiceRegistryInstallationInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceRegistryInstallationInput.
func (in *ServiceRegistryInstallationInput) DeepCopy() *ServiceRegistryInstallationInput {
	if in == nil {
		return nil
	}
	out := new(ServiceRegistryInstallationInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThreeScaleAPIcastInstallationInput) DeepCopyInto(out *ThreeScaleAPIcastInstallationInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThreeScaleAPIcastInstallationInput.
func (in *ThreeScaleAPIcastInstallationInput) DeepCopy() *ThreeScaleAPIcastInstallationInput {
	if in == nil {
		return nil
	}
	out := new(ThreeScaleAPIcastInstallationInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThreeScaleInstallationInput) DeepCopyInto(out *ThreeScaleInstallationInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThreeScaleInstallationInput.
func (in *ThreeScaleInstallationInput) DeepCopy() *ThreeScaleInstallationInput {
	if in == nil {
		return nil
	}
	out := new(ThreeScaleInstallationInput)
	in.DeepCopyInto(out)
	return out
}
