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

// Provides a VPC Endpoint connection notification resource.
// Connection notifications notify subscribers of VPC Endpoint events.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			topic, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"vpce.amazonaws.com",
//								},
//							},
//						},
//						Actions: []string{
//							"SNS:Publish",
//						},
//						Resources: []string{
//							"arn:aws:sns:*:*:vpce-notification-topic",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			topicTopic, err := sns.NewTopic(ctx, "topic", &sns.TopicArgs{
//				Name:   pulumi.String("vpce-notification-topic"),
//				Policy: pulumi.String(topic.Json),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := ec2.NewVpcEndpointService(ctx, "foo", &ec2.VpcEndpointServiceArgs{
//				AcceptanceRequired: pulumi.Bool(false),
//				NetworkLoadBalancerArns: pulumi.StringArray{
//					test.Arn,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcEndpointConnectionNotification(ctx, "foo", &ec2.VpcEndpointConnectionNotificationArgs{
//				VpcEndpointServiceId:      foo.ID(),
//				ConnectionNotificationArn: topicTopic.Arn,
//				ConnectionEvents: pulumi.StringArray{
//					pulumi.String("Accept"),
//					pulumi.String("Reject"),
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
// ## Import
//
// Using `pulumi import`, import VPC Endpoint connection notifications using the VPC endpoint connection notification `id`. For example:
//
// ```sh
// $ pulumi import aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification foo vpce-nfn-09e6ed3b4efba2263
// ```
type VpcEndpointConnectionNotification struct {
	pulumi.CustomResourceState

	// One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
	//
	// > **NOTE:** One of `vpcEndpointServiceId` or `vpcEndpointId` must be specified.
	ConnectionEvents pulumi.StringArrayOutput `pulumi:"connectionEvents"`
	// The ARN of the SNS topic for the notifications.
	ConnectionNotificationArn pulumi.StringOutput `pulumi:"connectionNotificationArn"`
	// The type of notification.
	NotificationType pulumi.StringOutput `pulumi:"notificationType"`
	// The state of the notification.
	State pulumi.StringOutput `pulumi:"state"`
	// The ID of the VPC Endpoint to receive notifications for.
	VpcEndpointId pulumi.StringPtrOutput `pulumi:"vpcEndpointId"`
	// The ID of the VPC Endpoint Service to receive notifications for.
	VpcEndpointServiceId pulumi.StringPtrOutput `pulumi:"vpcEndpointServiceId"`
}

// NewVpcEndpointConnectionNotification registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointConnectionNotification(ctx *pulumi.Context,
	name string, args *VpcEndpointConnectionNotificationArgs, opts ...pulumi.ResourceOption) (*VpcEndpointConnectionNotification, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionEvents == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionEvents'")
	}
	if args.ConnectionNotificationArn == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionNotificationArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcEndpointConnectionNotification
	err := ctx.RegisterResource("aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointConnectionNotification gets an existing VpcEndpointConnectionNotification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointConnectionNotification(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointConnectionNotificationState, opts ...pulumi.ResourceOption) (*VpcEndpointConnectionNotification, error) {
	var resource VpcEndpointConnectionNotification
	err := ctx.ReadResource("aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointConnectionNotification resources.
type vpcEndpointConnectionNotificationState struct {
	// One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
	//
	// > **NOTE:** One of `vpcEndpointServiceId` or `vpcEndpointId` must be specified.
	ConnectionEvents []string `pulumi:"connectionEvents"`
	// The ARN of the SNS topic for the notifications.
	ConnectionNotificationArn *string `pulumi:"connectionNotificationArn"`
	// The type of notification.
	NotificationType *string `pulumi:"notificationType"`
	// The state of the notification.
	State *string `pulumi:"state"`
	// The ID of the VPC Endpoint to receive notifications for.
	VpcEndpointId *string `pulumi:"vpcEndpointId"`
	// The ID of the VPC Endpoint Service to receive notifications for.
	VpcEndpointServiceId *string `pulumi:"vpcEndpointServiceId"`
}

type VpcEndpointConnectionNotificationState struct {
	// One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
	//
	// > **NOTE:** One of `vpcEndpointServiceId` or `vpcEndpointId` must be specified.
	ConnectionEvents pulumi.StringArrayInput
	// The ARN of the SNS topic for the notifications.
	ConnectionNotificationArn pulumi.StringPtrInput
	// The type of notification.
	NotificationType pulumi.StringPtrInput
	// The state of the notification.
	State pulumi.StringPtrInput
	// The ID of the VPC Endpoint to receive notifications for.
	VpcEndpointId pulumi.StringPtrInput
	// The ID of the VPC Endpoint Service to receive notifications for.
	VpcEndpointServiceId pulumi.StringPtrInput
}

func (VpcEndpointConnectionNotificationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointConnectionNotificationState)(nil)).Elem()
}

