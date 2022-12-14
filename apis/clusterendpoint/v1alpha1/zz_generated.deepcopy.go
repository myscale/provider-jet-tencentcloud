//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpoint) DeepCopyInto(out *ClusterEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpoint.
func (in *ClusterEndpoint) DeepCopy() *ClusterEndpoint {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpointList) DeepCopyInto(out *ClusterEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpointList.
func (in *ClusterEndpointList) DeepCopy() *ClusterEndpointList {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpointObservation) DeepCopyInto(out *ClusterEndpointObservation) {
	*out = *in
	if in.CertificationAuthority != nil {
		in, out := &in.CertificationAuthority, &out.CertificationAuthority
		*out = new(string)
		**out = **in
	}
	if in.ClusterDeployType != nil {
		in, out := &in.ClusterDeployType, &out.ClusterDeployType
		*out = new(string)
		**out = **in
	}
	if in.ClusterExternalEndpoint != nil {
		in, out := &in.ClusterExternalEndpoint, &out.ClusterExternalEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PgwEndpoint != nil {
		in, out := &in.PgwEndpoint, &out.PgwEndpoint
		*out = new(string)
		**out = **in
	}
	if in.UserName != nil {
		in, out := &in.UserName, &out.UserName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpointObservation.
func (in *ClusterEndpointObservation) DeepCopy() *ClusterEndpointObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpointParameters) DeepCopyInto(out *ClusterEndpointParameters) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterIdRefs != nil {
		in, out := &in.ClusterIdRefs, &out.ClusterIdRefs
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ClusterIdSelector != nil {
		in, out := &in.ClusterIdSelector, &out.ClusterIdSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterInternet != nil {
		in, out := &in.ClusterInternet, &out.ClusterInternet
		*out = new(bool)
		**out = **in
	}
	if in.ClusterInternetSecurityGroup != nil {
		in, out := &in.ClusterInternetSecurityGroup, &out.ClusterInternetSecurityGroup
		*out = new(string)
		**out = **in
	}
	if in.ClusterIntranet != nil {
		in, out := &in.ClusterIntranet, &out.ClusterIntranet
		*out = new(bool)
		**out = **in
	}
	if in.ClusterIntranetSubnetID != nil {
		in, out := &in.ClusterIntranetSubnetID, &out.ClusterIntranetSubnetID
		*out = new(string)
		**out = **in
	}
	if in.ClusterIntranetSubnetIdRefs != nil {
		in, out := &in.ClusterIntranetSubnetIdRefs, &out.ClusterIntranetSubnetIdRefs
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ClusterIntranetSubnetIdSelector != nil {
		in, out := &in.ClusterIntranetSubnetIdSelector, &out.ClusterIntranetSubnetIdSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagedClusterInternetSecurityPolicies != nil {
		in, out := &in.ManagedClusterInternetSecurityPolicies, &out.ManagedClusterInternetSecurityPolicies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpointParameters.
func (in *ClusterEndpointParameters) DeepCopy() *ClusterEndpointParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpointSpec) DeepCopyInto(out *ClusterEndpointSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpointSpec.
func (in *ClusterEndpointSpec) DeepCopy() *ClusterEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpointStatus) DeepCopyInto(out *ClusterEndpointStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpointStatus.
func (in *ClusterEndpointStatus) DeepCopy() *ClusterEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpointStatus)
	in.DeepCopyInto(out)
	return out
}
