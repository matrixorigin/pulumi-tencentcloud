// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package css

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SnapshotTemplate struct {
	pulumi.CustomResourceState

	// Cos application ID.
	CosAppId pulumi.IntOutput `pulumi:"cosAppId"`
	// Cos bucket name. Note: The CosBucket parameter value cannot include the - [appid] part.
	CosBucket pulumi.StringOutput `pulumi:"cosBucket"`
	// Cos file name. If it is empty, set according to the default value
	// {StreamID}-screenshot-{Hour}-{Minute}-{Second}-{Width}x{Height}{Ext}.
	CosFileName pulumi.StringPtrOutput `pulumi:"cosFileName"`
	// Cos Bucket folder prefix. If it is empty, set according to the default value /{Year}-{Month}-{Day}/.
	CosPrefix pulumi.StringPtrOutput `pulumi:"cosPrefix"`
	// Cos region.
	CosRegion pulumi.StringOutput `pulumi:"cosRegion"`
	// Description information. Maximum length: 1024 bytes. Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Screenshot height. Default: 0 (original height). Range: 0-2000.
	Height pulumi.IntPtrOutput `pulumi:"height"`
	// Whether porn is enabled, 0: not enabled, 1: enabled. Default: 0.
	PornFlag pulumi.IntPtrOutput `pulumi:"pornFlag"`
	// Screenshot interval, unit: s, default: 10s. Range: 2s~300s.
	SnapshotInterval pulumi.IntPtrOutput `pulumi:"snapshotInterval"`
	// Template name. Maximum length: 255 bytes. Only Chinese, English, numbers, `_`, `-` are supported.
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
	// Screenshot width. Default: 0 (original width). Range: 0-3000.
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewSnapshotTemplate registers a new resource with the given unique name, arguments, and options.
func NewSnapshotTemplate(ctx *pulumi.Context,
	name string, args *SnapshotTemplateArgs, opts ...pulumi.ResourceOption) (*SnapshotTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CosAppId == nil {
		return nil, errors.New("invalid value for required argument 'CosAppId'")
	}
	if args.CosBucket == nil {
		return nil, errors.New("invalid value for required argument 'CosBucket'")
	}
	if args.CosRegion == nil {
		return nil, errors.New("invalid value for required argument 'CosRegion'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SnapshotTemplate
	err := ctx.RegisterResource("tencentcloud:Css/snapshotTemplate:SnapshotTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotTemplate gets an existing SnapshotTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotTemplateState, opts ...pulumi.ResourceOption) (*SnapshotTemplate, error) {
	var resource SnapshotTemplate
	err := ctx.ReadResource("tencentcloud:Css/snapshotTemplate:SnapshotTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotTemplate resources.
type snapshotTemplateState struct {
	// Cos application ID.
	CosAppId *int `pulumi:"cosAppId"`
	// Cos bucket name. Note: The CosBucket parameter value cannot include the - [appid] part.
	CosBucket *string `pulumi:"cosBucket"`
	// Cos file name. If it is empty, set according to the default value
	// {StreamID}-screenshot-{Hour}-{Minute}-{Second}-{Width}x{Height}{Ext}.
	CosFileName *string `pulumi:"cosFileName"`
	// Cos Bucket folder prefix. If it is empty, set according to the default value /{Year}-{Month}-{Day}/.
	CosPrefix *string `pulumi:"cosPrefix"`
	// Cos region.
	CosRegion *string `pulumi:"cosRegion"`
	// Description information. Maximum length: 1024 bytes. Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
	Description *string `pulumi:"description"`
	// Screenshot height. Default: 0 (original height). Range: 0-2000.
	Height *int `pulumi:"height"`
	// Whether porn is enabled, 0: not enabled, 1: enabled. Default: 0.
	PornFlag *int `pulumi:"pornFlag"`
	// Screenshot interval, unit: s, default: 10s. Range: 2s~300s.
	SnapshotInterval *int `pulumi:"snapshotInterval"`
	// Template name. Maximum length: 255 bytes. Only Chinese, English, numbers, `_`, `-` are supported.
	TemplateName *string `pulumi:"templateName"`
	// Screenshot width. Default: 0 (original width). Range: 0-3000.
	Width *int `pulumi:"width"`
}

type SnapshotTemplateState struct {
	// Cos application ID.
	CosAppId pulumi.IntPtrInput
	// Cos bucket name. Note: The CosBucket parameter value cannot include the - [appid] part.
	CosBucket pulumi.StringPtrInput
	// Cos file name. If it is empty, set according to the default value
	// {StreamID}-screenshot-{Hour}-{Minute}-{Second}-{Width}x{Height}{Ext}.
	CosFileName pulumi.StringPtrInput
	// Cos Bucket folder prefix. If it is empty, set according to the default value /{Year}-{Month}-{Day}/.
	CosPrefix pulumi.StringPtrInput
	// Cos region.
	CosRegion pulumi.StringPtrInput
	// Description information. Maximum length: 1024 bytes. Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
	Description pulumi.StringPtrInput
	// Screenshot height. Default: 0 (original height). Range: 0-2000.
	Height pulumi.IntPtrInput
	// Whether porn is enabled, 0: not enabled, 1: enabled. Default: 0.
	PornFlag pulumi.IntPtrInput
	// Screenshot interval, unit: s, default: 10s. Range: 2s~300s.
	SnapshotInterval pulumi.IntPtrInput
	// Template name. Maximum length: 255 bytes. Only Chinese, English, numbers, `_`, `-` are supported.
	TemplateName pulumi.StringPtrInput
	// Screenshot width. Default: 0 (original width). Range: 0-3000.
	Width pulumi.IntPtrInput
}

func (SnapshotTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotTemplateState)(nil)).Elem()
}

type snapshotTemplateArgs struct {
	// Cos application ID.
	CosAppId int `pulumi:"cosAppId"`
	// Cos bucket name. Note: The CosBucket parameter value cannot include the - [appid] part.
	CosBucket string `pulumi:"cosBucket"`
	// Cos file name. If it is empty, set according to the default value
	// {StreamID}-screenshot-{Hour}-{Minute}-{Second}-{Width}x{Height}{Ext}.
	CosFileName *string `pulumi:"cosFileName"`
	// Cos Bucket folder prefix. If it is empty, set according to the default value /{Year}-{Month}-{Day}/.
	CosPrefix *string `pulumi:"cosPrefix"`
	// Cos region.
	CosRegion string `pulumi:"cosRegion"`
	// Description information. Maximum length: 1024 bytes. Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
	Description *string `pulumi:"description"`
	// Screenshot height. Default: 0 (original height). Range: 0-2000.
	Height *int `pulumi:"height"`
	// Whether porn is enabled, 0: not enabled, 1: enabled. Default: 0.
	PornFlag *int `pulumi:"pornFlag"`
	// Screenshot interval, unit: s, default: 10s. Range: 2s~300s.
	SnapshotInterval *int `pulumi:"snapshotInterval"`
	// Template name. Maximum length: 255 bytes. Only Chinese, English, numbers, `_`, `-` are supported.
	TemplateName string `pulumi:"templateName"`
	// Screenshot width. Default: 0 (original width). Range: 0-3000.
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a SnapshotTemplate resource.
type SnapshotTemplateArgs struct {
	// Cos application ID.
	CosAppId pulumi.IntInput
	// Cos bucket name. Note: The CosBucket parameter value cannot include the - [appid] part.
	CosBucket pulumi.StringInput
	// Cos file name. If it is empty, set according to the default value
	// {StreamID}-screenshot-{Hour}-{Minute}-{Second}-{Width}x{Height}{Ext}.
	CosFileName pulumi.StringPtrInput
	// Cos Bucket folder prefix. If it is empty, set according to the default value /{Year}-{Month}-{Day}/.
	CosPrefix pulumi.StringPtrInput
	// Cos region.
	CosRegion pulumi.StringInput
	// Description information. Maximum length: 1024 bytes. Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
	Description pulumi.StringPtrInput
	// Screenshot height. Default: 0 (original height). Range: 0-2000.
	Height pulumi.IntPtrInput
	// Whether porn is enabled, 0: not enabled, 1: enabled. Default: 0.
	PornFlag pulumi.IntPtrInput
	// Screenshot interval, unit: s, default: 10s. Range: 2s~300s.
	SnapshotInterval pulumi.IntPtrInput
	// Template name. Maximum length: 255 bytes. Only Chinese, English, numbers, `_`, `-` are supported.
	TemplateName pulumi.StringInput
	// Screenshot width. Default: 0 (original width). Range: 0-3000.
	Width pulumi.IntPtrInput
}

func (SnapshotTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotTemplateArgs)(nil)).Elem()
}

type SnapshotTemplateInput interface {
	pulumi.Input

	ToSnapshotTemplateOutput() SnapshotTemplateOutput
	ToSnapshotTemplateOutputWithContext(ctx context.Context) SnapshotTemplateOutput
}

func (*SnapshotTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotTemplate)(nil)).Elem()
}

func (i *SnapshotTemplate) ToSnapshotTemplateOutput() SnapshotTemplateOutput {
	return i.ToSnapshotTemplateOutputWithContext(context.Background())
}

func (i *SnapshotTemplate) ToSnapshotTemplateOutputWithContext(ctx context.Context) SnapshotTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotTemplateOutput)
}

