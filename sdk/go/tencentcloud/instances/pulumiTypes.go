// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package instances

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type GetInstanceInstanceList struct {
	AllocatePublicIp                   bool                              `pulumi:"allocatePublicIp"`
	AvailabilityZone                   string                            `pulumi:"availabilityZone"`
	CamRoleName                        string                            `pulumi:"camRoleName"`
	Cpu                                int                               `pulumi:"cpu"`
	CreateTime                         string                            `pulumi:"createTime"`
	DataDisks                          []GetInstanceInstanceListDataDisk `pulumi:"dataDisks"`
	ExpiredTime                        string                            `pulumi:"expiredTime"`
	ImageId                            string                            `pulumi:"imageId"`
	InstanceChargeType                 string                            `pulumi:"instanceChargeType"`
	InstanceChargeTypePrepaidRenewFlag string                            `pulumi:"instanceChargeTypePrepaidRenewFlag"`
	InstanceId                         string                            `pulumi:"instanceId"`
	InstanceName                       string                            `pulumi:"instanceName"`
	InstanceType                       string                            `pulumi:"instanceType"`
	InternetChargeType                 string                            `pulumi:"internetChargeType"`
	InternetMaxBandwidthOut            int                               `pulumi:"internetMaxBandwidthOut"`
	Memory                             int                               `pulumi:"memory"`
	PrivateIp                          string                            `pulumi:"privateIp"`
	ProjectId                          int                               `pulumi:"projectId"`
	PublicIp                           string                            `pulumi:"publicIp"`
	SecurityGroups                     []string                          `pulumi:"securityGroups"`
	Status                             string                            `pulumi:"status"`
	SubnetId                           string                            `pulumi:"subnetId"`
	SystemDiskId                       string                            `pulumi:"systemDiskId"`
	SystemDiskSize                     int                               `pulumi:"systemDiskSize"`
	SystemDiskType                     string                            `pulumi:"systemDiskType"`
	Tags                               map[string]interface{}            `pulumi:"tags"`
	VpcId                              string                            `pulumi:"vpcId"`
}

// GetInstanceInstanceListInput is an input type that accepts GetInstanceInstanceListArgs and GetInstanceInstanceListOutput values.
// You can construct a concrete instance of `GetInstanceInstanceListInput` via:
//
//	GetInstanceInstanceListArgs{...}
type GetInstanceInstanceListInput interface {
	pulumi.Input

	ToGetInstanceInstanceListOutput() GetInstanceInstanceListOutput
	ToGetInstanceInstanceListOutputWithContext(context.Context) GetInstanceInstanceListOutput
}

type GetInstanceInstanceListArgs struct {
	AllocatePublicIp                   pulumi.BoolInput                          `pulumi:"allocatePublicIp"`
	AvailabilityZone                   pulumi.StringInput                        `pulumi:"availabilityZone"`
	CamRoleName                        pulumi.StringInput                        `pulumi:"camRoleName"`
	Cpu                                pulumi.IntInput                           `pulumi:"cpu"`
	CreateTime                         pulumi.StringInput                        `pulumi:"createTime"`
	DataDisks                          GetInstanceInstanceListDataDiskArrayInput `pulumi:"dataDisks"`
	ExpiredTime                        pulumi.StringInput                        `pulumi:"expiredTime"`
	ImageId                            pulumi.StringInput                        `pulumi:"imageId"`
	InstanceChargeType                 pulumi.StringInput                        `pulumi:"instanceChargeType"`
	InstanceChargeTypePrepaidRenewFlag pulumi.StringInput                        `pulumi:"instanceChargeTypePrepaidRenewFlag"`
	InstanceId                         pulumi.StringInput                        `pulumi:"instanceId"`
	InstanceName                       pulumi.StringInput                        `pulumi:"instanceName"`
	InstanceType                       pulumi.StringInput                        `pulumi:"instanceType"`
	InternetChargeType                 pulumi.StringInput                        `pulumi:"internetChargeType"`
	InternetMaxBandwidthOut            pulumi.IntInput                           `pulumi:"internetMaxBandwidthOut"`
	Memory                             pulumi.IntInput                           `pulumi:"memory"`
	PrivateIp                          pulumi.StringInput                        `pulumi:"privateIp"`
	ProjectId                          pulumi.IntInput                           `pulumi:"projectId"`
	PublicIp                           pulumi.StringInput                        `pulumi:"publicIp"`
	SecurityGroups                     pulumi.StringArrayInput                   `pulumi:"securityGroups"`
	Status                             pulumi.StringInput                        `pulumi:"status"`
	SubnetId                           pulumi.StringInput                        `pulumi:"subnetId"`
	SystemDiskId                       pulumi.StringInput                        `pulumi:"systemDiskId"`
	SystemDiskSize                     pulumi.IntInput                           `pulumi:"systemDiskSize"`
	SystemDiskType                     pulumi.StringInput                        `pulumi:"systemDiskType"`
	Tags                               pulumi.MapInput                           `pulumi:"tags"`
	VpcId                              pulumi.StringInput                        `pulumi:"vpcId"`
}

