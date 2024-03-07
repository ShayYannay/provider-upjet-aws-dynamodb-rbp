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

type RuleInitParameters struct {

	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The name or description of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The objects to include in a rule (documented below).
	Predicate []RulePredicateInitParameters `json:"predicate,omitempty" tf:"predicate,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RuleObservation struct {

	// The ARN of the WAF Regional Rule.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the WAF Regional Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The name or description of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The objects to include in a rule (documented below).
	Predicate []RulePredicateObservation `json:"predicate,omitempty" tf:"predicate,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RuleParameters struct {

	// The name or description for the Amazon CloudWatch metric of this rule.
	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The name or description of the rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The objects to include in a rule (documented below).
	// +kubebuilder:validation:Optional
	Predicate []RulePredicateParameters `json:"predicate,omitempty" tf:"predicate,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RulePredicateInitParameters struct {

	// The unique identifier of a predicate, such as the ID of a ByteMatchSet or IPSet.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/wafregional/v1beta1.IPSet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DataID *string `json:"dataId,omitempty" tf:"data_id,omitempty"`

	// Reference to a IPSet in wafregional to populate dataId.
	// +kubebuilder:validation:Optional
	DataIDRef *v1.Reference `json:"dataIdRef,omitempty" tf:"-"`

	// Selector for a IPSet in wafregional to populate dataId.
	// +kubebuilder:validation:Optional
	DataIDSelector *v1.Selector `json:"dataIdSelector,omitempty" tf:"-"`

	// Whether to use the settings or the negated settings that you specified in the objects.
	Negated *bool `json:"negated,omitempty" tf:"negated,omitempty"`

	// The type of predicate in a rule. Valid values: ByteMatch, GeoMatch, IPMatch, RegexMatch, SizeConstraint, SqlInjectionMatch, or XssMatch
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RulePredicateObservation struct {

	// The unique identifier of a predicate, such as the ID of a ByteMatchSet or IPSet.
	DataID *string `json:"dataId,omitempty" tf:"data_id,omitempty"`

	// Whether to use the settings or the negated settings that you specified in the objects.
	Negated *bool `json:"negated,omitempty" tf:"negated,omitempty"`

	// The type of predicate in a rule. Valid values: ByteMatch, GeoMatch, IPMatch, RegexMatch, SizeConstraint, SqlInjectionMatch, or XssMatch
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RulePredicateParameters struct {

	// The unique identifier of a predicate, such as the ID of a ByteMatchSet or IPSet.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/wafregional/v1beta1.IPSet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataID *string `json:"dataId,omitempty" tf:"data_id,omitempty"`

	// Reference to a IPSet in wafregional to populate dataId.
	// +kubebuilder:validation:Optional
	DataIDRef *v1.Reference `json:"dataIdRef,omitempty" tf:"-"`

	// Selector for a IPSet in wafregional to populate dataId.
	// +kubebuilder:validation:Optional
	DataIDSelector *v1.Selector `json:"dataIdSelector,omitempty" tf:"-"`

	// Whether to use the settings or the negated settings that you specified in the objects.
	// +kubebuilder:validation:Optional
	Negated *bool `json:"negated" tf:"negated,omitempty"`

	// The type of predicate in a rule. Valid values: ByteMatch, GeoMatch, IPMatch, RegexMatch, SizeConstraint, SqlInjectionMatch, or XssMatch
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// RuleSpec defines the desired state of Rule
type RuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleParameters `json:"forProvider"`
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
	InitProvider RuleInitParameters `json:"initProvider,omitempty"`
}

// RuleStatus defines the observed state of Rule.
type RuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Rule is the Schema for the Rules API. Provides an AWS WAF Regional rule resource for use with ALB.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metricName) || (has(self.initProvider) && has(self.initProvider.metricName))",message="spec.forProvider.metricName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RuleSpec   `json:"spec"`
	Status RuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleList contains a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

// Repository type metadata.
var (
	Rule_Kind             = "Rule"
	Rule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rule_Kind}.String()
	Rule_KindAPIVersion   = Rule_Kind + "." + CRDGroupVersion.String()
	Rule_GroupVersionKind = CRDGroupVersion.WithKind(Rule_Kind)
)

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}
