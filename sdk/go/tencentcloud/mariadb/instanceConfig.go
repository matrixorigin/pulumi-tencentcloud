// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceConfig struct {
	pulumi.CustomResourceState

	// External network status, 0-closed; 1- Opening; Default not enabled.
	ExtranetAccess pulumi.IntOutput `pulumi:"extranetAccess"`
	// instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// RS proximity mode, 0- no strategy, 1- access to the nearest available zone.
	RsAccessStrategy pulumi.IntOutput `pulumi:"rsAccessStrategy"`
}

// NewInstanceConfig registers a new resource with the given unique name, arguments, and options.
func NewInstanceConfig(ctx *pulumi.Context,
	name string, args *InstanceConfigArgs, opts ...pulumi.ResourceOption) (*InstanceConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceConfig
	err := ctx.RegisterResource("tencentcloud:Mariadb/instanceConfig:InstanceConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceConfig gets an existing InstanceConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceConfigState, opts ...pulumi.ResourceOption) (*InstanceConfig, error) {
	var resource InstanceConfig
	err := ctx.ReadResource("tencentcloud:Mariadb/instanceConfig:InstanceConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceConfig resources.
type instanceConfigState struct {
	// External network status, 0-closed; 1- Opening; Default not enabled.
	ExtranetAccess *int `pulumi:"extranetAccess"`
	// instance id.
	InstanceId *string `pulumi:"instanceId"`
	// RS proximity mode, 0- no strategy, 1- access to the nearest available zone.
	RsAccessStrategy *int `pulumi:"rsAccessStrategy"`
}

type InstanceConfigState struct {
	// External network status, 0-closed; 1- Opening; Default not enabled.
	ExtranetAccess pulumi.IntPtrInput
	// instance id.
	InstanceId pulumi.StringPtrInput
	// RS proximity mode, 0- no strategy, 1- access to the nearest available zone.
	RsAccessStrategy pulumi.IntPtrInput
}

func (InstanceConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceConfigState)(nil)).Elem()
}

type instanceConfigArgs struct {
	// External network status, 0-closed; 1- Opening; Default not enabled.
	ExtranetAccess *int `pulumi:"extranetAccess"`
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// RS proximity mode, 0- no strategy, 1- access to the nearest available zone.
	RsAccessStrategy *int `pulumi:"rsAccessStrategy"`
}

// The set of arguments for constructing a InstanceConfig resource.
type InstanceConfigArgs struct {
	// External network status, 0-closed; 1- Opening; Default not enabled.
	ExtranetAccess pulumi.IntPtrInput
	// instance id.
	InstanceId pulumi.StringInput
	// RS proximity mode, 0- no strategy, 1- access to the nearest available zone.
	RsAccessStrategy pulumi.IntPtrInput
}

func (InstanceConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceConfigArgs)(nil)).Elem()
}

type InstanceConfigInput interface {
	pulumi.Input

	ToInstanceConfigOutput() InstanceConfigOutput
	ToInstanceConfigOutputWithContext(ctx context.Context) InstanceConfigOutput
}

func (*InstanceConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceConfig)(nil)).Elem()
}

func (i *InstanceConfig) ToInstanceConfigOutput() InstanceConfigOutput {
	return i.ToInstanceConfigOutputWithContext(context.Background())
}

func (i *InstanceConfig) ToInstanceConfigOutputWithContext(ctx context.Context) InstanceConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceConfigOutput)
}

// InstanceConfigArrayInput is an input type that accepts InstanceConfigArray and InstanceConfigArrayOutput values.
// You can construct a concrete instance of `InstanceConfigArrayInput` via:
//
//	InstanceConfigArray{ InstanceConfigArgs{...} }
type InstanceConfigArrayInput interface {
	pulumi.Input

	ToInstanceConfigArrayOutput() InstanceConfigArrayOutput
	ToInstanceConfigArrayOutputWithContext(context.Context) InstanceConfigArrayOutput
}

type InstanceConfigArray []InstanceConfigInput

func (InstanceConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceConfig)(nil)).Elem()
}

