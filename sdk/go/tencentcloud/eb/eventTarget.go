// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eb

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventTarget struct {
	pulumi.CustomResourceState

	// event bus id.
	EventBusId pulumi.StringOutput `pulumi:"eventBusId"`
	// event rule id.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// target description.
	TargetDescription EventTargetTargetDescriptionOutput `pulumi:"targetDescription"`
	// target type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEventTarget registers a new resource with the given unique name, arguments, and options.
func NewEventTarget(ctx *pulumi.Context,
	name string, args *EventTargetArgs, opts ...pulumi.ResourceOption) (*EventTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventBusId == nil {
		return nil, errors.New("invalid value for required argument 'EventBusId'")
	}
	if args.RuleId == nil {
		return nil, errors.New("invalid value for required argument 'RuleId'")
	}
	if args.TargetDescription == nil {
		return nil, errors.New("invalid value for required argument 'TargetDescription'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EventTarget
	err := ctx.RegisterResource("tencentcloud:Eb/eventTarget:EventTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventTarget gets an existing EventTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventTargetState, opts ...pulumi.ResourceOption) (*EventTarget, error) {
	var resource EventTarget
	err := ctx.ReadResource("tencentcloud:Eb/eventTarget:EventTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventTarget resources.
type eventTargetState struct {
	// event bus id.
	EventBusId *string `pulumi:"eventBusId"`
	// event rule id.
	RuleId *string `pulumi:"ruleId"`
	// target description.
	TargetDescription *EventTargetTargetDescription `pulumi:"targetDescription"`
	// target type.
	Type *string `pulumi:"type"`
}

type EventTargetState struct {
	// event bus id.
	EventBusId pulumi.StringPtrInput
	// event rule id.
	RuleId pulumi.StringPtrInput
	// target description.
	TargetDescription EventTargetTargetDescriptionPtrInput
	// target type.
	Type pulumi.StringPtrInput
}

func (EventTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventTargetState)(nil)).Elem()
}

type eventTargetArgs struct {
	// event bus id.
	EventBusId string `pulumi:"eventBusId"`
	// event rule id.
	RuleId string `pulumi:"ruleId"`
	// target description.
	TargetDescription EventTargetTargetDescription `pulumi:"targetDescription"`
	// target type.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a EventTarget resource.
type EventTargetArgs struct {
	// event bus id.
	EventBusId pulumi.StringInput
	// event rule id.
	RuleId pulumi.StringInput
	// target description.
	TargetDescription EventTargetTargetDescriptionInput
	// target type.
	Type pulumi.StringInput
}

func (EventTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventTargetArgs)(nil)).Elem()
}

type EventTargetInput interface {
	pulumi.Input

	ToEventTargetOutput() EventTargetOutput
	ToEventTargetOutputWithContext(ctx context.Context) EventTargetOutput
}

func (*EventTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**EventTarget)(nil)).Elem()
}

func (i *EventTarget) ToEventTargetOutput() EventTargetOutput {
	return i.ToEventTargetOutputWithContext(context.Background())
}

func (i *EventTarget) ToEventTargetOutputWithContext(ctx context.Context) EventTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTargetOutput)
}

// EventTargetArrayInput is an input type that accepts EventTargetArray and EventTargetArrayOutput values.
// You can construct a concrete instance of `EventTargetArrayInput` via:
//
//	EventTargetArray{ EventTargetArgs{...} }
type EventTargetArrayInput interface {
	pulumi.Input

	ToEventTargetArrayOutput() EventTargetArrayOutput
	ToEventTargetArrayOutputWithContext(context.Context) EventTargetArrayOutput
}

type EventTargetArray []EventTargetInput

func (EventTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventTarget)(nil)).Elem()
}

func (i EventTargetArray) ToEventTargetArrayOutput() EventTargetArrayOutput {
	return i.ToEventTargetArrayOutputWithContext(context.Background())
}

func (i EventTargetArray) ToEventTargetArrayOutputWithContext(ctx context.Context) EventTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTargetArrayOutput)
}

// EventTargetMapInput is an input type that accepts EventTargetMap and EventTargetMapOutput values.
// You can construct a concrete instance of `EventTargetMapInput` via:
//
//	EventTargetMap{ "key": EventTargetArgs{...} }
type EventTargetMapInput interface {
	pulumi.Input

	ToEventTargetMapOutput() EventTargetMapOutput
	ToEventTargetMapOutputWithContext(context.Context) EventTargetMapOutput
}

type EventTargetMap map[string]EventTargetInput

func (EventTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventTarget)(nil)).Elem()
}

func (i EventTargetMap) ToEventTargetMapOutput() EventTargetMapOutput {
	return i.ToEventTargetMapOutputWithContext(context.Background())
}

func (i EventTargetMap) ToEventTargetMapOutputWithContext(ctx context.Context) EventTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTargetMapOutput)
}

type EventTargetOutput struct{ *pulumi.OutputState }

func (EventTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventTarget)(nil)).Elem()
}

func (o EventTargetOutput) ToEventTargetOutput() EventTargetOutput {
	return o
}

func (o EventTargetOutput) ToEventTargetOutputWithContext(ctx context.Context) EventTargetOutput {
	return o
}

// event bus id.
func (o EventTargetOutput) EventBusId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventTarget) pulumi.StringOutput { return v.EventBusId }).(pulumi.StringOutput)
}

// event rule id.
func (o EventTargetOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventTarget) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// target description.
func (o EventTargetOutput) TargetDescription() EventTargetTargetDescriptionOutput {
	return o.ApplyT(func(v *EventTarget) EventTargetTargetDescriptionOutput { return v.TargetDescription }).(EventTargetTargetDescriptionOutput)
}

// target type.
func (o EventTargetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventTarget) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type EventTargetArrayOutput struct{ *pulumi.OutputState }

func (EventTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventTarget)(nil)).Elem()
}

func (o EventTargetArrayOutput) ToEventTargetArrayOutput() EventTargetArrayOutput {
	return o
}

func (o EventTargetArrayOutput) ToEventTargetArrayOutputWithContext(ctx context.Context) EventTargetArrayOutput {
	return o
}

func (o EventTargetArrayOutput) Index(i pulumi.IntInput) EventTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventTarget {
		return vs[0].([]*EventTarget)[vs[1].(int)]
	}).(EventTargetOutput)
}

type EventTargetMapOutput struct{ *pulumi.OutputState }

func (EventTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventTarget)(nil)).Elem()
}

func (o EventTargetMapOutput) ToEventTargetMapOutput() EventTargetMapOutput {
	return o
}

func (o EventTargetMapOutput) ToEventTargetMapOutputWithContext(ctx context.Context) EventTargetMapOutput {
	return o
}

func (o EventTargetMapOutput) MapIndex(k pulumi.StringInput) EventTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventTarget {
		return vs[0].(map[string]*EventTarget)[vs[1].(string)]
	}).(EventTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventTargetInput)(nil)).Elem(), &EventTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventTargetArrayInput)(nil)).Elem(), EventTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventTargetMapInput)(nil)).Elem(), EventTargetMap{})
	pulumi.RegisterOutputType(EventTargetOutput{})
	pulumi.RegisterOutputType(EventTargetArrayOutput{})
	pulumi.RegisterOutputType(EventTargetMapOutput{})
}
