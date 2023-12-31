// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetListKeys(ctx *pulumi.Context, args *GetListKeysArgs, opts ...pulumi.InvokeOption) (*GetListKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetListKeysResult
	err := ctx.Invoke("tencentcloud:Kms/getListKeys:getListKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getListKeys.
type GetListKeysArgs struct {
	HsmClusterId     *string `pulumi:"hsmClusterId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	Role             *int    `pulumi:"role"`
}

// A collection of values returned by getListKeys.
type GetListKeysResult struct {
	HsmClusterId *string `pulumi:"hsmClusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string           `pulumi:"id"`
	Keys             []GetListKeysKey `pulumi:"keys"`
	ResultOutputFile *string          `pulumi:"resultOutputFile"`
	Role             *int             `pulumi:"role"`
}

func GetListKeysOutput(ctx *pulumi.Context, args GetListKeysOutputArgs, opts ...pulumi.InvokeOption) GetListKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetListKeysResult, error) {
			args := v.(GetListKeysArgs)
			r, err := GetListKeys(ctx, &args, opts...)
			var s GetListKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetListKeysResultOutput)
}

// A collection of arguments for invoking getListKeys.
type GetListKeysOutputArgs struct {
	HsmClusterId     pulumi.StringPtrInput `pulumi:"hsmClusterId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	Role             pulumi.IntPtrInput    `pulumi:"role"`
}

func (GetListKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListKeysArgs)(nil)).Elem()
}

// A collection of values returned by getListKeys.
type GetListKeysResultOutput struct{ *pulumi.OutputState }

func (GetListKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListKeysResult)(nil)).Elem()
}

func (o GetListKeysResultOutput) ToGetListKeysResultOutput() GetListKeysResultOutput {
	return o
}

func (o GetListKeysResultOutput) ToGetListKeysResultOutputWithContext(ctx context.Context) GetListKeysResultOutput {
	return o
}

func (o GetListKeysResultOutput) HsmClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListKeysResult) *string { return v.HsmClusterId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetListKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetListKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetListKeysResultOutput) Keys() GetListKeysKeyArrayOutput {
	return o.ApplyT(func(v GetListKeysResult) []GetListKeysKey { return v.Keys }).(GetListKeysKeyArrayOutput)
}

func (o GetListKeysResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListKeysResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetListKeysResultOutput) Role() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetListKeysResult) *int { return v.Role }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetListKeysResultOutput{})
}