func (GetInstanceInstanceListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceInstanceList)(nil)).Elem()
}

func (i GetInstanceInstanceListArgs) ToGetInstanceInstanceListOutput() GetInstanceInstanceListOutput {
	return i.ToGetInstanceInstanceListOutputWithContext(context.Background())
}

func (i GetInstanceInstanceListArgs) ToGetInstanceInstanceListOutputWithContext(ctx context.Context) GetInstanceInstanceListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceInstanceListOutput)
}

// GetInstanceInstanceListArrayInput is an input type that accepts GetInstanceInstanceListArray and GetInstanceInstanceListArrayOutput values.
// You can construct a concrete instance of `GetInstanceInstanceListArrayInput` via:
//
//	GetInstanceInstanceListArray{ GetInstanceInstanceListArgs{...} }
type GetInstanceInstanceListArrayInput interface {
	pulumi.Input

	ToGetInstanceInstanceListArrayOutput() GetInstanceInstanceListArrayOutput
	ToGetInstanceInstanceListArrayOutputWithContext(context.Context) GetInstanceInstanceListArrayOutput
}

type GetInstanceInstanceListArray []GetInstanceInstanceListInput

func (GetInstanceInstanceListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceInstanceList)(nil)).Elem()
}

func (i GetInstanceInstanceListArray) ToGetInstanceInstanceListArrayOutput() GetInstanceInstanceListArrayOutput {
	return i.ToGetInstanceInstanceListArrayOutputWithContext(context.Background())
}

func (i GetInstanceInstanceListArray) ToGetInstanceInstanceListArrayOutputWithContext(ctx context.Context) GetInstanceInstanceListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceInstanceListArrayOutput)
}

type GetInstanceInstanceListOutput struct{ *pulumi.OutputState }

func (GetInstanceInstanceListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceInstanceList)(nil)).Elem()
}

func (o GetInstanceInstanceListOutput) ToGetInstanceInstanceListOutput() GetInstanceInstanceListOutput {
	return o
}

func (o GetInstanceInstanceListOutput) ToGetInstanceInstanceListOutputWithContext(ctx context.Context) GetInstanceInstanceListOutput {
	return o
}

func (o GetInstanceInstanceListOutput) AllocatePublicIp() pulumi.BoolOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) bool { return v.AllocatePublicIp }).(pulumi.BoolOutput)
}

func (o GetInstanceInstanceListOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) CamRoleName() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.CamRoleName }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) Cpu() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) int { return v.Cpu }).(pulumi.IntOutput)
}

func (o GetInstanceInstanceListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) DataDisks() GetInstanceInstanceListDataDiskArrayOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) []GetInstanceInstanceListDataDisk { return v.DataDisks }).(GetInstanceInstanceListDataDiskArrayOutput)
}

