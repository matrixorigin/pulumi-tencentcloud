// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oceanus

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetJobEvents(ctx *pulumi.Context, args *GetJobEventsArgs, opts ...pulumi.InvokeOption) (*GetJobEventsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetJobEventsResult
	err := ctx.Invoke("tencentcloud:Oceanus/getJobEvents:getJobEvents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getJobEvents.
type GetJobEventsArgs struct {
	EndTimestamp     int      `pulumi:"endTimestamp"`
	JobId            string   `pulumi:"jobId"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	StartTimestamp   int      `pulumi:"startTimestamp"`
	Types            []string `pulumi:"types"`
	WorkSpaceId      string   `pulumi:"workSpaceId"`
}

// A collection of values returned by getJobEvents.
type GetJobEventsResult struct {
	EndTimestamp int                 `pulumi:"endTimestamp"`
	Events       []GetJobEventsEvent `pulumi:"events"`
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	JobId            string   `pulumi:"jobId"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	RunningOrderIds  []int    `pulumi:"runningOrderIds"`
	StartTimestamp   int      `pulumi:"startTimestamp"`
	Types            []string `pulumi:"types"`
	WorkSpaceId      string   `pulumi:"workSpaceId"`
}

func GetJobEventsOutput(ctx *pulumi.Context, args GetJobEventsOutputArgs, opts ...pulumi.InvokeOption) GetJobEventsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetJobEventsResult, error) {
			args := v.(GetJobEventsArgs)
			r, err := GetJobEvents(ctx, &args, opts...)
			var s GetJobEventsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetJobEventsResultOutput)
}

// A collection of arguments for invoking getJobEvents.
type GetJobEventsOutputArgs struct {
	EndTimestamp     pulumi.IntInput         `pulumi:"endTimestamp"`
	JobId            pulumi.StringInput      `pulumi:"jobId"`
	ResultOutputFile pulumi.StringPtrInput   `pulumi:"resultOutputFile"`
	StartTimestamp   pulumi.IntInput         `pulumi:"startTimestamp"`
	Types            pulumi.StringArrayInput `pulumi:"types"`
	WorkSpaceId      pulumi.StringInput      `pulumi:"workSpaceId"`
}

func (GetJobEventsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetJobEventsArgs)(nil)).Elem()
}

// A collection of values returned by getJobEvents.
type GetJobEventsResultOutput struct{ *pulumi.OutputState }

func (GetJobEventsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetJobEventsResult)(nil)).Elem()
}

func (o GetJobEventsResultOutput) ToGetJobEventsResultOutput() GetJobEventsResultOutput {
	return o
}

func (o GetJobEventsResultOutput) ToGetJobEventsResultOutputWithContext(ctx context.Context) GetJobEventsResultOutput {
	return o
}

func (o GetJobEventsResultOutput) EndTimestamp() pulumi.IntOutput {
	return o.ApplyT(func(v GetJobEventsResult) int { return v.EndTimestamp }).(pulumi.IntOutput)
}

func (o GetJobEventsResultOutput) Events() GetJobEventsEventArrayOutput {
	return o.ApplyT(func(v GetJobEventsResult) []GetJobEventsEvent { return v.Events }).(GetJobEventsEventArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetJobEventsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetJobEventsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetJobEventsResultOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v GetJobEventsResult) string { return v.JobId }).(pulumi.StringOutput)
}

func (o GetJobEventsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetJobEventsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetJobEventsResultOutput) RunningOrderIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetJobEventsResult) []int { return v.RunningOrderIds }).(pulumi.IntArrayOutput)
}

func (o GetJobEventsResultOutput) StartTimestamp() pulumi.IntOutput {
	return o.ApplyT(func(v GetJobEventsResult) int { return v.StartTimestamp }).(pulumi.IntOutput)
}

func (o GetJobEventsResultOutput) Types() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetJobEventsResult) []string { return v.Types }).(pulumi.StringArrayOutput)
}

func (o GetJobEventsResultOutput) WorkSpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetJobEventsResult) string { return v.WorkSpaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetJobEventsResultOutput{})
}
