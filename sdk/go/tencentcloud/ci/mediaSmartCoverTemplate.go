// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ci

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MediaSmartCoverTemplate struct {
	pulumi.CustomResourceState

	// bucket name.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Smart Cover Parameters.
	SmartCover MediaSmartCoverTemplateSmartCoverOutput `pulumi:"smartCover"`
}

// NewMediaSmartCoverTemplate registers a new resource with the given unique name, arguments, and options.
func NewMediaSmartCoverTemplate(ctx *pulumi.Context,
	name string, args *MediaSmartCoverTemplateArgs, opts ...pulumi.ResourceOption) (*MediaSmartCoverTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.SmartCover == nil {
		return nil, errors.New("invalid value for required argument 'SmartCover'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MediaSmartCoverTemplate
	err := ctx.RegisterResource("tencentcloud:Ci/mediaSmartCoverTemplate:MediaSmartCoverTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMediaSmartCoverTemplate gets an existing MediaSmartCoverTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMediaSmartCoverTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MediaSmartCoverTemplateState, opts ...pulumi.ResourceOption) (*MediaSmartCoverTemplate, error) {
	var resource MediaSmartCoverTemplate
	err := ctx.ReadResource("tencentcloud:Ci/mediaSmartCoverTemplate:MediaSmartCoverTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MediaSmartCoverTemplate resources.
type mediaSmartCoverTemplateState struct {
	// bucket name.
	Bucket *string `pulumi:"bucket"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name *string `pulumi:"name"`
	// Smart Cover Parameters.
	SmartCover *MediaSmartCoverTemplateSmartCover `pulumi:"smartCover"`
}

type MediaSmartCoverTemplateState struct {
	// bucket name.
	Bucket pulumi.StringPtrInput
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringPtrInput
	// Smart Cover Parameters.
	SmartCover MediaSmartCoverTemplateSmartCoverPtrInput
}

func (MediaSmartCoverTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaSmartCoverTemplateState)(nil)).Elem()
}

type mediaSmartCoverTemplateArgs struct {
	// bucket name.
	Bucket string `pulumi:"bucket"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name *string `pulumi:"name"`
	// Smart Cover Parameters.
	SmartCover MediaSmartCoverTemplateSmartCover `pulumi:"smartCover"`
}

// The set of arguments for constructing a MediaSmartCoverTemplate resource.
type MediaSmartCoverTemplateArgs struct {
	// bucket name.
	Bucket pulumi.StringInput
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringPtrInput
	// Smart Cover Parameters.
	SmartCover MediaSmartCoverTemplateSmartCoverInput
}

func (MediaSmartCoverTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaSmartCoverTemplateArgs)(nil)).Elem()
}

type MediaSmartCoverTemplateInput interface {
	pulumi.Input

	ToMediaSmartCoverTemplateOutput() MediaSmartCoverTemplateOutput
	ToMediaSmartCoverTemplateOutputWithContext(ctx context.Context) MediaSmartCoverTemplateOutput
}

func (*MediaSmartCoverTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaSmartCoverTemplate)(nil)).Elem()
}

func (i *MediaSmartCoverTemplate) ToMediaSmartCoverTemplateOutput() MediaSmartCoverTemplateOutput {
	return i.ToMediaSmartCoverTemplateOutputWithContext(context.Background())
}

func (i *MediaSmartCoverTemplate) ToMediaSmartCoverTemplateOutputWithContext(ctx context.Context) MediaSmartCoverTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaSmartCoverTemplateOutput)
}

// MediaSmartCoverTemplateArrayInput is an input type that accepts MediaSmartCoverTemplateArray and MediaSmartCoverTemplateArrayOutput values.
// You can construct a concrete instance of `MediaSmartCoverTemplateArrayInput` via:
//
//	MediaSmartCoverTemplateArray{ MediaSmartCoverTemplateArgs{...} }
type MediaSmartCoverTemplateArrayInput interface {
	pulumi.Input

	ToMediaSmartCoverTemplateArrayOutput() MediaSmartCoverTemplateArrayOutput
	ToMediaSmartCoverTemplateArrayOutputWithContext(context.Context) MediaSmartCoverTemplateArrayOutput
}

type MediaSmartCoverTemplateArray []MediaSmartCoverTemplateInput

func (MediaSmartCoverTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MediaSmartCoverTemplate)(nil)).Elem()
}

func (i MediaSmartCoverTemplateArray) ToMediaSmartCoverTemplateArrayOutput() MediaSmartCoverTemplateArrayOutput {
	return i.ToMediaSmartCoverTemplateArrayOutputWithContext(context.Background())
}

func (i MediaSmartCoverTemplateArray) ToMediaSmartCoverTemplateArrayOutputWithContext(ctx context.Context) MediaSmartCoverTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaSmartCoverTemplateArrayOutput)
}

