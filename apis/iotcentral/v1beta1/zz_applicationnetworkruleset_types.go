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

type ApplicationNetworkRuleSetObservation struct {

	// The ID of the IoT Central Application Network Rule Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ApplicationNetworkRuleSetParameters struct {

	// Whether these IP Rules apply for device connectivity to IoT Hub and Device Provisioning Service associated with this IoT Central Application. Possible values are true, false. Defaults to true
	// +kubebuilder:validation:Optional
	ApplyToDevice *bool `json:"applyToDevice,omitempty" tf:"apply_to_device,omitempty"`

	// Specifies the default action for the IoT Central Application Network Rule Set. Possible values are Allow and Deny. Defaults to Deny.
	// +kubebuilder:validation:Optional
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// One or more ip_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	IPRule []IPRuleParameters `json:"ipRule,omitempty" tf:"ip_rule,omitempty"`

	// The ID of the IoT Central Application. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/iotcentral/v1beta1.Application
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IotcentralApplicationID *string `json:"iotcentralApplicationId,omitempty" tf:"iotcentral_application_id,omitempty"`

	// Reference to a Application in iotcentral to populate iotcentralApplicationId.
	// +kubebuilder:validation:Optional
	IotcentralApplicationIDRef *v1.Reference `json:"iotcentralApplicationIdRef,omitempty" tf:"-"`

	// Selector for a Application in iotcentral to populate iotcentralApplicationId.
	// +kubebuilder:validation:Optional
	IotcentralApplicationIDSelector *v1.Selector `json:"iotcentralApplicationIdSelector,omitempty" tf:"-"`
}

type IPRuleObservation struct {
}

type IPRuleParameters struct {

	// The IP address range in CIDR notation for the IP Rule.
	// +kubebuilder:validation:Required
	IPMask *string `json:"ipMask" tf:"ip_mask,omitempty"`

	// The name of the IP Rule
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// ApplicationNetworkRuleSetSpec defines the desired state of ApplicationNetworkRuleSet
type ApplicationNetworkRuleSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationNetworkRuleSetParameters `json:"forProvider"`
}

// ApplicationNetworkRuleSetStatus defines the observed state of ApplicationNetworkRuleSet.
type ApplicationNetworkRuleSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationNetworkRuleSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationNetworkRuleSet is the Schema for the ApplicationNetworkRuleSets API. Manages an IoT Central Application Network Rule Set.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApplicationNetworkRuleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationNetworkRuleSetSpec   `json:"spec"`
	Status            ApplicationNetworkRuleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationNetworkRuleSetList contains a list of ApplicationNetworkRuleSets
type ApplicationNetworkRuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationNetworkRuleSet `json:"items"`
}

// Repository type metadata.
var (
	ApplicationNetworkRuleSet_Kind             = "ApplicationNetworkRuleSet"
	ApplicationNetworkRuleSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationNetworkRuleSet_Kind}.String()
	ApplicationNetworkRuleSet_KindAPIVersion   = ApplicationNetworkRuleSet_Kind + "." + CRDGroupVersion.String()
	ApplicationNetworkRuleSet_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationNetworkRuleSet_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationNetworkRuleSet{}, &ApplicationNetworkRuleSetList{})
}
