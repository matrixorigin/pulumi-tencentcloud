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
    'GetPluginsResult',
    'AwaitableGetPluginsResult',
    'get_plugins',
    'get_plugins_output',
]

@pulumi.output_type
class GetPluginsResult:
    """
    A collection of values returned by getPlugins.
    """
    def __init__(__self__, environment_name=None, id=None, plugin_id=None, result_output_file=None, results=None, service_id=None):
        if environment_name and not isinstance(environment_name, str):
            raise TypeError("Expected argument 'environment_name' to be a str")
        pulumi.set(__self__, "environment_name", environment_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if plugin_id and not isinstance(plugin_id, str):
            raise TypeError("Expected argument 'plugin_id' to be a str")
        pulumi.set(__self__, "plugin_id", plugin_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)
        if service_id and not isinstance(service_id, str):
            raise TypeError("Expected argument 'service_id' to be a str")
        pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> str:
        return pulumi.get(self, "environment_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="pluginId")
    def plugin_id(self) -> str:
        return pulumi.get(self, "plugin_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetPluginsResultResult']:
        return pulumi.get(self, "results")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> str:
        return pulumi.get(self, "service_id")


class AwaitableGetPluginsResult(GetPluginsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPluginsResult(
            environment_name=self.environment_name,
            id=self.id,
            plugin_id=self.plugin_id,
            result_output_file=self.result_output_file,
            results=self.results,
            service_id=self.service_id)


def get_plugins(environment_name: Optional[str] = None,
                plugin_id: Optional[str] = None,
                result_output_file: Optional[str] = None,
                service_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPluginsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['environmentName'] = environment_name
    __args__['pluginId'] = plugin_id
    __args__['resultOutputFile'] = result_output_file
    __args__['serviceId'] = service_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:ApiGateway/getPlugins:getPlugins', __args__, opts=opts, typ=GetPluginsResult).value

    return AwaitableGetPluginsResult(
        environment_name=pulumi.get(__ret__, 'environment_name'),
        id=pulumi.get(__ret__, 'id'),
        plugin_id=pulumi.get(__ret__, 'plugin_id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        results=pulumi.get(__ret__, 'results'),
        service_id=pulumi.get(__ret__, 'service_id'))


@_utilities.lift_output_func(get_plugins)
def get_plugins_output(environment_name: Optional[pulumi.Input[str]] = None,
                       plugin_id: Optional[pulumi.Input[str]] = None,
                       result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       service_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPluginsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
