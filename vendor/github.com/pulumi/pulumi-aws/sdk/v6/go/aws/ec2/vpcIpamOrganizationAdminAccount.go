// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enables the IPAM Service and promotes a delegated administrator.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			delegated, err := aws.GetCallerIdentity(ctx, &aws.GetCallerIdentityArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcIpamOrganizationAdminAccount(ctx, "example", &ec2.VpcIpamOrganizationAdminAccountArgs{
//				DelegatedAdminAccountId: pulumi.String(delegated.AccountId),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Using `pulumi import`, import IPAMs using the delegate account `id`. For example:
//
// ```sh
// $ pulumi import aws:ec2/vpcIpamOrganizationAdminAccount:VpcIpamOrganizationAdminAccount example 12345678901
// ```
type VpcIpamOrganizationAdminAccount struct {
	pulumi.CustomResourceState

	// The Organizations ARN for the delegate account.
	Arn                     pulumi.StringOutput `pulumi:"arn"`
	DelegatedAdminAccountId pulumi.StringOutput `pulumi:"delegatedAdminAccountId"`
	// The Organizations email for the delegate account.
	Email pulumi.StringOutput `pulumi:"email"`
	// The Organizations name for the delegate account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The AWS service principal.
	ServicePrincipal pulumi.StringOutput `pulumi:"servicePrincipal"`
}

// NewVpcIpamOrganizationAdminAccount registers a new resource with the given unique name, arguments, and options.
func NewVpcIpamOrganizationAdminAccount(ctx *pulumi.Context,
	name string, args *VpcIpamOrganizationAdminAccountArgs, opts ...pulumi.ResourceOption) (*VpcIpamOrganizationAdminAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DelegatedAdminAccountId == nil {
		return nil, errors.New("invalid value for required argument 'DelegatedAdminAccountId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcIpamOrganizationAdminAccount
	err := ctx.RegisterResource("aws:ec2/vpcIpamOrganizationAdminAccount:VpcIpamOrganizationAdminAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcIpamOrganizationAdminAccount gets an existing VpcIpamOrganizationAdminAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcIpamOrganizationAdminAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcIpamOrganizationAdminAccountState, opts ...pulumi.ResourceOption) (*VpcIpamOrganizationAdminAccount, error) {
	var resource VpcIpamOrganizationAdminAccount
	err := ctx.ReadResource("aws:ec2/vpcIpamOrganizationAdminAccount:VpcIpamOrganizationAdminAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcIpamOrganizationAdminAccount resources.
type vpcIpamOrganizationAdminAccountState struct {
	// The Organizations ARN for the delegate account.
	Arn                     *string `pulumi:"arn"`
	DelegatedAdminAccountId *string `pulumi:"delegatedAdminAccountId"`
	// The Organizations email for the delegate account.
	Email *string `pulumi:"email"`
	// The Organizations name for the delegate account.
	Name *string `pulumi:"name"`
	// The AWS service principal.
	ServicePrincipal *string `pulumi:"servicePrincipal"`
}

type VpcIpamOrganizationAdminAccountState struct {
	// The Organizations ARN for the delegate account.
	Arn                     pulumi.StringPtrInput
	DelegatedAdminAccountId pulumi.StringPtrInput
	// The Organizations email for the delegate account.
	Email pulumi.StringPtrInput
	// The Organizations name for the delegate account.
	Name pulumi.StringPtrInput
	// The AWS service principal.
	ServicePrincipal pulumi.StringPtrInput
}

func (VpcIpamOrganizationAdminAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamOrganizationAdminAccountState)(nil)).Elem()
}

type vpcIpamOrganizationAdminAccountArgs struct {
	DelegatedAdminAccountId string `pulumi:"delegatedAdminAccountId"`
}

// The set of arguments for constructing a VpcIpamOrganizationAdminAccount resource.
type VpcIpamOrganizationAdminAccountArgs struct {
	DelegatedAdminAccountId pulumi.StringInput
}

func (VpcIpamOrganizationAdminAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamOrganizationAdminAccountArgs)(nil)).Elem()
}

type VpcIpamOrganizationAdminAccountInput interface {
	pulumi.Input

	ToVpcIpamOrganizationAdminAccountOutput() VpcIpamOrganizationAdminAccountOutput
	ToVpcIpamOrganizationAdminAccountOutputWithContext(ctx context.Context) VpcIpamOrganizationAdminAccountOutput
}

func (*VpcIpamOrganizationAdminAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpamOrganizationAdminAccount)(nil)).Elem()
}

