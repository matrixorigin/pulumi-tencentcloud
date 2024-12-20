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

__all__ = ['RocketmqInstanceArgs', 'RocketmqInstance']

@pulumi.input_type
class RocketmqInstanceArgs:
    def __init__(__self__, *,
                 instance_type: pulumi.Input[str],
                 sku_code: pulumi.Input[str],
                 subnet_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 enable_public: Optional[pulumi.Input[bool]] = None,
                 ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqInstanceIpRuleArgs']]]] = None,
                 message_retention: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a RocketmqInstance resource.
        :param pulumi.Input[str] instance_type: Instance type. Valid values: `EXPERIMENT`, `BASIC`, `PRO`, `PLATINUM`.
        :param pulumi.Input[str] sku_code: SKU code. Available specifications are as follows: experiment_500, basic_1k, basic_2k, basic_3k, basic_4k, basic_5k,
               basic_6k, basic_7k, basic_8k, basic_9k, basic_10k, pro_4k, pro_6k, pro_8k, pro_1w, pro_15k, pro_2w, pro_25k, pro_3w,
               pro_35k, pro_4w, pro_45k, pro_5w, pro_55k, pro_60k, pro_65k, pro_70k, pro_75k, pro_80k, pro_85k, pro_90k, pro_95k,
               pro_100k, platinum_1w, platinum_2w, platinum_3w, platinum_4w, platinum_5w, platinum_6w, platinum_7w, platinum_8w,
               platinum_9w, platinum_10w, platinum_12w, platinum_14w, platinum_16w, platinum_18w, platinum_20w, platinum_25w,
               platinum_30w, platinum_35w, platinum_40w, platinum_45w, platinum_50w, platinum_60w, platinum_70w, platinum_80w,
               platinum_90w, platinum_100w.
        :param pulumi.Input[str] subnet_id: Subnet id.
        :param pulumi.Input[str] vpc_id: VPC id.
        :param pulumi.Input[int] bandwidth: Public network bandwidth. `bandwidth` must be greater than zero when `enable_public` equal true.
        :param pulumi.Input[bool] enable_public: Whether to enable the public network. Must set `bandwidth` when `enable_public` equal true.
        :param pulumi.Input[Sequence[pulumi.Input['RocketmqInstanceIpRuleArgs']]] ip_rules: Public network access whitelist.
        :param pulumi.Input[int] message_retention: Message retention time in hours.
        :param pulumi.Input[str] name: Instance name.
        :param pulumi.Input[str] remark: Remark.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "sku_code", sku_code)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if enable_public is not None:
            pulumi.set(__self__, "enable_public", enable_public)
        if ip_rules is not None:
            pulumi.set(__self__, "ip_rules", ip_rules)
        if message_retention is not None:
            pulumi.set(__self__, "message_retention", message_retention)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        """
        Instance type. Valid values: `EXPERIMENT`, `BASIC`, `PRO`, `PLATINUM`.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="skuCode")
    def sku_code(self) -> pulumi.Input[str]:
        """
        SKU code. Available specifications are as follows: experiment_500, basic_1k, basic_2k, basic_3k, basic_4k, basic_5k,
        basic_6k, basic_7k, basic_8k, basic_9k, basic_10k, pro_4k, pro_6k, pro_8k, pro_1w, pro_15k, pro_2w, pro_25k, pro_3w,
        pro_35k, pro_4w, pro_45k, pro_5w, pro_55k, pro_60k, pro_65k, pro_70k, pro_75k, pro_80k, pro_85k, pro_90k, pro_95k,
        pro_100k, platinum_1w, platinum_2w, platinum_3w, platinum_4w, platinum_5w, platinum_6w, platinum_7w, platinum_8w,
        platinum_9w, platinum_10w, platinum_12w, platinum_14w, platinum_16w, platinum_18w, platinum_20w, platinum_25w,
        platinum_30w, platinum_35w, platinum_40w, platinum_45w, platinum_50w, platinum_60w, platinum_70w, platinum_80w,
        platinum_90w, platinum_100w.
        """
        return pulumi.get(self, "sku_code")

    @sku_code.setter
    def sku_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "sku_code", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        Subnet id.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        VPC id.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Public network bandwidth. `bandwidth` must be greater than zero when `enable_public` equal true.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="enablePublic")
    def enable_public(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the public network. Must set `bandwidth` when `enable_public` equal true.
        """
        return pulumi.get(self, "enable_public")

    @enable_public.setter
    def enable_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_public", value)

    @property
    @pulumi.getter(name="ipRules")
    def ip_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqInstanceIpRuleArgs']]]]:
        """
        Public network access whitelist.
        """
        return pulumi.get(self, "ip_rules")

    @ip_rules.setter
    def ip_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqInstanceIpRuleArgs']]]]):
        pulumi.set(self, "ip_rules", value)

    @property
    @pulumi.getter(name="messageRetention")
    def message_retention(self) -> Optional[pulumi.Input[int]]:
        """
        Message retention time in hours.
        """
        return pulumi.get(self, "message_retention")

    @message_retention.setter
    def message_retention(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_retention", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Instance name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remark.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _RocketmqInstanceState:
    def __init__(__self__, *,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 enable_public: Optional[pulumi.Input[bool]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqInstanceIpRuleArgs']]]] = None,
                 message_retention: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_end_point: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 sku_code: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_end_point: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RocketmqInstance resources.
        :param pulumi.Input[int] bandwidth: Public network bandwidth. `bandwidth` must be greater than zero when `enable_public` equal true.
        :param pulumi.Input[bool] enable_public: Whether to enable the public network. Must set `bandwidth` when `enable_public` equal true.
        :param pulumi.Input[str] instance_type: Instance type. Valid values: `EXPERIMENT`, `BASIC`, `PRO`, `PLATINUM`.
        :param pulumi.Input[Sequence[pulumi.Input['RocketmqInstanceIpRuleArgs']]] ip_rules: Public network access whitelist.
        :param pulumi.Input[int] message_retention: Message retention time in hours.
        :param pulumi.Input[str] name: Instance name.
        :param pulumi.Input[str] public_end_point: Public network access address.
        :param pulumi.Input[str] remark: Remark.
        :param pulumi.Input[str] sku_code: SKU code. Available specifications are as follows: experiment_500, basic_1k, basic_2k, basic_3k, basic_4k, basic_5k,
               basic_6k, basic_7k, basic_8k, basic_9k, basic_10k, pro_4k, pro_6k, pro_8k, pro_1w, pro_15k, pro_2w, pro_25k, pro_3w,
               pro_35k, pro_4w, pro_45k, pro_5w, pro_55k, pro_60k, pro_65k, pro_70k, pro_75k, pro_80k, pro_85k, pro_90k, pro_95k,
               pro_100k, platinum_1w, platinum_2w, platinum_3w, platinum_4w, platinum_5w, platinum_6w, platinum_7w, platinum_8w,
               platinum_9w, platinum_10w, platinum_12w, platinum_14w, platinum_16w, platinum_18w, platinum_20w, platinum_25w,
               platinum_30w, platinum_35w, platinum_40w, platinum_45w, platinum_50w, platinum_60w, platinum_70w, platinum_80w,
               platinum_90w, platinum_100w.
        :param pulumi.Input[str] subnet_id: Subnet id.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        :param pulumi.Input[str] vpc_end_point: VPC access address.
        :param pulumi.Input[str] vpc_id: VPC id.
        """
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if enable_public is not None:
            pulumi.set(__self__, "enable_public", enable_public)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if ip_rules is not None:
            pulumi.set(__self__, "ip_rules", ip_rules)
        if message_retention is not None:
            pulumi.set(__self__, "message_retention", message_retention)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public_end_point is not None:
            pulumi.set(__self__, "public_end_point", public_end_point)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if sku_code is not None:
            pulumi.set(__self__, "sku_code", sku_code)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_end_point is not None:
            pulumi.set(__self__, "vpc_end_point", vpc_end_point)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Public network bandwidth. `bandwidth` must be greater than zero when `enable_public` equal true.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="enablePublic")
    def enable_public(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the public network. Must set `bandwidth` when `enable_public` equal true.
        """
        return pulumi.get(self, "enable_public")

    @enable_public.setter
    def enable_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_public", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        Instance type. Valid values: `EXPERIMENT`, `BASIC`, `PRO`, `PLATINUM`.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="ipRules")
    def ip_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqInstanceIpRuleArgs']]]]:
        """
        Public network access whitelist.
        """
        return pulumi.get(self, "ip_rules")

    @ip_rules.setter
    def ip_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqInstanceIpRuleArgs']]]]):
        pulumi.set(self, "ip_rules", value)

    @property
    @pulumi.getter(name="messageRetention")
    def message_retention(self) -> Optional[pulumi.Input[int]]:
        """
        Message retention time in hours.
        """
        return pulumi.get(self, "message_retention")

    @message_retention.setter
    def message_retention(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_retention", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Instance name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="publicEndPoint")
    def public_end_point(self) -> Optional[pulumi.Input[str]]:
        """
        Public network access address.
        """
        return pulumi.get(self, "public_end_point")

    @public_end_point.setter
    def public_end_point(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_end_point", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remark.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="skuCode")
    def sku_code(self) -> Optional[pulumi.Input[str]]:
        """
        SKU code. Available specifications are as follows: experiment_500, basic_1k, basic_2k, basic_3k, basic_4k, basic_5k,
        basic_6k, basic_7k, basic_8k, basic_9k, basic_10k, pro_4k, pro_6k, pro_8k, pro_1w, pro_15k, pro_2w, pro_25k, pro_3w,
        pro_35k, pro_4w, pro_45k, pro_5w, pro_55k, pro_60k, pro_65k, pro_70k, pro_75k, pro_80k, pro_85k, pro_90k, pro_95k,
        pro_100k, platinum_1w, platinum_2w, platinum_3w, platinum_4w, platinum_5w, platinum_6w, platinum_7w, platinum_8w,
        platinum_9w, platinum_10w, platinum_12w, platinum_14w, platinum_16w, platinum_18w, platinum_20w, platinum_25w,
        platinum_30w, platinum_35w, platinum_40w, platinum_45w, platinum_50w, platinum_60w, platinum_70w, platinum_80w,
        platinum_90w, platinum_100w.
        """
        return pulumi.get(self, "sku_code")

    @sku_code.setter
    def sku_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sku_code", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet id.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcEndPoint")
    def vpc_end_point(self) -> Optional[pulumi.Input[str]]:
        """
        VPC access address.
        """
        return pulumi.get(self, "vpc_end_point")

    @vpc_end_point.setter
    def vpc_end_point(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_end_point", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC id.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class RocketmqInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 enable_public: Optional[pulumi.Input[bool]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RocketmqInstanceIpRuleArgs']]]]] = None,
                 message_retention: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 sku_code: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a RocketmqInstance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: Public network bandwidth. `bandwidth` must be greater than zero when `enable_public` equal true.
        :param pulumi.Input[bool] enable_public: Whether to enable the public network. Must set `bandwidth` when `enable_public` equal true.
        :param pulumi.Input[str] instance_type: Instance type. Valid values: `EXPERIMENT`, `BASIC`, `PRO`, `PLATINUM`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RocketmqInstanceIpRuleArgs']]]] ip_rules: Public network access whitelist.
        :param pulumi.Input[int] message_retention: Message retention time in hours.
        :param pulumi.Input[str] name: Instance name.
        :param pulumi.Input[str] remark: Remark.
        :param pulumi.Input[str] sku_code: SKU code. Available specifications are as follows: experiment_500, basic_1k, basic_2k, basic_3k, basic_4k, basic_5k,
               basic_6k, basic_7k, basic_8k, basic_9k, basic_10k, pro_4k, pro_6k, pro_8k, pro_1w, pro_15k, pro_2w, pro_25k, pro_3w,
               pro_35k, pro_4w, pro_45k, pro_5w, pro_55k, pro_60k, pro_65k, pro_70k, pro_75k, pro_80k, pro_85k, pro_90k, pro_95k,
               pro_100k, platinum_1w, platinum_2w, platinum_3w, platinum_4w, platinum_5w, platinum_6w, platinum_7w, platinum_8w,
               platinum_9w, platinum_10w, platinum_12w, platinum_14w, platinum_16w, platinum_18w, platinum_20w, platinum_25w,
               platinum_30w, platinum_35w, platinum_40w, platinum_45w, platinum_50w, platinum_60w, platinum_70w, platinum_80w,
               platinum_90w, platinum_100w.
        :param pulumi.Input[str] subnet_id: Subnet id.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        :param pulumi.Input[str] vpc_id: VPC id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RocketmqInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RocketmqInstance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RocketmqInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RocketmqInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 enable_public: Optional[pulumi.Input[bool]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RocketmqInstanceIpRuleArgs']]]]] = None,
                 message_retention: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 sku_code: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RocketmqInstanceArgs.__new__(RocketmqInstanceArgs)

            __props__.__dict__["bandwidth"] = bandwidth
            __props__.__dict__["enable_public"] = enable_public
            if instance_type is None and not opts.urn:
                raise TypeError("Missing required property 'instance_type'")
            __props__.__dict__["instance_type"] = instance_type
            __props__.__dict__["ip_rules"] = ip_rules
            __props__.__dict__["message_retention"] = message_retention
            __props__.__dict__["name"] = name
            __props__.__dict__["remark"] = remark
            if sku_code is None and not opts.urn:
                raise TypeError("Missing required property 'sku_code'")
            __props__.__dict__["sku_code"] = sku_code
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["public_end_point"] = None
            __props__.__dict__["vpc_end_point"] = None
        super(RocketmqInstance, __self__).__init__(
            'tencentcloud:Trocket/rocketmqInstance:RocketmqInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth: Optional[pulumi.Input[int]] = None,
            enable_public: Optional[pulumi.Input[bool]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RocketmqInstanceIpRuleArgs']]]]] = None,
            message_retention: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            public_end_point: Optional[pulumi.Input[str]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            sku_code: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vpc_end_point: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'RocketmqInstance':
        """
        Get an existing RocketmqInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: Public network bandwidth. `bandwidth` must be greater than zero when `enable_public` equal true.
        :param pulumi.Input[bool] enable_public: Whether to enable the public network. Must set `bandwidth` when `enable_public` equal true.
        :param pulumi.Input[str] instance_type: Instance type. Valid values: `EXPERIMENT`, `BASIC`, `PRO`, `PLATINUM`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RocketmqInstanceIpRuleArgs']]]] ip_rules: Public network access whitelist.
        :param pulumi.Input[int] message_retention: Message retention time in hours.
        :param pulumi.Input[str] name: Instance name.
        :param pulumi.Input[str] public_end_point: Public network access address.
        :param pulumi.Input[str] remark: Remark.
        :param pulumi.Input[str] sku_code: SKU code. Available specifications are as follows: experiment_500, basic_1k, basic_2k, basic_3k, basic_4k, basic_5k,
               basic_6k, basic_7k, basic_8k, basic_9k, basic_10k, pro_4k, pro_6k, pro_8k, pro_1w, pro_15k, pro_2w, pro_25k, pro_3w,
               pro_35k, pro_4w, pro_45k, pro_5w, pro_55k, pro_60k, pro_65k, pro_70k, pro_75k, pro_80k, pro_85k, pro_90k, pro_95k,
               pro_100k, platinum_1w, platinum_2w, platinum_3w, platinum_4w, platinum_5w, platinum_6w, platinum_7w, platinum_8w,
               platinum_9w, platinum_10w, platinum_12w, platinum_14w, platinum_16w, platinum_18w, platinum_20w, platinum_25w,
               platinum_30w, platinum_35w, platinum_40w, platinum_45w, platinum_50w, platinum_60w, platinum_70w, platinum_80w,
               platinum_90w, platinum_100w.
        :param pulumi.Input[str] subnet_id: Subnet id.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        :param pulumi.Input[str] vpc_end_point: VPC access address.
        :param pulumi.Input[str] vpc_id: VPC id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RocketmqInstanceState.__new__(_RocketmqInstanceState)

        __props__.__dict__["bandwidth"] = bandwidth
        __props__.__dict__["enable_public"] = enable_public
        __props__.__dict__["instance_type"] = instance_type
        __props__.__dict__["ip_rules"] = ip_rules
        __props__.__dict__["message_retention"] = message_retention
        __props__.__dict__["name"] = name
        __props__.__dict__["public_end_point"] = public_end_point
        __props__.__dict__["remark"] = remark
        __props__.__dict__["sku_code"] = sku_code
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_end_point"] = vpc_end_point
        __props__.__dict__["vpc_id"] = vpc_id
        return RocketmqInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[int]:
        """
        Public network bandwidth. `bandwidth` must be greater than zero when `enable_public` equal true.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="enablePublic")
    def enable_public(self) -> pulumi.Output[bool]:
        """
        Whether to enable the public network. Must set `bandwidth` when `enable_public` equal true.
        """
        return pulumi.get(self, "enable_public")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        """
        Instance type. Valid values: `EXPERIMENT`, `BASIC`, `PRO`, `PLATINUM`.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="ipRules")
    def ip_rules(self) -> pulumi.Output[Sequence['outputs.RocketmqInstanceIpRule']]:
        """
        Public network access whitelist.
        """
        return pulumi.get(self, "ip_rules")

    @property
    @pulumi.getter(name="messageRetention")
    def message_retention(self) -> pulumi.Output[int]:
        """
        Message retention time in hours.
        """
        return pulumi.get(self, "message_retention")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Instance name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicEndPoint")
    def public_end_point(self) -> pulumi.Output[str]:
        """
        Public network access address.
        """
        return pulumi.get(self, "public_end_point")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Remark.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="skuCode")
    def sku_code(self) -> pulumi.Output[str]:
        """
        SKU code. Available specifications are as follows: experiment_500, basic_1k, basic_2k, basic_3k, basic_4k, basic_5k,
        basic_6k, basic_7k, basic_8k, basic_9k, basic_10k, pro_4k, pro_6k, pro_8k, pro_1w, pro_15k, pro_2w, pro_25k, pro_3w,
        pro_35k, pro_4w, pro_45k, pro_5w, pro_55k, pro_60k, pro_65k, pro_70k, pro_75k, pro_80k, pro_85k, pro_90k, pro_95k,
        pro_100k, platinum_1w, platinum_2w, platinum_3w, platinum_4w, platinum_5w, platinum_6w, platinum_7w, platinum_8w,
        platinum_9w, platinum_10w, platinum_12w, platinum_14w, platinum_16w, platinum_18w, platinum_20w, platinum_25w,
        platinum_30w, platinum_35w, platinum_40w, platinum_45w, platinum_50w, platinum_60w, platinum_70w, platinum_80w,
        platinum_90w, platinum_100w.
        """
        return pulumi.get(self, "sku_code")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        Subnet id.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcEndPoint")
    def vpc_end_point(self) -> pulumi.Output[str]:
        """
        VPC access address.
        """
        return pulumi.get(self, "vpc_end_point")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        VPC id.
        """
        return pulumi.get(self, "vpc_id")