// SnapshotTemplateArrayInput is an input type that accepts SnapshotTemplateArray and SnapshotTemplateArrayOutput values.
// You can construct a concrete instance of `SnapshotTemplateArrayInput` via:
//
//	SnapshotTemplateArray{ SnapshotTemplateArgs{...} }
type SnapshotTemplateArrayInput interface {
	pulumi.Input

	ToSnapshotTemplateArrayOutput() SnapshotTemplateArrayOutput
	ToSnapshotTemplateArrayOutputWithContext(context.Context) SnapshotTemplateArrayOutput
}

type SnapshotTemplateArray []SnapshotTemplateInput

func (SnapshotTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotTemplate)(nil)).Elem()
}

func (i SnapshotTemplateArray) ToSnapshotTemplateArrayOutput() SnapshotTemplateArrayOutput {
	return i.ToSnapshotTemplateArrayOutputWithContext(context.Background())
}

func (i SnapshotTemplateArray) ToSnapshotTemplateArrayOutputWithContext(ctx context.Context) SnapshotTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotTemplateArrayOutput)
}

// SnapshotTemplateMapInput is an input type that accepts SnapshotTemplateMap and SnapshotTemplateMapOutput values.
// You can construct a concrete instance of `SnapshotTemplateMapInput` via:
//
//	SnapshotTemplateMap{ "key": SnapshotTemplateArgs{...} }
type SnapshotTemplateMapInput interface {
	pulumi.Input

	ToSnapshotTemplateMapOutput() SnapshotTemplateMapOutput
	ToSnapshotTemplateMapOutputWithContext(context.Context) SnapshotTemplateMapOutput
}

