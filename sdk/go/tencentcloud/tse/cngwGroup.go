// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tse

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CngwGroup struct {
	pulumi.CustomResourceState

	// description information of group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// gateway IDonly postpaid gateway supported.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// internet configration.
	InternetConfig CngwGroupInternetConfigPtrOutput `pulumi:"internetConfig"`
	// public network outbound traffic bandwidth,[1,2048]Mbps.
	InternetMaxBandwidthOut pulumi.IntPtrOutput `pulumi:"internetMaxBandwidthOut"`
	// gateway group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// group node configration.
	NodeConfig CngwGroupNodeConfigOutput `pulumi:"nodeConfig"`
	// subnet ID. Assign an IP address to the engine in the VPC subnet. Reference value:- subnet-ahde9me9.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewCngwGroup registers a new resource with the given unique name, arguments, and options.
func NewCngwGroup(ctx *pulumi.Context,
	name string, args *CngwGroupArgs, opts ...pulumi.ResourceOption) (*CngwGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.NodeConfig == nil {
		return nil, errors.New("invalid value for required argument 'NodeConfig'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CngwGroup
	err := ctx.RegisterResource("tencentcloud:Tse/cngwGroup:CngwGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCngwGroup gets an existing CngwGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCngwGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CngwGroupState, opts ...pulumi.ResourceOption) (*CngwGroup, error) {
	var resource CngwGroup
	err := ctx.ReadResource("tencentcloud:Tse/cngwGroup:CngwGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CngwGroup resources.
type cngwGroupState struct {
	// description information of group.
	Description *string `pulumi:"description"`
	// gateway IDonly postpaid gateway supported.
	GatewayId *string `pulumi:"gatewayId"`
	// internet configration.
	InternetConfig *CngwGroupInternetConfig `pulumi:"internetConfig"`
	// public network outbound traffic bandwidth,[1,2048]Mbps.
	InternetMaxBandwidthOut *int `pulumi:"internetMaxBandwidthOut"`
	// gateway group name.
	Name *string `pulumi:"name"`
	// group node configration.
	NodeConfig *CngwGroupNodeConfig `pulumi:"nodeConfig"`
	// subnet ID. Assign an IP address to the engine in the VPC subnet. Reference value:- subnet-ahde9me9.
	SubnetId *string `pulumi:"subnetId"`
}

type CngwGroupState struct {
	// description information of group.
	Description pulumi.StringPtrInput
	// gateway IDonly postpaid gateway supported.
	GatewayId pulumi.StringPtrInput
	// internet configration.
	InternetConfig CngwGroupInternetConfigPtrInput
	// public network outbound traffic bandwidth,[1,2048]Mbps.
	InternetMaxBandwidthOut pulumi.IntPtrInput
	// gateway group name.
	Name pulumi.StringPtrInput
	// group node configration.
	NodeConfig CngwGroupNodeConfigPtrInput
	// subnet ID. Assign an IP address to the engine in the VPC subnet. Reference value:- subnet-ahde9me9.
	SubnetId pulumi.StringPtrInput
}

func (CngwGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*cngwGroupState)(nil)).Elem()
}

type cngwGroupArgs struct {
	// description information of group.
	Description *string `pulumi:"description"`
	// gateway IDonly postpaid gateway supported.
	GatewayId string `pulumi:"gatewayId"`
	// internet configration.
	InternetConfig *CngwGroupInternetConfig `pulumi:"internetConfig"`
	// public network outbound traffic bandwidth,[1,2048]Mbps.
	InternetMaxBandwidthOut *int `pulumi:"internetMaxBandwidthOut"`
	// gateway group name.
	Name *string `pulumi:"name"`
	// group node configration.
	NodeConfig CngwGroupNodeConfig `pulumi:"nodeConfig"`
	// subnet ID. Assign an IP address to the engine in the VPC subnet. Reference value:- subnet-ahde9me9.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a CngwGroup resource.
type CngwGroupArgs struct {
	// description information of group.
	Description pulumi.StringPtrInput
	// gateway IDonly postpaid gateway supported.
	GatewayId pulumi.StringInput
	// internet configration.
	InternetConfig CngwGroupInternetConfigPtrInput
	// public network outbound traffic bandwidth,[1,2048]Mbps.
	InternetMaxBandwidthOut pulumi.IntPtrInput
	// gateway group name.
	Name pulumi.StringPtrInput
	// group node configration.
	NodeConfig CngwGroupNodeConfigInput
	// subnet ID. Assign an IP address to the engine in the VPC subnet. Reference value:- subnet-ahde9me9.
	SubnetId pulumi.StringInput
}

func (CngwGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cngwGroupArgs)(nil)).Elem()
}

type CngwGroupInput interface {
	pulumi.Input

	ToCngwGroupOutput() CngwGroupOutput
	ToCngwGroupOutputWithContext(ctx context.Context) CngwGroupOutput
}

func (*CngwGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**CngwGroup)(nil)).Elem()
}

func (i *CngwGroup) ToCngwGroupOutput() CngwGroupOutput {
	return i.ToCngwGroupOutputWithContext(context.Background())
}

func (i *CngwGroup) ToCngwGroupOutputWithContext(ctx context.Context) CngwGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CngwGroupOutput)
}

