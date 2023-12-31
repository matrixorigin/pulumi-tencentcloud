// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetUserClbRegions(ctx *pulumi.Context, args *GetUserClbRegionsArgs, opts ...pulumi.InvokeOption) (*GetUserClbRegionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUserClbRegionsResult
	err := ctx.Invoke("tencentcloud:Waf/getUserClbRegions:getUserClbRegions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserClbRegions.
type GetUserClbRegionsArgs struct {
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getUserClbRegions.
type GetUserClbRegionsResult struct {
	Datas []string `pulumi:"datas"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                      `pulumi:"id"`
	ResultOutputFile *string                     `pulumi:"resultOutputFile"`
	RichDatas        []GetUserClbRegionsRichData `pulumi:"richDatas"`
}

func GetUserClbRegionsOutput(ctx *pulumi.Context, args GetUserClbRegionsOutputArgs, opts ...pulumi.InvokeOption) GetUserClbRegionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserClbRegionsResult, error) {
			args := v.(GetUserClbRegionsArgs)
			r, err := GetUserClbRegions(ctx, &args, opts...)
			var s GetUserClbRegionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserClbRegionsResultOutput)
}

// A collection of arguments for invoking getUserClbRegions.
type GetUserClbRegionsOutputArgs struct {
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetUserClbRegionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserClbRegionsArgs)(nil)).Elem()
}

// A collection of values returned by getUserClbRegions.
type GetUserClbRegionsResultOutput struct{ *pulumi.OutputState }

func (GetUserClbRegionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserClbRegionsResult)(nil)).Elem()
}

func (o GetUserClbRegionsResultOutput) ToGetUserClbRegionsResultOutput() GetUserClbRegionsResultOutput {
	return o
}

func (o GetUserClbRegionsResultOutput) ToGetUserClbRegionsResultOutputWithContext(ctx context.Context) GetUserClbRegionsResultOutput {
	return o
}

func (o GetUserClbRegionsResultOutput) Datas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUserClbRegionsResult) []string { return v.Datas }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetUserClbRegionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserClbRegionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUserClbRegionsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserClbRegionsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetUserClbRegionsResultOutput) RichDatas() GetUserClbRegionsRichDataArrayOutput {
	return o.ApplyT(func(v GetUserClbRegionsResult) []GetUserClbRegionsRichData { return v.RichDatas }).(GetUserClbRegionsRichDataArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserClbRegionsResultOutput{})
}
