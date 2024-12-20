// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceBackupRule struct {
	pulumi.CustomResourceState

	// Set automatic backup method. Valid values: - 0: Logical backup; - 1: Physical backup; - 3: Snapshot backup (supported
	// only in cloud disk version).
	BackupMethod pulumi.IntOutput `pulumi:"backupMethod"`
	// Specify the number of days to save backup data. The default is 7 days, and the support settings are 7, 30, 90, 180, 365.
	BackupRetentionPeriod pulumi.IntOutput `pulumi:"backupRetentionPeriod"`
	// Set the start time for automatic backup. The value range is: [0,23]. For example, setting this parameter to 2 means that
	// backup starts at 02:00.
	BackupTime pulumi.IntOutput `pulumi:"backupTime"`
	// Instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewInstanceBackupRule registers a new resource with the given unique name, arguments, and options.
func NewInstanceBackupRule(ctx *pulumi.Context,
	name string, args *InstanceBackupRuleArgs, opts ...pulumi.ResourceOption) (*InstanceBackupRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupMethod == nil {
		return nil, errors.New("invalid value for required argument 'BackupMethod'")
	}
	if args.BackupTime == nil {
		return nil, errors.New("invalid value for required argument 'BackupTime'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceBackupRule
	err := ctx.RegisterResource("tencentcloud:Mongodb/instanceBackupRule:InstanceBackupRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceBackupRule gets an existing InstanceBackupRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceBackupRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceBackupRuleState, opts ...pulumi.ResourceOption) (*InstanceBackupRule, error) {
	var resource InstanceBackupRule
	err := ctx.ReadResource("tencentcloud:Mongodb/instanceBackupRule:InstanceBackupRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceBackupRule resources.
type instanceBackupRuleState struct {
	// Set automatic backup method. Valid values: - 0: Logical backup; - 1: Physical backup; - 3: Snapshot backup (supported
	// only in cloud disk version).
	BackupMethod *int `pulumi:"backupMethod"`
	// Specify the number of days to save backup data. The default is 7 days, and the support settings are 7, 30, 90, 180, 365.
	BackupRetentionPeriod *int `pulumi:"backupRetentionPeriod"`
	// Set the start time for automatic backup. The value range is: [0,23]. For example, setting this parameter to 2 means that
	// backup starts at 02:00.
	BackupTime *int `pulumi:"backupTime"`
	// Instance id.
	InstanceId *string `pulumi:"instanceId"`
}

type InstanceBackupRuleState struct {
	// Set automatic backup method. Valid values: - 0: Logical backup; - 1: Physical backup; - 3: Snapshot backup (supported
	// only in cloud disk version).
	BackupMethod pulumi.IntPtrInput
	// Specify the number of days to save backup data. The default is 7 days, and the support settings are 7, 30, 90, 180, 365.
	BackupRetentionPeriod pulumi.IntPtrInput
	// Set the start time for automatic backup. The value range is: [0,23]. For example, setting this parameter to 2 means that
	// backup starts at 02:00.
	BackupTime pulumi.IntPtrInput
	// Instance id.
	InstanceId pulumi.StringPtrInput
}

func (InstanceBackupRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceBackupRuleState)(nil)).Elem()
}

type instanceBackupRuleArgs struct {
	// Set automatic backup method. Valid values: - 0: Logical backup; - 1: Physical backup; - 3: Snapshot backup (supported
	// only in cloud disk version).
	BackupMethod int `pulumi:"backupMethod"`
	// Specify the number of days to save backup data. The default is 7 days, and the support settings are 7, 30, 90, 180, 365.
	BackupRetentionPeriod *int `pulumi:"backupRetentionPeriod"`
	// Set the start time for automatic backup. The value range is: [0,23]. For example, setting this parameter to 2 means that
	// backup starts at 02:00.
	BackupTime int `pulumi:"backupTime"`
	// Instance id.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a InstanceBackupRule resource.
type InstanceBackupRuleArgs struct {
	// Set automatic backup method. Valid values: - 0: Logical backup; - 1: Physical backup; - 3: Snapshot backup (supported
	// only in cloud disk version).
	BackupMethod pulumi.IntInput
	// Specify the number of days to save backup data. The default is 7 days, and the support settings are 7, 30, 90, 180, 365.
	BackupRetentionPeriod pulumi.IntPtrInput
	// Set the start time for automatic backup. The value range is: [0,23]. For example, setting this parameter to 2 means that
	// backup starts at 02:00.
	BackupTime pulumi.IntInput
	// Instance id.
	InstanceId pulumi.StringInput
}

func (InstanceBackupRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceBackupRuleArgs)(nil)).Elem()
}

type InstanceBackupRuleInput interface {
	pulumi.Input

	ToInstanceBackupRuleOutput() InstanceBackupRuleOutput
	ToInstanceBackupRuleOutputWithContext(ctx context.Context) InstanceBackupRuleOutput
}

func (*InstanceBackupRule) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceBackupRule)(nil)).Elem()
}

func (i *InstanceBackupRule) ToInstanceBackupRuleOutput() InstanceBackupRuleOutput {
	return i.ToInstanceBackupRuleOutputWithContext(context.Background())
}

func (i *InstanceBackupRule) ToInstanceBackupRuleOutputWithContext(ctx context.Context) InstanceBackupRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceBackupRuleOutput)
}

