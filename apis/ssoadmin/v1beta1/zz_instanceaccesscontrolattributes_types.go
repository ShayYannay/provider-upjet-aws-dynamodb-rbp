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

type AttributeInitParameters struct {

	// The name of the attribute associated with your identities in your identity source. This is used to map a specified attribute in your identity source with an attribute in AWS SSO.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The value used for mapping a specified attribute to an identity source. See AccessControlAttributeValue
	Value []ValueInitParameters `json:"value,omitempty" tf:"value,omitempty"`
}

type AttributeObservation struct {

	// The name of the attribute associated with your identities in your identity source. This is used to map a specified attribute in your identity source with an attribute in AWS SSO.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The value used for mapping a specified attribute to an identity source. See AccessControlAttributeValue
	Value []ValueObservation `json:"value,omitempty" tf:"value,omitempty"`
}

type AttributeParameters struct {

	// The name of the attribute associated with your identities in your identity source. This is used to map a specified attribute in your identity source with an attribute in AWS SSO.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// The value used for mapping a specified attribute to an identity source. See AccessControlAttributeValue
	// +kubebuilder:validation:Optional
	Value []ValueParameters `json:"value" tf:"value,omitempty"`
}

type InstanceAccessControlAttributesInitParameters struct {

	// See AccessControlAttribute for more details.
	Attribute []AttributeInitParameters `json:"attribute,omitempty" tf:"attribute,omitempty"`
}

type InstanceAccessControlAttributesObservation struct {

	// See AccessControlAttribute for more details.
	Attribute []AttributeObservation `json:"attribute,omitempty" tf:"attribute,omitempty"`

	// The identifier of the Instance Access Control Attribute instance_arn.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Resource Name (ARN) of the SSO Instance.
	InstanceArn *string `json:"instanceArn,omitempty" tf:"instance_arn,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	StatusReason *string `json:"statusReason,omitempty" tf:"status_reason,omitempty"`
}

type InstanceAccessControlAttributesParameters struct {

	// See AccessControlAttribute for more details.
	// +kubebuilder:validation:Optional
	Attribute []AttributeParameters `json:"attribute,omitempty" tf:"attribute,omitempty"`

	// The Amazon Resource Name (ARN) of the SSO Instance.
	// +kubebuilder:validation:Required
	InstanceArn *string `json:"instanceArn" tf:"instance_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type ValueInitParameters struct {

	// The identity source to use when mapping a specified attribute to AWS SSO.
	// +listType=set
	Source []*string `json:"source,omitempty" tf:"source,omitempty"`
}

type ValueObservation struct {

	// The identity source to use when mapping a specified attribute to AWS SSO.
	// +listType=set
	Source []*string `json:"source,omitempty" tf:"source,omitempty"`
}

type ValueParameters struct {

	// The identity source to use when mapping a specified attribute to AWS SSO.
	// +kubebuilder:validation:Optional
	// +listType=set
	Source []*string `json:"source" tf:"source,omitempty"`
}

// InstanceAccessControlAttributesSpec defines the desired state of InstanceAccessControlAttributes
type InstanceAccessControlAttributesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceAccessControlAttributesParameters `json:"forProvider"`
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
	InitProvider InstanceAccessControlAttributesInitParameters `json:"initProvider,omitempty"`
}

// InstanceAccessControlAttributesStatus defines the observed state of InstanceAccessControlAttributes.
type InstanceAccessControlAttributesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceAccessControlAttributesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InstanceAccessControlAttributes is the Schema for the InstanceAccessControlAttributess API. Provides a Single Sign-On (SSO) ABAC Resource: https://docs.aws.amazon.com/singlesignon/latest/userguide/abac.html
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InstanceAccessControlAttributes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.attribute) || (has(self.initProvider) && has(self.initProvider.attribute))",message="spec.forProvider.attribute is a required parameter"
	Spec   InstanceAccessControlAttributesSpec   `json:"spec"`
	Status InstanceAccessControlAttributesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceAccessControlAttributesList contains a list of InstanceAccessControlAttributess
type InstanceAccessControlAttributesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceAccessControlAttributes `json:"items"`
}

// Repository type metadata.
var (
	InstanceAccessControlAttributes_Kind             = "InstanceAccessControlAttributes"
	InstanceAccessControlAttributes_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceAccessControlAttributes_Kind}.String()
	InstanceAccessControlAttributes_KindAPIVersion   = InstanceAccessControlAttributes_Kind + "." + CRDGroupVersion.String()
	InstanceAccessControlAttributes_GroupVersionKind = CRDGroupVersion.WithKind(InstanceAccessControlAttributes_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceAccessControlAttributes{}, &InstanceAccessControlAttributesList{})
}