func (i *VpcIpamOrganizationAdminAccount) ToVpcIpamOrganizationAdminAccountOutput() VpcIpamOrganizationAdminAccountOutput {
	return i.ToVpcIpamOrganizationAdminAccountOutputWithContext(context.Background())
}

func (i *VpcIpamOrganizationAdminAccount) ToVpcIpamOrganizationAdminAccountOutputWithContext(ctx context.Context) VpcIpamOrganizationAdminAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamOrganizationAdminAccountOutput)
}

// VpcIpamOrganizationAdminAccountArrayInput is an input type that accepts VpcIpamOrganizationAdminAccountArray and VpcIpamOrganizationAdminAccountArrayOutput values.
// You can construct a concrete instance of `VpcIpamOrganizationAdminAccountArrayInput` via:
//
//	VpcIpamOrganizationAdminAccountArray{ VpcIpamOrganizationAdminAccountArgs{...} }
type VpcIpamOrganizationAdminAccountArrayInput interface {
	pulumi.Input

	ToVpcIpamOrganizationAdminAccountArrayOutput() VpcIpamOrganizationAdminAccountArrayOutput
	ToVpcIpamOrganizationAdminAccountArrayOutputWithContext(context.Context) VpcIpamOrganizationAdminAccountArrayOutput
}

type VpcIpamOrganizationAdminAccountArray []VpcIpamOrganizationAdminAccountInput

func (VpcIpamOrganizationAdminAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpamOrganizationAdminAccount)(nil)).Elem()
}

func (i VpcIpamOrganizationAdminAccountArray) ToVpcIpamOrganizationAdminAccountArrayOutput() VpcIpamOrganizationAdminAccountArrayOutput {
	return i.ToVpcIpamOrganizationAdminAccountArrayOutputWithContext(context.Background())
}

func (i VpcIpamOrganizationAdminAccountArray) ToVpcIpamOrganizationAdminAccountArrayOutputWithContext(ctx context.Context) VpcIpamOrganizationAdminAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamOrganizationAdminAccountArrayOutput)
}

// VpcIpamOrganizationAdminAccountMapInput is an input type that accepts VpcIpamOrganizationAdminAccountMap and VpcIpamOrganizationAdminAccountMapOutput values.
// You can construct a concrete instance of `VpcIpamOrganizationAdminAccountMapInput` via:
//
//	VpcIpamOrganizationAdminAccountMap{ "key": VpcIpamOrganizationAdminAccountArgs{...} }
type VpcIpamOrganizationAdminAccountMapInput interface {
	pulumi.Input

	ToVpcIpamOrganizationAdminAccountMapOutput() VpcIpamOrganizationAdminAccountMapOutput
	ToVpcIpamOrganizationAdminAccountMapOutputWithContext(context.Context) VpcIpamOrganizationAdminAccountMapOutput
}

type VpcIpamOrganizationAdminAccountMap map[string]VpcIpamOrganizationAdminAccountInput

func (VpcIpamOrganizationAdminAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpamOrganizationAdminAccount)(nil)).Elem()
}

func (i VpcIpamOrganizationAdminAccountMap) ToVpcIpamOrganizationAdminAccountMapOutput() VpcIpamOrganizationAdminAccountMapOutput {
	return i.ToVpcIpamOrganizationAdminAccountMapOutputWithContext(context.Background())
}

func (i VpcIpamOrganizationAdminAccountMap) ToVpcIpamOrganizationAdminAccountMapOutputWithContext(ctx context.Context) VpcIpamOrganizationAdminAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamOrganizationAdminAccountMapOutput)
}

