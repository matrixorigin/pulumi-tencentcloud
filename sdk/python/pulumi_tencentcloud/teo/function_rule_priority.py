# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FunctionRulePriorityArgs', 'FunctionRulePriority']

@pulumi.input_type
class FunctionRulePriorityArgs:
    def __init__(__self__, *,
                 function_id: pulumi.Input[str],
                 rule_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 zone_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a FunctionRulePriority resource.
        :param pulumi.Input[str] function_id: ID of the Function.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rule_ids: he list of rule IDs. It is required to include all rule IDs after adjusting their priorities. The execution order of
               multiple rules follows a top-down sequence. If not specified, the original priority order will be maintained.
        :param pulumi.Input[str] zone_id: ID of the site.
        """
        pulumi.set(__self__, "function_id", function_id)
        pulumi.set(__self__, "rule_ids", rule_ids)
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> pulumi.Input[str]:
        """
        ID of the Function.
        """
        return pulumi.get(self, "function_id")

    @function_id.setter
    def function_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_id", value)

    @property
    @pulumi.getter(name="ruleIds")
    def rule_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        he list of rule IDs. It is required to include all rule IDs after adjusting their priorities. The execution order of
        multiple rules follows a top-down sequence. If not specified, the original priority order will be maintained.
        """
        return pulumi.get(self, "rule_ids")

    @rule_ids.setter
    def rule_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "rule_ids", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        ID of the site.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)


@pulumi.input_type
class _FunctionRulePriorityState:
    def __init__(__self__, *,
                 function_id: Optional[pulumi.Input[str]] = None,
                 rule_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FunctionRulePriority resources.
        :param pulumi.Input[str] function_id: ID of the Function.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rule_ids: he list of rule IDs. It is required to include all rule IDs after adjusting their priorities. The execution order of
               multiple rules follows a top-down sequence. If not specified, the original priority order will be maintained.
        :param pulumi.Input[str] zone_id: ID of the site.
        """
        if function_id is not None:
            pulumi.set(__self__, "function_id", function_id)
        if rule_ids is not None:
            pulumi.set(__self__, "rule_ids", rule_ids)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Function.
        """
        return pulumi.get(self, "function_id")

    @function_id.setter
    def function_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_id", value)

    @property
    @pulumi.getter(name="ruleIds")
    def rule_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        he list of rule IDs. It is required to include all rule IDs after adjusting their priorities. The execution order of
        multiple rules follows a top-down sequence. If not specified, the original priority order will be maintained.
        """
        return pulumi.get(self, "rule_ids")

    @rule_ids.setter
    def rule_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "rule_ids", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the site.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class FunctionRulePriority(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 rule_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FunctionRulePriority resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] function_id: ID of the Function.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rule_ids: he list of rule IDs. It is required to include all rule IDs after adjusting their priorities. The execution order of
               multiple rules follows a top-down sequence. If not specified, the original priority order will be maintained.
        :param pulumi.Input[str] zone_id: ID of the site.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionRulePriorityArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FunctionRulePriority resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FunctionRulePriorityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionRulePriorityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 rule_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionRulePriorityArgs.__new__(FunctionRulePriorityArgs)

            if function_id is None and not opts.urn:
                raise TypeError("Missing required property 'function_id'")
            __props__.__dict__["function_id"] = function_id
            if rule_ids is None and not opts.urn:
                raise TypeError("Missing required property 'rule_ids'")
            __props__.__dict__["rule_ids"] = rule_ids
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
        super(FunctionRulePriority, __self__).__init__(
            'tencentcloud:Teo/functionRulePriority:FunctionRulePriority',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            function_id: Optional[pulumi.Input[str]] = None,
            rule_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'FunctionRulePriority':
        """
        Get an existing FunctionRulePriority resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] function_id: ID of the Function.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rule_ids: he list of rule IDs. It is required to include all rule IDs after adjusting their priorities. The execution order of
               multiple rules follows a top-down sequence. If not specified, the original priority order will be maintained.
        :param pulumi.Input[str] zone_id: ID of the site.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FunctionRulePriorityState.__new__(_FunctionRulePriorityState)

        __props__.__dict__["function_id"] = function_id
        __props__.__dict__["rule_ids"] = rule_ids
        __props__.__dict__["zone_id"] = zone_id
        return FunctionRulePriority(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> pulumi.Output[str]:
        """
        ID of the Function.
        """
        return pulumi.get(self, "function_id")

    @property
    @pulumi.getter(name="ruleIds")
    def rule_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        he list of rule IDs. It is required to include all rule IDs after adjusting their priorities. The execution order of
        multiple rules follows a top-down sequence. If not specified, the original priority order will be maintained.
        """
        return pulumi.get(self, "rule_ids")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        ID of the site.
        """
        return pulumi.get(self, "zone_id")
