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
    'GetInstanceProjectResult',
]

@pulumi.output_type
class GetInstanceProjectResult(dict):
    def __init__(__self__, *,
                 create_time: str,
                 creator_uin: int,
                 project_id: int,
                 project_info: str,
                 project_name: str):
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "creator_uin", creator_uin)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "project_info", project_info)
        pulumi.set(__self__, "project_name", project_name)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="creatorUin")
    def creator_uin(self) -> int:
        return pulumi.get(self, "creator_uin")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="projectInfo")
    def project_info(self) -> str:
        return pulumi.get(self, "project_info")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> str:
        return pulumi.get(self, "project_name")


