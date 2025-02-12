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

type ApplicationInsightsSmartDetectionRuleObservation struct {

	// The ID of the Application Insights Smart Detection Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ApplicationInsightsSmartDetectionRuleParameters struct {

	// Specifies a list of additional recipients that will be sent emails on this Application Insights Smart Detection Rule.
	// +kubebuilder:validation:Optional
	AdditionalEmailRecipients []*string `json:"additionalEmailRecipients,omitempty" tf:"additional_email_recipients,omitempty"`

	// The ID of the Application Insights component on which the Smart Detection Rule operates. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ApplicationInsightsID *string `json:"applicationInsightsId,omitempty" tf:"application_insights_id,omitempty"`

	// Reference to a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDRef *v1.Reference `json:"applicationInsightsIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDSelector *v1.Selector `json:"applicationInsightsIdSelector,omitempty" tf:"-"`

	// Is the Application Insights Smart Detection Rule enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the name of the Application Insights Smart Detection Rule. Valid values include Slow page load time, Slow server response time, Long dependency duration, Degradation in server response time, Degradation in dependency duration, Degradation in trace severity ratio, Abnormal rise in exception volume, Potential memory leak detected, Potential security issue detected and Abnormal rise in daily data volume, Long dependency duration, Degradation in server response time, Degradation in dependency duration, Degradation in trace severity ratio, Abnormal rise in exception volume, Potential memory leak detected, Potential security issue detected, Abnormal rise in daily data volume. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Do emails get sent to subscription owners? Defaults to true.
	// +kubebuilder:validation:Optional
	SendEmailsToSubscriptionOwners *bool `json:"sendEmailsToSubscriptionOwners,omitempty" tf:"send_emails_to_subscription_owners,omitempty"`
}

// ApplicationInsightsSmartDetectionRuleSpec defines the desired state of ApplicationInsightsSmartDetectionRule
type ApplicationInsightsSmartDetectionRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationInsightsSmartDetectionRuleParameters `json:"forProvider"`
}

// ApplicationInsightsSmartDetectionRuleStatus defines the observed state of ApplicationInsightsSmartDetectionRule.
type ApplicationInsightsSmartDetectionRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationInsightsSmartDetectionRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsSmartDetectionRule is the Schema for the ApplicationInsightsSmartDetectionRules API. Manages an Application Insights Smart Detection Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApplicationInsightsSmartDetectionRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationInsightsSmartDetectionRuleSpec   `json:"spec"`
	Status            ApplicationInsightsSmartDetectionRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsSmartDetectionRuleList contains a list of ApplicationInsightsSmartDetectionRules
type ApplicationInsightsSmartDetectionRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationInsightsSmartDetectionRule `json:"items"`
}

// Repository type metadata.
var (
	ApplicationInsightsSmartDetectionRule_Kind             = "ApplicationInsightsSmartDetectionRule"
	ApplicationInsightsSmartDetectionRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationInsightsSmartDetectionRule_Kind}.String()
	ApplicationInsightsSmartDetectionRule_KindAPIVersion   = ApplicationInsightsSmartDetectionRule_Kind + "." + CRDGroupVersion.String()
	ApplicationInsightsSmartDetectionRule_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationInsightsSmartDetectionRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationInsightsSmartDetectionRule{}, &ApplicationInsightsSmartDetectionRuleList{})
}
