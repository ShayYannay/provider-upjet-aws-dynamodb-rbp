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

type DomainPolicyInitParameters struct {

	// IAM policy document specifying the access policies for the domain
	AccessPolicies *string `json:"accessPolicies,omitempty" tf:"access_policies,omitempty"`

	// Name of the domain.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/opensearch/v1beta1.Domain
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Reference to a Domain in opensearch to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameRef *v1.Reference `json:"domainNameRef,omitempty" tf:"-"`

	// Selector for a Domain in opensearch to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameSelector *v1.Selector `json:"domainNameSelector,omitempty" tf:"-"`
}

type DomainPolicyObservation struct {

	// IAM policy document specifying the access policies for the domain
	AccessPolicies *string `json:"accessPolicies,omitempty" tf:"access_policies,omitempty"`

	// Name of the domain.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DomainPolicyParameters struct {

	// IAM policy document specifying the access policies for the domain
	// +kubebuilder:validation:Optional
	AccessPolicies *string `json:"accessPolicies,omitempty" tf:"access_policies,omitempty"`

	// Name of the domain.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/opensearch/v1beta1.Domain
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Reference to a Domain in opensearch to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameRef *v1.Reference `json:"domainNameRef,omitempty" tf:"-"`

	// Selector for a Domain in opensearch to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameSelector *v1.Selector `json:"domainNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DomainPolicySpec defines the desired state of DomainPolicy
type DomainPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainPolicyParameters `json:"forProvider"`
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
	InitProvider DomainPolicyInitParameters `json:"initProvider,omitempty"`
}

// DomainPolicyStatus defines the observed state of DomainPolicy.
type DomainPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DomainPolicy is the Schema for the DomainPolicys API. Provides an OpenSearch Domain Policy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DomainPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessPolicies) || (has(self.initProvider) && has(self.initProvider.accessPolicies))",message="spec.forProvider.accessPolicies is a required parameter"
	Spec   DomainPolicySpec   `json:"spec"`
	Status DomainPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainPolicyList contains a list of DomainPolicys
type DomainPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainPolicy `json:"items"`
}

// Repository type metadata.
var (
	DomainPolicy_Kind             = "DomainPolicy"
	DomainPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainPolicy_Kind}.String()
	DomainPolicy_KindAPIVersion   = DomainPolicy_Kind + "." + CRDGroupVersion.String()
	DomainPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DomainPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainPolicy{}, &DomainPolicyList{})
}
