// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAccessRegionsByDestRegion(ctx *pulumi.Context, args *GetAccessRegionsByDestRegionArgs, opts ...pulumi.InvokeOption) (*GetAccessRegionsByDestRegionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccessRegionsByDestRegionResult
	err := ctx.Invoke("tencentcloud:Gaap/getAccessRegionsByDestRegion:getAccessRegionsByDestRegion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessRegionsByDestRegion.
type GetAccessRegionsByDestRegionArgs struct {
	DestRegion       string  `pulumi:"destRegion"`
	IpAddressVersion *string `pulumi:"ipAddressVersion"`
	PackageType      *string `pulumi:"packageType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getAccessRegionsByDestRegion.
type GetAccessRegionsByDestRegionResult struct {
	AccessRegionSets []GetAccessRegionsByDestRegionAccessRegionSet `pulumi:"accessRegionSets"`
	DestRegion       string                                        `pulumi:"destRegion"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	IpAddressVersion *string `pulumi:"ipAddressVersion"`
	PackageType      *string `pulumi:"packageType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetAccessRegionsByDestRegionOutput(ctx *pulumi.Context, args GetAccessRegionsByDestRegionOutputArgs, opts ...pulumi.InvokeOption) GetAccessRegionsByDestRegionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccessRegionsByDestRegionResult, error) {
			args := v.(GetAccessRegionsByDestRegionArgs)
			r, err := GetAccessRegionsByDestRegion(ctx, &args, opts...)
			var s GetAccessRegionsByDestRegionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccessRegionsByDestRegionResultOutput)
}

// A collection of arguments for invoking getAccessRegionsByDestRegion.
type GetAccessRegionsByDestRegionOutputArgs struct {
	DestRegion       pulumi.StringInput    `pulumi:"destRegion"`
	IpAddressVersion pulumi.StringPtrInput `pulumi:"ipAddressVersion"`
	PackageType      pulumi.StringPtrInput `pulumi:"packageType"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetAccessRegionsByDestRegionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessRegionsByDestRegionArgs)(nil)).Elem()
}

// A collection of values returned by getAccessRegionsByDestRegion.
type GetAccessRegionsByDestRegionResultOutput struct{ *pulumi.OutputState }

func (GetAccessRegionsByDestRegionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessRegionsByDestRegionResult)(nil)).Elem()
}

func (o GetAccessRegionsByDestRegionResultOutput) ToGetAccessRegionsByDestRegionResultOutput() GetAccessRegionsByDestRegionResultOutput {
	return o
}

func (o GetAccessRegionsByDestRegionResultOutput) ToGetAccessRegionsByDestRegionResultOutputWithContext(ctx context.Context) GetAccessRegionsByDestRegionResultOutput {
	return o
}

func (o GetAccessRegionsByDestRegionResultOutput) AccessRegionSets() GetAccessRegionsByDestRegionAccessRegionSetArrayOutput {
	return o.ApplyT(func(v GetAccessRegionsByDestRegionResult) []GetAccessRegionsByDestRegionAccessRegionSet {
		return v.AccessRegionSets
	}).(GetAccessRegionsByDestRegionAccessRegionSetArrayOutput)
}

func (o GetAccessRegionsByDestRegionResultOutput) DestRegion() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessRegionsByDestRegionResult) string { return v.DestRegion }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccessRegionsByDestRegionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessRegionsByDestRegionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAccessRegionsByDestRegionResultOutput) IpAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessRegionsByDestRegionResult) *string { return v.IpAddressVersion }).(pulumi.StringPtrOutput)
}

func (o GetAccessRegionsByDestRegionResultOutput) PackageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessRegionsByDestRegionResult) *string { return v.PackageType }).(pulumi.StringPtrOutput)
}

func (o GetAccessRegionsByDestRegionResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessRegionsByDestRegionResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccessRegionsByDestRegionResultOutput{})
}
