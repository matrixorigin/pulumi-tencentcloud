// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAlarmHistory(ctx *pulumi.Context, args *GetAlarmHistoryArgs, opts ...pulumi.InvokeOption) (*GetAlarmHistoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAlarmHistoryResult
	err := ctx.Invoke("tencentcloud:Monitor/getAlarmHistory:getAlarmHistory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlarmHistory.
type GetAlarmHistoryArgs struct {
	AlarmLevels      []string                   `pulumi:"alarmLevels"`
	AlarmObject      *string                    `pulumi:"alarmObject"`
	AlarmStatuses    []string                   `pulumi:"alarmStatuses"`
	Content          *string                    `pulumi:"content"`
	EndTime          *int                       `pulumi:"endTime"`
	InstanceGroupIds []int                      `pulumi:"instanceGroupIds"`
	MetricNames      []string                   `pulumi:"metricNames"`
	Module           string                     `pulumi:"module"`
	MonitorTypes     []string                   `pulumi:"monitorTypes"`
	Namespaces       []GetAlarmHistoryNamespace `pulumi:"namespaces"`
	Order            *string                    `pulumi:"order"`
	PolicyIds        []string                   `pulumi:"policyIds"`
	PolicyName       *string                    `pulumi:"policyName"`
	ProjectIds       []int                      `pulumi:"projectIds"`
	ReceiverGroups   []int                      `pulumi:"receiverGroups"`
	ReceiverUids     []int                      `pulumi:"receiverUids"`
	ResultOutputFile *string                    `pulumi:"resultOutputFile"`
	StartTime        *int                       `pulumi:"startTime"`
}

// A collection of values returned by getAlarmHistory.
type GetAlarmHistoryResult struct {
	AlarmLevels   []string                 `pulumi:"alarmLevels"`
	AlarmObject   *string                  `pulumi:"alarmObject"`
	AlarmStatuses []string                 `pulumi:"alarmStatuses"`
	Content       *string                  `pulumi:"content"`
	EndTime       *int                     `pulumi:"endTime"`
	Histories     []GetAlarmHistoryHistory `pulumi:"histories"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                     `pulumi:"id"`
	InstanceGroupIds []int                      `pulumi:"instanceGroupIds"`
	MetricNames      []string                   `pulumi:"metricNames"`
	Module           string                     `pulumi:"module"`
	MonitorTypes     []string                   `pulumi:"monitorTypes"`
	Namespaces       []GetAlarmHistoryNamespace `pulumi:"namespaces"`
	Order            *string                    `pulumi:"order"`
	PolicyIds        []string                   `pulumi:"policyIds"`
	PolicyName       *string                    `pulumi:"policyName"`
	ProjectIds       []int                      `pulumi:"projectIds"`
	ReceiverGroups   []int                      `pulumi:"receiverGroups"`
	ReceiverUids     []int                      `pulumi:"receiverUids"`
	ResultOutputFile *string                    `pulumi:"resultOutputFile"`
	StartTime        *int                       `pulumi:"startTime"`
}

func GetAlarmHistoryOutput(ctx *pulumi.Context, args GetAlarmHistoryOutputArgs, opts ...pulumi.InvokeOption) GetAlarmHistoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAlarmHistoryResult, error) {
			args := v.(GetAlarmHistoryArgs)
			r, err := GetAlarmHistory(ctx, &args, opts...)
			var s GetAlarmHistoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAlarmHistoryResultOutput)
}

// A collection of arguments for invoking getAlarmHistory.
type GetAlarmHistoryOutputArgs struct {
	AlarmLevels      pulumi.StringArrayInput            `pulumi:"alarmLevels"`
	AlarmObject      pulumi.StringPtrInput              `pulumi:"alarmObject"`
	AlarmStatuses    pulumi.StringArrayInput            `pulumi:"alarmStatuses"`
	Content          pulumi.StringPtrInput              `pulumi:"content"`
	EndTime          pulumi.IntPtrInput                 `pulumi:"endTime"`
	InstanceGroupIds pulumi.IntArrayInput               `pulumi:"instanceGroupIds"`
	MetricNames      pulumi.StringArrayInput            `pulumi:"metricNames"`
	Module           pulumi.StringInput                 `pulumi:"module"`
	MonitorTypes     pulumi.StringArrayInput            `pulumi:"monitorTypes"`
	Namespaces       GetAlarmHistoryNamespaceArrayInput `pulumi:"namespaces"`
	Order            pulumi.StringPtrInput              `pulumi:"order"`
	PolicyIds        pulumi.StringArrayInput            `pulumi:"policyIds"`
	PolicyName       pulumi.StringPtrInput              `pulumi:"policyName"`
	ProjectIds       pulumi.IntArrayInput               `pulumi:"projectIds"`
	ReceiverGroups   pulumi.IntArrayInput               `pulumi:"receiverGroups"`
	ReceiverUids     pulumi.IntArrayInput               `pulumi:"receiverUids"`
	ResultOutputFile pulumi.StringPtrInput              `pulumi:"resultOutputFile"`
	StartTime        pulumi.IntPtrInput                 `pulumi:"startTime"`
}

func (GetAlarmHistoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlarmHistoryArgs)(nil)).Elem()
}

// A collection of values returned by getAlarmHistory.
type GetAlarmHistoryResultOutput struct{ *pulumi.OutputState }

func (GetAlarmHistoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlarmHistoryResult)(nil)).Elem()
}

func (o GetAlarmHistoryResultOutput) ToGetAlarmHistoryResultOutput() GetAlarmHistoryResultOutput {
	return o
}

func (o GetAlarmHistoryResultOutput) ToGetAlarmHistoryResultOutputWithContext(ctx context.Context) GetAlarmHistoryResultOutput {
	return o
}

func (o GetAlarmHistoryResultOutput) AlarmLevels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) []string { return v.AlarmLevels }).(pulumi.StringArrayOutput)
}

func (o GetAlarmHistoryResultOutput) AlarmObject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) *string { return v.AlarmObject }).(pulumi.StringPtrOutput)
}

func (o GetAlarmHistoryResultOutput) AlarmStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) []string { return v.AlarmStatuses }).(pulumi.StringArrayOutput)
}

func (o GetAlarmHistoryResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o GetAlarmHistoryResultOutput) EndTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) *int { return v.EndTime }).(pulumi.IntPtrOutput)
}

func (o GetAlarmHistoryResultOutput) Histories() GetAlarmHistoryHistoryArrayOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) []GetAlarmHistoryHistory { return v.Histories }).(GetAlarmHistoryHistoryArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAlarmHistoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAlarmHistoryResultOutput) InstanceGroupIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) []int { return v.InstanceGroupIds }).(pulumi.IntArrayOutput)
}

func (o GetAlarmHistoryResultOutput) MetricNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) []string { return v.MetricNames }).(pulumi.StringArrayOutput)
}

func (o GetAlarmHistoryResultOutput) Module() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) string { return v.Module }).(pulumi.StringOutput)
}

func (o GetAlarmHistoryResultOutput) MonitorTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) []string { return v.MonitorTypes }).(pulumi.StringArrayOutput)
}

func (o GetAlarmHistoryResultOutput) Namespaces() GetAlarmHistoryNamespaceArrayOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) []GetAlarmHistoryNamespace { return v.Namespaces }).(GetAlarmHistoryNamespaceArrayOutput)
}

func (o GetAlarmHistoryResultOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) *string { return v.Order }).(pulumi.StringPtrOutput)
}

func (o GetAlarmHistoryResultOutput) PolicyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) []string { return v.PolicyIds }).(pulumi.StringArrayOutput)
}

func (o GetAlarmHistoryResultOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

func (o GetAlarmHistoryResultOutput) ProjectIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) []int { return v.ProjectIds }).(pulumi.IntArrayOutput)
}

func (o GetAlarmHistoryResultOutput) ReceiverGroups() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) []int { return v.ReceiverGroups }).(pulumi.IntArrayOutput)
}

func (o GetAlarmHistoryResultOutput) ReceiverUids() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) []int { return v.ReceiverUids }).(pulumi.IntArrayOutput)
}

func (o GetAlarmHistoryResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAlarmHistoryResultOutput) StartTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAlarmHistoryResult) *int { return v.StartTime }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAlarmHistoryResultOutput{})
}
