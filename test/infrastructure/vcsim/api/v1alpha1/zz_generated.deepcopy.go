//go:build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEnvVarSpec) DeepCopyInto(out *ClusterEnvVarSpec) {
	*out = *in
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(string)
		**out = **in
	}
	if in.ControlPlaneMachines != nil {
		in, out := &in.ControlPlaneMachines, &out.ControlPlaneMachines
		*out = new(int32)
		**out = **in
	}
	if in.WorkerMachines != nil {
		in, out := &in.WorkerMachines, &out.WorkerMachines
		*out = new(int32)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(int32)
		**out = **in
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(int32)
		**out = **in
	}
	if in.Datastore != nil {
		in, out := &in.Datastore, &out.Datastore
		*out = new(int32)
		**out = **in
	}
	if in.PowerOffMode != nil {
		in, out := &in.PowerOffMode, &out.PowerOffMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEnvVarSpec.
func (in *ClusterEnvVarSpec) DeepCopy() *ClusterEnvVarSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterEnvVarSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryConfig) DeepCopyInto(out *ContentLibraryConfig) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContentLibraryItemConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryConfig.
func (in *ContentLibraryConfig) DeepCopy() *ContentLibraryConfig {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItemConfig) DeepCopyInto(out *ContentLibraryItemConfig) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]ContentLibraryItemFilesConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItemConfig.
func (in *ContentLibraryItemConfig) DeepCopy() *ContentLibraryItemConfig {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItemConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItemFilesConfig) DeepCopyInto(out *ContentLibraryItemFilesConfig) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItemFilesConfig.
func (in *ContentLibraryItemFilesConfig) DeepCopy() *ContentLibraryItemFilesConfig {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItemFilesConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneEndpoint) DeepCopyInto(out *ControlPlaneEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneEndpoint.
func (in *ControlPlaneEndpoint) DeepCopy() *ControlPlaneEndpoint {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControlPlaneEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneEndpointList) DeepCopyInto(out *ControlPlaneEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControlPlaneEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneEndpointList.
func (in *ControlPlaneEndpointList) DeepCopy() *ControlPlaneEndpointList {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControlPlaneEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneEndpointSpec) DeepCopyInto(out *ControlPlaneEndpointSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneEndpointSpec.
func (in *ControlPlaneEndpointSpec) DeepCopy() *ControlPlaneEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneEndpointStatus) DeepCopyInto(out *ControlPlaneEndpointStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneEndpointStatus.
func (in *ControlPlaneEndpointStatus) DeepCopy() *ControlPlaneEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvVar) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVarList) DeepCopyInto(out *EnvVarList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVarList.
func (in *EnvVarList) DeepCopy() *EnvVarList {
	if in == nil {
		return nil
	}
	out := new(EnvVarList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvVarList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVarSpec) DeepCopyInto(out *EnvVarSpec) {
	*out = *in
	if in.VCenterSimulator != nil {
		in, out := &in.VCenterSimulator, &out.VCenterSimulator
		*out = new(NamespacedRef)
		**out = **in
	}
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	in.Cluster.DeepCopyInto(&out.Cluster)
	if in.VMOperatorDependencies != nil {
		in, out := &in.VMOperatorDependencies, &out.VMOperatorDependencies
		*out = new(NamespacedRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVarSpec.
func (in *EnvVarSpec) DeepCopy() *EnvVarSpec {
	if in == nil {
		return nil
	}
	out := new(EnvVarSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVarStatus) DeepCopyInto(out *EnvVarStatus) {
	*out = *in
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVarStatus.
func (in *EnvVarStatus) DeepCopy() *EnvVarStatus {
	if in == nil {
		return nil
	}
	out := new(EnvVarStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedRef) DeepCopyInto(out *NamespacedRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedRef.
func (in *NamespacedRef) DeepCopy() *NamespacedRef {
	if in == nil {
		return nil
	}
	out := new(NamespacedRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClass) DeepCopyInto(out *StorageClass) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClass.
func (in *StorageClass) DeepCopy() *StorageClass {
	if in == nil {
		return nil
	}
	out := new(StorageClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VCenterSimulator) DeepCopyInto(out *VCenterSimulator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VCenterSimulator.
func (in *VCenterSimulator) DeepCopy() *VCenterSimulator {
	if in == nil {
		return nil
	}
	out := new(VCenterSimulator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VCenterSimulator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VCenterSimulatorList) DeepCopyInto(out *VCenterSimulatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VCenterSimulator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VCenterSimulatorList.
func (in *VCenterSimulatorList) DeepCopy() *VCenterSimulatorList {
	if in == nil {
		return nil
	}
	out := new(VCenterSimulatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VCenterSimulatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VCenterSimulatorModel) DeepCopyInto(out *VCenterSimulatorModel) {
	*out = *in
	if in.VSphereVersion != nil {
		in, out := &in.VSphereVersion, &out.VSphereVersion
		*out = new(string)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(int32)
		**out = **in
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(int32)
		**out = **in
	}
	if in.ClusterHost != nil {
		in, out := &in.ClusterHost, &out.ClusterHost
		*out = new(int32)
		**out = **in
	}
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(int32)
		**out = **in
	}
	if in.Datastore != nil {
		in, out := &in.Datastore, &out.Datastore
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VCenterSimulatorModel.
func (in *VCenterSimulatorModel) DeepCopy() *VCenterSimulatorModel {
	if in == nil {
		return nil
	}
	out := new(VCenterSimulatorModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VCenterSimulatorSpec) DeepCopyInto(out *VCenterSimulatorSpec) {
	*out = *in
	if in.Model != nil {
		in, out := &in.Model, &out.Model
		*out = new(VCenterSimulatorModel)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VCenterSimulatorSpec.
func (in *VCenterSimulatorSpec) DeepCopy() *VCenterSimulatorSpec {
	if in == nil {
		return nil
	}
	out := new(VCenterSimulatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VCenterSimulatorStatus) DeepCopyInto(out *VCenterSimulatorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VCenterSimulatorStatus.
func (in *VCenterSimulatorStatus) DeepCopy() *VCenterSimulatorStatus {
	if in == nil {
		return nil
	}
	out := new(VCenterSimulatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VCenterSpec) DeepCopyInto(out *VCenterSpec) {
	*out = *in
	in.ContentLibrary.DeepCopyInto(&out.ContentLibrary)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VCenterSpec.
func (in *VCenterSpec) DeepCopy() *VCenterSpec {
	if in == nil {
		return nil
	}
	out := new(VCenterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMOperatorDependencies) DeepCopyInto(out *VMOperatorDependencies) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMOperatorDependencies.
func (in *VMOperatorDependencies) DeepCopy() *VMOperatorDependencies {
	if in == nil {
		return nil
	}
	out := new(VMOperatorDependencies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMOperatorDependencies) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMOperatorDependenciesList) DeepCopyInto(out *VMOperatorDependenciesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VCenterSimulator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMOperatorDependenciesList.
func (in *VMOperatorDependenciesList) DeepCopy() *VMOperatorDependenciesList {
	if in == nil {
		return nil
	}
	out := new(VMOperatorDependenciesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMOperatorDependenciesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMOperatorDependenciesSpec) DeepCopyInto(out *VMOperatorDependenciesSpec) {
	*out = *in
	if in.OperatorRef != nil {
		in, out := &in.OperatorRef, &out.OperatorRef
		*out = new(VMOperatorRef)
		**out = **in
	}
	if in.VCenter != nil {
		in, out := &in.VCenter, &out.VCenter
		*out = new(VCenterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VCenterSimulatorRef != nil {
		in, out := &in.VCenterSimulatorRef, &out.VCenterSimulatorRef
		*out = new(NamespacedRef)
		**out = **in
	}
	if in.StorageClasses != nil {
		in, out := &in.StorageClasses, &out.StorageClasses
		*out = make([]StorageClass, len(*in))
		copy(*out, *in)
	}
	if in.VirtualMachineClasses != nil {
		in, out := &in.VirtualMachineClasses, &out.VirtualMachineClasses
		*out = make([]VirtualMachineClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMOperatorDependenciesSpec.
func (in *VMOperatorDependenciesSpec) DeepCopy() *VMOperatorDependenciesSpec {
	if in == nil {
		return nil
	}
	out := new(VMOperatorDependenciesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMOperatorDependenciesStatus) DeepCopyInto(out *VMOperatorDependenciesStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMOperatorDependenciesStatus.
func (in *VMOperatorDependenciesStatus) DeepCopy() *VMOperatorDependenciesStatus {
	if in == nil {
		return nil
	}
	out := new(VMOperatorDependenciesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMOperatorRef) DeepCopyInto(out *VMOperatorRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMOperatorRef.
func (in *VMOperatorRef) DeepCopy() *VMOperatorRef {
	if in == nil {
		return nil
	}
	out := new(VMOperatorRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineClass) DeepCopyInto(out *VirtualMachineClass) {
	*out = *in
	out.Memory = in.Memory.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineClass.
func (in *VirtualMachineClass) DeepCopy() *VirtualMachineClass {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineClass)
	in.DeepCopyInto(out)
	return out
}