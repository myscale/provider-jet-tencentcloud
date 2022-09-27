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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KubeConfig *string `json:"kubeConfig,omitempty" tf:"kube_config,omitempty"`
}

type ClusterParameters struct {

	// Description of EKS cluster.
	// +kubebuilder:validation:Optional
	ClusterDesc *string `json:"clusterDesc,omitempty" tf:"cluster_desc,omitempty"`

	// Name of EKS cluster.
	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// List of cluster custom DNS Server info.
	// +kubebuilder:validation:Optional
	DNSServers []DNSServersParameters `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Indicates whether to enable dns in user cluster, default value is `true`.
	// +kubebuilder:validation:Optional
	EnableVPCCoreDNS *bool `json:"enableVpcCoreDns,omitempty" tf:"enable_vpc_core_dns,omitempty"`

	// Extend parameters.
	// +kubebuilder:validation:Optional
	ExtraParam map[string]*string `json:"extraParam,omitempty" tf:"extra_param,omitempty"`

	// Cluster internal access LoadBalancer info.
	// +kubebuilder:validation:Optional
	InternalLB []InternalLBParameters `json:"internalLb,omitempty" tf:"internal_lb,omitempty"`

	// Kubernetes version of EKS cluster.
	// +kubebuilder:validation:Required
	K8SVersion *string `json:"k8sVersion" tf:"k8s_version,omitempty"`

	// Delete CBS after EKS cluster remove.
	// +kubebuilder:validation:Optional
	NeedDeleteCbs *bool `json:"needDeleteCbs,omitempty" tf:"need_delete_cbs,omitempty"`

	// Cluster public access LoadBalancer info.
	// +kubebuilder:validation:Optional
	PublicLB []PublicLBParameters `json:"publicLb,omitempty" tf:"public_lb,omitempty"`

	// Subnet id of service.
	// +kubebuilder:validation:Optional
	ServiceSubnetID *string `json:"serviceSubnetId,omitempty" tf:"service_subnet_id,omitempty"`

	// Subnet Ids for EKS cluster.
	// +kubebuilder:validation:Required
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`

	// Tags of EKS cluster.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Vpc Id of EKS cluster.
	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`
}

type DNSServersObservation struct {
}

type DNSServersParameters struct {

	// DNS Server domain. Empty indicates all domain.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// List of DNS Server IP address, pattern: "ip[:port]".
	// +kubebuilder:validation:Optional
	Servers []*string `json:"servers,omitempty" tf:"servers,omitempty"`
}

type InternalLBObservation struct {
}

type InternalLBParameters struct {

	// Indicates weather the internal access LB enabled.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// ID of subnet which related to Internal LB.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type PublicLBObservation struct {
	SecurityGroup *string `json:"securityGroup,omitempty" tf:"security_group,omitempty"`
}

type PublicLBParameters struct {

	// List of CIDRs which allowed to access.
	// +kubebuilder:validation:Optional
	AllowFromCidrs []*string `json:"allowFromCidrs,omitempty" tf:"allow_from_cidrs,omitempty"`

	// Indicates weather the public access LB enabled.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Extra param text json.
	// +kubebuilder:validation:Optional
	ExtraParam *string `json:"extraParam,omitempty" tf:"extra_param,omitempty"`

	// List of security allow IP or CIDRs, default deny all.
	// +kubebuilder:validation:Optional
	SecurityPolicies []*string `json:"securityPolicies,omitempty" tf:"security_policies,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}