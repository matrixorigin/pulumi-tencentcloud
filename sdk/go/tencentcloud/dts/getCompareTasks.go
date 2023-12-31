// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dts

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCompareTasks(ctx *pulumi.Context, args *GetCompareTasksArgs, opts ...pulumi.InvokeOption) (*GetCompareTasksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCompareTasksResult
	err := ctx.Invoke("tencentcloud:Dts/getCompareTasks:getCompareTasks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCompareTasks.
type GetCompareTasksArgs struct {
	JobId            string  `pulumi:"jobId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getCompareTasks.
type GetCompareTasksResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                `pulumi:"id"`
	JobId            string                `pulumi:"jobId"`
	Lists            []GetCompareTasksList `pulumi:"lists"`
	ResultOutputFile *string               `pulumi:"resultOutputFile"`
}

func GetCompareTasksOutput(ctx *pulumi.Context, args GetCompareTasksOutputArgs, opts ...pulumi.InvokeOption) GetCompareTasksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCompareTasksResult, error) {
			args := v.(GetCompareTasksArgs)
			r, err := GetCompareTasks(ctx, &args, opts...)
			var s GetCompareTasksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCompareTasksResultOutput)
}

// A collection of arguments for invoking getCompareTasks.
type GetCompareTasksOutputArgs struct {
	JobId            pulumi.StringInput    `pulumi:"jobId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetCompareTasksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCompareTasksArgs)(nil)).Elem()
}

// A collection of values returned by getCompareTasks.
type GetCompareTasksResultOutput struct{ *pulumi.OutputState }

func (GetCompareTasksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCompareTasksResult)(nil)).Elem()
}

func (o GetCompareTasksResultOutput) ToGetCompareTasksResultOutput() GetCompareTasksResultOutput {
	return o
}

func (o GetCompareTasksResultOutput) ToGetCompareTasksResultOutputWithContext(ctx context.Context) GetCompareTasksResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetCompareTasksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCompareTasksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCompareTasksResultOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCompareTasksResult) string { return v.JobId }).(pulumi.StringOutput)
}

func (o GetCompareTasksResultOutput) Lists() GetCompareTasksListArrayOutput {
	return o.ApplyT(func(v GetCompareTasksResult) []GetCompareTasksList { return v.Lists }).(GetCompareTasksListArrayOutput)
}

func (o GetCompareTasksResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCompareTasksResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCompareTasksResultOutput{})
}
