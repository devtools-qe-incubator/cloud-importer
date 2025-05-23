// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::OIDCProvider
func LookupOidcProvider(ctx *pulumi.Context, args *LookupOidcProviderArgs, opts ...pulumi.InvokeOption) (*LookupOidcProviderResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOidcProviderResult
	err := ctx.Invoke("aws-native:iam:getOidcProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOidcProviderArgs struct {
	// Amazon Resource Name (ARN) of the OIDC provider
	Arn string `pulumi:"arn"`
}

type LookupOidcProviderResult struct {
	// Amazon Resource Name (ARN) of the OIDC provider
	Arn *string `pulumi:"arn"`
	// A list of client IDs (also known as audiences) that are associated with the specified IAM OIDC provider resource object. For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html) .
	ClientIdList []string `pulumi:"clientIdList"`
	// A list of tags that are attached to the specified IAM OIDC provider. The returned list of tags is sorted by tag key. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide* .
	Tags []aws.Tag `pulumi:"tags"`
	// A list of certificate thumbprints that are associated with the specified IAM OIDC provider resource object. For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html) .
	//
	// This property is optional. If it is not included, IAM will retrieve and use the top intermediate certificate authority (CA) thumbprint of the OpenID Connect identity provider server certificate.
	ThumbprintList []string `pulumi:"thumbprintList"`
}

func LookupOidcProviderOutput(ctx *pulumi.Context, args LookupOidcProviderOutputArgs, opts ...pulumi.InvokeOption) LookupOidcProviderResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupOidcProviderResultOutput, error) {
			args := v.(LookupOidcProviderArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:iam:getOidcProvider", args, LookupOidcProviderResultOutput{}, options).(LookupOidcProviderResultOutput), nil
		}).(LookupOidcProviderResultOutput)
}

type LookupOidcProviderOutputArgs struct {
	// Amazon Resource Name (ARN) of the OIDC provider
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupOidcProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOidcProviderArgs)(nil)).Elem()
}

type LookupOidcProviderResultOutput struct{ *pulumi.OutputState }

func (LookupOidcProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOidcProviderResult)(nil)).Elem()
}

func (o LookupOidcProviderResultOutput) ToLookupOidcProviderResultOutput() LookupOidcProviderResultOutput {
	return o
}

func (o LookupOidcProviderResultOutput) ToLookupOidcProviderResultOutputWithContext(ctx context.Context) LookupOidcProviderResultOutput {
	return o
}

// Amazon Resource Name (ARN) of the OIDC provider
func (o LookupOidcProviderResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOidcProviderResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// A list of client IDs (also known as audiences) that are associated with the specified IAM OIDC provider resource object. For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html) .
func (o LookupOidcProviderResultOutput) ClientIdList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOidcProviderResult) []string { return v.ClientIdList }).(pulumi.StringArrayOutput)
}

// A list of tags that are attached to the specified IAM OIDC provider. The returned list of tags is sorted by tag key. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide* .
func (o LookupOidcProviderResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupOidcProviderResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// A list of certificate thumbprints that are associated with the specified IAM OIDC provider resource object. For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html) .
//
// This property is optional. If it is not included, IAM will retrieve and use the top intermediate certificate authority (CA) thumbprint of the OpenID Connect identity provider server certificate.
func (o LookupOidcProviderResultOutput) ThumbprintList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOidcProviderResult) []string { return v.ThumbprintList }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOidcProviderResultOutput{})
}
