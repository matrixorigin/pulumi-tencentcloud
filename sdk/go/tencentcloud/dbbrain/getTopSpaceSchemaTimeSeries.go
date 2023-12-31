// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbbrain

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTopSpaceSchemaTimeSeries(ctx *pulumi.Context, args *GetTopSpaceSchemaTimeSeriesArgs, opts ...pulumi.InvokeOption) (*GetTopSpaceSchemaTimeSeriesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTopSpaceSchemaTimeSeriesResult
	err := ctx.Invoke("tencentcloud:Dbbrain/getTopSpaceSchemaTimeSeries:getTopSpaceSchemaTimeSeries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTopSpaceSchemaTimeSeries.
type GetTopSpaceSchemaTimeSeriesArgs struct {
	EndDate          *string `pulumi:"endDate"`
	InstanceId       string  `pulumi:"instanceId"`
	Limit            *int    `pulumi:"limit"`
	Product          *string `pulumi:"product"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SortBy           *string `pulumi:"sortBy"`
	StartDate        *string `pulumi:"startDate"`
}

// A collection of values returned by getTopSpaceSchemaTimeSeries.
type GetTopSpaceSchemaTimeSeriesResult struct {
	EndDate *string `pulumi:"endDate"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string                                                `pulumi:"id"`
	InstanceId               string                                                `pulumi:"instanceId"`
	Limit                    *int                                                  `pulumi:"limit"`
	Product                  *string                                               `pulumi:"product"`
	ResultOutputFile         *string                                               `pulumi:"resultOutputFile"`
	SortBy                   *string                                               `pulumi:"sortBy"`
	StartDate                *string                                               `pulumi:"startDate"`
	TopSpaceSchemaTimeSeries []GetTopSpaceSchemaTimeSeriesTopSpaceSchemaTimeSeries `pulumi:"topSpaceSchemaTimeSeries"`
}

func GetTopSpaceSchemaTimeSeriesOutput(ctx *pulumi.Context, args GetTopSpaceSchemaTimeSeriesOutputArgs, opts ...pulumi.InvokeOption) GetTopSpaceSchemaTimeSeriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTopSpaceSchemaTimeSeriesResult, error) {
			args := v.(GetTopSpaceSchemaTimeSeriesArgs)
			r, err := GetTopSpaceSchemaTimeSeries(ctx, &args, opts...)
			var s GetTopSpaceSchemaTimeSeriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTopSpaceSchemaTimeSeriesResultOutput)
}

// A collection of arguments for invoking getTopSpaceSchemaTimeSeries.
type GetTopSpaceSchemaTimeSeriesOutputArgs struct {
	EndDate          pulumi.StringPtrInput `pulumi:"endDate"`
	InstanceId       pulumi.StringInput    `pulumi:"instanceId"`
	Limit            pulumi.IntPtrInput    `pulumi:"limit"`
	Product          pulumi.StringPtrInput `pulumi:"product"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	SortBy           pulumi.StringPtrInput `pulumi:"sortBy"`
	StartDate        pulumi.StringPtrInput `pulumi:"startDate"`
}

func (GetTopSpaceSchemaTimeSeriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTopSpaceSchemaTimeSeriesArgs)(nil)).Elem()
}

// A collection of values returned by getTopSpaceSchemaTimeSeries.
type GetTopSpaceSchemaTimeSeriesResultOutput struct{ *pulumi.OutputState }

func (GetTopSpaceSchemaTimeSeriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTopSpaceSchemaTimeSeriesResult)(nil)).Elem()
}

func (o GetTopSpaceSchemaTimeSeriesResultOutput) ToGetTopSpaceSchemaTimeSeriesResultOutput() GetTopSpaceSchemaTimeSeriesResultOutput {
	return o
}

func (o GetTopSpaceSchemaTimeSeriesResultOutput) ToGetTopSpaceSchemaTimeSeriesResultOutputWithContext(ctx context.Context) GetTopSpaceSchemaTimeSeriesResultOutput {
	return o
}

func (o GetTopSpaceSchemaTimeSeriesResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopSpaceSchemaTimeSeriesResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTopSpaceSchemaTimeSeriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTopSpaceSchemaTimeSeriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTopSpaceSchemaTimeSeriesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTopSpaceSchemaTimeSeriesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetTopSpaceSchemaTimeSeriesResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetTopSpaceSchemaTimeSeriesResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o GetTopSpaceSchemaTimeSeriesResultOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopSpaceSchemaTimeSeriesResult) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o GetTopSpaceSchemaTimeSeriesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopSpaceSchemaTimeSeriesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetTopSpaceSchemaTimeSeriesResultOutput) SortBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopSpaceSchemaTimeSeriesResult) *string { return v.SortBy }).(pulumi.StringPtrOutput)
}

func (o GetTopSpaceSchemaTimeSeriesResultOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopSpaceSchemaTimeSeriesResult) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o GetTopSpaceSchemaTimeSeriesResultOutput) TopSpaceSchemaTimeSeries() GetTopSpaceSchemaTimeSeriesTopSpaceSchemaTimeSeriesArrayOutput {
	return o.ApplyT(func(v GetTopSpaceSchemaTimeSeriesResult) []GetTopSpaceSchemaTimeSeriesTopSpaceSchemaTimeSeries {
		return v.TopSpaceSchemaTimeSeries
	}).(GetTopSpaceSchemaTimeSeriesTopSpaceSchemaTimeSeriesArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTopSpaceSchemaTimeSeriesResultOutput{})
}
