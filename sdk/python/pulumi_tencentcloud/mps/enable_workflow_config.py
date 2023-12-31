# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EnableWorkflowConfigArgs', 'EnableWorkflowConfig']

@pulumi.input_type
class EnableWorkflowConfigArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 workflow_id: pulumi.Input[int]):
        """
        The set of arguments for constructing a EnableWorkflowConfig resource.
        :param pulumi.Input[bool] enabled: true: enable; false: disable.
        :param pulumi.Input[int] workflow_id: Workflow ID.
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "workflow_id", workflow_id)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        true: enable; false: disable.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="workflowId")
    def workflow_id(self) -> pulumi.Input[int]:
        """
        Workflow ID.
        """
        return pulumi.get(self, "workflow_id")

    @workflow_id.setter
    def workflow_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "workflow_id", value)


@pulumi.input_type
class _EnableWorkflowConfigState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 workflow_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering EnableWorkflowConfig resources.
        :param pulumi.Input[bool] enabled: true: enable; false: disable.
        :param pulumi.Input[int] workflow_id: Workflow ID.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if workflow_id is not None:
            pulumi.set(__self__, "workflow_id", workflow_id)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        true: enable; false: disable.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="workflowId")
    def workflow_id(self) -> Optional[pulumi.Input[int]]:
        """
        Workflow ID.
        """
        return pulumi.get(self, "workflow_id")

    @workflow_id.setter
    def workflow_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "workflow_id", value)


class EnableWorkflowConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 workflow_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a EnableWorkflowConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: true: enable; false: disable.
        :param pulumi.Input[int] workflow_id: Workflow ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnableWorkflowConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a EnableWorkflowConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param EnableWorkflowConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnableWorkflowConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 workflow_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnableWorkflowConfigArgs.__new__(EnableWorkflowConfigArgs)

            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            if workflow_id is None and not opts.urn:
                raise TypeError("Missing required property 'workflow_id'")
            __props__.__dict__["workflow_id"] = workflow_id
        super(EnableWorkflowConfig, __self__).__init__(
            'tencentcloud:Mps/enableWorkflowConfig:EnableWorkflowConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            workflow_id: Optional[pulumi.Input[int]] = None) -> 'EnableWorkflowConfig':
        """
        Get an existing EnableWorkflowConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: true: enable; false: disable.
        :param pulumi.Input[int] workflow_id: Workflow ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnableWorkflowConfigState.__new__(_EnableWorkflowConfigState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["workflow_id"] = workflow_id
        return EnableWorkflowConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        true: enable; false: disable.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="workflowId")
    def workflow_id(self) -> pulumi.Output[int]:
        """
        Workflow ID.
        """
        return pulumi.get(self, "workflow_id")

