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

// Provides an EC2 Host resource. This allows Dedicated Hosts to be allocated, modified, and released.
//
// ## Example Usage
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
//			// Create a new host with instance type of c5.18xlarge with Auto Placement
//			// and Host Recovery enabled.
//			_, err := ec2.NewDedicatedHost(ctx, "test", &ec2.DedicatedHostArgs{
//				InstanceType:     pulumi.String("c5.18xlarge"),
//				AvailabilityZone: pulumi.String("us-west-2a"),
//				HostRecovery:     pulumi.String("on"),
//				AutoPlacement:    pulumi.String("on"),
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
// Using `pulumi import`, import hosts using the host `id`. For example:
//
// ```sh
// $ pulumi import aws:ec2/dedicatedHost:DedicatedHost example h-0385a99d0e4b20cbb
// ```
type DedicatedHost struct {
	pulumi.CustomResourceState

	// The ARN of the Dedicated Host.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
	AssetId pulumi.StringOutput `pulumi:"assetId"`
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
	AutoPlacement pulumi.StringPtrOutput `pulumi:"autoPlacement"`
	// The Availability Zone in which to allocate the Dedicated Host.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
	HostRecovery pulumi.StringPtrOutput `pulumi:"hostRecovery"`
	// Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instanceFamily` or `instanceType` must be specified.
	InstanceFamily pulumi.StringPtrOutput `pulumi:"instanceFamily"`
	// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instanceFamily` or `instanceType` must be specified.
	InstanceType pulumi.StringPtrOutput `pulumi:"instanceType"`
	// The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
	OutpostArn pulumi.StringPtrOutput `pulumi:"outpostArn"`
	// The ID of the AWS account that owns the Dedicated Host.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewDedicatedHost registers a new resource with the given unique name, arguments, and options.
func NewDedicatedHost(ctx *pulumi.Context,
	name string, args *DedicatedHostArgs, opts ...pulumi.ResourceOption) (*DedicatedHost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DedicatedHost
	err := ctx.RegisterResource("aws:ec2/dedicatedHost:DedicatedHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedHost gets an existing DedicatedHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedHostState, opts ...pulumi.ResourceOption) (*DedicatedHost, error) {
	var resource DedicatedHost
	err := ctx.ReadResource("aws:ec2/dedicatedHost:DedicatedHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedHost resources.
type dedicatedHostState struct {
	// The ARN of the Dedicated Host.
	Arn *string `pulumi:"arn"`
	// The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
	AssetId *string `pulumi:"assetId"`
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
	AutoPlacement *string `pulumi:"autoPlacement"`
	// The Availability Zone in which to allocate the Dedicated Host.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
	HostRecovery *string `pulumi:"hostRecovery"`
	// Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instanceFamily` or `instanceType` must be specified.
	InstanceFamily *string `pulumi:"instanceFamily"`
	// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instanceFamily` or `instanceType` must be specified.
	InstanceType *string `pulumi:"instanceType"`
	// The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
	OutpostArn *string `pulumi:"outpostArn"`
	// The ID of the AWS account that owns the Dedicated Host.
	OwnerId *string `pulumi:"ownerId"`
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type DedicatedHostState struct {
	// The ARN of the Dedicated Host.
	Arn pulumi.StringPtrInput
	// The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
	AssetId pulumi.StringPtrInput
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
	AutoPlacement pulumi.StringPtrInput
	// The Availability Zone in which to allocate the Dedicated Host.
	AvailabilityZone pulumi.StringPtrInput
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
	HostRecovery pulumi.StringPtrInput
	// Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instanceFamily` or `instanceType` must be specified.
	InstanceFamily pulumi.StringPtrInput
	// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instanceFamily` or `instanceType` must be specified.
	InstanceType pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
	OutpostArn pulumi.StringPtrInput
	// The ID of the AWS account that owns the Dedicated Host.
	OwnerId pulumi.StringPtrInput
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (DedicatedHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostState)(nil)).Elem()
}

type dedicatedHostArgs struct {
	// The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
	AssetId *string `pulumi:"assetId"`
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
	AutoPlacement *string `pulumi:"autoPlacement"`
	// The Availability Zone in which to allocate the Dedicated Host.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
	HostRecovery *string `pulumi:"hostRecovery"`
	// Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instanceFamily` or `instanceType` must be specified.
	InstanceFamily *string `pulumi:"instanceFamily"`
	// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instanceFamily` or `instanceType` must be specified.
	InstanceType *string `pulumi:"instanceType"`
	// The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
	OutpostArn *string `pulumi:"outpostArn"`
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DedicatedHost resource.
type DedicatedHostArgs struct {
	// The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
	AssetId pulumi.StringPtrInput
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
	AutoPlacement pulumi.StringPtrInput
	// The Availability Zone in which to allocate the Dedicated Host.
	AvailabilityZone pulumi.StringInput
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
	HostRecovery pulumi.StringPtrInput
	// Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instanceFamily` or `instanceType` must be specified.
	InstanceFamily pulumi.StringPtrInput
	// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instanceFamily` or `instanceType` must be specified.
	InstanceType pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
	OutpostArn pulumi.StringPtrInput
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (DedicatedHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostArgs)(nil)).Elem()
}

type DedicatedHostInput interface {
	pulumi.Input

	ToDedicatedHostOutput() DedicatedHostOutput
	ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput
}

func (*DedicatedHost) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHost)(nil)).Elem()
}

func (i *DedicatedHost) ToDedicatedHostOutput() DedicatedHostOutput {
	return i.ToDedicatedHostOutputWithContext(context.Background())
}

func (i *DedicatedHost) ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostOutput)
}

