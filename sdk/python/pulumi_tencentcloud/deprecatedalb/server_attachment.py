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

__all__ = ['ServerAttachmentArgs', 'ServerAttachment']

@pulumi.input_type
class ServerAttachmentArgs:
    def __init__(__self__, *,
                 backends: pulumi.Input[Sequence[pulumi.Input['ServerAttachmentBackendArgs']]],
                 listener_id: pulumi.Input[str],
                 loadbalancer_id: pulumi.Input[str],
                 location_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServerAttachment resource.
        :param pulumi.Input[Sequence[pulumi.Input['ServerAttachmentBackendArgs']]] backends: list of backend server.
        :param pulumi.Input[str] listener_id: listener ID.
        :param pulumi.Input[str] loadbalancer_id: loadbalancer ID.
        :param pulumi.Input[str] location_id: location ID, only support for layer 7 loadbalancer.
        """
        pulumi.set(__self__, "backends", backends)
        pulumi.set(__self__, "listener_id", listener_id)
        pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        if location_id is not None:
            pulumi.set(__self__, "location_id", location_id)

    @property
    @pulumi.getter
    def backends(self) -> pulumi.Input[Sequence[pulumi.Input['ServerAttachmentBackendArgs']]]:
        """
        list of backend server.
        """
        return pulumi.get(self, "backends")

    @backends.setter
    def backends(self, value: pulumi.Input[Sequence[pulumi.Input['ServerAttachmentBackendArgs']]]):
        pulumi.set(self, "backends", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Input[str]:
        """
        listener ID.
        """
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> pulumi.Input[str]:
        """
        loadbalancer ID.
        """
        return pulumi.get(self, "loadbalancer_id")

    @loadbalancer_id.setter
    def loadbalancer_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "loadbalancer_id", value)

    @property
    @pulumi.getter(name="locationId")
    def location_id(self) -> Optional[pulumi.Input[str]]:
        """
        location ID, only support for layer 7 loadbalancer.
        """
        return pulumi.get(self, "location_id")

    @location_id.setter
    def location_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location_id", value)


@pulumi.input_type
class _ServerAttachmentState:
    def __init__(__self__, *,
                 backends: Optional[pulumi.Input[Sequence[pulumi.Input['ServerAttachmentBackendArgs']]]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 location_id: Optional[pulumi.Input[str]] = None,
                 protocol_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServerAttachment resources.
        :param pulumi.Input[Sequence[pulumi.Input['ServerAttachmentBackendArgs']]] backends: list of backend server.
        :param pulumi.Input[str] listener_id: listener ID.
        :param pulumi.Input[str] loadbalancer_id: loadbalancer ID.
        :param pulumi.Input[str] location_id: location ID, only support for layer 7 loadbalancer.
        :param pulumi.Input[str] protocol_type: The protocol type, http or tcp.
        """
        if backends is not None:
            pulumi.set(__self__, "backends", backends)
        if listener_id is not None:
            pulumi.set(__self__, "listener_id", listener_id)
        if loadbalancer_id is not None:
            pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        if location_id is not None:
            pulumi.set(__self__, "location_id", location_id)
        if protocol_type is not None:
            pulumi.set(__self__, "protocol_type", protocol_type)

    @property
    @pulumi.getter
    def backends(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerAttachmentBackendArgs']]]]:
        """
        list of backend server.
        """
        return pulumi.get(self, "backends")

    @backends.setter
    def backends(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerAttachmentBackendArgs']]]]):
        pulumi.set(self, "backends", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[pulumi.Input[str]]:
        """
        listener ID.
        """
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> Optional[pulumi.Input[str]]:
        """
        loadbalancer ID.
        """
        return pulumi.get(self, "loadbalancer_id")

    @loadbalancer_id.setter
    def loadbalancer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "loadbalancer_id", value)

    @property
    @pulumi.getter(name="locationId")
    def location_id(self) -> Optional[pulumi.Input[str]]:
        """
        location ID, only support for layer 7 loadbalancer.
        """
        return pulumi.get(self, "location_id")

    @location_id.setter
    def location_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location_id", value)

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol type, http or tcp.
        """
        return pulumi.get(self, "protocol_type")

    @protocol_type.setter
    def protocol_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol_type", value)


class ServerAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backends: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerAttachmentBackendArgs']]]]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 location_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ServerAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerAttachmentBackendArgs']]]] backends: list of backend server.
        :param pulumi.Input[str] listener_id: listener ID.
        :param pulumi.Input[str] loadbalancer_id: loadbalancer ID.
        :param pulumi.Input[str] location_id: location ID, only support for layer 7 loadbalancer.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ServerAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ServerAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backends: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerAttachmentBackendArgs']]]]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 location_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServerAttachmentArgs.__new__(ServerAttachmentArgs)

            if backends is None and not opts.urn:
                raise TypeError("Missing required property 'backends'")
            __props__.__dict__["backends"] = backends
            if listener_id is None and not opts.urn:
                raise TypeError("Missing required property 'listener_id'")
            __props__.__dict__["listener_id"] = listener_id
            if loadbalancer_id is None and not opts.urn:
                raise TypeError("Missing required property 'loadbalancer_id'")
            __props__.__dict__["loadbalancer_id"] = loadbalancer_id
            __props__.__dict__["location_id"] = location_id
            __props__.__dict__["protocol_type"] = None
        super(ServerAttachment, __self__).__init__(
            'tencentcloud:Deprecatedalb/serverAttachment:ServerAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backends: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerAttachmentBackendArgs']]]]] = None,
            listener_id: Optional[pulumi.Input[str]] = None,
            loadbalancer_id: Optional[pulumi.Input[str]] = None,
            location_id: Optional[pulumi.Input[str]] = None,
            protocol_type: Optional[pulumi.Input[str]] = None) -> 'ServerAttachment':
        """
        Get an existing ServerAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerAttachmentBackendArgs']]]] backends: list of backend server.
        :param pulumi.Input[str] listener_id: listener ID.
        :param pulumi.Input[str] loadbalancer_id: loadbalancer ID.
        :param pulumi.Input[str] location_id: location ID, only support for layer 7 loadbalancer.
        :param pulumi.Input[str] protocol_type: The protocol type, http or tcp.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServerAttachmentState.__new__(_ServerAttachmentState)

        __props__.__dict__["backends"] = backends
        __props__.__dict__["listener_id"] = listener_id
        __props__.__dict__["loadbalancer_id"] = loadbalancer_id
        __props__.__dict__["location_id"] = location_id
        __props__.__dict__["protocol_type"] = protocol_type
        return ServerAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backends(self) -> pulumi.Output[Sequence['outputs.ServerAttachmentBackend']]:
        """
        list of backend server.
        """
        return pulumi.get(self, "backends")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Output[str]:
        """
        listener ID.
        """
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> pulumi.Output[str]:
        """
        loadbalancer ID.
        """
        return pulumi.get(self, "loadbalancer_id")

    @property
    @pulumi.getter(name="locationId")
    def location_id(self) -> pulumi.Output[str]:
        """
        location ID, only support for layer 7 loadbalancer.
        """
        return pulumi.get(self, "location_id")

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> pulumi.Output[str]:
        """
        The protocol type, http or tcp.
        """
        return pulumi.get(self, "protocol_type")

