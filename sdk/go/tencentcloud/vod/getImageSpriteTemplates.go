// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vod

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetImageSpriteTemplates(ctx *pulumi.Context, args *GetImageSpriteTemplatesArgs, opts ...pulumi.InvokeOption) (*GetImageSpriteTemplatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetImageSpriteTemplatesResult
	err := ctx.Invoke("tencentcloud:Vod/getImageSpriteTemplates:getImageSpriteTemplates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImageSpriteTemplates.
type GetImageSpriteTemplatesArgs struct {
	Definition       *string `pulumi:"definition"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SubAppId         *int    `pulumi:"subAppId"`
	Type             *string `pulumi:"type"`
}

// A collection of values returned by getImageSpriteTemplates.
type GetImageSpriteTemplatesResult struct {
	Definition *string `pulumi:"definition"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                                `pulumi:"id"`
	ResultOutputFile *string                               `pulumi:"resultOutputFile"`
	SubAppId         *int                                  `pulumi:"subAppId"`
	TemplateLists    []GetImageSpriteTemplatesTemplateList `pulumi:"templateLists"`
	Type             *string                               `pulumi:"type"`
}

func GetImageSpriteTemplatesOutput(ctx *pulumi.Context, args GetImageSpriteTemplatesOutputArgs, opts ...pulumi.InvokeOption) GetImageSpriteTemplatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetImageSpriteTemplatesResult, error) {
			args := v.(GetImageSpriteTemplatesArgs)
			r, err := GetImageSpriteTemplates(ctx, &args, opts...)
			var s GetImageSpriteTemplatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetImageSpriteTemplatesResultOutput)
}

// A collection of arguments for invoking getImageSpriteTemplates.
type GetImageSpriteTemplatesOutputArgs struct {
	Definition       pulumi.StringPtrInput `pulumi:"definition"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	SubAppId         pulumi.IntPtrInput    `pulumi:"subAppId"`
	Type             pulumi.StringPtrInput `pulumi:"type"`
}

func (GetImageSpriteTemplatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageSpriteTemplatesArgs)(nil)).Elem()
}

// A collection of values returned by getImageSpriteTemplates.
type GetImageSpriteTemplatesResultOutput struct{ *pulumi.OutputState }

func (GetImageSpriteTemplatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageSpriteTemplatesResult)(nil)).Elem()
}

func (o GetImageSpriteTemplatesResultOutput) ToGetImageSpriteTemplatesResultOutput() GetImageSpriteTemplatesResultOutput {
	return o
}

func (o GetImageSpriteTemplatesResultOutput) ToGetImageSpriteTemplatesResultOutputWithContext(ctx context.Context) GetImageSpriteTemplatesResultOutput {
	return o
}

func (o GetImageSpriteTemplatesResultOutput) Definition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageSpriteTemplatesResult) *string { return v.Definition }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetImageSpriteTemplatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageSpriteTemplatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetImageSpriteTemplatesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageSpriteTemplatesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetImageSpriteTemplatesResultOutput) SubAppId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetImageSpriteTemplatesResult) *int { return v.SubAppId }).(pulumi.IntPtrOutput)
}

func (o GetImageSpriteTemplatesResultOutput) TemplateLists() GetImageSpriteTemplatesTemplateListArrayOutput {
	return o.ApplyT(func(v GetImageSpriteTemplatesResult) []GetImageSpriteTemplatesTemplateList { return v.TemplateLists }).(GetImageSpriteTemplatesTemplateListArrayOutput)
}

func (o GetImageSpriteTemplatesResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageSpriteTemplatesResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetImageSpriteTemplatesResultOutput{})
}