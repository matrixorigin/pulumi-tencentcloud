# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['InstanceAccountArgs', 'InstanceAccount']

@pulumi.input_type
class InstanceAccountArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 mongo_user_password: pulumi.Input[str],
                 password: pulumi.Input[str],
                 user_name: pulumi.Input[str],
                 auth_roles: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceAccountAuthRoleArgs']]]] = None,
                 user_desc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceAccount resource.
        :param pulumi.Input[str] instance_id: Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
        :param pulumi.Input[str] mongo_user_password: The password corresponding to the mongouser account. mongouser is the system default account, which is the password set
               when creating an instance.
        :param pulumi.Input[str] password: New account password. Password complexity requirements are as follows: character length range [8,32]. Contains at least
               letters, numbers and special characters (exclamation point!, at@, pound sign #, percent sign %, caret ^, asterisk *,
               parentheses (), underscore _).
        :param pulumi.Input[str] user_name: The new account name. Its format requirements are as follows: character range [1,32]. Characters in the range of [A,Z],
               [a,z], [1,9] as well as underscore _ and dash - can be input.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceAccountAuthRoleArgs']]] auth_roles: The read and write permission information of the account.
        :param pulumi.Input[str] user_desc: Account remarks.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "mongo_user_password", mongo_user_password)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "user_name", user_name)
        if auth_roles is not None:
            pulumi.set(__self__, "auth_roles", auth_roles)
        if user_desc is not None:
            pulumi.set(__self__, "user_desc", user_desc)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="mongoUserPassword")
    def mongo_user_password(self) -> pulumi.Input[str]:
        """
        The password corresponding to the mongouser account. mongouser is the system default account, which is the password set
        when creating an instance.
        """
        return pulumi.get(self, "mongo_user_password")

    @mongo_user_password.setter
    def mongo_user_password(self, value: pulumi.Input[str]):
        pulumi.set(self, "mongo_user_password", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        New account password. Password complexity requirements are as follows: character length range [8,32]. Contains at least
        letters, numbers and special characters (exclamation point!, at@, pound sign #, percent sign %, caret ^, asterisk *,
        parentheses (), underscore _).
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        The new account name. Its format requirements are as follows: character range [1,32]. Characters in the range of [A,Z],
        [a,z], [1,9] as well as underscore _ and dash - can be input.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter(name="authRoles")
    def auth_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceAccountAuthRoleArgs']]]]:
        """
        The read and write permission information of the account.
        """
        return pulumi.get(self, "auth_roles")

    @auth_roles.setter
    def auth_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceAccountAuthRoleArgs']]]]):
        pulumi.set(self, "auth_roles", value)

    @property
    @pulumi.getter(name="userDesc")
    def user_desc(self) -> Optional[pulumi.Input[str]]:
        """
        Account remarks.
        """
        return pulumi.get(self, "user_desc")

    @user_desc.setter
    def user_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_desc", value)


@pulumi.input_type
class _InstanceAccountState:
    def __init__(__self__, *,
                 auth_roles: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceAccountAuthRoleArgs']]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 mongo_user_password: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 user_desc: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceAccount resources.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceAccountAuthRoleArgs']]] auth_roles: The read and write permission information of the account.
        :param pulumi.Input[str] instance_id: Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
        :param pulumi.Input[str] mongo_user_password: The password corresponding to the mongouser account. mongouser is the system default account, which is the password set
               when creating an instance.
        :param pulumi.Input[str] password: New account password. Password complexity requirements are as follows: character length range [8,32]. Contains at least
               letters, numbers and special characters (exclamation point!, at@, pound sign #, percent sign %, caret ^, asterisk *,
               parentheses (), underscore _).
        :param pulumi.Input[str] user_desc: Account remarks.
        :param pulumi.Input[str] user_name: The new account name. Its format requirements are as follows: character range [1,32]. Characters in the range of [A,Z],
               [a,z], [1,9] as well as underscore _ and dash - can be input.
        """
        if auth_roles is not None:
            pulumi.set(__self__, "auth_roles", auth_roles)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if mongo_user_password is not None:
            pulumi.set(__self__, "mongo_user_password", mongo_user_password)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if user_desc is not None:
            pulumi.set(__self__, "user_desc", user_desc)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="authRoles")
    def auth_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceAccountAuthRoleArgs']]]]:
        """
        The read and write permission information of the account.
        """
        return pulumi.get(self, "auth_roles")

    @auth_roles.setter
    def auth_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceAccountAuthRoleArgs']]]]):
        pulumi.set(self, "auth_roles", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="mongoUserPassword")
    def mongo_user_password(self) -> Optional[pulumi.Input[str]]:
        """
        The password corresponding to the mongouser account. mongouser is the system default account, which is the password set
        when creating an instance.
        """
        return pulumi.get(self, "mongo_user_password")

    @mongo_user_password.setter
    def mongo_user_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mongo_user_password", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        New account password. Password complexity requirements are as follows: character length range [8,32]. Contains at least
        letters, numbers and special characters (exclamation point!, at@, pound sign #, percent sign %, caret ^, asterisk *,
        parentheses (), underscore _).
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="userDesc")
    def user_desc(self) -> Optional[pulumi.Input[str]]:
        """
        Account remarks.
        """
        return pulumi.get(self, "user_desc")

    @user_desc.setter
    def user_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_desc", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        The new account name. Its format requirements are as follows: character range [1,32]. Characters in the range of [A,Z],
        [a,z], [1,9] as well as underscore _ and dash - can be input.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class InstanceAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceAccountAuthRoleArgs']]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 mongo_user_password: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 user_desc: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a InstanceAccount resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceAccountAuthRoleArgs']]]] auth_roles: The read and write permission information of the account.
        :param pulumi.Input[str] instance_id: Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
        :param pulumi.Input[str] mongo_user_password: The password corresponding to the mongouser account. mongouser is the system default account, which is the password set
               when creating an instance.
        :param pulumi.Input[str] password: New account password. Password complexity requirements are as follows: character length range [8,32]. Contains at least
               letters, numbers and special characters (exclamation point!, at@, pound sign #, percent sign %, caret ^, asterisk *,
               parentheses (), underscore _).
        :param pulumi.Input[str] user_desc: Account remarks.
        :param pulumi.Input[str] user_name: The new account name. Its format requirements are as follows: character range [1,32]. Characters in the range of [A,Z],
               [a,z], [1,9] as well as underscore _ and dash - can be input.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a InstanceAccount resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param InstanceAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceAccountAuthRoleArgs']]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 mongo_user_password: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 user_desc: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceAccountArgs.__new__(InstanceAccountArgs)

            __props__.__dict__["auth_roles"] = auth_roles
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if mongo_user_password is None and not opts.urn:
                raise TypeError("Missing required property 'mongo_user_password'")
            __props__.__dict__["mongo_user_password"] = mongo_user_password
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = password
            __props__.__dict__["user_desc"] = user_desc
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
        super(InstanceAccount, __self__).__init__(
            'tencentcloud:Mongodb/instanceAccount:InstanceAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceAccountAuthRoleArgs']]]]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            mongo_user_password: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            user_desc: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'InstanceAccount':
        """
        Get an existing InstanceAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceAccountAuthRoleArgs']]]] auth_roles: The read and write permission information of the account.
        :param pulumi.Input[str] instance_id: Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
        :param pulumi.Input[str] mongo_user_password: The password corresponding to the mongouser account. mongouser is the system default account, which is the password set
               when creating an instance.
        :param pulumi.Input[str] password: New account password. Password complexity requirements are as follows: character length range [8,32]. Contains at least
               letters, numbers and special characters (exclamation point!, at@, pound sign #, percent sign %, caret ^, asterisk *,
               parentheses (), underscore _).
        :param pulumi.Input[str] user_desc: Account remarks.
        :param pulumi.Input[str] user_name: The new account name. Its format requirements are as follows: character range [1,32]. Characters in the range of [A,Z],
               [a,z], [1,9] as well as underscore _ and dash - can be input.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceAccountState.__new__(_InstanceAccountState)

        __props__.__dict__["auth_roles"] = auth_roles
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["mongo_user_password"] = mongo_user_password
        __props__.__dict__["password"] = password
        __props__.__dict__["user_desc"] = user_desc
        __props__.__dict__["user_name"] = user_name
        return InstanceAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authRoles")
    def auth_roles(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceAccountAuthRole']]]:
        """
        The read and write permission information of the account.
        """
        return pulumi.get(self, "auth_roles")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="mongoUserPassword")
    def mongo_user_password(self) -> pulumi.Output[str]:
        """
        The password corresponding to the mongouser account. mongouser is the system default account, which is the password set
        when creating an instance.
        """
        return pulumi.get(self, "mongo_user_password")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        New account password. Password complexity requirements are as follows: character length range [8,32]. Contains at least
        letters, numbers and special characters (exclamation point!, at@, pound sign #, percent sign %, caret ^, asterisk *,
        parentheses (), underscore _).
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="userDesc")
    def user_desc(self) -> pulumi.Output[Optional[str]]:
        """
        Account remarks.
        """
        return pulumi.get(self, "user_desc")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        The new account name. Its format requirements are as follows: character range [1,32]. Characters in the range of [A,Z],
        [a,z], [1,9] as well as underscore _ and dash - can be input.
        """
        return pulumi.get(self, "user_name")
