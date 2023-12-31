// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CancelDcnJobOperation struct {
	pulumi.CustomResourceState

	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewCancelDcnJobOperation registers a new resource with the given unique name, arguments, and options.
func NewCancelDcnJobOperation(ctx *pulumi.Context,
	name string, args *CancelDcnJobOperationArgs, opts ...pulumi.ResourceOption) (*CancelDcnJobOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CancelDcnJobOperation
	err := ctx.RegisterResource("tencentcloud:Dcdb/cancelDcnJobOperation:CancelDcnJobOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCancelDcnJobOperation gets an existing CancelDcnJobOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCancelDcnJobOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CancelDcnJobOperationState, opts ...pulumi.ResourceOption) (*CancelDcnJobOperation, error) {
	var resource CancelDcnJobOperation
	err := ctx.ReadResource("tencentcloud:Dcdb/cancelDcnJobOperation:CancelDcnJobOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CancelDcnJobOperation resources.
type cancelDcnJobOperationState struct {
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
}

type CancelDcnJobOperationState struct {
	// Instance ID.
	InstanceId pulumi.StringPtrInput
}

func (CancelDcnJobOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*cancelDcnJobOperationState)(nil)).Elem()
}

type cancelDcnJobOperationArgs struct {
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a CancelDcnJobOperation resource.
type CancelDcnJobOperationArgs struct {
	// Instance ID.
	InstanceId pulumi.StringInput
}

func (CancelDcnJobOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cancelDcnJobOperationArgs)(nil)).Elem()
}

type CancelDcnJobOperationInput interface {
	pulumi.Input

	ToCancelDcnJobOperationOutput() CancelDcnJobOperationOutput
	ToCancelDcnJobOperationOutputWithContext(ctx context.Context) CancelDcnJobOperationOutput
}

func (*CancelDcnJobOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**CancelDcnJobOperation)(nil)).Elem()
}

func (i *CancelDcnJobOperation) ToCancelDcnJobOperationOutput() CancelDcnJobOperationOutput {
	return i.ToCancelDcnJobOperationOutputWithContext(context.Background())
}

func (i *CancelDcnJobOperation) ToCancelDcnJobOperationOutputWithContext(ctx context.Context) CancelDcnJobOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CancelDcnJobOperationOutput)
}

// CancelDcnJobOperationArrayInput is an input type that accepts CancelDcnJobOperationArray and CancelDcnJobOperationArrayOutput values.
// You can construct a concrete instance of `CancelDcnJobOperationArrayInput` via:
//
//	CancelDcnJobOperationArray{ CancelDcnJobOperationArgs{...} }
type CancelDcnJobOperationArrayInput interface {
	pulumi.Input

	ToCancelDcnJobOperationArrayOutput() CancelDcnJobOperationArrayOutput
	ToCancelDcnJobOperationArrayOutputWithContext(context.Context) CancelDcnJobOperationArrayOutput
}

type CancelDcnJobOperationArray []CancelDcnJobOperationInput

func (CancelDcnJobOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CancelDcnJobOperation)(nil)).Elem()
}

func (i CancelDcnJobOperationArray) ToCancelDcnJobOperationArrayOutput() CancelDcnJobOperationArrayOutput {
	return i.ToCancelDcnJobOperationArrayOutputWithContext(context.Background())
}

func (i CancelDcnJobOperationArray) ToCancelDcnJobOperationArrayOutputWithContext(ctx context.Context) CancelDcnJobOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CancelDcnJobOperationArrayOutput)
}

// CancelDcnJobOperationMapInput is an input type that accepts CancelDcnJobOperationMap and CancelDcnJobOperationMapOutput values.
// You can construct a concrete instance of `CancelDcnJobOperationMapInput` via:
//
//	CancelDcnJobOperationMap{ "key": CancelDcnJobOperationArgs{...} }
type CancelDcnJobOperationMapInput interface {
	pulumi.Input

	ToCancelDcnJobOperationMapOutput() CancelDcnJobOperationMapOutput
	ToCancelDcnJobOperationMapOutputWithContext(context.Context) CancelDcnJobOperationMapOutput
}

type CancelDcnJobOperationMap map[string]CancelDcnJobOperationInput

func (CancelDcnJobOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CancelDcnJobOperation)(nil)).Elem()
}

func (i CancelDcnJobOperationMap) ToCancelDcnJobOperationMapOutput() CancelDcnJobOperationMapOutput {
	return i.ToCancelDcnJobOperationMapOutputWithContext(context.Background())
}

func (i CancelDcnJobOperationMap) ToCancelDcnJobOperationMapOutputWithContext(ctx context.Context) CancelDcnJobOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CancelDcnJobOperationMapOutput)
}

type CancelDcnJobOperationOutput struct{ *pulumi.OutputState }

func (CancelDcnJobOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CancelDcnJobOperation)(nil)).Elem()
}

func (o CancelDcnJobOperationOutput) ToCancelDcnJobOperationOutput() CancelDcnJobOperationOutput {
	return o
}

func (o CancelDcnJobOperationOutput) ToCancelDcnJobOperationOutputWithContext(ctx context.Context) CancelDcnJobOperationOutput {
	return o
}

// Instance ID.
func (o CancelDcnJobOperationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CancelDcnJobOperation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type CancelDcnJobOperationArrayOutput struct{ *pulumi.OutputState }

func (CancelDcnJobOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CancelDcnJobOperation)(nil)).Elem()
}

func (o CancelDcnJobOperationArrayOutput) ToCancelDcnJobOperationArrayOutput() CancelDcnJobOperationArrayOutput {
	return o
}

func (o CancelDcnJobOperationArrayOutput) ToCancelDcnJobOperationArrayOutputWithContext(ctx context.Context) CancelDcnJobOperationArrayOutput {
	return o
}

func (o CancelDcnJobOperationArrayOutput) Index(i pulumi.IntInput) CancelDcnJobOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CancelDcnJobOperation {
		return vs[0].([]*CancelDcnJobOperation)[vs[1].(int)]
	}).(CancelDcnJobOperationOutput)
}

type CancelDcnJobOperationMapOutput struct{ *pulumi.OutputState }

func (CancelDcnJobOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CancelDcnJobOperation)(nil)).Elem()
}

func (o CancelDcnJobOperationMapOutput) ToCancelDcnJobOperationMapOutput() CancelDcnJobOperationMapOutput {
	return o
}

func (o CancelDcnJobOperationMapOutput) ToCancelDcnJobOperationMapOutputWithContext(ctx context.Context) CancelDcnJobOperationMapOutput {
	return o
}

func (o CancelDcnJobOperationMapOutput) MapIndex(k pulumi.StringInput) CancelDcnJobOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CancelDcnJobOperation {
		return vs[0].(map[string]*CancelDcnJobOperation)[vs[1].(string)]
	}).(CancelDcnJobOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CancelDcnJobOperationInput)(nil)).Elem(), &CancelDcnJobOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*CancelDcnJobOperationArrayInput)(nil)).Elem(), CancelDcnJobOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CancelDcnJobOperationMapInput)(nil)).Elem(), CancelDcnJobOperationMap{})
	pulumi.RegisterOutputType(CancelDcnJobOperationOutput{})
	pulumi.RegisterOutputType(CancelDcnJobOperationArrayOutput{})
	pulumi.RegisterOutputType(CancelDcnJobOperationMapOutput{})
}
