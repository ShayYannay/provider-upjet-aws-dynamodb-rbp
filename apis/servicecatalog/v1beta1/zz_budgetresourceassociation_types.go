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

type BudgetResourceAssociationInitParameters struct {

	// Budget name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/budgets/v1beta1.Budget
	BudgetName *string `json:"budgetName,omitempty" tf:"budget_name,omitempty"`

	// Reference to a Budget in budgets to populate budgetName.
	// +kubebuilder:validation:Optional
	BudgetNameRef *v1.Reference `json:"budgetNameRef,omitempty" tf:"-"`

	// Selector for a Budget in budgets to populate budgetName.
	// +kubebuilder:validation:Optional
	BudgetNameSelector *v1.Selector `json:"budgetNameSelector,omitempty" tf:"-"`

	// Resource identifier.
	// +crossplane:generate:reference:type=Product
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Product to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Product to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`
}

type BudgetResourceAssociationObservation struct {

	// Budget name.
	BudgetName *string `json:"budgetName,omitempty" tf:"budget_name,omitempty"`

	// Identifier of the association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource identifier.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type BudgetResourceAssociationParameters struct {

	// Budget name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/budgets/v1beta1.Budget
	// +kubebuilder:validation:Optional
	BudgetName *string `json:"budgetName,omitempty" tf:"budget_name,omitempty"`

	// Reference to a Budget in budgets to populate budgetName.
	// +kubebuilder:validation:Optional
	BudgetNameRef *v1.Reference `json:"budgetNameRef,omitempty" tf:"-"`

	// Selector for a Budget in budgets to populate budgetName.
	// +kubebuilder:validation:Optional
	BudgetNameSelector *v1.Selector `json:"budgetNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Resource identifier.
	// +crossplane:generate:reference:type=Product
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Product to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Product to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`
}

// BudgetResourceAssociationSpec defines the desired state of BudgetResourceAssociation
type BudgetResourceAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BudgetResourceAssociationParameters `json:"forProvider"`
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
	InitProvider BudgetResourceAssociationInitParameters `json:"initProvider,omitempty"`
}

// BudgetResourceAssociationStatus defines the observed state of BudgetResourceAssociation.
type BudgetResourceAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BudgetResourceAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BudgetResourceAssociation is the Schema for the BudgetResourceAssociations API. Manages a Service Catalog Budget Resource Association
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BudgetResourceAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BudgetResourceAssociationSpec   `json:"spec"`
	Status            BudgetResourceAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetResourceAssociationList contains a list of BudgetResourceAssociations
type BudgetResourceAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BudgetResourceAssociation `json:"items"`
}

// Repository type metadata.
var (
	BudgetResourceAssociation_Kind             = "BudgetResourceAssociation"
	BudgetResourceAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BudgetResourceAssociation_Kind}.String()
	BudgetResourceAssociation_KindAPIVersion   = BudgetResourceAssociation_Kind + "." + CRDGroupVersion.String()
	BudgetResourceAssociation_GroupVersionKind = CRDGroupVersion.WithKind(BudgetResourceAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&BudgetResourceAssociation{}, &BudgetResourceAssociationList{})
}
