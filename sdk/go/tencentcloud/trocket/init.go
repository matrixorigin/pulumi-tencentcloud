// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package trocket

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
	case "tencentcloud:Trocket/rocketmqConsumerGroup:RocketmqConsumerGroup":
		r = &RocketmqConsumerGroup{}
	case "tencentcloud:Trocket/rocketmqInstance:RocketmqInstance":
		r = &RocketmqInstance{}
	case "tencentcloud:Trocket/rocketmqRole:RocketmqRole":
		r = &RocketmqRole{}
	case "tencentcloud:Trocket/rocketmqTopic:RocketmqTopic":
		r = &RocketmqTopic{}
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
		"Trocket/rocketmqConsumerGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Trocket/rocketmqInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Trocket/rocketmqRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Trocket/rocketmqTopic",
		&module{version},
	)
}