func (o GetInstanceInstanceListOutput) ExpiredTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.ExpiredTime }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.ImageId }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) InstanceChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.InstanceChargeType }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) InstanceChargeTypePrepaidRenewFlag() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.InstanceChargeTypePrepaidRenewFlag }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) InternetChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.InternetChargeType }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) InternetMaxBandwidthOut() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) int { return v.InternetMaxBandwidthOut }).(pulumi.IntOutput)
}

func (o GetInstanceInstanceListOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) int { return v.Memory }).(pulumi.IntOutput)
}

func (o GetInstanceInstanceListOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.PrivateIp }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) int { return v.ProjectId }).(pulumi.IntOutput)
}

func (o GetInstanceInstanceListOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.PublicIp }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) []string { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

func (o GetInstanceInstanceListOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.Status }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) SystemDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.SystemDiskId }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) SystemDiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) int { return v.SystemDiskSize }).(pulumi.IntOutput)
}

func (o GetInstanceInstanceListOutput) SystemDiskType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.SystemDiskType }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetInstanceInstanceListOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceList) string { return v.VpcId }).(pulumi.StringOutput)
}

type GetInstanceInstanceListArrayOutput struct{ *pulumi.OutputState }

func (GetInstanceInstanceListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceInstanceList)(nil)).Elem()
}

func (o GetInstanceInstanceListArrayOutput) ToGetInstanceInstanceListArrayOutput() GetInstanceInstanceListArrayOutput {
	return o
}

func (o GetInstanceInstanceListArrayOutput) ToGetInstanceInstanceListArrayOutputWithContext(ctx context.Context) GetInstanceInstanceListArrayOutput {
	return o
}

func (o GetInstanceInstanceListArrayOutput) Index(i pulumi.IntInput) GetInstanceInstanceListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetInstanceInstanceList {
		return vs[0].([]GetInstanceInstanceList)[vs[1].(int)]
	}).(GetInstanceInstanceListOutput)
}

type GetInstanceInstanceListDataDisk struct {
	DataDiskId         string `pulumi:"dataDiskId"`
	DataDiskSize       int    `pulumi:"dataDiskSize"`
	DataDiskType       string `pulumi:"dataDiskType"`
	DeleteWithInstance bool   `pulumi:"deleteWithInstance"`
}

// GetInstanceInstanceListDataDiskInput is an input type that accepts GetInstanceInstanceListDataDiskArgs and GetInstanceInstanceListDataDiskOutput values.
// You can construct a concrete instance of `GetInstanceInstanceListDataDiskInput` via:
//
//	GetInstanceInstanceListDataDiskArgs{...}
type GetInstanceInstanceListDataDiskInput interface {
	pulumi.Input

	ToGetInstanceInstanceListDataDiskOutput() GetInstanceInstanceListDataDiskOutput
	ToGetInstanceInstanceListDataDiskOutputWithContext(context.Context) GetInstanceInstanceListDataDiskOutput
}

type GetInstanceInstanceListDataDiskArgs struct {
	DataDiskId         pulumi.StringInput `pulumi:"dataDiskId"`
	DataDiskSize       pulumi.IntInput    `pulumi:"dataDiskSize"`
	DataDiskType       pulumi.StringInput `pulumi:"dataDiskType"`
	DeleteWithInstance pulumi.BoolInput   `pulumi:"deleteWithInstance"`
}

func (GetInstanceInstanceListDataDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceInstanceListDataDisk)(nil)).Elem()
}

func (i GetInstanceInstanceListDataDiskArgs) ToGetInstanceInstanceListDataDiskOutput() GetInstanceInstanceListDataDiskOutput {
	return i.ToGetInstanceInstanceListDataDiskOutputWithContext(context.Background())
}

func (i GetInstanceInstanceListDataDiskArgs) ToGetInstanceInstanceListDataDiskOutputWithContext(ctx context.Context) GetInstanceInstanceListDataDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceInstanceListDataDiskOutput)
}

