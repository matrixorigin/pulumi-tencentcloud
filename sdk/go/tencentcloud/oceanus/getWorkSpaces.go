// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oceanus

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetWorkSpaces(ctx *pulumi.Context, args *GetWorkSpacesArgs, opts ...pulumi.InvokeOption) (*GetWorkSpacesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetWorkSpacesResult
	err := ctx.Invoke("tencentcloud:Oceanus/getWorkSpaces:getWorkSpaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWorkSpaces.
type GetWorkSpacesArgs struct {
	Filters          []GetWorkSpacesFilter `pulumi:"filters"`
	OrderType        *int                  `pulumi:"orderType"`
	ResultOutputFile *string               `pulumi:"resultOutputFile"`
}

// A collection of values returned by getWorkSpaces.
type GetWorkSpacesResult struct {
	Filters []GetWorkSpacesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                string                          `pulumi:"id"`
	OrderType         *int                            `pulumi:"orderType"`
	ResultOutputFile  *string                         `pulumi:"resultOutputFile"`
	WorkSpaceSetItems []GetWorkSpacesWorkSpaceSetItem `pulumi:"workSpaceSetItems"`
}

func GetWorkSpacesOutput(ctx *pulumi.Context, args GetWorkSpacesOutputArgs, opts ...pulumi.InvokeOption) GetWorkSpacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWorkSpacesResult, error) {
			args := v.(GetWorkSpacesArgs)
			r, err := GetWorkSpaces(ctx, &args, opts...)
			var s GetWorkSpacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWorkSpacesResultOutput)
}

// A collection of arguments for invoking getWorkSpaces.
type GetWorkSpacesOutputArgs struct {
	Filters          GetWorkSpacesFilterArrayInput `pulumi:"filters"`
	OrderType        pulumi.IntPtrInput            `pulumi:"orderType"`
	ResultOutputFile pulumi.StringPtrInput         `pulumi:"resultOutputFile"`
}

func (GetWorkSpacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkSpacesArgs)(nil)).Elem()
}

// A collection of values returned by getWorkSpaces.
type GetWorkSpacesResultOutput struct{ *pulumi.OutputState }

func (GetWorkSpacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkSpacesResult)(nil)).Elem()
}

func (o GetWorkSpacesResultOutput) ToGetWorkSpacesResultOutput() GetWorkSpacesResultOutput {
	return o
}

func (o GetWorkSpacesResultOutput) ToGetWorkSpacesResultOutputWithContext(ctx context.Context) GetWorkSpacesResultOutput {
	return o
}

func (o GetWorkSpacesResultOutput) Filters() GetWorkSpacesFilterArrayOutput {
	return o.ApplyT(func(v GetWorkSpacesResult) []GetWorkSpacesFilter { return v.Filters }).(GetWorkSpacesFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetWorkSpacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkSpacesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetWorkSpacesResultOutput) OrderType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetWorkSpacesResult) *int { return v.OrderType }).(pulumi.IntPtrOutput)
}

func (o GetWorkSpacesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWorkSpacesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetWorkSpacesResultOutput) WorkSpaceSetItems() GetWorkSpacesWorkSpaceSetItemArrayOutput {
	return o.ApplyT(func(v GetWorkSpacesResult) []GetWorkSpacesWorkSpaceSetItem { return v.WorkSpaceSetItems }).(GetWorkSpacesWorkSpaceSetItemArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWorkSpacesResultOutput{})
}
