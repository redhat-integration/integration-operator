// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMQStreamsInstallationPlan) DeepCopyInto(out *AMQStreamsInstallationPlan) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMQStreamsInstallationPlan.
func (in *AMQStreamsInstallationPlan) DeepCopy() *AMQStreamsInstallationPlan {
	if in == nil {
		return nil
	}
	out := new(AMQStreamsInstallationPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIDesignerInstallationPlan) DeepCopyInto(out *APIDesignerInstallationPlan) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIDesignerInstallationPlan.
func (in *APIDesignerInstallationPlan) DeepCopy() *APIDesignerInstallationPlan {
	if in == nil {
		return nil
	}
	out := new(APIDesignerInstallationPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CamelKInstallationPlan) DeepCopyInto(out *CamelKInstallationPlan) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CamelKInstallationPlan.
func (in *CamelKInstallationPlan) DeepCopy() *CamelKInstallationPlan {
	if in == nil {
		return nil
	}
	out := new(CamelKInstallationPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FuseOnlineInstallationPlan) DeepCopyInto(out *FuseOnlineInstallationPlan) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FuseOnlineInstallationPlan.
func (in *FuseOnlineInstallationPlan) DeepCopy() *FuseOnlineInstallationPlan {
	if in == nil {
		return nil
	}
	out := new(FuseOnlineInstallationPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Installation) DeepCopyInto(out *Installation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
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
func (in *InstallationPlan) DeepCopyInto(out *InstallationPlan) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallationPlan.
func (in *InstallationPlan) DeepCopy() *InstallationPlan {
	if in == nil {
		return nil
	}
	out := new(InstallationPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallationSpec) DeepCopyInto(out *InstallationSpec) {
	*out = *in
	out.ThreeScaleInstallationPlan = in.ThreeScaleInstallationPlan
	out.AMQStreamsInstallationPlan = in.AMQStreamsInstallationPlan
	out.APIDesignerInstallationPlan = in.APIDesignerInstallationPlan
	out.CamelKInstallationPlan = in.CamelKInstallationPlan
	out.FuseOnlineInstallationPlan = in.FuseOnlineInstallationPlan
	out.ServiceRegistryInstallationPlan = in.ServiceRegistryInstallationPlan
	out.SSOInstallationPlan = in.SSOInstallationPlan
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
	if in.ProductStatus != nil {
		in, out := &in.ProductStatus, &out.ProductStatus
		*out = make(map[string]ProductStatusValue, len(*in))
		for key, val := range *in {
			(*out)[key] = val
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
func (in *ProductStatusValue) DeepCopyInto(out *ProductStatusValue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductStatusValue.
func (in *ProductStatusValue) DeepCopy() *ProductStatusValue {
	if in == nil {
		return nil
	}
	out := new(ProductStatusValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSOInstallationPlan) DeepCopyInto(out *SSOInstallationPlan) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSOInstallationPlan.
func (in *SSOInstallationPlan) DeepCopy() *SSOInstallationPlan {
	if in == nil {
		return nil
	}
	out := new(SSOInstallationPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceRegistryInstallationPlan) DeepCopyInto(out *ServiceRegistryInstallationPlan) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceRegistryInstallationPlan.
func (in *ServiceRegistryInstallationPlan) DeepCopy() *ServiceRegistryInstallationPlan {
	if in == nil {
		return nil
	}
	out := new(ServiceRegistryInstallationPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThreeScaleInstallationPlan) DeepCopyInto(out *ThreeScaleInstallationPlan) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThreeScaleInstallationPlan.
func (in *ThreeScaleInstallationPlan) DeepCopy() *ThreeScaleInstallationPlan {
	if in == nil {
		return nil
	}
	out := new(ThreeScaleInstallationPlan)
	in.DeepCopyInto(out)
	return out
}
