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

__all__ = ['ParamTemplateArgs', 'ParamTemplate']

@pulumi.input_type
class ParamTemplateArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 param_lists: Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamListArgs']]]] = None,
                 template_id: Optional[pulumi.Input[int]] = None,
                 template_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ParamTemplate resource.
        :param pulumi.Input[str] description: The description of parameter template.
        :param pulumi.Input[str] engine_type: The engine type of instance, optional value is InnoDB or RocksDB, default to InnoDB.
        :param pulumi.Input[str] engine_version: The version of MySQL.
        :param pulumi.Input[str] name: The name of parameter template.
        :param pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamListArgs']]] param_lists: parameter list.
        :param pulumi.Input[int] template_id: The ID of source parameter template.
        :param pulumi.Input[str] template_type: The default type of parameter template, supported value is HIGH_STABILITY or HIGH_PERFORMANCE.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if engine_type is not None:
            pulumi.set(__self__, "engine_type", engine_type)
        if engine_version is not None:
            pulumi.set(__self__, "engine_version", engine_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if param_lists is not None:
            pulumi.set(__self__, "param_lists", param_lists)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)
        if template_type is not None:
            pulumi.set(__self__, "template_type", template_type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of parameter template.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> Optional[pulumi.Input[str]]:
        """
        The engine type of instance, optional value is InnoDB or RocksDB, default to InnoDB.
        """
        return pulumi.get(self, "engine_type")

    @engine_type.setter
    def engine_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_type", value)

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of MySQL.
        """
        return pulumi.get(self, "engine_version")

    @engine_version.setter
    def engine_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of parameter template.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="paramLists")
    def param_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamListArgs']]]]:
        """
        parameter list.
        """
        return pulumi.get(self, "param_lists")

    @param_lists.setter
    def param_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamListArgs']]]]):
        pulumi.set(self, "param_lists", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of source parameter template.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "template_id", value)

    @property
    @pulumi.getter(name="templateType")
    def template_type(self) -> Optional[pulumi.Input[str]]:
        """
        The default type of parameter template, supported value is HIGH_STABILITY or HIGH_PERFORMANCE.
        """
        return pulumi.get(self, "template_type")

    @template_type.setter
    def template_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_type", value)


@pulumi.input_type
class _ParamTemplateState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 param_lists: Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamListArgs']]]] = None,
                 template_id: Optional[pulumi.Input[int]] = None,
                 template_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ParamTemplate resources.
        :param pulumi.Input[str] description: The description of parameter template.
        :param pulumi.Input[str] engine_type: The engine type of instance, optional value is InnoDB or RocksDB, default to InnoDB.
        :param pulumi.Input[str] engine_version: The version of MySQL.
        :param pulumi.Input[str] name: The name of parameter template.
        :param pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamListArgs']]] param_lists: parameter list.
        :param pulumi.Input[int] template_id: The ID of source parameter template.
        :param pulumi.Input[str] template_type: The default type of parameter template, supported value is HIGH_STABILITY or HIGH_PERFORMANCE.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if engine_type is not None:
            pulumi.set(__self__, "engine_type", engine_type)
        if engine_version is not None:
            pulumi.set(__self__, "engine_version", engine_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if param_lists is not None:
            pulumi.set(__self__, "param_lists", param_lists)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)
        if template_type is not None:
            pulumi.set(__self__, "template_type", template_type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of parameter template.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> Optional[pulumi.Input[str]]:
        """
        The engine type of instance, optional value is InnoDB or RocksDB, default to InnoDB.
        """
        return pulumi.get(self, "engine_type")

    @engine_type.setter
    def engine_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_type", value)

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of MySQL.
        """
        return pulumi.get(self, "engine_version")

    @engine_version.setter
    def engine_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of parameter template.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="paramLists")
    def param_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamListArgs']]]]:
        """
        parameter list.
        """
        return pulumi.get(self, "param_lists")

    @param_lists.setter
    def param_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamListArgs']]]]):
        pulumi.set(self, "param_lists", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of source parameter template.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "template_id", value)

    @property
    @pulumi.getter(name="templateType")
    def template_type(self) -> Optional[pulumi.Input[str]]:
        """
        The default type of parameter template, supported value is HIGH_STABILITY or HIGH_PERFORMANCE.
        """
        return pulumi.get(self, "template_type")

    @template_type.setter
    def template_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_type", value)


class ParamTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 param_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParamTemplateParamListArgs']]]]] = None,
                 template_id: Optional[pulumi.Input[int]] = None,
                 template_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ParamTemplate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of parameter template.
        :param pulumi.Input[str] engine_type: The engine type of instance, optional value is InnoDB or RocksDB, default to InnoDB.
        :param pulumi.Input[str] engine_version: The version of MySQL.
        :param pulumi.Input[str] name: The name of parameter template.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParamTemplateParamListArgs']]]] param_lists: parameter list.
        :param pulumi.Input[int] template_id: The ID of source parameter template.
        :param pulumi.Input[str] template_type: The default type of parameter template, supported value is HIGH_STABILITY or HIGH_PERFORMANCE.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ParamTemplateArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ParamTemplate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ParamTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ParamTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 param_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParamTemplateParamListArgs']]]]] = None,
                 template_id: Optional[pulumi.Input[int]] = None,
                 template_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ParamTemplateArgs.__new__(ParamTemplateArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["engine_type"] = engine_type
            __props__.__dict__["engine_version"] = engine_version
            __props__.__dict__["name"] = name
            __props__.__dict__["param_lists"] = param_lists
            __props__.__dict__["template_id"] = template_id
            __props__.__dict__["template_type"] = template_type
        super(ParamTemplate, __self__).__init__(
            'tencentcloud:Mysql/paramTemplate:ParamTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            engine_type: Optional[pulumi.Input[str]] = None,
            engine_version: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            param_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParamTemplateParamListArgs']]]]] = None,
            template_id: Optional[pulumi.Input[int]] = None,
            template_type: Optional[pulumi.Input[str]] = None) -> 'ParamTemplate':
        """
        Get an existing ParamTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of parameter template.
        :param pulumi.Input[str] engine_type: The engine type of instance, optional value is InnoDB or RocksDB, default to InnoDB.
        :param pulumi.Input[str] engine_version: The version of MySQL.
        :param pulumi.Input[str] name: The name of parameter template.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParamTemplateParamListArgs']]]] param_lists: parameter list.
        :param pulumi.Input[int] template_id: The ID of source parameter template.
        :param pulumi.Input[str] template_type: The default type of parameter template, supported value is HIGH_STABILITY or HIGH_PERFORMANCE.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ParamTemplateState.__new__(_ParamTemplateState)

        __props__.__dict__["description"] = description
        __props__.__dict__["engine_type"] = engine_type
        __props__.__dict__["engine_version"] = engine_version
        __props__.__dict__["name"] = name
        __props__.__dict__["param_lists"] = param_lists
        __props__.__dict__["template_id"] = template_id
        __props__.__dict__["template_type"] = template_type
        return ParamTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of parameter template.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> pulumi.Output[Optional[str]]:
        """
        The engine type of instance, optional value is InnoDB or RocksDB, default to InnoDB.
        """
        return pulumi.get(self, "engine_type")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> pulumi.Output[Optional[str]]:
        """
        The version of MySQL.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of parameter template.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="paramLists")
    def param_lists(self) -> pulumi.Output[Sequence['outputs.ParamTemplateParamList']]:
        """
        parameter list.
        """
        return pulumi.get(self, "param_lists")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[int]:
        """
        The ID of source parameter template.
        """
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter(name="templateType")
    def template_type(self) -> pulumi.Output[Optional[str]]:
        """
        The default type of parameter template, supported value is HIGH_STABILITY or HIGH_PERFORMANCE.
        """
        return pulumi.get(self, "template_type")
