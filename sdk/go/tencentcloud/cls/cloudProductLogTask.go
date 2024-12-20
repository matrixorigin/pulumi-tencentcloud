// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cls

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudProductLogTask struct {
	pulumi.CustomResourceState

	// Cloud product identification, Values: CDS, CWP, CDB, TDSQL-C, MongoDB, TDStore, DCDB, MariaDB, PostgreSQL, BH, APIS.
	AssumerName pulumi.StringOutput `pulumi:"assumerName"`
	// Cloud product region. There are differences in the input format of different log types in different regions. Please
	// refer to the following example: - CDS(all log type): ap-guangzhou - CDB-AUDIT: gz - TDSQL-C-AUDIT: gz - MongoDB-AUDIT:
	// gz - MongoDB-SlowLog: ap-guangzhou - MongoDB-ErrorLog: ap-guangzhou - TDMYSQL-SLOW: gz - DCDB(all log type): gz -
	// MariaDB(all log type): gz - PostgreSQL(all log type): gz - BH(all log type): overseas-polaris(Domestic sites
	// overseas)/fsi-polaris(Domestic sites finance)/general-polaris(Domestic sites)/intl-sg-prod(International sites) -
	// APIS(all log type): gz.
	CloudProductRegion pulumi.StringOutput `pulumi:"cloudProductRegion"`
	// CLS target region.
	ClsRegion pulumi.StringOutput `pulumi:"clsRegion"`
	// Log configuration extension information, generally used to store additional log delivery configurations.
	Extend pulumi.StringPtrOutput `pulumi:"extend"`
	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Log type, Values: CDS-AUDIT, CDS-RISK, CDB-AUDIT, TDSQL-C-AUDIT, MongoDB-AUDIT, MongoDB-SlowLog, MongoDB-ErrorLog,
	// TDMYSQL-SLOW, DCDB-AUDIT, DCDB-SLOW, DCDB-ERROR, MariaDB-AUDIT, MariaDB-SLOW, MariaDB-ERROR, PostgreSQL-SLOW,
	// PostgreSQL-ERROR, PostgreSQL-AUDIT, BH-FILELOG, BH-COMMANDLOG, APIS-ACCESS.
	LogType pulumi.StringOutput `pulumi:"logType"`
	// Log set ID.
	LogsetId pulumi.StringOutput `pulumi:"logsetId"`
	// Log set name, it will be automatically created.
	LogsetName pulumi.StringPtrOutput `pulumi:"logsetName"`
	// Log theme ID.
	TopicId pulumi.StringOutput `pulumi:"topicId"`
	// The name of the log topic, it will be automatically created.
	TopicName pulumi.StringPtrOutput `pulumi:"topicName"`
}

