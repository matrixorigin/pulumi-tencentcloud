// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cwp

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMachinesSimple(ctx *pulumi.Context, args *GetMachinesSimpleArgs, opts ...pulumi.InvokeOption) (*GetMachinesSimpleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMachinesSimpleResult
	err := ctx.Invoke("tencentcloud:Cwp/getMachinesSimple:getMachinesSimple", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMachinesSimple.
type GetMachinesSimpleArgs struct {
	Filters          []GetMachinesSimpleFilter `pulumi:"filters"`
	MachineRegion    string                    `pulumi:"machineRegion"`
	MachineType      string                    `pulumi:"machineType"`
	ProjectIds       []int                     `pulumi:"projectIds"`
	ResultOutputFile *string                   `pulumi:"resultOutputFile"`
}

// A collection of values returned by getMachinesSimple.
type GetMachinesSimpleResult struct {
	Filters []GetMachinesSimpleFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                     `pulumi:"id"`
	MachineRegion    string                     `pulumi:"machineRegion"`
	MachineType      string                     `pulumi:"machineType"`
	Machines         []GetMachinesSimpleMachine `pulumi:"machines"`
	ProjectIds       []int                      `pulumi:"projectIds"`
	ResultOutputFile *string                    `pulumi:"resultOutputFile"`
}

func GetMachinesSimpleOutput(ctx *pulumi.Context, args GetMachinesSimpleOutputArgs, opts ...pulumi.InvokeOption) GetMachinesSimpleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMachinesSimpleResult, error) {
			args := v.(GetMachinesSimpleArgs)
			r, err := GetMachinesSimple(ctx, &args, opts...)
			var s GetMachinesSimpleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMachinesSimpleResultOutput)
}

// A collection of arguments for invoking getMachinesSimple.
type GetMachinesSimpleOutputArgs struct {
	Filters          GetMachinesSimpleFilterArrayInput `pulumi:"filters"`
	MachineRegion    pulumi.StringInput                `pulumi:"machineRegion"`
	MachineType      pulumi.StringInput                `pulumi:"machineType"`
	ProjectIds       pulumi.IntArrayInput              `pulumi:"projectIds"`
	ResultOutputFile pulumi.StringPtrInput             `pulumi:"resultOutputFile"`
}

func (GetMachinesSimpleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMachinesSimpleArgs)(nil)).Elem()
}

// A collection of values returned by getMachinesSimple.
type GetMachinesSimpleResultOutput struct{ *pulumi.OutputState }

func (GetMachinesSimpleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMachinesSimpleResult)(nil)).Elem()
}

func (o GetMachinesSimpleResultOutput) ToGetMachinesSimpleResultOutput() GetMachinesSimpleResultOutput {
	return o
}

func (o GetMachinesSimpleResultOutput) ToGetMachinesSimpleResultOutputWithContext(ctx context.Context) GetMachinesSimpleResultOutput {
	return o
}

func (o GetMachinesSimpleResultOutput) Filters() GetMachinesSimpleFilterArrayOutput {
	return o.ApplyT(func(v GetMachinesSimpleResult) []GetMachinesSimpleFilter { return v.Filters }).(GetMachinesSimpleFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMachinesSimpleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMachinesSimpleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetMachinesSimpleResultOutput) MachineRegion() pulumi.StringOutput {
	return o.ApplyT(func(v GetMachinesSimpleResult) string { return v.MachineRegion }).(pulumi.StringOutput)
}

func (o GetMachinesSimpleResultOutput) MachineType() pulumi.StringOutput {
	return o.ApplyT(func(v GetMachinesSimpleResult) string { return v.MachineType }).(pulumi.StringOutput)
}

func (o GetMachinesSimpleResultOutput) Machines() GetMachinesSimpleMachineArrayOutput {
	return o.ApplyT(func(v GetMachinesSimpleResult) []GetMachinesSimpleMachine { return v.Machines }).(GetMachinesSimpleMachineArrayOutput)
}

func (o GetMachinesSimpleResultOutput) ProjectIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetMachinesSimpleResult) []int { return v.ProjectIds }).(pulumi.IntArrayOutput)
}

func (o GetMachinesSimpleResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMachinesSimpleResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMachinesSimpleResultOutput{})
}