// GetInstanceInstanceListDataDiskArrayInput is an input type that accepts GetInstanceInstanceListDataDiskArray and GetInstanceInstanceListDataDiskArrayOutput values.
// You can construct a concrete instance of `GetInstanceInstanceListDataDiskArrayInput` via:
//
//	GetInstanceInstanceListDataDiskArray{ GetInstanceInstanceListDataDiskArgs{...} }
type GetInstanceInstanceListDataDiskArrayInput interface {
	pulumi.Input

	ToGetInstanceInstanceListDataDiskArrayOutput() GetInstanceInstanceListDataDiskArrayOutput
	ToGetInstanceInstanceListDataDiskArrayOutputWithContext(context.Context) GetInstanceInstanceListDataDiskArrayOutput
}

type GetInstanceInstanceListDataDiskArray []GetInstanceInstanceListDataDiskInput

func (GetInstanceInstanceListDataDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceInstanceListDataDisk)(nil)).Elem()
}

func (i GetInstanceInstanceListDataDiskArray) ToGetInstanceInstanceListDataDiskArrayOutput() GetInstanceInstanceListDataDiskArrayOutput {
	return i.ToGetInstanceInstanceListDataDiskArrayOutputWithContext(context.Background())
}

func (i GetInstanceInstanceListDataDiskArray) ToGetInstanceInstanceListDataDiskArrayOutputWithContext(ctx context.Context) GetInstanceInstanceListDataDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstanceInstanceListDataDiskArrayOutput)
}

type GetInstanceInstanceListDataDiskOutput struct{ *pulumi.OutputState }

func (GetInstanceInstanceListDataDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceInstanceListDataDisk)(nil)).Elem()
}

func (o GetInstanceInstanceListDataDiskOutput) ToGetInstanceInstanceListDataDiskOutput() GetInstanceInstanceListDataDiskOutput {
	return o
}

func (o GetInstanceInstanceListDataDiskOutput) ToGetInstanceInstanceListDataDiskOutputWithContext(ctx context.Context) GetInstanceInstanceListDataDiskOutput {
	return o
}

func (o GetInstanceInstanceListDataDiskOutput) DataDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceListDataDisk) string { return v.DataDiskId }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListDataDiskOutput) DataDiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstanceInstanceListDataDisk) int { return v.DataDiskSize }).(pulumi.IntOutput)
}

func (o GetInstanceInstanceListDataDiskOutput) DataDiskType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceInstanceListDataDisk) string { return v.DataDiskType }).(pulumi.StringOutput)
}

func (o GetInstanceInstanceListDataDiskOutput) DeleteWithInstance() pulumi.BoolOutput {
	return o.ApplyT(func(v GetInstanceInstanceListDataDisk) bool { return v.DeleteWithInstance }).(pulumi.BoolOutput)
}

type GetInstanceInstanceListDataDiskArrayOutput struct{ *pulumi.OutputState }

func (GetInstanceInstanceListDataDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstanceInstanceListDataDisk)(nil)).Elem()
}

func (o GetInstanceInstanceListDataDiskArrayOutput) ToGetInstanceInstanceListDataDiskArrayOutput() GetInstanceInstanceListDataDiskArrayOutput {
	return o
}

func (o GetInstanceInstanceListDataDiskArrayOutput) ToGetInstanceInstanceListDataDiskArrayOutputWithContext(ctx context.Context) GetInstanceInstanceListDataDiskArrayOutput {
	return o
}

func (o GetInstanceInstanceListDataDiskArrayOutput) Index(i pulumi.IntInput) GetInstanceInstanceListDataDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetInstanceInstanceListDataDisk {
		return vs[0].([]GetInstanceInstanceListDataDisk)[vs[1].(int)]
	}).(GetInstanceInstanceListDataDiskOutput)
}

