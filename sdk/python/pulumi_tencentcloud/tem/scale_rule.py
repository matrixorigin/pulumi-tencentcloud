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

__all__ = ['ScaleRuleArgs', 'ScaleRule']

@pulumi.input_type
class ScaleRuleArgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[str],
                 autoscaler: pulumi.Input['ScaleRuleAutoscalerArgs'],
                 environment_id: pulumi.Input[str],
                 workload_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ScaleRule resource.
        :param pulumi.Input[str] application_id: application ID.
        :param pulumi.Input['ScaleRuleAutoscalerArgs'] autoscaler: .
        :param pulumi.Input[str] environment_id: environment ID.
        :param pulumi.Input[str] workload_id: application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        """
        pulumi.set(__self__, "application_id", application_id)
        pulumi.set(__self__, "autoscaler", autoscaler)
        pulumi.set(__self__, "environment_id", environment_id)
        pulumi.set(__self__, "workload_id", workload_id)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Input[str]:
        """
        application ID.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter
    def autoscaler(self) -> pulumi.Input['ScaleRuleAutoscalerArgs']:
        """
        .
        """
        return pulumi.get(self, "autoscaler")

    @autoscaler.setter
    def autoscaler(self, value: pulumi.Input['ScaleRuleAutoscalerArgs']):
        pulumi.set(self, "autoscaler", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Input[str]:
        """
        environment ID.
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter(name="workloadId")
    def workload_id(self) -> pulumi.Input[str]:
        """
        application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        """
        return pulumi.get(self, "workload_id")

    @workload_id.setter
    def workload_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workload_id", value)


@pulumi.input_type
class _ScaleRuleState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 autoscaler: Optional[pulumi.Input['ScaleRuleAutoscalerArgs']] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 workload_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ScaleRule resources.
        :param pulumi.Input[str] application_id: application ID.
        :param pulumi.Input['ScaleRuleAutoscalerArgs'] autoscaler: .
        :param pulumi.Input[str] environment_id: environment ID.
        :param pulumi.Input[str] workload_id: application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if autoscaler is not None:
            pulumi.set(__self__, "autoscaler", autoscaler)
        if environment_id is not None:
            pulumi.set(__self__, "environment_id", environment_id)
        if workload_id is not None:
            pulumi.set(__self__, "workload_id", workload_id)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        application ID.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter
    def autoscaler(self) -> Optional[pulumi.Input['ScaleRuleAutoscalerArgs']]:
        """
        .
        """
        return pulumi.get(self, "autoscaler")

    @autoscaler.setter
    def autoscaler(self, value: Optional[pulumi.Input['ScaleRuleAutoscalerArgs']]):
        pulumi.set(self, "autoscaler", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> Optional[pulumi.Input[str]]:
        """
        environment ID.
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter(name="workloadId")
    def workload_id(self) -> Optional[pulumi.Input[str]]:
        """
        application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        """
        return pulumi.get(self, "workload_id")

    @workload_id.setter
    def workload_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workload_id", value)


class ScaleRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 autoscaler: Optional[pulumi.Input[pulumi.InputType['ScaleRuleAutoscalerArgs']]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 workload_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ScaleRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: application ID.
        :param pulumi.Input[pulumi.InputType['ScaleRuleAutoscalerArgs']] autoscaler: .
        :param pulumi.Input[str] environment_id: environment ID.
        :param pulumi.Input[str] workload_id: application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScaleRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ScaleRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ScaleRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScaleRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 autoscaler: Optional[pulumi.Input[pulumi.InputType['ScaleRuleAutoscalerArgs']]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 workload_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ScaleRuleArgs.__new__(ScaleRuleArgs)

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            if autoscaler is None and not opts.urn:
                raise TypeError("Missing required property 'autoscaler'")
            __props__.__dict__["autoscaler"] = autoscaler
            if environment_id is None and not opts.urn:
                raise TypeError("Missing required property 'environment_id'")
            __props__.__dict__["environment_id"] = environment_id
            if workload_id is None and not opts.urn:
                raise TypeError("Missing required property 'workload_id'")
            __props__.__dict__["workload_id"] = workload_id
        super(ScaleRule, __self__).__init__(
            'tencentcloud:Tem/scaleRule:ScaleRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            autoscaler: Optional[pulumi.Input[pulumi.InputType['ScaleRuleAutoscalerArgs']]] = None,
            environment_id: Optional[pulumi.Input[str]] = None,
            workload_id: Optional[pulumi.Input[str]] = None) -> 'ScaleRule':
        """
        Get an existing ScaleRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: application ID.
        :param pulumi.Input[pulumi.InputType['ScaleRuleAutoscalerArgs']] autoscaler: .
        :param pulumi.Input[str] environment_id: environment ID.
        :param pulumi.Input[str] workload_id: application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ScaleRuleState.__new__(_ScaleRuleState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["autoscaler"] = autoscaler
        __props__.__dict__["environment_id"] = environment_id
        __props__.__dict__["workload_id"] = workload_id
        return ScaleRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        application ID.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def autoscaler(self) -> pulumi.Output['outputs.ScaleRuleAutoscaler']:
        """
        .
        """
        return pulumi.get(self, "autoscaler")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[str]:
        """
        environment ID.
        """
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter(name="workloadId")
    def workload_id(self) -> pulumi.Output[str]:
        """
        application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        """
        return pulumi.get(self, "workload_id")
