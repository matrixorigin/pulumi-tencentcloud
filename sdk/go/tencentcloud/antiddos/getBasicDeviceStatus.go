// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package antiddos

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBasicDeviceStatus(ctx *pulumi.Context, args *GetBasicDeviceStatusArgs, opts ...pulumi.InvokeOption) (*GetBasicDeviceStatusResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBasicDeviceStatusResult
	err := ctx.Invoke("tencentcloud:Antiddos/getBasicDeviceStatus:getBasicDeviceStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBasicDeviceStatus.
type GetBasicDeviceStatusArgs struct {
	FilterRegion     *int     `pulumi:"filterRegion"`
	IdLists          []string `pulumi:"idLists"`
	IpLists          []string `pulumi:"ipLists"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
}

// A collection of values returned by getBasicDeviceStatus.
type GetBasicDeviceStatusResult struct {
	ClbDatas     []GetBasicDeviceStatusClbData `pulumi:"clbDatas"`
	Datas        []GetBasicDeviceStatusData    `pulumi:"datas"`
	FilterRegion *int                          `pulumi:"filterRegion"`
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	IdLists          []string `pulumi:"idLists"`
	IpLists          []string `pulumi:"ipLists"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
}

func GetBasicDeviceStatusOutput(ctx *pulumi.Context, args GetBasicDeviceStatusOutputArgs, opts ...pulumi.InvokeOption) GetBasicDeviceStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBasicDeviceStatusResult, error) {
			args := v.(GetBasicDeviceStatusArgs)
			r, err := GetBasicDeviceStatus(ctx, &args, opts...)
			var s GetBasicDeviceStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBasicDeviceStatusResultOutput)
}

// A collection of arguments for invoking getBasicDeviceStatus.
type GetBasicDeviceStatusOutputArgs struct {
	FilterRegion     pulumi.IntPtrInput      `pulumi:"filterRegion"`
	IdLists          pulumi.StringArrayInput `pulumi:"idLists"`
	IpLists          pulumi.StringArrayInput `pulumi:"ipLists"`
	ResultOutputFile pulumi.StringPtrInput   `pulumi:"resultOutputFile"`
}

func (GetBasicDeviceStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBasicDeviceStatusArgs)(nil)).Elem()
}

// A collection of values returned by getBasicDeviceStatus.
type GetBasicDeviceStatusResultOutput struct{ *pulumi.OutputState }

func (GetBasicDeviceStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBasicDeviceStatusResult)(nil)).Elem()
}

func (o GetBasicDeviceStatusResultOutput) ToGetBasicDeviceStatusResultOutput() GetBasicDeviceStatusResultOutput {
	return o
}

func (o GetBasicDeviceStatusResultOutput) ToGetBasicDeviceStatusResultOutputWithContext(ctx context.Context) GetBasicDeviceStatusResultOutput {
	return o
}

func (o GetBasicDeviceStatusResultOutput) ClbDatas() GetBasicDeviceStatusClbDataArrayOutput {
	return o.ApplyT(func(v GetBasicDeviceStatusResult) []GetBasicDeviceStatusClbData { return v.ClbDatas }).(GetBasicDeviceStatusClbDataArrayOutput)
}

func (o GetBasicDeviceStatusResultOutput) Datas() GetBasicDeviceStatusDataArrayOutput {
	return o.ApplyT(func(v GetBasicDeviceStatusResult) []GetBasicDeviceStatusData { return v.Datas }).(GetBasicDeviceStatusDataArrayOutput)
}

func (o GetBasicDeviceStatusResultOutput) FilterRegion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetBasicDeviceStatusResult) *int { return v.FilterRegion }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBasicDeviceStatusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBasicDeviceStatusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBasicDeviceStatusResultOutput) IdLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBasicDeviceStatusResult) []string { return v.IdLists }).(pulumi.StringArrayOutput)
}

func (o GetBasicDeviceStatusResultOutput) IpLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBasicDeviceStatusResult) []string { return v.IpLists }).(pulumi.StringArrayOutput)
}

func (o GetBasicDeviceStatusResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBasicDeviceStatusResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBasicDeviceStatusResultOutput{})
}
