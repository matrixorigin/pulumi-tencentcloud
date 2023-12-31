// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDomainVerifier(ctx *pulumi.Context, args *GetDomainVerifierArgs, opts ...pulumi.InvokeOption) (*GetDomainVerifierResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDomainVerifierResult
	err := ctx.Invoke("tencentcloud:Cdn/getDomainVerifier:getDomainVerifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomainVerifier.
type GetDomainVerifierArgs struct {
	AutoVerify       *bool   `pulumi:"autoVerify"`
	Domain           string  `pulumi:"domain"`
	FailedReason     *string `pulumi:"failedReason"`
	FreezeRecord     *bool   `pulumi:"freezeRecord"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	VerifyType       *string `pulumi:"verifyType"`
}

// A collection of values returned by getDomainVerifier.
type GetDomainVerifierResult struct {
	AutoVerify        *bool    `pulumi:"autoVerify"`
	Domain            string   `pulumi:"domain"`
	FailedReason      *string  `pulumi:"failedReason"`
	FileVerifyDomains []string `pulumi:"fileVerifyDomains"`
	FileVerifyName    string   `pulumi:"fileVerifyName"`
	FileVerifyUrl     string   `pulumi:"fileVerifyUrl"`
	FreezeRecord      *bool    `pulumi:"freezeRecord"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	Record           string  `pulumi:"record"`
	RecordType       string  `pulumi:"recordType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SubDomain        string  `pulumi:"subDomain"`
	VerifyResult     bool    `pulumi:"verifyResult"`
	VerifyType       *string `pulumi:"verifyType"`
}

func GetDomainVerifierOutput(ctx *pulumi.Context, args GetDomainVerifierOutputArgs, opts ...pulumi.InvokeOption) GetDomainVerifierResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainVerifierResult, error) {
			args := v.(GetDomainVerifierArgs)
			r, err := GetDomainVerifier(ctx, &args, opts...)
			var s GetDomainVerifierResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainVerifierResultOutput)
}

// A collection of arguments for invoking getDomainVerifier.
type GetDomainVerifierOutputArgs struct {
	AutoVerify       pulumi.BoolPtrInput   `pulumi:"autoVerify"`
	Domain           pulumi.StringInput    `pulumi:"domain"`
	FailedReason     pulumi.StringPtrInput `pulumi:"failedReason"`
	FreezeRecord     pulumi.BoolPtrInput   `pulumi:"freezeRecord"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	VerifyType       pulumi.StringPtrInput `pulumi:"verifyType"`
}

func (GetDomainVerifierOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainVerifierArgs)(nil)).Elem()
}

// A collection of values returned by getDomainVerifier.
type GetDomainVerifierResultOutput struct{ *pulumi.OutputState }

func (GetDomainVerifierResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainVerifierResult)(nil)).Elem()
}

func (o GetDomainVerifierResultOutput) ToGetDomainVerifierResultOutput() GetDomainVerifierResultOutput {
	return o
}

func (o GetDomainVerifierResultOutput) ToGetDomainVerifierResultOutputWithContext(ctx context.Context) GetDomainVerifierResultOutput {
	return o
}

func (o GetDomainVerifierResultOutput) AutoVerify() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) *bool { return v.AutoVerify }).(pulumi.BoolPtrOutput)
}

func (o GetDomainVerifierResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o GetDomainVerifierResultOutput) FailedReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) *string { return v.FailedReason }).(pulumi.StringPtrOutput)
}

func (o GetDomainVerifierResultOutput) FileVerifyDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) []string { return v.FileVerifyDomains }).(pulumi.StringArrayOutput)
}

func (o GetDomainVerifierResultOutput) FileVerifyName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) string { return v.FileVerifyName }).(pulumi.StringOutput)
}

func (o GetDomainVerifierResultOutput) FileVerifyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) string { return v.FileVerifyUrl }).(pulumi.StringOutput)
}

func (o GetDomainVerifierResultOutput) FreezeRecord() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) *bool { return v.FreezeRecord }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDomainVerifierResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDomainVerifierResultOutput) Record() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) string { return v.Record }).(pulumi.StringOutput)
}

func (o GetDomainVerifierResultOutput) RecordType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) string { return v.RecordType }).(pulumi.StringOutput)
}

func (o GetDomainVerifierResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetDomainVerifierResultOutput) SubDomain() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) string { return v.SubDomain }).(pulumi.StringOutput)
}

func (o GetDomainVerifierResultOutput) VerifyResult() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) bool { return v.VerifyResult }).(pulumi.BoolOutput)
}

func (o GetDomainVerifierResultOutput) VerifyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainVerifierResult) *string { return v.VerifyType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainVerifierResultOutput{})
}
