// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAlarmAllNamespaces(ctx *pulumi.Context, args *GetAlarmAllNamespacesArgs, opts ...pulumi.InvokeOption) (*GetAlarmAllNamespacesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAlarmAllNamespacesResult
	err := ctx.Invoke("tencentcloud:Monitor/getAlarmAllNamespaces:getAlarmAllNamespaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlarmAllNamespaces.
type GetAlarmAllNamespacesArgs struct {
	Ids              []string `pulumi:"ids"`
	Module           string   `pulumi:"module"`
	MonitorTypes     []string `pulumi:"monitorTypes"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	SceneType        string   `pulumi:"sceneType"`
}

// A collection of values returned by getAlarmAllNamespaces.
type GetAlarmAllNamespacesResult struct {
	CommonNamespaces     []GetAlarmAllNamespacesCommonNamespace      `pulumi:"commonNamespaces"`
	CustomNamespacesNews []GetAlarmAllNamespacesCustomNamespacesNews `pulumi:"customNamespacesNews"`
	// The provider-assigned unique ID for this managed resource.
	Id                string                                   `pulumi:"id"`
	Ids               []string                                 `pulumi:"ids"`
	Module            string                                   `pulumi:"module"`
	MonitorTypes      []string                                 `pulumi:"monitorTypes"`
	QceNamespacesNews []GetAlarmAllNamespacesQceNamespacesNews `pulumi:"qceNamespacesNews"`
	ResultOutputFile  *string                                  `pulumi:"resultOutputFile"`
	SceneType         string                                   `pulumi:"sceneType"`
}

func GetAlarmAllNamespacesOutput(ctx *pulumi.Context, args GetAlarmAllNamespacesOutputArgs, opts ...pulumi.InvokeOption) GetAlarmAllNamespacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAlarmAllNamespacesResult, error) {
			args := v.(GetAlarmAllNamespacesArgs)
			r, err := GetAlarmAllNamespaces(ctx, &args, opts...)
			var s GetAlarmAllNamespacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAlarmAllNamespacesResultOutput)
}

// A collection of arguments for invoking getAlarmAllNamespaces.
type GetAlarmAllNamespacesOutputArgs struct {
	Ids              pulumi.StringArrayInput `pulumi:"ids"`
	Module           pulumi.StringInput      `pulumi:"module"`
	MonitorTypes     pulumi.StringArrayInput `pulumi:"monitorTypes"`
	ResultOutputFile pulumi.StringPtrInput   `pulumi:"resultOutputFile"`
	SceneType        pulumi.StringInput      `pulumi:"sceneType"`
}

func (GetAlarmAllNamespacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlarmAllNamespacesArgs)(nil)).Elem()
}

// A collection of values returned by getAlarmAllNamespaces.
type GetAlarmAllNamespacesResultOutput struct{ *pulumi.OutputState }

func (GetAlarmAllNamespacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlarmAllNamespacesResult)(nil)).Elem()
}

func (o GetAlarmAllNamespacesResultOutput) ToGetAlarmAllNamespacesResultOutput() GetAlarmAllNamespacesResultOutput {
	return o
}

func (o GetAlarmAllNamespacesResultOutput) ToGetAlarmAllNamespacesResultOutputWithContext(ctx context.Context) GetAlarmAllNamespacesResultOutput {
	return o
}

func (o GetAlarmAllNamespacesResultOutput) CommonNamespaces() GetAlarmAllNamespacesCommonNamespaceArrayOutput {
	return o.ApplyT(func(v GetAlarmAllNamespacesResult) []GetAlarmAllNamespacesCommonNamespace { return v.CommonNamespaces }).(GetAlarmAllNamespacesCommonNamespaceArrayOutput)
}

func (o GetAlarmAllNamespacesResultOutput) CustomNamespacesNews() GetAlarmAllNamespacesCustomNamespacesNewsArrayOutput {
	return o.ApplyT(func(v GetAlarmAllNamespacesResult) []GetAlarmAllNamespacesCustomNamespacesNews {
		return v.CustomNamespacesNews
	}).(GetAlarmAllNamespacesCustomNamespacesNewsArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAlarmAllNamespacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlarmAllNamespacesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAlarmAllNamespacesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAlarmAllNamespacesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAlarmAllNamespacesResultOutput) Module() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlarmAllNamespacesResult) string { return v.Module }).(pulumi.StringOutput)
}

func (o GetAlarmAllNamespacesResultOutput) MonitorTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAlarmAllNamespacesResult) []string { return v.MonitorTypes }).(pulumi.StringArrayOutput)
}

func (o GetAlarmAllNamespacesResultOutput) QceNamespacesNews() GetAlarmAllNamespacesQceNamespacesNewsArrayOutput {
	return o.ApplyT(func(v GetAlarmAllNamespacesResult) []GetAlarmAllNamespacesQceNamespacesNews {
		return v.QceNamespacesNews
	}).(GetAlarmAllNamespacesQceNamespacesNewsArrayOutput)
}

func (o GetAlarmAllNamespacesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmAllNamespacesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAlarmAllNamespacesResultOutput) SceneType() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlarmAllNamespacesResult) string { return v.SceneType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAlarmAllNamespacesResultOutput{})
}