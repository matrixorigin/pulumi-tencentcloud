# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EmbedTokenApplyArgs', 'EmbedTokenApply']

@pulumi.input_type
class EmbedTokenApplyArgs:
    def __init__(__self__, *,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 page_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 ticket_num: Optional[pulumi.Input[int]] = None,
                 user_corp_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EmbedTokenApply resource.
        :param pulumi.Input[str] expire_time: Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
        :param pulumi.Input[int] page_id: Sharing page id, this is empty value 0 when embedding the board.
        :param pulumi.Input[int] project_id: Share project id.
        :param pulumi.Input[str] scope: Page means embedding the page, and panel means embedding the entire board.
        :param pulumi.Input[int] ticket_num: Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
        :param pulumi.Input[str] user_corp_id: User enterprise ID (for multi-user only).
        :param pulumi.Input[str] user_id: UserId (for multi-user only).
        """
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if page_id is not None:
            pulumi.set(__self__, "page_id", page_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if ticket_num is not None:
            pulumi.set(__self__, "ticket_num", ticket_num)
        if user_corp_id is not None:
            pulumi.set(__self__, "user_corp_id", user_corp_id)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter(name="pageId")
    def page_id(self) -> Optional[pulumi.Input[int]]:
        """
        Sharing page id, this is empty value 0 when embedding the board.
        """
        return pulumi.get(self, "page_id")

    @page_id.setter
    def page_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "page_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Share project id.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Page means embedding the page, and panel means embedding the entire board.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="ticketNum")
    def ticket_num(self) -> Optional[pulumi.Input[int]]:
        """
        Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
        """
        return pulumi.get(self, "ticket_num")

    @ticket_num.setter
    def ticket_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ticket_num", value)

    @property
    @pulumi.getter(name="userCorpId")
    def user_corp_id(self) -> Optional[pulumi.Input[str]]:
        """
        User enterprise ID (for multi-user only).
        """
        return pulumi.get(self, "user_corp_id")

    @user_corp_id.setter
    def user_corp_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_corp_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        UserId (for multi-user only).
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


@pulumi.input_type
class _EmbedTokenApplyState:
    def __init__(__self__, *,
                 bi_token: Optional[pulumi.Input[str]] = None,
                 create_at: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 page_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 ticket_num: Optional[pulumi.Input[int]] = None,
                 udpate_at: Optional[pulumi.Input[str]] = None,
                 user_corp_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EmbedTokenApply resources.
        :param pulumi.Input[str] bi_token: Create the generated token.
        :param pulumi.Input[str] create_at: Create time.
        :param pulumi.Input[str] expire_time: Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
        :param pulumi.Input[int] page_id: Sharing page id, this is empty value 0 when embedding the board.
        :param pulumi.Input[int] project_id: Share project id.
        :param pulumi.Input[str] scope: Page means embedding the page, and panel means embedding the entire board.
        :param pulumi.Input[int] ticket_num: Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
        :param pulumi.Input[str] udpate_at: Upadte time.
        :param pulumi.Input[str] user_corp_id: User enterprise ID (for multi-user only).
        :param pulumi.Input[str] user_id: UserId (for multi-user only).
        """
        if bi_token is not None:
            pulumi.set(__self__, "bi_token", bi_token)
        if create_at is not None:
            pulumi.set(__self__, "create_at", create_at)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if page_id is not None:
            pulumi.set(__self__, "page_id", page_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if ticket_num is not None:
            pulumi.set(__self__, "ticket_num", ticket_num)
        if udpate_at is not None:
            pulumi.set(__self__, "udpate_at", udpate_at)
        if user_corp_id is not None:
            pulumi.set(__self__, "user_corp_id", user_corp_id)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="biToken")
    def bi_token(self) -> Optional[pulumi.Input[str]]:
        """
        Create the generated token.
        """
        return pulumi.get(self, "bi_token")

    @bi_token.setter
    def bi_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bi_token", value)

    @property
    @pulumi.getter(name="createAt")
    def create_at(self) -> Optional[pulumi.Input[str]]:
        """
        Create time.
        """
        return pulumi.get(self, "create_at")

    @create_at.setter
    def create_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_at", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter(name="pageId")
    def page_id(self) -> Optional[pulumi.Input[int]]:
        """
        Sharing page id, this is empty value 0 when embedding the board.
        """
        return pulumi.get(self, "page_id")

    @page_id.setter
    def page_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "page_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Share project id.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Page means embedding the page, and panel means embedding the entire board.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="ticketNum")
    def ticket_num(self) -> Optional[pulumi.Input[int]]:
        """
        Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
        """
        return pulumi.get(self, "ticket_num")

    @ticket_num.setter
    def ticket_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ticket_num", value)

    @property
    @pulumi.getter(name="udpateAt")
    def udpate_at(self) -> Optional[pulumi.Input[str]]:
        """
        Upadte time.
        """
        return pulumi.get(self, "udpate_at")

    @udpate_at.setter
    def udpate_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "udpate_at", value)

    @property
    @pulumi.getter(name="userCorpId")
    def user_corp_id(self) -> Optional[pulumi.Input[str]]:
        """
        User enterprise ID (for multi-user only).
        """
        return pulumi.get(self, "user_corp_id")

    @user_corp_id.setter
    def user_corp_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_corp_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        UserId (for multi-user only).
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class EmbedTokenApply(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 page_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 ticket_num: Optional[pulumi.Input[int]] = None,
                 user_corp_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a EmbedTokenApply resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expire_time: Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
        :param pulumi.Input[int] page_id: Sharing page id, this is empty value 0 when embedding the board.
        :param pulumi.Input[int] project_id: Share project id.
        :param pulumi.Input[str] scope: Page means embedding the page, and panel means embedding the entire board.
        :param pulumi.Input[int] ticket_num: Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
        :param pulumi.Input[str] user_corp_id: User enterprise ID (for multi-user only).
        :param pulumi.Input[str] user_id: UserId (for multi-user only).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EmbedTokenApplyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a EmbedTokenApply resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param EmbedTokenApplyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EmbedTokenApplyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 page_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 ticket_num: Optional[pulumi.Input[int]] = None,
                 user_corp_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EmbedTokenApplyArgs.__new__(EmbedTokenApplyArgs)

            __props__.__dict__["expire_time"] = expire_time
            __props__.__dict__["page_id"] = page_id
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["scope"] = scope
            __props__.__dict__["ticket_num"] = ticket_num
            __props__.__dict__["user_corp_id"] = user_corp_id
            __props__.__dict__["user_id"] = user_id
            __props__.__dict__["bi_token"] = None
            __props__.__dict__["create_at"] = None
            __props__.__dict__["udpate_at"] = None
        super(EmbedTokenApply, __self__).__init__(
            'tencentcloud:Bi/embedTokenApply:EmbedTokenApply',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bi_token: Optional[pulumi.Input[str]] = None,
            create_at: Optional[pulumi.Input[str]] = None,
            expire_time: Optional[pulumi.Input[str]] = None,
            page_id: Optional[pulumi.Input[int]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            ticket_num: Optional[pulumi.Input[int]] = None,
            udpate_at: Optional[pulumi.Input[str]] = None,
            user_corp_id: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'EmbedTokenApply':
        """
        Get an existing EmbedTokenApply resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bi_token: Create the generated token.
        :param pulumi.Input[str] create_at: Create time.
        :param pulumi.Input[str] expire_time: Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
        :param pulumi.Input[int] page_id: Sharing page id, this is empty value 0 when embedding the board.
        :param pulumi.Input[int] project_id: Share project id.
        :param pulumi.Input[str] scope: Page means embedding the page, and panel means embedding the entire board.
        :param pulumi.Input[int] ticket_num: Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
        :param pulumi.Input[str] udpate_at: Upadte time.
        :param pulumi.Input[str] user_corp_id: User enterprise ID (for multi-user only).
        :param pulumi.Input[str] user_id: UserId (for multi-user only).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EmbedTokenApplyState.__new__(_EmbedTokenApplyState)

        __props__.__dict__["bi_token"] = bi_token
        __props__.__dict__["create_at"] = create_at
        __props__.__dict__["expire_time"] = expire_time
        __props__.__dict__["page_id"] = page_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["scope"] = scope
        __props__.__dict__["ticket_num"] = ticket_num
        __props__.__dict__["udpate_at"] = udpate_at
        __props__.__dict__["user_corp_id"] = user_corp_id
        __props__.__dict__["user_id"] = user_id
        return EmbedTokenApply(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="biToken")
    def bi_token(self) -> pulumi.Output[str]:
        """
        Create the generated token.
        """
        return pulumi.get(self, "bi_token")

    @property
    @pulumi.getter(name="createAt")
    def create_at(self) -> pulumi.Output[str]:
        """
        Create time.
        """
        return pulumi.get(self, "create_at")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[Optional[str]]:
        """
        Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="pageId")
    def page_id(self) -> pulumi.Output[Optional[int]]:
        """
        Sharing page id, this is empty value 0 when embedding the board.
        """
        return pulumi.get(self, "page_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[int]]:
        """
        Share project id.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[Optional[str]]:
        """
        Page means embedding the page, and panel means embedding the entire board.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter(name="ticketNum")
    def ticket_num(self) -> pulumi.Output[Optional[int]]:
        """
        Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
        """
        return pulumi.get(self, "ticket_num")

    @property
    @pulumi.getter(name="udpateAt")
    def udpate_at(self) -> pulumi.Output[str]:
        """
        Upadte time.
        """
        return pulumi.get(self, "udpate_at")

    @property
    @pulumi.getter(name="userCorpId")
    def user_corp_id(self) -> pulumi.Output[Optional[str]]:
        """
        User enterprise ID (for multi-user only).
        """
        return pulumi.get(self, "user_corp_id")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[Optional[str]]:
        """
        UserId (for multi-user only).
        """
        return pulumi.get(self, "user_id")

