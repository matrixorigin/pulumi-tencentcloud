// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bi

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type DatasourceCloudServiceType struct {
	InstanceId string `pulumi:"instanceId"`
	Region     string `pulumi:"region"`
	Type       string `pulumi:"type"`
}

// DatasourceCloudServiceTypeInput is an input type that accepts DatasourceCloudServiceTypeArgs and DatasourceCloudServiceTypeOutput values.
// You can construct a concrete instance of `DatasourceCloudServiceTypeInput` via:
//
//	DatasourceCloudServiceTypeArgs{...}
type DatasourceCloudServiceTypeInput interface {
	pulumi.Input

	ToDatasourceCloudServiceTypeOutput() DatasourceCloudServiceTypeOutput
	ToDatasourceCloudServiceTypeOutputWithContext(context.Context) DatasourceCloudServiceTypeOutput
}

type DatasourceCloudServiceTypeArgs struct {
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	Region     pulumi.StringInput `pulumi:"region"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (DatasourceCloudServiceTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasourceCloudServiceType)(nil)).Elem()
}

func (i DatasourceCloudServiceTypeArgs) ToDatasourceCloudServiceTypeOutput() DatasourceCloudServiceTypeOutput {
	return i.ToDatasourceCloudServiceTypeOutputWithContext(context.Background())
}

func (i DatasourceCloudServiceTypeArgs) ToDatasourceCloudServiceTypeOutputWithContext(ctx context.Context) DatasourceCloudServiceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasourceCloudServiceTypeOutput)
}

func (i DatasourceCloudServiceTypeArgs) ToDatasourceCloudServiceTypePtrOutput() DatasourceCloudServiceTypePtrOutput {
	return i.ToDatasourceCloudServiceTypePtrOutputWithContext(context.Background())
}

func (i DatasourceCloudServiceTypeArgs) ToDatasourceCloudServiceTypePtrOutputWithContext(ctx context.Context) DatasourceCloudServiceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasourceCloudServiceTypeOutput).ToDatasourceCloudServiceTypePtrOutputWithContext(ctx)
}

// DatasourceCloudServiceTypePtrInput is an input type that accepts DatasourceCloudServiceTypeArgs, DatasourceCloudServiceTypePtr and DatasourceCloudServiceTypePtrOutput values.
// You can construct a concrete instance of `DatasourceCloudServiceTypePtrInput` via:
//
//	        DatasourceCloudServiceTypeArgs{...}
//
//	or:
//
//	        nil
type DatasourceCloudServiceTypePtrInput interface {
	pulumi.Input

	ToDatasourceCloudServiceTypePtrOutput() DatasourceCloudServiceTypePtrOutput
	ToDatasourceCloudServiceTypePtrOutputWithContext(context.Context) DatasourceCloudServiceTypePtrOutput
}

type datasourceCloudServiceTypePtrType DatasourceCloudServiceTypeArgs

func DatasourceCloudServiceTypePtr(v *DatasourceCloudServiceTypeArgs) DatasourceCloudServiceTypePtrInput {
	return (*datasourceCloudServiceTypePtrType)(v)
}

func (*datasourceCloudServiceTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasourceCloudServiceType)(nil)).Elem()
}

func (i *datasourceCloudServiceTypePtrType) ToDatasourceCloudServiceTypePtrOutput() DatasourceCloudServiceTypePtrOutput {
	return i.ToDatasourceCloudServiceTypePtrOutputWithContext(context.Background())
}

func (i *datasourceCloudServiceTypePtrType) ToDatasourceCloudServiceTypePtrOutputWithContext(ctx context.Context) DatasourceCloudServiceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasourceCloudServiceTypePtrOutput)
}

type DatasourceCloudServiceTypeOutput struct{ *pulumi.OutputState }

func (DatasourceCloudServiceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasourceCloudServiceType)(nil)).Elem()
}

func (o DatasourceCloudServiceTypeOutput) ToDatasourceCloudServiceTypeOutput() DatasourceCloudServiceTypeOutput {
	return o
}

func (o DatasourceCloudServiceTypeOutput) ToDatasourceCloudServiceTypeOutputWithContext(ctx context.Context) DatasourceCloudServiceTypeOutput {
	return o
}

func (o DatasourceCloudServiceTypeOutput) ToDatasourceCloudServiceTypePtrOutput() DatasourceCloudServiceTypePtrOutput {
	return o.ToDatasourceCloudServiceTypePtrOutputWithContext(context.Background())
}

func (o DatasourceCloudServiceTypeOutput) ToDatasourceCloudServiceTypePtrOutputWithContext(ctx context.Context) DatasourceCloudServiceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatasourceCloudServiceType) *DatasourceCloudServiceType {
		return &v
	}).(DatasourceCloudServiceTypePtrOutput)
}

func (o DatasourceCloudServiceTypeOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v DatasourceCloudServiceType) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o DatasourceCloudServiceTypeOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v DatasourceCloudServiceType) string { return v.Region }).(pulumi.StringOutput)
}

func (o DatasourceCloudServiceTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DatasourceCloudServiceType) string { return v.Type }).(pulumi.StringOutput)
}

type DatasourceCloudServiceTypePtrOutput struct{ *pulumi.OutputState }

func (DatasourceCloudServiceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasourceCloudServiceType)(nil)).Elem()
}

func (o DatasourceCloudServiceTypePtrOutput) ToDatasourceCloudServiceTypePtrOutput() DatasourceCloudServiceTypePtrOutput {
	return o
}

func (o DatasourceCloudServiceTypePtrOutput) ToDatasourceCloudServiceTypePtrOutputWithContext(ctx context.Context) DatasourceCloudServiceTypePtrOutput {
	return o
}

func (o DatasourceCloudServiceTypePtrOutput) Elem() DatasourceCloudServiceTypeOutput {
	return o.ApplyT(func(v *DatasourceCloudServiceType) DatasourceCloudServiceType {
		if v != nil {
			return *v
		}
		var ret DatasourceCloudServiceType
		return ret
	}).(DatasourceCloudServiceTypeOutput)
}

func (o DatasourceCloudServiceTypePtrOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceCloudServiceType) *string {
		if v == nil {
			return nil
		}
		return &v.InstanceId
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceCloudServiceTypePtrOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceCloudServiceType) *string {
		if v == nil {
			return nil
		}
		return &v.Region
	}).(pulumi.StringPtrOutput)
}

func (o DatasourceCloudServiceTypePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasourceCloudServiceType) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type GetProjectList struct {
	Apply            bool                       `pulumi:"apply"`
	AuthLists        []string                   `pulumi:"authLists"`
	ColorCode        string                     `pulumi:"colorCode"`
	ConfigLists      []GetProjectListConfigList `pulumi:"configLists"`
	CorpId           string                     `pulumi:"corpId"`
	CreatedAt        string                     `pulumi:"createdAt"`
	CreatedUser      string                     `pulumi:"createdUser"`
	Id               int                        `pulumi:"id"`
	IsExternalManage bool                       `pulumi:"isExternalManage"`
	LastModifyName   string                     `pulumi:"lastModifyName"`
	Logo             string                     `pulumi:"logo"`
	ManagePlatform   string                     `pulumi:"managePlatform"`
	Mark             string                     `pulumi:"mark"`
	MemberCount      int                        `pulumi:"memberCount"`
	Name             string                     `pulumi:"name"`
	PageCount        int                        `pulumi:"pageCount"`
	PanelScope       string                     `pulumi:"panelScope"`
	Seed             string                     `pulumi:"seed"`
	Source           string                     `pulumi:"source"`
	UpdatedAt        string                     `pulumi:"updatedAt"`
	UpdatedUser      string                     `pulumi:"updatedUser"`
}

// GetProjectListInput is an input type that accepts GetProjectListArgs and GetProjectListOutput values.
// You can construct a concrete instance of `GetProjectListInput` via:
//
//	GetProjectListArgs{...}
type GetProjectListInput interface {
	pulumi.Input

	ToGetProjectListOutput() GetProjectListOutput
	ToGetProjectListOutputWithContext(context.Context) GetProjectListOutput
}

type GetProjectListArgs struct {
	Apply            pulumi.BoolInput                   `pulumi:"apply"`
	AuthLists        pulumi.StringArrayInput            `pulumi:"authLists"`
	ColorCode        pulumi.StringInput                 `pulumi:"colorCode"`
	ConfigLists      GetProjectListConfigListArrayInput `pulumi:"configLists"`
	CorpId           pulumi.StringInput                 `pulumi:"corpId"`
	CreatedAt        pulumi.StringInput                 `pulumi:"createdAt"`
	CreatedUser      pulumi.StringInput                 `pulumi:"createdUser"`
	Id               pulumi.IntInput                    `pulumi:"id"`
	IsExternalManage pulumi.BoolInput                   `pulumi:"isExternalManage"`
	LastModifyName   pulumi.StringInput                 `pulumi:"lastModifyName"`
	Logo             pulumi.StringInput                 `pulumi:"logo"`
	ManagePlatform   pulumi.StringInput                 `pulumi:"managePlatform"`
	Mark             pulumi.StringInput                 `pulumi:"mark"`
	MemberCount      pulumi.IntInput                    `pulumi:"memberCount"`
	Name             pulumi.StringInput                 `pulumi:"name"`
	PageCount        pulumi.IntInput                    `pulumi:"pageCount"`
	PanelScope       pulumi.StringInput                 `pulumi:"panelScope"`
	Seed             pulumi.StringInput                 `pulumi:"seed"`
	Source           pulumi.StringInput                 `pulumi:"source"`
	UpdatedAt        pulumi.StringInput                 `pulumi:"updatedAt"`
	UpdatedUser      pulumi.StringInput                 `pulumi:"updatedUser"`
}

func (GetProjectListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectList)(nil)).Elem()
}

func (i GetProjectListArgs) ToGetProjectListOutput() GetProjectListOutput {
	return i.ToGetProjectListOutputWithContext(context.Background())
}

func (i GetProjectListArgs) ToGetProjectListOutputWithContext(ctx context.Context) GetProjectListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectListOutput)
}

// GetProjectListArrayInput is an input type that accepts GetProjectListArray and GetProjectListArrayOutput values.
// You can construct a concrete instance of `GetProjectListArrayInput` via:
//
//	GetProjectListArray{ GetProjectListArgs{...} }
type GetProjectListArrayInput interface {
	pulumi.Input

	ToGetProjectListArrayOutput() GetProjectListArrayOutput
	ToGetProjectListArrayOutputWithContext(context.Context) GetProjectListArrayOutput
}

type GetProjectListArray []GetProjectListInput

func (GetProjectListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProjectList)(nil)).Elem()
}

func (i GetProjectListArray) ToGetProjectListArrayOutput() GetProjectListArrayOutput {
	return i.ToGetProjectListArrayOutputWithContext(context.Background())
}

func (i GetProjectListArray) ToGetProjectListArrayOutputWithContext(ctx context.Context) GetProjectListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectListArrayOutput)
}

type GetProjectListOutput struct{ *pulumi.OutputState }

func (GetProjectListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectList)(nil)).Elem()
}

func (o GetProjectListOutput) ToGetProjectListOutput() GetProjectListOutput {
	return o
}

func (o GetProjectListOutput) ToGetProjectListOutputWithContext(ctx context.Context) GetProjectListOutput {
	return o
}

func (o GetProjectListOutput) Apply() pulumi.BoolOutput {
	return o.ApplyT(func(v GetProjectList) bool { return v.Apply }).(pulumi.BoolOutput)
}

func (o GetProjectListOutput) AuthLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProjectList) []string { return v.AuthLists }).(pulumi.StringArrayOutput)
}

func (o GetProjectListOutput) ColorCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.ColorCode }).(pulumi.StringOutput)
}

func (o GetProjectListOutput) ConfigLists() GetProjectListConfigListArrayOutput {
	return o.ApplyT(func(v GetProjectList) []GetProjectListConfigList { return v.ConfigLists }).(GetProjectListConfigListArrayOutput)
}

func (o GetProjectListOutput) CorpId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.CorpId }).(pulumi.StringOutput)
}

func (o GetProjectListOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o GetProjectListOutput) CreatedUser() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.CreatedUser }).(pulumi.StringOutput)
}

func (o GetProjectListOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetProjectList) int { return v.Id }).(pulumi.IntOutput)
}

func (o GetProjectListOutput) IsExternalManage() pulumi.BoolOutput {
	return o.ApplyT(func(v GetProjectList) bool { return v.IsExternalManage }).(pulumi.BoolOutput)
}

func (o GetProjectListOutput) LastModifyName() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.LastModifyName }).(pulumi.StringOutput)
}

func (o GetProjectListOutput) Logo() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.Logo }).(pulumi.StringOutput)
}

func (o GetProjectListOutput) ManagePlatform() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.ManagePlatform }).(pulumi.StringOutput)
}

func (o GetProjectListOutput) Mark() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.Mark }).(pulumi.StringOutput)
}

func (o GetProjectListOutput) MemberCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetProjectList) int { return v.MemberCount }).(pulumi.IntOutput)
}

func (o GetProjectListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetProjectListOutput) PageCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetProjectList) int { return v.PageCount }).(pulumi.IntOutput)
}

func (o GetProjectListOutput) PanelScope() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.PanelScope }).(pulumi.StringOutput)
}

func (o GetProjectListOutput) Seed() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.Seed }).(pulumi.StringOutput)
}

func (o GetProjectListOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.Source }).(pulumi.StringOutput)
}

func (o GetProjectListOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o GetProjectListOutput) UpdatedUser() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectList) string { return v.UpdatedUser }).(pulumi.StringOutput)
}

type GetProjectListArrayOutput struct{ *pulumi.OutputState }

func (GetProjectListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProjectList)(nil)).Elem()
}

func (o GetProjectListArrayOutput) ToGetProjectListArrayOutput() GetProjectListArrayOutput {
	return o
}

func (o GetProjectListArrayOutput) ToGetProjectListArrayOutputWithContext(ctx context.Context) GetProjectListArrayOutput {
	return o
}

func (o GetProjectListArrayOutput) Index(i pulumi.IntInput) GetProjectListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProjectList {
		return vs[0].([]GetProjectList)[vs[1].(int)]
	}).(GetProjectListOutput)
}

type GetProjectListConfigList struct {
	Components  []GetProjectListConfigListComponent `pulumi:"components"`
	ModuleGroup string                              `pulumi:"moduleGroup"`
}

// GetProjectListConfigListInput is an input type that accepts GetProjectListConfigListArgs and GetProjectListConfigListOutput values.
// You can construct a concrete instance of `GetProjectListConfigListInput` via:
//
//	GetProjectListConfigListArgs{...}
type GetProjectListConfigListInput interface {
	pulumi.Input

	ToGetProjectListConfigListOutput() GetProjectListConfigListOutput
	ToGetProjectListConfigListOutputWithContext(context.Context) GetProjectListConfigListOutput
}

type GetProjectListConfigListArgs struct {
	Components  GetProjectListConfigListComponentArrayInput `pulumi:"components"`
	ModuleGroup pulumi.StringInput                          `pulumi:"moduleGroup"`
}

func (GetProjectListConfigListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectListConfigList)(nil)).Elem()
}

func (i GetProjectListConfigListArgs) ToGetProjectListConfigListOutput() GetProjectListConfigListOutput {
	return i.ToGetProjectListConfigListOutputWithContext(context.Background())
}

func (i GetProjectListConfigListArgs) ToGetProjectListConfigListOutputWithContext(ctx context.Context) GetProjectListConfigListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectListConfigListOutput)
}

// GetProjectListConfigListArrayInput is an input type that accepts GetProjectListConfigListArray and GetProjectListConfigListArrayOutput values.
// You can construct a concrete instance of `GetProjectListConfigListArrayInput` via:
//
//	GetProjectListConfigListArray{ GetProjectListConfigListArgs{...} }
type GetProjectListConfigListArrayInput interface {
	pulumi.Input

	ToGetProjectListConfigListArrayOutput() GetProjectListConfigListArrayOutput
	ToGetProjectListConfigListArrayOutputWithContext(context.Context) GetProjectListConfigListArrayOutput
}

type GetProjectListConfigListArray []GetProjectListConfigListInput

func (GetProjectListConfigListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProjectListConfigList)(nil)).Elem()
}

func (i GetProjectListConfigListArray) ToGetProjectListConfigListArrayOutput() GetProjectListConfigListArrayOutput {
	return i.ToGetProjectListConfigListArrayOutputWithContext(context.Background())
}

func (i GetProjectListConfigListArray) ToGetProjectListConfigListArrayOutputWithContext(ctx context.Context) GetProjectListConfigListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectListConfigListArrayOutput)
}

type GetProjectListConfigListOutput struct{ *pulumi.OutputState }

func (GetProjectListConfigListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectListConfigList)(nil)).Elem()
}

func (o GetProjectListConfigListOutput) ToGetProjectListConfigListOutput() GetProjectListConfigListOutput {
	return o
}

func (o GetProjectListConfigListOutput) ToGetProjectListConfigListOutputWithContext(ctx context.Context) GetProjectListConfigListOutput {
	return o
}

func (o GetProjectListConfigListOutput) Components() GetProjectListConfigListComponentArrayOutput {
	return o.ApplyT(func(v GetProjectListConfigList) []GetProjectListConfigListComponent { return v.Components }).(GetProjectListConfigListComponentArrayOutput)
}

func (o GetProjectListConfigListOutput) ModuleGroup() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectListConfigList) string { return v.ModuleGroup }).(pulumi.StringOutput)
}

type GetProjectListConfigListArrayOutput struct{ *pulumi.OutputState }

func (GetProjectListConfigListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProjectListConfigList)(nil)).Elem()
}

func (o GetProjectListConfigListArrayOutput) ToGetProjectListConfigListArrayOutput() GetProjectListConfigListArrayOutput {
	return o
}

func (o GetProjectListConfigListArrayOutput) ToGetProjectListConfigListArrayOutputWithContext(ctx context.Context) GetProjectListConfigListArrayOutput {
	return o
}

func (o GetProjectListConfigListArrayOutput) Index(i pulumi.IntInput) GetProjectListConfigListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProjectListConfigList {
		return vs[0].([]GetProjectListConfigList)[vs[1].(int)]
	}).(GetProjectListConfigListOutput)
}

type GetProjectListConfigListComponent struct {
	IncludeType string `pulumi:"includeType"`
	ModuleId    string `pulumi:"moduleId"`
	Params      string `pulumi:"params"`
}

// GetProjectListConfigListComponentInput is an input type that accepts GetProjectListConfigListComponentArgs and GetProjectListConfigListComponentOutput values.
// You can construct a concrete instance of `GetProjectListConfigListComponentInput` via:
//
//	GetProjectListConfigListComponentArgs{...}
type GetProjectListConfigListComponentInput interface {
	pulumi.Input

	ToGetProjectListConfigListComponentOutput() GetProjectListConfigListComponentOutput
	ToGetProjectListConfigListComponentOutputWithContext(context.Context) GetProjectListConfigListComponentOutput
}

type GetProjectListConfigListComponentArgs struct {
	IncludeType pulumi.StringInput `pulumi:"includeType"`
	ModuleId    pulumi.StringInput `pulumi:"moduleId"`
	Params      pulumi.StringInput `pulumi:"params"`
}

func (GetProjectListConfigListComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectListConfigListComponent)(nil)).Elem()
}

func (i GetProjectListConfigListComponentArgs) ToGetProjectListConfigListComponentOutput() GetProjectListConfigListComponentOutput {
	return i.ToGetProjectListConfigListComponentOutputWithContext(context.Background())
}

func (i GetProjectListConfigListComponentArgs) ToGetProjectListConfigListComponentOutputWithContext(ctx context.Context) GetProjectListConfigListComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectListConfigListComponentOutput)
}

// GetProjectListConfigListComponentArrayInput is an input type that accepts GetProjectListConfigListComponentArray and GetProjectListConfigListComponentArrayOutput values.
// You can construct a concrete instance of `GetProjectListConfigListComponentArrayInput` via:
//
//	GetProjectListConfigListComponentArray{ GetProjectListConfigListComponentArgs{...} }
type GetProjectListConfigListComponentArrayInput interface {
	pulumi.Input

	ToGetProjectListConfigListComponentArrayOutput() GetProjectListConfigListComponentArrayOutput
	ToGetProjectListConfigListComponentArrayOutputWithContext(context.Context) GetProjectListConfigListComponentArrayOutput
}

type GetProjectListConfigListComponentArray []GetProjectListConfigListComponentInput

func (GetProjectListConfigListComponentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProjectListConfigListComponent)(nil)).Elem()
}

func (i GetProjectListConfigListComponentArray) ToGetProjectListConfigListComponentArrayOutput() GetProjectListConfigListComponentArrayOutput {
	return i.ToGetProjectListConfigListComponentArrayOutputWithContext(context.Background())
}

func (i GetProjectListConfigListComponentArray) ToGetProjectListConfigListComponentArrayOutputWithContext(ctx context.Context) GetProjectListConfigListComponentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectListConfigListComponentArrayOutput)
}

type GetProjectListConfigListComponentOutput struct{ *pulumi.OutputState }

func (GetProjectListConfigListComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectListConfigListComponent)(nil)).Elem()
}

func (o GetProjectListConfigListComponentOutput) ToGetProjectListConfigListComponentOutput() GetProjectListConfigListComponentOutput {
	return o
}

func (o GetProjectListConfigListComponentOutput) ToGetProjectListConfigListComponentOutputWithContext(ctx context.Context) GetProjectListConfigListComponentOutput {
	return o
}

func (o GetProjectListConfigListComponentOutput) IncludeType() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectListConfigListComponent) string { return v.IncludeType }).(pulumi.StringOutput)
}

func (o GetProjectListConfigListComponentOutput) ModuleId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectListConfigListComponent) string { return v.ModuleId }).(pulumi.StringOutput)
}

func (o GetProjectListConfigListComponentOutput) Params() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectListConfigListComponent) string { return v.Params }).(pulumi.StringOutput)
}

type GetProjectListConfigListComponentArrayOutput struct{ *pulumi.OutputState }

func (GetProjectListConfigListComponentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProjectListConfigListComponent)(nil)).Elem()
}

func (o GetProjectListConfigListComponentArrayOutput) ToGetProjectListConfigListComponentArrayOutput() GetProjectListConfigListComponentArrayOutput {
	return o
}

func (o GetProjectListConfigListComponentArrayOutput) ToGetProjectListConfigListComponentArrayOutputWithContext(ctx context.Context) GetProjectListConfigListComponentArrayOutput {
	return o
}

func (o GetProjectListConfigListComponentArrayOutput) Index(i pulumi.IntInput) GetProjectListConfigListComponentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProjectListConfigListComponent {
		return vs[0].([]GetProjectListConfigListComponent)[vs[1].(int)]
	}).(GetProjectListConfigListComponentOutput)
}

type GetUserProjectList struct {
	AreaCode       string `pulumi:"areaCode"`
	CorpId         string `pulumi:"corpId"`
	CreatedAt      string `pulumi:"createdAt"`
	CreatedUser    string `pulumi:"createdUser"`
	Email          string `pulumi:"email"`
	FirstModify    int    `pulumi:"firstModify"`
	GlobalUserName string `pulumi:"globalUserName"`
	LastLogin      string `pulumi:"lastLogin"`
	Mobile         string `pulumi:"mobile"`
	PhoneNumber    string `pulumi:"phoneNumber"`
	Status         int    `pulumi:"status"`
	UpdatedAt      string `pulumi:"updatedAt"`
	UpdatedUser    string `pulumi:"updatedUser"`
	UserId         string `pulumi:"userId"`
	UserName       string `pulumi:"userName"`
}

// GetUserProjectListInput is an input type that accepts GetUserProjectListArgs and GetUserProjectListOutput values.
// You can construct a concrete instance of `GetUserProjectListInput` via:
//
//	GetUserProjectListArgs{...}
type GetUserProjectListInput interface {
	pulumi.Input

	ToGetUserProjectListOutput() GetUserProjectListOutput
	ToGetUserProjectListOutputWithContext(context.Context) GetUserProjectListOutput
}

type GetUserProjectListArgs struct {
	AreaCode       pulumi.StringInput `pulumi:"areaCode"`
	CorpId         pulumi.StringInput `pulumi:"corpId"`
	CreatedAt      pulumi.StringInput `pulumi:"createdAt"`
	CreatedUser    pulumi.StringInput `pulumi:"createdUser"`
	Email          pulumi.StringInput `pulumi:"email"`
	FirstModify    pulumi.IntInput    `pulumi:"firstModify"`
	GlobalUserName pulumi.StringInput `pulumi:"globalUserName"`
	LastLogin      pulumi.StringInput `pulumi:"lastLogin"`
	Mobile         pulumi.StringInput `pulumi:"mobile"`
	PhoneNumber    pulumi.StringInput `pulumi:"phoneNumber"`
	Status         pulumi.IntInput    `pulumi:"status"`
	UpdatedAt      pulumi.StringInput `pulumi:"updatedAt"`
	UpdatedUser    pulumi.StringInput `pulumi:"updatedUser"`
	UserId         pulumi.StringInput `pulumi:"userId"`
	UserName       pulumi.StringInput `pulumi:"userName"`
}

func (GetUserProjectListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserProjectList)(nil)).Elem()
}

func (i GetUserProjectListArgs) ToGetUserProjectListOutput() GetUserProjectListOutput {
	return i.ToGetUserProjectListOutputWithContext(context.Background())
}

func (i GetUserProjectListArgs) ToGetUserProjectListOutputWithContext(ctx context.Context) GetUserProjectListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetUserProjectListOutput)
}

// GetUserProjectListArrayInput is an input type that accepts GetUserProjectListArray and GetUserProjectListArrayOutput values.
// You can construct a concrete instance of `GetUserProjectListArrayInput` via:
//
//	GetUserProjectListArray{ GetUserProjectListArgs{...} }
type GetUserProjectListArrayInput interface {
	pulumi.Input

	ToGetUserProjectListArrayOutput() GetUserProjectListArrayOutput
	ToGetUserProjectListArrayOutputWithContext(context.Context) GetUserProjectListArrayOutput
}

type GetUserProjectListArray []GetUserProjectListInput

func (GetUserProjectListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetUserProjectList)(nil)).Elem()
}

func (i GetUserProjectListArray) ToGetUserProjectListArrayOutput() GetUserProjectListArrayOutput {
	return i.ToGetUserProjectListArrayOutputWithContext(context.Background())
}

func (i GetUserProjectListArray) ToGetUserProjectListArrayOutputWithContext(ctx context.Context) GetUserProjectListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetUserProjectListArrayOutput)
}

type GetUserProjectListOutput struct{ *pulumi.OutputState }

func (GetUserProjectListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserProjectList)(nil)).Elem()
}

func (o GetUserProjectListOutput) ToGetUserProjectListOutput() GetUserProjectListOutput {
	return o
}

func (o GetUserProjectListOutput) ToGetUserProjectListOutputWithContext(ctx context.Context) GetUserProjectListOutput {
	return o
}

func (o GetUserProjectListOutput) AreaCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserProjectList) string { return v.AreaCode }).(pulumi.StringOutput)
}

func (o GetUserProjectListOutput) CorpId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserProjectList) string { return v.CorpId }).(pulumi.StringOutput)
}

func (o GetUserProjectListOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserProjectList) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o GetUserProjectListOutput) CreatedUser() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserProjectList) string { return v.CreatedUser }).(pulumi.StringOutput)
}

func (o GetUserProjectListOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserProjectList) string { return v.Email }).(pulumi.StringOutput)
}

func (o GetUserProjectListOutput) FirstModify() pulumi.IntOutput {
	return o.ApplyT(func(v GetUserProjectList) int { return v.FirstModify }).(pulumi.IntOutput)
}

func (o GetUserProjectListOutput) GlobalUserName() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserProjectList) string { return v.GlobalUserName }).(pulumi.StringOutput)
}

func (o GetUserProjectListOutput) LastLogin() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserProjectList) string { return v.LastLogin }).(pulumi.StringOutput)
}

func (o GetUserProjectListOutput) Mobile() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserProjectList) string { return v.Mobile }).(pulumi.StringOutput)
}

func (o GetUserProjectListOutput) PhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserProjectList) string { return v.PhoneNumber }).(pulumi.StringOutput)
}

func (o GetUserProjectListOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v GetUserProjectList) int { return v.Status }).(pulumi.IntOutput)
}

func (o GetUserProjectListOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserProjectList) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o GetUserProjectListOutput) UpdatedUser() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserProjectList) string { return v.UpdatedUser }).(pulumi.StringOutput)
}

func (o GetUserProjectListOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserProjectList) string { return v.UserId }).(pulumi.StringOutput)
}

func (o GetUserProjectListOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserProjectList) string { return v.UserName }).(pulumi.StringOutput)
}

type GetUserProjectListArrayOutput struct{ *pulumi.OutputState }

func (GetUserProjectListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetUserProjectList)(nil)).Elem()
}

func (o GetUserProjectListArrayOutput) ToGetUserProjectListArrayOutput() GetUserProjectListArrayOutput {
	return o
}

func (o GetUserProjectListArrayOutput) ToGetUserProjectListArrayOutputWithContext(ctx context.Context) GetUserProjectListArrayOutput {
	return o
}

func (o GetUserProjectListArrayOutput) Index(i pulumi.IntInput) GetUserProjectListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetUserProjectList {
		return vs[0].([]GetUserProjectList)[vs[1].(int)]
	}).(GetUserProjectListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasourceCloudServiceTypeInput)(nil)).Elem(), DatasourceCloudServiceTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasourceCloudServiceTypePtrInput)(nil)).Elem(), DatasourceCloudServiceTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectListInput)(nil)).Elem(), GetProjectListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectListArrayInput)(nil)).Elem(), GetProjectListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectListConfigListInput)(nil)).Elem(), GetProjectListConfigListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectListConfigListArrayInput)(nil)).Elem(), GetProjectListConfigListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectListConfigListComponentInput)(nil)).Elem(), GetProjectListConfigListComponentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectListConfigListComponentArrayInput)(nil)).Elem(), GetProjectListConfigListComponentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetUserProjectListInput)(nil)).Elem(), GetUserProjectListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetUserProjectListArrayInput)(nil)).Elem(), GetUserProjectListArray{})
	pulumi.RegisterOutputType(DatasourceCloudServiceTypeOutput{})
	pulumi.RegisterOutputType(DatasourceCloudServiceTypePtrOutput{})
	pulumi.RegisterOutputType(GetProjectListOutput{})
	pulumi.RegisterOutputType(GetProjectListArrayOutput{})
	pulumi.RegisterOutputType(GetProjectListConfigListOutput{})
	pulumi.RegisterOutputType(GetProjectListConfigListArrayOutput{})
	pulumi.RegisterOutputType(GetProjectListConfigListComponentOutput{})
	pulumi.RegisterOutputType(GetProjectListConfigListComponentArrayOutput{})
	pulumi.RegisterOutputType(GetUserProjectListOutput{})
	pulumi.RegisterOutputType(GetUserProjectListArrayOutput{})
}
