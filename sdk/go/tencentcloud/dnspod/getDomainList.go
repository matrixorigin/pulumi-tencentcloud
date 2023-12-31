// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dnspod

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDomainList(ctx *pulumi.Context, args *GetDomainListArgs, opts ...pulumi.InvokeOption) (*GetDomainListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDomainListResult
	err := ctx.Invoke("tencentcloud:Dnspod/getDomainList:getDomainList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomainList.
type GetDomainListArgs struct {
	GroupIds         []int              `pulumi:"groupIds"`
	Keyword          *string            `pulumi:"keyword"`
	Packages         []string           `pulumi:"packages"`
	ProjectId        *int               `pulumi:"projectId"`
	RecordCountBegin *int               `pulumi:"recordCountBegin"`
	RecordCountEnd   *int               `pulumi:"recordCountEnd"`
	Remark           *string            `pulumi:"remark"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
	SortField        *string            `pulumi:"sortField"`
	SortType         *string            `pulumi:"sortType"`
	Statuses         []string           `pulumi:"statuses"`
	Tags             []GetDomainListTag `pulumi:"tags"`
	Type             string             `pulumi:"type"`
	UpdatedAtBegin   *string            `pulumi:"updatedAtBegin"`
	UpdatedAtEnd     *string            `pulumi:"updatedAtEnd"`
}

// A collection of values returned by getDomainList.
type GetDomainListResult struct {
	DomainLists []GetDomainListDomainList `pulumi:"domainLists"`
	GroupIds    []int                     `pulumi:"groupIds"`
	// The provider-assigned unique ID for this managed resource.
	Id               string             `pulumi:"id"`
	Keyword          *string            `pulumi:"keyword"`
	Packages         []string           `pulumi:"packages"`
	ProjectId        *int               `pulumi:"projectId"`
	RecordCountBegin *int               `pulumi:"recordCountBegin"`
	RecordCountEnd   *int               `pulumi:"recordCountEnd"`
	Remark           *string            `pulumi:"remark"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
	SortField        *string            `pulumi:"sortField"`
	SortType         *string            `pulumi:"sortType"`
	Statuses         []string           `pulumi:"statuses"`
	Tags             []GetDomainListTag `pulumi:"tags"`
	Type             string             `pulumi:"type"`
	UpdatedAtBegin   *string            `pulumi:"updatedAtBegin"`
	UpdatedAtEnd     *string            `pulumi:"updatedAtEnd"`
}

func GetDomainListOutput(ctx *pulumi.Context, args GetDomainListOutputArgs, opts ...pulumi.InvokeOption) GetDomainListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainListResult, error) {
			args := v.(GetDomainListArgs)
			r, err := GetDomainList(ctx, &args, opts...)
			var s GetDomainListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainListResultOutput)
}

// A collection of arguments for invoking getDomainList.
type GetDomainListOutputArgs struct {
	GroupIds         pulumi.IntArrayInput       `pulumi:"groupIds"`
	Keyword          pulumi.StringPtrInput      `pulumi:"keyword"`
	Packages         pulumi.StringArrayInput    `pulumi:"packages"`
	ProjectId        pulumi.IntPtrInput         `pulumi:"projectId"`
	RecordCountBegin pulumi.IntPtrInput         `pulumi:"recordCountBegin"`
	RecordCountEnd   pulumi.IntPtrInput         `pulumi:"recordCountEnd"`
	Remark           pulumi.StringPtrInput      `pulumi:"remark"`
	ResultOutputFile pulumi.StringPtrInput      `pulumi:"resultOutputFile"`
	SortField        pulumi.StringPtrInput      `pulumi:"sortField"`
	SortType         pulumi.StringPtrInput      `pulumi:"sortType"`
	Statuses         pulumi.StringArrayInput    `pulumi:"statuses"`
	Tags             GetDomainListTagArrayInput `pulumi:"tags"`
	Type             pulumi.StringInput         `pulumi:"type"`
	UpdatedAtBegin   pulumi.StringPtrInput      `pulumi:"updatedAtBegin"`
	UpdatedAtEnd     pulumi.StringPtrInput      `pulumi:"updatedAtEnd"`
}

func (GetDomainListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainListArgs)(nil)).Elem()
}

// A collection of values returned by getDomainList.
type GetDomainListResultOutput struct{ *pulumi.OutputState }

func (GetDomainListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainListResult)(nil)).Elem()
}

func (o GetDomainListResultOutput) ToGetDomainListResultOutput() GetDomainListResultOutput {
	return o
}

func (o GetDomainListResultOutput) ToGetDomainListResultOutputWithContext(ctx context.Context) GetDomainListResultOutput {
	return o
}

func (o GetDomainListResultOutput) DomainLists() GetDomainListDomainListArrayOutput {
	return o.ApplyT(func(v GetDomainListResult) []GetDomainListDomainList { return v.DomainLists }).(GetDomainListDomainListArrayOutput)
}

func (o GetDomainListResultOutput) GroupIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetDomainListResult) []int { return v.GroupIds }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDomainListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDomainListResultOutput) Keyword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainListResult) *string { return v.Keyword }).(pulumi.StringPtrOutput)
}

func (o GetDomainListResultOutput) Packages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDomainListResult) []string { return v.Packages }).(pulumi.StringArrayOutput)
}

func (o GetDomainListResultOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDomainListResult) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

func (o GetDomainListResultOutput) RecordCountBegin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDomainListResult) *int { return v.RecordCountBegin }).(pulumi.IntPtrOutput)
}

func (o GetDomainListResultOutput) RecordCountEnd() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDomainListResult) *int { return v.RecordCountEnd }).(pulumi.IntPtrOutput)
}

func (o GetDomainListResultOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainListResult) *string { return v.Remark }).(pulumi.StringPtrOutput)
}

func (o GetDomainListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetDomainListResultOutput) SortField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainListResult) *string { return v.SortField }).(pulumi.StringPtrOutput)
}

func (o GetDomainListResultOutput) SortType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainListResult) *string { return v.SortType }).(pulumi.StringPtrOutput)
}

func (o GetDomainListResultOutput) Statuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDomainListResult) []string { return v.Statuses }).(pulumi.StringArrayOutput)
}

func (o GetDomainListResultOutput) Tags() GetDomainListTagArrayOutput {
	return o.ApplyT(func(v GetDomainListResult) []GetDomainListTag { return v.Tags }).(GetDomainListTagArrayOutput)
}

func (o GetDomainListResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainListResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetDomainListResultOutput) UpdatedAtBegin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainListResult) *string { return v.UpdatedAtBegin }).(pulumi.StringPtrOutput)
}

func (o GetDomainListResultOutput) UpdatedAtEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainListResult) *string { return v.UpdatedAtEnd }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainListResultOutput{})
}
