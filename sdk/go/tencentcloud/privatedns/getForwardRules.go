// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetForwardRules(ctx *pulumi.Context, args *GetForwardRulesArgs, opts ...pulumi.InvokeOption) (*GetForwardRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetForwardRulesResult
	err := ctx.Invoke("tencentcloud:PrivateDns/getForwardRules:getForwardRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getForwardRules.
type GetForwardRulesArgs struct {
	Filters          []GetForwardRulesFilter `pulumi:"filters"`
	ResultOutputFile *string                 `pulumi:"resultOutputFile"`
}

// A collection of values returned by getForwardRules.
type GetForwardRulesResult struct {
	Filters         []GetForwardRulesFilter         `pulumi:"filters"`
	ForwardRuleSets []GetForwardRulesForwardRuleSet `pulumi:"forwardRuleSets"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetForwardRulesOutput(ctx *pulumi.Context, args GetForwardRulesOutputArgs, opts ...pulumi.InvokeOption) GetForwardRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetForwardRulesResult, error) {
			args := v.(GetForwardRulesArgs)
			r, err := GetForwardRules(ctx, &args, opts...)
			var s GetForwardRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetForwardRulesResultOutput)
}

// A collection of arguments for invoking getForwardRules.
type GetForwardRulesOutputArgs struct {
	Filters          GetForwardRulesFilterArrayInput `pulumi:"filters"`
	ResultOutputFile pulumi.StringPtrInput           `pulumi:"resultOutputFile"`
}

func (GetForwardRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetForwardRulesArgs)(nil)).Elem()
}

// A collection of values returned by getForwardRules.
type GetForwardRulesResultOutput struct{ *pulumi.OutputState }

func (GetForwardRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetForwardRulesResult)(nil)).Elem()
}

func (o GetForwardRulesResultOutput) ToGetForwardRulesResultOutput() GetForwardRulesResultOutput {
	return o
}

func (o GetForwardRulesResultOutput) ToGetForwardRulesResultOutputWithContext(ctx context.Context) GetForwardRulesResultOutput {
	return o
}

func (o GetForwardRulesResultOutput) Filters() GetForwardRulesFilterArrayOutput {
	return o.ApplyT(func(v GetForwardRulesResult) []GetForwardRulesFilter { return v.Filters }).(GetForwardRulesFilterArrayOutput)
}

func (o GetForwardRulesResultOutput) ForwardRuleSets() GetForwardRulesForwardRuleSetArrayOutput {
	return o.ApplyT(func(v GetForwardRulesResult) []GetForwardRulesForwardRuleSet { return v.ForwardRuleSets }).(GetForwardRulesForwardRuleSetArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetForwardRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetForwardRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetForwardRulesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetForwardRulesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetForwardRulesResultOutput{})
}