type vpcEndpointConnectionNotificationArgs struct {
	// One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
	//
	// > **NOTE:** One of `vpcEndpointServiceId` or `vpcEndpointId` must be specified.
	ConnectionEvents []string `pulumi:"connectionEvents"`
	// The ARN of the SNS topic for the notifications.
	ConnectionNotificationArn string `pulumi:"connectionNotificationArn"`
	// The ID of the VPC Endpoint to receive notifications for.
	VpcEndpointId *string `pulumi:"vpcEndpointId"`
	// The ID of the VPC Endpoint Service to receive notifications for.
	VpcEndpointServiceId *string `pulumi:"vpcEndpointServiceId"`
}

// The set of arguments for constructing a VpcEndpointConnectionNotification resource.
type VpcEndpointConnectionNotificationArgs struct {
	// One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
	//
	// > **NOTE:** One of `vpcEndpointServiceId` or `vpcEndpointId` must be specified.
	ConnectionEvents pulumi.StringArrayInput
	// The ARN of the SNS topic for the notifications.
	ConnectionNotificationArn pulumi.StringInput
	// The ID of the VPC Endpoint to receive notifications for.
	VpcEndpointId pulumi.StringPtrInput
	// The ID of the VPC Endpoint Service to receive notifications for.
	VpcEndpointServiceId pulumi.StringPtrInput
}

func (VpcEndpointConnectionNotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointConnectionNotificationArgs)(nil)).Elem()
}

type VpcEndpointConnectionNotificationInput interface {
	pulumi.Input

	ToVpcEndpointConnectionNotificationOutput() VpcEndpointConnectionNotificationOutput
	ToVpcEndpointConnectionNotificationOutputWithContext(ctx context.Context) VpcEndpointConnectionNotificationOutput
}

func (*VpcEndpointConnectionNotification) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointConnectionNotification)(nil)).Elem()
}

func (i *VpcEndpointConnectionNotification) ToVpcEndpointConnectionNotificationOutput() VpcEndpointConnectionNotificationOutput {
	return i.ToVpcEndpointConnectionNotificationOutputWithContext(context.Background())
}

func (i *VpcEndpointConnectionNotification) ToVpcEndpointConnectionNotificationOutputWithContext(ctx context.Context) VpcEndpointConnectionNotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointConnectionNotificationOutput)
}

// VpcEndpointConnectionNotificationArrayInput is an input type that accepts VpcEndpointConnectionNotificationArray and VpcEndpointConnectionNotificationArrayOutput values.
// You can construct a concrete instance of `VpcEndpointConnectionNotificationArrayInput` via:
//
//	VpcEndpointConnectionNotificationArray{ VpcEndpointConnectionNotificationArgs{...} }
type VpcEndpointConnectionNotificationArrayInput interface {
	pulumi.Input

	ToVpcEndpointConnectionNotificationArrayOutput() VpcEndpointConnectionNotificationArrayOutput
	ToVpcEndpointConnectionNotificationArrayOutputWithContext(context.Context) VpcEndpointConnectionNotificationArrayOutput
}

type VpcEndpointConnectionNotificationArray []VpcEndpointConnectionNotificationInput

func (VpcEndpointConnectionNotificationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointConnectionNotification)(nil)).Elem()
}

func (i VpcEndpointConnectionNotificationArray) ToVpcEndpointConnectionNotificationArrayOutput() VpcEndpointConnectionNotificationArrayOutput {
	return i.ToVpcEndpointConnectionNotificationArrayOutputWithContext(context.Background())
}

func (i VpcEndpointConnectionNotificationArray) ToVpcEndpointConnectionNotificationArrayOutputWithContext(ctx context.Context) VpcEndpointConnectionNotificationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointConnectionNotificationArrayOutput)
}

// VpcEndpointConnectionNotificationMapInput is an input type that accepts VpcEndpointConnectionNotificationMap and VpcEndpointConnectionNotificationMapOutput values.
// You can construct a concrete instance of `VpcEndpointConnectionNotificationMapInput` via:
//
//	VpcEndpointConnectionNotificationMap{ "key": VpcEndpointConnectionNotificationArgs{...} }
type VpcEndpointConnectionNotificationMapInput interface {
	pulumi.Input

	ToVpcEndpointConnectionNotificationMapOutput() VpcEndpointConnectionNotificationMapOutput
	ToVpcEndpointConnectionNotificationMapOutputWithContext(context.Context) VpcEndpointConnectionNotificationMapOutput
}

type VpcEndpointConnectionNotificationMap map[string]VpcEndpointConnectionNotificationInput

func (VpcEndpointConnectionNotificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointConnectionNotification)(nil)).Elem()
}