func (i InstanceConfigArray) ToInstanceConfigArrayOutput() InstanceConfigArrayOutput {
	return i.ToInstanceConfigArrayOutputWithContext(context.Background())
}

func (i InstanceConfigArray) ToInstanceConfigArrayOutputWithContext(ctx context.Context) InstanceConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceConfigArrayOutput)
}

// InstanceConfigMapInput is an input type that accepts InstanceConfigMap and InstanceConfigMapOutput values.
// You can construct a concrete instance of `InstanceConfigMapInput` via:
//
//	InstanceConfigMap{ "key": InstanceConfigArgs{...} }
type InstanceConfigMapInput interface {
	pulumi.Input

	ToInstanceConfigMapOutput() InstanceConfigMapOutput
	ToInstanceConfigMapOutputWithContext(context.Context) InstanceConfigMapOutput
}

type InstanceConfigMap map[string]InstanceConfigInput

func (InstanceConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceConfig)(nil)).Elem()
}

func (i InstanceConfigMap) ToInstanceConfigMapOutput() InstanceConfigMapOutput {
	return i.ToInstanceConfigMapOutputWithContext(context.Background())
}

func (i InstanceConfigMap) ToInstanceConfigMapOutputWithContext(ctx context.Context) InstanceConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceConfigMapOutput)
}

type InstanceConfigOutput struct{ *pulumi.OutputState }

func (InstanceConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceConfig)(nil)).Elem()
}

func (o InstanceConfigOutput) ToInstanceConfigOutput() InstanceConfigOutput {
	return o
}

func (o InstanceConfigOutput) ToInstanceConfigOutputWithContext(ctx context.Context) InstanceConfigOutput {
	return o
}

// External network status, 0-closed; 1- Opening; Default not enabled.
func (o InstanceConfigOutput) ExtranetAccess() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceConfig) pulumi.IntOutput { return v.ExtranetAccess }).(pulumi.IntOutput)
}

// instance id.
func (o InstanceConfigOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceConfig) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// RS proximity mode, 0- no strategy, 1- access to the nearest available zone.
func (o InstanceConfigOutput) RsAccessStrategy() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceConfig) pulumi.IntOutput { return v.RsAccessStrategy }).(pulumi.IntOutput)
}

type InstanceConfigArrayOutput struct{ *pulumi.OutputState }

func (InstanceConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceConfig)(nil)).Elem()
}

func (o InstanceConfigArrayOutput) ToInstanceConfigArrayOutput() InstanceConfigArrayOutput {
	return o
}

func (o InstanceConfigArrayOutput) ToInstanceConfigArrayOutputWithContext(ctx context.Context) InstanceConfigArrayOutput {
	return o
}

func (o InstanceConfigArrayOutput) Index(i pulumi.IntInput) InstanceConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceConfig {
		return vs[0].([]*InstanceConfig)[vs[1].(int)]
	}).(InstanceConfigOutput)
}

type InstanceConfigMapOutput struct{ *pulumi.OutputState }

func (InstanceConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceConfig)(nil)).Elem()
}

func (o InstanceConfigMapOutput) ToInstanceConfigMapOutput() InstanceConfigMapOutput {
	return o
}

func (o InstanceConfigMapOutput) ToInstanceConfigMapOutputWithContext(ctx context.Context) InstanceConfigMapOutput {
	return o
}

func (o InstanceConfigMapOutput) MapIndex(k pulumi.StringInput) InstanceConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceConfig {
		return vs[0].(map[string]*InstanceConfig)[vs[1].(string)]
	}).(InstanceConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceConfigInput)(nil)).Elem(), &InstanceConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceConfigArrayInput)(nil)).Elem(), InstanceConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceConfigMapInput)(nil)).Elem(), InstanceConfigMap{})
	pulumi.RegisterOutputType(InstanceConfigOutput{})
	pulumi.RegisterOutputType(InstanceConfigArrayOutput{})
	pulumi.RegisterOutputType(InstanceConfigMapOutput{})
}