type SnapshotTemplateMap map[string]SnapshotTemplateInput

func (SnapshotTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotTemplate)(nil)).Elem()
}

func (i SnapshotTemplateMap) ToSnapshotTemplateMapOutput() SnapshotTemplateMapOutput {
	return i.ToSnapshotTemplateMapOutputWithContext(context.Background())
}

func (i SnapshotTemplateMap) ToSnapshotTemplateMapOutputWithContext(ctx context.Context) SnapshotTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotTemplateMapOutput)
}

type SnapshotTemplateOutput struct{ *pulumi.OutputState }

func (SnapshotTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotTemplate)(nil)).Elem()
}

func (o SnapshotTemplateOutput) ToSnapshotTemplateOutput() SnapshotTemplateOutput {
	return o
}

func (o SnapshotTemplateOutput) ToSnapshotTemplateOutputWithContext(ctx context.Context) SnapshotTemplateOutput {
	return o
}

// Cos application ID.
func (o SnapshotTemplateOutput) CosAppId() pulumi.IntOutput {
	return o.ApplyT(func(v *SnapshotTemplate) pulumi.IntOutput { return v.CosAppId }).(pulumi.IntOutput)
}

// Cos bucket name. Note: The CosBucket parameter value cannot include the - [appid] part.
func (o SnapshotTemplateOutput) CosBucket() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotTemplate) pulumi.StringOutput { return v.CosBucket }).(pulumi.StringOutput)
}

