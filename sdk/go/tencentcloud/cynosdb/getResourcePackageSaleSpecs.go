// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetResourcePackageSaleSpecs(ctx *pulumi.Context, args *GetResourcePackageSaleSpecsArgs, opts ...pulumi.InvokeOption) (*GetResourcePackageSaleSpecsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetResourcePackageSaleSpecsResult
	err := ctx.Invoke("tencentcloud:Cynosdb/getResourcePackageSaleSpecs:getResourcePackageSaleSpecs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourcePackageSaleSpecs.
type GetResourcePackageSaleSpecsArgs struct {
	InstanceType     string  `pulumi:"instanceType"`
	PackageRegion    string  `pulumi:"packageRegion"`
	PackageType      string  `pulumi:"packageType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getResourcePackageSaleSpecs.
type GetResourcePackageSaleSpecsResult struct {
	Details []GetResourcePackageSaleSpecsDetail `pulumi:"details"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	InstanceType     string  `pulumi:"instanceType"`
	PackageRegion    string  `pulumi:"packageRegion"`
	PackageType      string  `pulumi:"packageType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetResourcePackageSaleSpecsOutput(ctx *pulumi.Context, args GetResourcePackageSaleSpecsOutputArgs, opts ...pulumi.InvokeOption) GetResourcePackageSaleSpecsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetResourcePackageSaleSpecsResult, error) {
			args := v.(GetResourcePackageSaleSpecsArgs)
			r, err := GetResourcePackageSaleSpecs(ctx, &args, opts...)
			var s GetResourcePackageSaleSpecsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetResourcePackageSaleSpecsResultOutput)
}

// A collection of arguments for invoking getResourcePackageSaleSpecs.
type GetResourcePackageSaleSpecsOutputArgs struct {
	InstanceType     pulumi.StringInput    `pulumi:"instanceType"`
	PackageRegion    pulumi.StringInput    `pulumi:"packageRegion"`
	PackageType      pulumi.StringInput    `pulumi:"packageType"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetResourcePackageSaleSpecsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourcePackageSaleSpecsArgs)(nil)).Elem()
}

// A collection of values returned by getResourcePackageSaleSpecs.
type GetResourcePackageSaleSpecsResultOutput struct{ *pulumi.OutputState }

func (GetResourcePackageSaleSpecsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourcePackageSaleSpecsResult)(nil)).Elem()
}

func (o GetResourcePackageSaleSpecsResultOutput) ToGetResourcePackageSaleSpecsResultOutput() GetResourcePackageSaleSpecsResultOutput {
	return o
}

func (o GetResourcePackageSaleSpecsResultOutput) ToGetResourcePackageSaleSpecsResultOutputWithContext(ctx context.Context) GetResourcePackageSaleSpecsResultOutput {
	return o
}

func (o GetResourcePackageSaleSpecsResultOutput) Details() GetResourcePackageSaleSpecsDetailArrayOutput {
	return o.ApplyT(func(v GetResourcePackageSaleSpecsResult) []GetResourcePackageSaleSpecsDetail { return v.Details }).(GetResourcePackageSaleSpecsDetailArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetResourcePackageSaleSpecsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourcePackageSaleSpecsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetResourcePackageSaleSpecsResultOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourcePackageSaleSpecsResult) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o GetResourcePackageSaleSpecsResultOutput) PackageRegion() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourcePackageSaleSpecsResult) string { return v.PackageRegion }).(pulumi.StringOutput)
}

func (o GetResourcePackageSaleSpecsResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourcePackageSaleSpecsResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o GetResourcePackageSaleSpecsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetResourcePackageSaleSpecsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetResourcePackageSaleSpecsResultOutput{})
}
