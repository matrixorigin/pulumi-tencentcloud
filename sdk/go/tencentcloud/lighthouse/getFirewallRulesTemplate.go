// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lighthouse

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetFirewallRulesTemplate(ctx *pulumi.Context, args *GetFirewallRulesTemplateArgs, opts ...pulumi.InvokeOption) (*GetFirewallRulesTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFirewallRulesTemplateResult
	err := ctx.Invoke("tencentcloud:Lighthouse/getFirewallRulesTemplate:getFirewallRulesTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallRulesTemplate.
type GetFirewallRulesTemplateArgs struct {
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getFirewallRulesTemplate.
type GetFirewallRulesTemplateResult struct {
	FirewallRuleSets []GetFirewallRulesTemplateFirewallRuleSet `pulumi:"firewallRuleSets"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetFirewallRulesTemplateOutput(ctx *pulumi.Context, args GetFirewallRulesTemplateOutputArgs, opts ...pulumi.InvokeOption) GetFirewallRulesTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallRulesTemplateResult, error) {
			args := v.(GetFirewallRulesTemplateArgs)
			r, err := GetFirewallRulesTemplate(ctx, &args, opts...)
			var s GetFirewallRulesTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallRulesTemplateResultOutput)
}

// A collection of arguments for invoking getFirewallRulesTemplate.
type GetFirewallRulesTemplateOutputArgs struct {
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetFirewallRulesTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallRulesTemplateArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallRulesTemplate.
type GetFirewallRulesTemplateResultOutput struct{ *pulumi.OutputState }

func (GetFirewallRulesTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallRulesTemplateResult)(nil)).Elem()
}

func (o GetFirewallRulesTemplateResultOutput) ToGetFirewallRulesTemplateResultOutput() GetFirewallRulesTemplateResultOutput {
	return o
}

func (o GetFirewallRulesTemplateResultOutput) ToGetFirewallRulesTemplateResultOutputWithContext(ctx context.Context) GetFirewallRulesTemplateResultOutput {
	return o
}

func (o GetFirewallRulesTemplateResultOutput) FirewallRuleSets() GetFirewallRulesTemplateFirewallRuleSetArrayOutput {
	return o.ApplyT(func(v GetFirewallRulesTemplateResult) []GetFirewallRulesTemplateFirewallRuleSet {
		return v.FirewallRuleSets
	}).(GetFirewallRulesTemplateFirewallRuleSetArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallRulesTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallRulesTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFirewallRulesTemplateResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallRulesTemplateResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallRulesTemplateResultOutput{})
}