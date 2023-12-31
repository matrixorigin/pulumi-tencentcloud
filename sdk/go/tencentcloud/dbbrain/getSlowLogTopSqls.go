// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbbrain

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSlowLogTopSqls(ctx *pulumi.Context, args *GetSlowLogTopSqlsArgs, opts ...pulumi.InvokeOption) (*GetSlowLogTopSqlsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSlowLogTopSqlsResult
	err := ctx.Invoke("tencentcloud:Dbbrain/getSlowLogTopSqls:getSlowLogTopSqls", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSlowLogTopSqls.
type GetSlowLogTopSqlsArgs struct {
	EndTime          string                        `pulumi:"endTime"`
	InstanceId       string                        `pulumi:"instanceId"`
	OrderBy          *string                       `pulumi:"orderBy"`
	Product          *string                       `pulumi:"product"`
	ResultOutputFile *string                       `pulumi:"resultOutputFile"`
	SchemaLists      []GetSlowLogTopSqlsSchemaList `pulumi:"schemaLists"`
	SortBy           *string                       `pulumi:"sortBy"`
	StartTime        string                        `pulumi:"startTime"`
}

// A collection of values returned by getSlowLogTopSqls.
type GetSlowLogTopSqlsResult struct {
	EndTime string `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                        `pulumi:"id"`
	InstanceId       string                        `pulumi:"instanceId"`
	OrderBy          *string                       `pulumi:"orderBy"`
	Product          *string                       `pulumi:"product"`
	ResultOutputFile *string                       `pulumi:"resultOutputFile"`
	Rows             []GetSlowLogTopSqlsRow        `pulumi:"rows"`
	SchemaLists      []GetSlowLogTopSqlsSchemaList `pulumi:"schemaLists"`
	SortBy           *string                       `pulumi:"sortBy"`
	StartTime        string                        `pulumi:"startTime"`
}

func GetSlowLogTopSqlsOutput(ctx *pulumi.Context, args GetSlowLogTopSqlsOutputArgs, opts ...pulumi.InvokeOption) GetSlowLogTopSqlsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSlowLogTopSqlsResult, error) {
			args := v.(GetSlowLogTopSqlsArgs)
			r, err := GetSlowLogTopSqls(ctx, &args, opts...)
			var s GetSlowLogTopSqlsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSlowLogTopSqlsResultOutput)
}

// A collection of arguments for invoking getSlowLogTopSqls.
type GetSlowLogTopSqlsOutputArgs struct {
	EndTime          pulumi.StringInput                    `pulumi:"endTime"`
	InstanceId       pulumi.StringInput                    `pulumi:"instanceId"`
	OrderBy          pulumi.StringPtrInput                 `pulumi:"orderBy"`
	Product          pulumi.StringPtrInput                 `pulumi:"product"`
	ResultOutputFile pulumi.StringPtrInput                 `pulumi:"resultOutputFile"`
	SchemaLists      GetSlowLogTopSqlsSchemaListArrayInput `pulumi:"schemaLists"`
	SortBy           pulumi.StringPtrInput                 `pulumi:"sortBy"`
	StartTime        pulumi.StringInput                    `pulumi:"startTime"`
}

func (GetSlowLogTopSqlsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowLogTopSqlsArgs)(nil)).Elem()
}

// A collection of values returned by getSlowLogTopSqls.
type GetSlowLogTopSqlsResultOutput struct{ *pulumi.OutputState }

func (GetSlowLogTopSqlsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowLogTopSqlsResult)(nil)).Elem()
}

func (o GetSlowLogTopSqlsResultOutput) ToGetSlowLogTopSqlsResultOutput() GetSlowLogTopSqlsResultOutput {
	return o
}

func (o GetSlowLogTopSqlsResultOutput) ToGetSlowLogTopSqlsResultOutputWithContext(ctx context.Context) GetSlowLogTopSqlsResultOutput {
	return o
}

func (o GetSlowLogTopSqlsResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSlowLogTopSqlsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSlowLogTopSqlsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetSlowLogTopSqlsResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogTopSqlsResultOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogTopSqlsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogTopSqlsResultOutput) Rows() GetSlowLogTopSqlsRowArrayOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) []GetSlowLogTopSqlsRow { return v.Rows }).(GetSlowLogTopSqlsRowArrayOutput)
}

func (o GetSlowLogTopSqlsResultOutput) SchemaLists() GetSlowLogTopSqlsSchemaListArrayOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) []GetSlowLogTopSqlsSchemaList { return v.SchemaLists }).(GetSlowLogTopSqlsSchemaListArrayOutput)
}

func (o GetSlowLogTopSqlsResultOutput) SortBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) *string { return v.SortBy }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogTopSqlsResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSlowLogTopSqlsResultOutput{})
}
