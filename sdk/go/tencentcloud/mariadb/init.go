// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

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
	case "tencentcloud:Mariadb/account:Account":
		r = &Account{}
	case "tencentcloud:Mariadb/accountPrivileges:AccountPrivileges":
		r = &AccountPrivileges{}
	case "tencentcloud:Mariadb/backupTime:BackupTime":
		r = &BackupTime{}
	case "tencentcloud:Mariadb/cancelDcnJob:CancelDcnJob":
		r = &CancelDcnJob{}
	case "tencentcloud:Mariadb/dedicatedclusterDbInstance:DedicatedclusterDbInstance":
		r = &DedicatedclusterDbInstance{}
	case "tencentcloud:Mariadb/encryptAttributes:EncryptAttributes":
		r = &EncryptAttributes{}
	case "tencentcloud:Mariadb/flushBinlog:FlushBinlog":
		r = &FlushBinlog{}
	case "tencentcloud:Mariadb/hourDbInstance:HourDbInstance":
		r = &HourDbInstance{}
	case "tencentcloud:Mariadb/instance:Instance":
		r = &Instance{}
	case "tencentcloud:Mariadb/instanceConfig:InstanceConfig":
		r = &InstanceConfig{}
	case "tencentcloud:Mariadb/logFileRetentionPeriod:LogFileRetentionPeriod":
		r = &LogFileRetentionPeriod{}
	case "tencentcloud:Mariadb/operateHourDbInstance:OperateHourDbInstance":
		r = &OperateHourDbInstance{}
	case "tencentcloud:Mariadb/parameters:Parameters":
		r = &Parameters{}
	case "tencentcloud:Mariadb/renewInstance:RenewInstance":
		r = &RenewInstance{}
	case "tencentcloud:Mariadb/restartInstance:RestartInstance":
		r = &RestartInstance{}
	case "tencentcloud:Mariadb/securityGroups:SecurityGroups":
		r = &SecurityGroups{}
	case "tencentcloud:Mariadb/switchHa:SwitchHa":
		r = &SwitchHa{}
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
		"Mariadb/account",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/accountPrivileges",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/backupTime",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/cancelDcnJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/dedicatedclusterDbInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/encryptAttributes",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/flushBinlog",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/hourDbInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/instanceConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/logFileRetentionPeriod",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/operateHourDbInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/parameters",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/renewInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/restartInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/securityGroups",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Mariadb/switchHa",
		&module{version},
	)
}