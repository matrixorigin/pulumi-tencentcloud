// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clickhouse

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetInstanceNodes(ctx *pulumi.Context, args *GetInstanceNodesArgs, opts ...pulumi.InvokeOption) (*GetInstanceNodesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstanceNodesResult
	err := ctx.Invoke("tencentcloud:Clickhouse/getInstanceNodes:getInstanceNodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceNodes.
type GetInstanceNodesArgs struct {
	DisplayPolicy    *string `pulumi:"displayPolicy"`
	ForceAll         *bool   `pulumi:"forceAll"`
	InstanceId       string  `pulumi:"instanceId"`
	NodeRole         *string `pulumi:"nodeRole"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstanceNodes.
type GetInstanceNodesResult struct {
	DisplayPolicy *string `pulumi:"displayPolicy"`
	ForceAll      *bool   `pulumi:"forceAll"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string                              `pulumi:"id"`
	InstanceId         string                              `pulumi:"instanceId"`
	InstanceNodesLists []GetInstanceNodesInstanceNodesList `pulumi:"instanceNodesLists"`
	NodeRole           *string                             `pulumi:"nodeRole"`
	ResultOutputFile   *string                             `pulumi:"resultOutputFile"`
}

func GetInstanceNodesOutput(ctx *pulumi.Context, args GetInstanceNodesOutputArgs, opts ...pulumi.InvokeOption) GetInstanceNodesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceNodesResult, error) {
			args := v.(GetInstanceNodesArgs)
			r, err := GetInstanceNodes(ctx, &args, opts...)
			var s GetInstanceNodesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceNodesResultOutput)
}

// A collection of arguments for invoking getInstanceNodes.
type GetInstanceNodesOutputArgs struct {
	DisplayPolicy    pulumi.StringPtrInput `pulumi:"displayPolicy"`
	ForceAll         pulumi.BoolPtrInput   `pulumi:"forceAll"`
	InstanceId       pulumi.StringInput    `pulumi:"instanceId"`
	NodeRole         pulumi.StringPtrInput `pulumi:"nodeRole"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetInstanceNodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceNodesArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceNodes.
type GetInstanceNodesResultOutput struct{ *pulumi.OutputState }

func (GetInstanceNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceNodesResult)(nil)).Elem()
}

func (o GetInstanceNodesResultOutput) ToGetInstanceNodesResultOutput() GetInstanceNodesResultOutput {
	return o
}

func (o GetInstanceNodesResultOutput) ToGetInstanceNodesResultOutputWithContext(ctx context.Context) GetInstanceNodesResultOutput {
	return o
}

func (o GetInstanceNodesResultOutput) DisplayPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceNodesResult) *string { return v.DisplayPolicy }).(pulumi.StringPtrOutput)
}

func (o GetInstanceNodesResultOutput) ForceAll() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetInstanceNodesResult) *bool { return v.ForceAll }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceNodesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceNodesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstanceNodesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceNodesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetInstanceNodesResultOutput) InstanceNodesLists() GetInstanceNodesInstanceNodesListArrayOutput {
	return o.ApplyT(func(v GetInstanceNodesResult) []GetInstanceNodesInstanceNodesList { return v.InstanceNodesLists }).(GetInstanceNodesInstanceNodesListArrayOutput)
}

func (o GetInstanceNodesResultOutput) NodeRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceNodesResult) *string { return v.NodeRole }).(pulumi.StringPtrOutput)
}

func (o GetInstanceNodesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceNodesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceNodesResultOutput{})
}
