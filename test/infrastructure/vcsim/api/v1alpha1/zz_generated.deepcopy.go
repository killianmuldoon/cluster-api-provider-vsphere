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
	in.Cluster.DeepCopyInto(&out.Cluster)
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
