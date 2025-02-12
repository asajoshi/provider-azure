/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IPGroupObservation struct {

	// A firewall_ids block as defined below.
	FirewallIds []*string `json:"firewallIds,omitempty" tf:"firewall_ids,omitempty"`

	// A firewall_policy_ids block as defined below.
	FirewallPolicyIds []*string `json:"firewallPolicyIds,omitempty" tf:"firewall_policy_ids,omitempty"`

	// The ID of the IP group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IPGroupParameters struct {

	// A list of CIDRs or IP addresses.
	// +kubebuilder:validation:Optional
	Cidrs []*string `json:"cidrs,omitempty" tf:"cidrs,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the resource group in which to create the IP group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// IPGroupSpec defines the desired state of IPGroup
type IPGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPGroupParameters `json:"forProvider"`
}

// IPGroupStatus defines the observed state of IPGroup.
type IPGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IPGroup is the Schema for the IPGroups API. Manages an IP group which contains a list of CIDRs and/or IP addresses.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IPGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IPGroupSpec   `json:"spec"`
	Status            IPGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPGroupList contains a list of IPGroups
type IPGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPGroup `json:"items"`
}

// Repository type metadata.
var (
	IPGroup_Kind             = "IPGroup"
	IPGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPGroup_Kind}.String()
	IPGroup_KindAPIVersion   = IPGroup_Kind + "." + CRDGroupVersion.String()
	IPGroup_GroupVersionKind = CRDGroupVersion.WithKind(IPGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&IPGroup{}, &IPGroupList{})
}
