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

type LifecyclePolicyInitParameters struct {

	// Description of the policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// JSON policy document to use as the content for the new policy.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Type of lifecycle policy. Must be retention.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LifecyclePolicyObservation struct {

	// Description of the policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// JSON policy document to use as the content for the new policy.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Version of the policy.
	PolicyVersion *string `json:"policyVersion,omitempty" tf:"policy_version,omitempty"`

	// Type of lifecycle policy. Must be retention.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LifecyclePolicyParameters struct {

	// Description of the policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// JSON policy document to use as the content for the new policy.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Type of lifecycle policy. Must be retention.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// LifecyclePolicySpec defines the desired state of LifecyclePolicy
type LifecyclePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LifecyclePolicyParameters `json:"forProvider"`
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
	InitProvider LifecyclePolicyInitParameters `json:"initProvider,omitempty"`
}

// LifecyclePolicyStatus defines the observed state of LifecyclePolicy.
type LifecyclePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LifecyclePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LifecyclePolicy is the Schema for the LifecyclePolicys API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   LifecyclePolicySpec   `json:"spec"`
	Status LifecyclePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LifecyclePolicyList contains a list of LifecyclePolicys
type LifecyclePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LifecyclePolicy `json:"items"`
}

// Repository type metadata.
var (
	LifecyclePolicy_Kind             = "LifecyclePolicy"
	LifecyclePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LifecyclePolicy_Kind}.String()
	LifecyclePolicy_KindAPIVersion   = LifecyclePolicy_Kind + "." + CRDGroupVersion.String()
	LifecyclePolicy_GroupVersionKind = CRDGroupVersion.WithKind(LifecyclePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&LifecyclePolicy{}, &LifecyclePolicyList{})
}
