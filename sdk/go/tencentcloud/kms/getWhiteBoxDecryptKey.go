// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetWhiteBoxDecryptKey(ctx *pulumi.Context, args *GetWhiteBoxDecryptKeyArgs, opts ...pulumi.InvokeOption) (*GetWhiteBoxDecryptKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetWhiteBoxDecryptKeyResult
	err := ctx.Invoke("tencentcloud:Kms/getWhiteBoxDecryptKey:getWhiteBoxDecryptKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWhiteBoxDecryptKey.
type GetWhiteBoxDecryptKeyArgs struct {
	KeyId            string  `pulumi:"keyId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getWhiteBoxDecryptKey.
type GetWhiteBoxDecryptKeyResult struct {
	DecryptKey string `pulumi:"decryptKey"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	KeyId            string  `pulumi:"keyId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetWhiteBoxDecryptKeyOutput(ctx *pulumi.Context, args GetWhiteBoxDecryptKeyOutputArgs, opts ...pulumi.InvokeOption) GetWhiteBoxDecryptKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWhiteBoxDecryptKeyResult, error) {
			args := v.(GetWhiteBoxDecryptKeyArgs)
			r, err := GetWhiteBoxDecryptKey(ctx, &args, opts...)
			var s GetWhiteBoxDecryptKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWhiteBoxDecryptKeyResultOutput)
}

// A collection of arguments for invoking getWhiteBoxDecryptKey.
type GetWhiteBoxDecryptKeyOutputArgs struct {
	KeyId            pulumi.StringInput    `pulumi:"keyId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetWhiteBoxDecryptKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWhiteBoxDecryptKeyArgs)(nil)).Elem()
}

// A collection of values returned by getWhiteBoxDecryptKey.
type GetWhiteBoxDecryptKeyResultOutput struct{ *pulumi.OutputState }

func (GetWhiteBoxDecryptKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWhiteBoxDecryptKeyResult)(nil)).Elem()
}

func (o GetWhiteBoxDecryptKeyResultOutput) ToGetWhiteBoxDecryptKeyResultOutput() GetWhiteBoxDecryptKeyResultOutput {
	return o
}

func (o GetWhiteBoxDecryptKeyResultOutput) ToGetWhiteBoxDecryptKeyResultOutputWithContext(ctx context.Context) GetWhiteBoxDecryptKeyResultOutput {
	return o
}

func (o GetWhiteBoxDecryptKeyResultOutput) DecryptKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetWhiteBoxDecryptKeyResult) string { return v.DecryptKey }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetWhiteBoxDecryptKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetWhiteBoxDecryptKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetWhiteBoxDecryptKeyResultOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetWhiteBoxDecryptKeyResult) string { return v.KeyId }).(pulumi.StringOutput)
}

func (o GetWhiteBoxDecryptKeyResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWhiteBoxDecryptKeyResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWhiteBoxDecryptKeyResultOutput{})
}