// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiDetail(ctx *pulumi.Context, args *LookupApiDetailArgs, opts ...pulumi.InvokeOption) (*LookupApiDetailResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApiDetailResult
	err := ctx.Invoke("tencentcloud:Tsf/getApiDetail:getApiDetail", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApiDetail.
type LookupApiDetailArgs struct {
	ApplicationId    string  `pulumi:"applicationId"`
	Method           string  `pulumi:"method"`
	MicroserviceId   string  `pulumi:"microserviceId"`
	Path             string  `pulumi:"path"`
	PkgVersion       string  `pulumi:"pkgVersion"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getApiDetail.
type LookupApiDetailResult struct {
	ApplicationId string `pulumi:"applicationId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string               `pulumi:"id"`
	Method           string               `pulumi:"method"`
	MicroserviceId   string               `pulumi:"microserviceId"`
	Path             string               `pulumi:"path"`
	PkgVersion       string               `pulumi:"pkgVersion"`
	ResultOutputFile *string              `pulumi:"resultOutputFile"`
	Results          []GetApiDetailResult `pulumi:"results"`
}

func LookupApiDetailOutput(ctx *pulumi.Context, args LookupApiDetailOutputArgs, opts ...pulumi.InvokeOption) LookupApiDetailResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiDetailResult, error) {
			args := v.(LookupApiDetailArgs)
			r, err := LookupApiDetail(ctx, &args, opts...)
			var s LookupApiDetailResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiDetailResultOutput)
}

// A collection of arguments for invoking getApiDetail.
type LookupApiDetailOutputArgs struct {
	ApplicationId    pulumi.StringInput    `pulumi:"applicationId"`
	Method           pulumi.StringInput    `pulumi:"method"`
	MicroserviceId   pulumi.StringInput    `pulumi:"microserviceId"`
	Path             pulumi.StringInput    `pulumi:"path"`
	PkgVersion       pulumi.StringInput    `pulumi:"pkgVersion"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (LookupApiDetailOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiDetailArgs)(nil)).Elem()
}

// A collection of values returned by getApiDetail.
type LookupApiDetailResultOutput struct{ *pulumi.OutputState }

func (LookupApiDetailResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiDetailResult)(nil)).Elem()
}

func (o LookupApiDetailResultOutput) ToLookupApiDetailResultOutput() LookupApiDetailResultOutput {
	return o
}

func (o LookupApiDetailResultOutput) ToLookupApiDetailResultOutputWithContext(ctx context.Context) LookupApiDetailResultOutput {
	return o
}

func (o LookupApiDetailResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDetailResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupApiDetailResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDetailResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiDetailResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDetailResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o LookupApiDetailResultOutput) MicroserviceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDetailResult) string { return v.MicroserviceId }).(pulumi.StringOutput)
}

func (o LookupApiDetailResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDetailResult) string { return v.Path }).(pulumi.StringOutput)
}

func (o LookupApiDetailResultOutput) PkgVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDetailResult) string { return v.PkgVersion }).(pulumi.StringOutput)
}

func (o LookupApiDetailResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiDetailResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupApiDetailResultOutput) Results() GetApiDetailResultArrayOutput {
	return o.ApplyT(func(v LookupApiDetailResult) []GetApiDetailResult { return v.Results }).(GetApiDetailResultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiDetailResultOutput{})
}