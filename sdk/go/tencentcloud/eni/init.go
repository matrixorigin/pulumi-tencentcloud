// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eni

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
	case "tencentcloud:Eni/attachment:Attachment":
		r = &Attachment{}
	case "tencentcloud:Eni/instance:Instance":
		r = &Instance{}
	case "tencentcloud:Eni/ipv4Address:Ipv4Address":
		r = &Ipv4Address{}
	case "tencentcloud:Eni/ipv6Address:Ipv6Address":
		r = &Ipv6Address{}
	case "tencentcloud:Eni/sgAttachment:SgAttachment":
		r = &SgAttachment{}
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
		"Eni/attachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Eni/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Eni/ipv4Address",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Eni/ipv6Address",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Eni/sgAttachment",
		&module{version},
	)
}
