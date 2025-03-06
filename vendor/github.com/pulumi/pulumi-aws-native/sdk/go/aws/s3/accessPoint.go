// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::S3::AccessPoint resource is an Amazon S3 resource type that you can use to access buckets.
type AccessPoint struct {
	pulumi.CustomResourceState

	// The alias of this Access Point. This alias can be used for compatibility purposes with other AWS services and third-party applications.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// The Amazon Resource Name (ARN) of the specified accesspoint.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the bucket that you want to associate this Access Point with.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The AWS account ID associated with the S3 bucket associated with this access point.
	BucketAccountId pulumi.StringPtrOutput `pulumi:"bucketAccountId"`
	// The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.
	NetworkOrigin AccessPointNetworkOriginOutput `pulumi:"networkOrigin"`
	// The Access Point Policy you want to apply to this access point.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3::AccessPoint` for more information about the expected schema for this property.
	Policy pulumi.AnyOutput `pulumi:"policy"`
	// The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
	PublicAccessBlockConfiguration AccessPointPublicAccessBlockConfigurationPtrOutput `pulumi:"publicAccessBlockConfiguration"`
	// If you include this field, Amazon S3 restricts access to this Access Point to requests from the specified Virtual Private Cloud (VPC).
	VpcConfiguration AccessPointVpcConfigurationPtrOutput `pulumi:"vpcConfiguration"`
}

// NewAccessPoint registers a new resource with the given unique name, arguments, and options.
func NewAccessPoint(ctx *pulumi.Context,
	name string, args *AccessPointArgs, opts ...pulumi.ResourceOption) (*AccessPoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"bucket",
		"bucketAccountId",
		"name",
		"vpcConfiguration",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccessPoint
	err := ctx.RegisterResource("aws-native:s3:AccessPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPoint gets an existing AccessPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPointState, opts ...pulumi.ResourceOption) (*AccessPoint, error) {
	var resource AccessPoint
	err := ctx.ReadResource("aws-native:s3:AccessPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPoint resources.
type accessPointState struct {
}

type AccessPointState struct {
}

func (AccessPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPointState)(nil)).Elem()
}

type accessPointArgs struct {
	// The name of the bucket that you want to associate this Access Point with.
	Bucket string `pulumi:"bucket"`
	// The AWS account ID associated with the S3 bucket associated with this access point.
	BucketAccountId *string `pulumi:"bucketAccountId"`
	// The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.
	Name *string `pulumi:"name"`
	// The Access Point Policy you want to apply to this access point.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3::AccessPoint` for more information about the expected schema for this property.
	Policy interface{} `pulumi:"policy"`
	// The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
	PublicAccessBlockConfiguration *AccessPointPublicAccessBlockConfiguration `pulumi:"publicAccessBlockConfiguration"`
	// If you include this field, Amazon S3 restricts access to this Access Point to requests from the specified Virtual Private Cloud (VPC).
	VpcConfiguration *AccessPointVpcConfiguration `pulumi:"vpcConfiguration"`
}

// The set of arguments for constructing a AccessPoint resource.
type AccessPointArgs struct {
	// The name of the bucket that you want to associate this Access Point with.
	Bucket pulumi.StringInput
	// The AWS account ID associated with the S3 bucket associated with this access point.
	BucketAccountId pulumi.StringPtrInput
	// The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.
	Name pulumi.StringPtrInput
	// The Access Point Policy you want to apply to this access point.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3::AccessPoint` for more information about the expected schema for this property.
	Policy pulumi.Input
	// The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
	PublicAccessBlockConfiguration AccessPointPublicAccessBlockConfigurationPtrInput
	// If you include this field, Amazon S3 restricts access to this Access Point to requests from the specified Virtual Private Cloud (VPC).
	VpcConfiguration AccessPointVpcConfigurationPtrInput
}

func (AccessPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPointArgs)(nil)).Elem()
}

type AccessPointInput interface {
	pulumi.Input

	ToAccessPointOutput() AccessPointOutput
	ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput
}

func (*AccessPoint) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPoint)(nil)).Elem()
}

func (i *AccessPoint) ToAccessPointOutput() AccessPointOutput {
	return i.ToAccessPointOutputWithContext(context.Background())
}

func (i *AccessPoint) ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointOutput)
}

type AccessPointOutput struct{ *pulumi.OutputState }

func (AccessPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPoint)(nil)).Elem()
}

func (o AccessPointOutput) ToAccessPointOutput() AccessPointOutput {
	return o
}

func (o AccessPointOutput) ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput {
	return o
}

// The alias of this Access Point. This alias can be used for compatibility purposes with other AWS services and third-party applications.
func (o AccessPointOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the specified accesspoint.
func (o AccessPointOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The name of the bucket that you want to associate this Access Point with.
func (o AccessPointOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The AWS account ID associated with the S3 bucket associated with this access point.
func (o AccessPointOutput) BucketAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringPtrOutput { return v.BucketAccountId }).(pulumi.StringPtrOutput)
}

// The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.
func (o AccessPointOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.
func (o AccessPointOutput) NetworkOrigin() AccessPointNetworkOriginOutput {
	return o.ApplyT(func(v *AccessPoint) AccessPointNetworkOriginOutput { return v.NetworkOrigin }).(AccessPointNetworkOriginOutput)
}

// The Access Point Policy you want to apply to this access point.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3::AccessPoint` for more information about the expected schema for this property.
func (o AccessPointOutput) Policy() pulumi.AnyOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.AnyOutput { return v.Policy }).(pulumi.AnyOutput)
}

// The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
func (o AccessPointOutput) PublicAccessBlockConfiguration() AccessPointPublicAccessBlockConfigurationPtrOutput {
	return o.ApplyT(func(v *AccessPoint) AccessPointPublicAccessBlockConfigurationPtrOutput {
		return v.PublicAccessBlockConfiguration
	}).(AccessPointPublicAccessBlockConfigurationPtrOutput)
}

// If you include this field, Amazon S3 restricts access to this Access Point to requests from the specified Virtual Private Cloud (VPC).
func (o AccessPointOutput) VpcConfiguration() AccessPointVpcConfigurationPtrOutput {
	return o.ApplyT(func(v *AccessPoint) AccessPointVpcConfigurationPtrOutput { return v.VpcConfiguration }).(AccessPointVpcConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPointInput)(nil)).Elem(), &AccessPoint{})
	pulumi.RegisterOutputType(AccessPointOutput{})
}
