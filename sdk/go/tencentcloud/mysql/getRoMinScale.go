// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRoMinScale(ctx *pulumi.Context, args *GetRoMinScaleArgs, opts ...pulumi.InvokeOption) (*GetRoMinScaleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRoMinScaleResult
	err := ctx.Invoke("tencentcloud:Mysql/getRoMinScale:getRoMinScale", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRoMinScale.
type GetRoMinScaleArgs struct {
	MasterInstanceId *string `pulumi:"masterInstanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	RoInstanceId     *string `pulumi:"roInstanceId"`
}

// A collection of values returned by getRoMinScale.
type GetRoMinScaleResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	MasterInstanceId *string `pulumi:"masterInstanceId"`
	Memory           int     `pulumi:"memory"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	RoInstanceId     *string `pulumi:"roInstanceId"`
	Volume           int     `pulumi:"volume"`
}

func GetRoMinScaleOutput(ctx *pulumi.Context, args GetRoMinScaleOutputArgs, opts ...pulumi.InvokeOption) GetRoMinScaleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRoMinScaleResult, error) {
			args := v.(GetRoMinScaleArgs)
			r, err := GetRoMinScale(ctx, &args, opts...)
			var s GetRoMinScaleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRoMinScaleResultOutput)
}

// A collection of arguments for invoking getRoMinScale.
type GetRoMinScaleOutputArgs struct {
	MasterInstanceId pulumi.StringPtrInput `pulumi:"masterInstanceId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	RoInstanceId     pulumi.StringPtrInput `pulumi:"roInstanceId"`
}

func (GetRoMinScaleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRoMinScaleArgs)(nil)).Elem()
}

// A collection of values returned by getRoMinScale.
type GetRoMinScaleResultOutput struct{ *pulumi.OutputState }

func (GetRoMinScaleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRoMinScaleResult)(nil)).Elem()
}

func (o GetRoMinScaleResultOutput) ToGetRoMinScaleResultOutput() GetRoMinScaleResultOutput {
	return o
}

func (o GetRoMinScaleResultOutput) ToGetRoMinScaleResultOutputWithContext(ctx context.Context) GetRoMinScaleResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetRoMinScaleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRoMinScaleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRoMinScaleResultOutput) MasterInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRoMinScaleResult) *string { return v.MasterInstanceId }).(pulumi.StringPtrOutput)
}

func (o GetRoMinScaleResultOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v GetRoMinScaleResult) int { return v.Memory }).(pulumi.IntOutput)
}

func (o GetRoMinScaleResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRoMinScaleResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetRoMinScaleResultOutput) RoInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRoMinScaleResult) *string { return v.RoInstanceId }).(pulumi.StringPtrOutput)
}

func (o GetRoMinScaleResultOutput) Volume() pulumi.IntOutput {
	return o.ApplyT(func(v GetRoMinScaleResult) int { return v.Volume }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRoMinScaleResultOutput{})
}
