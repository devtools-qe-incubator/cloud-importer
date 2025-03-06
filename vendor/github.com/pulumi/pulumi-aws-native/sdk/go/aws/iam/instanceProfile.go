// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new instance profile. For information about instance profiles, see [Using instance profiles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html).
//
//	For information about the number of instance profiles you can create, see [object quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *User Guide*.
//
// ## Example Usage
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/ec2"
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myIamInstanceProfile, err := iam.NewInstanceProfile(ctx, "myIamInstanceProfile", &iam.InstanceProfileArgs{
//				InstanceProfileName: pulumi.String("MyIamInstanceProfile"),
//				Path:                pulumi.String("/"),
//				Roles: pulumi.StringArray{
//					pulumi.String("MyAdminRole"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewLaunchTemplate(ctx, "myLaunchTemplate", &ec2.LaunchTemplateArgs{
//				LaunchTemplateName: pulumi.String("MyLaunchTemplate"),
//				LaunchTemplateData: &ec2.LaunchTemplateDataArgs{
//					IamInstanceProfile: &ec2.LaunchTemplateIamInstanceProfileArgs{
//						Arn: myIamInstanceProfile.Arn,
//					},
//					DisableApiTermination: pulumi.Bool(true),
//					ImageId:               pulumi.String("ami-04d5cc9b88example"),
//					InstanceType:          pulumi.String("t2.micro"),
//					KeyName:               pulumi.String("MyKeyPair"),
//					SecurityGroupIds: pulumi.StringArray{
//						pulumi.String("sg-083cd3bfb8example"),
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
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/ec2"
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myIamInstanceProfile, err := iam.NewInstanceProfile(ctx, "myIamInstanceProfile", &iam.InstanceProfileArgs{
//				InstanceProfileName: pulumi.String("MyIamInstanceProfile"),
//				Path:                pulumi.String("/"),
//				Roles: pulumi.StringArray{
//					pulumi.String("MyAdminRole"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewLaunchTemplate(ctx, "myLaunchTemplate", &ec2.LaunchTemplateArgs{
//				LaunchTemplateName: pulumi.String("MyLaunchTemplate"),
//				LaunchTemplateData: &ec2.LaunchTemplateDataArgs{
//					IamInstanceProfile: &ec2.LaunchTemplateIamInstanceProfileArgs{
//						Arn: myIamInstanceProfile.Arn,
//					},
//					DisableApiTermination: pulumi.Bool(true),
//					ImageId:               pulumi.String("ami-04d5cc9b88example"),
//					InstanceType:          pulumi.String("t2.micro"),
//					KeyName:               pulumi.String("MyKeyPair"),
//					SecurityGroupIds: pulumi.StringArray{
//						pulumi.String("sg-083cd3bfb8example"),
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
type InstanceProfile struct {
	pulumi.CustomResourceState

	// Returns the Amazon Resource Name (ARN) for the instance profile. For example:
	//
	// `{"Fn::GetAtt" : ["MyProfile", "Arn"] }`
	//
	// This returns a value such as `arn:aws:iam::1234567890:instance-profile/MyProfile-ASDNSDLKJ` .
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the instance profile to create.
	//  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	InstanceProfileName pulumi.StringPtrOutput `pulumi:"instanceProfileName"`
	// The path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*.
	//  This parameter is optional. If it is not included, it defaults to a slash (/).
	//  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (``\u0021``) through the DEL character (``\u007F``), including most punctuation characters, digits, and upper and lowercased letters.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The name of the role to associate with the instance profile. Only one role can be assigned to an EC2 instance at a time, and all applications on the instance share the same role and permissions.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
}

// NewInstanceProfile registers a new resource with the given unique name, arguments, and options.
func NewInstanceProfile(ctx *pulumi.Context,
	name string, args *InstanceProfileArgs, opts ...pulumi.ResourceOption) (*InstanceProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"instanceProfileName",
		"path",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceProfile
	err := ctx.RegisterResource("aws-native:iam:InstanceProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceProfile gets an existing InstanceProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceProfileState, opts ...pulumi.ResourceOption) (*InstanceProfile, error) {
	var resource InstanceProfile
	err := ctx.ReadResource("aws-native:iam:InstanceProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceProfile resources.
type instanceProfileState struct {
}

type InstanceProfileState struct {
}

func (InstanceProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceProfileState)(nil)).Elem()
}

type instanceProfileArgs struct {
	// The name of the instance profile to create.
	//  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	InstanceProfileName *string `pulumi:"instanceProfileName"`
	// The path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*.
	//  This parameter is optional. If it is not included, it defaults to a slash (/).
	//  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (``\u0021``) through the DEL character (``\u007F``), including most punctuation characters, digits, and upper and lowercased letters.
	Path *string `pulumi:"path"`
	// The name of the role to associate with the instance profile. Only one role can be assigned to an EC2 instance at a time, and all applications on the instance share the same role and permissions.
	Roles []string `pulumi:"roles"`
}

// The set of arguments for constructing a InstanceProfile resource.
type InstanceProfileArgs struct {
	// The name of the instance profile to create.
	//  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	InstanceProfileName pulumi.StringPtrInput
	// The path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*.
	//  This parameter is optional. If it is not included, it defaults to a slash (/).
	//  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (``\u0021``) through the DEL character (``\u007F``), including most punctuation characters, digits, and upper and lowercased letters.
	Path pulumi.StringPtrInput
	// The name of the role to associate with the instance profile. Only one role can be assigned to an EC2 instance at a time, and all applications on the instance share the same role and permissions.
	Roles pulumi.StringArrayInput
}

func (InstanceProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceProfileArgs)(nil)).Elem()
}

type InstanceProfileInput interface {
	pulumi.Input

	ToInstanceProfileOutput() InstanceProfileOutput
	ToInstanceProfileOutputWithContext(ctx context.Context) InstanceProfileOutput
}

func (*InstanceProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceProfile)(nil)).Elem()
}

func (i *InstanceProfile) ToInstanceProfileOutput() InstanceProfileOutput {
	return i.ToInstanceProfileOutputWithContext(context.Background())
}

func (i *InstanceProfile) ToInstanceProfileOutputWithContext(ctx context.Context) InstanceProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceProfileOutput)
}

