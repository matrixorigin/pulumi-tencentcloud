// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRenewalPrice(ctx *pulumi.Context, args *GetRenewalPriceArgs, opts ...pulumi.InvokeOption) (*GetRenewalPriceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRenewalPriceResult
	err := ctx.Invoke("tencentcloud:Mariadb/getRenewalPrice:getRenewalPrice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRenewalPrice.
type GetRenewalPriceArgs struct {
	AmountUnit       *string `pulumi:"amountUnit"`
	InstanceId       string  `pulumi:"instanceId"`
	Period           *int    `pulumi:"period"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getRenewalPrice.
type GetRenewalPriceResult struct {
	AmountUnit *string `pulumi:"amountUnit"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	InstanceId       string  `pulumi:"instanceId"`
	OriginalPrice    int     `pulumi:"originalPrice"`
	Period           *int    `pulumi:"period"`
	Price            int     `pulumi:"price"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetRenewalPriceOutput(ctx *pulumi.Context, args GetRenewalPriceOutputArgs, opts ...pulumi.InvokeOption) GetRenewalPriceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRenewalPriceResult, error) {
			args := v.(GetRenewalPriceArgs)
			r, err := GetRenewalPrice(ctx, &args, opts...)
			var s GetRenewalPriceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRenewalPriceResultOutput)
}

// A collection of arguments for invoking getRenewalPrice.
type GetRenewalPriceOutputArgs struct {
	AmountUnit       pulumi.StringPtrInput `pulumi:"amountUnit"`
	InstanceId       pulumi.StringInput    `pulumi:"instanceId"`
	Period           pulumi.IntPtrInput    `pulumi:"period"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetRenewalPriceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRenewalPriceArgs)(nil)).Elem()
}

// A collection of values returned by getRenewalPrice.
type GetRenewalPriceResultOutput struct{ *pulumi.OutputState }

func (GetRenewalPriceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRenewalPriceResult)(nil)).Elem()
}

func (o GetRenewalPriceResultOutput) ToGetRenewalPriceResultOutput() GetRenewalPriceResultOutput {
	return o
}

func (o GetRenewalPriceResultOutput) ToGetRenewalPriceResultOutputWithContext(ctx context.Context) GetRenewalPriceResultOutput {
	return o
}

func (o GetRenewalPriceResultOutput) AmountUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) *string { return v.AmountUnit }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRenewalPriceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRenewalPriceResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetRenewalPriceResultOutput) OriginalPrice() pulumi.IntOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) int { return v.OriginalPrice }).(pulumi.IntOutput)
}

func (o GetRenewalPriceResultOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) *int { return v.Period }).(pulumi.IntPtrOutput)
}

func (o GetRenewalPriceResultOutput) Price() pulumi.IntOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) int { return v.Price }).(pulumi.IntOutput)
}

func (o GetRenewalPriceResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRenewalPriceResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRenewalPriceResultOutput{})
}
