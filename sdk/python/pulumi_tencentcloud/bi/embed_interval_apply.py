# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EmbedIntervalApplyArgs', 'EmbedIntervalApply']

@pulumi.input_type
class EmbedIntervalApplyArgs:
    def __init__(__self__, *,
                 bi_token: Optional[pulumi.Input[str]] = None,
                 page_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EmbedIntervalApply resource.
        :param pulumi.Input[str] bi_token: Token that needs to be applied for extension.
        :param pulumi.Input[int] page_id: Sharing page id, this is empty value 0 when embedding the board.
        :param pulumi.Input[int] project_id: Sharing project id, required.
        :param pulumi.Input[str] scope: Choose panel or page.
        """
        if bi_token is not None:
            pulumi.set(__self__, "bi_token", bi_token)
        if page_id is not None:
            pulumi.set(__self__, "page_id", page_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter(name="biToken")
    def bi_token(self) -> Optional[pulumi.Input[str]]:
        """
        Token that needs to be applied for extension.
        """
        return pulumi.get(self, "bi_token")

    @bi_token.setter
    def bi_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bi_token", value)

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
        Sharing project id, required.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Choose panel or page.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)


@pulumi.input_type
class _EmbedIntervalApplyState:
    def __init__(__self__, *,
                 bi_token: Optional[pulumi.Input[str]] = None,
                 page_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EmbedIntervalApply resources.
        :param pulumi.Input[str] bi_token: Token that needs to be applied for extension.
        :param pulumi.Input[int] page_id: Sharing page id, this is empty value 0 when embedding the board.
        :param pulumi.Input[int] project_id: Sharing project id, required.
        :param pulumi.Input[str] scope: Choose panel or page.
        """
        if bi_token is not None:
            pulumi.set(__self__, "bi_token", bi_token)
        if page_id is not None:
            pulumi.set(__self__, "page_id", page_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter(name="biToken")
    def bi_token(self) -> Optional[pulumi.Input[str]]:
        """
        Token that needs to be applied for extension.
        """
        return pulumi.get(self, "bi_token")

    @bi_token.setter
    def bi_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bi_token", value)

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
        Sharing project id, required.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Choose panel or page.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)


class EmbedIntervalApply(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bi_token: Optional[pulumi.Input[str]] = None,
                 page_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a EmbedIntervalApply resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bi_token: Token that needs to be applied for extension.
        :param pulumi.Input[int] page_id: Sharing page id, this is empty value 0 when embedding the board.
        :param pulumi.Input[int] project_id: Sharing project id, required.
        :param pulumi.Input[str] scope: Choose panel or page.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EmbedIntervalApplyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a EmbedIntervalApply resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param EmbedIntervalApplyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EmbedIntervalApplyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bi_token: Optional[pulumi.Input[str]] = None,
                 page_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EmbedIntervalApplyArgs.__new__(EmbedIntervalApplyArgs)

            __props__.__dict__["bi_token"] = bi_token
            __props__.__dict__["page_id"] = page_id
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["scope"] = scope
        super(EmbedIntervalApply, __self__).__init__(
            'tencentcloud:Bi/embedIntervalApply:EmbedIntervalApply',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bi_token: Optional[pulumi.Input[str]] = None,
            page_id: Optional[pulumi.Input[int]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            scope: Optional[pulumi.Input[str]] = None) -> 'EmbedIntervalApply':
        """
        Get an existing EmbedIntervalApply resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bi_token: Token that needs to be applied for extension.
        :param pulumi.Input[int] page_id: Sharing page id, this is empty value 0 when embedding the board.
        :param pulumi.Input[int] project_id: Sharing project id, required.
        :param pulumi.Input[str] scope: Choose panel or page.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EmbedIntervalApplyState.__new__(_EmbedIntervalApplyState)

        __props__.__dict__["bi_token"] = bi_token
        __props__.__dict__["page_id"] = page_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["scope"] = scope
        return EmbedIntervalApply(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="biToken")
    def bi_token(self) -> pulumi.Output[Optional[str]]:
        """
        Token that needs to be applied for extension.
        """
        return pulumi.get(self, "bi_token")

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
        Sharing project id, required.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[Optional[str]]:
        """
        Choose panel or page.
        """
        return pulumi.get(self, "scope")
