// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGatewayFlowMonitorDetail(ctx *pulumi.Context, args *GetGatewayFlowMonitorDetailArgs, opts ...pulumi.InvokeOption) (*GetGatewayFlowMonitorDetailResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGatewayFlowMonitorDetailResult
	err := ctx.Invoke("tencentcloud:Vpc/getGatewayFlowMonitorDetail:getGatewayFlowMonitorDetail", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayFlowMonitorDetail.
type GetGatewayFlowMonitorDetailArgs struct {
	DirectConnectGatewayId *string `pulumi:"directConnectGatewayId"`
	NatId                  *string `pulumi:"natId"`
	OrderDirection         *string `pulumi:"orderDirection"`
	OrderField             *string `pulumi:"orderField"`
	PeeringConnectionId    *string `pulumi:"peeringConnectionId"`
	ResultOutputFile       *string `pulumi:"resultOutputFile"`
	TimePoint              string  `pulumi:"timePoint"`
	VpnId                  *string `pulumi:"vpnId"`
}

// A collection of values returned by getGatewayFlowMonitorDetail.
type GetGatewayFlowMonitorDetailResult struct {
	DirectConnectGatewayId       *string                                                  `pulumi:"directConnectGatewayId"`
	GatewayFlowMonitorDetailSets []GetGatewayFlowMonitorDetailGatewayFlowMonitorDetailSet `pulumi:"gatewayFlowMonitorDetailSets"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string  `pulumi:"id"`
	NatId               *string `pulumi:"natId"`
	OrderDirection      *string `pulumi:"orderDirection"`
	OrderField          *string `pulumi:"orderField"`
	PeeringConnectionId *string `pulumi:"peeringConnectionId"`
	ResultOutputFile    *string `pulumi:"resultOutputFile"`
	TimePoint           string  `pulumi:"timePoint"`
	VpnId               *string `pulumi:"vpnId"`
}

func GetGatewayFlowMonitorDetailOutput(ctx *pulumi.Context, args GetGatewayFlowMonitorDetailOutputArgs, opts ...pulumi.InvokeOption) GetGatewayFlowMonitorDetailResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGatewayFlowMonitorDetailResult, error) {
			args := v.(GetGatewayFlowMonitorDetailArgs)
			r, err := GetGatewayFlowMonitorDetail(ctx, &args, opts...)
			var s GetGatewayFlowMonitorDetailResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGatewayFlowMonitorDetailResultOutput)
}

// A collection of arguments for invoking getGatewayFlowMonitorDetail.
type GetGatewayFlowMonitorDetailOutputArgs struct {
	DirectConnectGatewayId pulumi.StringPtrInput `pulumi:"directConnectGatewayId"`
	NatId                  pulumi.StringPtrInput `pulumi:"natId"`
	OrderDirection         pulumi.StringPtrInput `pulumi:"orderDirection"`
	OrderField             pulumi.StringPtrInput `pulumi:"orderField"`
	PeeringConnectionId    pulumi.StringPtrInput `pulumi:"peeringConnectionId"`
	ResultOutputFile       pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	TimePoint              pulumi.StringInput    `pulumi:"timePoint"`
	VpnId                  pulumi.StringPtrInput `pulumi:"vpnId"`
}

func (GetGatewayFlowMonitorDetailOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewayFlowMonitorDetailArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayFlowMonitorDetail.
type GetGatewayFlowMonitorDetailResultOutput struct{ *pulumi.OutputState }

func (GetGatewayFlowMonitorDetailResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewayFlowMonitorDetailResult)(nil)).Elem()
}

func (o GetGatewayFlowMonitorDetailResultOutput) ToGetGatewayFlowMonitorDetailResultOutput() GetGatewayFlowMonitorDetailResultOutput {
	return o
}

func (o GetGatewayFlowMonitorDetailResultOutput) ToGetGatewayFlowMonitorDetailResultOutputWithContext(ctx context.Context) GetGatewayFlowMonitorDetailResultOutput {
	return o
}

func (o GetGatewayFlowMonitorDetailResultOutput) DirectConnectGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewayFlowMonitorDetailResult) *string { return v.DirectConnectGatewayId }).(pulumi.StringPtrOutput)
}

func (o GetGatewayFlowMonitorDetailResultOutput) GatewayFlowMonitorDetailSets() GetGatewayFlowMonitorDetailGatewayFlowMonitorDetailSetArrayOutput {
	return o.ApplyT(func(v GetGatewayFlowMonitorDetailResult) []GetGatewayFlowMonitorDetailGatewayFlowMonitorDetailSet {
		return v.GatewayFlowMonitorDetailSets
	}).(GetGatewayFlowMonitorDetailGatewayFlowMonitorDetailSetArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGatewayFlowMonitorDetailResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewayFlowMonitorDetailResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGatewayFlowMonitorDetailResultOutput) NatId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewayFlowMonitorDetailResult) *string { return v.NatId }).(pulumi.StringPtrOutput)
}

func (o GetGatewayFlowMonitorDetailResultOutput) OrderDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewayFlowMonitorDetailResult) *string { return v.OrderDirection }).(pulumi.StringPtrOutput)
}

func (o GetGatewayFlowMonitorDetailResultOutput) OrderField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewayFlowMonitorDetailResult) *string { return v.OrderField }).(pulumi.StringPtrOutput)
}

func (o GetGatewayFlowMonitorDetailResultOutput) PeeringConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewayFlowMonitorDetailResult) *string { return v.PeeringConnectionId }).(pulumi.StringPtrOutput)
}

func (o GetGatewayFlowMonitorDetailResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewayFlowMonitorDetailResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetGatewayFlowMonitorDetailResultOutput) TimePoint() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewayFlowMonitorDetailResult) string { return v.TimePoint }).(pulumi.StringOutput)
}

func (o GetGatewayFlowMonitorDetailResultOutput) VpnId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewayFlowMonitorDetailResult) *string { return v.VpnId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGatewayFlowMonitorDetailResultOutput{})
}
