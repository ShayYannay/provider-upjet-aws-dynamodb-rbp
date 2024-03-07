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

type CatalogDataInitParameters struct {

	// A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format.
	AboutText *string `json:"aboutText,omitempty" tf:"about_text,omitempty"`

	// The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: ARM, ARM 64, x86, x86-64
	// +listType=set
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute.
	LogoImageBlob *string `json:"logoImageBlob,omitempty" tf:"logo_image_blob,omitempty"`

	// The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters: Linux, Windows
	// +listType=set
	OperatingSystems []*string `json:"operatingSystems,omitempty" tf:"operating_systems,omitempty"`

	// Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format.
	UsageText *string `json:"usageText,omitempty" tf:"usage_text,omitempty"`
}

type CatalogDataObservation struct {

	// A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format.
	AboutText *string `json:"aboutText,omitempty" tf:"about_text,omitempty"`

	// The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: ARM, ARM 64, x86, x86-64
	// +listType=set
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute.
	LogoImageBlob *string `json:"logoImageBlob,omitempty" tf:"logo_image_blob,omitempty"`

	// The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters: Linux, Windows
	// +listType=set
	OperatingSystems []*string `json:"operatingSystems,omitempty" tf:"operating_systems,omitempty"`

	// Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format.
	UsageText *string `json:"usageText,omitempty" tf:"usage_text,omitempty"`
}

type CatalogDataParameters struct {

	// A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format.
	// +kubebuilder:validation:Optional
	AboutText *string `json:"aboutText,omitempty" tf:"about_text,omitempty"`

	// The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: ARM, ARM 64, x86, x86-64
	// +kubebuilder:validation:Optional
	// +listType=set
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute.
	// +kubebuilder:validation:Optional
	LogoImageBlob *string `json:"logoImageBlob,omitempty" tf:"logo_image_blob,omitempty"`

	// The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters: Linux, Windows
	// +kubebuilder:validation:Optional
	// +listType=set
	OperatingSystems []*string `json:"operatingSystems,omitempty" tf:"operating_systems,omitempty"`

	// Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format.
	// +kubebuilder:validation:Optional
	UsageText *string `json:"usageText,omitempty" tf:"usage_text,omitempty"`
}

type RepositoryInitParameters struct {

	// Catalog data configuration for the repository. See below for schema.
	CatalogData []CatalogDataInitParameters `json:"catalogData,omitempty" tf:"catalog_data,omitempty"`

	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RepositoryObservation struct {

	// Full ARN of the repository.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Catalog data configuration for the repository. See below for schema.
	CatalogData []CatalogDataObservation `json:"catalogData,omitempty" tf:"catalog_data,omitempty"`

	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The repository name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The registry ID where the repository was created.
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// The URI of the repository.
	RepositoryURI *string `json:"repositoryUri,omitempty" tf:"repository_uri,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RepositoryParameters struct {

	// Catalog data configuration for the repository. See below for schema.
	// +kubebuilder:validation:Optional
	CatalogData []CatalogDataParameters `json:"catalogData,omitempty" tf:"catalog_data,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryParameters `json:"forProvider"`
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
	InitProvider RepositoryInitParameters `json:"initProvider,omitempty"`
}

// RepositoryStatus defines the observed state of Repository.
type RepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Repository is the Schema for the Repositorys API. Provides a Public Elastic Container Registry Repository.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositorySpec   `json:"spec"`
	Status            RepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryList contains a list of Repositorys
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

// Repository type metadata.
var (
	Repository_Kind             = "Repository"
	Repository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Repository_Kind}.String()
	Repository_KindAPIVersion   = Repository_Kind + "." + CRDGroupVersion.String()
	Repository_GroupVersionKind = CRDGroupVersion.WithKind(Repository_Kind)
)

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}
