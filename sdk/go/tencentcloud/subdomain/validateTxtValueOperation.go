// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package subdomain

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ValidateTxtValueOperation struct {
	pulumi.CustomResourceState

	// The domain name for which TXT records need to be added.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The subdomain to add Zone domain.
	DomainZone pulumi.StringOutput `pulumi:"domainZone"`
	// Record types need to be added.
	RecordType pulumi.StringOutput `pulumi:"recordType"`
	// Host records that need to be added to TXT records.
	Subdomain pulumi.StringOutput `pulumi:"subdomain"`
	// The record value of the TXT record needs to be added.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewValidateTxtValueOperation registers a new resource with the given unique name, arguments, and options.
func NewValidateTxtValueOperation(ctx *pulumi.Context,
	name string, args *ValidateTxtValueOperationArgs, opts ...pulumi.ResourceOption) (*ValidateTxtValueOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainZone == nil {
		return nil, errors.New("invalid value for required argument 'DomainZone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ValidateTxtValueOperation
	err := ctx.RegisterResource("tencentcloud:Subdomain/validateTxtValueOperation:ValidateTxtValueOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetValidateTxtValueOperation gets an existing ValidateTxtValueOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetValidateTxtValueOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ValidateTxtValueOperationState, opts ...pulumi.ResourceOption) (*ValidateTxtValueOperation, error) {
	var resource ValidateTxtValueOperation
	err := ctx.ReadResource("tencentcloud:Subdomain/validateTxtValueOperation:ValidateTxtValueOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ValidateTxtValueOperation resources.
type validateTxtValueOperationState struct {
	// The domain name for which TXT records need to be added.
	Domain *string `pulumi:"domain"`
	// The subdomain to add Zone domain.
	DomainZone *string `pulumi:"domainZone"`
	// Record types need to be added.
	RecordType *string `pulumi:"recordType"`
	// Host records that need to be added to TXT records.
	Subdomain *string `pulumi:"subdomain"`
	// The record value of the TXT record needs to be added.
	Value *string `pulumi:"value"`
}

type ValidateTxtValueOperationState struct {
	// The domain name for which TXT records need to be added.
	Domain pulumi.StringPtrInput
	// The subdomain to add Zone domain.
	DomainZone pulumi.StringPtrInput
	// Record types need to be added.
	RecordType pulumi.StringPtrInput
	// Host records that need to be added to TXT records.
	Subdomain pulumi.StringPtrInput
	// The record value of the TXT record needs to be added.
	Value pulumi.StringPtrInput
}

func (ValidateTxtValueOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*validateTxtValueOperationState)(nil)).Elem()
}

type validateTxtValueOperationArgs struct {
	// The subdomain to add Zone domain.
	DomainZone string `pulumi:"domainZone"`
}

// The set of arguments for constructing a ValidateTxtValueOperation resource.
type ValidateTxtValueOperationArgs struct {
	// The subdomain to add Zone domain.
	DomainZone pulumi.StringInput
}

func (ValidateTxtValueOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*validateTxtValueOperationArgs)(nil)).Elem()
}

type ValidateTxtValueOperationInput interface {
	pulumi.Input

	ToValidateTxtValueOperationOutput() ValidateTxtValueOperationOutput
	ToValidateTxtValueOperationOutputWithContext(ctx context.Context) ValidateTxtValueOperationOutput
}

func (*ValidateTxtValueOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**ValidateTxtValueOperation)(nil)).Elem()
}

func (i *ValidateTxtValueOperation) ToValidateTxtValueOperationOutput() ValidateTxtValueOperationOutput {
	return i.ToValidateTxtValueOperationOutputWithContext(context.Background())
}

func (i *ValidateTxtValueOperation) ToValidateTxtValueOperationOutputWithContext(ctx context.Context) ValidateTxtValueOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidateTxtValueOperationOutput)
}

// ValidateTxtValueOperationArrayInput is an input type that accepts ValidateTxtValueOperationArray and ValidateTxtValueOperationArrayOutput values.
// You can construct a concrete instance of `ValidateTxtValueOperationArrayInput` via:
//
//	ValidateTxtValueOperationArray{ ValidateTxtValueOperationArgs{...} }
type ValidateTxtValueOperationArrayInput interface {
	pulumi.Input

	ToValidateTxtValueOperationArrayOutput() ValidateTxtValueOperationArrayOutput
	ToValidateTxtValueOperationArrayOutputWithContext(context.Context) ValidateTxtValueOperationArrayOutput
}

type ValidateTxtValueOperationArray []ValidateTxtValueOperationInput

func (ValidateTxtValueOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ValidateTxtValueOperation)(nil)).Elem()
}

func (i ValidateTxtValueOperationArray) ToValidateTxtValueOperationArrayOutput() ValidateTxtValueOperationArrayOutput {
	return i.ToValidateTxtValueOperationArrayOutputWithContext(context.Background())
}

func (i ValidateTxtValueOperationArray) ToValidateTxtValueOperationArrayOutputWithContext(ctx context.Context) ValidateTxtValueOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidateTxtValueOperationArrayOutput)
}

