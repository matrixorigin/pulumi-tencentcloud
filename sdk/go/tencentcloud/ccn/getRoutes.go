// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ccn

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoutes(ctx *pulumi.Context, args *LookupRoutesArgs, opts ...pulumi.InvokeOption) (*LookupRoutesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRoutesResult
	err := ctx.Invoke("tencentcloud:Ccn/getRoutes:getRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRoutes.
type LookupRoutesArgs struct {
	CcnId            string            `pulumi:"ccnId"`
	Filters          []GetRoutesFilter `pulumi:"filters"`
	ResultOutputFile *string           `pulumi:"resultOutputFile"`
}

// A collection of values returned by getRoutes.
type LookupRoutesResult struct {
	CcnId   string            `pulumi:"ccnId"`
	Filters []GetRoutesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string               `pulumi:"id"`
	ResultOutputFile *string              `pulumi:"resultOutputFile"`
	RouteLists       []GetRoutesRouteList `pulumi:"routeLists"`
}

func LookupRoutesOutput(ctx *pulumi.Context, args LookupRoutesOutputArgs, opts ...pulumi.InvokeOption) LookupRoutesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoutesResult, error) {
			args := v.(LookupRoutesArgs)
			r, err := LookupRoutes(ctx, &args, opts...)
			var s LookupRoutesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoutesResultOutput)
}

// A collection of arguments for invoking getRoutes.
type LookupRoutesOutputArgs struct {
	CcnId            pulumi.StringInput        `pulumi:"ccnId"`
	Filters          GetRoutesFilterArrayInput `pulumi:"filters"`
	ResultOutputFile pulumi.StringPtrInput     `pulumi:"resultOutputFile"`
}

func (LookupRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutesArgs)(nil)).Elem()
}

// A collection of values returned by getRoutes.
type LookupRoutesResultOutput struct{ *pulumi.OutputState }

func (LookupRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutesResult)(nil)).Elem()
}

func (o LookupRoutesResultOutput) ToLookupRoutesResultOutput() LookupRoutesResultOutput {
	return o
}

func (o LookupRoutesResultOutput) ToLookupRoutesResultOutputWithContext(ctx context.Context) LookupRoutesResultOutput {
	return o
}

func (o LookupRoutesResultOutput) CcnId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutesResult) string { return v.CcnId }).(pulumi.StringOutput)
}

func (o LookupRoutesResultOutput) Filters() GetRoutesFilterArrayOutput {
	return o.ApplyT(func(v LookupRoutesResult) []GetRoutesFilter { return v.Filters }).(GetRoutesFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRoutesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRoutesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoutesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupRoutesResultOutput) RouteLists() GetRoutesRouteListArrayOutput {
	return o.ApplyT(func(v LookupRoutesResult) []GetRoutesRouteList { return v.RouteLists }).(GetRoutesRouteListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoutesResultOutput{})
}
