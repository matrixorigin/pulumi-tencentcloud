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

__all__ = ['PutEventsArgs', 'PutEvents']

@pulumi.input_type
class PutEventsArgs:
    def __init__(__self__, *,
                 event_bus_id: pulumi.Input[str],
                 event_lists: pulumi.Input[Sequence[pulumi.Input['PutEventsEventListArgs']]]):
        """
        The set of arguments for constructing a PutEvents resource.
        :param pulumi.Input[str] event_bus_id: event bus Id.
        :param pulumi.Input[Sequence[pulumi.Input['PutEventsEventListArgs']]] event_lists: event list.
        """
        pulumi.set(__self__, "event_bus_id", event_bus_id)
        pulumi.set(__self__, "event_lists", event_lists)

    @property
    @pulumi.getter(name="eventBusId")
    def event_bus_id(self) -> pulumi.Input[str]:
        """
        event bus Id.
        """
        return pulumi.get(self, "event_bus_id")

    @event_bus_id.setter
    def event_bus_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "event_bus_id", value)

    @property
    @pulumi.getter(name="eventLists")
    def event_lists(self) -> pulumi.Input[Sequence[pulumi.Input['PutEventsEventListArgs']]]:
        """
        event list.
        """
        return pulumi.get(self, "event_lists")

    @event_lists.setter
    def event_lists(self, value: pulumi.Input[Sequence[pulumi.Input['PutEventsEventListArgs']]]):
        pulumi.set(self, "event_lists", value)


@pulumi.input_type
class _PutEventsState:
    def __init__(__self__, *,
                 event_bus_id: Optional[pulumi.Input[str]] = None,
                 event_lists: Optional[pulumi.Input[Sequence[pulumi.Input['PutEventsEventListArgs']]]] = None):
        """
        Input properties used for looking up and filtering PutEvents resources.
        :param pulumi.Input[str] event_bus_id: event bus Id.
        :param pulumi.Input[Sequence[pulumi.Input['PutEventsEventListArgs']]] event_lists: event list.
        """
        if event_bus_id is not None:
            pulumi.set(__self__, "event_bus_id", event_bus_id)
        if event_lists is not None:
            pulumi.set(__self__, "event_lists", event_lists)

    @property
    @pulumi.getter(name="eventBusId")
    def event_bus_id(self) -> Optional[pulumi.Input[str]]:
        """
        event bus Id.
        """
        return pulumi.get(self, "event_bus_id")

    @event_bus_id.setter
    def event_bus_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_bus_id", value)

    @property
    @pulumi.getter(name="eventLists")
    def event_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PutEventsEventListArgs']]]]:
        """
        event list.
        """
        return pulumi.get(self, "event_lists")

    @event_lists.setter
    def event_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PutEventsEventListArgs']]]]):
        pulumi.set(self, "event_lists", value)


class PutEvents(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 event_bus_id: Optional[pulumi.Input[str]] = None,
                 event_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PutEventsEventListArgs']]]]] = None,
                 __props__=None):
        """
        Create a PutEvents resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] event_bus_id: event bus Id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PutEventsEventListArgs']]]] event_lists: event list.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PutEventsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PutEvents resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PutEventsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PutEventsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 event_bus_id: Optional[pulumi.Input[str]] = None,
                 event_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PutEventsEventListArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PutEventsArgs.__new__(PutEventsArgs)

            if event_bus_id is None and not opts.urn:
                raise TypeError("Missing required property 'event_bus_id'")
            __props__.__dict__["event_bus_id"] = event_bus_id
            if event_lists is None and not opts.urn:
                raise TypeError("Missing required property 'event_lists'")
            __props__.__dict__["event_lists"] = event_lists
        super(PutEvents, __self__).__init__(
            'tencentcloud:Eb/putEvents:PutEvents',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            event_bus_id: Optional[pulumi.Input[str]] = None,
            event_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PutEventsEventListArgs']]]]] = None) -> 'PutEvents':
        """
        Get an existing PutEvents resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] event_bus_id: event bus Id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PutEventsEventListArgs']]]] event_lists: event list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PutEventsState.__new__(_PutEventsState)

        __props__.__dict__["event_bus_id"] = event_bus_id
        __props__.__dict__["event_lists"] = event_lists
        return PutEvents(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="eventBusId")
    def event_bus_id(self) -> pulumi.Output[str]:
        """
        event bus Id.
        """
        return pulumi.get(self, "event_bus_id")

    @property
    @pulumi.getter(name="eventLists")
    def event_lists(self) -> pulumi.Output[Sequence['outputs.PutEventsEventList']]:
        """
        event list.
        """
        return pulumi.get(self, "event_lists")
