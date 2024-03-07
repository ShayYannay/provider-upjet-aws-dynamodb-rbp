// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type HSMConfigurationInitParameters struct {

	// A text description of the HSM configuration to be created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP address that the Amazon Redshift cluster must use to access the HSM.
	HSMIPAddress *string `json:"hsmIpAddress,omitempty" tf:"hsm_ip_address,omitempty"`

	// The name of the partition in the HSM where the Amazon Redshift clusters will store their database encryption keys.
	HSMPartitionName *string `json:"hsmPartitionName,omitempty" tf:"hsm_partition_name,omitempty"`

	// The HSMs public certificate file. When using Cloud HSM, the file name is server.pem.
	HSMServerPublicCertificate *string `json:"hsmServerPublicCertificate,omitempty" tf:"hsm_server_public_certificate,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HSMConfigurationObservation struct {

	// Amazon Resource Name (ARN) of the Hsm Client Certificate.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A text description of the HSM configuration to be created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP address that the Amazon Redshift cluster must use to access the HSM.
	HSMIPAddress *string `json:"hsmIpAddress,omitempty" tf:"hsm_ip_address,omitempty"`

	// The name of the partition in the HSM where the Amazon Redshift clusters will store their database encryption keys.
	HSMPartitionName *string `json:"hsmPartitionName,omitempty" tf:"hsm_partition_name,omitempty"`

	// The HSMs public certificate file. When using Cloud HSM, the file name is server.pem.
	HSMServerPublicCertificate *string `json:"hsmServerPublicCertificate,omitempty" tf:"hsm_server_public_certificate,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type HSMConfigurationParameters struct {

	// A text description of the HSM configuration to be created.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP address that the Amazon Redshift cluster must use to access the HSM.
	// +kubebuilder:validation:Optional
	HSMIPAddress *string `json:"hsmIpAddress,omitempty" tf:"hsm_ip_address,omitempty"`

	// The name of the partition in the HSM where the Amazon Redshift clusters will store their database encryption keys.
	// +kubebuilder:validation:Optional
	HSMPartitionName *string `json:"hsmPartitionName,omitempty" tf:"hsm_partition_name,omitempty"`

	// The password required to access the HSM partition.
	// +kubebuilder:validation:Optional
	HSMPartitionPasswordSecretRef v1.SecretKeySelector `json:"hsmPartitionPasswordSecretRef" tf:"-"`

	// The HSMs public certificate file. When using Cloud HSM, the file name is server.pem.
	// +kubebuilder:validation:Optional
	HSMServerPublicCertificate *string `json:"hsmServerPublicCertificate,omitempty" tf:"hsm_server_public_certificate,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// HSMConfigurationSpec defines the desired state of HSMConfiguration
type HSMConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HSMConfigurationParameters `json:"forProvider"`
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
	InitProvider HSMConfigurationInitParameters `json:"initProvider,omitempty"`
}

// HSMConfigurationStatus defines the observed state of HSMConfiguration.
type HSMConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HSMConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HSMConfiguration is the Schema for the HSMConfigurations API. Creates an HSM configuration that contains the information required by an Amazon Redshift cluster to store and use database encryption keys in a Hardware Security Module (HSM).
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type HSMConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hsmIpAddress) || (has(self.initProvider) && has(self.initProvider.hsmIpAddress))",message="spec.forProvider.hsmIpAddress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hsmPartitionName) || (has(self.initProvider) && has(self.initProvider.hsmPartitionName))",message="spec.forProvider.hsmPartitionName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hsmPartitionPasswordSecretRef)",message="spec.forProvider.hsmPartitionPasswordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hsmServerPublicCertificate) || (has(self.initProvider) && has(self.initProvider.hsmServerPublicCertificate))",message="spec.forProvider.hsmServerPublicCertificate is a required parameter"
	Spec   HSMConfigurationSpec   `json:"spec"`
	Status HSMConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HSMConfigurationList contains a list of HSMConfigurations
type HSMConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HSMConfiguration `json:"items"`
}

// Repository type metadata.
var (
	HSMConfiguration_Kind             = "HSMConfiguration"
	HSMConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HSMConfiguration_Kind}.String()
	HSMConfiguration_KindAPIVersion   = HSMConfiguration_Kind + "." + CRDGroupVersion.String()
	HSMConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(HSMConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&HSMConfiguration{}, &HSMConfigurationList{})
}
