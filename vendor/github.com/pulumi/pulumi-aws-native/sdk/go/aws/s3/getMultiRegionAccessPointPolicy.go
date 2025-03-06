// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The policy to be attached to a Multi Region Access Point
func LookupMultiRegionAccessPointPolicy(ctx *pulumi.Context, args *LookupMultiRegionAccessPointPolicyArgs, opts ...pulumi.InvokeOption) (*LookupMultiRegionAccessPointPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMultiRegionAccessPointPolicyResult
	err := ctx.Invoke("aws-native:s3:getMultiRegionAccessPointPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMultiRegionAccessPointPolicyArgs struct {
	// The name of the Multi Region Access Point to apply policy
	MrapName string `pulumi:"mrapName"`
}

type LookupMultiRegionAccessPointPolicyResult struct {
	// Policy document to apply to a Multi Region Access Point
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3::MultiRegionAccessPointPolicy` for more information about the expected schema for this property.
	Policy interface{} `pulumi:"policy"`
	// The Policy Status associated with this Multi Region Access Point
	PolicyStatus *PolicyStatusProperties `pulumi:"policyStatus"`
}

func LookupMultiRegionAccessPointPolicyOutput(ctx *pulumi.Context, args LookupMultiRegionAccessPointPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupMultiRegionAccessPointPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMultiRegionAccessPointPolicyResultOutput, error) {
			args := v.(LookupMultiRegionAccessPointPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:s3:getMultiRegionAccessPointPolicy", args, LookupMultiRegionAccessPointPolicyResultOutput{}, options).(LookupMultiRegionAccessPointPolicyResultOutput), nil
		}).(LookupMultiRegionAccessPointPolicyResultOutput)
}

type LookupMultiRegionAccessPointPolicyOutputArgs struct {
	// The name of the Multi Region Access Point to apply policy
	MrapName pulumi.StringInput `pulumi:"mrapName"`
}

func (LookupMultiRegionAccessPointPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMultiRegionAccessPointPolicyArgs)(nil)).Elem()
}

type LookupMultiRegionAccessPointPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupMultiRegionAccessPointPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMultiRegionAccessPointPolicyResult)(nil)).Elem()
}

func (o LookupMultiRegionAccessPointPolicyResultOutput) ToLookupMultiRegionAccessPointPolicyResultOutput() LookupMultiRegionAccessPointPolicyResultOutput {
	return o
}

func (o LookupMultiRegionAccessPointPolicyResultOutput) ToLookupMultiRegionAccessPointPolicyResultOutputWithContext(ctx context.Context) LookupMultiRegionAccessPointPolicyResultOutput {
	return o
}

// Policy document to apply to a Multi Region Access Point
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3::MultiRegionAccessPointPolicy` for more information about the expected schema for this property.
func (o LookupMultiRegionAccessPointPolicyResultOutput) Policy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupMultiRegionAccessPointPolicyResult) interface{} { return v.Policy }).(pulumi.AnyOutput)
}

// The Policy Status associated with this Multi Region Access Point
func (o LookupMultiRegionAccessPointPolicyResultOutput) PolicyStatus() PolicyStatusPropertiesPtrOutput {
	return o.ApplyT(func(v LookupMultiRegionAccessPointPolicyResult) *PolicyStatusProperties { return v.PolicyStatus }).(PolicyStatusPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMultiRegionAccessPointPolicyResultOutput{})
}
