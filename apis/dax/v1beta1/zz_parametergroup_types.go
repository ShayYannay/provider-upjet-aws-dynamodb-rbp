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

type ParameterGroupInitParameters struct {

	// A description of the parameter group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  The parameters of the parameter group.
	Parameters []ParametersInitParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type ParameterGroupObservation struct {

	// A description of the parameter group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the parameter group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// –  The parameters of the parameter group.
	Parameters []ParametersObservation `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type ParameterGroupParameters struct {

	// A description of the parameter group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  The parameters of the parameter group.
	// +kubebuilder:validation:Optional
	Parameters []ParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type ParametersInitParameters struct {

	// The name of the parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value for the parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParametersObservation struct {

	// The name of the parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value for the parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParametersParameters struct {

	// The name of the parameter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value for the parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// ParameterGroupSpec defines the desired state of ParameterGroup
type ParameterGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ParameterGroupParameters `json:"forProvider"`
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
	InitProvider ParameterGroupInitParameters `json:"initProvider,omitempty"`
}

// ParameterGroupStatus defines the observed state of ParameterGroup.
type ParameterGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ParameterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ParameterGroup is the Schema for the ParameterGroups API. Provides an DAX Parameter Group resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ParameterGroupSpec   `json:"spec"`
	Status            ParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ParameterGroupList contains a list of ParameterGroups
type ParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ParameterGroup `json:"items"`
}

// Repository type metadata.
var (
	ParameterGroup_Kind             = "ParameterGroup"
	ParameterGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ParameterGroup_Kind}.String()
	ParameterGroup_KindAPIVersion   = ParameterGroup_Kind + "." + CRDGroupVersion.String()
	ParameterGroup_GroupVersionKind = CRDGroupVersion.WithKind(ParameterGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ParameterGroup{}, &ParameterGroupList{})
}
