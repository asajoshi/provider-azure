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

type AccountObservation struct {

	// The endpoint used to connect to the CosmosDB account.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Specifies a geo_location resource, used to define where data should be replicated with the failover_priority 0 specifying the primary location. Value is a geo_location block as defined below.
	// +kubebuilder:validation:Required
	GeoLocation []GeoLocationObservation `json:"geoLocation,omitempty" tf:"geo_location,omitempty"`

	// The ID of the virtual network subnet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// A list of read endpoints available for this CosmosDB account.
	ReadEndpoints []*string `json:"readEndpoints,omitempty" tf:"read_endpoints,omitempty"`

	// A list of write endpoints available for this CosmosDB account.
	WriteEndpoints []*string `json:"writeEndpoints,omitempty" tf:"write_endpoints,omitempty"`
}

type AccountParameters struct {

	// Is write operations on metadata resources (databases, containers, throughput) via account keys enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	AccessKeyMetadataWritesEnabled *bool `json:"accessKeyMetadataWritesEnabled,omitempty" tf:"access_key_metadata_writes_enabled,omitempty"`

	// An analytical_storage block as defined below.
	// +kubebuilder:validation:Optional
	AnalyticalStorage []AnalyticalStorageParameters `json:"analyticalStorage,omitempty" tf:"analytical_storage,omitempty"`

	// Enable Analytical Storage option for this Cosmos DB account. Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AnalyticalStorageEnabled *bool `json:"analyticalStorageEnabled,omitempty" tf:"analytical_storage_enabled,omitempty"`

	// A backup block as defined below.
	// +kubebuilder:validation:Optional
	Backup []BackupParameters `json:"backup,omitempty" tf:"backup,omitempty"`

	// The capabilities which should be enabled for this Cosmos DB account. Value is a capabilities block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Capabilities []CapabilitiesParameters `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// A capacity block as defined below.
	// +kubebuilder:validation:Optional
	Capacity []CapacityParameters `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// Specifies a consistency_policy resource, used to define the consistency policy for this CosmosDB account.
	// +kubebuilder:validation:Required
	ConsistencyPolicy []ConsistencyPolicyParameters `json:"consistencyPolicy" tf:"consistency_policy,omitempty"`

	// A cors_rule block as defined below.
	// +kubebuilder:validation:Optional
	CorsRule []CorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// The creation mode for the CosmosDB Account. Possible values are Default and Restore. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// The default identity for accessing Key Vault. Possible values are FirstPartyIdentity, SystemAssignedIdentity or start with UserAssignedIdentity. Defaults to FirstPartyIdentity.
	// +kubebuilder:validation:Optional
	DefaultIdentityType *string `json:"defaultIdentityType,omitempty" tf:"default_identity_type,omitempty"`

	// Enable automatic fail over for this Cosmos DB account.
	// +kubebuilder:validation:Optional
	EnableAutomaticFailover *bool `json:"enableAutomaticFailover,omitempty" tf:"enable_automatic_failover,omitempty"`

	// Enable Free Tier pricing option for this Cosmos DB account. Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	EnableFreeTier *bool `json:"enableFreeTier,omitempty" tf:"enable_free_tier,omitempty"`

	// Enable multiple write locations for this Cosmos DB account.
	// +kubebuilder:validation:Optional
	EnableMultipleWriteLocations *bool `json:"enableMultipleWriteLocations,omitempty" tf:"enable_multiple_write_locations,omitempty"`

	// Specifies a geo_location resource, used to define where data should be replicated with the failover_priority 0 specifying the primary location. Value is a geo_location block as defined below.
	// +kubebuilder:validation:Required
	GeoLocation []GeoLocationParameters `json:"geoLocation" tf:"geo_location,omitempty"`

	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	// +kubebuilder:validation:Optional
	IPRangeFilter *string `json:"ipRangeFilter,omitempty" tf:"ip_range_filter,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Enables virtual network filtering for this Cosmos DB account.
	// +kubebuilder:validation:Optional
	IsVirtualNetworkFilterEnabled *bool `json:"isVirtualNetworkFilterEnabled,omitempty" tf:"is_virtual_network_filter_enabled,omitempty"`

	// A versionless Key Vault Key ID for CMK encryption. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Specifies the Kind of CosmosDB to create - possible values are GlobalDocumentDB and MongoDB. Defaults to GlobalDocumentDB. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Disable local authentication and ensure only MSI and AAD can be used exclusively for authentication. Defaults to false. Can be set only when using the SQL API.
	// +kubebuilder:validation:Optional
	LocalAuthenticationDisabled *bool `json:"localAuthenticationDisabled,omitempty" tf:"local_authentication_disabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The Server Version of a MongoDB account. Possible values are 4.2, 4.0, 3.6, and 3.2.
	// +kubebuilder:validation:Optional
	MongoServerVersion *string `json:"mongoServerVersion,omitempty" tf:"mongo_server_version,omitempty"`

	// If Azure services can bypass ACLs. Defaults to false.
	// +kubebuilder:validation:Optional
	NetworkACLBypassForAzureServices *bool `json:"networkAclBypassForAzureServices,omitempty" tf:"network_acl_bypass_for_azure_services,omitempty"`

	// The list of resource Ids for Network Acl Bypass for this Cosmos DB account.
	// +kubebuilder:validation:Optional
	NetworkACLBypassIds []*string `json:"networkAclBypassIds,omitempty" tf:"network_acl_bypass_ids,omitempty"`

	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to Standard.
	// +kubebuilder:validation:Required
	OfferType *string `json:"offerType" tf:"offer_type,omitempty"`

	// Whether or not public network access is allowed for this CosmosDB account.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A restore block as defined below.
	// +kubebuilder:validation:Optional
	Restore []RestoreParameters `json:"restore,omitempty" tf:"restore,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a virtual_network_rules resource, used to define which subnets are allowed to access this CosmosDB account.
	// +kubebuilder:validation:Optional
	VirtualNetworkRule []VirtualNetworkRuleParameters `json:"virtualNetworkRule,omitempty" tf:"virtual_network_rule,omitempty"`
}

