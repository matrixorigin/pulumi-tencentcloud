# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApplyDiskBackupArgs', 'ApplyDiskBackup']

@pulumi.input_type
class ApplyDiskBackupArgs:
    def __init__(__self__, *,
                 disk_backup_id: pulumi.Input[str],
                 disk_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ApplyDiskBackup resource.
        :param pulumi.Input[str] disk_backup_id: Disk backup ID.
        :param pulumi.Input[str] disk_id: Disk ID.
        """
        pulumi.set(__self__, "disk_backup_id", disk_backup_id)
        pulumi.set(__self__, "disk_id", disk_id)

    @property
    @pulumi.getter(name="diskBackupId")
    def disk_backup_id(self) -> pulumi.Input[str]:
        """
        Disk backup ID.
        """
        return pulumi.get(self, "disk_backup_id")

    @disk_backup_id.setter
    def disk_backup_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "disk_backup_id", value)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> pulumi.Input[str]:
        """
        Disk ID.
        """
        return pulumi.get(self, "disk_id")

    @disk_id.setter
    def disk_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "disk_id", value)


@pulumi.input_type
class _ApplyDiskBackupState:
    def __init__(__self__, *,
                 disk_backup_id: Optional[pulumi.Input[str]] = None,
                 disk_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplyDiskBackup resources.
        :param pulumi.Input[str] disk_backup_id: Disk backup ID.
        :param pulumi.Input[str] disk_id: Disk ID.
        """
        if disk_backup_id is not None:
            pulumi.set(__self__, "disk_backup_id", disk_backup_id)
        if disk_id is not None:
            pulumi.set(__self__, "disk_id", disk_id)

    @property
    @pulumi.getter(name="diskBackupId")
    def disk_backup_id(self) -> Optional[pulumi.Input[str]]:
        """
        Disk backup ID.
        """
        return pulumi.get(self, "disk_backup_id")

    @disk_backup_id.setter
    def disk_backup_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_backup_id", value)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> Optional[pulumi.Input[str]]:
        """
        Disk ID.
        """
        return pulumi.get(self, "disk_id")

    @disk_id.setter
    def disk_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_id", value)


class ApplyDiskBackup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk_backup_id: Optional[pulumi.Input[str]] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ApplyDiskBackup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_backup_id: Disk backup ID.
        :param pulumi.Input[str] disk_id: Disk ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplyDiskBackupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ApplyDiskBackup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ApplyDiskBackupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplyDiskBackupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk_backup_id: Optional[pulumi.Input[str]] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplyDiskBackupArgs.__new__(ApplyDiskBackupArgs)

            if disk_backup_id is None and not opts.urn:
                raise TypeError("Missing required property 'disk_backup_id'")
            __props__.__dict__["disk_backup_id"] = disk_backup_id
            if disk_id is None and not opts.urn:
                raise TypeError("Missing required property 'disk_id'")
            __props__.__dict__["disk_id"] = disk_id
        super(ApplyDiskBackup, __self__).__init__(
            'tencentcloud:Lighthouse/applyDiskBackup:ApplyDiskBackup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disk_backup_id: Optional[pulumi.Input[str]] = None,
            disk_id: Optional[pulumi.Input[str]] = None) -> 'ApplyDiskBackup':
        """
        Get an existing ApplyDiskBackup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_backup_id: Disk backup ID.
        :param pulumi.Input[str] disk_id: Disk ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplyDiskBackupState.__new__(_ApplyDiskBackupState)

        __props__.__dict__["disk_backup_id"] = disk_backup_id
        __props__.__dict__["disk_id"] = disk_id
        return ApplyDiskBackup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="diskBackupId")
    def disk_backup_id(self) -> pulumi.Output[str]:
        """
        Disk backup ID.
        """
        return pulumi.get(self, "disk_backup_id")

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> pulumi.Output[str]:
        """
        Disk ID.
        """
        return pulumi.get(self, "disk_id")
