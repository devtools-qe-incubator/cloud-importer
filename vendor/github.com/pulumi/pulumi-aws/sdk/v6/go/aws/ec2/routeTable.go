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

// Provides a resource to create a VPC routing table.
//
// > **NOTE on Route Tables and Routes:** This provider currently
// provides both a standalone Route resource and a Route Table resource with routes
// defined in-line. At this time you cannot use a Route Table with in-line routes
// in conjunction with any Route resources. Doing so will cause
// a conflict of rule settings and will overwrite rules.
//
// > **NOTE on `gatewayId` and `natGatewayId`:** The AWS API is very forgiving with these two
// attributes and the `ec2.RouteTable` resource can be created with a NAT ID specified as a Gateway ID attribute.
// This _will_ lead to a permanent diff between your configuration and statefile, as the API returns the correct
// parameters in the returned route table. If you're experiencing constant diffs in your `ec2.RouteTable` resources,
// the first thing to check is whether or not you're specifying a NAT ID instead of a Gateway ID, or vice-versa.
//
// > **NOTE on `propagatingVgws` and the `ec2.VpnGatewayRoutePropagation` resource:**
// If the `propagatingVgws` argument is present, it's not supported to _also_
// define route propagations using `ec2.VpnGatewayRoutePropagation`, since
// this resource will delete any propagating gateways not explicitly listed in
// `propagatingVgws`. Omit this argument when defining route propagation using
// the separate resource.
//
// ## Example Usage
//
// ### Basic example
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
//			_, err := ec2.NewRouteTable(ctx, "example", &ec2.RouteTableArgs{
//				VpcId: pulumi.Any(exampleAwsVpc.Id),
//				Routes: ec2.RouteTableRouteArray{
//					&ec2.RouteTableRouteArgs{
//						CidrBlock: pulumi.String("10.0.1.0/24"),
//						GatewayId: pulumi.Any(exampleAwsInternetGateway.Id),
//					},
//					&ec2.RouteTableRouteArgs{
//						Ipv6CidrBlock:       pulumi.String("::/0"),
//						EgressOnlyGatewayId: pulumi.Any(exampleAwsEgressOnlyInternetGateway.Id),
//					},
//				},
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("example"),
//				},
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
// To subsequently remove all managed routes:
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
//			_, err := ec2.NewRouteTable(ctx, "example", &ec2.RouteTableArgs{
//				VpcId:  pulumi.Any(exampleAwsVpc.Id),
//				Routes: ec2.RouteTableRouteArray{},
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("example"),
//				},
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
// ### Adopting an existing local route
//
// AWS creates certain routes that the AWS provider mostly ignores. You can manage them by importing or adopting them. See Import below for information on importing. This example shows adopting a route and then updating its target.
//
// First, adopt an existing AWS-created route:
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
//			test, err := ec2.NewVpc(ctx, "test", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.1.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewRouteTable(ctx, "test", &ec2.RouteTableArgs{
//				VpcId: test.ID(),
//				Routes: ec2.RouteTableRouteArray{
//					&ec2.RouteTableRouteArgs{
//						CidrBlock: pulumi.String("10.1.0.0/16"),
//						GatewayId: pulumi.String("local"),
//					},
//				},
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
// Next, update the target of the route:
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
//			test, err := ec2.NewVpc(ctx, "test", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.1.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			testSubnet, err := ec2.NewSubnet(ctx, "test", &ec2.SubnetArgs{
//				CidrBlock: pulumi.String("10.1.1.0/24"),
//				VpcId:     test.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			testNetworkInterface, err := ec2.NewNetworkInterface(ctx, "test", &ec2.NetworkInterfaceArgs{
//				SubnetId: testSubnet.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewRouteTable(ctx, "test", &ec2.RouteTableArgs{
//				VpcId: test.ID(),
//				Routes: ec2.RouteTableRouteArray{
//					&ec2.RouteTableRouteArgs{
//						CidrBlock:          test.CidrBlock,
//						NetworkInterfaceId: testNetworkInterface.ID(),
//					},
//				},
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
// The target could then be updated again back to `local`.
//
// ## Import
//
// Using `pulumi import`, import Route Tables using the route table `id`. For example:
//
// ```sh
// $ pulumi import aws:ec2/routeTable:RouteTable public_rt rtb-4e616f6d69
// ```
type RouteTable struct {
	pulumi.CustomResourceState

	// The ARN of the route table.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the AWS account that owns the route table.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.StringArrayOutput `pulumi:"propagatingVgws"`
	// A list of route objects. Their keys are documented below.
	// This means that omitting this argument is interpreted as ignoring any existing routes. To remove all managed routes an empty list should be specified. See the example above.
	Routes RouteTableRouteArrayOutput `pulumi:"routes"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewRouteTable registers a new resource with the given unique name, arguments, and options.
func NewRouteTable(ctx *pulumi.Context,
	name string, args *RouteTableArgs, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouteTable
	err := ctx.RegisterResource("aws:ec2/routeTable:RouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTable gets an existing RouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteTableState, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	var resource RouteTable
	err := ctx.ReadResource("aws:ec2/routeTable:RouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteTable resources.
type routeTableState struct {
	// The ARN of the route table.
	Arn *string `pulumi:"arn"`
	// The ID of the AWS account that owns the route table.
	OwnerId *string `pulumi:"ownerId"`
	// A list of virtual gateways for propagation.
	PropagatingVgws []string `pulumi:"propagatingVgws"`
	// A list of route objects. Their keys are documented below.
	// This means that omitting this argument is interpreted as ignoring any existing routes. To remove all managed routes an empty list should be specified. See the example above.
	Routes []RouteTableRoute `pulumi:"routes"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

type RouteTableState struct {
	// The ARN of the route table.
	Arn pulumi.StringPtrInput
	// The ID of the AWS account that owns the route table.
	OwnerId pulumi.StringPtrInput
	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.StringArrayInput
	// A list of route objects. Their keys are documented below.
	// This means that omitting this argument is interpreted as ignoring any existing routes. To remove all managed routes an empty list should be specified. See the example above.
	Routes RouteTableRouteArrayInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The VPC ID.
	VpcId pulumi.StringPtrInput
}

func (RouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableState)(nil)).Elem()
}

