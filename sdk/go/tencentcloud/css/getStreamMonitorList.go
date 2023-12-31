// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package css

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetStreamMonitorList(ctx *pulumi.Context, args *GetStreamMonitorListArgs, opts ...pulumi.InvokeOption) (*GetStreamMonitorListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetStreamMonitorListResult
	err := ctx.Invoke("tencentcloud:Css/getStreamMonitorList:getStreamMonitorList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStreamMonitorList.
type GetStreamMonitorListArgs struct {
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getStreamMonitorList.
type GetStreamMonitorListResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                 string                                  `pulumi:"id"`
	LiveStreamMonitors []GetStreamMonitorListLiveStreamMonitor `pulumi:"liveStreamMonitors"`
	ResultOutputFile   *string                                 `pulumi:"resultOutputFile"`
}

func GetStreamMonitorListOutput(ctx *pulumi.Context, args GetStreamMonitorListOutputArgs, opts ...pulumi.InvokeOption) GetStreamMonitorListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStreamMonitorListResult, error) {
			args := v.(GetStreamMonitorListArgs)
			r, err := GetStreamMonitorList(ctx, &args, opts...)
			var s GetStreamMonitorListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetStreamMonitorListResultOutput)
}

// A collection of arguments for invoking getStreamMonitorList.
type GetStreamMonitorListOutputArgs struct {
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetStreamMonitorListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStreamMonitorListArgs)(nil)).Elem()
}

// A collection of values returned by getStreamMonitorList.
type GetStreamMonitorListResultOutput struct{ *pulumi.OutputState }

func (GetStreamMonitorListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStreamMonitorListResult)(nil)).Elem()
}

func (o GetStreamMonitorListResultOutput) ToGetStreamMonitorListResultOutput() GetStreamMonitorListResultOutput {
	return o
}

func (o GetStreamMonitorListResultOutput) ToGetStreamMonitorListResultOutputWithContext(ctx context.Context) GetStreamMonitorListResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetStreamMonitorListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStreamMonitorListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStreamMonitorListResultOutput) LiveStreamMonitors() GetStreamMonitorListLiveStreamMonitorArrayOutput {
	return o.ApplyT(func(v GetStreamMonitorListResult) []GetStreamMonitorListLiveStreamMonitor {
		return v.LiveStreamMonitors
	}).(GetStreamMonitorListLiveStreamMonitorArrayOutput)
}

func (o GetStreamMonitorListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStreamMonitorListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStreamMonitorListResultOutput{})
}
