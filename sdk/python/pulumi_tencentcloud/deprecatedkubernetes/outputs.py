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
    'AsScalingGroupAutoScalingConfig',
    'AsScalingGroupAutoScalingConfigDataDisk',
    'AsScalingGroupAutoScalingGroup',
    'AsScalingGroupAutoScalingGroupForwardBalancerId',
    'AsScalingGroupAutoScalingGroupForwardBalancerIdTargetAttribute',
]

@pulumi.output_type
class AsScalingGroupAutoScalingConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "configurationName":
            suggest = "configuration_name"
        elif key == "instanceType":
            suggest = "instance_type"
        elif key == "dataDisks":
            suggest = "data_disks"
        elif key == "enhancedMonitorService":
            suggest = "enhanced_monitor_service"
        elif key == "enhancedSecurityService":
            suggest = "enhanced_security_service"
        elif key == "instanceTags":
            suggest = "instance_tags"
        elif key == "internetChargeType":
            suggest = "internet_charge_type"
        elif key == "internetMaxBandwidthOut":
            suggest = "internet_max_bandwidth_out"
        elif key == "keyIds":
            suggest = "key_ids"
        elif key == "projectId":
            suggest = "project_id"
        elif key == "publicIpAssigned":
            suggest = "public_ip_assigned"
        elif key == "securityGroupIds":
            suggest = "security_group_ids"
        elif key == "systemDiskSize":
            suggest = "system_disk_size"
        elif key == "systemDiskType":
            suggest = "system_disk_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AsScalingGroupAutoScalingConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AsScalingGroupAutoScalingConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AsScalingGroupAutoScalingConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 configuration_name: str,
                 instance_type: str,
                 data_disks: Optional[Sequence['outputs.AsScalingGroupAutoScalingConfigDataDisk']] = None,
                 enhanced_monitor_service: Optional[bool] = None,
                 enhanced_security_service: Optional[bool] = None,
                 instance_tags: Optional[Mapping[str, Any]] = None,
                 internet_charge_type: Optional[str] = None,
                 internet_max_bandwidth_out: Optional[int] = None,
                 key_ids: Optional[Sequence[str]] = None,
                 password: Optional[str] = None,
                 project_id: Optional[int] = None,
                 public_ip_assigned: Optional[bool] = None,
                 security_group_ids: Optional[Sequence[str]] = None,
                 system_disk_size: Optional[int] = None,
                 system_disk_type: Optional[str] = None):
        pulumi.set(__self__, "configuration_name", configuration_name)
        pulumi.set(__self__, "instance_type", instance_type)
        if data_disks is not None:
            pulumi.set(__self__, "data_disks", data_disks)
        if enhanced_monitor_service is not None:
            pulumi.set(__self__, "enhanced_monitor_service", enhanced_monitor_service)
        if enhanced_security_service is not None:
            pulumi.set(__self__, "enhanced_security_service", enhanced_security_service)
        if instance_tags is not None:
            pulumi.set(__self__, "instance_tags", instance_tags)
        if internet_charge_type is not None:
            pulumi.set(__self__, "internet_charge_type", internet_charge_type)
        if internet_max_bandwidth_out is not None:
            pulumi.set(__self__, "internet_max_bandwidth_out", internet_max_bandwidth_out)
        if key_ids is not None:
            pulumi.set(__self__, "key_ids", key_ids)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if public_ip_assigned is not None:
            pulumi.set(__self__, "public_ip_assigned", public_ip_assigned)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if system_disk_size is not None:
            pulumi.set(__self__, "system_disk_size", system_disk_size)
        if system_disk_type is not None:
            pulumi.set(__self__, "system_disk_type", system_disk_type)

    @property
    @pulumi.getter(name="configurationName")
    def configuration_name(self) -> str:
        return pulumi.get(self, "configuration_name")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="dataDisks")
    def data_disks(self) -> Optional[Sequence['outputs.AsScalingGroupAutoScalingConfigDataDisk']]:
        return pulumi.get(self, "data_disks")

    @property
    @pulumi.getter(name="enhancedMonitorService")
    def enhanced_monitor_service(self) -> Optional[bool]:
        return pulumi.get(self, "enhanced_monitor_service")

    @property
    @pulumi.getter(name="enhancedSecurityService")
    def enhanced_security_service(self) -> Optional[bool]:
        return pulumi.get(self, "enhanced_security_service")

    @property
    @pulumi.getter(name="instanceTags")
    def instance_tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "instance_tags")

    @property
    @pulumi.getter(name="internetChargeType")
    def internet_charge_type(self) -> Optional[str]:
        return pulumi.get(self, "internet_charge_type")

    @property
    @pulumi.getter(name="internetMaxBandwidthOut")
    def internet_max_bandwidth_out(self) -> Optional[int]:
        return pulumi.get(self, "internet_max_bandwidth_out")

    @property
    @pulumi.getter(name="keyIds")
    def key_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "key_ids")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[int]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="publicIpAssigned")
    def public_ip_assigned(self) -> Optional[bool]:
        return pulumi.get(self, "public_ip_assigned")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="systemDiskSize")
    def system_disk_size(self) -> Optional[int]:
        return pulumi.get(self, "system_disk_size")

    @property
    @pulumi.getter(name="systemDiskType")
    def system_disk_type(self) -> Optional[str]:
        return pulumi.get(self, "system_disk_type")


