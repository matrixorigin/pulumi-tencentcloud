// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TrafficPackage struct {
	pulumi.CustomResourceState

	// Created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Remaining amount.
	RemainingAmount pulumi.Float64Output `pulumi:"remainingAmount"`
	// Traffic Package Amount, eg: 10,20,50,512,1024,5120,51200,60,300,600,3072,6144,30720,61440,307200.
	TrafficAmount pulumi.IntOutput `pulumi:"trafficAmount"`
	// Used amount.
	UsedAmount pulumi.Float64Output `pulumi:"usedAmount"`
}

// NewTrafficPackage registers a new resource with the given unique name, arguments, and options.
func NewTrafficPackage(ctx *pulumi.Context,
	name string, args *TrafficPackageArgs, opts ...pulumi.ResourceOption) (*TrafficPackage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TrafficAmount == nil {
		return nil, errors.New("invalid value for required argument 'TrafficAmount'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TrafficPackage
	err := ctx.RegisterResource("tencentcloud:Vpc/trafficPackage:TrafficPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficPackage gets an existing TrafficPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficPackageState, opts ...pulumi.ResourceOption) (*TrafficPackage, error) {
	var resource TrafficPackage
	err := ctx.ReadResource("tencentcloud:Vpc/trafficPackage:TrafficPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficPackage resources.
type trafficPackageState struct {
	// Created time.
	CreatedTime *string `pulumi:"createdTime"`
	// Remaining amount.
	RemainingAmount *float64 `pulumi:"remainingAmount"`
	// Traffic Package Amount, eg: 10,20,50,512,1024,5120,51200,60,300,600,3072,6144,30720,61440,307200.
	TrafficAmount *int `pulumi:"trafficAmount"`
	// Used amount.
	UsedAmount *float64 `pulumi:"usedAmount"`
}

type TrafficPackageState struct {
	// Created time.
	CreatedTime pulumi.StringPtrInput
	// Remaining amount.
	RemainingAmount pulumi.Float64PtrInput
	// Traffic Package Amount, eg: 10,20,50,512,1024,5120,51200,60,300,600,3072,6144,30720,61440,307200.
	TrafficAmount pulumi.IntPtrInput
	// Used amount.
	UsedAmount pulumi.Float64PtrInput
}

func (TrafficPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficPackageState)(nil)).Elem()
}

type trafficPackageArgs struct {
	// Traffic Package Amount, eg: 10,20,50,512,1024,5120,51200,60,300,600,3072,6144,30720,61440,307200.
	TrafficAmount int `pulumi:"trafficAmount"`
}

// The set of arguments for constructing a TrafficPackage resource.
type TrafficPackageArgs struct {
	// Traffic Package Amount, eg: 10,20,50,512,1024,5120,51200,60,300,600,3072,6144,30720,61440,307200.
	TrafficAmount pulumi.IntInput
}

func (TrafficPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficPackageArgs)(nil)).Elem()
}

type TrafficPackageInput interface {
	pulumi.Input

	ToTrafficPackageOutput() TrafficPackageOutput
	ToTrafficPackageOutputWithContext(ctx context.Context) TrafficPackageOutput
}

func (*TrafficPackage) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficPackage)(nil)).Elem()
}

func (i *TrafficPackage) ToTrafficPackageOutput() TrafficPackageOutput {
	return i.ToTrafficPackageOutputWithContext(context.Background())
}

func (i *TrafficPackage) ToTrafficPackageOutputWithContext(ctx context.Context) TrafficPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficPackageOutput)
}

// TrafficPackageArrayInput is an input type that accepts TrafficPackageArray and TrafficPackageArrayOutput values.
// You can construct a concrete instance of `TrafficPackageArrayInput` via:
//
//	TrafficPackageArray{ TrafficPackageArgs{...} }
type TrafficPackageArrayInput interface {
	pulumi.Input

	ToTrafficPackageArrayOutput() TrafficPackageArrayOutput
	ToTrafficPackageArrayOutputWithContext(context.Context) TrafficPackageArrayOutput
}

type TrafficPackageArray []TrafficPackageInput

func (TrafficPackageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficPackage)(nil)).Elem()
}

func (i TrafficPackageArray) ToTrafficPackageArrayOutput() TrafficPackageArrayOutput {
	return i.ToTrafficPackageArrayOutputWithContext(context.Background())
}

