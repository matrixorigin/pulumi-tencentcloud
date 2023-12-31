// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetStatisticsReport(ctx *pulumi.Context, args *GetStatisticsReportArgs, opts ...pulumi.InvokeOption) (*GetStatisticsReportResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetStatisticsReportResult
	err := ctx.Invoke("tencentcloud:Ses/getStatisticsReport:getStatisticsReport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStatisticsReport.
type GetStatisticsReportArgs struct {
	Domain               *string `pulumi:"domain"`
	EndDate              string  `pulumi:"endDate"`
	ReceivingMailboxType *string `pulumi:"receivingMailboxType"`
	ResultOutputFile     *string `pulumi:"resultOutputFile"`
	StartDate            string  `pulumi:"startDate"`
}

// A collection of values returned by getStatisticsReport.
type GetStatisticsReportResult struct {
	DailyVolumes []GetStatisticsReportDailyVolume `pulumi:"dailyVolumes"`
	Domain       *string                          `pulumi:"domain"`
	EndDate      string                           `pulumi:"endDate"`
	// The provider-assigned unique ID for this managed resource.
	Id                   string                             `pulumi:"id"`
	OverallVolumes       []GetStatisticsReportOverallVolume `pulumi:"overallVolumes"`
	ReceivingMailboxType *string                            `pulumi:"receivingMailboxType"`
	ResultOutputFile     *string                            `pulumi:"resultOutputFile"`
	StartDate            string                             `pulumi:"startDate"`
}

func GetStatisticsReportOutput(ctx *pulumi.Context, args GetStatisticsReportOutputArgs, opts ...pulumi.InvokeOption) GetStatisticsReportResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStatisticsReportResult, error) {
			args := v.(GetStatisticsReportArgs)
			r, err := GetStatisticsReport(ctx, &args, opts...)
			var s GetStatisticsReportResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetStatisticsReportResultOutput)
}

// A collection of arguments for invoking getStatisticsReport.
type GetStatisticsReportOutputArgs struct {
	Domain               pulumi.StringPtrInput `pulumi:"domain"`
	EndDate              pulumi.StringInput    `pulumi:"endDate"`
	ReceivingMailboxType pulumi.StringPtrInput `pulumi:"receivingMailboxType"`
	ResultOutputFile     pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	StartDate            pulumi.StringInput    `pulumi:"startDate"`
}

func (GetStatisticsReportOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStatisticsReportArgs)(nil)).Elem()
}

// A collection of values returned by getStatisticsReport.
type GetStatisticsReportResultOutput struct{ *pulumi.OutputState }

func (GetStatisticsReportResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStatisticsReportResult)(nil)).Elem()
}

func (o GetStatisticsReportResultOutput) ToGetStatisticsReportResultOutput() GetStatisticsReportResultOutput {
	return o
}

func (o GetStatisticsReportResultOutput) ToGetStatisticsReportResultOutputWithContext(ctx context.Context) GetStatisticsReportResultOutput {
	return o
}

func (o GetStatisticsReportResultOutput) DailyVolumes() GetStatisticsReportDailyVolumeArrayOutput {
	return o.ApplyT(func(v GetStatisticsReportResult) []GetStatisticsReportDailyVolume { return v.DailyVolumes }).(GetStatisticsReportDailyVolumeArrayOutput)
}

func (o GetStatisticsReportResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStatisticsReportResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o GetStatisticsReportResultOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetStatisticsReportResult) string { return v.EndDate }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStatisticsReportResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStatisticsReportResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStatisticsReportResultOutput) OverallVolumes() GetStatisticsReportOverallVolumeArrayOutput {
	return o.ApplyT(func(v GetStatisticsReportResult) []GetStatisticsReportOverallVolume { return v.OverallVolumes }).(GetStatisticsReportOverallVolumeArrayOutput)
}

func (o GetStatisticsReportResultOutput) ReceivingMailboxType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStatisticsReportResult) *string { return v.ReceivingMailboxType }).(pulumi.StringPtrOutput)
}

func (o GetStatisticsReportResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStatisticsReportResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetStatisticsReportResultOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetStatisticsReportResult) string { return v.StartDate }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStatisticsReportResultOutput{})
}