// InstanceBackupRuleArrayInput is an input type that accepts InstanceBackupRuleArray and InstanceBackupRuleArrayOutput values.
// You can construct a concrete instance of `InstanceBackupRuleArrayInput` via:
//
//	InstanceBackupRuleArray{ InstanceBackupRuleArgs{...} }
type InstanceBackupRuleArrayInput interface {
	pulumi.Input

	ToInstanceBackupRuleArrayOutput() InstanceBackupRuleArrayOutput
	ToInstanceBackupRuleArrayOutputWithContext(context.Context) InstanceBackupRuleArrayOutput
}

type InstanceBackupRuleArray []InstanceBackupRuleInput

func (InstanceBackupRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceBackupRule)(nil)).Elem()
}

func (i InstanceBackupRuleArray) ToInstanceBackupRuleArrayOutput() InstanceBackupRuleArrayOutput {
	return i.ToInstanceBackupRuleArrayOutputWithContext(context.Background())
}

func (i InstanceBackupRuleArray) ToInstanceBackupRuleArrayOutputWithContext(ctx context.Context) InstanceBackupRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceBackupRuleArrayOutput)
}

// InstanceBackupRuleMapInput is an input type that accepts InstanceBackupRuleMap and InstanceBackupRuleMapOutput values.
// You can construct a concrete instance of `InstanceBackupRuleMapInput` via:
//
//	InstanceBackupRuleMap{ "key": InstanceBackupRuleArgs{...} }
type InstanceBackupRuleMapInput interface {
	pulumi.Input

	ToInstanceBackupRuleMapOutput() InstanceBackupRuleMapOutput
	ToInstanceBackupRuleMapOutputWithContext(context.Context) InstanceBackupRuleMapOutput
}

type InstanceBackupRuleMap map[string]InstanceBackupRuleInput

func (InstanceBackupRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceBackupRule)(nil)).Elem()
}

func (i InstanceBackupRuleMap) ToInstanceBackupRuleMapOutput() InstanceBackupRuleMapOutput {
	return i.ToInstanceBackupRuleMapOutputWithContext(context.Background())
}

func (i InstanceBackupRuleMap) ToInstanceBackupRuleMapOutputWithContext(ctx context.Context) InstanceBackupRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceBackupRuleMapOutput)
}

type InstanceBackupRuleOutput struct{ *pulumi.OutputState }

func (InstanceBackupRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceBackupRule)(nil)).Elem()
}

func (o InstanceBackupRuleOutput) ToInstanceBackupRuleOutput() InstanceBackupRuleOutput {
	return o
}

func (o InstanceBackupRuleOutput) ToInstanceBackupRuleOutputWithContext(ctx context.Context) InstanceBackupRuleOutput {
	return o
}

// Set automatic backup method. Valid values: - 0: Logical backup; - 1: Physical backup; - 3: Snapshot backup (supported
// only in cloud disk version).
func (o InstanceBackupRuleOutput) BackupMethod() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceBackupRule) pulumi.IntOutput { return v.BackupMethod }).(pulumi.IntOutput)
}

// Specify the number of days to save backup data. The default is 7 days, and the support settings are 7, 30, 90, 180, 365.
func (o InstanceBackupRuleOutput) BackupRetentionPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceBackupRule) pulumi.IntOutput { return v.BackupRetentionPeriod }).(pulumi.IntOutput)
}

// Set the start time for automatic backup. The value range is: [0,23]. For example, setting this parameter to 2 means that
// backup starts at 02:00.
func (o InstanceBackupRuleOutput) BackupTime() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceBackupRule) pulumi.IntOutput { return v.BackupTime }).(pulumi.IntOutput)
}

// Instance id.
func (o InstanceBackupRuleOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceBackupRule) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type InstanceBackupRuleArrayOutput struct{ *pulumi.OutputState }

func (InstanceBackupRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceBackupRule)(nil)).Elem()
}

func (o InstanceBackupRuleArrayOutput) ToInstanceBackupRuleArrayOutput() InstanceBackupRuleArrayOutput {
	return o
}

func (o InstanceBackupRuleArrayOutput) ToInstanceBackupRuleArrayOutputWithContext(ctx context.Context) InstanceBackupRuleArrayOutput {
	return o
}

func (o InstanceBackupRuleArrayOutput) Index(i pulumi.IntInput) InstanceBackupRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceBackupRule {
		return vs[0].([]*InstanceBackupRule)[vs[1].(int)]
	}).(InstanceBackupRuleOutput)
}

type InstanceBackupRuleMapOutput struct{ *pulumi.OutputState }

func (InstanceBackupRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceBackupRule)(nil)).Elem()
}

func (o InstanceBackupRuleMapOutput) ToInstanceBackupRuleMapOutput() InstanceBackupRuleMapOutput {
	return o
}

func (o InstanceBackupRuleMapOutput) ToInstanceBackupRuleMapOutputWithContext(ctx context.Context) InstanceBackupRuleMapOutput {
	return o
}

func (o InstanceBackupRuleMapOutput) MapIndex(k pulumi.StringInput) InstanceBackupRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceBackupRule {
		return vs[0].(map[string]*InstanceBackupRule)[vs[1].(string)]
	}).(InstanceBackupRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceBackupRuleInput)(nil)).Elem(), &InstanceBackupRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceBackupRuleArrayInput)(nil)).Elem(), InstanceBackupRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceBackupRuleMapInput)(nil)).Elem(), InstanceBackupRuleMap{})
	pulumi.RegisterOutputType(InstanceBackupRuleOutput{})
	pulumi.RegisterOutputType(InstanceBackupRuleArrayOutput{})
	pulumi.RegisterOutputType(InstanceBackupRuleMapOutput{})
}
