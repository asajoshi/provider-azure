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

type PrivateDNSResolverObservation struct {

	// The ID of the DNS Resolver.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivateDNSResolverParameters struct {

	// Specifies the Azure Region where the Private DNS Resolver should exist. Changing this forces a new Private DNS Resolver to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Specifies the name of the Resource Group where the Private DNS Resolver should exist. Changing this forces a new Private DNS Resolver to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Private DNS Resolver.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the Virtual Network that is linked to the Private DNS Resolver. Changing this forces a new Private DNS Resolver to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualNetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`

	// Reference to a VirtualNetwork in network to populate virtualNetworkId.
	// +kubebuilder:validation:Optional
	VirtualNetworkIDRef *v1.Reference `json:"virtualNetworkIdRef,omitempty" tf:"-"`

	// Selector for a VirtualNetwork in network to populate virtualNetworkId.
	// +kubebuilder:validation:Optional
	VirtualNetworkIDSelector *v1.Selector `json:"virtualNetworkIdSelector,omitempty" tf:"-"`
}

// PrivateDNSResolverSpec defines the desired state of PrivateDNSResolver
type PrivateDNSResolverSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSResolverParameters `json:"forProvider"`
}

// PrivateDNSResolverStatus defines the observed state of PrivateDNSResolver.
type PrivateDNSResolverStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSResolverObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSResolver is the Schema for the PrivateDNSResolvers API. Manages a Private DNS Resolver.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateDNSResolver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDNSResolverSpec   `json:"spec"`
	Status            PrivateDNSResolverStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSResolverList contains a list of PrivateDNSResolvers
type PrivateDNSResolverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSResolver `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSResolver_Kind             = "PrivateDNSResolver"
	PrivateDNSResolver_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSResolver_Kind}.String()
	PrivateDNSResolver_KindAPIVersion   = PrivateDNSResolver_Kind + "." + CRDGroupVersion.String()
	PrivateDNSResolver_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSResolver_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSResolver{}, &PrivateDNSResolverList{})
}
