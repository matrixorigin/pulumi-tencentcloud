// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package antiddos

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetListListener(ctx *pulumi.Context, args *GetListListenerArgs, opts ...pulumi.InvokeOption) (*GetListListenerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetListListenerResult
	err := ctx.Invoke("tencentcloud:Antiddos/getListListener:getListListener", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getListListener.
type GetListListenerArgs struct {
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getListListener.
type GetListListenerResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                          `pulumi:"id"`
	Layer4Listeners  []GetListListenerLayer4Listener `pulumi:"layer4Listeners"`
	Layer7Listeners  []GetListListenerLayer7Listener `pulumi:"layer7Listeners"`
	ResultOutputFile *string                         `pulumi:"resultOutputFile"`
}

func GetListListenerOutput(ctx *pulumi.Context, args GetListListenerOutputArgs, opts ...pulumi.InvokeOption) GetListListenerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetListListenerResult, error) {
			args := v.(GetListListenerArgs)
			r, err := GetListListener(ctx, &args, opts...)
			var s GetListListenerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetListListenerResultOutput)
}

// A collection of arguments for invoking getListListener.
type GetListListenerOutputArgs struct {
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetListListenerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListListenerArgs)(nil)).Elem()
}

// A collection of values returned by getListListener.
type GetListListenerResultOutput struct{ *pulumi.OutputState }

func (GetListListenerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListListenerResult)(nil)).Elem()
}

func (o GetListListenerResultOutput) ToGetListListenerResultOutput() GetListListenerResultOutput {
	return o
}

func (o GetListListenerResultOutput) ToGetListListenerResultOutputWithContext(ctx context.Context) GetListListenerResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetListListenerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetListListenerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetListListenerResultOutput) Layer4Listeners() GetListListenerLayer4ListenerArrayOutput {
	return o.ApplyT(func(v GetListListenerResult) []GetListListenerLayer4Listener { return v.Layer4Listeners }).(GetListListenerLayer4ListenerArrayOutput)
}

func (o GetListListenerResultOutput) Layer7Listeners() GetListListenerLayer7ListenerArrayOutput {
	return o.ApplyT(func(v GetListListenerResult) []GetListListenerLayer7Listener { return v.Layer7Listeners }).(GetListListenerLayer7ListenerArrayOutput)
}

func (o GetListListenerResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListListenerResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetListListenerResultOutput{})
}
