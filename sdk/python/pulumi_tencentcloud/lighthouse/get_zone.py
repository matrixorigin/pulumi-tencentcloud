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
    'GetZoneResult',
    'AwaitableGetZoneResult',
    'get_zone',
    'get_zone_output',
]

@pulumi.output_type
class GetZoneResult:
    """
    A collection of values returned by getZone.
    """
    def __init__(__self__, id=None, order=None, order_field=None, result_output_file=None, zone_info_sets=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if order and not isinstance(order, str):
            raise TypeError("Expected argument 'order' to be a str")
        pulumi.set(__self__, "order", order)
        if order_field and not isinstance(order_field, str):
            raise TypeError("Expected argument 'order_field' to be a str")
        pulumi.set(__self__, "order_field", order_field)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if zone_info_sets and not isinstance(zone_info_sets, list):
            raise TypeError("Expected argument 'zone_info_sets' to be a list")
        pulumi.set(__self__, "zone_info_sets", zone_info_sets)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def order(self) -> Optional[str]:
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="orderField")
    def order_field(self) -> Optional[str]:
        return pulumi.get(self, "order_field")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="zoneInfoSets")
    def zone_info_sets(self) -> Sequence['outputs.GetZoneZoneInfoSetResult']:
        return pulumi.get(self, "zone_info_sets")


class AwaitableGetZoneResult(GetZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZoneResult(
            id=self.id,
            order=self.order,
            order_field=self.order_field,
            result_output_file=self.result_output_file,
            zone_info_sets=self.zone_info_sets)


def get_zone(order: Optional[str] = None,
             order_field: Optional[str] = None,
             result_output_file: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZoneResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['order'] = order
    __args__['orderField'] = order_field
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Lighthouse/getZone:getZone', __args__, opts=opts, typ=GetZoneResult).value

    return AwaitableGetZoneResult(
        id=pulumi.get(__ret__, 'id'),
        order=pulumi.get(__ret__, 'order'),
        order_field=pulumi.get(__ret__, 'order_field'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        zone_info_sets=pulumi.get(__ret__, 'zone_info_sets'))


@_utilities.lift_output_func(get_zone)
def get_zone_output(order: Optional[pulumi.Input[Optional[str]]] = None,
                    order_field: Optional[pulumi.Input[Optional[str]]] = None,
                    result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetZoneResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...