# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PluginAttachmentArgs', 'PluginAttachment']

@pulumi.input_type
class PluginAttachmentArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[str],
                 environment_name: pulumi.Input[str],
                 plugin_id: pulumi.Input[str],
                 service_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a PluginAttachment resource.
        :param pulumi.Input[str] api_id: Id of API.
        :param pulumi.Input[str] environment_name: Name of Environment.
        :param pulumi.Input[str] plugin_id: Id of Plugin.
        :param pulumi.Input[str] service_id: Id of Service.
        """
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "environment_name", environment_name)
        pulumi.set(__self__, "plugin_id", plugin_id)
        pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        """
        Id of API.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> pulumi.Input[str]:
        """
        Name of Environment.
        """
        return pulumi.get(self, "environment_name")

    @environment_name.setter
    def environment_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment_name", value)

    @property
    @pulumi.getter(name="pluginId")
    def plugin_id(self) -> pulumi.Input[str]:
        """
        Id of Plugin.
        """
        return pulumi.get(self, "plugin_id")

    @plugin_id.setter
    def plugin_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "plugin_id", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Input[str]:
        """
        Id of Service.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_id", value)


@pulumi.input_type
class _PluginAttachmentState:
    def __init__(__self__, *,
                 api_id: Optional[pulumi.Input[str]] = None,
                 environment_name: Optional[pulumi.Input[str]] = None,
                 plugin_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PluginAttachment resources.
        :param pulumi.Input[str] api_id: Id of API.
        :param pulumi.Input[str] environment_name: Name of Environment.
        :param pulumi.Input[str] plugin_id: Id of Plugin.
        :param pulumi.Input[str] service_id: Id of Service.
        """
        if api_id is not None:
            pulumi.set(__self__, "api_id", api_id)
        if environment_name is not None:
            pulumi.set(__self__, "environment_name", environment_name)
        if plugin_id is not None:
            pulumi.set(__self__, "plugin_id", plugin_id)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of API.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Environment.
        """
        return pulumi.get(self, "environment_name")

    @environment_name.setter
    def environment_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment_name", value)

    @property
    @pulumi.getter(name="pluginId")
    def plugin_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of Plugin.
        """
        return pulumi.get(self, "plugin_id")

    @plugin_id.setter
    def plugin_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plugin_id", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of Service.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_id", value)


class PluginAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 environment_name: Optional[pulumi.Input[str]] = None,
                 plugin_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a PluginAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: Id of API.
        :param pulumi.Input[str] environment_name: Name of Environment.
        :param pulumi.Input[str] plugin_id: Id of Plugin.
        :param pulumi.Input[str] service_id: Id of Service.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PluginAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PluginAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PluginAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PluginAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 environment_name: Optional[pulumi.Input[str]] = None,
                 plugin_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PluginAttachmentArgs.__new__(PluginAttachmentArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            if environment_name is None and not opts.urn:
                raise TypeError("Missing required property 'environment_name'")
            __props__.__dict__["environment_name"] = environment_name
            if plugin_id is None and not opts.urn:
                raise TypeError("Missing required property 'plugin_id'")
            __props__.__dict__["plugin_id"] = plugin_id
            if service_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_id'")
            __props__.__dict__["service_id"] = service_id
        super(PluginAttachment, __self__).__init__(
            'tencentcloud:ApiGateway/pluginAttachment:PluginAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_id: Optional[pulumi.Input[str]] = None,
            environment_name: Optional[pulumi.Input[str]] = None,
            plugin_id: Optional[pulumi.Input[str]] = None,
            service_id: Optional[pulumi.Input[str]] = None) -> 'PluginAttachment':
        """
        Get an existing PluginAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: Id of API.
        :param pulumi.Input[str] environment_name: Name of Environment.
        :param pulumi.Input[str] plugin_id: Id of Plugin.
        :param pulumi.Input[str] service_id: Id of Service.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PluginAttachmentState.__new__(_PluginAttachmentState)

        __props__.__dict__["api_id"] = api_id
        __props__.__dict__["environment_name"] = environment_name
        __props__.__dict__["plugin_id"] = plugin_id
        __props__.__dict__["service_id"] = service_id
        return PluginAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        """
        Id of API.
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> pulumi.Output[str]:
        """
        Name of Environment.
        """
        return pulumi.get(self, "environment_name")

    @property
    @pulumi.getter(name="pluginId")
    def plugin_id(self) -> pulumi.Output[str]:
        """
        Id of Plugin.
        """
        return pulumi.get(self, "plugin_id")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        """
        Id of Service.
        """
        return pulumi.get(self, "service_id")
