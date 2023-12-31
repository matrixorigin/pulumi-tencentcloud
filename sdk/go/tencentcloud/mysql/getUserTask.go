// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetUserTask(ctx *pulumi.Context, args *GetUserTaskArgs, opts ...pulumi.InvokeOption) (*GetUserTaskResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUserTaskResult
	err := ctx.Invoke("tencentcloud:Mysql/getUserTask:getUserTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserTask.
type GetUserTaskArgs struct {
	AsyncRequestId   *string  `pulumi:"asyncRequestId"`
	InstanceId       *string  `pulumi:"instanceId"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	StartTimeBegin   *string  `pulumi:"startTimeBegin"`
	StartTimeEnd     *string  `pulumi:"startTimeEnd"`
	TaskStatuses     []string `pulumi:"taskStatuses"`
	TaskTypes        []string `pulumi:"taskTypes"`
}

// A collection of values returned by getUserTask.
type GetUserTaskResult struct {
	AsyncRequestId *string `pulumi:"asyncRequestId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string            `pulumi:"id"`
	InstanceId       *string           `pulumi:"instanceId"`
	Items            []GetUserTaskItem `pulumi:"items"`
	ResultOutputFile *string           `pulumi:"resultOutputFile"`
	StartTimeBegin   *string           `pulumi:"startTimeBegin"`
	StartTimeEnd     *string           `pulumi:"startTimeEnd"`
	TaskStatuses     []string          `pulumi:"taskStatuses"`
	TaskTypes        []string          `pulumi:"taskTypes"`
}

func GetUserTaskOutput(ctx *pulumi.Context, args GetUserTaskOutputArgs, opts ...pulumi.InvokeOption) GetUserTaskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserTaskResult, error) {
			args := v.(GetUserTaskArgs)
			r, err := GetUserTask(ctx, &args, opts...)
			var s GetUserTaskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserTaskResultOutput)
}

// A collection of arguments for invoking getUserTask.
type GetUserTaskOutputArgs struct {
	AsyncRequestId   pulumi.StringPtrInput   `pulumi:"asyncRequestId"`
	InstanceId       pulumi.StringPtrInput   `pulumi:"instanceId"`
	ResultOutputFile pulumi.StringPtrInput   `pulumi:"resultOutputFile"`
	StartTimeBegin   pulumi.StringPtrInput   `pulumi:"startTimeBegin"`
	StartTimeEnd     pulumi.StringPtrInput   `pulumi:"startTimeEnd"`
	TaskStatuses     pulumi.StringArrayInput `pulumi:"taskStatuses"`
	TaskTypes        pulumi.StringArrayInput `pulumi:"taskTypes"`
}

func (GetUserTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserTaskArgs)(nil)).Elem()
}

// A collection of values returned by getUserTask.
type GetUserTaskResultOutput struct{ *pulumi.OutputState }

func (GetUserTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserTaskResult)(nil)).Elem()
}

func (o GetUserTaskResultOutput) ToGetUserTaskResultOutput() GetUserTaskResultOutput {
	return o
}

func (o GetUserTaskResultOutput) ToGetUserTaskResultOutputWithContext(ctx context.Context) GetUserTaskResultOutput {
	return o
}

func (o GetUserTaskResultOutput) AsyncRequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserTaskResult) *string { return v.AsyncRequestId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetUserTaskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserTaskResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUserTaskResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserTaskResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GetUserTaskResultOutput) Items() GetUserTaskItemArrayOutput {
	return o.ApplyT(func(v GetUserTaskResult) []GetUserTaskItem { return v.Items }).(GetUserTaskItemArrayOutput)
}

func (o GetUserTaskResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserTaskResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetUserTaskResultOutput) StartTimeBegin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserTaskResult) *string { return v.StartTimeBegin }).(pulumi.StringPtrOutput)
}

func (o GetUserTaskResultOutput) StartTimeEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserTaskResult) *string { return v.StartTimeEnd }).(pulumi.StringPtrOutput)
}

func (o GetUserTaskResultOutput) TaskStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUserTaskResult) []string { return v.TaskStatuses }).(pulumi.StringArrayOutput)
}

func (o GetUserTaskResultOutput) TaskTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUserTaskResult) []string { return v.TaskTypes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserTaskResultOutput{})
}
