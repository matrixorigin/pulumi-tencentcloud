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
    'GetGroupsPlacementGroupListResult',
]

@pulumi.output_type
class GetGroupsPlacementGroupListResult(dict):
    def __init__(__self__, *,
                 create_time: str,
                 current_num: int,
                 cvm_quota_total: int,
                 instance_ids: Sequence[str],
                 name: str,
                 placement_group_id: str,
                 type: str):
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "current_num", current_num)
        pulumi.set(__self__, "cvm_quota_total", cvm_quota_total)
        pulumi.set(__self__, "instance_ids", instance_ids)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "placement_group_id", placement_group_id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="currentNum")
    def current_num(self) -> int:
        return pulumi.get(self, "current_num")

    @property
    @pulumi.getter(name="cvmQuotaTotal")
    def cvm_quota_total(self) -> int:
        return pulumi.get(self, "cvm_quota_total")

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Sequence[str]:
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="placementGroupId")
    def placement_group_id(self) -> str:
        return pulumi.get(self, "placement_group_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


