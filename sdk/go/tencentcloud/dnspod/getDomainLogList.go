// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dnspod

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDomainLogList(ctx *pulumi.Context, args *GetDomainLogListArgs, opts ...pulumi.InvokeOption) (*GetDomainLogListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDomainLogListResult
	err := ctx.Invoke("tencentcloud:Dnspod/getDomainLogList:getDomainLogList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomainLogList.
type GetDomainLogListArgs struct {
	Domain           string  `pulumi:"domain"`
	DomainId         *int    `pulumi:"domainId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDomainLogList.
type GetDomainLogListResult struct {
	Domain   string `pulumi:"domain"`
	DomainId *int   `pulumi:"domainId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	LogLists         []string `pulumi:"logLists"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
}

func GetDomainLogListOutput(ctx *pulumi.Context, args GetDomainLogListOutputArgs, opts ...pulumi.InvokeOption) GetDomainLogListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainLogListResult, error) {
			args := v.(GetDomainLogListArgs)
			r, err := GetDomainLogList(ctx, &args, opts...)
			var s GetDomainLogListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainLogListResultOutput)
}

// A collection of arguments for invoking getDomainLogList.
type GetDomainLogListOutputArgs struct {
	Domain           pulumi.StringInput    `pulumi:"domain"`
	DomainId         pulumi.IntPtrInput    `pulumi:"domainId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDomainLogListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainLogListArgs)(nil)).Elem()
}

// A collection of values returned by getDomainLogList.
type GetDomainLogListResultOutput struct{ *pulumi.OutputState }

func (GetDomainLogListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainLogListResult)(nil)).Elem()
}

func (o GetDomainLogListResultOutput) ToGetDomainLogListResultOutput() GetDomainLogListResultOutput {
	return o
}

func (o GetDomainLogListResultOutput) ToGetDomainLogListResultOutputWithContext(ctx context.Context) GetDomainLogListResultOutput {
	return o
}

func (o GetDomainLogListResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainLogListResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o GetDomainLogListResultOutput) DomainId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDomainLogListResult) *int { return v.DomainId }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDomainLogListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainLogListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDomainLogListResultOutput) LogLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDomainLogListResult) []string { return v.LogLists }).(pulumi.StringArrayOutput)
}

func (o GetDomainLogListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainLogListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainLogListResultOutput{})
}