type GetSetInstanceList struct {
	AllocatePublicIp                   bool                         `pulumi:"allocatePublicIp"`
	AvailabilityZone                   string                       `pulumi:"availabilityZone"`
	CamRoleName                        string                       `pulumi:"camRoleName"`
	Cpu                                int                          `pulumi:"cpu"`
	CreateTime                         string                       `pulumi:"createTime"`
	DataDisks                          []GetSetInstanceListDataDisk `pulumi:"dataDisks"`
	ExpiredTime                        string                       `pulumi:"expiredTime"`
	ImageId                            string                       `pulumi:"imageId"`
	InstanceChargeType                 string                       `pulumi:"instanceChargeType"`
	InstanceChargeTypePrepaidRenewFlag string                       `pulumi:"instanceChargeTypePrepaidRenewFlag"`
	InstanceId                         string                       `pulumi:"instanceId"`
	InstanceName                       string                       `pulumi:"instanceName"`
	InstanceType                       string                       `pulumi:"instanceType"`
	InternetChargeType                 string                       `pulumi:"internetChargeType"`
	InternetMaxBandwidthOut            int                          `pulumi:"internetMaxBandwidthOut"`
	Memory                             int                          `pulumi:"memory"`
	PrivateIp                          string                       `pulumi:"privateIp"`
	ProjectId                          int                          `pulumi:"projectId"`
	PublicIp                           string                       `pulumi:"publicIp"`
	SecurityGroups                     []string                     `pulumi:"securityGroups"`
	Status                             string                       `pulumi:"status"`
	SubnetId                           string                       `pulumi:"subnetId"`
	SystemDiskId                       string                       `pulumi:"systemDiskId"`
	SystemDiskSize                     int                          `pulumi:"systemDiskSize"`
	SystemDiskType                     string                       `pulumi:"systemDiskType"`
	Tags                               map[string]interface{}       `pulumi:"tags"`
	VpcId                              string                       `pulumi:"vpcId"`
}

// GetSetInstanceListInput is an input type that accepts GetSetInstanceListArgs and GetSetInstanceListOutput values.
// You can construct a concrete instance of `GetSetInstanceListInput` via:
//
//	GetSetInstanceListArgs{...}
type GetSetInstanceListInput interface {
	pulumi.Input

	ToGetSetInstanceListOutput() GetSetInstanceListOutput
	ToGetSetInstanceListOutputWithContext(context.Context) GetSetInstanceListOutput
}

type GetSetInstanceListArgs struct {
	AllocatePublicIp                   pulumi.BoolInput                     `pulumi:"allocatePublicIp"`
	AvailabilityZone                   pulumi.StringInput                   `pulumi:"availabilityZone"`
	CamRoleName                        pulumi.StringInput                   `pulumi:"camRoleName"`
	Cpu                                pulumi.IntInput                      `pulumi:"cpu"`
	CreateTime                         pulumi.StringInput                   `pulumi:"createTime"`
	DataDisks                          GetSetInstanceListDataDiskArrayInput `pulumi:"dataDisks"`
	ExpiredTime                        pulumi.StringInput                   `pulumi:"expiredTime"`
	ImageId                            pulumi.StringInput                   `pulumi:"imageId"`
	InstanceChargeType                 pulumi.StringInput                   `pulumi:"instanceChargeType"`
	InstanceChargeTypePrepaidRenewFlag pulumi.StringInput                   `pulumi:"instanceChargeTypePrepaidRenewFlag"`
	InstanceId                         pulumi.StringInput                   `pulumi:"instanceId"`
	InstanceName                       pulumi.StringInput                   `pulumi:"instanceName"`
	InstanceType                       pulumi.StringInput                   `pulumi:"instanceType"`
	InternetChargeType                 pulumi.StringInput                   `pulumi:"internetChargeType"`
	InternetMaxBandwidthOut            pulumi.IntInput                      `pulumi:"internetMaxBandwidthOut"`
	Memory                             pulumi.IntInput                      `pulumi:"memory"`
	PrivateIp                          pulumi.StringInput                   `pulumi:"privateIp"`
	ProjectId                          pulumi.IntInput                      `pulumi:"projectId"`
	PublicIp                           pulumi.StringInput                   `pulumi:"publicIp"`
	SecurityGroups                     pulumi.StringArrayInput              `pulumi:"securityGroups"`
	Status                             pulumi.StringInput                   `pulumi:"status"`
	SubnetId                           pulumi.StringInput                   `pulumi:"subnetId"`
	SystemDiskId                       pulumi.StringInput                   `pulumi:"systemDiskId"`
	SystemDiskSize                     pulumi.IntInput                      `pulumi:"systemDiskSize"`
	SystemDiskType                     pulumi.StringInput                   `pulumi:"systemDiskType"`
	Tags                               pulumi.MapInput                      `pulumi:"tags"`
	VpcId                              pulumi.StringInput                   `pulumi:"vpcId"`
}

