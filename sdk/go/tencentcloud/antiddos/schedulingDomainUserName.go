// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package antiddos

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SchedulingDomainUserName struct {
	pulumi.CustomResourceState

	// user cname.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// domain name.
	DomainUserName pulumi.StringOutput `pulumi:"domainUserName"`
}

// NewSchedulingDomainUserName registers a new resource with the given unique name, arguments, and options.
func NewSchedulingDomainUserName(ctx *pulumi.Context,
	name string, args *SchedulingDomainUserNameArgs, opts ...pulumi.ResourceOption) (*SchedulingDomainUserName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.DomainUserName == nil {
		return nil, errors.New("invalid value for required argument 'DomainUserName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SchedulingDomainUserName
	err := ctx.RegisterResource("tencentcloud:Antiddos/schedulingDomainUserName:SchedulingDomainUserName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchedulingDomainUserName gets an existing SchedulingDomainUserName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchedulingDomainUserName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchedulingDomainUserNameState, opts ...pulumi.ResourceOption) (*SchedulingDomainUserName, error) {
	var resource SchedulingDomainUserName
	err := ctx.ReadResource("tencentcloud:Antiddos/schedulingDomainUserName:SchedulingDomainUserName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SchedulingDomainUserName resources.
type schedulingDomainUserNameState struct {
	// user cname.
	DomainName *string `pulumi:"domainName"`
	// domain name.
	DomainUserName *string `pulumi:"domainUserName"`
}

type SchedulingDomainUserNameState struct {
	// user cname.
	DomainName pulumi.StringPtrInput
	// domain name.
	DomainUserName pulumi.StringPtrInput
}

func (SchedulingDomainUserNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*schedulingDomainUserNameState)(nil)).Elem()
}

type schedulingDomainUserNameArgs struct {
	// user cname.
	DomainName string `pulumi:"domainName"`
	// domain name.
	DomainUserName string `pulumi:"domainUserName"`
}

// The set of arguments for constructing a SchedulingDomainUserName resource.
type SchedulingDomainUserNameArgs struct {
	// user cname.
	DomainName pulumi.StringInput
	// domain name.
	DomainUserName pulumi.StringInput
}

func (SchedulingDomainUserNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schedulingDomainUserNameArgs)(nil)).Elem()
}

type SchedulingDomainUserNameInput interface {
	pulumi.Input

	ToSchedulingDomainUserNameOutput() SchedulingDomainUserNameOutput
	ToSchedulingDomainUserNameOutputWithContext(ctx context.Context) SchedulingDomainUserNameOutput
}

func (*SchedulingDomainUserName) ElementType() reflect.Type {
	return reflect.TypeOf((**SchedulingDomainUserName)(nil)).Elem()
}

func (i *SchedulingDomainUserName) ToSchedulingDomainUserNameOutput() SchedulingDomainUserNameOutput {
	return i.ToSchedulingDomainUserNameOutputWithContext(context.Background())
}

func (i *SchedulingDomainUserName) ToSchedulingDomainUserNameOutputWithContext(ctx context.Context) SchedulingDomainUserNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulingDomainUserNameOutput)
}

// SchedulingDomainUserNameArrayInput is an input type that accepts SchedulingDomainUserNameArray and SchedulingDomainUserNameArrayOutput values.
// You can construct a concrete instance of `SchedulingDomainUserNameArrayInput` via:
//
//	SchedulingDomainUserNameArray{ SchedulingDomainUserNameArgs{...} }
type SchedulingDomainUserNameArrayInput interface {
	pulumi.Input

	ToSchedulingDomainUserNameArrayOutput() SchedulingDomainUserNameArrayOutput
	ToSchedulingDomainUserNameArrayOutputWithContext(context.Context) SchedulingDomainUserNameArrayOutput
}

type SchedulingDomainUserNameArray []SchedulingDomainUserNameInput

func (SchedulingDomainUserNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SchedulingDomainUserName)(nil)).Elem()
}

func (i SchedulingDomainUserNameArray) ToSchedulingDomainUserNameArrayOutput() SchedulingDomainUserNameArrayOutput {
	return i.ToSchedulingDomainUserNameArrayOutputWithContext(context.Background())
}

func (i SchedulingDomainUserNameArray) ToSchedulingDomainUserNameArrayOutputWithContext(ctx context.Context) SchedulingDomainUserNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulingDomainUserNameArrayOutput)
}

