// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetProxyGroupStatistics(ctx *pulumi.Context, args *GetProxyGroupStatisticsArgs, opts ...pulumi.InvokeOption) (*GetProxyGroupStatisticsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProxyGroupStatisticsResult
	err := ctx.Invoke("tencentcloud:Gaap/getProxyGroupStatistics:getProxyGroupStatistics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProxyGroupStatistics.
type GetProxyGroupStatisticsArgs struct {
	EndTime          string   `pulumi:"endTime"`
	Granularity      int      `pulumi:"granularity"`
	GroupId          string   `pulumi:"groupId"`
	MetricNames      []string `pulumi:"metricNames"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	StartTime        string   `pulumi:"startTime"`
}

// A collection of values returned by getProxyGroupStatistics.
type GetProxyGroupStatisticsResult struct {
	EndTime     string `pulumi:"endTime"`
	Granularity int    `pulumi:"granularity"`
	GroupId     string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                                  `pulumi:"id"`
	MetricNames      []string                                `pulumi:"metricNames"`
	ResultOutputFile *string                                 `pulumi:"resultOutputFile"`
	StartTime        string                                  `pulumi:"startTime"`
	StatisticsDatas  []GetProxyGroupStatisticsStatisticsData `pulumi:"statisticsDatas"`
}

func GetProxyGroupStatisticsOutput(ctx *pulumi.Context, args GetProxyGroupStatisticsOutputArgs, opts ...pulumi.InvokeOption) GetProxyGroupStatisticsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProxyGroupStatisticsResult, error) {
			args := v.(GetProxyGroupStatisticsArgs)
			r, err := GetProxyGroupStatistics(ctx, &args, opts...)
			var s GetProxyGroupStatisticsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProxyGroupStatisticsResultOutput)
}

// A collection of arguments for invoking getProxyGroupStatistics.
type GetProxyGroupStatisticsOutputArgs struct {
	EndTime          pulumi.StringInput      `pulumi:"endTime"`
	Granularity      pulumi.IntInput         `pulumi:"granularity"`
	GroupId          pulumi.StringInput      `pulumi:"groupId"`
	MetricNames      pulumi.StringArrayInput `pulumi:"metricNames"`
	ResultOutputFile pulumi.StringPtrInput   `pulumi:"resultOutputFile"`
	StartTime        pulumi.StringInput      `pulumi:"startTime"`
}

func (GetProxyGroupStatisticsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxyGroupStatisticsArgs)(nil)).Elem()
}

// A collection of values returned by getProxyGroupStatistics.
type GetProxyGroupStatisticsResultOutput struct{ *pulumi.OutputState }

func (GetProxyGroupStatisticsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxyGroupStatisticsResult)(nil)).Elem()
}

func (o GetProxyGroupStatisticsResultOutput) ToGetProxyGroupStatisticsResultOutput() GetProxyGroupStatisticsResultOutput {
	return o
}

func (o GetProxyGroupStatisticsResultOutput) ToGetProxyGroupStatisticsResultOutputWithContext(ctx context.Context) GetProxyGroupStatisticsResultOutput {
	return o
}

func (o GetProxyGroupStatisticsResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxyGroupStatisticsResult) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o GetProxyGroupStatisticsResultOutput) Granularity() pulumi.IntOutput {
	return o.ApplyT(func(v GetProxyGroupStatisticsResult) int { return v.Granularity }).(pulumi.IntOutput)
}

func (o GetProxyGroupStatisticsResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxyGroupStatisticsResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProxyGroupStatisticsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxyGroupStatisticsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetProxyGroupStatisticsResultOutput) MetricNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProxyGroupStatisticsResult) []string { return v.MetricNames }).(pulumi.StringArrayOutput)
}

func (o GetProxyGroupStatisticsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProxyGroupStatisticsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetProxyGroupStatisticsResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxyGroupStatisticsResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o GetProxyGroupStatisticsResultOutput) StatisticsDatas() GetProxyGroupStatisticsStatisticsDataArrayOutput {
	return o.ApplyT(func(v GetProxyGroupStatisticsResult) []GetProxyGroupStatisticsStatisticsData {
		return v.StatisticsDatas
	}).(GetProxyGroupStatisticsStatisticsDataArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProxyGroupStatisticsResultOutput{})
}
