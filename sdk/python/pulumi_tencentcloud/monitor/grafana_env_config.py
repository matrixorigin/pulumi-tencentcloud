# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GrafanaEnvConfigArgs', 'GrafanaEnvConfig']

@pulumi.input_type
class GrafanaEnvConfigArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 envs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a GrafanaEnvConfig resource.
        :param pulumi.Input[str] instance_id: Grafana instance ID.
        :param pulumi.Input[Mapping[str, Any]] envs: Environment variables.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if envs is not None:
            pulumi.set(__self__, "envs", envs)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Grafana instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def envs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Environment variables.
        """
        return pulumi.get(self, "envs")

    @envs.setter
    def envs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "envs", value)


@pulumi.input_type
class _GrafanaEnvConfigState:
    def __init__(__self__, *,
                 envs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GrafanaEnvConfig resources.
        :param pulumi.Input[Mapping[str, Any]] envs: Environment variables.
        :param pulumi.Input[str] instance_id: Grafana instance ID.
        """
        if envs is not None:
            pulumi.set(__self__, "envs", envs)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter
    def envs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Environment variables.
        """
        return pulumi.get(self, "envs")

    @envs.setter
    def envs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "envs", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Grafana instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class GrafanaEnvConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 envs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a GrafanaEnvConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] envs: Environment variables.
        :param pulumi.Input[str] instance_id: Grafana instance ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GrafanaEnvConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a GrafanaEnvConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param GrafanaEnvConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GrafanaEnvConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 envs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GrafanaEnvConfigArgs.__new__(GrafanaEnvConfigArgs)

            __props__.__dict__["envs"] = envs
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(GrafanaEnvConfig, __self__).__init__(
            'tencentcloud:Monitor/grafanaEnvConfig:GrafanaEnvConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            envs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'GrafanaEnvConfig':
        """
        Get an existing GrafanaEnvConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] envs: Environment variables.
        :param pulumi.Input[str] instance_id: Grafana instance ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GrafanaEnvConfigState.__new__(_GrafanaEnvConfigState)

        __props__.__dict__["envs"] = envs
        __props__.__dict__["instance_id"] = instance_id
        return GrafanaEnvConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def envs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Environment variables.
        """
        return pulumi.get(self, "envs")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Grafana instance ID.
        """
        return pulumi.get(self, "instance_id")