func (GetSetInstanceListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSetInstanceList)(nil)).Elem()
}

func (i GetSetInstanceListArgs) ToGetSetInstanceListOutput() GetSetInstanceListOutput {
	return i.ToGetSetInstanceListOutputWithContext(context.Background())
}

func (i GetSetInstanceListArgs) ToGetSetInstanceListOutputWithContext(ctx context.Context) GetSetInstanceListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSetInstanceListOutput)
}

// GetSetInstanceListArrayInput is an input type that accepts GetSetInstanceListArray and GetSetInstanceListArrayOutput values.
// You can construct a concrete instance of `GetSetInstanceListArrayInput` via:
//
//	GetSetInstanceListArray{ GetSetInstanceListArgs{...} }
type GetSetInstanceListArrayInput interface {
	pulumi.Input

	ToGetSetInstanceListArrayOutput() GetSetInstanceListArrayOutput
	ToGetSetInstanceListArrayOutputWithContext(context.Context) GetSetInstanceListArrayOutput
}

type GetSetInstanceListArray []GetSetInstanceListInput

func (GetSetInstanceListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSetInstanceList)(nil)).Elem()
}

func (i GetSetInstanceListArray) ToGetSetInstanceListArrayOutput() GetSetInstanceListArrayOutput {
	return i.ToGetSetInstanceListArrayOutputWithContext(context.Background())
}

func (i GetSetInstanceListArray) ToGetSetInstanceListArrayOutputWithContext(ctx context.Context) GetSetInstanceListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSetInstanceListArrayOutput)
}

type GetSetInstanceListOutput struct{ *pulumi.OutputState }

func (GetSetInstanceListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSetInstanceList)(nil)).Elem()
}

func (o GetSetInstanceListOutput) ToGetSetInstanceListOutput() GetSetInstanceListOutput {
	return o
}

func (o GetSetInstanceListOutput) ToGetSetInstanceListOutputWithContext(ctx context.Context) GetSetInstanceListOutput {
	return o
}

func (o GetSetInstanceListOutput) AllocatePublicIp() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSetInstanceList) bool { return v.AllocatePublicIp }).(pulumi.BoolOutput)
}

func (o GetSetInstanceListOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) CamRoleName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.CamRoleName }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) Cpu() pulumi.IntOutput {
	return o.ApplyT(func(v GetSetInstanceList) int { return v.Cpu }).(pulumi.IntOutput)
}

func (o GetSetInstanceListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) DataDisks() GetSetInstanceListDataDiskArrayOutput {
	return o.ApplyT(func(v GetSetInstanceList) []GetSetInstanceListDataDisk { return v.DataDisks }).(GetSetInstanceListDataDiskArrayOutput)
}

func (o GetSetInstanceListOutput) ExpiredTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.ExpiredTime }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.ImageId }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) InstanceChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.InstanceChargeType }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) InstanceChargeTypePrepaidRenewFlag() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.InstanceChargeTypePrepaidRenewFlag }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) InternetChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.InternetChargeType }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) InternetMaxBandwidthOut() pulumi.IntOutput {
	return o.ApplyT(func(v GetSetInstanceList) int { return v.InternetMaxBandwidthOut }).(pulumi.IntOutput)
}

func (o GetSetInstanceListOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v GetSetInstanceList) int { return v.Memory }).(pulumi.IntOutput)
}

func (o GetSetInstanceListOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.PrivateIp }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v GetSetInstanceList) int { return v.ProjectId }).(pulumi.IntOutput)
}

func (o GetSetInstanceListOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.PublicIp }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSetInstanceList) []string { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

func (o GetSetInstanceListOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.Status }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) SystemDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.SystemDiskId }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) SystemDiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetSetInstanceList) int { return v.SystemDiskSize }).(pulumi.IntOutput)
}

