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

type AttachmentsSourceInitParameters struct {

	// The key describing the location of an attachment to a document. Valid key types include: SourceUrl and S3FileUrl
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The name of the document attachment file
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value describing the location of an attachment to a document
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AttachmentsSourceObservation struct {

	// The key describing the location of an attachment to a document. Valid key types include: SourceUrl and S3FileUrl
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The name of the document attachment file
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value describing the location of an attachment to a document
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type AttachmentsSourceParameters struct {

	// The key describing the location of an attachment to a document. Valid key types include: SourceUrl and S3FileUrl
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// The name of the document attachment file
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value describing the location of an attachment to a document
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type DocumentInitParameters struct {

	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	AttachmentsSource []AttachmentsSourceInitParameters `json:"attachmentsSource,omitempty" tf:"attachments_source,omitempty"`

	// The JSON or YAML content of the document.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The format of the document. Valid document types include: JSON and YAML
	DocumentFormat *string `json:"documentFormat,omitempty" tf:"document_format,omitempty"`

	// The type of the document. Valid document types include: Automation, Command, Package, Policy, and Session
	DocumentType *string `json:"documentType,omitempty" tf:"document_type,omitempty"`

	// Additional Permissions to attach to the document. See Permissions below for details.
	// +mapType=granular
	Permissions map[string]*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`

	// A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.
	VersionName *string `json:"versionName,omitempty" tf:"version_name,omitempty"`
}

type DocumentObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	AttachmentsSource []AttachmentsSourceObservation `json:"attachmentsSource,omitempty" tf:"attachments_source,omitempty"`

	// The JSON or YAML content of the document.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The date the document was created.
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// The default version of the document.
	DefaultVersion *string `json:"defaultVersion,omitempty" tf:"default_version,omitempty"`

	// The description of the document.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The format of the document. Valid document types include: JSON and YAML
	DocumentFormat *string `json:"documentFormat,omitempty" tf:"document_format,omitempty"`

	// The type of the document. Valid document types include: Automation, Command, Package, Policy, and Session
	DocumentType *string `json:"documentType,omitempty" tf:"document_type,omitempty"`

	// The document version.
	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`

	// The sha1 or sha256 of the document content
	Hash *string `json:"hash,omitempty" tf:"hash,omitempty"`

	// "Sha1" "Sha256". The hashing algorithm used when hashing the content.
	HashType *string `json:"hashType,omitempty" tf:"hash_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The latest version of the document.
	LatestVersion *string `json:"latestVersion,omitempty" tf:"latest_version,omitempty"`

	// The AWS user account of the person who created the document.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The parameters that are available to this document.
	Parameter []ParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Additional Permissions to attach to the document. See Permissions below for details.
	// +mapType=granular
	Permissions map[string]*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// A list of OS platforms compatible with this SSM document, either "Windows" or "Linux".
	PlatformTypes []*string `json:"platformTypes,omitempty" tf:"platform_types,omitempty"`

	// The schema version of the document.
	SchemaVersion *string `json:"schemaVersion,omitempty" tf:"schema_version,omitempty"`

	// "Creating", "Active" or "Deleting". The current status of the document.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`

	// A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.
	VersionName *string `json:"versionName,omitempty" tf:"version_name,omitempty"`
}

type DocumentParameters struct {

	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	// +kubebuilder:validation:Optional
	AttachmentsSource []AttachmentsSourceParameters `json:"attachmentsSource,omitempty" tf:"attachments_source,omitempty"`

	// The JSON or YAML content of the document.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The format of the document. Valid document types include: JSON and YAML
	// +kubebuilder:validation:Optional
	DocumentFormat *string `json:"documentFormat,omitempty" tf:"document_format,omitempty"`

	// The type of the document. Valid document types include: Automation, Command, Package, Policy, and Session
	// +kubebuilder:validation:Optional
	DocumentType *string `json:"documentType,omitempty" tf:"document_type,omitempty"`

	// Additional Permissions to attach to the document. See Permissions below for details.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Permissions map[string]*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	// +kubebuilder:validation:Optional
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`

	// A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.
	// +kubebuilder:validation:Optional
	VersionName *string `json:"versionName,omitempty" tf:"version_name,omitempty"`
}

type ParameterInitParameters struct {
}

type ParameterObservation struct {
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// The description of the document.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the document.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The permission type for the document. The permission type can be Share.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ParameterParameters struct {
}

// DocumentSpec defines the desired state of Document
type DocumentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DocumentParameters `json:"forProvider"`
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
	InitProvider DocumentInitParameters `json:"initProvider,omitempty"`
}

// DocumentStatus defines the observed state of Document.
type DocumentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DocumentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Document is the Schema for the Documents API. Provides an SSM Document resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Document struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.content) || (has(self.initProvider) && has(self.initProvider.content))",message="spec.forProvider.content is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.documentType) || (has(self.initProvider) && has(self.initProvider.documentType))",message="spec.forProvider.documentType is a required parameter"
	Spec   DocumentSpec   `json:"spec"`
	Status DocumentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentList contains a list of Documents
type DocumentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Document `json:"items"`
}

// Repository type metadata.
var (
	Document_Kind             = "Document"
	Document_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Document_Kind}.String()
	Document_KindAPIVersion   = Document_Kind + "." + CRDGroupVersion.String()
	Document_GroupVersionKind = CRDGroupVersion.WithKind(Document_Kind)
)

func init() {
	SchemeBuilder.Register(&Document{}, &DocumentList{})
}
