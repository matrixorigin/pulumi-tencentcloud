// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDomains(ctx *pulumi.Context, args *GetDomainsArgs, opts ...pulumi.InvokeOption) (*GetDomainsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDomainsResult
	err := ctx.Invoke("tencentcloud:Cdn/getDomains:getDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomains.
type GetDomainsArgs struct {
	Domain             *string `pulumi:"domain"`
	FullUrlCache       *bool   `pulumi:"fullUrlCache"`
	HttpsSwitch        *string `pulumi:"httpsSwitch"`
	OriginPullProtocol *string `pulumi:"originPullProtocol"`
	ResultOutputFile   *string `pulumi:"resultOutputFile"`
	ServiceType        *string `pulumi:"serviceType"`
}

// A collection of values returned by getDomains.
type GetDomainsResult struct {
	Domain       *string                `pulumi:"domain"`
	DomainLists  []GetDomainsDomainList `pulumi:"domainLists"`
	FullUrlCache *bool                  `pulumi:"fullUrlCache"`
	HttpsSwitch  *string                `pulumi:"httpsSwitch"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	OriginPullProtocol *string `pulumi:"originPullProtocol"`
	ResultOutputFile   *string `pulumi:"resultOutputFile"`
	ServiceType        *string `pulumi:"serviceType"`
}

func GetDomainsOutput(ctx *pulumi.Context, args GetDomainsOutputArgs, opts ...pulumi.InvokeOption) GetDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainsResult, error) {
			args := v.(GetDomainsArgs)
			r, err := GetDomains(ctx, &args, opts...)
			var s GetDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainsResultOutput)
}

// A collection of arguments for invoking getDomains.
type GetDomainsOutputArgs struct {
	Domain             pulumi.StringPtrInput `pulumi:"domain"`
	FullUrlCache       pulumi.BoolPtrInput   `pulumi:"fullUrlCache"`
	HttpsSwitch        pulumi.StringPtrInput `pulumi:"httpsSwitch"`
	OriginPullProtocol pulumi.StringPtrInput `pulumi:"originPullProtocol"`
	ResultOutputFile   pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	ServiceType        pulumi.StringPtrInput `pulumi:"serviceType"`
}

func (GetDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsArgs)(nil)).Elem()
}

// A collection of values returned by getDomains.
type GetDomainsResultOutput struct{ *pulumi.OutputState }

func (GetDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsResult)(nil)).Elem()
}

func (o GetDomainsResultOutput) ToGetDomainsResultOutput() GetDomainsResultOutput {
	return o
}

func (o GetDomainsResultOutput) ToGetDomainsResultOutputWithContext(ctx context.Context) GetDomainsResultOutput {
	return o
}

func (o GetDomainsResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o GetDomainsResultOutput) DomainLists() GetDomainsDomainListArrayOutput {
	return o.ApplyT(func(v GetDomainsResult) []GetDomainsDomainList { return v.DomainLists }).(GetDomainsDomainListArrayOutput)
}

func (o GetDomainsResultOutput) FullUrlCache() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *bool { return v.FullUrlCache }).(pulumi.BoolPtrOutput)
}

func (o GetDomainsResultOutput) HttpsSwitch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *string { return v.HttpsSwitch }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDomainsResultOutput) OriginPullProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *string { return v.OriginPullProtocol }).(pulumi.StringPtrOutput)
}

func (o GetDomainsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetDomainsResultOutput) ServiceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *string { return v.ServiceType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainsResultOutput{})
}