func (i TrafficPackageArray) ToTrafficPackageArrayOutputWithContext(ctx context.Context) TrafficPackageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficPackageArrayOutput)
}

// TrafficPackageMapInput is an input type that accepts TrafficPackageMap and TrafficPackageMapOutput values.
// You can construct a concrete instance of `TrafficPackageMapInput` via:
//
//	TrafficPackageMap{ "key": TrafficPackageArgs{...} }
type TrafficPackageMapInput interface {
	pulumi.Input

	ToTrafficPackageMapOutput() TrafficPackageMapOutput
	ToTrafficPackageMapOutputWithContext(context.Context) TrafficPackageMapOutput
}

type TrafficPackageMap map[string]TrafficPackageInput

func (TrafficPackageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficPackage)(nil)).Elem()
}

func (i TrafficPackageMap) ToTrafficPackageMapOutput() TrafficPackageMapOutput {
	return i.ToTrafficPackageMapOutputWithContext(context.Background())
}

func (i TrafficPackageMap) ToTrafficPackageMapOutputWithContext(ctx context.Context) TrafficPackageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficPackageMapOutput)
}

type TrafficPackageOutput struct{ *pulumi.OutputState }

func (TrafficPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficPackage)(nil)).Elem()
}

func (o TrafficPackageOutput) ToTrafficPackageOutput() TrafficPackageOutput {
	return o
}

func (o TrafficPackageOutput) ToTrafficPackageOutputWithContext(ctx context.Context) TrafficPackageOutput {
	return o
}

// Created time.
func (o TrafficPackageOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficPackage) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// Remaining amount.
func (o TrafficPackageOutput) RemainingAmount() pulumi.Float64Output {
	return o.ApplyT(func(v *TrafficPackage) pulumi.Float64Output { return v.RemainingAmount }).(pulumi.Float64Output)
}

// Traffic Package Amount, eg: 10,20,50,512,1024,5120,51200,60,300,600,3072,6144,30720,61440,307200.
func (o TrafficPackageOutput) TrafficAmount() pulumi.IntOutput {
	return o.ApplyT(func(v *TrafficPackage) pulumi.IntOutput { return v.TrafficAmount }).(pulumi.IntOutput)
}

// Used amount.
func (o TrafficPackageOutput) UsedAmount() pulumi.Float64Output {
	return o.ApplyT(func(v *TrafficPackage) pulumi.Float64Output { return v.UsedAmount }).(pulumi.Float64Output)
}

type TrafficPackageArrayOutput struct{ *pulumi.OutputState }

func (TrafficPackageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficPackage)(nil)).Elem()
}

func (o TrafficPackageArrayOutput) ToTrafficPackageArrayOutput() TrafficPackageArrayOutput {
	return o
}

func (o TrafficPackageArrayOutput) ToTrafficPackageArrayOutputWithContext(ctx context.Context) TrafficPackageArrayOutput {
	return o
}

func (o TrafficPackageArrayOutput) Index(i pulumi.IntInput) TrafficPackageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrafficPackage {
		return vs[0].([]*TrafficPackage)[vs[1].(int)]
	}).(TrafficPackageOutput)
}

type TrafficPackageMapOutput struct{ *pulumi.OutputState }

func (TrafficPackageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficPackage)(nil)).Elem()
}

func (o TrafficPackageMapOutput) ToTrafficPackageMapOutput() TrafficPackageMapOutput {
	return o
}

func (o TrafficPackageMapOutput) ToTrafficPackageMapOutputWithContext(ctx context.Context) TrafficPackageMapOutput {
	return o
}

func (o TrafficPackageMapOutput) MapIndex(k pulumi.StringInput) TrafficPackageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrafficPackage {
		return vs[0].(map[string]*TrafficPackage)[vs[1].(string)]
	}).(TrafficPackageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficPackageInput)(nil)).Elem(), &TrafficPackage{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficPackageArrayInput)(nil)).Elem(), TrafficPackageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficPackageMapInput)(nil)).Elem(), TrafficPackageMap{})
	pulumi.RegisterOutputType(TrafficPackageOutput{})
	pulumi.RegisterOutputType(TrafficPackageArrayOutput{})
	pulumi.RegisterOutputType(TrafficPackageMapOutput{})
}