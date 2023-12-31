// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clickhouse

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSpec(ctx *pulumi.Context, args *GetSpecArgs, opts ...pulumi.InvokeOption) (*GetSpecResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSpecResult
	err := ctx.Invoke("tencentcloud:Clickhouse/getSpec:getSpec", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSpec.
type GetSpecArgs struct {
	IsElastic        *bool   `pulumi:"isElastic"`
	PayMode          *string `pulumi:"payMode"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	Zone             string  `pulumi:"zone"`
}

// A collection of values returned by getSpec.
type GetSpecResult struct {
	AttachCbsSpecs []GetSpecAttachCbsSpec `pulumi:"attachCbsSpecs"`
	CommonSpecs    []GetSpecCommonSpec    `pulumi:"commonSpecs"`
	DataSpecs      []GetSpecDataSpec      `pulumi:"dataSpecs"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	IsElastic        *bool   `pulumi:"isElastic"`
	PayMode          *string `pulumi:"payMode"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	Zone             string  `pulumi:"zone"`
}

func GetSpecOutput(ctx *pulumi.Context, args GetSpecOutputArgs, opts ...pulumi.InvokeOption) GetSpecResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSpecResult, error) {
			args := v.(GetSpecArgs)
			r, err := GetSpec(ctx, &args, opts...)
			var s GetSpecResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSpecResultOutput)
}

// A collection of arguments for invoking getSpec.
type GetSpecOutputArgs struct {
	IsElastic        pulumi.BoolPtrInput   `pulumi:"isElastic"`
	PayMode          pulumi.StringPtrInput `pulumi:"payMode"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	Zone             pulumi.StringInput    `pulumi:"zone"`
}

func (GetSpecOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSpecArgs)(nil)).Elem()
}

// A collection of values returned by getSpec.
type GetSpecResultOutput struct{ *pulumi.OutputState }

func (GetSpecResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSpecResult)(nil)).Elem()
}

func (o GetSpecResultOutput) ToGetSpecResultOutput() GetSpecResultOutput {
	return o
}

func (o GetSpecResultOutput) ToGetSpecResultOutputWithContext(ctx context.Context) GetSpecResultOutput {
	return o
}

func (o GetSpecResultOutput) AttachCbsSpecs() GetSpecAttachCbsSpecArrayOutput {
	return o.ApplyT(func(v GetSpecResult) []GetSpecAttachCbsSpec { return v.AttachCbsSpecs }).(GetSpecAttachCbsSpecArrayOutput)
}

func (o GetSpecResultOutput) CommonSpecs() GetSpecCommonSpecArrayOutput {
	return o.ApplyT(func(v GetSpecResult) []GetSpecCommonSpec { return v.CommonSpecs }).(GetSpecCommonSpecArrayOutput)
}

func (o GetSpecResultOutput) DataSpecs() GetSpecDataSpecArrayOutput {
	return o.ApplyT(func(v GetSpecResult) []GetSpecDataSpec { return v.DataSpecs }).(GetSpecDataSpecArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSpecResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSpecResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSpecResultOutput) IsElastic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSpecResult) *bool { return v.IsElastic }).(pulumi.BoolPtrOutput)
}

func (o GetSpecResultOutput) PayMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSpecResult) *string { return v.PayMode }).(pulumi.StringPtrOutput)
}

func (o GetSpecResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSpecResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetSpecResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetSpecResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSpecResultOutput{})
}
