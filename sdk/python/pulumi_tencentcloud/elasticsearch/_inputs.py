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
    'DiagnoseDiagnoseJobMetaArgs',
    'InstanceEsAclArgs',
    'InstanceMultiZoneInfoArgs',
    'InstanceNodeInfoListArgs',
    'InstanceWebNodeTypeInfoArgs',
    'LogstashOperationDurationArgs',
    'LogstashPipelinePipelineArgs',
]

@pulumi.input_type
class DiagnoseDiagnoseJobMetaArgs:
    def __init__(__self__, *,
                 job_description: Optional[pulumi.Input[str]] = None,
                 job_name: Optional[pulumi.Input[str]] = None,
                 job_zh_name: Optional[pulumi.Input[str]] = None):
        if job_description is not None:
            pulumi.set(__self__, "job_description", job_description)
        if job_name is not None:
            pulumi.set(__self__, "job_name", job_name)
        if job_zh_name is not None:
            pulumi.set(__self__, "job_zh_name", job_zh_name)

    @property
    @pulumi.getter(name="jobDescription")
    def job_description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "job_description")

    @job_description.setter
    def job_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_description", value)

    @property
    @pulumi.getter(name="jobName")
    def job_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "job_name")

    @job_name.setter
    def job_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_name", value)

    @property
    @pulumi.getter(name="jobZhName")
    def job_zh_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "job_zh_name")

    @job_zh_name.setter
    def job_zh_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_zh_name", value)


@pulumi.input_type
class InstanceEsAclArgs:
    def __init__(__self__, *,
                 black_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 white_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if black_lists is not None:
            pulumi.set(__self__, "black_lists", black_lists)
        if white_lists is not None:
            pulumi.set(__self__, "white_lists", white_lists)

    @property
    @pulumi.getter(name="blackLists")
    def black_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "black_lists")

    @black_lists.setter
    def black_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "black_lists", value)

    @property
    @pulumi.getter(name="whiteLists")
    def white_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "white_lists")

    @white_lists.setter
    def white_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "white_lists", value)


