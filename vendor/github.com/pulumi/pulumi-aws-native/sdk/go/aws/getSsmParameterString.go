// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSsmParameterString(ctx *pulumi.Context, args *GetSsmParameterStringArgs, opts ...pulumi.InvokeOption) (*GetSsmParameterStringResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSsmParameterStringResult
	err := ctx.Invoke("aws-native:index:getSsmParameterString", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSsmParameterStringArgs struct {
	Name string `pulumi:"name"`
}

type GetSsmParameterStringResult struct {
	Value string `pulumi:"value"`
}

func GetSsmParameterStringOutput(ctx *pulumi.Context, args GetSsmParameterStringOutputArgs, opts ...pulumi.InvokeOption) GetSsmParameterStringResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetSsmParameterStringResultOutput, error) {
			args := v.(GetSsmParameterStringArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:index:getSsmParameterString", args, GetSsmParameterStringResultOutput{}, options).(GetSsmParameterStringResultOutput), nil
		}).(GetSsmParameterStringResultOutput)
}

type GetSsmParameterStringOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetSsmParameterStringOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSsmParameterStringArgs)(nil)).Elem()
}

type GetSsmParameterStringResultOutput struct{ *pulumi.OutputState }

func (GetSsmParameterStringResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSsmParameterStringResult)(nil)).Elem()
}

func (o GetSsmParameterStringResultOutput) ToGetSsmParameterStringResultOutput() GetSsmParameterStringResultOutput {
	return o
}

func (o GetSsmParameterStringResultOutput) ToGetSsmParameterStringResultOutputWithContext(ctx context.Context) GetSsmParameterStringResultOutput {
	return o
}

func (o GetSsmParameterStringResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetSsmParameterStringResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSsmParameterStringResultOutput{})
}
