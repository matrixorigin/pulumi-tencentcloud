// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetLogBackups(ctx *pulumi.Context, args *GetLogBackupsArgs, opts ...pulumi.InvokeOption) (*GetLogBackupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLogBackupsResult
	err := ctx.Invoke("tencentcloud:Postgresql/getLogBackups:getLogBackups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogBackups.
type GetLogBackupsArgs struct {
	Filters          []GetLogBackupsFilter `pulumi:"filters"`
	MaxFinishTime    *string               `pulumi:"maxFinishTime"`
	MinFinishTime    *string               `pulumi:"minFinishTime"`
	OrderBy          *string               `pulumi:"orderBy"`
	OrderByType      *string               `pulumi:"orderByType"`
	ResultOutputFile *string               `pulumi:"resultOutputFile"`
}

// A collection of values returned by getLogBackups.
type GetLogBackupsResult struct {
	Filters []GetLogBackupsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                      `pulumi:"id"`
	LogBackupSets    []GetLogBackupsLogBackupSet `pulumi:"logBackupSets"`
	MaxFinishTime    *string                     `pulumi:"maxFinishTime"`
	MinFinishTime    *string                     `pulumi:"minFinishTime"`
	OrderBy          *string                     `pulumi:"orderBy"`
	OrderByType      *string                     `pulumi:"orderByType"`
	ResultOutputFile *string                     `pulumi:"resultOutputFile"`
}

func GetLogBackupsOutput(ctx *pulumi.Context, args GetLogBackupsOutputArgs, opts ...pulumi.InvokeOption) GetLogBackupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLogBackupsResult, error) {
			args := v.(GetLogBackupsArgs)
			r, err := GetLogBackups(ctx, &args, opts...)
			var s GetLogBackupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLogBackupsResultOutput)
}

// A collection of arguments for invoking getLogBackups.
type GetLogBackupsOutputArgs struct {
	Filters          GetLogBackupsFilterArrayInput `pulumi:"filters"`
	MaxFinishTime    pulumi.StringPtrInput         `pulumi:"maxFinishTime"`
	MinFinishTime    pulumi.StringPtrInput         `pulumi:"minFinishTime"`
	OrderBy          pulumi.StringPtrInput         `pulumi:"orderBy"`
	OrderByType      pulumi.StringPtrInput         `pulumi:"orderByType"`
	ResultOutputFile pulumi.StringPtrInput         `pulumi:"resultOutputFile"`
}

func (GetLogBackupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogBackupsArgs)(nil)).Elem()
}

// A collection of values returned by getLogBackups.
type GetLogBackupsResultOutput struct{ *pulumi.OutputState }

func (GetLogBackupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogBackupsResult)(nil)).Elem()
}

func (o GetLogBackupsResultOutput) ToGetLogBackupsResultOutput() GetLogBackupsResultOutput {
	return o
}

func (o GetLogBackupsResultOutput) ToGetLogBackupsResultOutputWithContext(ctx context.Context) GetLogBackupsResultOutput {
	return o
}

func (o GetLogBackupsResultOutput) Filters() GetLogBackupsFilterArrayOutput {
	return o.ApplyT(func(v GetLogBackupsResult) []GetLogBackupsFilter { return v.Filters }).(GetLogBackupsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLogBackupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogBackupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLogBackupsResultOutput) LogBackupSets() GetLogBackupsLogBackupSetArrayOutput {
	return o.ApplyT(func(v GetLogBackupsResult) []GetLogBackupsLogBackupSet { return v.LogBackupSets }).(GetLogBackupsLogBackupSetArrayOutput)
}

func (o GetLogBackupsResultOutput) MaxFinishTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogBackupsResult) *string { return v.MaxFinishTime }).(pulumi.StringPtrOutput)
}

func (o GetLogBackupsResultOutput) MinFinishTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogBackupsResult) *string { return v.MinFinishTime }).(pulumi.StringPtrOutput)
}

func (o GetLogBackupsResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogBackupsResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetLogBackupsResultOutput) OrderByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogBackupsResult) *string { return v.OrderByType }).(pulumi.StringPtrOutput)
}

func (o GetLogBackupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogBackupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogBackupsResultOutput{})
}
