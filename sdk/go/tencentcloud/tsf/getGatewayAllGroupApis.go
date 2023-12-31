// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayAllGroupApis(ctx *pulumi.Context, args *LookupGatewayAllGroupApisArgs, opts ...pulumi.InvokeOption) (*LookupGatewayAllGroupApisResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayAllGroupApisResult
	err := ctx.Invoke("tencentcloud:Tsf/getGatewayAllGroupApis:getGatewayAllGroupApis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayAllGroupApis.
type LookupGatewayAllGroupApisArgs struct {
	GatewayDeployGroupId string  `pulumi:"gatewayDeployGroupId"`
	ResultOutputFile     *string `pulumi:"resultOutputFile"`
	SearchWord           *string `pulumi:"searchWord"`
}

// A collection of values returned by getGatewayAllGroupApis.
type LookupGatewayAllGroupApisResult struct {
	GatewayDeployGroupId string `pulumi:"gatewayDeployGroupId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                         `pulumi:"id"`
	ResultOutputFile *string                        `pulumi:"resultOutputFile"`
	Results          []GetGatewayAllGroupApisResult `pulumi:"results"`
	SearchWord       *string                        `pulumi:"searchWord"`
}

func LookupGatewayAllGroupApisOutput(ctx *pulumi.Context, args LookupGatewayAllGroupApisOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayAllGroupApisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayAllGroupApisResult, error) {
			args := v.(LookupGatewayAllGroupApisArgs)
			r, err := LookupGatewayAllGroupApis(ctx, &args, opts...)
			var s LookupGatewayAllGroupApisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayAllGroupApisResultOutput)
}

// A collection of arguments for invoking getGatewayAllGroupApis.
type LookupGatewayAllGroupApisOutputArgs struct {
	GatewayDeployGroupId pulumi.StringInput    `pulumi:"gatewayDeployGroupId"`
	ResultOutputFile     pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	SearchWord           pulumi.StringPtrInput `pulumi:"searchWord"`
}

func (LookupGatewayAllGroupApisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayAllGroupApisArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayAllGroupApis.
type LookupGatewayAllGroupApisResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayAllGroupApisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayAllGroupApisResult)(nil)).Elem()
}

func (o LookupGatewayAllGroupApisResultOutput) ToLookupGatewayAllGroupApisResultOutput() LookupGatewayAllGroupApisResultOutput {
	return o
}

func (o LookupGatewayAllGroupApisResultOutput) ToLookupGatewayAllGroupApisResultOutputWithContext(ctx context.Context) LookupGatewayAllGroupApisResultOutput {
	return o
}

func (o LookupGatewayAllGroupApisResultOutput) GatewayDeployGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayAllGroupApisResult) string { return v.GatewayDeployGroupId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGatewayAllGroupApisResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayAllGroupApisResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayAllGroupApisResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayAllGroupApisResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupGatewayAllGroupApisResultOutput) Results() GetGatewayAllGroupApisResultArrayOutput {
	return o.ApplyT(func(v LookupGatewayAllGroupApisResult) []GetGatewayAllGroupApisResult { return v.Results }).(GetGatewayAllGroupApisResultArrayOutput)
}

func (o LookupGatewayAllGroupApisResultOutput) SearchWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayAllGroupApisResult) *string { return v.SearchWord }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayAllGroupApisResultOutput{})
}