@pulumi.output_type
class AsScalingGroupAutoScalingConfigDataDisk(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "diskSize":
            suggest = "disk_size"
        elif key == "diskType":
            suggest = "disk_type"
        elif key == "snapshotId":
            suggest = "snapshot_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AsScalingGroupAutoScalingConfigDataDisk. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AsScalingGroupAutoScalingConfigDataDisk.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AsScalingGroupAutoScalingConfigDataDisk.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 disk_size: Optional[int] = None,
                 disk_type: Optional[str] = None,
                 snapshot_id: Optional[str] = None):
        if disk_size is not None:
            pulumi.set(__self__, "disk_size", disk_size)
        if disk_type is not None:
            pulumi.set(__self__, "disk_type", disk_type)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)

    @property
    @pulumi.getter(name="diskSize")
    def disk_size(self) -> Optional[int]:
        return pulumi.get(self, "disk_size")

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> Optional[str]:
        return pulumi.get(self, "disk_type")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[str]:
        return pulumi.get(self, "snapshot_id")


@pulumi.output_type
class AsScalingGroupAutoScalingGroup(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxSize":
            suggest = "max_size"
        elif key == "minSize":
            suggest = "min_size"
        elif key == "scalingGroupName":
            suggest = "scaling_group_name"
        elif key == "vpcId":
            suggest = "vpc_id"
        elif key == "defaultCooldown":
            suggest = "default_cooldown"
        elif key == "desiredCapacity":
            suggest = "desired_capacity"
        elif key == "forwardBalancerIds":
            suggest = "forward_balancer_ids"
        elif key == "loadBalancerIds":
            suggest = "load_balancer_ids"
        elif key == "projectId":
            suggest = "project_id"
        elif key == "retryPolicy":
            suggest = "retry_policy"
        elif key == "subnetIds":
            suggest = "subnet_ids"
        elif key == "terminationPolicies":
            suggest = "termination_policies"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AsScalingGroupAutoScalingGroup. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AsScalingGroupAutoScalingGroup.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AsScalingGroupAutoScalingGroup.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 max_size: int,
                 min_size: int,
                 scaling_group_name: str,
                 vpc_id: str,
                 default_cooldown: Optional[int] = None,
                 desired_capacity: Optional[int] = None,
                 forward_balancer_ids: Optional[Sequence['outputs.AsScalingGroupAutoScalingGroupForwardBalancerId']] = None,
                 load_balancer_ids: Optional[Sequence[str]] = None,
                 project_id: Optional[int] = None,
                 retry_policy: Optional[str] = None,
                 subnet_ids: Optional[Sequence[str]] = None,
                 tags: Optional[Mapping[str, Any]] = None,
                 termination_policies: Optional[str] = None,
                 zones: Optional[Sequence[str]] = None):
        pulumi.set(__self__, "max_size", max_size)
        pulumi.set(__self__, "min_size", min_size)
        pulumi.set(__self__, "scaling_group_name", scaling_group_name)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if default_cooldown is not None:
            pulumi.set(__self__, "default_cooldown", default_cooldown)
        if desired_capacity is not None:
            pulumi.set(__self__, "desired_capacity", desired_capacity)
        if forward_balancer_ids is not None:
            pulumi.set(__self__, "forward_balancer_ids", forward_balancer_ids)
        if load_balancer_ids is not None:
            pulumi.set(__self__, "load_balancer_ids", load_balancer_ids)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if retry_policy is not None:
            pulumi.set(__self__, "retry_policy", retry_policy)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if termination_policies is not None:
            pulumi.set(__self__, "termination_policies", termination_policies)
        if zones is not None:
            pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> int:
        return pulumi.get(self, "max_size")

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> int:
        return pulumi.get(self, "min_size")

    @property
    @pulumi.getter(name="scalingGroupName")
    def scaling_group_name(self) -> str:
        return pulumi.get(self, "scaling_group_name")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="defaultCooldown")
    def default_cooldown(self) -> Optional[int]:
        return pulumi.get(self, "default_cooldown")

    @property
    @pulumi.getter(name="desiredCapacity")
    def desired_capacity(self) -> Optional[int]:
        return pulumi.get(self, "desired_capacity")

    @property
    @pulumi.getter(name="forwardBalancerIds")
    def forward_balancer_ids(self) -> Optional[Sequence['outputs.AsScalingGroupAutoScalingGroupForwardBalancerId']]:
        return pulumi.get(self, "forward_balancer_ids")

    @property
    @pulumi.getter(name="loadBalancerIds")
    def load_balancer_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "load_balancer_ids")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[int]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="retryPolicy")
    def retry_policy(self) -> Optional[str]:
        return pulumi.get(self, "retry_policy")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="terminationPolicies")
    def termination_policies(self) -> Optional[str]:
        return pulumi.get(self, "termination_policies")

    @property
    @pulumi.getter
    def zones(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "zones")


