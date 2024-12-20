// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cls

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "tencentcloud:Cls/alarm:Alarm":
		r = &Alarm{}
	case "tencentcloud:Cls/alarmNotice:AlarmNotice":
		r = &AlarmNotice{}
	case "tencentcloud:Cls/ckafkaConsumer:CkafkaConsumer":
		r = &CkafkaConsumer{}
	case "tencentcloud:Cls/cloudProductLogTask:CloudProductLogTask":
		r = &CloudProductLogTask{}
	case "tencentcloud:Cls/config:Config":
		r = &Config{}
	case "tencentcloud:Cls/configAttachment:ConfigAttachment":
		r = &ConfigAttachment{}
	case "tencentcloud:Cls/configExtra:ConfigExtra":
		r = &ConfigExtra{}
	case "tencentcloud:Cls/cosRecharge:CosRecharge":
		r = &CosRecharge{}
	case "tencentcloud:Cls/cosShipper:CosShipper":
		r = &CosShipper{}
	case "tencentcloud:Cls/dataTransform:DataTransform":
		r = &DataTransform{}
	case "tencentcloud:Cls/export:Export":
		r = &Export{}
	case "tencentcloud:Cls/index:Index":
		r = &Index{}
	case "tencentcloud:Cls/kafkaRecharge:KafkaRecharge":
		r = &KafkaRecharge{}
	case "tencentcloud:Cls/logset:Logset":
		r = &Logset{}
	case "tencentcloud:Cls/machineGroup:MachineGroup":
		r = &MachineGroup{}
	case "tencentcloud:Cls/noticeContent:NoticeContent":
		r = &NoticeContent{}
	case "tencentcloud:Cls/scheduledSql:ScheduledSql":
		r = &ScheduledSql{}
	case "tencentcloud:Cls/topic:Topic":
		r = &Topic{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/alarm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/alarmNotice",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/ckafkaConsumer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/cloudProductLogTask",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/config",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/configAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/configExtra",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/cosRecharge",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/cosShipper",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/dataTransform",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/export",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/index",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/kafkaRecharge",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/logset",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/machineGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/noticeContent",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/scheduledSql",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cls/topic",
		&module{version},
	)
}
