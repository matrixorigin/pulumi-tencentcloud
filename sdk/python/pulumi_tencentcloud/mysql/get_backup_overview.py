# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetBackupOverviewResult',
    'AwaitableGetBackupOverviewResult',
    'get_backup_overview',
    'get_backup_overview_output',
]

@pulumi.output_type
class GetBackupOverviewResult:
    """
    A collection of values returned by getBackupOverview.
    """
    def __init__(__self__, backup_archive_volume=None, backup_count=None, backup_standby_volume=None, backup_volume=None, billing_volume=None, free_volume=None, id=None, product=None, remote_backup_volume=None, result_output_file=None):
        if backup_archive_volume and not isinstance(backup_archive_volume, int):
            raise TypeError("Expected argument 'backup_archive_volume' to be a int")
        pulumi.set(__self__, "backup_archive_volume", backup_archive_volume)
        if backup_count and not isinstance(backup_count, int):
            raise TypeError("Expected argument 'backup_count' to be a int")
        pulumi.set(__self__, "backup_count", backup_count)
        if backup_standby_volume and not isinstance(backup_standby_volume, int):
            raise TypeError("Expected argument 'backup_standby_volume' to be a int")
        pulumi.set(__self__, "backup_standby_volume", backup_standby_volume)
        if backup_volume and not isinstance(backup_volume, int):
            raise TypeError("Expected argument 'backup_volume' to be a int")
        pulumi.set(__self__, "backup_volume", backup_volume)
        if billing_volume and not isinstance(billing_volume, int):
            raise TypeError("Expected argument 'billing_volume' to be a int")
        pulumi.set(__self__, "billing_volume", billing_volume)
        if free_volume and not isinstance(free_volume, int):
            raise TypeError("Expected argument 'free_volume' to be a int")
        pulumi.set(__self__, "free_volume", free_volume)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if product and not isinstance(product, str):
            raise TypeError("Expected argument 'product' to be a str")
        pulumi.set(__self__, "product", product)
        if remote_backup_volume and not isinstance(remote_backup_volume, int):
            raise TypeError("Expected argument 'remote_backup_volume' to be a int")
        pulumi.set(__self__, "remote_backup_volume", remote_backup_volume)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="backupArchiveVolume")
    def backup_archive_volume(self) -> int:
        return pulumi.get(self, "backup_archive_volume")

    @property
    @pulumi.getter(name="backupCount")
    def backup_count(self) -> int:
        return pulumi.get(self, "backup_count")

    @property
    @pulumi.getter(name="backupStandbyVolume")
    def backup_standby_volume(self) -> int:
        return pulumi.get(self, "backup_standby_volume")

    @property
    @pulumi.getter(name="backupVolume")
    def backup_volume(self) -> int:
        return pulumi.get(self, "backup_volume")

    @property
    @pulumi.getter(name="billingVolume")
    def billing_volume(self) -> int:
        return pulumi.get(self, "billing_volume")

    @property
    @pulumi.getter(name="freeVolume")
    def free_volume(self) -> int:
        return pulumi.get(self, "free_volume")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def product(self) -> str:
        return pulumi.get(self, "product")

    @property
    @pulumi.getter(name="remoteBackupVolume")
    def remote_backup_volume(self) -> int:
        return pulumi.get(self, "remote_backup_volume")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetBackupOverviewResult(GetBackupOverviewResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupOverviewResult(
            backup_archive_volume=self.backup_archive_volume,
            backup_count=self.backup_count,
            backup_standby_volume=self.backup_standby_volume,
            backup_volume=self.backup_volume,
            billing_volume=self.billing_volume,
            free_volume=self.free_volume,
            id=self.id,
            product=self.product,
            remote_backup_volume=self.remote_backup_volume,
            result_output_file=self.result_output_file)


def get_backup_overview(product: Optional[str] = None,
                        result_output_file: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupOverviewResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['product'] = product
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Mysql/getBackupOverview:getBackupOverview', __args__, opts=opts, typ=GetBackupOverviewResult).value

    return AwaitableGetBackupOverviewResult(
        backup_archive_volume=pulumi.get(__ret__, 'backup_archive_volume'),
        backup_count=pulumi.get(__ret__, 'backup_count'),
        backup_standby_volume=pulumi.get(__ret__, 'backup_standby_volume'),
        backup_volume=pulumi.get(__ret__, 'backup_volume'),
        billing_volume=pulumi.get(__ret__, 'billing_volume'),
        free_volume=pulumi.get(__ret__, 'free_volume'),
        id=pulumi.get(__ret__, 'id'),
        product=pulumi.get(__ret__, 'product'),
        remote_backup_volume=pulumi.get(__ret__, 'remote_backup_volume'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_backup_overview)
def get_backup_overview_output(product: Optional[pulumi.Input[str]] = None,
                               result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackupOverviewResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