type AnalyticalStorageObservation struct {
}

type AnalyticalStorageParameters struct {

	// The schema type of the Analytical Storage for this Cosmos DB account. Possible values are FullFidelity and WellDefined.
	// +kubebuilder:validation:Required
	SchemaType *string `json:"schemaType" tf:"schema_type,omitempty"`
}

type BackupObservation struct {
}

type BackupParameters struct {

	// The interval in minutes between two backups. This is configurable only when type is Periodic. Possible values are between 60 and 1440.
	// +kubebuilder:validation:Optional
	IntervalInMinutes *float64 `json:"intervalInMinutes,omitempty" tf:"interval_in_minutes,omitempty"`

	// The time in hours that each backup is retained. This is configurable only when type is Periodic. Possible values are between 8 and 720.
	// +kubebuilder:validation:Optional
	RetentionInHours *float64 `json:"retentionInHours,omitempty" tf:"retention_in_hours,omitempty"`

	// The storage redundancy which is used to indicate type of backup residency. This is configurable only when type is Periodic. Possible values are Geo, Local and Zone.
	// +kubebuilder:validation:Optional
	StorageRedundancy *string `json:"storageRedundancy,omitempty" tf:"storage_redundancy,omitempty"`

	// The type of the backup. Possible values are Continuous and Periodic. Defaults to Periodic. Migration of Periodic to Continuous is one-way, changing Continuous to Periodic forces a new resource to be created.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type CapabilitiesObservation struct {
}

type CapabilitiesParameters struct {

	// Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type CapacityObservation struct {
}

type CapacityParameters struct {

	// The total throughput limit imposed on this Cosmos DB account (RU/s). Possible values are at least -1. -1 means no limit.
	// +kubebuilder:validation:Required
	TotalThroughputLimit *float64 `json:"totalThroughputLimit" tf:"total_throughput_limit,omitempty"`
}

type ConsistencyPolicyObservation struct {
}

type ConsistencyPolicyParameters struct {

	// The Consistency Level to use for this CosmosDB Account - can be either BoundedStaleness, Eventual, Session, Strong or ConsistentPrefix.
	// +kubebuilder:validation:Required
	ConsistencyLevel *string `json:"consistencyLevel" tf:"consistency_level,omitempty"`

	// When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400 (1 day). Defaults to 5. Required when consistency_level is set to BoundedStaleness.
	// +kubebuilder:validation:Optional
	MaxIntervalInSeconds *float64 `json:"maxIntervalInSeconds,omitempty" tf:"max_interval_in_seconds,omitempty"`

	// When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated. Accepted range for this value is 10 – 2147483647. Defaults to 100. Required when consistency_level is set to BoundedStaleness.
	// +kubebuilder:validation:Optional
	MaxStalenessPrefix *float64 `json:"maxStalenessPrefix,omitempty" tf:"max_staleness_prefix,omitempty"`
}

