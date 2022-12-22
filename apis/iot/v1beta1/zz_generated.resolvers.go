/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/sns/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this LoggingOptions.
func (mg *LoggingOptions) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PolicyAttachment.
func (mg *PolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Policy),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyRef,
		Selector:     mg.Spec.ForProvider.PolicySelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Policy")
	}
	mg.Spec.ForProvider.Policy = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Target),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.TargetRef,
		Selector:     mg.Spec.ForProvider.TargetSelector,
		To: reference.To{
			List:    &CertificateList{},
			Managed: &Certificate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Target")
	}
	mg.Spec.ForProvider.Target = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ProvisioningTemplate.
func (mg *ProvisioningTemplate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProvisioningRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ProvisioningRoleArnRef,
		Selector:     mg.Spec.ForProvider.ProvisioningRoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProvisioningRoleArn")
	}
	mg.Spec.ForProvider.ProvisioningRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProvisioningRoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RoleAlias.
func (mg *RoleAlias) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ThingGroup.
func (mg *ThingGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ParentGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ParentGroupNameRef,
		Selector:     mg.Spec.ForProvider.ParentGroupNameSelector,
		To: reference.To{
			List:    &ThingGroupList{},
			Managed: &ThingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ParentGroupName")
	}
	mg.Spec.ForProvider.ParentGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ThingPrincipalAttachment.
func (mg *ThingPrincipalAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Principal),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.PrincipalRef,
		Selector:     mg.Spec.ForProvider.PrincipalSelector,
		To: reference.To{
			List:    &CertificateList{},
			Managed: &Certificate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Principal")
	}
	mg.Spec.ForProvider.Principal = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrincipalRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Thing),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ThingRef,
		Selector:     mg.Spec.ForProvider.ThingSelector,
		To: reference.To{
			List:    &ThingList{},
			Managed: &Thing{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Thing")
	}
	mg.Spec.ForProvider.Thing = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ThingRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TopicRule.
func (mg *TopicRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ErrorAction); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ErrorAction[i3].Sns); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].RoleArnRef,
				Selector:     mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].RoleArnSelector,
				To: reference.To{
					List:    &v1beta1.RoleList{},
					Managed: &v1beta1.Role{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].RoleArn")
			}
			mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ErrorAction); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ErrorAction[i3].Sns); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].TargetArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].TargetArnRef,
				Selector:     mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].TargetArnSelector,
				To: reference.To{
					List:    &v1beta11.TopicList{},
					Managed: &v1beta11.Topic{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].TargetArn")
			}
			mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].TargetArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].TargetArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Sns); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Sns[i3].RoleArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Sns[i3].RoleArnRef,
			Selector:     mg.Spec.ForProvider.Sns[i3].RoleArnSelector,
			To: reference.To{
				List:    &v1beta1.RoleList{},
				Managed: &v1beta1.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Sns[i3].RoleArn")
		}
		mg.Spec.ForProvider.Sns[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Sns[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Sns); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Sns[i3].TargetArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Sns[i3].TargetArnRef,
			Selector:     mg.Spec.ForProvider.Sns[i3].TargetArnSelector,
			To: reference.To{
				List:    &v1beta11.TopicList{},
				Managed: &v1beta11.Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Sns[i3].TargetArn")
		}
		mg.Spec.ForProvider.Sns[i3].TargetArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Sns[i3].TargetArnRef = rsp.ResolvedReference

	}

	return nil
}
