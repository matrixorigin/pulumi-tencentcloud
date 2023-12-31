// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package placement

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type GetGroupsPlacementGroupList struct {
	CreateTime       string   `pulumi:"createTime"`
	CurrentNum       int      `pulumi:"currentNum"`
	CvmQuotaTotal    int      `pulumi:"cvmQuotaTotal"`
	InstanceIds      []string `pulumi:"instanceIds"`
	Name             string   `pulumi:"name"`
	PlacementGroupId string   `pulumi:"placementGroupId"`
	Type             string   `pulumi:"type"`
}

// GetGroupsPlacementGroupListInput is an input type that accepts GetGroupsPlacementGroupListArgs and GetGroupsPlacementGroupListOutput values.
// You can construct a concrete instance of `GetGroupsPlacementGroupListInput` via:
//
//	GetGroupsPlacementGroupListArgs{...}
type GetGroupsPlacementGroupListInput interface {
	pulumi.Input

	ToGetGroupsPlacementGroupListOutput() GetGroupsPlacementGroupListOutput
	ToGetGroupsPlacementGroupListOutputWithContext(context.Context) GetGroupsPlacementGroupListOutput
}

type GetGroupsPlacementGroupListArgs struct {
	CreateTime       pulumi.StringInput      `pulumi:"createTime"`
	CurrentNum       pulumi.IntInput         `pulumi:"currentNum"`
	CvmQuotaTotal    pulumi.IntInput         `pulumi:"cvmQuotaTotal"`
	InstanceIds      pulumi.StringArrayInput `pulumi:"instanceIds"`
	Name             pulumi.StringInput      `pulumi:"name"`
	PlacementGroupId pulumi.StringInput      `pulumi:"placementGroupId"`
	Type             pulumi.StringInput      `pulumi:"type"`
}

func (GetGroupsPlacementGroupListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsPlacementGroupList)(nil)).Elem()
}

func (i GetGroupsPlacementGroupListArgs) ToGetGroupsPlacementGroupListOutput() GetGroupsPlacementGroupListOutput {
	return i.ToGetGroupsPlacementGroupListOutputWithContext(context.Background())
}

func (i GetGroupsPlacementGroupListArgs) ToGetGroupsPlacementGroupListOutputWithContext(ctx context.Context) GetGroupsPlacementGroupListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupsPlacementGroupListOutput)
}

// GetGroupsPlacementGroupListArrayInput is an input type that accepts GetGroupsPlacementGroupListArray and GetGroupsPlacementGroupListArrayOutput values.
// You can construct a concrete instance of `GetGroupsPlacementGroupListArrayInput` via:
//
//	GetGroupsPlacementGroupListArray{ GetGroupsPlacementGroupListArgs{...} }
type GetGroupsPlacementGroupListArrayInput interface {
	pulumi.Input

	ToGetGroupsPlacementGroupListArrayOutput() GetGroupsPlacementGroupListArrayOutput
	ToGetGroupsPlacementGroupListArrayOutputWithContext(context.Context) GetGroupsPlacementGroupListArrayOutput
}

type GetGroupsPlacementGroupListArray []GetGroupsPlacementGroupListInput

func (GetGroupsPlacementGroupListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupsPlacementGroupList)(nil)).Elem()
}

func (i GetGroupsPlacementGroupListArray) ToGetGroupsPlacementGroupListArrayOutput() GetGroupsPlacementGroupListArrayOutput {
	return i.ToGetGroupsPlacementGroupListArrayOutputWithContext(context.Background())
}

func (i GetGroupsPlacementGroupListArray) ToGetGroupsPlacementGroupListArrayOutputWithContext(ctx context.Context) GetGroupsPlacementGroupListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupsPlacementGroupListArrayOutput)
}

type GetGroupsPlacementGroupListOutput struct{ *pulumi.OutputState }

func (GetGroupsPlacementGroupListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsPlacementGroupList)(nil)).Elem()
}

func (o GetGroupsPlacementGroupListOutput) ToGetGroupsPlacementGroupListOutput() GetGroupsPlacementGroupListOutput {
	return o
}

func (o GetGroupsPlacementGroupListOutput) ToGetGroupsPlacementGroupListOutputWithContext(ctx context.Context) GetGroupsPlacementGroupListOutput {
	return o
}

func (o GetGroupsPlacementGroupListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsPlacementGroupList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o GetGroupsPlacementGroupListOutput) CurrentNum() pulumi.IntOutput {
	return o.ApplyT(func(v GetGroupsPlacementGroupList) int { return v.CurrentNum }).(pulumi.IntOutput)
}

func (o GetGroupsPlacementGroupListOutput) CvmQuotaTotal() pulumi.IntOutput {
	return o.ApplyT(func(v GetGroupsPlacementGroupList) int { return v.CvmQuotaTotal }).(pulumi.IntOutput)
}

func (o GetGroupsPlacementGroupListOutput) InstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGroupsPlacementGroupList) []string { return v.InstanceIds }).(pulumi.StringArrayOutput)
}

func (o GetGroupsPlacementGroupListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsPlacementGroupList) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetGroupsPlacementGroupListOutput) PlacementGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsPlacementGroupList) string { return v.PlacementGroupId }).(pulumi.StringOutput)
}

func (o GetGroupsPlacementGroupListOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsPlacementGroupList) string { return v.Type }).(pulumi.StringOutput)
}

type GetGroupsPlacementGroupListArrayOutput struct{ *pulumi.OutputState }

func (GetGroupsPlacementGroupListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupsPlacementGroupList)(nil)).Elem()
}

func (o GetGroupsPlacementGroupListArrayOutput) ToGetGroupsPlacementGroupListArrayOutput() GetGroupsPlacementGroupListArrayOutput {
	return o
}

func (o GetGroupsPlacementGroupListArrayOutput) ToGetGroupsPlacementGroupListArrayOutputWithContext(ctx context.Context) GetGroupsPlacementGroupListArrayOutput {
	return o
}

func (o GetGroupsPlacementGroupListArrayOutput) Index(i pulumi.IntInput) GetGroupsPlacementGroupListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetGroupsPlacementGroupList {
		return vs[0].([]GetGroupsPlacementGroupList)[vs[1].(int)]
	}).(GetGroupsPlacementGroupListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetGroupsPlacementGroupListInput)(nil)).Elem(), GetGroupsPlacementGroupListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGroupsPlacementGroupListArrayInput)(nil)).Elem(), GetGroupsPlacementGroupListArray{})
	pulumi.RegisterOutputType(GetGroupsPlacementGroupListOutput{})
	pulumi.RegisterOutputType(GetGroupsPlacementGroupListArrayOutput{})
}
