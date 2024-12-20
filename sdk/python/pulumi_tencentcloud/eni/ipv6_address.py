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

__all__ = ['Ipv6AddressArgs', 'Ipv6Address']

@pulumi.input_type
class Ipv6AddressArgs:
    def __init__(__self__, *,
                 network_interface_id: pulumi.Input[str],
                 ipv6_address_count: Optional[pulumi.Input[int]] = None,
                 ipv6_addresses: Optional[pulumi.Input[Sequence[pulumi.Input['Ipv6AddressIpv6AddressArgs']]]] = None):
        """
        The set of arguments for constructing a Ipv6Address resource.
        :param pulumi.Input[str] network_interface_id: ENI instance `ID`, in the form of `eni-m6dyj72l`.
        :param pulumi.Input[int] ipv6_address_count: The number of automatically assigned IPv6 addresses and the total number of private IP addresses cannot exceed the
               quota. This should be combined with the input parameter `ipv6_addresses` for quota calculation. At least one of them,
               either this or 'Ipv6Addresses', must be provided.
        :param pulumi.Input[Sequence[pulumi.Input['Ipv6AddressIpv6AddressArgs']]] ipv6_addresses: The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter
               `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
        """
        pulumi.set(__self__, "network_interface_id", network_interface_id)
        if ipv6_address_count is not None:
            pulumi.set(__self__, "ipv6_address_count", ipv6_address_count)
        if ipv6_addresses is not None:
            pulumi.set(__self__, "ipv6_addresses", ipv6_addresses)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Input[str]:
        """
        ENI instance `ID`, in the form of `eni-m6dyj72l`.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_interface_id", value)

    @property
    @pulumi.getter(name="ipv6AddressCount")
    def ipv6_address_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of automatically assigned IPv6 addresses and the total number of private IP addresses cannot exceed the
        quota. This should be combined with the input parameter `ipv6_addresses` for quota calculation. At least one of them,
        either this or 'Ipv6Addresses', must be provided.
        """
        return pulumi.get(self, "ipv6_address_count")

    @ipv6_address_count.setter
    def ipv6_address_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ipv6_address_count", value)

    @property
    @pulumi.getter(name="ipv6Addresses")
    def ipv6_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['Ipv6AddressIpv6AddressArgs']]]]:
        """
        The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter
        `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
        """
        return pulumi.get(self, "ipv6_addresses")

    @ipv6_addresses.setter
    def ipv6_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['Ipv6AddressIpv6AddressArgs']]]]):
        pulumi.set(self, "ipv6_addresses", value)


