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

__all__ = ['BackupStrategyArgs', 'BackupStrategy']

@pulumi.input_type
class BackupStrategyArgs:
    def __init__(__self__, *,
                 data_backup_strategies: pulumi.Input[Sequence[pulumi.Input['BackupStrategyDataBackupStrategyArgs']]],
                 instance_id: pulumi.Input[str],
                 meta_backup_strategies: Optional[pulumi.Input[Sequence[pulumi.Input['BackupStrategyMetaBackupStrategyArgs']]]] = None):
        """
        The set of arguments for constructing a BackupStrategy resource.
        :param pulumi.Input[Sequence[pulumi.Input['BackupStrategyDataBackupStrategyArgs']]] data_backup_strategies: Data backup strategy.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[Sequence[pulumi.Input['BackupStrategyMetaBackupStrategyArgs']]] meta_backup_strategies: Metadata backup strategy.
        """
        pulumi.set(__self__, "data_backup_strategies", data_backup_strategies)
        pulumi.set(__self__, "instance_id", instance_id)
        if meta_backup_strategies is not None:
            pulumi.set(__self__, "meta_backup_strategies", meta_backup_strategies)

    @property
    @pulumi.getter(name="dataBackupStrategies")
    def data_backup_strategies(self) -> pulumi.Input[Sequence[pulumi.Input['BackupStrategyDataBackupStrategyArgs']]]:
        """
        Data backup strategy.
        """
        return pulumi.get(self, "data_backup_strategies")

    @data_backup_strategies.setter
    def data_backup_strategies(self, value: pulumi.Input[Sequence[pulumi.Input['BackupStrategyDataBackupStrategyArgs']]]):
        pulumi.set(self, "data_backup_strategies", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="metaBackupStrategies")
    def meta_backup_strategies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BackupStrategyMetaBackupStrategyArgs']]]]:
        """
        Metadata backup strategy.
        """
        return pulumi.get(self, "meta_backup_strategies")

    @meta_backup_strategies.setter
    def meta_backup_strategies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BackupStrategyMetaBackupStrategyArgs']]]]):
        pulumi.set(self, "meta_backup_strategies", value)


@pulumi.input_type
class _BackupStrategyState:
    def __init__(__self__, *,
                 data_backup_strategies: Optional[pulumi.Input[Sequence[pulumi.Input['BackupStrategyDataBackupStrategyArgs']]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 meta_backup_strategies: Optional[pulumi.Input[Sequence[pulumi.Input['BackupStrategyMetaBackupStrategyArgs']]]] = None):
        """
        Input properties used for looking up and filtering BackupStrategy resources.
        :param pulumi.Input[Sequence[pulumi.Input['BackupStrategyDataBackupStrategyArgs']]] data_backup_strategies: Data backup strategy.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[Sequence[pulumi.Input['BackupStrategyMetaBackupStrategyArgs']]] meta_backup_strategies: Metadata backup strategy.
        """
        if data_backup_strategies is not None:
            pulumi.set(__self__, "data_backup_strategies", data_backup_strategies)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if meta_backup_strategies is not None:
            pulumi.set(__self__, "meta_backup_strategies", meta_backup_strategies)

    @property
    @pulumi.getter(name="dataBackupStrategies")
    def data_backup_strategies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BackupStrategyDataBackupStrategyArgs']]]]:
        """
        Data backup strategy.
        """
        return pulumi.get(self, "data_backup_strategies")

    @data_backup_strategies.setter
    def data_backup_strategies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BackupStrategyDataBackupStrategyArgs']]]]):
        pulumi.set(self, "data_backup_strategies", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="metaBackupStrategies")
    def meta_backup_strategies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BackupStrategyMetaBackupStrategyArgs']]]]:
        """
        Metadata backup strategy.
        """
        return pulumi.get(self, "meta_backup_strategies")

    @meta_backup_strategies.setter
    def meta_backup_strategies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BackupStrategyMetaBackupStrategyArgs']]]]):
        pulumi.set(self, "meta_backup_strategies", value)


class BackupStrategy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_backup_strategies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupStrategyDataBackupStrategyArgs']]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 meta_backup_strategies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupStrategyMetaBackupStrategyArgs']]]]] = None,
                 __props__=None):
        """
        Create a BackupStrategy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupStrategyDataBackupStrategyArgs']]]] data_backup_strategies: Data backup strategy.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupStrategyMetaBackupStrategyArgs']]]] meta_backup_strategies: Metadata backup strategy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackupStrategyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a BackupStrategy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param BackupStrategyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackupStrategyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_backup_strategies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupStrategyDataBackupStrategyArgs']]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 meta_backup_strategies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupStrategyMetaBackupStrategyArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackupStrategyArgs.__new__(BackupStrategyArgs)

            if data_backup_strategies is None and not opts.urn:
                raise TypeError("Missing required property 'data_backup_strategies'")
            __props__.__dict__["data_backup_strategies"] = data_backup_strategies
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["meta_backup_strategies"] = meta_backup_strategies
        super(BackupStrategy, __self__).__init__(
            'tencentcloud:Clickhouse/backupStrategy:BackupStrategy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data_backup_strategies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupStrategyDataBackupStrategyArgs']]]]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            meta_backup_strategies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupStrategyMetaBackupStrategyArgs']]]]] = None) -> 'BackupStrategy':
        """
        Get an existing BackupStrategy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupStrategyDataBackupStrategyArgs']]]] data_backup_strategies: Data backup strategy.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupStrategyMetaBackupStrategyArgs']]]] meta_backup_strategies: Metadata backup strategy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BackupStrategyState.__new__(_BackupStrategyState)

        __props__.__dict__["data_backup_strategies"] = data_backup_strategies
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["meta_backup_strategies"] = meta_backup_strategies
        return BackupStrategy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataBackupStrategies")
    def data_backup_strategies(self) -> pulumi.Output[Sequence['outputs.BackupStrategyDataBackupStrategy']]:
        """
        Data backup strategy.
        """
        return pulumi.get(self, "data_backup_strategies")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="metaBackupStrategies")
    def meta_backup_strategies(self) -> pulumi.Output[Sequence['outputs.BackupStrategyMetaBackupStrategy']]:
        """
        Metadata backup strategy.
        """
        return pulumi.get(self, "meta_backup_strategies")
