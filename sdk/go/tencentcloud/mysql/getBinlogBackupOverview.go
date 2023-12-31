// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBinlogBackupOverview(ctx *pulumi.Context, args *GetBinlogBackupOverviewArgs, opts ...pulumi.InvokeOption) (*GetBinlogBackupOverviewResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBinlogBackupOverviewResult
	err := ctx.Invoke("tencentcloud:Mysql/getBinlogBackupOverview:getBinlogBackupOverview", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBinlogBackupOverview.
type GetBinlogBackupOverviewArgs struct {
	Product          string  `pulumi:"product"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getBinlogBackupOverview.
type GetBinlogBackupOverviewResult struct {
	BinlogArchiveCount  int `pulumi:"binlogArchiveCount"`
	BinlogArchiveVolume int `pulumi:"binlogArchiveVolume"`
	BinlogBackupCount   int `pulumi:"binlogBackupCount"`
	BinlogBackupVolume  int `pulumi:"binlogBackupVolume"`
	BinlogStandbyCount  int `pulumi:"binlogStandbyCount"`
	BinlogStandbyVolume int `pulumi:"binlogStandbyVolume"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	Product            string  `pulumi:"product"`
	RemoteBinlogCount  int     `pulumi:"remoteBinlogCount"`
	RemoteBinlogVolume int     `pulumi:"remoteBinlogVolume"`
	ResultOutputFile   *string `pulumi:"resultOutputFile"`
}

func GetBinlogBackupOverviewOutput(ctx *pulumi.Context, args GetBinlogBackupOverviewOutputArgs, opts ...pulumi.InvokeOption) GetBinlogBackupOverviewResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBinlogBackupOverviewResult, error) {
			args := v.(GetBinlogBackupOverviewArgs)
			r, err := GetBinlogBackupOverview(ctx, &args, opts...)
			var s GetBinlogBackupOverviewResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBinlogBackupOverviewResultOutput)
}

// A collection of arguments for invoking getBinlogBackupOverview.
type GetBinlogBackupOverviewOutputArgs struct {
	Product          pulumi.StringInput    `pulumi:"product"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetBinlogBackupOverviewOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBinlogBackupOverviewArgs)(nil)).Elem()
}

// A collection of values returned by getBinlogBackupOverview.
type GetBinlogBackupOverviewResultOutput struct{ *pulumi.OutputState }

func (GetBinlogBackupOverviewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBinlogBackupOverviewResult)(nil)).Elem()
}

func (o GetBinlogBackupOverviewResultOutput) ToGetBinlogBackupOverviewResultOutput() GetBinlogBackupOverviewResultOutput {
	return o
}

func (o GetBinlogBackupOverviewResultOutput) ToGetBinlogBackupOverviewResultOutputWithContext(ctx context.Context) GetBinlogBackupOverviewResultOutput {
	return o
}

func (o GetBinlogBackupOverviewResultOutput) BinlogArchiveCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetBinlogBackupOverviewResult) int { return v.BinlogArchiveCount }).(pulumi.IntOutput)
}

func (o GetBinlogBackupOverviewResultOutput) BinlogArchiveVolume() pulumi.IntOutput {
	return o.ApplyT(func(v GetBinlogBackupOverviewResult) int { return v.BinlogArchiveVolume }).(pulumi.IntOutput)
}

func (o GetBinlogBackupOverviewResultOutput) BinlogBackupCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetBinlogBackupOverviewResult) int { return v.BinlogBackupCount }).(pulumi.IntOutput)
}

func (o GetBinlogBackupOverviewResultOutput) BinlogBackupVolume() pulumi.IntOutput {
	return o.ApplyT(func(v GetBinlogBackupOverviewResult) int { return v.BinlogBackupVolume }).(pulumi.IntOutput)
}

func (o GetBinlogBackupOverviewResultOutput) BinlogStandbyCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetBinlogBackupOverviewResult) int { return v.BinlogStandbyCount }).(pulumi.IntOutput)
}

func (o GetBinlogBackupOverviewResultOutput) BinlogStandbyVolume() pulumi.IntOutput {
	return o.ApplyT(func(v GetBinlogBackupOverviewResult) int { return v.BinlogStandbyVolume }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBinlogBackupOverviewResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBinlogBackupOverviewResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBinlogBackupOverviewResultOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v GetBinlogBackupOverviewResult) string { return v.Product }).(pulumi.StringOutput)
}

func (o GetBinlogBackupOverviewResultOutput) RemoteBinlogCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetBinlogBackupOverviewResult) int { return v.RemoteBinlogCount }).(pulumi.IntOutput)
}

func (o GetBinlogBackupOverviewResultOutput) RemoteBinlogVolume() pulumi.IntOutput {
	return o.ApplyT(func(v GetBinlogBackupOverviewResult) int { return v.RemoteBinlogVolume }).(pulumi.IntOutput)
}

func (o GetBinlogBackupOverviewResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBinlogBackupOverviewResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBinlogBackupOverviewResultOutput{})
}