// CngwGroupArrayInput is an input type that accepts CngwGroupArray and CngwGroupArrayOutput values.
// You can construct a concrete instance of `CngwGroupArrayInput` via:
//
//	CngwGroupArray{ CngwGroupArgs{...} }
type CngwGroupArrayInput interface {
	pulumi.Input

	ToCngwGroupArrayOutput() CngwGroupArrayOutput
	ToCngwGroupArrayOutputWithContext(context.Context) CngwGroupArrayOutput
}

type CngwGroupArray []CngwGroupInput

func (CngwGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CngwGroup)(nil)).Elem()
}

func (i CngwGroupArray) ToCngwGroupArrayOutput() CngwGroupArrayOutput {
	return i.ToCngwGroupArrayOutputWithContext(context.Background())
}

func (i CngwGroupArray) ToCngwGroupArrayOutputWithContext(ctx context.Context) CngwGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CngwGroupArrayOutput)
}

// CngwGroupMapInput is an input type that accepts CngwGroupMap and CngwGroupMapOutput values.
// You can construct a concrete instance of `CngwGroupMapInput` via:
//
//	CngwGroupMap{ "key": CngwGroupArgs{...} }
type CngwGroupMapInput interface {
	pulumi.Input

	ToCngwGroupMapOutput() CngwGroupMapOutput
	ToCngwGroupMapOutputWithContext(context.Context) CngwGroupMapOutput
}

type CngwGroupMap map[string]CngwGroupInput

func (CngwGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CngwGroup)(nil)).Elem()
}

func (i CngwGroupMap) ToCngwGroupMapOutput() CngwGroupMapOutput {
	return i.ToCngwGroupMapOutputWithContext(context.Background())
}

func (i CngwGroupMap) ToCngwGroupMapOutputWithContext(ctx context.Context) CngwGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CngwGroupMapOutput)
}

type CngwGroupOutput struct{ *pulumi.OutputState }

func (CngwGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CngwGroup)(nil)).Elem()
}

func (o CngwGroupOutput) ToCngwGroupOutput() CngwGroupOutput {
	return o
}

func (o CngwGroupOutput) ToCngwGroupOutputWithContext(ctx context.Context) CngwGroupOutput {
	return o
}

// description information of group.
func (o CngwGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CngwGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// gateway IDonly postpaid gateway supported.
func (o CngwGroupOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *CngwGroup) pulumi.StringOutput { return v.GatewayId }).(pulumi.StringOutput)
}

// internet configration.
func (o CngwGroupOutput) InternetConfig() CngwGroupInternetConfigPtrOutput {
	return o.ApplyT(func(v *CngwGroup) CngwGroupInternetConfigPtrOutput { return v.InternetConfig }).(CngwGroupInternetConfigPtrOutput)
}

// public network outbound traffic bandwidth,[1,2048]Mbps.
func (o CngwGroupOutput) InternetMaxBandwidthOut() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CngwGroup) pulumi.IntPtrOutput { return v.InternetMaxBandwidthOut }).(pulumi.IntPtrOutput)
}

// gateway group name.
func (o CngwGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CngwGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// group node configration.
func (o CngwGroupOutput) NodeConfig() CngwGroupNodeConfigOutput {
	return o.ApplyT(func(v *CngwGroup) CngwGroupNodeConfigOutput { return v.NodeConfig }).(CngwGroupNodeConfigOutput)
}

// subnet ID. Assign an IP address to the engine in the VPC subnet. Reference value:- subnet-ahde9me9.
func (o CngwGroupOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *CngwGroup) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

type CngwGroupArrayOutput struct{ *pulumi.OutputState }

func (CngwGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CngwGroup)(nil)).Elem()
}

func (o CngwGroupArrayOutput) ToCngwGroupArrayOutput() CngwGroupArrayOutput {
	return o
}

func (o CngwGroupArrayOutput) ToCngwGroupArrayOutputWithContext(ctx context.Context) CngwGroupArrayOutput {
	return o
}

func (o CngwGroupArrayOutput) Index(i pulumi.IntInput) CngwGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CngwGroup {
		return vs[0].([]*CngwGroup)[vs[1].(int)]
	}).(CngwGroupOutput)
}

type CngwGroupMapOutput struct{ *pulumi.OutputState }

func (CngwGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CngwGroup)(nil)).Elem()
}

func (o CngwGroupMapOutput) ToCngwGroupMapOutput() CngwGroupMapOutput {
	return o
}

func (o CngwGroupMapOutput) ToCngwGroupMapOutputWithContext(ctx context.Context) CngwGroupMapOutput {
	return o
}

func (o CngwGroupMapOutput) MapIndex(k pulumi.StringInput) CngwGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CngwGroup {
		return vs[0].(map[string]*CngwGroup)[vs[1].(string)]
	}).(CngwGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CngwGroupInput)(nil)).Elem(), &CngwGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*CngwGroupArrayInput)(nil)).Elem(), CngwGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CngwGroupMapInput)(nil)).Elem(), CngwGroupMap{})
	pulumi.RegisterOutputType(CngwGroupOutput{})
	pulumi.RegisterOutputType(CngwGroupArrayOutput{})
	pulumi.RegisterOutputType(CngwGroupMapOutput{})
}