func (o GetSetInstanceListOutput) SystemDiskType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.SystemDiskType }).(pulumi.StringOutput)
}

func (o GetSetInstanceListOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetSetInstanceList) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetSetInstanceListOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceList) string { return v.VpcId }).(pulumi.StringOutput)
}

type GetSetInstanceListArrayOutput struct{ *pulumi.OutputState }

func (GetSetInstanceListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSetInstanceList)(nil)).Elem()
}

func (o GetSetInstanceListArrayOutput) ToGetSetInstanceListArrayOutput() GetSetInstanceListArrayOutput {
	return o
}

func (o GetSetInstanceListArrayOutput) ToGetSetInstanceListArrayOutputWithContext(ctx context.Context) GetSetInstanceListArrayOutput {
	return o
}

func (o GetSetInstanceListArrayOutput) Index(i pulumi.IntInput) GetSetInstanceListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSetInstanceList {
		return vs[0].([]GetSetInstanceList)[vs[1].(int)]
	}).(GetSetInstanceListOutput)
}

type GetSetInstanceListDataDisk struct {
	DataDiskId         string `pulumi:"dataDiskId"`
	DataDiskSize       int    `pulumi:"dataDiskSize"`
	DataDiskType       string `pulumi:"dataDiskType"`
	DeleteWithInstance bool   `pulumi:"deleteWithInstance"`
}

// GetSetInstanceListDataDiskInput is an input type that accepts GetSetInstanceListDataDiskArgs and GetSetInstanceListDataDiskOutput values.
// You can construct a concrete instance of `GetSetInstanceListDataDiskInput` via:
//
//	GetSetInstanceListDataDiskArgs{...}
type GetSetInstanceListDataDiskInput interface {
	pulumi.Input

	ToGetSetInstanceListDataDiskOutput() GetSetInstanceListDataDiskOutput
	ToGetSetInstanceListDataDiskOutputWithContext(context.Context) GetSetInstanceListDataDiskOutput
}

type GetSetInstanceListDataDiskArgs struct {
	DataDiskId         pulumi.StringInput `pulumi:"dataDiskId"`
	DataDiskSize       pulumi.IntInput    `pulumi:"dataDiskSize"`
	DataDiskType       pulumi.StringInput `pulumi:"dataDiskType"`
	DeleteWithInstance pulumi.BoolInput   `pulumi:"deleteWithInstance"`
}

func (GetSetInstanceListDataDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSetInstanceListDataDisk)(nil)).Elem()
}

func (i GetSetInstanceListDataDiskArgs) ToGetSetInstanceListDataDiskOutput() GetSetInstanceListDataDiskOutput {
	return i.ToGetSetInstanceListDataDiskOutputWithContext(context.Background())
}

func (i GetSetInstanceListDataDiskArgs) ToGetSetInstanceListDataDiskOutputWithContext(ctx context.Context) GetSetInstanceListDataDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSetInstanceListDataDiskOutput)
}

// GetSetInstanceListDataDiskArrayInput is an input type that accepts GetSetInstanceListDataDiskArray and GetSetInstanceListDataDiskArrayOutput values.
// You can construct a concrete instance of `GetSetInstanceListDataDiskArrayInput` via:
//
//	GetSetInstanceListDataDiskArray{ GetSetInstanceListDataDiskArgs{...} }
type GetSetInstanceListDataDiskArrayInput interface {
	pulumi.Input

	ToGetSetInstanceListDataDiskArrayOutput() GetSetInstanceListDataDiskArrayOutput
	ToGetSetInstanceListDataDiskArrayOutputWithContext(context.Context) GetSetInstanceListDataDiskArrayOutput
}

type GetSetInstanceListDataDiskArray []GetSetInstanceListDataDiskInput

func (GetSetInstanceListDataDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSetInstanceListDataDisk)(nil)).Elem()
}