@pulumi.output_type
class AsScalingGroupAutoScalingGroupForwardBalancerId(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "listenerId":
            suggest = "listener_id"
        elif key == "loadBalancerId":
            suggest = "load_balancer_id"
        elif key == "targetAttributes":
            suggest = "target_attributes"
        elif key == "ruleId":
            suggest = "rule_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AsScalingGroupAutoScalingGroupForwardBalancerId. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AsScalingGroupAutoScalingGroupForwardBalancerId.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AsScalingGroupAutoScalingGroupForwardBalancerId.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 listener_id: str,
                 load_balancer_id: str,
                 target_attributes: Sequence['outputs.AsScalingGroupAutoScalingGroupForwardBalancerIdTargetAttribute'],
                 rule_id: Optional[str] = None):
        pulumi.set(__self__, "listener_id", listener_id)
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        pulumi.set(__self__, "target_attributes", target_attributes)
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> str:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> str:
        return pulumi.get(self, "load_balancer_id")

    @property
    @pulumi.getter(name="targetAttributes")
    def target_attributes(self) -> Sequence['outputs.AsScalingGroupAutoScalingGroupForwardBalancerIdTargetAttribute']:
        return pulumi.get(self, "target_attributes")

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[str]:
        return pulumi.get(self, "rule_id")


@pulumi.output_type
class AsScalingGroupAutoScalingGroupForwardBalancerIdTargetAttribute(dict):
    def __init__(__self__, *,
                 port: int,
                 weight: int):
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def weight(self) -> int:
        return pulumi.get(self, "weight")

