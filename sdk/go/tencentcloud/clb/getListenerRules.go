// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetListenerRules(ctx *pulumi.Context, args *GetListenerRulesArgs, opts ...pulumi.InvokeOption) (*GetListenerRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetListenerRulesResult
	err := ctx.Invoke("tencentcloud:Clb/getListenerRules:getListenerRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getListenerRules.
type GetListenerRulesArgs struct {
	ClbId            string  `pulumi:"clbId"`
	Domain           *string `pulumi:"domain"`
	ListenerId       string  `pulumi:"listenerId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	RuleId           *string `pulumi:"ruleId"`
	Scheduler        *string `pulumi:"scheduler"`
	Url              *string `pulumi:"url"`
}

// A collection of values returned by getListenerRules.
type GetListenerRulesResult struct {
	ClbId  string  `pulumi:"clbId"`
	Domain *string `pulumi:"domain"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                     `pulumi:"id"`
	ListenerId       string                     `pulumi:"listenerId"`
	ResultOutputFile *string                    `pulumi:"resultOutputFile"`
	RuleId           *string                    `pulumi:"ruleId"`
	RuleLists        []GetListenerRulesRuleList `pulumi:"ruleLists"`
	Scheduler        *string                    `pulumi:"scheduler"`
	Url              *string                    `pulumi:"url"`
}

func GetListenerRulesOutput(ctx *pulumi.Context, args GetListenerRulesOutputArgs, opts ...pulumi.InvokeOption) GetListenerRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetListenerRulesResult, error) {
			args := v.(GetListenerRulesArgs)
			r, err := GetListenerRules(ctx, &args, opts...)
			var s GetListenerRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetListenerRulesResultOutput)
}

// A collection of arguments for invoking getListenerRules.
type GetListenerRulesOutputArgs struct {
	ClbId            pulumi.StringInput    `pulumi:"clbId"`
	Domain           pulumi.StringPtrInput `pulumi:"domain"`
	ListenerId       pulumi.StringInput    `pulumi:"listenerId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	RuleId           pulumi.StringPtrInput `pulumi:"ruleId"`
	Scheduler        pulumi.StringPtrInput `pulumi:"scheduler"`
	Url              pulumi.StringPtrInput `pulumi:"url"`
}

func (GetListenerRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListenerRulesArgs)(nil)).Elem()
}

// A collection of values returned by getListenerRules.
type GetListenerRulesResultOutput struct{ *pulumi.OutputState }

func (GetListenerRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListenerRulesResult)(nil)).Elem()
}

func (o GetListenerRulesResultOutput) ToGetListenerRulesResultOutput() GetListenerRulesResultOutput {
	return o
}

func (o GetListenerRulesResultOutput) ToGetListenerRulesResultOutputWithContext(ctx context.Context) GetListenerRulesResultOutput {
	return o
}

func (o GetListenerRulesResultOutput) ClbId() pulumi.StringOutput {
	return o.ApplyT(func(v GetListenerRulesResult) string { return v.ClbId }).(pulumi.StringOutput)
}

func (o GetListenerRulesResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListenerRulesResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetListenerRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetListenerRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetListenerRulesResultOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetListenerRulesResult) string { return v.ListenerId }).(pulumi.StringOutput)
}

func (o GetListenerRulesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListenerRulesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetListenerRulesResultOutput) RuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListenerRulesResult) *string { return v.RuleId }).(pulumi.StringPtrOutput)
}

func (o GetListenerRulesResultOutput) RuleLists() GetListenerRulesRuleListArrayOutput {
	return o.ApplyT(func(v GetListenerRulesResult) []GetListenerRulesRuleList { return v.RuleLists }).(GetListenerRulesRuleListArrayOutput)
}

func (o GetListenerRulesResultOutput) Scheduler() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListenerRulesResult) *string { return v.Scheduler }).(pulumi.StringPtrOutput)
}

func (o GetListenerRulesResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListenerRulesResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetListenerRulesResultOutput{})
}
