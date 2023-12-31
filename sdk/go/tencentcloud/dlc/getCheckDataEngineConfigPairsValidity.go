// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dlc

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCheckDataEngineConfigPairsValidity(ctx *pulumi.Context, args *GetCheckDataEngineConfigPairsValidityArgs, opts ...pulumi.InvokeOption) (*GetCheckDataEngineConfigPairsValidityResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCheckDataEngineConfigPairsValidityResult
	err := ctx.Invoke("tencentcloud:Dlc/getCheckDataEngineConfigPairsValidity:getCheckDataEngineConfigPairsValidity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCheckDataEngineConfigPairsValidity.
type GetCheckDataEngineConfigPairsValidityArgs struct {
	ChildImageVersionId   *string                                                     `pulumi:"childImageVersionId"`
	DataEngineConfigPairs []GetCheckDataEngineConfigPairsValidityDataEngineConfigPair `pulumi:"dataEngineConfigPairs"`
	ImageVersionId        *string                                                     `pulumi:"imageVersionId"`
	ResultOutputFile      *string                                                     `pulumi:"resultOutputFile"`
}

// A collection of values returned by getCheckDataEngineConfigPairsValidity.
type GetCheckDataEngineConfigPairsValidityResult struct {
	ChildImageVersionId   *string                                                     `pulumi:"childImageVersionId"`
	DataEngineConfigPairs []GetCheckDataEngineConfigPairsValidityDataEngineConfigPair `pulumi:"dataEngineConfigPairs"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string   `pulumi:"id"`
	ImageVersionId     *string  `pulumi:"imageVersionId"`
	IsAvailable        bool     `pulumi:"isAvailable"`
	ResultOutputFile   *string  `pulumi:"resultOutputFile"`
	UnavailableConfigs []string `pulumi:"unavailableConfigs"`
}

func GetCheckDataEngineConfigPairsValidityOutput(ctx *pulumi.Context, args GetCheckDataEngineConfigPairsValidityOutputArgs, opts ...pulumi.InvokeOption) GetCheckDataEngineConfigPairsValidityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCheckDataEngineConfigPairsValidityResult, error) {
			args := v.(GetCheckDataEngineConfigPairsValidityArgs)
			r, err := GetCheckDataEngineConfigPairsValidity(ctx, &args, opts...)
			var s GetCheckDataEngineConfigPairsValidityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCheckDataEngineConfigPairsValidityResultOutput)
}

// A collection of arguments for invoking getCheckDataEngineConfigPairsValidity.
type GetCheckDataEngineConfigPairsValidityOutputArgs struct {
	ChildImageVersionId   pulumi.StringPtrInput                                               `pulumi:"childImageVersionId"`
	DataEngineConfigPairs GetCheckDataEngineConfigPairsValidityDataEngineConfigPairArrayInput `pulumi:"dataEngineConfigPairs"`
	ImageVersionId        pulumi.StringPtrInput                                               `pulumi:"imageVersionId"`
	ResultOutputFile      pulumi.StringPtrInput                                               `pulumi:"resultOutputFile"`
}

func (GetCheckDataEngineConfigPairsValidityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCheckDataEngineConfigPairsValidityArgs)(nil)).Elem()
}

// A collection of values returned by getCheckDataEngineConfigPairsValidity.
type GetCheckDataEngineConfigPairsValidityResultOutput struct{ *pulumi.OutputState }

func (GetCheckDataEngineConfigPairsValidityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCheckDataEngineConfigPairsValidityResult)(nil)).Elem()
}

func (o GetCheckDataEngineConfigPairsValidityResultOutput) ToGetCheckDataEngineConfigPairsValidityResultOutput() GetCheckDataEngineConfigPairsValidityResultOutput {
	return o
}

func (o GetCheckDataEngineConfigPairsValidityResultOutput) ToGetCheckDataEngineConfigPairsValidityResultOutputWithContext(ctx context.Context) GetCheckDataEngineConfigPairsValidityResultOutput {
	return o
}

func (o GetCheckDataEngineConfigPairsValidityResultOutput) ChildImageVersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCheckDataEngineConfigPairsValidityResult) *string { return v.ChildImageVersionId }).(pulumi.StringPtrOutput)
}

func (o GetCheckDataEngineConfigPairsValidityResultOutput) DataEngineConfigPairs() GetCheckDataEngineConfigPairsValidityDataEngineConfigPairArrayOutput {
	return o.ApplyT(func(v GetCheckDataEngineConfigPairsValidityResult) []GetCheckDataEngineConfigPairsValidityDataEngineConfigPair {
		return v.DataEngineConfigPairs
	}).(GetCheckDataEngineConfigPairsValidityDataEngineConfigPairArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCheckDataEngineConfigPairsValidityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCheckDataEngineConfigPairsValidityResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCheckDataEngineConfigPairsValidityResultOutput) ImageVersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCheckDataEngineConfigPairsValidityResult) *string { return v.ImageVersionId }).(pulumi.StringPtrOutput)
}

func (o GetCheckDataEngineConfigPairsValidityResultOutput) IsAvailable() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCheckDataEngineConfigPairsValidityResult) bool { return v.IsAvailable }).(pulumi.BoolOutput)
}

func (o GetCheckDataEngineConfigPairsValidityResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCheckDataEngineConfigPairsValidityResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetCheckDataEngineConfigPairsValidityResultOutput) UnavailableConfigs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCheckDataEngineConfigPairsValidityResult) []string { return v.UnavailableConfigs }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCheckDataEngineConfigPairsValidityResultOutput{})
}