// Cos file name. If it is empty, set according to the default value
// {StreamID}-screenshot-{Hour}-{Minute}-{Second}-{Width}x{Height}{Ext}.
func (o SnapshotTemplateOutput) CosFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotTemplate) pulumi.StringPtrOutput { return v.CosFileName }).(pulumi.StringPtrOutput)
}

// Cos Bucket folder prefix. If it is empty, set according to the default value /{Year}-{Month}-{Day}/.
func (o SnapshotTemplateOutput) CosPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotTemplate) pulumi.StringPtrOutput { return v.CosPrefix }).(pulumi.StringPtrOutput)
}

// Cos region.
func (o SnapshotTemplateOutput) CosRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotTemplate) pulumi.StringOutput { return v.CosRegion }).(pulumi.StringOutput)
}

// Description information. Maximum length: 1024 bytes. Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
func (o SnapshotTemplateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotTemplate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Screenshot height. Default: 0 (original height). Range: 0-2000.
func (o SnapshotTemplateOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotTemplate) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// Whether porn is enabled, 0: not enabled, 1: enabled. Default: 0.
func (o SnapshotTemplateOutput) PornFlag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotTemplate) pulumi.IntPtrOutput { return v.PornFlag }).(pulumi.IntPtrOutput)
}

// Screenshot interval, unit: s, default: 10s. Range: 2s~300s.
func (o SnapshotTemplateOutput) SnapshotInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotTemplate) pulumi.IntPtrOutput { return v.SnapshotInterval }).(pulumi.IntPtrOutput)
}

// Template name. Maximum length: 255 bytes. Only Chinese, English, numbers, `_`, `-` are supported.
func (o SnapshotTemplateOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotTemplate) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

// Screenshot width. Default: 0 (original width). Range: 0-3000.
func (o SnapshotTemplateOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotTemplate) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type SnapshotTemplateArrayOutput struct{ *pulumi.OutputState }

func (SnapshotTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotTemplate)(nil)).Elem()
}

func (o SnapshotTemplateArrayOutput) ToSnapshotTemplateArrayOutput() SnapshotTemplateArrayOutput {
	return o
}

func (o SnapshotTemplateArrayOutput) ToSnapshotTemplateArrayOutputWithContext(ctx context.Context) SnapshotTemplateArrayOutput {
	return o
}

func (o SnapshotTemplateArrayOutput) Index(i pulumi.IntInput) SnapshotTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnapshotTemplate {
		return vs[0].([]*SnapshotTemplate)[vs[1].(int)]
	}).(SnapshotTemplateOutput)
}

type SnapshotTemplateMapOutput struct{ *pulumi.OutputState }

func (SnapshotTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotTemplate)(nil)).Elem()
}

func (o SnapshotTemplateMapOutput) ToSnapshotTemplateMapOutput() SnapshotTemplateMapOutput {
	return o
}

func (o SnapshotTemplateMapOutput) ToSnapshotTemplateMapOutputWithContext(ctx context.Context) SnapshotTemplateMapOutput {
	return o
}

func (o SnapshotTemplateMapOutput) MapIndex(k pulumi.StringInput) SnapshotTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnapshotTemplate {
		return vs[0].(map[string]*SnapshotTemplate)[vs[1].(string)]
	}).(SnapshotTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotTemplateInput)(nil)).Elem(), &SnapshotTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotTemplateArrayInput)(nil)).Elem(), SnapshotTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotTemplateMapInput)(nil)).Elem(), SnapshotTemplateMap{})
	pulumi.RegisterOutputType(SnapshotTemplateOutput{})
	pulumi.RegisterOutputType(SnapshotTemplateArrayOutput{})
	pulumi.RegisterOutputType(SnapshotTemplateMapOutput{})
}