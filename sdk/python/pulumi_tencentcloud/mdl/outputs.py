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
    'StreamLiveInputInputSetting',
]

@pulumi.output_type
class StreamLiveInputInputSetting(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "appName":
            suggest = "app_name"
        elif key == "delayTime":
            suggest = "delay_time"
        elif key == "inputAddress":
            suggest = "input_address"
        elif key == "inputDomain":
            suggest = "input_domain"
        elif key == "sourceType":
            suggest = "source_type"
        elif key == "sourceUrl":
            suggest = "source_url"
        elif key == "streamName":
            suggest = "stream_name"
        elif key == "userName":
            suggest = "user_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StreamLiveInputInputSetting. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StreamLiveInputInputSetting.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StreamLiveInputInputSetting.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 app_name: Optional[str] = None,
                 delay_time: Optional[int] = None,
                 input_address: Optional[str] = None,
                 input_domain: Optional[str] = None,
                 password: Optional[str] = None,
                 source_type: Optional[str] = None,
                 source_url: Optional[str] = None,
                 stream_name: Optional[str] = None,
                 user_name: Optional[str] = None):
        if app_name is not None:
            pulumi.set(__self__, "app_name", app_name)
        if delay_time is not None:
            pulumi.set(__self__, "delay_time", delay_time)
        if input_address is not None:
            pulumi.set(__self__, "input_address", input_address)
        if input_domain is not None:
            pulumi.set(__self__, "input_domain", input_domain)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if source_type is not None:
            pulumi.set(__self__, "source_type", source_type)
        if source_url is not None:
            pulumi.set(__self__, "source_url", source_url)
        if stream_name is not None:
            pulumi.set(__self__, "stream_name", stream_name)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> Optional[str]:
        return pulumi.get(self, "app_name")

    @property
    @pulumi.getter(name="delayTime")
    def delay_time(self) -> Optional[int]:
        return pulumi.get(self, "delay_time")

    @property
    @pulumi.getter(name="inputAddress")
    def input_address(self) -> Optional[str]:
        return pulumi.get(self, "input_address")

    @property
    @pulumi.getter(name="inputDomain")
    def input_domain(self) -> Optional[str]:
        return pulumi.get(self, "input_domain")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional[str]:
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter(name="sourceUrl")
    def source_url(self) -> Optional[str]:
        return pulumi.get(self, "source_url")

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> Optional[str]:
        return pulumi.get(self, "stream_name")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[str]:
        return pulumi.get(self, "user_name")