// MediaSmartCoverTemplateMapInput is an input type that accepts MediaSmartCoverTemplateMap and MediaSmartCoverTemplateMapOutput values.
// You can construct a concrete instance of `MediaSmartCoverTemplateMapInput` via:
//
//	MediaSmartCoverTemplateMap{ "key": MediaSmartCoverTemplateArgs{...} }
type MediaSmartCoverTemplateMapInput interface {
	pulumi.Input

	ToMediaSmartCoverTemplateMapOutput() MediaSmartCoverTemplateMapOutput
	ToMediaSmartCoverTemplateMapOutputWithContext(context.Context) MediaSmartCoverTemplateMapOutput
}

type MediaSmartCoverTemplateMap map[string]MediaSmartCoverTemplateInput

func (MediaSmartCoverTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MediaSmartCoverTemplate)(nil)).Elem()
}

func (i MediaSmartCoverTemplateMap) ToMediaSmartCoverTemplateMapOutput() MediaSmartCoverTemplateMapOutput {
	return i.ToMediaSmartCoverTemplateMapOutputWithContext(context.Background())
}

func (i MediaSmartCoverTemplateMap) ToMediaSmartCoverTemplateMapOutputWithContext(ctx context.Context) MediaSmartCoverTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaSmartCoverTemplateMapOutput)
}

type MediaSmartCoverTemplateOutput struct{ *pulumi.OutputState }

func (MediaSmartCoverTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaSmartCoverTemplate)(nil)).Elem()
}

func (o MediaSmartCoverTemplateOutput) ToMediaSmartCoverTemplateOutput() MediaSmartCoverTemplateOutput {
	return o
}

func (o MediaSmartCoverTemplateOutput) ToMediaSmartCoverTemplateOutputWithContext(ctx context.Context) MediaSmartCoverTemplateOutput {
	return o
}

// bucket name.
func (o MediaSmartCoverTemplateOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaSmartCoverTemplate) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
func (o MediaSmartCoverTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaSmartCoverTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Smart Cover Parameters.
func (o MediaSmartCoverTemplateOutput) SmartCover() MediaSmartCoverTemplateSmartCoverOutput {
	return o.ApplyT(func(v *MediaSmartCoverTemplate) MediaSmartCoverTemplateSmartCoverOutput { return v.SmartCover }).(MediaSmartCoverTemplateSmartCoverOutput)
}

type MediaSmartCoverTemplateArrayOutput struct{ *pulumi.OutputState }

func (MediaSmartCoverTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MediaSmartCoverTemplate)(nil)).Elem()
}

func (o MediaSmartCoverTemplateArrayOutput) ToMediaSmartCoverTemplateArrayOutput() MediaSmartCoverTemplateArrayOutput {
	return o
}

func (o MediaSmartCoverTemplateArrayOutput) ToMediaSmartCoverTemplateArrayOutputWithContext(ctx context.Context) MediaSmartCoverTemplateArrayOutput {
	return o
}

func (o MediaSmartCoverTemplateArrayOutput) Index(i pulumi.IntInput) MediaSmartCoverTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MediaSmartCoverTemplate {
		return vs[0].([]*MediaSmartCoverTemplate)[vs[1].(int)]
	}).(MediaSmartCoverTemplateOutput)
}

type MediaSmartCoverTemplateMapOutput struct{ *pulumi.OutputState }

func (MediaSmartCoverTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MediaSmartCoverTemplate)(nil)).Elem()
}

func (o MediaSmartCoverTemplateMapOutput) ToMediaSmartCoverTemplateMapOutput() MediaSmartCoverTemplateMapOutput {
	return o
}

func (o MediaSmartCoverTemplateMapOutput) ToMediaSmartCoverTemplateMapOutputWithContext(ctx context.Context) MediaSmartCoverTemplateMapOutput {
	return o
}

func (o MediaSmartCoverTemplateMapOutput) MapIndex(k pulumi.StringInput) MediaSmartCoverTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MediaSmartCoverTemplate {
		return vs[0].(map[string]*MediaSmartCoverTemplate)[vs[1].(string)]
	}).(MediaSmartCoverTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MediaSmartCoverTemplateInput)(nil)).Elem(), &MediaSmartCoverTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaSmartCoverTemplateArrayInput)(nil)).Elem(), MediaSmartCoverTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaSmartCoverTemplateMapInput)(nil)).Elem(), MediaSmartCoverTemplateMap{})
	pulumi.RegisterOutputType(MediaSmartCoverTemplateOutput{})
	pulumi.RegisterOutputType(MediaSmartCoverTemplateArrayOutput{})
	pulumi.RegisterOutputType(MediaSmartCoverTemplateMapOutput{})
}