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
    'AuditTrackFilters',
    'AuditTrackFiltersResourceField',
    'AuditTrackStorage',
]

@pulumi.output_type
class AuditTrackFilters(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "resourceFields":
            suggest = "resource_fields"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AuditTrackFilters. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AuditTrackFilters.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AuditTrackFilters.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 resource_fields: Optional[Sequence['outputs.AuditTrackFiltersResourceField']] = None):
        if resource_fields is not None:
            pulumi.set(__self__, "resource_fields", resource_fields)

    @property
    @pulumi.getter(name="resourceFields")
    def resource_fields(self) -> Optional[Sequence['outputs.AuditTrackFiltersResourceField']]:
        return pulumi.get(self, "resource_fields")


@pulumi.output_type
class AuditTrackFiltersResourceField(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "actionType":
            suggest = "action_type"
        elif key == "eventNames":
            suggest = "event_names"
        elif key == "resourceType":
            suggest = "resource_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AuditTrackFiltersResourceField. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AuditTrackFiltersResourceField.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AuditTrackFiltersResourceField.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action_type: str,
                 event_names: Sequence[str],
                 resource_type: str):
        pulumi.set(__self__, "action_type", action_type)
        pulumi.set(__self__, "event_names", event_names)
        pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="actionType")
    def action_type(self) -> str:
        return pulumi.get(self, "action_type")

    @property
    @pulumi.getter(name="eventNames")
    def event_names(self) -> Sequence[str]:
        return pulumi.get(self, "event_names")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        return pulumi.get(self, "resource_type")


@pulumi.output_type
class AuditTrackStorage(dict):
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
            pulumi.log.warn(f"Key '{key}' not found in AuditTrackStorage. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AuditTrackStorage.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AuditTrackStorage.__key_warning(key)
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

