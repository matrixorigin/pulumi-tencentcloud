// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRealServersStatus(ctx *pulumi.Context, args *GetRealServersStatusArgs, opts ...pulumi.InvokeOption) (*GetRealServersStatusResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRealServersStatusResult
	err := ctx.Invoke("tencentcloud:Gaap/getRealServersStatus:getRealServersStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRealServersStatus.
type GetRealServersStatusArgs struct {
	RealServerIds    []string `pulumi:"realServerIds"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
}

// A collection of values returned by getRealServersStatus.
type GetRealServersStatusResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                   string                                    `pulumi:"id"`
	RealServerIds        []string                                  `pulumi:"realServerIds"`
	RealServerStatusSets []GetRealServersStatusRealServerStatusSet `pulumi:"realServerStatusSets"`
	ResultOutputFile     *string                                   `pulumi:"resultOutputFile"`
}

func GetRealServersStatusOutput(ctx *pulumi.Context, args GetRealServersStatusOutputArgs, opts ...pulumi.InvokeOption) GetRealServersStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRealServersStatusResult, error) {
			args := v.(GetRealServersStatusArgs)
			r, err := GetRealServersStatus(ctx, &args, opts...)
			var s GetRealServersStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRealServersStatusResultOutput)
}

// A collection of arguments for invoking getRealServersStatus.
type GetRealServersStatusOutputArgs struct {
	RealServerIds    pulumi.StringArrayInput `pulumi:"realServerIds"`
	ResultOutputFile pulumi.StringPtrInput   `pulumi:"resultOutputFile"`
}

func (GetRealServersStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRealServersStatusArgs)(nil)).Elem()
}

// A collection of values returned by getRealServersStatus.
type GetRealServersStatusResultOutput struct{ *pulumi.OutputState }

func (GetRealServersStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRealServersStatusResult)(nil)).Elem()
}

func (o GetRealServersStatusResultOutput) ToGetRealServersStatusResultOutput() GetRealServersStatusResultOutput {
	return o
}

func (o GetRealServersStatusResultOutput) ToGetRealServersStatusResultOutputWithContext(ctx context.Context) GetRealServersStatusResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetRealServersStatusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRealServersStatusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRealServersStatusResultOutput) RealServerIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRealServersStatusResult) []string { return v.RealServerIds }).(pulumi.StringArrayOutput)
}

func (o GetRealServersStatusResultOutput) RealServerStatusSets() GetRealServersStatusRealServerStatusSetArrayOutput {
	return o.ApplyT(func(v GetRealServersStatusResult) []GetRealServersStatusRealServerStatusSet {
		return v.RealServerStatusSets
	}).(GetRealServersStatusRealServerStatusSetArrayOutput)
}

func (o GetRealServersStatusResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRealServersStatusResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRealServersStatusResultOutput{})
}