type VpcIpamOrganizationAdminAccountOutput struct{ *pulumi.OutputState }

func (VpcIpamOrganizationAdminAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpamOrganizationAdminAccount)(nil)).Elem()
}

func (o VpcIpamOrganizationAdminAccountOutput) ToVpcIpamOrganizationAdminAccountOutput() VpcIpamOrganizationAdminAccountOutput {
	return o
}

func (o VpcIpamOrganizationAdminAccountOutput) ToVpcIpamOrganizationAdminAccountOutputWithContext(ctx context.Context) VpcIpamOrganizationAdminAccountOutput {
	return o
}

// The Organizations ARN for the delegate account.
func (o VpcIpamOrganizationAdminAccountOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamOrganizationAdminAccount) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o VpcIpamOrganizationAdminAccountOutput) DelegatedAdminAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamOrganizationAdminAccount) pulumi.StringOutput { return v.DelegatedAdminAccountId }).(pulumi.StringOutput)
}

// The Organizations email for the delegate account.
func (o VpcIpamOrganizationAdminAccountOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamOrganizationAdminAccount) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The Organizations name for the delegate account.
func (o VpcIpamOrganizationAdminAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamOrganizationAdminAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The AWS service principal.
func (o VpcIpamOrganizationAdminAccountOutput) ServicePrincipal() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamOrganizationAdminAccount) pulumi.StringOutput { return v.ServicePrincipal }).(pulumi.StringOutput)
}

type VpcIpamOrganizationAdminAccountArrayOutput struct{ *pulumi.OutputState }

func (VpcIpamOrganizationAdminAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpamOrganizationAdminAccount)(nil)).Elem()
}

func (o VpcIpamOrganizationAdminAccountArrayOutput) ToVpcIpamOrganizationAdminAccountArrayOutput() VpcIpamOrganizationAdminAccountArrayOutput {
	return o
}

func (o VpcIpamOrganizationAdminAccountArrayOutput) ToVpcIpamOrganizationAdminAccountArrayOutputWithContext(ctx context.Context) VpcIpamOrganizationAdminAccountArrayOutput {
	return o
}

func (o VpcIpamOrganizationAdminAccountArrayOutput) Index(i pulumi.IntInput) VpcIpamOrganizationAdminAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcIpamOrganizationAdminAccount {
		return vs[0].([]*VpcIpamOrganizationAdminAccount)[vs[1].(int)]
	}).(VpcIpamOrganizationAdminAccountOutput)
}

type VpcIpamOrganizationAdminAccountMapOutput struct{ *pulumi.OutputState }

func (VpcIpamOrganizationAdminAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpamOrganizationAdminAccount)(nil)).Elem()
}

func (o VpcIpamOrganizationAdminAccountMapOutput) ToVpcIpamOrganizationAdminAccountMapOutput() VpcIpamOrganizationAdminAccountMapOutput {
	return o
}

func (o VpcIpamOrganizationAdminAccountMapOutput) ToVpcIpamOrganizationAdminAccountMapOutputWithContext(ctx context.Context) VpcIpamOrganizationAdminAccountMapOutput {
	return o
}

func (o VpcIpamOrganizationAdminAccountMapOutput) MapIndex(k pulumi.StringInput) VpcIpamOrganizationAdminAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcIpamOrganizationAdminAccount {
		return vs[0].(map[string]*VpcIpamOrganizationAdminAccount)[vs[1].(string)]
	}).(VpcIpamOrganizationAdminAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamOrganizationAdminAccountInput)(nil)).Elem(), &VpcIpamOrganizationAdminAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamOrganizationAdminAccountArrayInput)(nil)).Elem(), VpcIpamOrganizationAdminAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamOrganizationAdminAccountMapInput)(nil)).Elem(), VpcIpamOrganizationAdminAccountMap{})
	pulumi.RegisterOutputType(VpcIpamOrganizationAdminAccountOutput{})
	pulumi.RegisterOutputType(VpcIpamOrganizationAdminAccountArrayOutput{})
	pulumi.RegisterOutputType(VpcIpamOrganizationAdminAccountMapOutput{})
}
