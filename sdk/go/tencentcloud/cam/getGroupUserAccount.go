// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGroupUserAccount(ctx *pulumi.Context, args *GetGroupUserAccountArgs, opts ...pulumi.InvokeOption) (*GetGroupUserAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGroupUserAccountResult
	err := ctx.Invoke("tencentcloud:Cam/getGroupUserAccount:getGroupUserAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroupUserAccount.
type GetGroupUserAccountArgs struct {
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	Rp               *int    `pulumi:"rp"`
	SubUin           *int    `pulumi:"subUin"`
	Uid              *int    `pulumi:"uid"`
}

// A collection of values returned by getGroupUserAccount.
type GetGroupUserAccountResult struct {
	GroupInfos []GetGroupUserAccountGroupInfo `pulumi:"groupInfos"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	Rp               *int    `pulumi:"rp"`
	SubUin           *int    `pulumi:"subUin"`
	TotalNum         int     `pulumi:"totalNum"`
	Uid              *int    `pulumi:"uid"`
}

func GetGroupUserAccountOutput(ctx *pulumi.Context, args GetGroupUserAccountOutputArgs, opts ...pulumi.InvokeOption) GetGroupUserAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupUserAccountResult, error) {
			args := v.(GetGroupUserAccountArgs)
			r, err := GetGroupUserAccount(ctx, &args, opts...)
			var s GetGroupUserAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGroupUserAccountResultOutput)
}

// A collection of arguments for invoking getGroupUserAccount.
type GetGroupUserAccountOutputArgs struct {
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	Rp               pulumi.IntPtrInput    `pulumi:"rp"`
	SubUin           pulumi.IntPtrInput    `pulumi:"subUin"`
	Uid              pulumi.IntPtrInput    `pulumi:"uid"`
}

func (GetGroupUserAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupUserAccountArgs)(nil)).Elem()
}

// A collection of values returned by getGroupUserAccount.
type GetGroupUserAccountResultOutput struct{ *pulumi.OutputState }

func (GetGroupUserAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupUserAccountResult)(nil)).Elem()
}

func (o GetGroupUserAccountResultOutput) ToGetGroupUserAccountResultOutput() GetGroupUserAccountResultOutput {
	return o
}

func (o GetGroupUserAccountResultOutput) ToGetGroupUserAccountResultOutputWithContext(ctx context.Context) GetGroupUserAccountResultOutput {
	return o
}

func (o GetGroupUserAccountResultOutput) GroupInfos() GetGroupUserAccountGroupInfoArrayOutput {
	return o.ApplyT(func(v GetGroupUserAccountResult) []GetGroupUserAccountGroupInfo { return v.GroupInfos }).(GetGroupUserAccountGroupInfoArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupUserAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupUserAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGroupUserAccountResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupUserAccountResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetGroupUserAccountResultOutput) Rp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetGroupUserAccountResult) *int { return v.Rp }).(pulumi.IntPtrOutput)
}

func (o GetGroupUserAccountResultOutput) SubUin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetGroupUserAccountResult) *int { return v.SubUin }).(pulumi.IntPtrOutput)
}

func (o GetGroupUserAccountResultOutput) TotalNum() pulumi.IntOutput {
	return o.ApplyT(func(v GetGroupUserAccountResult) int { return v.TotalNum }).(pulumi.IntOutput)
}

func (o GetGroupUserAccountResultOutput) Uid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetGroupUserAccountResult) *int { return v.Uid }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupUserAccountResultOutput{})
}