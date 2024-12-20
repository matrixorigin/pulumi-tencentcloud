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

__all__ = [
    'GetPrivateZoneListResult',
    'AwaitableGetPrivateZoneListResult',
    'get_private_zone_list',
    'get_private_zone_list_output',
]

@pulumi.output_type
class GetPrivateZoneListResult:
    """
    A collection of values returned by getPrivateZoneList.
    """
    def __init__(__self__, filters=None, id=None, private_zone_sets=None, result_output_file=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if private_zone_sets and not isinstance(private_zone_sets, list):
            raise TypeError("Expected argument 'private_zone_sets' to be a list")
        pulumi.set(__self__, "private_zone_sets", private_zone_sets)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetPrivateZoneListFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="privateZoneSets")
    def private_zone_sets(self) -> Sequence['outputs.GetPrivateZoneListPrivateZoneSetResult']:
        return pulumi.get(self, "private_zone_sets")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetPrivateZoneListResult(GetPrivateZoneListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivateZoneListResult(
            filters=self.filters,
            id=self.id,
            private_zone_sets=self.private_zone_sets,
            result_output_file=self.result_output_file)


def get_private_zone_list(filters: Optional[Sequence[pulumi.InputType['GetPrivateZoneListFilterArgs']]] = None,
                          result_output_file: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivateZoneListResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:PrivateDns/getPrivateZoneList:getPrivateZoneList', __args__, opts=opts, typ=GetPrivateZoneListResult).value

    return AwaitableGetPrivateZoneListResult(
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        private_zone_sets=pulumi.get(__ret__, 'private_zone_sets'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_private_zone_list)
def get_private_zone_list_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetPrivateZoneListFilterArgs']]]]] = None,
                                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPrivateZoneListResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