type routeTableArgs struct {
	// A list of virtual gateways for propagation.
	PropagatingVgws []string `pulumi:"propagatingVgws"`
	// A list of route objects. Their keys are documented below.
	// This means that omitting this argument is interpreted as ignoring any existing routes. To remove all managed routes an empty list should be specified. See the example above.
	Routes []RouteTableRoute `pulumi:"routes"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The VPC ID.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a RouteTable resource.
type RouteTableArgs struct {
	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.StringArrayInput
	// A list of route objects. Their keys are documented below.
	// This means that omitting this argument is interpreted as ignoring any existing routes. To remove all managed routes an empty list should be specified. See the example above.
	Routes RouteTableRouteArrayInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The VPC ID.
	VpcId pulumi.StringInput
}

func (RouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableArgs)(nil)).Elem()
}

type RouteTableInput interface {
	pulumi.Input

	ToRouteTableOutput() RouteTableOutput
	ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput
}

func (*RouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil)).Elem()
}

func (i *RouteTable) ToRouteTableOutput() RouteTableOutput {
	return i.ToRouteTableOutputWithContext(context.Background())
}

func (i *RouteTable) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableOutput)
}

// RouteTableArrayInput is an input type that accepts RouteTableArray and RouteTableArrayOutput values.
// You can construct a concrete instance of `RouteTableArrayInput` via:
//
//	RouteTableArray{ RouteTableArgs{...} }
type RouteTableArrayInput interface {
	pulumi.Input

	ToRouteTableArrayOutput() RouteTableArrayOutput
	ToRouteTableArrayOutputWithContext(context.Context) RouteTableArrayOutput
}

type RouteTableArray []RouteTableInput

func (RouteTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteTable)(nil)).Elem()
}

func (i RouteTableArray) ToRouteTableArrayOutput() RouteTableArrayOutput {
	return i.ToRouteTableArrayOutputWithContext(context.Background())
}

func (i RouteTableArray) ToRouteTableArrayOutputWithContext(ctx context.Context) RouteTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableArrayOutput)
}

// RouteTableMapInput is an input type that accepts RouteTableMap and RouteTableMapOutput values.
// You can construct a concrete instance of `RouteTableMapInput` via:
//
//	RouteTableMap{ "key": RouteTableArgs{...} }
type RouteTableMapInput interface {
	pulumi.Input

	ToRouteTableMapOutput() RouteTableMapOutput
	ToRouteTableMapOutputWithContext(context.Context) RouteTableMapOutput
}

type RouteTableMap map[string]RouteTableInput

func (RouteTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteTable)(nil)).Elem()
}

func (i RouteTableMap) ToRouteTableMapOutput() RouteTableMapOutput {
	return i.ToRouteTableMapOutputWithContext(context.Background())
}

func (i RouteTableMap) ToRouteTableMapOutputWithContext(ctx context.Context) RouteTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableMapOutput)
}

type RouteTableOutput struct{ *pulumi.OutputState }

func (RouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil)).Elem()
}

func (o RouteTableOutput) ToRouteTableOutput() RouteTableOutput {
	return o
}

func (o RouteTableOutput) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return o
}

// The ARN of the route table.
func (o RouteTableOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ID of the AWS account that owns the route table.
func (o RouteTableOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// A list of virtual gateways for propagation.
func (o RouteTableOutput) PropagatingVgws() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringArrayOutput { return v.PropagatingVgws }).(pulumi.StringArrayOutput)
}

// A list of route objects. Their keys are documented below.
// This means that omitting this argument is interpreted as ignoring any existing routes. To remove all managed routes an empty list should be specified. See the example above.
func (o RouteTableOutput) Routes() RouteTableRouteArrayOutput {
	return o.ApplyT(func(v *RouteTable) RouteTableRouteArrayOutput { return v.Routes }).(RouteTableRouteArrayOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o RouteTableOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o RouteTableOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The VPC ID.
func (o RouteTableOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type RouteTableArrayOutput struct{ *pulumi.OutputState }

func (RouteTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteTable)(nil)).Elem()
}

func (o RouteTableArrayOutput) ToRouteTableArrayOutput() RouteTableArrayOutput {
	return o
}

func (o RouteTableArrayOutput) ToRouteTableArrayOutputWithContext(ctx context.Context) RouteTableArrayOutput {
	return o
}

func (o RouteTableArrayOutput) Index(i pulumi.IntInput) RouteTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouteTable {
		return vs[0].([]*RouteTable)[vs[1].(int)]
	}).(RouteTableOutput)
}

type RouteTableMapOutput struct{ *pulumi.OutputState }

func (RouteTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteTable)(nil)).Elem()
}

func (o RouteTableMapOutput) ToRouteTableMapOutput() RouteTableMapOutput {
	return o
}

func (o RouteTableMapOutput) ToRouteTableMapOutputWithContext(ctx context.Context) RouteTableMapOutput {
	return o
}

func (o RouteTableMapOutput) MapIndex(k pulumi.StringInput) RouteTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouteTable {
		return vs[0].(map[string]*RouteTable)[vs[1].(string)]
	}).(RouteTableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTableInput)(nil)).Elem(), &RouteTable{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTableArrayInput)(nil)).Elem(), RouteTableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTableMapInput)(nil)).Elem(), RouteTableMap{})
	pulumi.RegisterOutputType(RouteTableOutput{})
	pulumi.RegisterOutputType(RouteTableArrayOutput{})
	pulumi.RegisterOutputType(RouteTableMapOutput{})
}
