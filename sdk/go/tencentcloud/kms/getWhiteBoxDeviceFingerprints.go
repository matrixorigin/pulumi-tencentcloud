// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetWhiteBoxDeviceFingerprints(ctx *pulumi.Context, args *GetWhiteBoxDeviceFingerprintsArgs, opts ...pulumi.InvokeOption) (*GetWhiteBoxDeviceFingerprintsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetWhiteBoxDeviceFingerprintsResult
	err := ctx.Invoke("tencentcloud:Kms/getWhiteBoxDeviceFingerprints:getWhiteBoxDeviceFingerprints", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWhiteBoxDeviceFingerprints.
type GetWhiteBoxDeviceFingerprintsArgs struct {
	KeyId            string  `pulumi:"keyId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getWhiteBoxDeviceFingerprints.
type GetWhiteBoxDeviceFingerprintsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                              `pulumi:"id"`
	KeyId            string                              `pulumi:"keyId"`
	Lists            []GetWhiteBoxDeviceFingerprintsList `pulumi:"lists"`
	ResultOutputFile *string                             `pulumi:"resultOutputFile"`
}

func GetWhiteBoxDeviceFingerprintsOutput(ctx *pulumi.Context, args GetWhiteBoxDeviceFingerprintsOutputArgs, opts ...pulumi.InvokeOption) GetWhiteBoxDeviceFingerprintsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWhiteBoxDeviceFingerprintsResult, error) {
			args := v.(GetWhiteBoxDeviceFingerprintsArgs)
			r, err := GetWhiteBoxDeviceFingerprints(ctx, &args, opts...)
			var s GetWhiteBoxDeviceFingerprintsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWhiteBoxDeviceFingerprintsResultOutput)
}

// A collection of arguments for invoking getWhiteBoxDeviceFingerprints.
type GetWhiteBoxDeviceFingerprintsOutputArgs struct {
	KeyId            pulumi.StringInput    `pulumi:"keyId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetWhiteBoxDeviceFingerprintsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWhiteBoxDeviceFingerprintsArgs)(nil)).Elem()
}

// A collection of values returned by getWhiteBoxDeviceFingerprints.
type GetWhiteBoxDeviceFingerprintsResultOutput struct{ *pulumi.OutputState }

func (GetWhiteBoxDeviceFingerprintsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWhiteBoxDeviceFingerprintsResult)(nil)).Elem()
}

func (o GetWhiteBoxDeviceFingerprintsResultOutput) ToGetWhiteBoxDeviceFingerprintsResultOutput() GetWhiteBoxDeviceFingerprintsResultOutput {
	return o
}

func (o GetWhiteBoxDeviceFingerprintsResultOutput) ToGetWhiteBoxDeviceFingerprintsResultOutputWithContext(ctx context.Context) GetWhiteBoxDeviceFingerprintsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetWhiteBoxDeviceFingerprintsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetWhiteBoxDeviceFingerprintsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetWhiteBoxDeviceFingerprintsResultOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetWhiteBoxDeviceFingerprintsResult) string { return v.KeyId }).(pulumi.StringOutput)
}

func (o GetWhiteBoxDeviceFingerprintsResultOutput) Lists() GetWhiteBoxDeviceFingerprintsListArrayOutput {
	return o.ApplyT(func(v GetWhiteBoxDeviceFingerprintsResult) []GetWhiteBoxDeviceFingerprintsList { return v.Lists }).(GetWhiteBoxDeviceFingerprintsListArrayOutput)
}

func (o GetWhiteBoxDeviceFingerprintsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWhiteBoxDeviceFingerprintsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWhiteBoxDeviceFingerprintsResultOutput{})
}