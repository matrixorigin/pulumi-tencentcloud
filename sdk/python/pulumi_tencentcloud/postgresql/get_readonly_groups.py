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
    'GetReadonlyGroupsResult',
    'AwaitableGetReadonlyGroupsResult',
    'get_readonly_groups',
    'get_readonly_groups_output',
]

@pulumi.output_type
class GetReadonlyGroupsResult:
    """
    A collection of values returned by getReadonlyGroups.
    """
    def __init__(__self__, filters=None, id=None, order_by=None, order_by_type=None, read_only_group_lists=None, result_output_file=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        pulumi.set(__self__, "order_by", order_by)
        if order_by_type and not isinstance(order_by_type, str):
            raise TypeError("Expected argument 'order_by_type' to be a str")
        pulumi.set(__self__, "order_by_type", order_by_type)
        if read_only_group_lists and not isinstance(read_only_group_lists, list):
            raise TypeError("Expected argument 'read_only_group_lists' to be a list")
        pulumi.set(__self__, "read_only_group_lists", read_only_group_lists)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetReadonlyGroupsFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[str]:
        return pulumi.get(self, "order_by")

    @property
    @pulumi.getter(name="orderByType")
    def order_by_type(self) -> Optional[str]:
        return pulumi.get(self, "order_by_type")

    @property
    @pulumi.getter(name="readOnlyGroupLists")
    def read_only_group_lists(self) -> Sequence['outputs.GetReadonlyGroupsReadOnlyGroupListResult']:
        return pulumi.get(self, "read_only_group_lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetReadonlyGroupsResult(GetReadonlyGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReadonlyGroupsResult(
            filters=self.filters,
            id=self.id,
            order_by=self.order_by,
            order_by_type=self.order_by_type,
            read_only_group_lists=self.read_only_group_lists,
            result_output_file=self.result_output_file)


def get_readonly_groups(filters: Optional[Sequence[pulumi.InputType['GetReadonlyGroupsFilterArgs']]] = None,
                        order_by: Optional[str] = None,
                        order_by_type: Optional[str] = None,
                        result_output_file: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReadonlyGroupsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['orderBy'] = order_by
    __args__['orderByType'] = order_by_type
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Postgresql/getReadonlyGroups:getReadonlyGroups', __args__, opts=opts, typ=GetReadonlyGroupsResult).value

    return AwaitableGetReadonlyGroupsResult(
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        order_by=pulumi.get(__ret__, 'order_by'),
        order_by_type=pulumi.get(__ret__, 'order_by_type'),
        read_only_group_lists=pulumi.get(__ret__, 'read_only_group_lists'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_readonly_groups)
def get_readonly_groups_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetReadonlyGroupsFilterArgs']]]]] = None,
                               order_by: Optional[pulumi.Input[Optional[str]]] = None,
                               order_by_type: Optional[pulumi.Input[Optional[str]]] = None,
                               result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetReadonlyGroupsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...