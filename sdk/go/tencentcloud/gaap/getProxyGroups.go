// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetProxyGroups(ctx *pulumi.Context, args *GetProxyGroupsArgs, opts ...pulumi.InvokeOption) (*GetProxyGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProxyGroupsResult
	err := ctx.Invoke("tencentcloud:Gaap/getProxyGroups:getProxyGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProxyGroups.
type GetProxyGroupsArgs struct {
	Filters          []GetProxyGroupsFilter `pulumi:"filters"`
	ProjectId        int                    `pulumi:"projectId"`
	ResultOutputFile *string                `pulumi:"resultOutputFile"`
	TagSets          []GetProxyGroupsTagSet `pulumi:"tagSets"`
}

// A collection of values returned by getProxyGroups.
type GetProxyGroupsResult struct {
	Filters []GetProxyGroupsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                         `pulumi:"id"`
	ProjectId        int                            `pulumi:"projectId"`
	ProxyGroupLists  []GetProxyGroupsProxyGroupList `pulumi:"proxyGroupLists"`
	ResultOutputFile *string                        `pulumi:"resultOutputFile"`
	TagSets          []GetProxyGroupsTagSet         `pulumi:"tagSets"`
}

func GetProxyGroupsOutput(ctx *pulumi.Context, args GetProxyGroupsOutputArgs, opts ...pulumi.InvokeOption) GetProxyGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProxyGroupsResult, error) {
			args := v.(GetProxyGroupsArgs)
			r, err := GetProxyGroups(ctx, &args, opts...)
			var s GetProxyGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProxyGroupsResultOutput)
}

// A collection of arguments for invoking getProxyGroups.
type GetProxyGroupsOutputArgs struct {
	Filters          GetProxyGroupsFilterArrayInput `pulumi:"filters"`
	ProjectId        pulumi.IntInput                `pulumi:"projectId"`
	ResultOutputFile pulumi.StringPtrInput          `pulumi:"resultOutputFile"`
	TagSets          GetProxyGroupsTagSetArrayInput `pulumi:"tagSets"`
}

func (GetProxyGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxyGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getProxyGroups.
type GetProxyGroupsResultOutput struct{ *pulumi.OutputState }

func (GetProxyGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxyGroupsResult)(nil)).Elem()
}

func (o GetProxyGroupsResultOutput) ToGetProxyGroupsResultOutput() GetProxyGroupsResultOutput {
	return o
}

func (o GetProxyGroupsResultOutput) ToGetProxyGroupsResultOutputWithContext(ctx context.Context) GetProxyGroupsResultOutput {
	return o
}

func (o GetProxyGroupsResultOutput) Filters() GetProxyGroupsFilterArrayOutput {
	return o.ApplyT(func(v GetProxyGroupsResult) []GetProxyGroupsFilter { return v.Filters }).(GetProxyGroupsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProxyGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxyGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetProxyGroupsResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v GetProxyGroupsResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

func (o GetProxyGroupsResultOutput) ProxyGroupLists() GetProxyGroupsProxyGroupListArrayOutput {
	return o.ApplyT(func(v GetProxyGroupsResult) []GetProxyGroupsProxyGroupList { return v.ProxyGroupLists }).(GetProxyGroupsProxyGroupListArrayOutput)
}

func (o GetProxyGroupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProxyGroupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetProxyGroupsResultOutput) TagSets() GetProxyGroupsTagSetArrayOutput {
	return o.ApplyT(func(v GetProxyGroupsResult) []GetProxyGroupsTagSet { return v.TagSets }).(GetProxyGroupsTagSetArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProxyGroupsResultOutput{})
}