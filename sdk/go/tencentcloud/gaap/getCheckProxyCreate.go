// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCheckProxyCreate(ctx *pulumi.Context, args *GetCheckProxyCreateArgs, opts ...pulumi.InvokeOption) (*GetCheckProxyCreateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCheckProxyCreateResult
	err := ctx.Invoke("tencentcloud:Gaap/getCheckProxyCreate:getCheckProxyCreate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCheckProxyCreate.
type GetCheckProxyCreateArgs struct {
	AccessRegion     string  `pulumi:"accessRegion"`
	Bandwidth        int     `pulumi:"bandwidth"`
	Concurrent       int     `pulumi:"concurrent"`
	GroupId          *string `pulumi:"groupId"`
	IpAddressVersion *string `pulumi:"ipAddressVersion"`
	NetworkType      *string `pulumi:"networkType"`
	PackageType      *string `pulumi:"packageType"`
	RealServerRegion string  `pulumi:"realServerRegion"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getCheckProxyCreate.
type GetCheckProxyCreateResult struct {
	AccessRegion string  `pulumi:"accessRegion"`
	Bandwidth    int     `pulumi:"bandwidth"`
	CheckFlag    int     `pulumi:"checkFlag"`
	Concurrent   int     `pulumi:"concurrent"`
	GroupId      *string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	IpAddressVersion *string `pulumi:"ipAddressVersion"`
	NetworkType      *string `pulumi:"networkType"`
	PackageType      *string `pulumi:"packageType"`
	RealServerRegion string  `pulumi:"realServerRegion"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetCheckProxyCreateOutput(ctx *pulumi.Context, args GetCheckProxyCreateOutputArgs, opts ...pulumi.InvokeOption) GetCheckProxyCreateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCheckProxyCreateResult, error) {
			args := v.(GetCheckProxyCreateArgs)
			r, err := GetCheckProxyCreate(ctx, &args, opts...)
			var s GetCheckProxyCreateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCheckProxyCreateResultOutput)
}

// A collection of arguments for invoking getCheckProxyCreate.
type GetCheckProxyCreateOutputArgs struct {
	AccessRegion     pulumi.StringInput    `pulumi:"accessRegion"`
	Bandwidth        pulumi.IntInput       `pulumi:"bandwidth"`
	Concurrent       pulumi.IntInput       `pulumi:"concurrent"`
	GroupId          pulumi.StringPtrInput `pulumi:"groupId"`
	IpAddressVersion pulumi.StringPtrInput `pulumi:"ipAddressVersion"`
	NetworkType      pulumi.StringPtrInput `pulumi:"networkType"`
	PackageType      pulumi.StringPtrInput `pulumi:"packageType"`
	RealServerRegion pulumi.StringInput    `pulumi:"realServerRegion"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetCheckProxyCreateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCheckProxyCreateArgs)(nil)).Elem()
}

// A collection of values returned by getCheckProxyCreate.
type GetCheckProxyCreateResultOutput struct{ *pulumi.OutputState }

func (GetCheckProxyCreateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCheckProxyCreateResult)(nil)).Elem()
}

func (o GetCheckProxyCreateResultOutput) ToGetCheckProxyCreateResultOutput() GetCheckProxyCreateResultOutput {
	return o
}

func (o GetCheckProxyCreateResultOutput) ToGetCheckProxyCreateResultOutputWithContext(ctx context.Context) GetCheckProxyCreateResultOutput {
	return o
}

func (o GetCheckProxyCreateResultOutput) AccessRegion() pulumi.StringOutput {
	return o.ApplyT(func(v GetCheckProxyCreateResult) string { return v.AccessRegion }).(pulumi.StringOutput)
}

func (o GetCheckProxyCreateResultOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetCheckProxyCreateResult) int { return v.Bandwidth }).(pulumi.IntOutput)
}

func (o GetCheckProxyCreateResultOutput) CheckFlag() pulumi.IntOutput {
	return o.ApplyT(func(v GetCheckProxyCreateResult) int { return v.CheckFlag }).(pulumi.IntOutput)
}

func (o GetCheckProxyCreateResultOutput) Concurrent() pulumi.IntOutput {
	return o.ApplyT(func(v GetCheckProxyCreateResult) int { return v.Concurrent }).(pulumi.IntOutput)
}

func (o GetCheckProxyCreateResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCheckProxyCreateResult) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCheckProxyCreateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCheckProxyCreateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCheckProxyCreateResultOutput) IpAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCheckProxyCreateResult) *string { return v.IpAddressVersion }).(pulumi.StringPtrOutput)
}

func (o GetCheckProxyCreateResultOutput) NetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCheckProxyCreateResult) *string { return v.NetworkType }).(pulumi.StringPtrOutput)
}

func (o GetCheckProxyCreateResultOutput) PackageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCheckProxyCreateResult) *string { return v.PackageType }).(pulumi.StringPtrOutput)
}

func (o GetCheckProxyCreateResultOutput) RealServerRegion() pulumi.StringOutput {
	return o.ApplyT(func(v GetCheckProxyCreateResult) string { return v.RealServerRegion }).(pulumi.StringOutput)
}

func (o GetCheckProxyCreateResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCheckProxyCreateResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCheckProxyCreateResultOutput{})
}
