// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deprecatedroute

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTable(ctx *pulumi.Context, args *GetTableArgs, opts ...pulumi.InvokeOption) (*GetTableResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTableResult
	err := ctx.Invoke("tencentcloud:Deprecatedroute/getTable:getTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTable.
type GetTableArgs struct {
	Name         *string `pulumi:"name"`
	RouteTableId string  `pulumi:"routeTableId"`
}

// A collection of values returned by getTable.
type GetTableResult struct {
	CreateTime string `pulumi:"createTime"`
	// The provider-assigned unique ID for this managed resource.
	Id           string          `pulumi:"id"`
	Name         *string         `pulumi:"name"`
	RouteTableId string          `pulumi:"routeTableId"`
	Routes       []GetTableRoute `pulumi:"routes"`
	SubnetNum    int             `pulumi:"subnetNum"`
	VpcId        string          `pulumi:"vpcId"`
}

func GetTableOutput(ctx *pulumi.Context, args GetTableOutputArgs, opts ...pulumi.InvokeOption) GetTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTableResult, error) {
			args := v.(GetTableArgs)
			r, err := GetTable(ctx, &args, opts...)
			var s GetTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTableResultOutput)
}

// A collection of arguments for invoking getTable.
type GetTableOutputArgs struct {
	Name         pulumi.StringPtrInput `pulumi:"name"`
	RouteTableId pulumi.StringInput    `pulumi:"routeTableId"`
}

func (GetTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTableArgs)(nil)).Elem()
}

// A collection of values returned by getTable.
type GetTableResultOutput struct{ *pulumi.OutputState }

func (GetTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTableResult)(nil)).Elem()
}

func (o GetTableResultOutput) ToGetTableResultOutput() GetTableResultOutput {
	return o
}

func (o GetTableResultOutput) ToGetTableResultOutputWithContext(ctx context.Context) GetTableResultOutput {
	return o
}

func (o GetTableResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetTableResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTableResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTableResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetTableResultOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTableResult) string { return v.RouteTableId }).(pulumi.StringOutput)
}

func (o GetTableResultOutput) Routes() GetTableRouteArrayOutput {
	return o.ApplyT(func(v GetTableResult) []GetTableRoute { return v.Routes }).(GetTableRouteArrayOutput)
}

func (o GetTableResultOutput) SubnetNum() pulumi.IntOutput {
	return o.ApplyT(func(v GetTableResult) int { return v.SubnetNum }).(pulumi.IntOutput)
}

func (o GetTableResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTableResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTableResultOutput{})
}
