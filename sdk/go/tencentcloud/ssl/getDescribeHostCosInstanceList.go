// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDescribeHostCosInstanceList(ctx *pulumi.Context, args *GetDescribeHostCosInstanceListArgs, opts ...pulumi.InvokeOption) (*GetDescribeHostCosInstanceListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDescribeHostCosInstanceListResult
	err := ctx.Invoke("tencentcloud:Ssl/getDescribeHostCosInstanceList:getDescribeHostCosInstanceList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDescribeHostCosInstanceList.
type GetDescribeHostCosInstanceListArgs struct {
	CertificateId    string                                 `pulumi:"certificateId"`
	Filters          []GetDescribeHostCosInstanceListFilter `pulumi:"filters"`
	IsCache          *int                                   `pulumi:"isCache"`
	ResourceType     string                                 `pulumi:"resourceType"`
	ResultOutputFile *string                                `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDescribeHostCosInstanceList.
type GetDescribeHostCosInstanceListResult struct {
	AsyncCacheTime string                                 `pulumi:"asyncCacheTime"`
	AsyncOffset    int                                    `pulumi:"asyncOffset"`
	AsyncTotalNum  int                                    `pulumi:"asyncTotalNum"`
	CertificateId  string                                 `pulumi:"certificateId"`
	Filters        []GetDescribeHostCosInstanceListFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                                       `pulumi:"id"`
	InstanceLists    []GetDescribeHostCosInstanceListInstanceList `pulumi:"instanceLists"`
	IsCache          *int                                         `pulumi:"isCache"`
	ResourceType     string                                       `pulumi:"resourceType"`
	ResultOutputFile *string                                      `pulumi:"resultOutputFile"`
}

func GetDescribeHostCosInstanceListOutput(ctx *pulumi.Context, args GetDescribeHostCosInstanceListOutputArgs, opts ...pulumi.InvokeOption) GetDescribeHostCosInstanceListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDescribeHostCosInstanceListResult, error) {
			args := v.(GetDescribeHostCosInstanceListArgs)
			r, err := GetDescribeHostCosInstanceList(ctx, &args, opts...)
			var s GetDescribeHostCosInstanceListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDescribeHostCosInstanceListResultOutput)
}

// A collection of arguments for invoking getDescribeHostCosInstanceList.
type GetDescribeHostCosInstanceListOutputArgs struct {
	CertificateId    pulumi.StringInput                             `pulumi:"certificateId"`
	Filters          GetDescribeHostCosInstanceListFilterArrayInput `pulumi:"filters"`
	IsCache          pulumi.IntPtrInput                             `pulumi:"isCache"`
	ResourceType     pulumi.StringInput                             `pulumi:"resourceType"`
	ResultOutputFile pulumi.StringPtrInput                          `pulumi:"resultOutputFile"`
}

func (GetDescribeHostCosInstanceListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeHostCosInstanceListArgs)(nil)).Elem()
}

// A collection of values returned by getDescribeHostCosInstanceList.
type GetDescribeHostCosInstanceListResultOutput struct{ *pulumi.OutputState }

func (GetDescribeHostCosInstanceListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDescribeHostCosInstanceListResult)(nil)).Elem()
}

func (o GetDescribeHostCosInstanceListResultOutput) ToGetDescribeHostCosInstanceListResultOutput() GetDescribeHostCosInstanceListResultOutput {
	return o
}

func (o GetDescribeHostCosInstanceListResultOutput) ToGetDescribeHostCosInstanceListResultOutputWithContext(ctx context.Context) GetDescribeHostCosInstanceListResultOutput {
	return o
}

func (o GetDescribeHostCosInstanceListResultOutput) AsyncCacheTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostCosInstanceListResult) string { return v.AsyncCacheTime }).(pulumi.StringOutput)
}

func (o GetDescribeHostCosInstanceListResultOutput) AsyncOffset() pulumi.IntOutput {
	return o.ApplyT(func(v GetDescribeHostCosInstanceListResult) int { return v.AsyncOffset }).(pulumi.IntOutput)
}

func (o GetDescribeHostCosInstanceListResultOutput) AsyncTotalNum() pulumi.IntOutput {
	return o.ApplyT(func(v GetDescribeHostCosInstanceListResult) int { return v.AsyncTotalNum }).(pulumi.IntOutput)
}

func (o GetDescribeHostCosInstanceListResultOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostCosInstanceListResult) string { return v.CertificateId }).(pulumi.StringOutput)
}

func (o GetDescribeHostCosInstanceListResultOutput) Filters() GetDescribeHostCosInstanceListFilterArrayOutput {
	return o.ApplyT(func(v GetDescribeHostCosInstanceListResult) []GetDescribeHostCosInstanceListFilter { return v.Filters }).(GetDescribeHostCosInstanceListFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDescribeHostCosInstanceListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostCosInstanceListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDescribeHostCosInstanceListResultOutput) InstanceLists() GetDescribeHostCosInstanceListInstanceListArrayOutput {
	return o.ApplyT(func(v GetDescribeHostCosInstanceListResult) []GetDescribeHostCosInstanceListInstanceList {
		return v.InstanceLists
	}).(GetDescribeHostCosInstanceListInstanceListArrayOutput)
}

func (o GetDescribeHostCosInstanceListResultOutput) IsCache() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDescribeHostCosInstanceListResult) *int { return v.IsCache }).(pulumi.IntPtrOutput)
}

func (o GetDescribeHostCosInstanceListResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDescribeHostCosInstanceListResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o GetDescribeHostCosInstanceListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDescribeHostCosInstanceListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDescribeHostCosInstanceListResultOutput{})
}
