// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data source for managing VPC IPAMs.
//
// ## Example Usage
//
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.GetVpcIpams(ctx, &ec2.GetVpcIpamsArgs{
//				IpamIds: []string{
//					"ipam-abcd1234",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### Filter by `tags`
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.GetVpcIpams(ctx, &ec2.GetVpcIpamsArgs{
//				Filters: []ec2.GetVpcIpamsFilter{
//					{
//						Name: "tags.Some",
//						Values: []string{
//							"Value",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### Filter by `tier`
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.GetVpcIpams(ctx, &ec2.GetVpcIpamsArgs{
//				Filters: []ec2.GetVpcIpamsFilter{
//					{
//						Name: "tier",
//						Values: []string{
//							"free",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetVpcIpams(ctx *pulumi.Context, args *GetVpcIpamsArgs, opts ...pulumi.InvokeOption) (*GetVpcIpamsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVpcIpamsResult
	err := ctx.Invoke("aws:ec2/getVpcIpams:getVpcIpams", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcIpams.
type GetVpcIpamsArgs struct {
	// Custom filter block as described below.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Filters []GetVpcIpamsFilter `pulumi:"filters"`
	// IDs of the IPAM resources to query for.
	IpamIds []string `pulumi:"ipamIds"`
}

// A collection of values returned by getVpcIpams.
type GetVpcIpamsResult struct {
	Filters []GetVpcIpamsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id      string   `pulumi:"id"`
	IpamIds []string `pulumi:"ipamIds"`
	// List of IPAM resources matching the provided arguments.
	Ipams []GetVpcIpamsIpam `pulumi:"ipams"`
}

func GetVpcIpamsOutput(ctx *pulumi.Context, args GetVpcIpamsOutputArgs, opts ...pulumi.InvokeOption) GetVpcIpamsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetVpcIpamsResultOutput, error) {
			args := v.(GetVpcIpamsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws:ec2/getVpcIpams:getVpcIpams", args, GetVpcIpamsResultOutput{}, options).(GetVpcIpamsResultOutput), nil
		}).(GetVpcIpamsResultOutput)
}

// A collection of arguments for invoking getVpcIpams.
type GetVpcIpamsOutputArgs struct {
	// Custom filter block as described below.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Filters GetVpcIpamsFilterArrayInput `pulumi:"filters"`
	// IDs of the IPAM resources to query for.
	IpamIds pulumi.StringArrayInput `pulumi:"ipamIds"`
}

func (GetVpcIpamsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcIpamsArgs)(nil)).Elem()
}

// A collection of values returned by getVpcIpams.
type GetVpcIpamsResultOutput struct{ *pulumi.OutputState }

func (GetVpcIpamsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcIpamsResult)(nil)).Elem()
}

func (o GetVpcIpamsResultOutput) ToGetVpcIpamsResultOutput() GetVpcIpamsResultOutput {
	return o
}

func (o GetVpcIpamsResultOutput) ToGetVpcIpamsResultOutputWithContext(ctx context.Context) GetVpcIpamsResultOutput {
	return o
}

func (o GetVpcIpamsResultOutput) Filters() GetVpcIpamsFilterArrayOutput {
	return o.ApplyT(func(v GetVpcIpamsResult) []GetVpcIpamsFilter { return v.Filters }).(GetVpcIpamsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpcIpamsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcIpamsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVpcIpamsResultOutput) IpamIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpcIpamsResult) []string { return v.IpamIds }).(pulumi.StringArrayOutput)
}

// List of IPAM resources matching the provided arguments.
func (o GetVpcIpamsResultOutput) Ipams() GetVpcIpamsIpamArrayOutput {
	return o.ApplyT(func(v GetVpcIpamsResult) []GetVpcIpamsIpam { return v.Ipams }).(GetVpcIpamsIpamArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpcIpamsResultOutput{})
}
