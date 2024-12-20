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
    'TrackStorage',
    'GetCosRegionsAuditCosRegionListResult',
    'GetEventsEventResult',
    'GetEventsEventResourcesResult',
    'GetEventsLookupAttributeResult',
    'GetKeyAliasAuditKeyAliasListResult',
]

@pulumi.output_type
class TrackStorage(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "storageName":
            suggest = "storage_name"
        elif key == "storagePrefix":
            suggest = "storage_prefix"
        elif key == "storageRegion":
            suggest = "storage_region"
        elif key == "storageType":
            suggest = "storage_type"
        elif key == "storageAccountId":
            suggest = "storage_account_id"
        elif key == "storageAppId":
            suggest = "storage_app_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TrackStorage. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TrackStorage.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TrackStorage.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 storage_name: str,
                 storage_prefix: str,
                 storage_region: str,
                 storage_type: str,
                 storage_account_id: Optional[str] = None,
                 storage_app_id: Optional[str] = None):
        pulumi.set(__self__, "storage_name", storage_name)
        pulumi.set(__self__, "storage_prefix", storage_prefix)
        pulumi.set(__self__, "storage_region", storage_region)
        pulumi.set(__self__, "storage_type", storage_type)
        if storage_account_id is not None:
            pulumi.set(__self__, "storage_account_id", storage_account_id)
        if storage_app_id is not None:
            pulumi.set(__self__, "storage_app_id", storage_app_id)

    @property
    @pulumi.getter(name="storageName")
    def storage_name(self) -> str:
        return pulumi.get(self, "storage_name")

    @property
    @pulumi.getter(name="storagePrefix")
    def storage_prefix(self) -> str:
        return pulumi.get(self, "storage_prefix")

    @property
    @pulumi.getter(name="storageRegion")
    def storage_region(self) -> str:
        return pulumi.get(self, "storage_region")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> str:
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> Optional[str]:
        return pulumi.get(self, "storage_account_id")

    @property
    @pulumi.getter(name="storageAppId")
    def storage_app_id(self) -> Optional[str]:
        return pulumi.get(self, "storage_app_id")


@pulumi.output_type
class GetCosRegionsAuditCosRegionListResult(dict):
    def __init__(__self__, *,
                 cos_region: str,
                 cos_region_name: str):
        pulumi.set(__self__, "cos_region", cos_region)
        pulumi.set(__self__, "cos_region_name", cos_region_name)

    @property
    @pulumi.getter(name="cosRegion")
    def cos_region(self) -> str:
        return pulumi.get(self, "cos_region")

    @property
    @pulumi.getter(name="cosRegionName")
    def cos_region_name(self) -> str:
        return pulumi.get(self, "cos_region_name")


