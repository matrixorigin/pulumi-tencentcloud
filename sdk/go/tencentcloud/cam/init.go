// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

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
	case "tencentcloud:Cam/accessKey:AccessKey":
		r = &AccessKey{}
	case "tencentcloud:Cam/group:Group":
		r = &Group{}
	case "tencentcloud:Cam/groupMembership:GroupMembership":
		r = &GroupMembership{}
	case "tencentcloud:Cam/groupPolicyAttachment:GroupPolicyAttachment":
		r = &GroupPolicyAttachment{}
	case "tencentcloud:Cam/mfaFlag:MfaFlag":
		r = &MfaFlag{}
	case "tencentcloud:Cam/oidcSso:OidcSso":
		r = &OidcSso{}
	case "tencentcloud:Cam/policy:Policy":
		r = &Policy{}
	case "tencentcloud:Cam/policyByName:PolicyByName":
		r = &PolicyByName{}
	case "tencentcloud:Cam/policyVersion:PolicyVersion":
		r = &PolicyVersion{}
	case "tencentcloud:Cam/role:Role":
		r = &Role{}
	case "tencentcloud:Cam/roleByName:RoleByName":
		r = &RoleByName{}
	case "tencentcloud:Cam/rolePermissionBoundaryAttachment:RolePermissionBoundaryAttachment":
		r = &RolePermissionBoundaryAttachment{}
	case "tencentcloud:Cam/rolePolicyAttachment:RolePolicyAttachment":
		r = &RolePolicyAttachment{}
	case "tencentcloud:Cam/rolePolicyAttachmentByName:RolePolicyAttachmentByName":
		r = &RolePolicyAttachmentByName{}
	case "tencentcloud:Cam/roleSso:RoleSso":
		r = &RoleSso{}
	case "tencentcloud:Cam/samlProvider:SamlProvider":
		r = &SamlProvider{}
	case "tencentcloud:Cam/serviceLinkedRole:ServiceLinkedRole":
		r = &ServiceLinkedRole{}
	case "tencentcloud:Cam/setPolicyVersionConfig:SetPolicyVersionConfig":
		r = &SetPolicyVersionConfig{}
	case "tencentcloud:Cam/tagRoleAttachment:TagRoleAttachment":
		r = &TagRoleAttachment{}
	case "tencentcloud:Cam/user:User":
		r = &User{}
	case "tencentcloud:Cam/userPermissionBoundaryAttachment:UserPermissionBoundaryAttachment":
		r = &UserPermissionBoundaryAttachment{}
	case "tencentcloud:Cam/userPolicyAttachment:UserPolicyAttachment":
		r = &UserPolicyAttachment{}
	case "tencentcloud:Cam/userSamlConfig:UserSamlConfig":
		r = &UserSamlConfig{}
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
		"Cam/accessKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/groupMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/groupPolicyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/mfaFlag",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/oidcSso",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/policyByName",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/policyVersion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/role",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/roleByName",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/rolePermissionBoundaryAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/rolePolicyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/rolePolicyAttachmentByName",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/roleSso",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/samlProvider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/serviceLinkedRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/setPolicyVersionConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/tagRoleAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/userPermissionBoundaryAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/userPolicyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cam/userSamlConfig",
		&module{version},
	)
}