// SchedulingDomainUserNameMapInput is an input type that accepts SchedulingDomainUserNameMap and SchedulingDomainUserNameMapOutput values.
// You can construct a concrete instance of `SchedulingDomainUserNameMapInput` via:
//
//	SchedulingDomainUserNameMap{ "key": SchedulingDomainUserNameArgs{...} }
type SchedulingDomainUserNameMapInput interface {
	pulumi.Input

	ToSchedulingDomainUserNameMapOutput() SchedulingDomainUserNameMapOutput
	ToSchedulingDomainUserNameMapOutputWithContext(context.Context) SchedulingDomainUserNameMapOutput
}

type SchedulingDomainUserNameMap map[string]SchedulingDomainUserNameInput

func (SchedulingDomainUserNameMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SchedulingDomainUserName)(nil)).Elem()
}

func (i SchedulingDomainUserNameMap) ToSchedulingDomainUserNameMapOutput() SchedulingDomainUserNameMapOutput {
	return i.ToSchedulingDomainUserNameMapOutputWithContext(context.Background())
}

func (i SchedulingDomainUserNameMap) ToSchedulingDomainUserNameMapOutputWithContext(ctx context.Context) SchedulingDomainUserNameMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulingDomainUserNameMapOutput)
}

type SchedulingDomainUserNameOutput struct{ *pulumi.OutputState }

func (SchedulingDomainUserNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchedulingDomainUserName)(nil)).Elem()
}

func (o SchedulingDomainUserNameOutput) ToSchedulingDomainUserNameOutput() SchedulingDomainUserNameOutput {
	return o
}

func (o SchedulingDomainUserNameOutput) ToSchedulingDomainUserNameOutputWithContext(ctx context.Context) SchedulingDomainUserNameOutput {
	return o
}

// user cname.
func (o SchedulingDomainUserNameOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *SchedulingDomainUserName) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// domain name.
func (o SchedulingDomainUserNameOutput) DomainUserName() pulumi.StringOutput {
	return o.ApplyT(func(v *SchedulingDomainUserName) pulumi.StringOutput { return v.DomainUserName }).(pulumi.StringOutput)
}

type SchedulingDomainUserNameArrayOutput struct{ *pulumi.OutputState }

func (SchedulingDomainUserNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SchedulingDomainUserName)(nil)).Elem()
}

func (o SchedulingDomainUserNameArrayOutput) ToSchedulingDomainUserNameArrayOutput() SchedulingDomainUserNameArrayOutput {
	return o
}

func (o SchedulingDomainUserNameArrayOutput) ToSchedulingDomainUserNameArrayOutputWithContext(ctx context.Context) SchedulingDomainUserNameArrayOutput {
	return o
}

func (o SchedulingDomainUserNameArrayOutput) Index(i pulumi.IntInput) SchedulingDomainUserNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SchedulingDomainUserName {
		return vs[0].([]*SchedulingDomainUserName)[vs[1].(int)]
	}).(SchedulingDomainUserNameOutput)
}

type SchedulingDomainUserNameMapOutput struct{ *pulumi.OutputState }

func (SchedulingDomainUserNameMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SchedulingDomainUserName)(nil)).Elem()
}

func (o SchedulingDomainUserNameMapOutput) ToSchedulingDomainUserNameMapOutput() SchedulingDomainUserNameMapOutput {
	return o
}

func (o SchedulingDomainUserNameMapOutput) ToSchedulingDomainUserNameMapOutputWithContext(ctx context.Context) SchedulingDomainUserNameMapOutput {
	return o
}

func (o SchedulingDomainUserNameMapOutput) MapIndex(k pulumi.StringInput) SchedulingDomainUserNameOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SchedulingDomainUserName {
		return vs[0].(map[string]*SchedulingDomainUserName)[vs[1].(string)]
	}).(SchedulingDomainUserNameOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchedulingDomainUserNameInput)(nil)).Elem(), &SchedulingDomainUserName{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchedulingDomainUserNameArrayInput)(nil)).Elem(), SchedulingDomainUserNameArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchedulingDomainUserNameMapInput)(nil)).Elem(), SchedulingDomainUserNameMap{})
	pulumi.RegisterOutputType(SchedulingDomainUserNameOutput{})
	pulumi.RegisterOutputType(SchedulingDomainUserNameArrayOutput{})
	pulumi.RegisterOutputType(SchedulingDomainUserNameMapOutput{})
}