type InstanceProfileOutput struct{ *pulumi.OutputState }

func (InstanceProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceProfile)(nil)).Elem()
}

func (o InstanceProfileOutput) ToInstanceProfileOutput() InstanceProfileOutput {
	return o
}

func (o InstanceProfileOutput) ToInstanceProfileOutputWithContext(ctx context.Context) InstanceProfileOutput {
	return o
}

// Returns the Amazon Resource Name (ARN) for the instance profile. For example:
//
// `{"Fn::GetAtt" : ["MyProfile", "Arn"] }`
//
// This returns a value such as `arn:aws:iam::1234567890:instance-profile/MyProfile-ASDNSDLKJ` .
func (o InstanceProfileOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceProfile) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The name of the instance profile to create.
//
//	This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
func (o InstanceProfileOutput) InstanceProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceProfile) pulumi.StringPtrOutput { return v.InstanceProfileName }).(pulumi.StringPtrOutput)
}

// The path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*.
//
//	This parameter is optional. If it is not included, it defaults to a slash (/).
//	This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (``\u0021``) through the DEL character (``\u007F``), including most punctuation characters, digits, and upper and lowercased letters.
func (o InstanceProfileOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceProfile) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// The name of the role to associate with the instance profile. Only one role can be assigned to an EC2 instance at a time, and all applications on the instance share the same role and permissions.
func (o InstanceProfileOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstanceProfile) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceProfileInput)(nil)).Elem(), &InstanceProfile{})
	pulumi.RegisterOutputType(InstanceProfileOutput{})
}
