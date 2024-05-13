// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AMICopyEBSBlockDeviceInitParameters struct {
}

type AMICopyEBSBlockDeviceObservation struct {
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Region-unique name for the AMI.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Whether the destination snapshots of the copied image should be encrypted. Defaults to false
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// ARN of the AMI.
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// ID of the created AMI.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type AMICopyEBSBlockDeviceParameters struct {
}

type AMICopyEphemeralBlockDeviceInitParameters struct {
}

type AMICopyEphemeralBlockDeviceObservation struct {

	// Region-unique name for the AMI.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Region-unique name for the AMI.
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type AMICopyEphemeralBlockDeviceParameters struct {
}

type AMICopyInitParameters struct {
	DeprecationTime *string `json:"deprecationTime,omitempty" tf:"deprecation_time,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ARN of the Outpost to which to copy the AMI.
	// Only specify this parameter when copying an AMI from an AWS Region to an Outpost. The AMI must be in the Region of the destination Outpost.
	DestinationOutpostArn *string `json:"destinationOutpostArn,omitempty" tf:"destination_outpost_arn,omitempty"`

	EBSBlockDevice []AMICopyEBSBlockDeviceInitParameters `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	// Whether the destination snapshots of the copied image should be encrypted. Defaults to false
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	EphemeralBlockDevice []AMICopyEphemeralBlockDeviceInitParameters `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`

	// Full ARN of the KMS Key to use when encrypting the snapshots of an image during a copy operation. If not specified, then the default AWS KMS Key will be used
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Region-unique name for the AMI.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Id of the AMI to copy. This id must be valid in the region
	// given by source_ami_region.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.AMI
	SourceAMIID *string `json:"sourceAmiId,omitempty" tf:"source_ami_id,omitempty"`

	// Reference to a AMI in ec2 to populate sourceAmiId.
	// +kubebuilder:validation:Optional
	SourceAMIIDRef *v1.Reference `json:"sourceAmiIdRef,omitempty" tf:"-"`

	// Selector for a AMI in ec2 to populate sourceAmiId.
	// +kubebuilder:validation:Optional
	SourceAMIIDSelector *v1.Selector `json:"sourceAmiIdSelector,omitempty" tf:"-"`

	// Region from which the AMI will be copied. This may be the
	// same as the AWS provider region in order to create a copy within the same region.
	SourceAMIRegion *string `json:"sourceAmiRegion,omitempty" tf:"source_ami_region,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AMICopyObservation struct {
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// ARN of the AMI.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	BootMode *string `json:"bootMode,omitempty" tf:"boot_mode,omitempty"`

	DeprecationTime *string `json:"deprecationTime,omitempty" tf:"deprecation_time,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ARN of the Outpost to which to copy the AMI.
	// Only specify this parameter when copying an AMI from an AWS Region to an Outpost. The AMI must be in the Region of the destination Outpost.
	DestinationOutpostArn *string `json:"destinationOutpostArn,omitempty" tf:"destination_outpost_arn,omitempty"`

	EBSBlockDevice []AMICopyEBSBlockDeviceObservation `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	EnaSupport *bool `json:"enaSupport,omitempty" tf:"ena_support,omitempty"`

	// Whether the destination snapshots of the copied image should be encrypted. Defaults to false
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	EphemeralBlockDevice []AMICopyEphemeralBlockDeviceObservation `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`

	Hypervisor *string `json:"hypervisor,omitempty" tf:"hypervisor,omitempty"`

	// ID of the created AMI.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ImageLocation *string `json:"imageLocation,omitempty" tf:"image_location,omitempty"`

	ImageOwnerAlias *string `json:"imageOwnerAlias,omitempty" tf:"image_owner_alias,omitempty"`

	ImageType *string `json:"imageType,omitempty" tf:"image_type,omitempty"`

	ImdsSupport *string `json:"imdsSupport,omitempty" tf:"imds_support,omitempty"`

	// Full ARN of the KMS Key to use when encrypting the snapshots of an image during a copy operation. If not specified, then the default AWS KMS Key will be used
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// ID of the created AMI.
	KernelID *string `json:"kernelId,omitempty" tf:"kernel_id,omitempty"`

	ManageEBSSnapshots *bool `json:"manageEbsSnapshots,omitempty" tf:"manage_ebs_snapshots,omitempty"`

	// Region-unique name for the AMI.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the created AMI.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	PlatformDetails *string `json:"platformDetails,omitempty" tf:"platform_details,omitempty"`

	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// ID of the created AMI.
	RamdiskID *string `json:"ramdiskId,omitempty" tf:"ramdisk_id,omitempty"`

	// Region-unique name for the AMI.
	RootDeviceName *string `json:"rootDeviceName,omitempty" tf:"root_device_name,omitempty"`

	// ID of the created AMI.
	RootSnapshotID *string `json:"rootSnapshotId,omitempty" tf:"root_snapshot_id,omitempty"`

	// Id of the AMI to copy. This id must be valid in the region
	// given by source_ami_region.
	SourceAMIID *string `json:"sourceAmiId,omitempty" tf:"source_ami_id,omitempty"`

	// Region from which the AMI will be copied. This may be the
	// same as the AWS provider region in order to create a copy within the same region.
	SourceAMIRegion *string `json:"sourceAmiRegion,omitempty" tf:"source_ami_region,omitempty"`

	SriovNetSupport *string `json:"sriovNetSupport,omitempty" tf:"sriov_net_support,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	TpmSupport *string `json:"tpmSupport,omitempty" tf:"tpm_support,omitempty"`

	UsageOperation *string `json:"usageOperation,omitempty" tf:"usage_operation,omitempty"`

	VirtualizationType *string `json:"virtualizationType,omitempty" tf:"virtualization_type,omitempty"`
}

type AMICopyParameters struct {

	// +kubebuilder:validation:Optional
	DeprecationTime *string `json:"deprecationTime,omitempty" tf:"deprecation_time,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ARN of the Outpost to which to copy the AMI.
	// Only specify this parameter when copying an AMI from an AWS Region to an Outpost. The AMI must be in the Region of the destination Outpost.
	// +kubebuilder:validation:Optional
	DestinationOutpostArn *string `json:"destinationOutpostArn,omitempty" tf:"destination_outpost_arn,omitempty"`

	// +kubebuilder:validation:Optional
	EBSBlockDevice []AMICopyEBSBlockDeviceParameters `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	// Whether the destination snapshots of the copied image should be encrypted. Defaults to false
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	EphemeralBlockDevice []AMICopyEphemeralBlockDeviceParameters `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`

	// Full ARN of the KMS Key to use when encrypting the snapshots of an image during a copy operation. If not specified, then the default AWS KMS Key will be used
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Region-unique name for the AMI.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Id of the AMI to copy. This id must be valid in the region
	// given by source_ami_region.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.AMI
	// +kubebuilder:validation:Optional
	SourceAMIID *string `json:"sourceAmiId,omitempty" tf:"source_ami_id,omitempty"`

	// Reference to a AMI in ec2 to populate sourceAmiId.
	// +kubebuilder:validation:Optional
	SourceAMIIDRef *v1.Reference `json:"sourceAmiIdRef,omitempty" tf:"-"`

	// Selector for a AMI in ec2 to populate sourceAmiId.
	// +kubebuilder:validation:Optional
	SourceAMIIDSelector *v1.Selector `json:"sourceAmiIdSelector,omitempty" tf:"-"`

	// Region from which the AMI will be copied. This may be the
	// same as the AWS provider region in order to create a copy within the same region.
	// +kubebuilder:validation:Optional
	SourceAMIRegion *string `json:"sourceAmiRegion,omitempty" tf:"source_ami_region,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// AMICopySpec defines the desired state of AMICopy
type AMICopySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AMICopyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AMICopyInitParameters `json:"initProvider,omitempty"`
}

// AMICopyStatus defines the observed state of AMICopy.
type AMICopyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AMICopyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AMICopy is the Schema for the AMICopys API. Duplicates an existing Amazon Machine Image (AMI)
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AMICopy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourceAmiRegion) || (has(self.initProvider) && has(self.initProvider.sourceAmiRegion))",message="spec.forProvider.sourceAmiRegion is a required parameter"
	Spec   AMICopySpec   `json:"spec"`
	Status AMICopyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AMICopyList contains a list of AMICopys
type AMICopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AMICopy `json:"items"`
}

// Repository type metadata.
var (
	AMICopy_Kind             = "AMICopy"
	AMICopy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AMICopy_Kind}.String()
	AMICopy_KindAPIVersion   = AMICopy_Kind + "." + CRDGroupVersion.String()
	AMICopy_GroupVersionKind = CRDGroupVersion.WithKind(AMICopy_Kind)
)

func init() {
	SchemeBuilder.Register(&AMICopy{}, &AMICopyList{})
}
