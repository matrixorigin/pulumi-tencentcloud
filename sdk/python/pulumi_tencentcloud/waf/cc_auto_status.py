# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CcAutoStatusArgs', 'CcAutoStatus']

@pulumi.input_type
class CcAutoStatusArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 edition: pulumi.Input[str]):
        """
        The set of arguments for constructing a CcAutoStatus resource.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] edition: Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "edition", edition)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        Domain.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def edition(self) -> pulumi.Input[str]:
        """
        Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        """
        return pulumi.get(self, "edition")

    @edition.setter
    def edition(self, value: pulumi.Input[str]):
        pulumi.set(self, "edition", value)


@pulumi.input_type
class _CcAutoStatusState:
    def __init__(__self__, *,
                 domain: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering CcAutoStatus resources.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] edition: Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        :param pulumi.Input[int] status: cc auto status, 1 means open, 0 means close.
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if edition is not None:
            pulumi.set(__self__, "edition", edition)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def edition(self) -> Optional[pulumi.Input[str]]:
        """
        Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        """
        return pulumi.get(self, "edition")

    @edition.setter
    def edition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edition", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[int]]:
        """
        cc auto status, 1 means open, 0 means close.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "status", value)


class CcAutoStatus(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a CcAutoStatus resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] edition: Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CcAutoStatusArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CcAutoStatus resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CcAutoStatusArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CcAutoStatusArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CcAutoStatusArgs.__new__(CcAutoStatusArgs)

            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if edition is None and not opts.urn:
                raise TypeError("Missing required property 'edition'")
            __props__.__dict__["edition"] = edition
            __props__.__dict__["status"] = None
        super(CcAutoStatus, __self__).__init__(
            'tencentcloud:Waf/ccAutoStatus:CcAutoStatus',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain: Optional[pulumi.Input[str]] = None,
            edition: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[int]] = None) -> 'CcAutoStatus':
        """
        Get an existing CcAutoStatus resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] edition: Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        :param pulumi.Input[int] status: cc auto status, 1 means open, 0 means close.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CcAutoStatusState.__new__(_CcAutoStatusState)

        __props__.__dict__["domain"] = domain
        __props__.__dict__["edition"] = edition
        __props__.__dict__["status"] = status
        return CcAutoStatus(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        Domain.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def edition(self) -> pulumi.Output[str]:
        """
        Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        """
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[int]:
        """
        cc auto status, 1 means open, 0 means close.
        """
        return pulumi.get(self, "status")
