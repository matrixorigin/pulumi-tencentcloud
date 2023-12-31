// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deprecatednats

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type GetInstanceNat struct {
	AssignedEipSets []string `pulumi:"assignedEipSets"`
	Bandwidth       int      `pulumi:"bandwidth"`
	CreateTime      string   `pulumi:"createTime"`
	Id              string   `pulumi:"id"`
	MaxConcurrent   int      `pulumi:"maxConcurrent"`
	Name            string   `pulumi:"name"`
	State           int      `pulumi:"state"`
	VpcId           string   `pulumi:"vpcId"`
}

// GetInstanceNatInput is an input type that accepts GetInstanceNatArgs and GetInstanceNatOutput values.
// You can construct a concrete instance of `GetInstanceNatInput` via:
//
//	GetInstanceNatArgs{...}
type GetInstanceNatInput interface {
	pulumi.Input

	ToGetInstanceNatOutput() GetInstanceNatOutput
	ToGetInstanceNatOutputWithContext(context.Context) GetInstanceNatOutput
}

type GetInstanceNatArgs struct {
	AssignedEipSets pulumi.StringArrayInput `pulumi:"assignedEipSets"`
	Bandwidth       pulumi.IntInput         `pulumi:"bandwidth"`
	CreateTime      pulumi.StringInput      `pulumi:"createTime"`
	Id              pulumi.StringInput      `pulumi:"id"`
	MaxConcurrent   pulumi.IntInput         `pulumi:"maxConcurrent"`
	Name            pulumi.StringInput      `pulumi:"name"`
	State           pulumi.IntInput         `pulumi:"state"`
	VpcId           pulumi.StringInput      `pulumi:"vpcId"`
}

func (GetInstanceNatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceNat)(nil)).Elem()
}

func (i GetInstanceNatArgs) ToGetInstanceNatOutput() GetInstanceNatOutput {
	return i.ToGetInstanceNatOutputWithContext(context.Background())
}

func (i GetInstanceNatArgs) ToGetInstanceNatOutputWithContext(ctx context.Context) GetInstanceNatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceNatOutput)
}

// GetInstanceNatArrayInput is an input type that accepts GetInstanceNatArray and GetInstanceNatArrayOutput values.
// You can construct a concrete instance of `GetInstanceNatArrayInput` via:
//
//	GetInstanceNatArray{ GetInstanceNatArgs{...} }
type GetInstanceNatArrayInput interface {
	pulumi.Input

	ToGetInstanceNatArrayOutput() GetInstanceNatArrayOutput
	ToGetInstanceNatArrayOutputWithContext(context.Context) GetInstanceNatArrayOutput
}

type GetInstanceNatArray []GetInstanceNatInput

func (GetInstanceNatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceNat)(nil)).Elem()
}

func (i GetInstanceNatArray) ToGetInstanceNatArrayOutput() GetInstanceNatArrayOutput {
	return i.ToGetInstanceNatArrayOutputWithContext(context.Background())
}

func (i GetInstanceNatArray) ToGetInstanceNatArrayOutputWithContext(ctx context.Context) GetInstanceNatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceNatArrayOutput)
}

type GetInstanceNatOutput struct{ *pulumi.OutputState }

func (GetInstanceNatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceNat)(nil)).Elem()
}

func (o GetInstanceNatOutput) ToGetInstanceNatOutput() GetInstanceNatOutput {
	return o
}

func (o GetInstanceNatOutput) ToGetInstanceNatOutputWithContext(ctx context.Context) GetInstanceNatOutput {
	return o
}

func (o GetInstanceNatOutput) AssignedEipSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceNat) []string { return v.AssignedEipSets }).(pulumi.StringArrayOutput)
}

func (o GetInstanceNatOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstanceNat) int { return v.Bandwidth }).(pulumi.IntOutput)
}

func (o GetInstanceNatOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceNat) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o GetInstanceNatOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceNat) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstanceNatOutput) MaxConcurrent() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstanceNat) int { return v.MaxConcurrent }).(pulumi.IntOutput)
}

func (o GetInstanceNatOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceNat) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetInstanceNatOutput) State() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstanceNat) int { return v.State }).(pulumi.IntOutput)
}

func (o GetInstanceNatOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceNat) string { return v.VpcId }).(pulumi.StringOutput)
}

type GetInstanceNatArrayOutput struct{ *pulumi.OutputState }

func (GetInstanceNatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceNat)(nil)).Elem()
}

func (o GetInstanceNatArrayOutput) ToGetInstanceNatArrayOutput() GetInstanceNatArrayOutput {
	return o
}

func (o GetInstanceNatArrayOutput) ToGetInstanceNatArrayOutputWithContext(ctx context.Context) GetInstanceNatArrayOutput {
	return o
}

func (o GetInstanceNatArrayOutput) Index(i pulumi.IntInput) GetInstanceNatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetInstanceNat {
		return vs[0].([]GetInstanceNat)[vs[1].(int)]
	}).(GetInstanceNatOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstanceNatInput)(nil)).Elem(), GetInstanceNatArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstanceNatArrayInput)(nil)).Elem(), GetInstanceNatArray{})
	pulumi.RegisterOutputType(GetInstanceNatOutput{})
	pulumi.RegisterOutputType(GetInstanceNatArrayOutput{})
}
