// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cls

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigAttachment struct {
	pulumi.CustomResourceState

	// Collection configuration id.
	ConfigId pulumi.StringOutput `pulumi:"configId"`
	// Machine group id.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
}

// NewConfigAttachment registers a new resource with the given unique name, arguments, and options.
func NewConfigAttachment(ctx *pulumi.Context,
	name string, args *ConfigAttachmentArgs, opts ...pulumi.ResourceOption) (*ConfigAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConfigAttachment
	err := ctx.RegisterResource("tencentcloud:Cls/configAttachment:ConfigAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigAttachment gets an existing ConfigAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigAttachmentState, opts ...pulumi.ResourceOption) (*ConfigAttachment, error) {
	var resource ConfigAttachment
	err := ctx.ReadResource("tencentcloud:Cls/configAttachment:ConfigAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigAttachment resources.
type configAttachmentState struct {
	// Collection configuration id.
	ConfigId *string `pulumi:"configId"`
	// Machine group id.
	GroupId *string `pulumi:"groupId"`
}

type ConfigAttachmentState struct {
	// Collection configuration id.
	ConfigId pulumi.StringPtrInput
	// Machine group id.
	GroupId pulumi.StringPtrInput
}

func (ConfigAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*configAttachmentState)(nil)).Elem()
}

type configAttachmentArgs struct {
	// Collection configuration id.
	ConfigId string `pulumi:"configId"`
	// Machine group id.
	GroupId string `pulumi:"groupId"`
}

// The set of arguments for constructing a ConfigAttachment resource.
type ConfigAttachmentArgs struct {
	// Collection configuration id.
	ConfigId pulumi.StringInput
	// Machine group id.
	GroupId pulumi.StringInput
}

func (ConfigAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configAttachmentArgs)(nil)).Elem()
}

type ConfigAttachmentInput interface {
	pulumi.Input

	ToConfigAttachmentOutput() ConfigAttachmentOutput
	ToConfigAttachmentOutputWithContext(ctx context.Context) ConfigAttachmentOutput
}

func (*ConfigAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigAttachment)(nil)).Elem()
}

func (i *ConfigAttachment) ToConfigAttachmentOutput() ConfigAttachmentOutput {
	return i.ToConfigAttachmentOutputWithContext(context.Background())
}

func (i *ConfigAttachment) ToConfigAttachmentOutputWithContext(ctx context.Context) ConfigAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigAttachmentOutput)
}

// ConfigAttachmentArrayInput is an input type that accepts ConfigAttachmentArray and ConfigAttachmentArrayOutput values.
// You can construct a concrete instance of `ConfigAttachmentArrayInput` via:
//
//	ConfigAttachmentArray{ ConfigAttachmentArgs{...} }
type ConfigAttachmentArrayInput interface {
	pulumi.Input

	ToConfigAttachmentArrayOutput() ConfigAttachmentArrayOutput
	ToConfigAttachmentArrayOutputWithContext(context.Context) ConfigAttachmentArrayOutput
}

type ConfigAttachmentArray []ConfigAttachmentInput

func (ConfigAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigAttachment)(nil)).Elem()
}

func (i ConfigAttachmentArray) ToConfigAttachmentArrayOutput() ConfigAttachmentArrayOutput {
	return i.ToConfigAttachmentArrayOutputWithContext(context.Background())
}

func (i ConfigAttachmentArray) ToConfigAttachmentArrayOutputWithContext(ctx context.Context) ConfigAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigAttachmentArrayOutput)
}

// ConfigAttachmentMapInput is an input type that accepts ConfigAttachmentMap and ConfigAttachmentMapOutput values.
// You can construct a concrete instance of `ConfigAttachmentMapInput` via:
//
//	ConfigAttachmentMap{ "key": ConfigAttachmentArgs{...} }
type ConfigAttachmentMapInput interface {
	pulumi.Input

	ToConfigAttachmentMapOutput() ConfigAttachmentMapOutput
	ToConfigAttachmentMapOutputWithContext(context.Context) ConfigAttachmentMapOutput
}

type ConfigAttachmentMap map[string]ConfigAttachmentInput

func (ConfigAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigAttachment)(nil)).Elem()
}

func (i ConfigAttachmentMap) ToConfigAttachmentMapOutput() ConfigAttachmentMapOutput {
	return i.ToConfigAttachmentMapOutputWithContext(context.Background())
}

func (i ConfigAttachmentMap) ToConfigAttachmentMapOutputWithContext(ctx context.Context) ConfigAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigAttachmentMapOutput)
}

type ConfigAttachmentOutput struct{ *pulumi.OutputState }

func (ConfigAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigAttachment)(nil)).Elem()
}

func (o ConfigAttachmentOutput) ToConfigAttachmentOutput() ConfigAttachmentOutput {
	return o
}

func (o ConfigAttachmentOutput) ToConfigAttachmentOutputWithContext(ctx context.Context) ConfigAttachmentOutput {
	return o
}

// Collection configuration id.
func (o ConfigAttachmentOutput) ConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigAttachment) pulumi.StringOutput { return v.ConfigId }).(pulumi.StringOutput)
}

// Machine group id.
func (o ConfigAttachmentOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigAttachment) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

type ConfigAttachmentArrayOutput struct{ *pulumi.OutputState }

func (ConfigAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigAttachment)(nil)).Elem()
}

func (o ConfigAttachmentArrayOutput) ToConfigAttachmentArrayOutput() ConfigAttachmentArrayOutput {
	return o
}

func (o ConfigAttachmentArrayOutput) ToConfigAttachmentArrayOutputWithContext(ctx context.Context) ConfigAttachmentArrayOutput {
	return o
}

func (o ConfigAttachmentArrayOutput) Index(i pulumi.IntInput) ConfigAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigAttachment {
		return vs[0].([]*ConfigAttachment)[vs[1].(int)]
	}).(ConfigAttachmentOutput)
}

type ConfigAttachmentMapOutput struct{ *pulumi.OutputState }

func (ConfigAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigAttachment)(nil)).Elem()
}

func (o ConfigAttachmentMapOutput) ToConfigAttachmentMapOutput() ConfigAttachmentMapOutput {
	return o
}

func (o ConfigAttachmentMapOutput) ToConfigAttachmentMapOutputWithContext(ctx context.Context) ConfigAttachmentMapOutput {
	return o
}

func (o ConfigAttachmentMapOutput) MapIndex(k pulumi.StringInput) ConfigAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigAttachment {
		return vs[0].(map[string]*ConfigAttachment)[vs[1].(string)]
	}).(ConfigAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigAttachmentInput)(nil)).Elem(), &ConfigAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigAttachmentArrayInput)(nil)).Elem(), ConfigAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigAttachmentMapInput)(nil)).Elem(), ConfigAttachmentMap{})
	pulumi.RegisterOutputType(ConfigAttachmentOutput{})
	pulumi.RegisterOutputType(ConfigAttachmentArrayOutput{})
	pulumi.RegisterOutputType(ConfigAttachmentMapOutput{})
}