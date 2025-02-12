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

type BackupPolicyVMWorkloadObservation struct {

	// The ID of the Azure VM Workload Backup Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackupPolicyVMWorkloadParameters struct {

	// One or more protection_policy blocks as defined below.
	// +kubebuilder:validation:Required
	ProtectionPolicy []ProtectionPolicyParameters `json:"protectionPolicy" tf:"protection_policy,omitempty"`

	// The name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	// +kubebuilder:validation:Optional
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// The name of the resource group in which to create the VM Workload Backup Policy. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A settings block as defined below.
	// +kubebuilder:validation:Required
	Settings []SettingsParameters `json:"settings" tf:"settings,omitempty"`

	// The VM Workload type for the Backup Policy. Possible values are SQLDataBase and SAPHanaDatabase. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	WorkloadType *string `json:"workloadType" tf:"workload_type,omitempty"`
}

type ProtectionPolicyBackupObservation struct {
}

type ProtectionPolicyBackupParameters struct {

	// The backup frequency for the VM Workload Backup Policy. Possible values are Daily and Weekly.
	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// The backup frequency in minutes for the VM Workload Backup Policy. Possible values are 15, 30, 60, 120, 240, 480, 720 and 1440.
	// +kubebuilder:validation:Optional
	FrequencyInMinutes *float64 `json:"frequencyInMinutes,omitempty" tf:"frequency_in_minutes,omitempty"`

	// The time of day to perform the backup in 24hour format.
	// +kubebuilder:validation:Optional
	Time *string `json:"time,omitempty" tf:"time,omitempty"`

	// The weekday backups to retain. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	// +kubebuilder:validation:Optional
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type ProtectionPolicyObservation struct {
}

type ProtectionPolicyParameters struct {

	// A backup block as defined below.
	// +kubebuilder:validation:Required
	Backup []ProtectionPolicyBackupParameters `json:"backup" tf:"backup,omitempty"`

	// The type of the VM Workload Backup Policy. Possible values are Differential, Full, Incremental and Log.
	// +kubebuilder:validation:Required
	PolicyType *string `json:"policyType" tf:"policy_type,omitempty"`

	// A retention_daily block as defined below.
	// +kubebuilder:validation:Optional
	RetentionDaily []ProtectionPolicyRetentionDailyParameters `json:"retentionDaily,omitempty" tf:"retention_daily,omitempty"`

	// A retention_monthly block as defined below.
	// +kubebuilder:validation:Optional
	RetentionMonthly []ProtectionPolicyRetentionMonthlyParameters `json:"retentionMonthly,omitempty" tf:"retention_monthly,omitempty"`

	// A retention_weekly block as defined below.
	// +kubebuilder:validation:Optional
	RetentionWeekly []ProtectionPolicyRetentionWeeklyParameters `json:"retentionWeekly,omitempty" tf:"retention_weekly,omitempty"`

	// A retention_yearly block as defined below.
	// +kubebuilder:validation:Optional
	RetentionYearly []ProtectionPolicyRetentionYearlyParameters `json:"retentionYearly,omitempty" tf:"retention_yearly,omitempty"`

	// A simple_retention block as defined below.
	// +kubebuilder:validation:Optional
	SimpleRetention []SimpleRetentionParameters `json:"simpleRetention,omitempty" tf:"simple_retention,omitempty"`
}

type ProtectionPolicyRetentionDailyObservation struct {
}

type ProtectionPolicyRetentionDailyParameters struct {

	// The count that is used to count retention duration with duration type Days. Possible values are between 7 and 35.
	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`
}

type ProtectionPolicyRetentionMonthlyObservation struct {
}

type ProtectionPolicyRetentionMonthlyParameters struct {

	// The count that is used to count retention duration with duration type Days. Possible values are between 7 and 35.
	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`

	// The retention schedule format type for yearly retention policy. Possible values are Daily and Weekly.
	// +kubebuilder:validation:Required
	FormatType *string `json:"formatType" tf:"format_type,omitempty"`

	// The monthday backups to retain. Possible values are between 0 and 28.
	// +kubebuilder:validation:Optional
	Monthdays []*float64 `json:"monthdays,omitempty" tf:"monthdays,omitempty"`

	// The weekday backups to retain. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	// +kubebuilder:validation:Optional
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Possible values are First, Second, Third, Fourth, Last.
	// +kubebuilder:validation:Optional
	Weeks []*string `json:"weeks,omitempty" tf:"weeks,omitempty"`
}

type ProtectionPolicyRetentionWeeklyObservation struct {
}

type ProtectionPolicyRetentionWeeklyParameters struct {

	// The count that is used to count retention duration with duration type Days. Possible values are between 7 and 35.
	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`

	// The weekday backups to retain. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	// +kubebuilder:validation:Required
	Weekdays []*string `json:"weekdays" tf:"weekdays,omitempty"`
}

type ProtectionPolicyRetentionYearlyObservation struct {
}

type ProtectionPolicyRetentionYearlyParameters struct {

	// The count that is used to count retention duration with duration type Days. Possible values are between 7 and 35.
	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`

	// The retention schedule format type for yearly retention policy. Possible values are Daily and Weekly.
	// +kubebuilder:validation:Required
	FormatType *string `json:"formatType" tf:"format_type,omitempty"`

	// The monthday backups to retain. Possible values are between 0 and 28.
	// +kubebuilder:validation:Optional
	Monthdays []*float64 `json:"monthdays,omitempty" tf:"monthdays,omitempty"`

	// The months of the year to retain backups of. Possible values are January, February, March, April, May, June, July, August, September, October, November and December.
	// +kubebuilder:validation:Required
	Months []*string `json:"months" tf:"months,omitempty"`

	// The weekday backups to retain. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	// +kubebuilder:validation:Optional
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Possible values are First, Second, Third, Fourth, Last.
	// +kubebuilder:validation:Optional
	Weeks []*string `json:"weeks,omitempty" tf:"weeks,omitempty"`
}

type SettingsObservation struct {
}

type SettingsParameters struct {

	// The compression setting for the VM Workload Backup Policy. Defaults to false.
	// +kubebuilder:validation:Optional
	CompressionEnabled *bool `json:"compressionEnabled,omitempty" tf:"compression_enabled,omitempty"`

	// The timezone for the VM Workload Backup Policy. The possible values are defined here.
	// +kubebuilder:validation:Required
	TimeZone *string `json:"timeZone" tf:"time_zone,omitempty"`
}

type SimpleRetentionObservation struct {
}

type SimpleRetentionParameters struct {

	// The count that is used to count retention duration with duration type Days. Possible values are between 7 and 35.
	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`
}

// BackupPolicyVMWorkloadSpec defines the desired state of BackupPolicyVMWorkload
type BackupPolicyVMWorkloadSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupPolicyVMWorkloadParameters `json:"forProvider"`
}

// BackupPolicyVMWorkloadStatus defines the observed state of BackupPolicyVMWorkload.
type BackupPolicyVMWorkloadStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupPolicyVMWorkloadObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyVMWorkload is the Schema for the BackupPolicyVMWorkloads API. Manages an Azure VM Workload Backup Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupPolicyVMWorkload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupPolicyVMWorkloadSpec   `json:"spec"`
	Status            BackupPolicyVMWorkloadStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyVMWorkloadList contains a list of BackupPolicyVMWorkloads
type BackupPolicyVMWorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPolicyVMWorkload `json:"items"`
}

// Repository type metadata.
var (
	BackupPolicyVMWorkload_Kind             = "BackupPolicyVMWorkload"
	BackupPolicyVMWorkload_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupPolicyVMWorkload_Kind}.String()
	BackupPolicyVMWorkload_KindAPIVersion   = BackupPolicyVMWorkload_Kind + "." + CRDGroupVersion.String()
	BackupPolicyVMWorkload_GroupVersionKind = CRDGroupVersion.WithKind(BackupPolicyVMWorkload_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupPolicyVMWorkload{}, &BackupPolicyVMWorkloadList{})
}
