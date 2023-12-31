// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRocketmqGroup(ctx *pulumi.Context, args *LookupRocketmqGroupArgs, opts ...pulumi.InvokeOption) (*LookupRocketmqGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRocketmqGroupResult
	err := ctx.Invoke("tencentcloud:Tdmq/getRocketmqGroup:getRocketmqGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRocketmqGroup.
type LookupRocketmqGroupArgs struct {
	ClusterId        string  `pulumi:"clusterId"`
	FilterGroup      *string `pulumi:"filterGroup"`
	FilterOneGroup   *string `pulumi:"filterOneGroup"`
	FilterTopic      *string `pulumi:"filterTopic"`
	NamespaceId      string  `pulumi:"namespaceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getRocketmqGroup.
type LookupRocketmqGroupResult struct {
	ClusterId      string                  `pulumi:"clusterId"`
	FilterGroup    *string                 `pulumi:"filterGroup"`
	FilterOneGroup *string                 `pulumi:"filterOneGroup"`
	FilterTopic    *string                 `pulumi:"filterTopic"`
	Groups         []GetRocketmqGroupGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	NamespaceId      string  `pulumi:"namespaceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func LookupRocketmqGroupOutput(ctx *pulumi.Context, args LookupRocketmqGroupOutputArgs, opts ...pulumi.InvokeOption) LookupRocketmqGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRocketmqGroupResult, error) {
			args := v.(LookupRocketmqGroupArgs)
			r, err := LookupRocketmqGroup(ctx, &args, opts...)
			var s LookupRocketmqGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRocketmqGroupResultOutput)
}

// A collection of arguments for invoking getRocketmqGroup.
type LookupRocketmqGroupOutputArgs struct {
	ClusterId        pulumi.StringInput    `pulumi:"clusterId"`
	FilterGroup      pulumi.StringPtrInput `pulumi:"filterGroup"`
	FilterOneGroup   pulumi.StringPtrInput `pulumi:"filterOneGroup"`
	FilterTopic      pulumi.StringPtrInput `pulumi:"filterTopic"`
	NamespaceId      pulumi.StringInput    `pulumi:"namespaceId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (LookupRocketmqGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRocketmqGroupArgs)(nil)).Elem()
}

// A collection of values returned by getRocketmqGroup.
type LookupRocketmqGroupResultOutput struct{ *pulumi.OutputState }

func (LookupRocketmqGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRocketmqGroupResult)(nil)).Elem()
}

func (o LookupRocketmqGroupResultOutput) ToLookupRocketmqGroupResultOutput() LookupRocketmqGroupResultOutput {
	return o
}

func (o LookupRocketmqGroupResultOutput) ToLookupRocketmqGroupResultOutputWithContext(ctx context.Context) LookupRocketmqGroupResultOutput {
	return o
}

func (o LookupRocketmqGroupResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRocketmqGroupResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupRocketmqGroupResultOutput) FilterGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRocketmqGroupResult) *string { return v.FilterGroup }).(pulumi.StringPtrOutput)
}

func (o LookupRocketmqGroupResultOutput) FilterOneGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRocketmqGroupResult) *string { return v.FilterOneGroup }).(pulumi.StringPtrOutput)
}

func (o LookupRocketmqGroupResultOutput) FilterTopic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRocketmqGroupResult) *string { return v.FilterTopic }).(pulumi.StringPtrOutput)
}

func (o LookupRocketmqGroupResultOutput) Groups() GetRocketmqGroupGroupArrayOutput {
	return o.ApplyT(func(v LookupRocketmqGroupResult) []GetRocketmqGroupGroup { return v.Groups }).(GetRocketmqGroupGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRocketmqGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRocketmqGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRocketmqGroupResultOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRocketmqGroupResult) string { return v.NamespaceId }).(pulumi.StringOutput)
}

func (o LookupRocketmqGroupResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRocketmqGroupResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRocketmqGroupResultOutput{})
}