// NewCloudProductLogTask registers a new resource with the given unique name, arguments, and options.
func NewCloudProductLogTask(ctx *pulumi.Context,
	name string, args *CloudProductLogTaskArgs, opts ...pulumi.ResourceOption) (*CloudProductLogTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssumerName == nil {
		return nil, errors.New("invalid value for required argument 'AssumerName'")
	}
	if args.CloudProductRegion == nil {
		return nil, errors.New("invalid value for required argument 'CloudProductRegion'")
	}
	if args.ClsRegion == nil {
		return nil, errors.New("invalid value for required argument 'ClsRegion'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.LogType == nil {
		return nil, errors.New("invalid value for required argument 'LogType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudProductLogTask
	err := ctx.RegisterResource("tencentcloud:Cls/cloudProductLogTask:CloudProductLogTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProductLogTask gets an existing CloudProductLogTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProductLogTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProductLogTaskState, opts ...pulumi.ResourceOption) (*CloudProductLogTask, error) {
	var resource CloudProductLogTask
	err := ctx.ReadResource("tencentcloud:Cls/cloudProductLogTask:CloudProductLogTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProductLogTask resources.
type cloudProductLogTaskState struct {
	// Cloud product identification, Values: CDS, CWP, CDB, TDSQL-C, MongoDB, TDStore, DCDB, MariaDB, PostgreSQL, BH, APIS.
	AssumerName *string `pulumi:"assumerName"`
	// Cloud product region. There are differences in the input format of different log types in different regions. Please
	// refer to the following example: - CDS(all log type): ap-guangzhou - CDB-AUDIT: gz - TDSQL-C-AUDIT: gz - MongoDB-AUDIT:
	// gz - MongoDB-SlowLog: ap-guangzhou - MongoDB-ErrorLog: ap-guangzhou - TDMYSQL-SLOW: gz - DCDB(all log type): gz -
	// MariaDB(all log type): gz - PostgreSQL(all log type): gz - BH(all log type): overseas-polaris(Domestic sites
	// overseas)/fsi-polaris(Domestic sites finance)/general-polaris(Domestic sites)/intl-sg-prod(International sites) -
	// APIS(all log type): gz.
	CloudProductRegion *string `pulumi:"cloudProductRegion"`
	// CLS target region.
	ClsRegion *string `pulumi:"clsRegion"`
	// Log configuration extension information, generally used to store additional log delivery configurations.
	Extend *string `pulumi:"extend"`
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Log type, Values: CDS-AUDIT, CDS-RISK, CDB-AUDIT, TDSQL-C-AUDIT, MongoDB-AUDIT, MongoDB-SlowLog, MongoDB-ErrorLog,
	// TDMYSQL-SLOW, DCDB-AUDIT, DCDB-SLOW, DCDB-ERROR, MariaDB-AUDIT, MariaDB-SLOW, MariaDB-ERROR, PostgreSQL-SLOW,
	// PostgreSQL-ERROR, PostgreSQL-AUDIT, BH-FILELOG, BH-COMMANDLOG, APIS-ACCESS.
	LogType *string `pulumi:"logType"`
	// Log set ID.
	LogsetId *string `pulumi:"logsetId"`
	// Log set name, it will be automatically created.
	LogsetName *string `pulumi:"logsetName"`
	// Log theme ID.
	TopicId *string `pulumi:"topicId"`
	// The name of the log topic, it will be automatically created.
	TopicName *string `pulumi:"topicName"`
}

type CloudProductLogTaskState struct {
	// Cloud product identification, Values: CDS, CWP, CDB, TDSQL-C, MongoDB, TDStore, DCDB, MariaDB, PostgreSQL, BH, APIS.
	AssumerName pulumi.StringPtrInput
	// Cloud product region. There are differences in the input format of different log types in different regions. Please
	// refer to the following example: - CDS(all log type): ap-guangzhou - CDB-AUDIT: gz - TDSQL-C-AUDIT: gz - MongoDB-AUDIT:
	// gz - MongoDB-SlowLog: ap-guangzhou - MongoDB-ErrorLog: ap-guangzhou - TDMYSQL-SLOW: gz - DCDB(all log type): gz -
	// MariaDB(all log type): gz - PostgreSQL(all log type): gz - BH(all log type): overseas-polaris(Domestic sites
	// overseas)/fsi-polaris(Domestic sites finance)/general-polaris(Domestic sites)/intl-sg-prod(International sites) -
	// APIS(all log type): gz.
	CloudProductRegion pulumi.StringPtrInput
	// CLS target region.
	ClsRegion pulumi.StringPtrInput
	// Log configuration extension information, generally used to store additional log delivery configurations.
	Extend pulumi.StringPtrInput
	// Instance ID.
	InstanceId pulumi.StringPtrInput
	// Log type, Values: CDS-AUDIT, CDS-RISK, CDB-AUDIT, TDSQL-C-AUDIT, MongoDB-AUDIT, MongoDB-SlowLog, MongoDB-ErrorLog,
	// TDMYSQL-SLOW, DCDB-AUDIT, DCDB-SLOW, DCDB-ERROR, MariaDB-AUDIT, MariaDB-SLOW, MariaDB-ERROR, PostgreSQL-SLOW,
	// PostgreSQL-ERROR, PostgreSQL-AUDIT, BH-FILELOG, BH-COMMANDLOG, APIS-ACCESS.
	LogType pulumi.StringPtrInput
	// Log set ID.
	LogsetId pulumi.StringPtrInput
	// Log set name, it will be automatically created.
	LogsetName pulumi.StringPtrInput
	// Log theme ID.
	TopicId pulumi.StringPtrInput
	// The name of the log topic, it will be automatically created.
	TopicName pulumi.StringPtrInput
}

func (CloudProductLogTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProductLogTaskState)(nil)).Elem()
}

type cloudProductLogTaskArgs struct {
	// Cloud product identification, Values: CDS, CWP, CDB, TDSQL-C, MongoDB, TDStore, DCDB, MariaDB, PostgreSQL, BH, APIS.
	AssumerName string `pulumi:"assumerName"`
	// Cloud product region. There are differences in the input format of different log types in different regions. Please
	// refer to the following example: - CDS(all log type): ap-guangzhou - CDB-AUDIT: gz - TDSQL-C-AUDIT: gz - MongoDB-AUDIT:
	// gz - MongoDB-SlowLog: ap-guangzhou - MongoDB-ErrorLog: ap-guangzhou - TDMYSQL-SLOW: gz - DCDB(all log type): gz -
	// MariaDB(all log type): gz - PostgreSQL(all log type): gz - BH(all log type): overseas-polaris(Domestic sites
	// overseas)/fsi-polaris(Domestic sites finance)/general-polaris(Domestic sites)/intl-sg-prod(International sites) -
	// APIS(all log type): gz.
	CloudProductRegion string `pulumi:"cloudProductRegion"`
	// CLS target region.
	ClsRegion string `pulumi:"clsRegion"`
	// Log configuration extension information, generally used to store additional log delivery configurations.
	Extend *string `pulumi:"extend"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Log type, Values: CDS-AUDIT, CDS-RISK, CDB-AUDIT, TDSQL-C-AUDIT, MongoDB-AUDIT, MongoDB-SlowLog, MongoDB-ErrorLog,
	// TDMYSQL-SLOW, DCDB-AUDIT, DCDB-SLOW, DCDB-ERROR, MariaDB-AUDIT, MariaDB-SLOW, MariaDB-ERROR, PostgreSQL-SLOW,
	// PostgreSQL-ERROR, PostgreSQL-AUDIT, BH-FILELOG, BH-COMMANDLOG, APIS-ACCESS.
	LogType string `pulumi:"logType"`
	// Log set name, it will be automatically created.
	LogsetName *string `pulumi:"logsetName"`
	// The name of the log topic, it will be automatically created.
	TopicName *string `pulumi:"topicName"`
}

// The set of arguments for constructing a CloudProductLogTask resource.
type CloudProductLogTaskArgs struct {
	// Cloud product identification, Values: CDS, CWP, CDB, TDSQL-C, MongoDB, TDStore, DCDB, MariaDB, PostgreSQL, BH, APIS.
	AssumerName pulumi.StringInput
	// Cloud product region. There are differences in the input format of different log types in different regions. Please
	// refer to the following example: - CDS(all log type): ap-guangzhou - CDB-AUDIT: gz - TDSQL-C-AUDIT: gz - MongoDB-AUDIT:
	// gz - MongoDB-SlowLog: ap-guangzhou - MongoDB-ErrorLog: ap-guangzhou - TDMYSQL-SLOW: gz - DCDB(all log type): gz -
	// MariaDB(all log type): gz - PostgreSQL(all log type): gz - BH(all log type): overseas-polaris(Domestic sites
	// overseas)/fsi-polaris(Domestic sites finance)/general-polaris(Domestic sites)/intl-sg-prod(International sites) -
	// APIS(all log type): gz.
	CloudProductRegion pulumi.StringInput
	// CLS target region.
	ClsRegion pulumi.StringInput
	// Log configuration extension information, generally used to store additional log delivery configurations.
	Extend pulumi.StringPtrInput
	// Instance ID.
	InstanceId pulumi.StringInput
	// Log type, Values: CDS-AUDIT, CDS-RISK, CDB-AUDIT, TDSQL-C-AUDIT, MongoDB-AUDIT, MongoDB-SlowLog, MongoDB-ErrorLog,
	// TDMYSQL-SLOW, DCDB-AUDIT, DCDB-SLOW, DCDB-ERROR, MariaDB-AUDIT, MariaDB-SLOW, MariaDB-ERROR, PostgreSQL-SLOW,
	// PostgreSQL-ERROR, PostgreSQL-AUDIT, BH-FILELOG, BH-COMMANDLOG, APIS-ACCESS.
	LogType pulumi.StringInput
	// Log set name, it will be automatically created.
	LogsetName pulumi.StringPtrInput
	// The name of the log topic, it will be automatically created.
	TopicName pulumi.StringPtrInput
}

func (CloudProductLogTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProductLogTaskArgs)(nil)).Elem()
}

type CloudProductLogTaskInput interface {
	pulumi.Input

	ToCloudProductLogTaskOutput() CloudProductLogTaskOutput
	ToCloudProductLogTaskOutputWithContext(ctx context.Context) CloudProductLogTaskOutput
}

func (*CloudProductLogTask) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProductLogTask)(nil)).Elem()
}

func (i *CloudProductLogTask) ToCloudProductLogTaskOutput() CloudProductLogTaskOutput {
	return i.ToCloudProductLogTaskOutputWithContext(context.Background())
}

func (i *CloudProductLogTask) ToCloudProductLogTaskOutputWithContext(ctx context.Context) CloudProductLogTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProductLogTaskOutput)
}

// CloudProductLogTaskArrayInput is an input type that accepts CloudProductLogTaskArray and CloudProductLogTaskArrayOutput values.
// You can construct a concrete instance of `CloudProductLogTaskArrayInput` via:
//
//	CloudProductLogTaskArray{ CloudProductLogTaskArgs{...} }
type CloudProductLogTaskArrayInput interface {
	pulumi.Input

	ToCloudProductLogTaskArrayOutput() CloudProductLogTaskArrayOutput
	ToCloudProductLogTaskArrayOutputWithContext(context.Context) CloudProductLogTaskArrayOutput
}

type CloudProductLogTaskArray []CloudProductLogTaskInput

func (CloudProductLogTaskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProductLogTask)(nil)).Elem()
}

func (i CloudProductLogTaskArray) ToCloudProductLogTaskArrayOutput() CloudProductLogTaskArrayOutput {
	return i.ToCloudProductLogTaskArrayOutputWithContext(context.Background())
}

func (i CloudProductLogTaskArray) ToCloudProductLogTaskArrayOutputWithContext(ctx context.Context) CloudProductLogTaskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProductLogTaskArrayOutput)
}

// CloudProductLogTaskMapInput is an input type that accepts CloudProductLogTaskMap and CloudProductLogTaskMapOutput values.
// You can construct a concrete instance of `CloudProductLogTaskMapInput` via:
//
//	CloudProductLogTaskMap{ "key": CloudProductLogTaskArgs{...} }
type CloudProductLogTaskMapInput interface {
	pulumi.Input

	ToCloudProductLogTaskMapOutput() CloudProductLogTaskMapOutput
	ToCloudProductLogTaskMapOutputWithContext(context.Context) CloudProductLogTaskMapOutput
}

type CloudProductLogTaskMap map[string]CloudProductLogTaskInput

func (CloudProductLogTaskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProductLogTask)(nil)).Elem()
}

func (i CloudProductLogTaskMap) ToCloudProductLogTaskMapOutput() CloudProductLogTaskMapOutput {
	return i.ToCloudProductLogTaskMapOutputWithContext(context.Background())
}

func (i CloudProductLogTaskMap) ToCloudProductLogTaskMapOutputWithContext(ctx context.Context) CloudProductLogTaskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProductLogTaskMapOutput)
}

type CloudProductLogTaskOutput struct{ *pulumi.OutputState }

func (CloudProductLogTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProductLogTask)(nil)).Elem()
}

func (o CloudProductLogTaskOutput) ToCloudProductLogTaskOutput() CloudProductLogTaskOutput {
	return o
}

func (o CloudProductLogTaskOutput) ToCloudProductLogTaskOutputWithContext(ctx context.Context) CloudProductLogTaskOutput {
	return o
}

// Cloud product identification, Values: CDS, CWP, CDB, TDSQL-C, MongoDB, TDStore, DCDB, MariaDB, PostgreSQL, BH, APIS.
func (o CloudProductLogTaskOutput) AssumerName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProductLogTask) pulumi.StringOutput { return v.AssumerName }).(pulumi.StringOutput)
}

// Cloud product region. There are differences in the input format of different log types in different regions. Please
// refer to the following example: - CDS(all log type): ap-guangzhou - CDB-AUDIT: gz - TDSQL-C-AUDIT: gz - MongoDB-AUDIT:
// gz - MongoDB-SlowLog: ap-guangzhou - MongoDB-ErrorLog: ap-guangzhou - TDMYSQL-SLOW: gz - DCDB(all log type): gz -
// MariaDB(all log type): gz - PostgreSQL(all log type): gz - BH(all log type): overseas-polaris(Domestic sites
// overseas)/fsi-polaris(Domestic sites finance)/general-polaris(Domestic sites)/intl-sg-prod(International sites) -
// APIS(all log type): gz.
func (o CloudProductLogTaskOutput) CloudProductRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProductLogTask) pulumi.StringOutput { return v.CloudProductRegion }).(pulumi.StringOutput)
}

// CLS target region.
func (o CloudProductLogTaskOutput) ClsRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProductLogTask) pulumi.StringOutput { return v.ClsRegion }).(pulumi.StringOutput)
}

// Log configuration extension information, generally used to store additional log delivery configurations.
func (o CloudProductLogTaskOutput) Extend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProductLogTask) pulumi.StringPtrOutput { return v.Extend }).(pulumi.StringPtrOutput)
}

// Instance ID.
func (o CloudProductLogTaskOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProductLogTask) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Log type, Values: CDS-AUDIT, CDS-RISK, CDB-AUDIT, TDSQL-C-AUDIT, MongoDB-AUDIT, MongoDB-SlowLog, MongoDB-ErrorLog,
// TDMYSQL-SLOW, DCDB-AUDIT, DCDB-SLOW, DCDB-ERROR, MariaDB-AUDIT, MariaDB-SLOW, MariaDB-ERROR, PostgreSQL-SLOW,
// PostgreSQL-ERROR, PostgreSQL-AUDIT, BH-FILELOG, BH-COMMANDLOG, APIS-ACCESS.
func (o CloudProductLogTaskOutput) LogType() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProductLogTask) pulumi.StringOutput { return v.LogType }).(pulumi.StringOutput)
}

// Log set ID.
func (o CloudProductLogTaskOutput) LogsetId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProductLogTask) pulumi.StringOutput { return v.LogsetId }).(pulumi.StringOutput)
}

// Log set name, it will be automatically created.
func (o CloudProductLogTaskOutput) LogsetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProductLogTask) pulumi.StringPtrOutput { return v.LogsetName }).(pulumi.StringPtrOutput)
}

// Log theme ID.
func (o CloudProductLogTaskOutput) TopicId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProductLogTask) pulumi.StringOutput { return v.TopicId }).(pulumi.StringOutput)
}

// The name of the log topic, it will be automatically created.
func (o CloudProductLogTaskOutput) TopicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudProductLogTask) pulumi.StringPtrOutput { return v.TopicName }).(pulumi.StringPtrOutput)
}

type CloudProductLogTaskArrayOutput struct{ *pulumi.OutputState }

func (CloudProductLogTaskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProductLogTask)(nil)).Elem()
}

func (o CloudProductLogTaskArrayOutput) ToCloudProductLogTaskArrayOutput() CloudProductLogTaskArrayOutput {
	return o
}

func (o CloudProductLogTaskArrayOutput) ToCloudProductLogTaskArrayOutputWithContext(ctx context.Context) CloudProductLogTaskArrayOutput {
	return o
}

func (o CloudProductLogTaskArrayOutput) Index(i pulumi.IntInput) CloudProductLogTaskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudProductLogTask {
		return vs[0].([]*CloudProductLogTask)[vs[1].(int)]
	}).(CloudProductLogTaskOutput)
}

type CloudProductLogTaskMapOutput struct{ *pulumi.OutputState }

func (CloudProductLogTaskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProductLogTask)(nil)).Elem()
}

func (o CloudProductLogTaskMapOutput) ToCloudProductLogTaskMapOutput() CloudProductLogTaskMapOutput {
	return o
}

func (o CloudProductLogTaskMapOutput) ToCloudProductLogTaskMapOutputWithContext(ctx context.Context) CloudProductLogTaskMapOutput {
	return o
}

func (o CloudProductLogTaskMapOutput) MapIndex(k pulumi.StringInput) CloudProductLogTaskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudProductLogTask {
		return vs[0].(map[string]*CloudProductLogTask)[vs[1].(string)]
	}).(CloudProductLogTaskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProductLogTaskInput)(nil)).Elem(), &CloudProductLogTask{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProductLogTaskArrayInput)(nil)).Elem(), CloudProductLogTaskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProductLogTaskMapInput)(nil)).Elem(), CloudProductLogTaskMap{})
	pulumi.RegisterOutputType(CloudProductLogTaskOutput{})
	pulumi.RegisterOutputType(CloudProductLogTaskArrayOutput{})
	pulumi.RegisterOutputType(CloudProductLogTaskMapOutput{})
}
