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
    'GetSceneResult',
    'AwaitableGetSceneResult',
    'get_scene',
    'get_scene_output',
]

@pulumi.output_type
class GetSceneResult:
    """
    A collection of values returned by getScene.
    """
    def __init__(__self__, id=None, limit=None, offset=None, result_output_file=None, scene_ids=None, scene_sets=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if limit and not isinstance(limit, int):
            raise TypeError("Expected argument 'limit' to be a int")
        pulumi.set(__self__, "limit", limit)
        if offset and not isinstance(offset, int):
            raise TypeError("Expected argument 'offset' to be a int")
        pulumi.set(__self__, "offset", offset)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if scene_ids and not isinstance(scene_ids, list):
            raise TypeError("Expected argument 'scene_ids' to be a list")
        pulumi.set(__self__, "scene_ids", scene_ids)
        if scene_sets and not isinstance(scene_sets, list):
            raise TypeError("Expected argument 'scene_sets' to be a list")
        pulumi.set(__self__, "scene_sets", scene_sets)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def limit(self) -> Optional[int]:
        return pulumi.get(self, "limit")

    @property
    @pulumi.getter
    def offset(self) -> Optional[int]:
        return pulumi.get(self, "offset")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="sceneIds")
    def scene_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "scene_ids")

    @property
    @pulumi.getter(name="sceneSets")
    def scene_sets(self) -> Sequence['outputs.GetSceneSceneSetResult']:
        return pulumi.get(self, "scene_sets")


class AwaitableGetSceneResult(GetSceneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSceneResult(
            id=self.id,
            limit=self.limit,
            offset=self.offset,
            result_output_file=self.result_output_file,
            scene_ids=self.scene_ids,
            scene_sets=self.scene_sets)


def get_scene(limit: Optional[int] = None,
              offset: Optional[int] = None,
              result_output_file: Optional[str] = None,
              scene_ids: Optional[Sequence[str]] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSceneResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['limit'] = limit
    __args__['offset'] = offset
    __args__['resultOutputFile'] = result_output_file
    __args__['sceneIds'] = scene_ids
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Lighthouse/getScene:getScene', __args__, opts=opts, typ=GetSceneResult).value

    return AwaitableGetSceneResult(
        id=pulumi.get(__ret__, 'id'),
        limit=pulumi.get(__ret__, 'limit'),
        offset=pulumi.get(__ret__, 'offset'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        scene_ids=pulumi.get(__ret__, 'scene_ids'),
        scene_sets=pulumi.get(__ret__, 'scene_sets'))


@_utilities.lift_output_func(get_scene)
def get_scene_output(limit: Optional[pulumi.Input[Optional[int]]] = None,
                     offset: Optional[pulumi.Input[Optional[int]]] = None,
                     result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     scene_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSceneResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...