@pulumi.input_type
class _Ipv6AddressState:
    def __init__(__self__, *,
                 ipv6_address_count: Optional[pulumi.Input[int]] = None,
                 ipv6_addresses: Optional[pulumi.Input[Sequence[pulumi.Input['Ipv6AddressIpv6AddressArgs']]]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ipv6Address resources.
        :param pulumi.Input[int] ipv6_address_count: The number of automatically assigned IPv6 addresses and the total number of private IP addresses cannot exceed the
               quota. This should be combined with the input parameter `ipv6_addresses` for quota calculation. At least one of them,
               either this or 'Ipv6Addresses', must be provided.
        :param pulumi.Input[Sequence[pulumi.Input['Ipv6AddressIpv6AddressArgs']]] ipv6_addresses: The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter
               `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
        :param pulumi.Input[str] network_interface_id: ENI instance `ID`, in the form of `eni-m6dyj72l`.
        """
        if ipv6_address_count is not None:
            pulumi.set(__self__, "ipv6_address_count", ipv6_address_count)
        if ipv6_addresses is not None:
            pulumi.set(__self__, "ipv6_addresses", ipv6_addresses)
        if network_interface_id is not None:
            pulumi.set(__self__, "network_interface_id", network_interface_id)

    @property
    @pulumi.getter(name="ipv6AddressCount")
    def ipv6_address_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of automatically assigned IPv6 addresses and the total number of private IP addresses cannot exceed the
        quota. This should be combined with the input parameter `ipv6_addresses` for quota calculation. At least one of them,
        either this or 'Ipv6Addresses', must be provided.
        """
        return pulumi.get(self, "ipv6_address_count")

    @ipv6_address_count.setter
    def ipv6_address_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ipv6_address_count", value)

    @property
    @pulumi.getter(name="ipv6Addresses")
    def ipv6_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['Ipv6AddressIpv6AddressArgs']]]]:
        """
        The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter
        `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
        """
        return pulumi.get(self, "ipv6_addresses")

    @ipv6_addresses.setter
    def ipv6_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['Ipv6AddressIpv6AddressArgs']]]]):
        pulumi.set(self, "ipv6_addresses", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        ENI instance `ID`, in the form of `eni-m6dyj72l`.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_interface_id", value)


class Ipv6Address(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ipv6_address_count: Optional[pulumi.Input[int]] = None,
                 ipv6_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['Ipv6AddressIpv6AddressArgs']]]]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Ipv6Address resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ipv6_address_count: The number of automatically assigned IPv6 addresses and the total number of private IP addresses cannot exceed the
               quota. This should be combined with the input parameter `ipv6_addresses` for quota calculation. At least one of them,
               either this or 'Ipv6Addresses', must be provided.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['Ipv6AddressIpv6AddressArgs']]]] ipv6_addresses: The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter
               `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
        :param pulumi.Input[str] network_interface_id: ENI instance `ID`, in the form of `eni-m6dyj72l`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Ipv6AddressArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Ipv6Address resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param Ipv6AddressArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Ipv6AddressArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ipv6_address_count: Optional[pulumi.Input[int]] = None,
                 ipv6_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['Ipv6AddressIpv6AddressArgs']]]]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Ipv6AddressArgs.__new__(Ipv6AddressArgs)

            __props__.__dict__["ipv6_address_count"] = ipv6_address_count
            __props__.__dict__["ipv6_addresses"] = ipv6_addresses
            if network_interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_interface_id'")
            __props__.__dict__["network_interface_id"] = network_interface_id
        super(Ipv6Address, __self__).__init__(
            'tencentcloud:Eni/ipv6Address:Ipv6Address',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ipv6_address_count: Optional[pulumi.Input[int]] = None,
            ipv6_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['Ipv6AddressIpv6AddressArgs']]]]] = None,
            network_interface_id: Optional[pulumi.Input[str]] = None) -> 'Ipv6Address':
        """
        Get an existing Ipv6Address resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ipv6_address_count: The number of automatically assigned IPv6 addresses and the total number of private IP addresses cannot exceed the
               quota. This should be combined with the input parameter `ipv6_addresses` for quota calculation. At least one of them,
               either this or 'Ipv6Addresses', must be provided.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['Ipv6AddressIpv6AddressArgs']]]] ipv6_addresses: The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter
               `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
        :param pulumi.Input[str] network_interface_id: ENI instance `ID`, in the form of `eni-m6dyj72l`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Ipv6AddressState.__new__(_Ipv6AddressState)

        __props__.__dict__["ipv6_address_count"] = ipv6_address_count
        __props__.__dict__["ipv6_addresses"] = ipv6_addresses
        __props__.__dict__["network_interface_id"] = network_interface_id
        return Ipv6Address(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ipv6AddressCount")
    def ipv6_address_count(self) -> pulumi.Output[int]:
        """
        The number of automatically assigned IPv6 addresses and the total number of private IP addresses cannot exceed the
        quota. This should be combined with the input parameter `ipv6_addresses` for quota calculation. At least one of them,
        either this or 'Ipv6Addresses', must be provided.
        """
        return pulumi.get(self, "ipv6_address_count")

    @property
    @pulumi.getter(name="ipv6Addresses")
    def ipv6_addresses(self) -> pulumi.Output[Sequence['outputs.Ipv6AddressIpv6Address']]:
        """
        The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter
        `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
        """
        return pulumi.get(self, "ipv6_addresses")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[str]:
        """
        ENI instance `ID`, in the form of `eni-m6dyj72l`.
        """
        return pulumi.get(self, "network_interface_id")