@pulumi.input_type
class InstanceMultiZoneInfoArgs:
    def __init__(__self__, *,
                 availability_zone: pulumi.Input[str],
                 subnet_id: pulumi.Input[str]):
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Input[str]:
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class InstanceNodeInfoListArgs:
    def __init__(__self__, *,
                 node_num: pulumi.Input[int],
                 node_type: pulumi.Input[str],
                 disk_size: Optional[pulumi.Input[int]] = None,
                 disk_type: Optional[pulumi.Input[str]] = None,
                 encrypt: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "node_num", node_num)
        pulumi.set(__self__, "node_type", node_type)
        if disk_size is not None:
            pulumi.set(__self__, "disk_size", disk_size)
        if disk_type is not None:
            pulumi.set(__self__, "disk_type", disk_type)
        if encrypt is not None:
            pulumi.set(__self__, "encrypt", encrypt)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="nodeNum")
    def node_num(self) -> pulumi.Input[int]:
        return pulumi.get(self, "node_num")

    @node_num.setter
    def node_num(self, value: pulumi.Input[int]):
        pulumi.set(self, "node_num", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_type", value)

    @property
    @pulumi.getter(name="diskSize")
    def disk_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "disk_size")

    @disk_size.setter
    def disk_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_size", value)

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "disk_type")

    @disk_type.setter
    def disk_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_type", value)

    @property
    @pulumi.getter
    def encrypt(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "encrypt")

    @encrypt.setter
    def encrypt(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "encrypt", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class InstanceWebNodeTypeInfoArgs:
    def __init__(__self__, *,
                 node_num: pulumi.Input[int],
                 node_type: pulumi.Input[str]):
        pulumi.set(__self__, "node_num", node_num)
        pulumi.set(__self__, "node_type", node_type)

    @property
    @pulumi.getter(name="nodeNum")
    def node_num(self) -> pulumi.Input[int]:
        return pulumi.get(self, "node_num")

    @node_num.setter
    def node_num(self, value: pulumi.Input[int]):
        pulumi.set(self, "node_num", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_type", value)


@pulumi.input_type
class LogstashOperationDurationArgs:
    def __init__(__self__, *,
                 periods: pulumi.Input[Sequence[pulumi.Input[int]]],
                 time_end: pulumi.Input[str],
                 time_start: pulumi.Input[str],
                 time_zone: pulumi.Input[str]):
        pulumi.set(__self__, "periods", periods)
        pulumi.set(__self__, "time_end", time_end)
        pulumi.set(__self__, "time_start", time_start)
        pulumi.set(__self__, "time_zone", time_zone)

    @property
    @pulumi.getter
    def periods(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        return pulumi.get(self, "periods")

    @periods.setter
    def periods(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "periods", value)

    @property
    @pulumi.getter(name="timeEnd")
    def time_end(self) -> pulumi.Input[str]:
        return pulumi.get(self, "time_end")

    @time_end.setter
    def time_end(self, value: pulumi.Input[str]):
        pulumi.set(self, "time_end", value)

    @property
    @pulumi.getter(name="timeStart")
    def time_start(self) -> pulumi.Input[str]:
        return pulumi.get(self, "time_start")

    @time_start.setter
    def time_start(self, value: pulumi.Input[str]):
        pulumi.set(self, "time_start", value)

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> pulumi.Input[str]:
        return pulumi.get(self, "time_zone")

    @time_zone.setter
    def time_zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "time_zone", value)


@pulumi.input_type
class LogstashPipelinePipelineArgs:
    def __init__(__self__, *,
                 batch_delay: pulumi.Input[int],
                 batch_size: pulumi.Input[int],
                 config: pulumi.Input[str],
                 pipeline_desc: pulumi.Input[str],
                 pipeline_id: pulumi.Input[str],
                 queue_check_point_writes: pulumi.Input[int],
                 queue_max_bytes: pulumi.Input[str],
                 queue_type: pulumi.Input[str],
                 workers: pulumi.Input[int]):
        pulumi.set(__self__, "batch_delay", batch_delay)
        pulumi.set(__self__, "batch_size", batch_size)
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "pipeline_desc", pipeline_desc)
        pulumi.set(__self__, "pipeline_id", pipeline_id)
        pulumi.set(__self__, "queue_check_point_writes", queue_check_point_writes)
        pulumi.set(__self__, "queue_max_bytes", queue_max_bytes)
        pulumi.set(__self__, "queue_type", queue_type)
        pulumi.set(__self__, "workers", workers)

    @property
    @pulumi.getter(name="batchDelay")
    def batch_delay(self) -> pulumi.Input[int]:
        return pulumi.get(self, "batch_delay")

    @batch_delay.setter
    def batch_delay(self, value: pulumi.Input[int]):
        pulumi.set(self, "batch_delay", value)

    @property
    @pulumi.getter(name="batchSize")
    def batch_size(self) -> pulumi.Input[int]:
        return pulumi.get(self, "batch_size")

    @batch_size.setter
    def batch_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "batch_size", value)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input[str]:
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input[str]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="pipelineDesc")
    def pipeline_desc(self) -> pulumi.Input[str]:
        return pulumi.get(self, "pipeline_desc")

    @pipeline_desc.setter
    def pipeline_desc(self, value: pulumi.Input[str]):
        pulumi.set(self, "pipeline_desc", value)

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "pipeline_id")

    @pipeline_id.setter
    def pipeline_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "pipeline_id", value)

    @property
    @pulumi.getter(name="queueCheckPointWrites")
    def queue_check_point_writes(self) -> pulumi.Input[int]:
        return pulumi.get(self, "queue_check_point_writes")

    @queue_check_point_writes.setter
    def queue_check_point_writes(self, value: pulumi.Input[int]):
        pulumi.set(self, "queue_check_point_writes", value)

    @property
    @pulumi.getter(name="queueMaxBytes")
    def queue_max_bytes(self) -> pulumi.Input[str]:
        return pulumi.get(self, "queue_max_bytes")

    @queue_max_bytes.setter
    def queue_max_bytes(self, value: pulumi.Input[str]):
        pulumi.set(self, "queue_max_bytes", value)

    @property
    @pulumi.getter(name="queueType")
    def queue_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "queue_type")

    @queue_type.setter
    def queue_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "queue_type", value)

    @property
    @pulumi.getter
    def workers(self) -> pulumi.Input[int]:
        return pulumi.get(self, "workers")

    @workers.setter
    def workers(self, value: pulumi.Input[int]):
        pulumi.set(self, "workers", value)

