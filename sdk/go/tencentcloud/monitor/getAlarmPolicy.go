// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAlarmPolicy(ctx *pulumi.Context, args *LookupAlarmPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAlarmPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAlarmPolicyResult
	err := ctx.Invoke("tencentcloud:Monitor/getAlarmPolicy:getAlarmPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlarmPolicy.
type LookupAlarmPolicyArgs struct {
	Dimensions            *string                     `pulumi:"dimensions"`
	Enables               []int                       `pulumi:"enables"`
	Field                 *string                     `pulumi:"field"`
	InstanceGroupId       *int                        `pulumi:"instanceGroupId"`
	Module                string                      `pulumi:"module"`
	MonitorTypes          []string                    `pulumi:"monitorTypes"`
	Namespaces            []string                    `pulumi:"namespaces"`
	NeedCorrespondence    *int                        `pulumi:"needCorrespondence"`
	NotBindAll            *int                        `pulumi:"notBindAll"`
	NotBindingNoticeRule  *int                        `pulumi:"notBindingNoticeRule"`
	NotInstanceGroup      *int                        `pulumi:"notInstanceGroup"`
	NoticeIds             []string                    `pulumi:"noticeIds"`
	OneClickPolicyTypes   []string                    `pulumi:"oneClickPolicyTypes"`
	Order                 *string                     `pulumi:"order"`
	PolicyName            *string                     `pulumi:"policyName"`
	PolicyTypes           []string                    `pulumi:"policyTypes"`
	ProjectIds            []int                       `pulumi:"projectIds"`
	PromInsId             *string                     `pulumi:"promInsId"`
	ReceiverGroups        []int                       `pulumi:"receiverGroups"`
	ReceiverOnCallFormIds []string                    `pulumi:"receiverOnCallFormIds"`
	ReceiverUids          []int                       `pulumi:"receiverUids"`
	ResultOutputFile      *string                     `pulumi:"resultOutputFile"`
	RuleTypes             []string                    `pulumi:"ruleTypes"`
	TriggerTasks          []GetAlarmPolicyTriggerTask `pulumi:"triggerTasks"`
}

// A collection of values returned by getAlarmPolicy.
type LookupAlarmPolicyResult struct {
	Dimensions *string `pulumi:"dimensions"`
	Enables    []int   `pulumi:"enables"`
	Field      *string `pulumi:"field"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string                      `pulumi:"id"`
	InstanceGroupId       *int                        `pulumi:"instanceGroupId"`
	Module                string                      `pulumi:"module"`
	MonitorTypes          []string                    `pulumi:"monitorTypes"`
	Namespaces            []string                    `pulumi:"namespaces"`
	NeedCorrespondence    *int                        `pulumi:"needCorrespondence"`
	NotBindAll            *int                        `pulumi:"notBindAll"`
	NotBindingNoticeRule  *int                        `pulumi:"notBindingNoticeRule"`
	NotInstanceGroup      *int                        `pulumi:"notInstanceGroup"`
	NoticeIds             []string                    `pulumi:"noticeIds"`
	OneClickPolicyTypes   []string                    `pulumi:"oneClickPolicyTypes"`
	Order                 *string                     `pulumi:"order"`
	Policies              []GetAlarmPolicyPolicy      `pulumi:"policies"`
	PolicyName            *string                     `pulumi:"policyName"`
	PolicyTypes           []string                    `pulumi:"policyTypes"`
	ProjectIds            []int                       `pulumi:"projectIds"`
	PromInsId             *string                     `pulumi:"promInsId"`
	ReceiverGroups        []int                       `pulumi:"receiverGroups"`
	ReceiverOnCallFormIds []string                    `pulumi:"receiverOnCallFormIds"`
	ReceiverUids          []int                       `pulumi:"receiverUids"`
	ResultOutputFile      *string                     `pulumi:"resultOutputFile"`
	RuleTypes             []string                    `pulumi:"ruleTypes"`
	TriggerTasks          []GetAlarmPolicyTriggerTask `pulumi:"triggerTasks"`
}

func LookupAlarmPolicyOutput(ctx *pulumi.Context, args LookupAlarmPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupAlarmPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAlarmPolicyResult, error) {
			args := v.(LookupAlarmPolicyArgs)
			r, err := LookupAlarmPolicy(ctx, &args, opts...)
			var s LookupAlarmPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAlarmPolicyResultOutput)
}

// A collection of arguments for invoking getAlarmPolicy.
type LookupAlarmPolicyOutputArgs struct {
	Dimensions            pulumi.StringPtrInput               `pulumi:"dimensions"`
	Enables               pulumi.IntArrayInput                `pulumi:"enables"`
	Field                 pulumi.StringPtrInput               `pulumi:"field"`
	InstanceGroupId       pulumi.IntPtrInput                  `pulumi:"instanceGroupId"`
	Module                pulumi.StringInput                  `pulumi:"module"`
	MonitorTypes          pulumi.StringArrayInput             `pulumi:"monitorTypes"`
	Namespaces            pulumi.StringArrayInput             `pulumi:"namespaces"`
	NeedCorrespondence    pulumi.IntPtrInput                  `pulumi:"needCorrespondence"`
	NotBindAll            pulumi.IntPtrInput                  `pulumi:"notBindAll"`
	NotBindingNoticeRule  pulumi.IntPtrInput                  `pulumi:"notBindingNoticeRule"`
	NotInstanceGroup      pulumi.IntPtrInput                  `pulumi:"notInstanceGroup"`
	NoticeIds             pulumi.StringArrayInput             `pulumi:"noticeIds"`
	OneClickPolicyTypes   pulumi.StringArrayInput             `pulumi:"oneClickPolicyTypes"`
	Order                 pulumi.StringPtrInput               `pulumi:"order"`
	PolicyName            pulumi.StringPtrInput               `pulumi:"policyName"`
	PolicyTypes           pulumi.StringArrayInput             `pulumi:"policyTypes"`
	ProjectIds            pulumi.IntArrayInput                `pulumi:"projectIds"`
	PromInsId             pulumi.StringPtrInput               `pulumi:"promInsId"`
	ReceiverGroups        pulumi.IntArrayInput                `pulumi:"receiverGroups"`
	ReceiverOnCallFormIds pulumi.StringArrayInput             `pulumi:"receiverOnCallFormIds"`
	ReceiverUids          pulumi.IntArrayInput                `pulumi:"receiverUids"`
	ResultOutputFile      pulumi.StringPtrInput               `pulumi:"resultOutputFile"`
	RuleTypes             pulumi.StringArrayInput             `pulumi:"ruleTypes"`
	TriggerTasks          GetAlarmPolicyTriggerTaskArrayInput `pulumi:"triggerTasks"`
}

func (LookupAlarmPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlarmPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getAlarmPolicy.
type LookupAlarmPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupAlarmPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlarmPolicyResult)(nil)).Elem()
}

func (o LookupAlarmPolicyResultOutput) ToLookupAlarmPolicyResultOutput() LookupAlarmPolicyResultOutput {
	return o
}

func (o LookupAlarmPolicyResultOutput) ToLookupAlarmPolicyResultOutputWithContext(ctx context.Context) LookupAlarmPolicyResultOutput {
	return o
}

func (o LookupAlarmPolicyResultOutput) Dimensions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *string { return v.Dimensions }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) Enables() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []int { return v.Enables }).(pulumi.IntArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) Field() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *string { return v.Field }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAlarmPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAlarmPolicyResultOutput) InstanceGroupId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *int { return v.InstanceGroupId }).(pulumi.IntPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) Module() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) string { return v.Module }).(pulumi.StringOutput)
}

func (o LookupAlarmPolicyResultOutput) MonitorTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.MonitorTypes }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) Namespaces() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.Namespaces }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) NeedCorrespondence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *int { return v.NeedCorrespondence }).(pulumi.IntPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) NotBindAll() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *int { return v.NotBindAll }).(pulumi.IntPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) NotBindingNoticeRule() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *int { return v.NotBindingNoticeRule }).(pulumi.IntPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) NotInstanceGroup() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *int { return v.NotInstanceGroup }).(pulumi.IntPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) NoticeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.NoticeIds }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) OneClickPolicyTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.OneClickPolicyTypes }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *string { return v.Order }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) Policies() GetAlarmPolicyPolicyArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []GetAlarmPolicyPolicy { return v.Policies }).(GetAlarmPolicyPolicyArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) PolicyTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.PolicyTypes }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) ProjectIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []int { return v.ProjectIds }).(pulumi.IntArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) PromInsId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *string { return v.PromInsId }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) ReceiverGroups() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []int { return v.ReceiverGroups }).(pulumi.IntArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) ReceiverOnCallFormIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.ReceiverOnCallFormIds }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) ReceiverUids() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []int { return v.ReceiverUids }).(pulumi.IntArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) RuleTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.RuleTypes }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) TriggerTasks() GetAlarmPolicyTriggerTaskArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []GetAlarmPolicyTriggerTask { return v.TriggerTasks }).(GetAlarmPolicyTriggerTaskArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlarmPolicyResultOutput{})
}