// ValidateTxtValueOperationMapInput is an input type that accepts ValidateTxtValueOperationMap and ValidateTxtValueOperationMapOutput values.
// You can construct a concrete instance of `ValidateTxtValueOperationMapInput` via:
//
//	ValidateTxtValueOperationMap{ "key": ValidateTxtValueOperationArgs{...} }
type ValidateTxtValueOperationMapInput interface {
	pulumi.Input

	ToValidateTxtValueOperationMapOutput() ValidateTxtValueOperationMapOutput
	ToValidateTxtValueOperationMapOutputWithContext(context.Context) ValidateTxtValueOperationMapOutput
}

type ValidateTxtValueOperationMap map[string]ValidateTxtValueOperationInput

func (ValidateTxtValueOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ValidateTxtValueOperation)(nil)).Elem()
}

func (i ValidateTxtValueOperationMap) ToValidateTxtValueOperationMapOutput() ValidateTxtValueOperationMapOutput {
	return i.ToValidateTxtValueOperationMapOutputWithContext(context.Background())
}

func (i ValidateTxtValueOperationMap) ToValidateTxtValueOperationMapOutputWithContext(ctx context.Context) ValidateTxtValueOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidateTxtValueOperationMapOutput)
}

type ValidateTxtValueOperationOutput struct{ *pulumi.OutputState }

func (ValidateTxtValueOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ValidateTxtValueOperation)(nil)).Elem()
}

func (o ValidateTxtValueOperationOutput) ToValidateTxtValueOperationOutput() ValidateTxtValueOperationOutput {
	return o
}

func (o ValidateTxtValueOperationOutput) ToValidateTxtValueOperationOutputWithContext(ctx context.Context) ValidateTxtValueOperationOutput {
	return o
}

// The domain name for which TXT records need to be added.
func (o ValidateTxtValueOperationOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *ValidateTxtValueOperation) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The subdomain to add Zone domain.
func (o ValidateTxtValueOperationOutput) DomainZone() pulumi.StringOutput {
	return o.ApplyT(func(v *ValidateTxtValueOperation) pulumi.StringOutput { return v.DomainZone }).(pulumi.StringOutput)
}

// Record types need to be added.
func (o ValidateTxtValueOperationOutput) RecordType() pulumi.StringOutput {
	return o.ApplyT(func(v *ValidateTxtValueOperation) pulumi.StringOutput { return v.RecordType }).(pulumi.StringOutput)
}

// Host records that need to be added to TXT records.
func (o ValidateTxtValueOperationOutput) Subdomain() pulumi.StringOutput {
	return o.ApplyT(func(v *ValidateTxtValueOperation) pulumi.StringOutput { return v.Subdomain }).(pulumi.StringOutput)
}

// The record value of the TXT record needs to be added.
func (o ValidateTxtValueOperationOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ValidateTxtValueOperation) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type ValidateTxtValueOperationArrayOutput struct{ *pulumi.OutputState }

func (ValidateTxtValueOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ValidateTxtValueOperation)(nil)).Elem()
}

func (o ValidateTxtValueOperationArrayOutput) ToValidateTxtValueOperationArrayOutput() ValidateTxtValueOperationArrayOutput {
	return o
}

func (o ValidateTxtValueOperationArrayOutput) ToValidateTxtValueOperationArrayOutputWithContext(ctx context.Context) ValidateTxtValueOperationArrayOutput {
	return o
}

func (o ValidateTxtValueOperationArrayOutput) Index(i pulumi.IntInput) ValidateTxtValueOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ValidateTxtValueOperation {
		return vs[0].([]*ValidateTxtValueOperation)[vs[1].(int)]
	}).(ValidateTxtValueOperationOutput)
}

type ValidateTxtValueOperationMapOutput struct{ *pulumi.OutputState }

func (ValidateTxtValueOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ValidateTxtValueOperation)(nil)).Elem()
}

func (o ValidateTxtValueOperationMapOutput) ToValidateTxtValueOperationMapOutput() ValidateTxtValueOperationMapOutput {
	return o
}

func (o ValidateTxtValueOperationMapOutput) ToValidateTxtValueOperationMapOutputWithContext(ctx context.Context) ValidateTxtValueOperationMapOutput {
	return o
}

func (o ValidateTxtValueOperationMapOutput) MapIndex(k pulumi.StringInput) ValidateTxtValueOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ValidateTxtValueOperation {
		return vs[0].(map[string]*ValidateTxtValueOperation)[vs[1].(string)]
	}).(ValidateTxtValueOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ValidateTxtValueOperationInput)(nil)).Elem(), &ValidateTxtValueOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ValidateTxtValueOperationArrayInput)(nil)).Elem(), ValidateTxtValueOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ValidateTxtValueOperationMapInput)(nil)).Elem(), ValidateTxtValueOperationMap{})
	pulumi.RegisterOutputType(ValidateTxtValueOperationOutput{})
	pulumi.RegisterOutputType(ValidateTxtValueOperationArrayOutput{})
	pulumi.RegisterOutputType(ValidateTxtValueOperationMapOutput{})
}
