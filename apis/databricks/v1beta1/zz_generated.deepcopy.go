//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessConnector) DeepCopyInto(out *AccessConnector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessConnector.
func (in *AccessConnector) DeepCopy() *AccessConnector {
	if in == nil {
		return nil
	}
	out := new(AccessConnector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessConnector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessConnectorList) DeepCopyInto(out *AccessConnectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessConnector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessConnectorList.
func (in *AccessConnectorList) DeepCopy() *AccessConnectorList {
	if in == nil {
		return nil
	}
	out := new(AccessConnectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessConnectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessConnectorObservation) DeepCopyInto(out *AccessConnectorObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessConnectorObservation.
func (in *AccessConnectorObservation) DeepCopy() *AccessConnectorObservation {
	if in == nil {
		return nil
	}
	out := new(AccessConnectorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessConnectorParameters) DeepCopyInto(out *AccessConnectorParameters) {
	*out = *in
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessConnectorParameters.
func (in *AccessConnectorParameters) DeepCopy() *AccessConnectorParameters {
	if in == nil {
		return nil
	}
	out := new(AccessConnectorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessConnectorSpec) DeepCopyInto(out *AccessConnectorSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessConnectorSpec.
func (in *AccessConnectorSpec) DeepCopy() *AccessConnectorSpec {
	if in == nil {
		return nil
	}
	out := new(AccessConnectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessConnectorStatus) DeepCopyInto(out *AccessConnectorStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessConnectorStatus.
func (in *AccessConnectorStatus) DeepCopy() *AccessConnectorStatus {
	if in == nil {
		return nil
	}
	out := new(AccessConnectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomParametersObservation) DeepCopyInto(out *CustomParametersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomParametersObservation.
func (in *CustomParametersObservation) DeepCopy() *CustomParametersObservation {
	if in == nil {
		return nil
	}
	out := new(CustomParametersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomParametersParameters) DeepCopyInto(out *CustomParametersParameters) {
	*out = *in
	if in.MachineLearningWorkspaceID != nil {
		in, out := &in.MachineLearningWorkspaceID, &out.MachineLearningWorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.NATGatewayName != nil {
		in, out := &in.NATGatewayName, &out.NATGatewayName
		*out = new(string)
		**out = **in
	}
	if in.NoPublicIP != nil {
		in, out := &in.NoPublicIP, &out.NoPublicIP
		*out = new(bool)
		**out = **in
	}
	if in.PrivateSubnetName != nil {
		in, out := &in.PrivateSubnetName, &out.PrivateSubnetName
		*out = new(string)
		**out = **in
	}
	if in.PrivateSubnetNameRef != nil {
		in, out := &in.PrivateSubnetNameRef, &out.PrivateSubnetNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateSubnetNameSelector != nil {
		in, out := &in.PrivateSubnetNameSelector, &out.PrivateSubnetNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateSubnetNetworkSecurityGroupAssociationID != nil {
		in, out := &in.PrivateSubnetNetworkSecurityGroupAssociationID, &out.PrivateSubnetNetworkSecurityGroupAssociationID
		*out = new(string)
		**out = **in
	}
	if in.PublicIPName != nil {
		in, out := &in.PublicIPName, &out.PublicIPName
		*out = new(string)
		**out = **in
	}
	if in.PublicSubnetName != nil {
		in, out := &in.PublicSubnetName, &out.PublicSubnetName
		*out = new(string)
		**out = **in
	}
	if in.PublicSubnetNameRef != nil {
		in, out := &in.PublicSubnetNameRef, &out.PublicSubnetNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.PublicSubnetNameSelector != nil {
		in, out := &in.PublicSubnetNameSelector, &out.PublicSubnetNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PublicSubnetNetworkSecurityGroupAssociationID != nil {
		in, out := &in.PublicSubnetNetworkSecurityGroupAssociationID, &out.PublicSubnetNetworkSecurityGroupAssociationID
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountName != nil {
		in, out := &in.StorageAccountName, &out.StorageAccountName
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountSkuName != nil {
		in, out := &in.StorageAccountSkuName, &out.StorageAccountSkuName
		*out = new(string)
		**out = **in
	}
	if in.VirtualNetworkID != nil {
		in, out := &in.VirtualNetworkID, &out.VirtualNetworkID
		*out = new(string)
		**out = **in
	}
	if in.VnetAddressPrefix != nil {
		in, out := &in.VnetAddressPrefix, &out.VnetAddressPrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomParametersParameters.
func (in *CustomParametersParameters) DeepCopy() *CustomParametersParameters {
	if in == nil {
		return nil
	}
	out := new(CustomParametersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityObservation) DeepCopyInto(out *IdentityObservation) {
	*out = *in
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityObservation.
func (in *IdentityObservation) DeepCopy() *IdentityObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityParameters) DeepCopyInto(out *IdentityParameters) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityParameters.
func (in *IdentityParameters) DeepCopy() *IdentityParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedDiskIdentityObservation) DeepCopyInto(out *ManagedDiskIdentityObservation) {
	*out = *in
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedDiskIdentityObservation.
func (in *ManagedDiskIdentityObservation) DeepCopy() *ManagedDiskIdentityObservation {
	if in == nil {
		return nil
	}
	out := new(ManagedDiskIdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedDiskIdentityParameters) DeepCopyInto(out *ManagedDiskIdentityParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedDiskIdentityParameters.
func (in *ManagedDiskIdentityParameters) DeepCopy() *ManagedDiskIdentityParameters {
	if in == nil {
		return nil
	}
	out := new(ManagedDiskIdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageAccountIdentityObservation) DeepCopyInto(out *StorageAccountIdentityObservation) {
	*out = *in
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageAccountIdentityObservation.
func (in *StorageAccountIdentityObservation) DeepCopy() *StorageAccountIdentityObservation {
	if in == nil {
		return nil
	}
	out := new(StorageAccountIdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageAccountIdentityParameters) DeepCopyInto(out *StorageAccountIdentityParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageAccountIdentityParameters.
func (in *StorageAccountIdentityParameters) DeepCopy() *StorageAccountIdentityParameters {
	if in == nil {
		return nil
	}
	out := new(StorageAccountIdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspace) DeepCopyInto(out *Workspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace.
func (in *Workspace) DeepCopy() *Workspace {
	if in == nil {
		return nil
	}
	out := new(Workspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCustomerManagedKey) DeepCopyInto(out *WorkspaceCustomerManagedKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCustomerManagedKey.
func (in *WorkspaceCustomerManagedKey) DeepCopy() *WorkspaceCustomerManagedKey {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCustomerManagedKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceCustomerManagedKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCustomerManagedKeyList) DeepCopyInto(out *WorkspaceCustomerManagedKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkspaceCustomerManagedKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCustomerManagedKeyList.
func (in *WorkspaceCustomerManagedKeyList) DeepCopy() *WorkspaceCustomerManagedKeyList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCustomerManagedKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceCustomerManagedKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCustomerManagedKeyObservation) DeepCopyInto(out *WorkspaceCustomerManagedKeyObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCustomerManagedKeyObservation.
func (in *WorkspaceCustomerManagedKeyObservation) DeepCopy() *WorkspaceCustomerManagedKeyObservation {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCustomerManagedKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCustomerManagedKeyParameters) DeepCopyInto(out *WorkspaceCustomerManagedKeyParameters) {
	*out = *in
	if in.KeyVaultKeyID != nil {
		in, out := &in.KeyVaultKeyID, &out.KeyVaultKeyID
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultKeyIDRef != nil {
		in, out := &in.KeyVaultKeyIDRef, &out.KeyVaultKeyIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultKeyIDSelector != nil {
		in, out := &in.KeyVaultKeyIDSelector, &out.KeyVaultKeyIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkspaceID != nil {
		in, out := &in.WorkspaceID, &out.WorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceIDRef != nil {
		in, out := &in.WorkspaceIDRef, &out.WorkspaceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkspaceIDSelector != nil {
		in, out := &in.WorkspaceIDSelector, &out.WorkspaceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCustomerManagedKeyParameters.
func (in *WorkspaceCustomerManagedKeyParameters) DeepCopy() *WorkspaceCustomerManagedKeyParameters {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCustomerManagedKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCustomerManagedKeySpec) DeepCopyInto(out *WorkspaceCustomerManagedKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCustomerManagedKeySpec.
func (in *WorkspaceCustomerManagedKeySpec) DeepCopy() *WorkspaceCustomerManagedKeySpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCustomerManagedKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCustomerManagedKeyStatus) DeepCopyInto(out *WorkspaceCustomerManagedKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCustomerManagedKeyStatus.
func (in *WorkspaceCustomerManagedKeyStatus) DeepCopy() *WorkspaceCustomerManagedKeyStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCustomerManagedKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceList) DeepCopyInto(out *WorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceList.
func (in *WorkspaceList) DeepCopy() *WorkspaceList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceObservation) DeepCopyInto(out *WorkspaceObservation) {
	*out = *in
	if in.DiskEncryptionSetID != nil {
		in, out := &in.DiskEncryptionSetID, &out.DiskEncryptionSetID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ManagedDiskIdentity != nil {
		in, out := &in.ManagedDiskIdentity, &out.ManagedDiskIdentity
		*out = make([]ManagedDiskIdentityObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ManagedResourceGroupID != nil {
		in, out := &in.ManagedResourceGroupID, &out.ManagedResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountIdentity != nil {
		in, out := &in.StorageAccountIdentity, &out.StorageAccountIdentity
		*out = make([]StorageAccountIdentityObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WorkspaceID != nil {
		in, out := &in.WorkspaceID, &out.WorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceURL != nil {
		in, out := &in.WorkspaceURL, &out.WorkspaceURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceObservation.
func (in *WorkspaceObservation) DeepCopy() *WorkspaceObservation {
	if in == nil {
		return nil
	}
	out := new(WorkspaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceParameters) DeepCopyInto(out *WorkspaceParameters) {
	*out = *in
	if in.CustomParameters != nil {
		in, out := &in.CustomParameters, &out.CustomParameters
		*out = make([]CustomParametersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomerManagedKeyEnabled != nil {
		in, out := &in.CustomerManagedKeyEnabled, &out.CustomerManagedKeyEnabled
		*out = new(bool)
		**out = **in
	}
	if in.InfrastructureEncryptionEnabled != nil {
		in, out := &in.InfrastructureEncryptionEnabled, &out.InfrastructureEncryptionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.LoadBalancerBackendAddressPoolID != nil {
		in, out := &in.LoadBalancerBackendAddressPoolID, &out.LoadBalancerBackendAddressPoolID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ManagedDiskCmkKeyVaultKeyID != nil {
		in, out := &in.ManagedDiskCmkKeyVaultKeyID, &out.ManagedDiskCmkKeyVaultKeyID
		*out = new(string)
		**out = **in
	}
	if in.ManagedDiskCmkRotationToLatestVersionEnabled != nil {
		in, out := &in.ManagedDiskCmkRotationToLatestVersionEnabled, &out.ManagedDiskCmkRotationToLatestVersionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ManagedResourceGroupName != nil {
		in, out := &in.ManagedResourceGroupName, &out.ManagedResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ManagedResourceGroupNameRef != nil {
		in, out := &in.ManagedResourceGroupNameRef, &out.ManagedResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagedResourceGroupNameSelector != nil {
		in, out := &in.ManagedResourceGroupNameSelector, &out.ManagedResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagedServicesCmkKeyVaultKeyID != nil {
		in, out := &in.ManagedServicesCmkKeyVaultKeyID, &out.ManagedServicesCmkKeyVaultKeyID
		*out = new(string)
		**out = **in
	}
	if in.NetworkSecurityGroupRulesRequired != nil {
		in, out := &in.NetworkSecurityGroupRulesRequired, &out.NetworkSecurityGroupRulesRequired
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceParameters.
func (in *WorkspaceParameters) DeepCopy() *WorkspaceParameters {
	if in == nil {
		return nil
	}
	out := new(WorkspaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpec) DeepCopyInto(out *WorkspaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpec.
func (in *WorkspaceSpec) DeepCopy() *WorkspaceSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceStatus) DeepCopyInto(out *WorkspaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceStatus.
func (in *WorkspaceStatus) DeepCopy() *WorkspaceStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceStatus)
	in.DeepCopyInto(out)
	return out
}
