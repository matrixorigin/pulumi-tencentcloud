// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUsersResult
	err := ctx.Invoke("tencentcloud:Cam/getUsers:getUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsers.
type GetUsersArgs struct {
	ConsoleLogin     *bool   `pulumi:"consoleLogin"`
	CountryCode      *string `pulumi:"countryCode"`
	Email            *string `pulumi:"email"`
	Name             *string `pulumi:"name"`
	PhoneNum         *string `pulumi:"phoneNum"`
	Remark           *string `pulumi:"remark"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	Uid              *int    `pulumi:"uid"`
	Uin              *int    `pulumi:"uin"`
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	ConsoleLogin *bool   `pulumi:"consoleLogin"`
	CountryCode  *string `pulumi:"countryCode"`
	Email        *string `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id               string             `pulumi:"id"`
	Name             *string            `pulumi:"name"`
	PhoneNum         *string            `pulumi:"phoneNum"`
	Remark           *string            `pulumi:"remark"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
	Uid              *int               `pulumi:"uid"`
	Uin              *int               `pulumi:"uin"`
	UserLists        []GetUsersUserList `pulumi:"userLists"`
}

func GetUsersOutput(ctx *pulumi.Context, args GetUsersOutputArgs, opts ...pulumi.InvokeOption) GetUsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUsersResult, error) {
			args := v.(GetUsersArgs)
			r, err := GetUsers(ctx, &args, opts...)
			var s GetUsersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUsersResultOutput)
}

// A collection of arguments for invoking getUsers.
type GetUsersOutputArgs struct {
	ConsoleLogin     pulumi.BoolPtrInput   `pulumi:"consoleLogin"`
	CountryCode      pulumi.StringPtrInput `pulumi:"countryCode"`
	Email            pulumi.StringPtrInput `pulumi:"email"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	PhoneNum         pulumi.StringPtrInput `pulumi:"phoneNum"`
	Remark           pulumi.StringPtrInput `pulumi:"remark"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	Uid              pulumi.IntPtrInput    `pulumi:"uid"`
	Uin              pulumi.IntPtrInput    `pulumi:"uin"`
}

func (GetUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersArgs)(nil)).Elem()
}

// A collection of values returned by getUsers.
type GetUsersResultOutput struct{ *pulumi.OutputState }

func (GetUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersResult)(nil)).Elem()
}

func (o GetUsersResultOutput) ToGetUsersResultOutput() GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ToGetUsersResultOutputWithContext(ctx context.Context) GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ConsoleLogin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *bool { return v.ConsoleLogin }).(pulumi.BoolPtrOutput)
}

func (o GetUsersResultOutput) CountryCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.CountryCode }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Email }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUsersResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) PhoneNum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.PhoneNum }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Remark }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) Uid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *int { return v.Uid }).(pulumi.IntPtrOutput)
}

func (o GetUsersResultOutput) Uin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *int { return v.Uin }).(pulumi.IntPtrOutput)
}

func (o GetUsersResultOutput) UserLists() GetUsersUserListArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []GetUsersUserList { return v.UserLists }).(GetUsersUserListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUsersResultOutput{})
}
