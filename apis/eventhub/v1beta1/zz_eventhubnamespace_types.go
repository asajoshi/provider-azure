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

type EventHubNamespaceObservation struct {

	// The EventHub Namespace ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`
}

type EventHubNamespaceParameters struct {

	// Is Auto Inflate enabled for the EventHub Namespace?
	// +kubebuilder:validation:Optional
	AutoInflateEnabled *bool `json:"autoInflateEnabled,omitempty" tf:"auto_inflate_enabled,omitempty"`

	// Specifies the Capacity / Throughput Units for a Standard SKU namespace. Default capacity has a maximum of 2, but can be increased in blocks of 2 on a committed purchase basis. Defaults to 1.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// Specifies the ID of the EventHub Dedicated Cluster where this Namespace should created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DedicatedClusterID *string `json:"dedicatedClusterId,omitempty" tf:"dedicated_cluster_id,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Is SAS authentication enabled for the EventHub Namespace? Defaults to true.
	// +kubebuilder:validation:Optional
	LocalAuthenticationEnabled *bool `json:"localAuthenticationEnabled,omitempty" tf:"local_authentication_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Specifies the maximum number of throughput units when Auto Inflate is Enabled. Valid values range from 1 - 20.
	// +kubebuilder:validation:Optional
	MaximumThroughputUnits *float64 `json:"maximumThroughputUnits,omitempty" tf:"maximum_throughput_units,omitempty"`

	// The minimum supported TLS version for this EventHub Namespace. Valid values are: 1.0, 1.1 and 1.2. The current default minimum TLS version is 1.2.
	// +kubebuilder:validation:Optional
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// A network_rulesets block as defined below.
	// +kubebuilder:validation:Optional
	NetworkRulesets []NetworkRulesetsParameters `json:"networkRulesets,omitempty" tf:"network_rulesets,omitempty"`

	// Is public network access enabled for the EventHub Namespace? Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Defines which tier to use. Valid options are Basic, Standard, and Premium. Please note that setting this field to Premium will force the creation of a new resource.
	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies if the EventHub Namespace should be Zone Redundant (created across Availability Zones). Changing this forces a new resource to be created. Defaults to false.
	// +kubebuilder:validation:Optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type IPRuleObservation struct {
}

type IPRuleParameters struct {

	// The action to take when the rule is matched. Possible values are Allow.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action"`

	// The IP mask to match on.
	// +kubebuilder:validation:Optional
	IPMask *string `json:"ipMask,omitempty" tf:"ip_mask"`
}

type IdentityObservation struct {

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this EventHub namespace.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Event Hub Namespace. Possible values are SystemAssigned or UserAssigned.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type NetworkRulesetsObservation struct {
}

type NetworkRulesetsParameters struct {

	// The default action to take when a rule is not matched. Possible values are Allow and Deny.
	// +kubebuilder:validation:Optional
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action"`

	// One or more ip_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	IPRule []IPRuleParameters `json:"ipRule,omitempty" tf:"ip_rule"`

	// Is public network access enabled for the EventHub Namespace? Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`

	// Whether Trusted Microsoft Services are allowed to bypass firewall.
	// +kubebuilder:validation:Optional
	TrustedServiceAccessEnabled *bool `json:"trustedServiceAccessEnabled,omitempty" tf:"trusted_service_access_enabled"`

	// One or more virtual_network_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	VirtualNetworkRule []VirtualNetworkRuleParameters `json:"virtualNetworkRule,omitempty" tf:"virtual_network_rule"`
}

type VirtualNetworkRuleObservation struct {
}

type VirtualNetworkRuleParameters struct {

	// Are missing virtual network service endpoints ignored?
	// +kubebuilder:validation:Optional
	IgnoreMissingVirtualNetworkServiceEndpoint *bool `json:"ignoreMissingVirtualNetworkServiceEndpoint,omitempty" tf:"ignore_missing_virtual_network_service_endpoint"`

	// The id of the subnet to match on.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// EventHubNamespaceSpec defines the desired state of EventHubNamespace
type EventHubNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventHubNamespaceParameters `json:"forProvider"`
}

// EventHubNamespaceStatus defines the observed state of EventHubNamespace.
type EventHubNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventHubNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventHubNamespace is the Schema for the EventHubNamespaces API. Manages an EventHub Namespace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EventHubNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventHubNamespaceSpec   `json:"spec"`
	Status            EventHubNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventHubNamespaceList contains a list of EventHubNamespaces
type EventHubNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventHubNamespace `json:"items"`
}

// Repository type metadata.
var (
	EventHubNamespace_Kind             = "EventHubNamespace"
	EventHubNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventHubNamespace_Kind}.String()
	EventHubNamespace_KindAPIVersion   = EventHubNamespace_Kind + "." + CRDGroupVersion.String()
	EventHubNamespace_GroupVersionKind = CRDGroupVersion.WithKind(EventHubNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&EventHubNamespace{}, &EventHubNamespaceList{})
}