type CorsRuleObservation struct {
}

type CorsRuleParameters struct {

	// A list of headers that are allowed to be a part of the cross-origin request.
	// +kubebuilder:validation:Required
	AllowedHeaders []*string `json:"allowedHeaders" tf:"allowed_headers,omitempty"`

	// A list of HTTP headers that are allowed to be executed by the origin. Valid options are  DELETE, GET, HEAD, MERGE, POST, OPTIONS, PUT or PATCH.
	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// A list of origin domains that will be allowed by CORS.
	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// A list of response headers that are exposed to CORS clients.
	// +kubebuilder:validation:Required
	ExposedHeaders []*string `json:"exposedHeaders" tf:"exposed_headers,omitempty"`

	// The number of seconds the client should cache a preflight response.
	// +kubebuilder:validation:Required
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds" tf:"max_age_in_seconds,omitempty"`
}

type DatabaseObservation struct {
}

type DatabaseParameters struct {

	// A list of the collection names for the restore request. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CollectionNames []*string `json:"collectionNames,omitempty" tf:"collection_names,omitempty"`

	// Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type GeoLocationObservation struct {

	// The ID of the virtual network subnet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GeoLocationParameters struct {

	// The failover priority of the region. A failover priority of 0 indicates a write region. The maximum value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the regions in which the database account exists. Changing this causes the location to be re-provisioned and cannot be changed for the location with failover priority 0.
	// +kubebuilder:validation:Required
	FailoverPriority *float64 `json:"failoverPriority" tf:"failover_priority,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Should zone redundancy be enabled for this region? Defaults to false.
	// +kubebuilder:validation:Optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type IdentityObservation struct {

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// Specifies the type of Managed Service Identity that should be configured on this Cosmos Account. The only possible value is SystemAssigned.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RestoreObservation struct {
}

type RestoreParameters struct {

	// A database block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Database []DatabaseParameters `json:"database,omitempty" tf:"database,omitempty"`

	// The creation time of the database or the collection (Datetime Format RFC 3339). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	RestoreTimestampInUtc *string `json:"restoreTimestampInUtc" tf:"restore_timestamp_in_utc,omitempty"`

	// The resource ID of the restorable database account from which the restore has to be initiated. The example is /subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{restorableDatabaseAccountName}. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Account
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceCosmosDBAccountID *string `json:"sourceCosmosdbAccountId,omitempty" tf:"source_cosmosdb_account_id,omitempty"`

	// Reference to a Account to populate sourceCosmosdbAccountId.
	// +kubebuilder:validation:Optional
	SourceCosmosDBAccountIDRef *v1.Reference `json:"sourceCosmosdbAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account to populate sourceCosmosdbAccountId.
	// +kubebuilder:validation:Optional
	SourceCosmosDBAccountIDSelector *v1.Selector `json:"sourceCosmosdbAccountIdSelector,omitempty" tf:"-"`
}

type VirtualNetworkRuleObservation struct {
}

type VirtualNetworkRuleParameters struct {

	// The ID of the virtual network subnet.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// If set to true, the specified subnet will be added as a virtual network rule even if its CosmosDB service endpoint is not active. Defaults to false.
	// +kubebuilder:validation:Optional
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`
}

// AccountSpec defines the desired state of Account
type AccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountParameters `json:"forProvider"`
}

// AccountStatus defines the observed state of Account.
type AccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Account is the Schema for the Accounts API. Manages a CosmosDB (formally DocumentDB) Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountSpec   `json:"spec"`
	Status            AccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountList contains a list of Accounts
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

// Repository type metadata.
var (
	Account_Kind             = "Account"
	Account_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Account_Kind}.String()
	Account_KindAPIVersion   = Account_Kind + "." + CRDGroupVersion.String()
	Account_GroupVersionKind = CRDGroupVersion.WithKind(Account_Kind)
)

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}
