# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CheckCertificateChainOperationArgs', 'CheckCertificateChainOperation']

@pulumi.input_type
class CheckCertificateChainOperationArgs:
    def __init__(__self__, *,
                 certificate_chain: pulumi.Input[str]):
        """
        The set of arguments for constructing a CheckCertificateChainOperation resource.
        :param pulumi.Input[str] certificate_chain: The certificate chain to check.
        """
        pulumi.set(__self__, "certificate_chain", certificate_chain)

    @property
    @pulumi.getter(name="certificateChain")
    def certificate_chain(self) -> pulumi.Input[str]:
        """
        The certificate chain to check.
        """
        return pulumi.get(self, "certificate_chain")

    @certificate_chain.setter
    def certificate_chain(self, value: pulumi.Input[str]):
        pulumi.set(self, "certificate_chain", value)


@pulumi.input_type
class _CheckCertificateChainOperationState:
    def __init__(__self__, *,
                 certificate_chain: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CheckCertificateChainOperation resources.
        :param pulumi.Input[str] certificate_chain: The certificate chain to check.
        """
        if certificate_chain is not None:
            pulumi.set(__self__, "certificate_chain", certificate_chain)

    @property
    @pulumi.getter(name="certificateChain")
    def certificate_chain(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate chain to check.
        """
        return pulumi.get(self, "certificate_chain")

    @certificate_chain.setter
    def certificate_chain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_chain", value)


class CheckCertificateChainOperation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_chain: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a CheckCertificateChainOperation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_chain: The certificate chain to check.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CheckCertificateChainOperationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CheckCertificateChainOperation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CheckCertificateChainOperationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CheckCertificateChainOperationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_chain: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CheckCertificateChainOperationArgs.__new__(CheckCertificateChainOperationArgs)

            if certificate_chain is None and not opts.urn:
                raise TypeError("Missing required property 'certificate_chain'")
            __props__.__dict__["certificate_chain"] = certificate_chain
        super(CheckCertificateChainOperation, __self__).__init__(
            'tencentcloud:Ssl/checkCertificateChainOperation:CheckCertificateChainOperation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate_chain: Optional[pulumi.Input[str]] = None) -> 'CheckCertificateChainOperation':
        """
        Get an existing CheckCertificateChainOperation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_chain: The certificate chain to check.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CheckCertificateChainOperationState.__new__(_CheckCertificateChainOperationState)

        __props__.__dict__["certificate_chain"] = certificate_chain
        return CheckCertificateChainOperation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certificateChain")
    def certificate_chain(self) -> pulumi.Output[str]:
        """
        The certificate chain to check.
        """
        return pulumi.get(self, "certificate_chain")
