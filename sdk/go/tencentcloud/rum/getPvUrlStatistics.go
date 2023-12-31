// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rum

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPvUrlStatistics(ctx *pulumi.Context, args *GetPvUrlStatisticsArgs, opts ...pulumi.InvokeOption) (*GetPvUrlStatisticsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPvUrlStatisticsResult
	err := ctx.Invoke("tencentcloud:Rum/getPvUrlStatistics:getPvUrlStatistics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPvUrlStatistics.
type GetPvUrlStatisticsArgs struct {
	Area             *string `pulumi:"area"`
	Brand            *string `pulumi:"brand"`
	Browser          *string `pulumi:"browser"`
	Device           *string `pulumi:"device"`
	EndTime          int     `pulumi:"endTime"`
	Engine           *string `pulumi:"engine"`
	Env              *string `pulumi:"env"`
	ExtFirst         *string `pulumi:"extFirst"`
	ExtSecond        *string `pulumi:"extSecond"`
	ExtThird         *string `pulumi:"extThird"`
	From             *string `pulumi:"from"`
	GroupByType      *int    `pulumi:"groupByType"`
	IsAbroad         *string `pulumi:"isAbroad"`
	Isp              *string `pulumi:"isp"`
	Level            *string `pulumi:"level"`
	NetType          *string `pulumi:"netType"`
	Os               *string `pulumi:"os"`
	Platform         *string `pulumi:"platform"`
	ProjectId        int     `pulumi:"projectId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	StartTime        int     `pulumi:"startTime"`
	Type             string  `pulumi:"type"`
	VersionNum       *string `pulumi:"versionNum"`
}

// A collection of values returned by getPvUrlStatistics.
type GetPvUrlStatisticsResult struct {
	Area        *string `pulumi:"area"`
	Brand       *string `pulumi:"brand"`
	Browser     *string `pulumi:"browser"`
	Device      *string `pulumi:"device"`
	EndTime     int     `pulumi:"endTime"`
	Engine      *string `pulumi:"engine"`
	Env         *string `pulumi:"env"`
	ExtFirst    *string `pulumi:"extFirst"`
	ExtSecond   *string `pulumi:"extSecond"`
	ExtThird    *string `pulumi:"extThird"`
	From        *string `pulumi:"from"`
	GroupByType *int    `pulumi:"groupByType"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	IsAbroad         *string `pulumi:"isAbroad"`
	Isp              *string `pulumi:"isp"`
	Level            *string `pulumi:"level"`
	NetType          *string `pulumi:"netType"`
	Os               *string `pulumi:"os"`
	Platform         *string `pulumi:"platform"`
	ProjectId        int     `pulumi:"projectId"`
	Result           string  `pulumi:"result"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	StartTime        int     `pulumi:"startTime"`
	Type             string  `pulumi:"type"`
	VersionNum       *string `pulumi:"versionNum"`
}

func GetPvUrlStatisticsOutput(ctx *pulumi.Context, args GetPvUrlStatisticsOutputArgs, opts ...pulumi.InvokeOption) GetPvUrlStatisticsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPvUrlStatisticsResult, error) {
			args := v.(GetPvUrlStatisticsArgs)
			r, err := GetPvUrlStatistics(ctx, &args, opts...)
			var s GetPvUrlStatisticsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPvUrlStatisticsResultOutput)
}

// A collection of arguments for invoking getPvUrlStatistics.
type GetPvUrlStatisticsOutputArgs struct {
	Area             pulumi.StringPtrInput `pulumi:"area"`
	Brand            pulumi.StringPtrInput `pulumi:"brand"`
	Browser          pulumi.StringPtrInput `pulumi:"browser"`
	Device           pulumi.StringPtrInput `pulumi:"device"`
	EndTime          pulumi.IntInput       `pulumi:"endTime"`
	Engine           pulumi.StringPtrInput `pulumi:"engine"`
	Env              pulumi.StringPtrInput `pulumi:"env"`
	ExtFirst         pulumi.StringPtrInput `pulumi:"extFirst"`
	ExtSecond        pulumi.StringPtrInput `pulumi:"extSecond"`
	ExtThird         pulumi.StringPtrInput `pulumi:"extThird"`
	From             pulumi.StringPtrInput `pulumi:"from"`
	GroupByType      pulumi.IntPtrInput    `pulumi:"groupByType"`
	IsAbroad         pulumi.StringPtrInput `pulumi:"isAbroad"`
	Isp              pulumi.StringPtrInput `pulumi:"isp"`
	Level            pulumi.StringPtrInput `pulumi:"level"`
	NetType          pulumi.StringPtrInput `pulumi:"netType"`
	Os               pulumi.StringPtrInput `pulumi:"os"`
	Platform         pulumi.StringPtrInput `pulumi:"platform"`
	ProjectId        pulumi.IntInput       `pulumi:"projectId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	StartTime        pulumi.IntInput       `pulumi:"startTime"`
	Type             pulumi.StringInput    `pulumi:"type"`
	VersionNum       pulumi.StringPtrInput `pulumi:"versionNum"`
}

func (GetPvUrlStatisticsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPvUrlStatisticsArgs)(nil)).Elem()
}

// A collection of values returned by getPvUrlStatistics.
type GetPvUrlStatisticsResultOutput struct{ *pulumi.OutputState }

func (GetPvUrlStatisticsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPvUrlStatisticsResult)(nil)).Elem()
}

func (o GetPvUrlStatisticsResultOutput) ToGetPvUrlStatisticsResultOutput() GetPvUrlStatisticsResultOutput {
	return o
}

func (o GetPvUrlStatisticsResultOutput) ToGetPvUrlStatisticsResultOutputWithContext(ctx context.Context) GetPvUrlStatisticsResultOutput {
	return o
}

func (o GetPvUrlStatisticsResultOutput) Area() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.Area }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) Brand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.Brand }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) Browser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.Browser }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) EndTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) int { return v.EndTime }).(pulumi.IntOutput)
}

func (o GetPvUrlStatisticsResultOutput) Engine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.Engine }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) Env() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.Env }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) ExtFirst() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.ExtFirst }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) ExtSecond() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.ExtSecond }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) ExtThird() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.ExtThird }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.From }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) GroupByType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *int { return v.GroupByType }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPvUrlStatisticsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPvUrlStatisticsResultOutput) IsAbroad() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.IsAbroad }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) Isp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.Isp }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) NetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.NetType }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.Os }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) Platform() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.Platform }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

func (o GetPvUrlStatisticsResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) string { return v.Result }).(pulumi.StringOutput)
}

func (o GetPvUrlStatisticsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetPvUrlStatisticsResultOutput) StartTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) int { return v.StartTime }).(pulumi.IntOutput)
}

func (o GetPvUrlStatisticsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetPvUrlStatisticsResultOutput) VersionNum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPvUrlStatisticsResult) *string { return v.VersionNum }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPvUrlStatisticsResultOutput{})
}
