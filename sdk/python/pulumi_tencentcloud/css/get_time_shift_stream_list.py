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

__all__ = [
    'GetTimeShiftStreamListResult',
    'AwaitableGetTimeShiftStreamListResult',
    'get_time_shift_stream_list',
    'get_time_shift_stream_list_output',
]

@pulumi.output_type
class GetTimeShiftStreamListResult:
    """
    A collection of values returned by getTimeShiftStreamList.
    """
    def __init__(__self__, domain=None, domain_group=None, end_time=None, id=None, result_output_file=None, start_time=None, stream_lists=None, stream_name=None, total_size=None):
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if domain_group and not isinstance(domain_group, str):
            raise TypeError("Expected argument 'domain_group' to be a str")
        pulumi.set(__self__, "domain_group", domain_group)
        if end_time and not isinstance(end_time, int):
            raise TypeError("Expected argument 'end_time' to be a int")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if start_time and not isinstance(start_time, int):
            raise TypeError("Expected argument 'start_time' to be a int")
        pulumi.set(__self__, "start_time", start_time)
        if stream_lists and not isinstance(stream_lists, list):
            raise TypeError("Expected argument 'stream_lists' to be a list")
        pulumi.set(__self__, "stream_lists", stream_lists)
        if stream_name and not isinstance(stream_name, str):
            raise TypeError("Expected argument 'stream_name' to be a str")
        pulumi.set(__self__, "stream_name", stream_name)
        if total_size and not isinstance(total_size, int):
            raise TypeError("Expected argument 'total_size' to be a int")
        pulumi.set(__self__, "total_size", total_size)

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="domainGroup")
    def domain_group(self) -> Optional[str]:
        return pulumi.get(self, "domain_group")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> int:
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> int:
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter(name="streamLists")
    def stream_lists(self) -> Sequence['outputs.GetTimeShiftStreamListStreamListResult']:
        return pulumi.get(self, "stream_lists")

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> Optional[str]:
        return pulumi.get(self, "stream_name")

    @property
    @pulumi.getter(name="totalSize")
    def total_size(self) -> int:
        return pulumi.get(self, "total_size")


class AwaitableGetTimeShiftStreamListResult(GetTimeShiftStreamListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTimeShiftStreamListResult(
            domain=self.domain,
            domain_group=self.domain_group,
            end_time=self.end_time,
            id=self.id,
            result_output_file=self.result_output_file,
            start_time=self.start_time,
            stream_lists=self.stream_lists,
            stream_name=self.stream_name,
            total_size=self.total_size)


def get_time_shift_stream_list(domain: Optional[str] = None,
                               domain_group: Optional[str] = None,
                               end_time: Optional[int] = None,
                               result_output_file: Optional[str] = None,
                               start_time: Optional[int] = None,
                               stream_name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTimeShiftStreamListResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['domainGroup'] = domain_group
    __args__['endTime'] = end_time
    __args__['resultOutputFile'] = result_output_file
    __args__['startTime'] = start_time
    __args__['streamName'] = stream_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Css/getTimeShiftStreamList:getTimeShiftStreamList', __args__, opts=opts, typ=GetTimeShiftStreamListResult).value

    return AwaitableGetTimeShiftStreamListResult(
        domain=pulumi.get(__ret__, 'domain'),
        domain_group=pulumi.get(__ret__, 'domain_group'),
        end_time=pulumi.get(__ret__, 'end_time'),
        id=pulumi.get(__ret__, 'id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        start_time=pulumi.get(__ret__, 'start_time'),
        stream_lists=pulumi.get(__ret__, 'stream_lists'),
        stream_name=pulumi.get(__ret__, 'stream_name'),
        total_size=pulumi.get(__ret__, 'total_size'))


@_utilities.lift_output_func(get_time_shift_stream_list)
def get_time_shift_stream_list_output(domain: Optional[pulumi.Input[Optional[str]]] = None,
                                      domain_group: Optional[pulumi.Input[Optional[str]]] = None,
                                      end_time: Optional[pulumi.Input[int]] = None,
                                      result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                      start_time: Optional[pulumi.Input[int]] = None,
                                      stream_name: Optional[pulumi.Input[Optional[str]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTimeShiftStreamListResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...