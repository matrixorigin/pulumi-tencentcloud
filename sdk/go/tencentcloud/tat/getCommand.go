// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tat

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCommand(ctx *pulumi.Context, args *LookupCommandArgs, opts ...pulumi.InvokeOption) (*LookupCommandResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCommandResult
	err := ctx.Invoke("tencentcloud:Tat/getCommand:getCommand", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCommand.
type LookupCommandArgs struct {
	CommandId        *string `pulumi:"commandId"`
	CommandName      *string `pulumi:"commandName"`
	CommandType      *string `pulumi:"commandType"`
	CreatedBy        *string `pulumi:"createdBy"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getCommand.
type LookupCommandResult struct {
	CommandId   *string                `pulumi:"commandId"`
	CommandName *string                `pulumi:"commandName"`
	CommandSets []GetCommandCommandSet `pulumi:"commandSets"`
	CommandType *string                `pulumi:"commandType"`
	CreatedBy   *string                `pulumi:"createdBy"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func LookupCommandOutput(ctx *pulumi.Context, args LookupCommandOutputArgs, opts ...pulumi.InvokeOption) LookupCommandResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCommandResult, error) {
			args := v.(LookupCommandArgs)
			r, err := LookupCommand(ctx, &args, opts...)
			var s LookupCommandResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCommandResultOutput)
}

// A collection of arguments for invoking getCommand.
type LookupCommandOutputArgs struct {
	CommandId        pulumi.StringPtrInput `pulumi:"commandId"`
	CommandName      pulumi.StringPtrInput `pulumi:"commandName"`
	CommandType      pulumi.StringPtrInput `pulumi:"commandType"`
	CreatedBy        pulumi.StringPtrInput `pulumi:"createdBy"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (LookupCommandOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommandArgs)(nil)).Elem()
}

// A collection of values returned by getCommand.
type LookupCommandResultOutput struct{ *pulumi.OutputState }

func (LookupCommandResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommandResult)(nil)).Elem()
}

func (o LookupCommandResultOutput) ToLookupCommandResultOutput() LookupCommandResultOutput {
	return o
}

func (o LookupCommandResultOutput) ToLookupCommandResultOutputWithContext(ctx context.Context) LookupCommandResultOutput {
	return o
}

func (o LookupCommandResultOutput) CommandId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCommandResult) *string { return v.CommandId }).(pulumi.StringPtrOutput)
}

func (o LookupCommandResultOutput) CommandName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCommandResult) *string { return v.CommandName }).(pulumi.StringPtrOutput)
}

func (o LookupCommandResultOutput) CommandSets() GetCommandCommandSetArrayOutput {
	return o.ApplyT(func(v LookupCommandResult) []GetCommandCommandSet { return v.CommandSets }).(GetCommandCommandSetArrayOutput)
}

func (o LookupCommandResultOutput) CommandType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCommandResult) *string { return v.CommandType }).(pulumi.StringPtrOutput)
}

func (o LookupCommandResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCommandResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCommandResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommandResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCommandResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCommandResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCommandResultOutput{})
}