func (i VpcEndpointConnectionNotificationMap) ToVpcEndpointConnectionNotificationMapOutput() VpcEndpointConnectionNotificationMapOutput {
	return i.ToVpcEndpointConnectionNotificationMapOutputWithContext(context.Background())
}

func (i VpcEndpointConnectionNotificationMap) ToVpcEndpointConnectionNotificationMapOutputWithContext(ctx context.Context) VpcEndpointConnectionNotificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointConnectionNotificationMapOutput)
}

type VpcEndpointConnectionNotificationOutput struct{ *pulumi.OutputState }

func (VpcEndpointConnectionNotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointConnectionNotification)(nil)).Elem()
}

func (o VpcEndpointConnectionNotificationOutput) ToVpcEndpointConnectionNotificationOutput() VpcEndpointConnectionNotificationOutput {
	return o
}

func (o VpcEndpointConnectionNotificationOutput) ToVpcEndpointConnectionNotificationOutputWithContext(ctx context.Context) VpcEndpointConnectionNotificationOutput {
	return o
}

// One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
//
// > **NOTE:** One of `vpcEndpointServiceId` or `vpcEndpointId` must be specified.
func (o VpcEndpointConnectionNotificationOutput) ConnectionEvents() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpointConnectionNotification) pulumi.StringArrayOutput { return v.ConnectionEvents }).(pulumi.StringArrayOutput)
}

// The ARN of the SNS topic for the notifications.
func (o VpcEndpointConnectionNotificationOutput) ConnectionNotificationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointConnectionNotification) pulumi.StringOutput { return v.ConnectionNotificationArn }).(pulumi.StringOutput)
}

// The type of notification.
func (o VpcEndpointConnectionNotificationOutput) NotificationType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointConnectionNotification) pulumi.StringOutput { return v.NotificationType }).(pulumi.StringOutput)
}

// The state of the notification.
func (o VpcEndpointConnectionNotificationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointConnectionNotification) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The ID of the VPC Endpoint to receive notifications for.
func (o VpcEndpointConnectionNotificationOutput) VpcEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcEndpointConnectionNotification) pulumi.StringPtrOutput { return v.VpcEndpointId }).(pulumi.StringPtrOutput)
}

// The ID of the VPC Endpoint Service to receive notifications for.
func (o VpcEndpointConnectionNotificationOutput) VpcEndpointServiceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcEndpointConnectionNotification) pulumi.StringPtrOutput { return v.VpcEndpointServiceId }).(pulumi.StringPtrOutput)
}

type VpcEndpointConnectionNotificationArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointConnectionNotificationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointConnectionNotification)(nil)).Elem()
}

func (o VpcEndpointConnectionNotificationArrayOutput) ToVpcEndpointConnectionNotificationArrayOutput() VpcEndpointConnectionNotificationArrayOutput {
	return o
}

func (o VpcEndpointConnectionNotificationArrayOutput) ToVpcEndpointConnectionNotificationArrayOutputWithContext(ctx context.Context) VpcEndpointConnectionNotificationArrayOutput {
	return o
}

func (o VpcEndpointConnectionNotificationArrayOutput) Index(i pulumi.IntInput) VpcEndpointConnectionNotificationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcEndpointConnectionNotification {
		return vs[0].([]*VpcEndpointConnectionNotification)[vs[1].(int)]
	}).(VpcEndpointConnectionNotificationOutput)
}

type VpcEndpointConnectionNotificationMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointConnectionNotificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointConnectionNotification)(nil)).Elem()
}

func (o VpcEndpointConnectionNotificationMapOutput) ToVpcEndpointConnectionNotificationMapOutput() VpcEndpointConnectionNotificationMapOutput {
	return o
}

func (o VpcEndpointConnectionNotificationMapOutput) ToVpcEndpointConnectionNotificationMapOutputWithContext(ctx context.Context) VpcEndpointConnectionNotificationMapOutput {
	return o
}

func (o VpcEndpointConnectionNotificationMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointConnectionNotificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcEndpointConnectionNotification {
		return vs[0].(map[string]*VpcEndpointConnectionNotification)[vs[1].(string)]
	}).(VpcEndpointConnectionNotificationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointConnectionNotificationInput)(nil)).Elem(), &VpcEndpointConnectionNotification{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointConnectionNotificationArrayInput)(nil)).Elem(), VpcEndpointConnectionNotificationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointConnectionNotificationMapInput)(nil)).Elem(), VpcEndpointConnectionNotificationMap{})
	pulumi.RegisterOutputType(VpcEndpointConnectionNotificationOutput{})
	pulumi.RegisterOutputType(VpcEndpointConnectionNotificationArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointConnectionNotificationMapOutput{})
}