@pulumi.output_type
class GetEventsEventResult(dict):
    def __init__(__self__, *,
                 account_id: Optional[int] = None,
                 cloud_audit_event: Optional[str] = None,
                 error_code: Optional[int] = None,
                 event_id: Optional[str] = None,
                 event_name: Optional[str] = None,
                 event_name_cn: Optional[str] = None,
                 event_region: Optional[str] = None,
                 event_source: Optional[str] = None,
                 event_time: Optional[str] = None,
                 location: Optional[str] = None,
                 request_id: Optional[str] = None,
                 resource_region: Optional[str] = None,
                 resource_type_cn: Optional[str] = None,
                 resources: Optional['outputs.GetEventsEventResourcesResult'] = None,
                 secret_id: Optional[str] = None,
                 source_ip_address: Optional[str] = None,
                 username: Optional[str] = None):
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if cloud_audit_event is not None:
            pulumi.set(__self__, "cloud_audit_event", cloud_audit_event)
        if error_code is not None:
            pulumi.set(__self__, "error_code", error_code)
        if event_id is not None:
            pulumi.set(__self__, "event_id", event_id)
        if event_name is not None:
            pulumi.set(__self__, "event_name", event_name)
        if event_name_cn is not None:
            pulumi.set(__self__, "event_name_cn", event_name_cn)
        if event_region is not None:
            pulumi.set(__self__, "event_region", event_region)
        if event_source is not None:
            pulumi.set(__self__, "event_source", event_source)
        if event_time is not None:
            pulumi.set(__self__, "event_time", event_time)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if resource_region is not None:
            pulumi.set(__self__, "resource_region", resource_region)
        if resource_type_cn is not None:
            pulumi.set(__self__, "resource_type_cn", resource_type_cn)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if secret_id is not None:
            pulumi.set(__self__, "secret_id", secret_id)
        if source_ip_address is not None:
            pulumi.set(__self__, "source_ip_address", source_ip_address)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[int]:
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="cloudAuditEvent")
    def cloud_audit_event(self) -> Optional[str]:
        return pulumi.get(self, "cloud_audit_event")

    @property
    @pulumi.getter(name="errorCode")
    def error_code(self) -> Optional[int]:
        return pulumi.get(self, "error_code")

    @property
    @pulumi.getter(name="eventId")
    def event_id(self) -> Optional[str]:
        return pulumi.get(self, "event_id")

    @property
    @pulumi.getter(name="eventName")
    def event_name(self) -> Optional[str]:
        return pulumi.get(self, "event_name")

    @property
    @pulumi.getter(name="eventNameCn")
    def event_name_cn(self) -> Optional[str]:
        return pulumi.get(self, "event_name_cn")

    @property
    @pulumi.getter(name="eventRegion")
    def event_region(self) -> Optional[str]:
        return pulumi.get(self, "event_region")

    @property
    @pulumi.getter(name="eventSource")
    def event_source(self) -> Optional[str]:
        return pulumi.get(self, "event_source")

    @property
    @pulumi.getter(name="eventTime")
    def event_time(self) -> Optional[str]:
        return pulumi.get(self, "event_time")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[str]:
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter(name="resourceRegion")
    def resource_region(self) -> Optional[str]:
        return pulumi.get(self, "resource_region")

    @property
    @pulumi.getter(name="resourceTypeCn")
    def resource_type_cn(self) -> Optional[str]:
        return pulumi.get(self, "resource_type_cn")

    @property
    @pulumi.getter
    def resources(self) -> Optional['outputs.GetEventsEventResourcesResult']:
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[str]:
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter(name="sourceIpAddress")
    def source_ip_address(self) -> Optional[str]:
        return pulumi.get(self, "source_ip_address")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        return pulumi.get(self, "username")


@pulumi.output_type
class GetEventsEventResourcesResult(dict):
    def __init__(__self__, *,
                 resource_name: Optional[str] = None,
                 resource_type: Optional[str] = None):
        if resource_name is not None:
            pulumi.set(__self__, "resource_name", resource_name)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> Optional[str]:
        return pulumi.get(self, "resource_name")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[str]:
        return pulumi.get(self, "resource_type")


@pulumi.output_type
class GetEventsLookupAttributeResult(dict):
    def __init__(__self__, *,
                 attribute_key: str,
                 attribute_value: Optional[str] = None):
        pulumi.set(__self__, "attribute_key", attribute_key)
        if attribute_value is not None:
            pulumi.set(__self__, "attribute_value", attribute_value)

    @property
    @pulumi.getter(name="attributeKey")
    def attribute_key(self) -> str:
        return pulumi.get(self, "attribute_key")

    @property
    @pulumi.getter(name="attributeValue")
    def attribute_value(self) -> Optional[str]:
        return pulumi.get(self, "attribute_value")


@pulumi.output_type
class GetKeyAliasAuditKeyAliasListResult(dict):
    def __init__(__self__, *,
                 key_alias: str,
                 key_id: str):
        pulumi.set(__self__, "key_alias", key_alias)
        pulumi.set(__self__, "key_id", key_id)

    @property
    @pulumi.getter(name="keyAlias")
    def key_alias(self) -> str:
        return pulumi.get(self, "key_alias")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> str:
        return pulumi.get(self, "key_id")