// DedicatedHostArrayInput is an input type that accepts DedicatedHostArray and DedicatedHostArrayOutput values.
// You can construct a concrete instance of `DedicatedHostArrayInput` via:
//
//	DedicatedHostArray{ DedicatedHostArgs{...} }
type DedicatedHostArrayInput interface {
	pulumi.Input

	ToDedicatedHostArrayOutput() DedicatedHostArrayOutput
	ToDedicatedHostArrayOutputWithContext(context.Context) DedicatedHostArrayOutput
}

type DedicatedHostArray []DedicatedHostInput

func (DedicatedHostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedHost)(nil)).Elem()
}

func (i DedicatedHostArray) ToDedicatedHostArrayOutput() DedicatedHostArrayOutput {
	return i.ToDedicatedHostArrayOutputWithContext(context.Background())
}

func (i DedicatedHostArray) ToDedicatedHostArrayOutputWithContext(ctx context.Context) DedicatedHostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostArrayOutput)
}

// DedicatedHostMapInput is an input type that accepts DedicatedHostMap and DedicatedHostMapOutput values.
// You can construct a concrete instance of `DedicatedHostMapInput` via:
//
//	DedicatedHostMap{ "key": DedicatedHostArgs{...} }
type DedicatedHostMapInput interface {
	pulumi.Input

	ToDedicatedHostMapOutput() DedicatedHostMapOutput
	ToDedicatedHostMapOutputWithContext(context.Context) DedicatedHostMapOutput
}

type DedicatedHostMap map[string]DedicatedHostInput

func (DedicatedHostMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedHost)(nil)).Elem()
}

func (i DedicatedHostMap) ToDedicatedHostMapOutput() DedicatedHostMapOutput {
	return i.ToDedicatedHostMapOutputWithContext(context.Background())
}

func (i DedicatedHostMap) ToDedicatedHostMapOutputWithContext(ctx context.Context) DedicatedHostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostMapOutput)
}

type DedicatedHostOutput struct{ *pulumi.OutputState }

func (DedicatedHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHost)(nil)).Elem()
}

func (o DedicatedHostOutput) ToDedicatedHostOutput() DedicatedHostOutput {
	return o
}

func (o DedicatedHostOutput) ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput {
	return o
}

// The ARN of the Dedicated Host.
func (o DedicatedHostOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
func (o DedicatedHostOutput) AssetId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.AssetId }).(pulumi.StringOutput)
}

// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
func (o DedicatedHostOutput) AutoPlacement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringPtrOutput { return v.AutoPlacement }).(pulumi.StringPtrOutput)
}

// The Availability Zone in which to allocate the Dedicated Host.
func (o DedicatedHostOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
func (o DedicatedHostOutput) HostRecovery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringPtrOutput { return v.HostRecovery }).(pulumi.StringPtrOutput)
}

// Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instanceFamily` or `instanceType` must be specified.
func (o DedicatedHostOutput) InstanceFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringPtrOutput { return v.InstanceFamily }).(pulumi.StringPtrOutput)
}

// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instanceFamily` or `instanceType` must be specified.
func (o DedicatedHostOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringPtrOutput { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
func (o DedicatedHostOutput) OutpostArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringPtrOutput { return v.OutpostArn }).(pulumi.StringPtrOutput)
}

// The ID of the AWS account that owns the Dedicated Host.
func (o DedicatedHostOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o DedicatedHostOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o DedicatedHostOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type DedicatedHostArrayOutput struct{ *pulumi.OutputState }

func (DedicatedHostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedHost)(nil)).Elem()
}

func (o DedicatedHostArrayOutput) ToDedicatedHostArrayOutput() DedicatedHostArrayOutput {
	return o
}

func (o DedicatedHostArrayOutput) ToDedicatedHostArrayOutputWithContext(ctx context.Context) DedicatedHostArrayOutput {
	return o
}

func (o DedicatedHostArrayOutput) Index(i pulumi.IntInput) DedicatedHostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DedicatedHost {
		return vs[0].([]*DedicatedHost)[vs[1].(int)]
	}).(DedicatedHostOutput)
}

type DedicatedHostMapOutput struct{ *pulumi.OutputState }

func (DedicatedHostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedHost)(nil)).Elem()
}

func (o DedicatedHostMapOutput) ToDedicatedHostMapOutput() DedicatedHostMapOutput {
	return o
}

func (o DedicatedHostMapOutput) ToDedicatedHostMapOutputWithContext(ctx context.Context) DedicatedHostMapOutput {
	return o
}

func (o DedicatedHostMapOutput) MapIndex(k pulumi.StringInput) DedicatedHostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DedicatedHost {
		return vs[0].(map[string]*DedicatedHost)[vs[1].(string)]
	}).(DedicatedHostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedHostInput)(nil)).Elem(), &DedicatedHost{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedHostArrayInput)(nil)).Elem(), DedicatedHostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedHostMapInput)(nil)).Elem(), DedicatedHostMap{})
	pulumi.RegisterOutputType(DedicatedHostOutput{})
	pulumi.RegisterOutputType(DedicatedHostArrayOutput{})
	pulumi.RegisterOutputType(DedicatedHostMapOutput{})
}
