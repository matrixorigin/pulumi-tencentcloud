// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tat

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInvoker(ctx *pulumi.Context, args *LookupInvokerArgs, opts ...pulumi.InvokeOption) (*LookupInvokerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInvokerResult
	err := ctx.Invoke("tencentcloud:Tat/getInvoker:getInvoker", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInvoker.
type LookupInvokerArgs struct {
	CommandId        *string `pulumi:"commandId"`
	InvokerId        *string `pulumi:"invokerId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	Type             *string `pulumi:"type"`
}

// A collection of values returned by getInvoker.
type LookupInvokerResult struct {
	CommandId *string `pulumi:"commandId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                 `pulumi:"id"`
	InvokerId        *string                `pulumi:"invokerId"`
	InvokerSets      []GetInvokerInvokerSet `pulumi:"invokerSets"`
	ResultOutputFile *string                `pulumi:"resultOutputFile"`
	Type             *string                `pulumi:"type"`
}

func LookupInvokerOutput(ctx *pulumi.Context, args LookupInvokerOutputArgs, opts ...pulumi.InvokeOption) LookupInvokerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInvokerResult, error) {
			args := v.(LookupInvokerArgs)
			r, err := LookupInvoker(ctx, &args, opts...)
			var s LookupInvokerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInvokerResultOutput)
}

// A collection of arguments for invoking getInvoker.
type LookupInvokerOutputArgs struct {
	CommandId        pulumi.StringPtrInput `pulumi:"commandId"`
	InvokerId        pulumi.StringPtrInput `pulumi:"invokerId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	Type             pulumi.StringPtrInput `pulumi:"type"`
}

func (LookupInvokerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInvokerArgs)(nil)).Elem()
}

// A collection of values returned by getInvoker.
type LookupInvokerResultOutput struct{ *pulumi.OutputState }

func (LookupInvokerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInvokerResult)(nil)).Elem()
}

func (o LookupInvokerResultOutput) ToLookupInvokerResultOutput() LookupInvokerResultOutput {
	return o
}

func (o LookupInvokerResultOutput) ToLookupInvokerResultOutputWithContext(ctx context.Context) LookupInvokerResultOutput {
	return o
}

func (o LookupInvokerResultOutput) CommandId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInvokerResult) *string { return v.CommandId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInvokerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvokerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInvokerResultOutput) InvokerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInvokerResult) *string { return v.InvokerId }).(pulumi.StringPtrOutput)
}

func (o LookupInvokerResultOutput) InvokerSets() GetInvokerInvokerSetArrayOutput {
	return o.ApplyT(func(v LookupInvokerResult) []GetInvokerInvokerSet { return v.InvokerSets }).(GetInvokerInvokerSetArrayOutput)
}

func (o LookupInvokerResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInvokerResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupInvokerResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInvokerResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInvokerResultOutput{})
}