func (i GetSetInstanceListDataDiskArray) ToGetSetInstanceListDataDiskArrayOutput() GetSetInstanceListDataDiskArrayOutput {
	return i.ToGetSetInstanceListDataDiskArrayOutputWithContext(context.Background())
}

func (i GetSetInstanceListDataDiskArray) ToGetSetInstanceListDataDiskArrayOutputWithContext(ctx context.Context) GetSetInstanceListDataDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSetInstanceListDataDiskArrayOutput)
}

type GetSetInstanceListDataDiskOutput struct{ *pulumi.OutputState }

func (GetSetInstanceListDataDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSetInstanceListDataDisk)(nil)).Elem()
}

func (o GetSetInstanceListDataDiskOutput) ToGetSetInstanceListDataDiskOutput() GetSetInstanceListDataDiskOutput {
	return o
}

func (o GetSetInstanceListDataDiskOutput) ToGetSetInstanceListDataDiskOutputWithContext(ctx context.Context) GetSetInstanceListDataDiskOutput {
	return o
}

func (o GetSetInstanceListDataDiskOutput) DataDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceListDataDisk) string { return v.DataDiskId }).(pulumi.StringOutput)
}

func (o GetSetInstanceListDataDiskOutput) DataDiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetSetInstanceListDataDisk) int { return v.DataDiskSize }).(pulumi.IntOutput)
}

func (o GetSetInstanceListDataDiskOutput) DataDiskType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSetInstanceListDataDisk) string { return v.DataDiskType }).(pulumi.StringOutput)
}

func (o GetSetInstanceListDataDiskOutput) DeleteWithInstance() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSetInstanceListDataDisk) bool { return v.DeleteWithInstance }).(pulumi.BoolOutput)
}

type GetSetInstanceListDataDiskArrayOutput struct{ *pulumi.OutputState }

func (GetSetInstanceListDataDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSetInstanceListDataDisk)(nil)).Elem()
}

func (o GetSetInstanceListDataDiskArrayOutput) ToGetSetInstanceListDataDiskArrayOutput() GetSetInstanceListDataDiskArrayOutput {
	return o
}

func (o GetSetInstanceListDataDiskArrayOutput) ToGetSetInstanceListDataDiskArrayOutputWithContext(ctx context.Context) GetSetInstanceListDataDiskArrayOutput {
	return o
}

func (o GetSetInstanceListDataDiskArrayOutput) Index(i pulumi.IntInput) GetSetInstanceListDataDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSetInstanceListDataDisk {
		return vs[0].([]GetSetInstanceListDataDisk)[vs[1].(int)]
	}).(GetSetInstanceListDataDiskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstanceInstanceListInput)(nil)).Elem(), GetInstanceInstanceListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstanceInstanceListArrayInput)(nil)).Elem(), GetInstanceInstanceListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstanceInstanceListDataDiskInput)(nil)).Elem(), GetInstanceInstanceListDataDiskArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstanceInstanceListDataDiskArrayInput)(nil)).Elem(), GetInstanceInstanceListDataDiskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSetInstanceListInput)(nil)).Elem(), GetSetInstanceListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSetInstanceListArrayInput)(nil)).Elem(), GetSetInstanceListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSetInstanceListDataDiskInput)(nil)).Elem(), GetSetInstanceListDataDiskArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSetInstanceListDataDiskArrayInput)(nil)).Elem(), GetSetInstanceListDataDiskArray{})
	pulumi.RegisterOutputType(GetInstanceInstanceListOutput{})
	pulumi.RegisterOutputType(GetInstanceInstanceListArrayOutput{})
	pulumi.RegisterOutputType(GetInstanceInstanceListDataDiskOutput{})
	pulumi.RegisterOutputType(GetInstanceInstanceListDataDiskArrayOutput{})
	pulumi.RegisterOutputType(GetSetInstanceListOutput{})
	pulumi.RegisterOutputType(GetSetInstanceListArrayOutput{})
	pulumi.RegisterOutputType(GetSetInstanceListDataDiskOutput{})
	pulumi.RegisterOutputType(GetSetInstanceListDataDiskArrayOutput{})
}