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
    'AccessLogConfigClsArgs',
    'AccessLogConfigSelectedRangeArgs',
    'AccessLogConfigSelectedRangeItemArgs',
    'ClusterAttachmentClusterListArgs',
    'MeshConfigArgs',
    'MeshConfigInjectArgs',
    'MeshConfigIstioArgs',
    'MeshConfigIstioSmartDnsArgs',
    'MeshConfigIstioTracingArgs',
    'MeshConfigIstioTracingApmArgs',
    'MeshConfigIstioTracingZipkinArgs',
    'MeshConfigPrometheusArgs',
    'MeshConfigPrometheusCustomPromArgs',
    'MeshConfigSidecarResourcesArgs',
    'MeshConfigSidecarResourcesLimitArgs',
    'MeshConfigSidecarResourcesRequestArgs',
    'MeshConfigTracingArgs',
    'MeshConfigTracingApmArgs',
    'MeshConfigTracingZipkinArgs',
    'MeshTagListArgs',
    'PrometheusAttachmentPrometheusArgs',
    'PrometheusAttachmentPrometheusCustomPromArgs',
    'TracingConfigApmArgs',
    'TracingConfigZipkinArgs',
]

@pulumi.input_type
class AccessLogConfigClsArgs:
    def __init__(__self__, *,
                 enable: pulumi.Input[bool],
                 log_set: Optional[pulumi.Input[str]] = None,
                 topic: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "enable", enable)
        if log_set is not None:
            pulumi.set(__self__, "log_set", log_set)
        if topic is not None:
            pulumi.set(__self__, "topic", topic)

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="logSet")
    def log_set(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "log_set")

    @log_set.setter
    def log_set(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_set", value)

    @property
    @pulumi.getter
    def topic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic", value)


@pulumi.input_type
class AccessLogConfigSelectedRangeArgs:
    def __init__(__self__, *,
                 all: Optional[pulumi.Input[bool]] = None,
                 items: Optional[pulumi.Input[Sequence[pulumi.Input['AccessLogConfigSelectedRangeItemArgs']]]] = None):
        if all is not None:
            pulumi.set(__self__, "all", all)
        if items is not None:
            pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def all(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "all")

    @all.setter
    def all(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "all", value)

    @property
    @pulumi.getter
    def items(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AccessLogConfigSelectedRangeItemArgs']]]]:
        return pulumi.get(self, "items")

    @items.setter
    def items(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AccessLogConfigSelectedRangeItemArgs']]]]):
        pulumi.set(self, "items", value)


@pulumi.input_type
class AccessLogConfigSelectedRangeItemArgs:
    def __init__(__self__, *,
                 gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        if gateways is not None:
            pulumi.set(__self__, "gateways", gateways)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def gateways(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "gateways")

    @gateways.setter
    def gateways(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "gateways", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class ClusterAttachmentClusterListArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 region: pulumi.Input[str],
                 role: pulumi.Input[str],
                 type: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 subnet_id: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class MeshConfigArgs:
    def __init__(__self__, *,
                 inject: Optional[pulumi.Input['MeshConfigInjectArgs']] = None,
                 istio: Optional[pulumi.Input['MeshConfigIstioArgs']] = None,
                 prometheus: Optional[pulumi.Input['MeshConfigPrometheusArgs']] = None,
                 sidecar_resources: Optional[pulumi.Input['MeshConfigSidecarResourcesArgs']] = None,
                 tracing: Optional[pulumi.Input['MeshConfigTracingArgs']] = None):
        if inject is not None:
            pulumi.set(__self__, "inject", inject)
        if istio is not None:
            pulumi.set(__self__, "istio", istio)
        if prometheus is not None:
            pulumi.set(__self__, "prometheus", prometheus)
        if sidecar_resources is not None:
            pulumi.set(__self__, "sidecar_resources", sidecar_resources)
        if tracing is not None:
            pulumi.set(__self__, "tracing", tracing)

    @property
    @pulumi.getter
    def inject(self) -> Optional[pulumi.Input['MeshConfigInjectArgs']]:
        return pulumi.get(self, "inject")

    @inject.setter
    def inject(self, value: Optional[pulumi.Input['MeshConfigInjectArgs']]):
        pulumi.set(self, "inject", value)

    @property
    @pulumi.getter
    def istio(self) -> Optional[pulumi.Input['MeshConfigIstioArgs']]:
        return pulumi.get(self, "istio")

    @istio.setter
    def istio(self, value: Optional[pulumi.Input['MeshConfigIstioArgs']]):
        pulumi.set(self, "istio", value)

    @property
    @pulumi.getter
    def prometheus(self) -> Optional[pulumi.Input['MeshConfigPrometheusArgs']]:
        return pulumi.get(self, "prometheus")

    @prometheus.setter
    def prometheus(self, value: Optional[pulumi.Input['MeshConfigPrometheusArgs']]):
        pulumi.set(self, "prometheus", value)

    @property
    @pulumi.getter(name="sidecarResources")
    def sidecar_resources(self) -> Optional[pulumi.Input['MeshConfigSidecarResourcesArgs']]:
        return pulumi.get(self, "sidecar_resources")

    @sidecar_resources.setter
    def sidecar_resources(self, value: Optional[pulumi.Input['MeshConfigSidecarResourcesArgs']]):
        pulumi.set(self, "sidecar_resources", value)

    @property
    @pulumi.getter
    def tracing(self) -> Optional[pulumi.Input['MeshConfigTracingArgs']]:
        return pulumi.get(self, "tracing")

    @tracing.setter
    def tracing(self, value: Optional[pulumi.Input['MeshConfigTracingArgs']]):
        pulumi.set(self, "tracing", value)


@pulumi.input_type
class MeshConfigInjectArgs:
    def __init__(__self__, *,
                 exclude_ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hold_application_until_proxy_starts: Optional[pulumi.Input[bool]] = None,
                 hold_proxy_until_application_ends: Optional[pulumi.Input[bool]] = None):
        if exclude_ip_ranges is not None:
            pulumi.set(__self__, "exclude_ip_ranges", exclude_ip_ranges)
        if hold_application_until_proxy_starts is not None:
            pulumi.set(__self__, "hold_application_until_proxy_starts", hold_application_until_proxy_starts)
        if hold_proxy_until_application_ends is not None:
            pulumi.set(__self__, "hold_proxy_until_application_ends", hold_proxy_until_application_ends)

    @property
    @pulumi.getter(name="excludeIpRanges")
    def exclude_ip_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "exclude_ip_ranges")

    @exclude_ip_ranges.setter
    def exclude_ip_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "exclude_ip_ranges", value)

    @property
    @pulumi.getter(name="holdApplicationUntilProxyStarts")
    def hold_application_until_proxy_starts(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "hold_application_until_proxy_starts")

    @hold_application_until_proxy_starts.setter
    def hold_application_until_proxy_starts(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "hold_application_until_proxy_starts", value)

    @property
    @pulumi.getter(name="holdProxyUntilApplicationEnds")
    def hold_proxy_until_application_ends(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "hold_proxy_until_application_ends")

    @hold_proxy_until_application_ends.setter
    def hold_proxy_until_application_ends(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "hold_proxy_until_application_ends", value)


@pulumi.input_type
class MeshConfigIstioArgs:
    def __init__(__self__, *,
                 outbound_traffic_policy: pulumi.Input[str],
                 disable_http_retry: Optional[pulumi.Input[bool]] = None,
                 disable_policy_checks: Optional[pulumi.Input[bool]] = None,
                 enable_pilot_http: Optional[pulumi.Input[bool]] = None,
                 smart_dns: Optional[pulumi.Input['MeshConfigIstioSmartDnsArgs']] = None,
                 tracing: Optional[pulumi.Input['MeshConfigIstioTracingArgs']] = None):
        pulumi.set(__self__, "outbound_traffic_policy", outbound_traffic_policy)
        if disable_http_retry is not None:
            pulumi.set(__self__, "disable_http_retry", disable_http_retry)
        if disable_policy_checks is not None:
            pulumi.set(__self__, "disable_policy_checks", disable_policy_checks)
        if enable_pilot_http is not None:
            pulumi.set(__self__, "enable_pilot_http", enable_pilot_http)
        if smart_dns is not None:
            pulumi.set(__self__, "smart_dns", smart_dns)
        if tracing is not None:
            pulumi.set(__self__, "tracing", tracing)

    @property
    @pulumi.getter(name="outboundTrafficPolicy")
    def outbound_traffic_policy(self) -> pulumi.Input[str]:
        return pulumi.get(self, "outbound_traffic_policy")

    @outbound_traffic_policy.setter
    def outbound_traffic_policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "outbound_traffic_policy", value)

    @property
    @pulumi.getter(name="disableHttpRetry")
    def disable_http_retry(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disable_http_retry")

    @disable_http_retry.setter
    def disable_http_retry(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_http_retry", value)

    @property
    @pulumi.getter(name="disablePolicyChecks")
    def disable_policy_checks(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disable_policy_checks")

    @disable_policy_checks.setter
    def disable_policy_checks(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_policy_checks", value)

    @property
    @pulumi.getter(name="enablePilotHttp")
    def enable_pilot_http(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_pilot_http")

    @enable_pilot_http.setter
    def enable_pilot_http(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_pilot_http", value)

    @property
    @pulumi.getter(name="smartDns")
    def smart_dns(self) -> Optional[pulumi.Input['MeshConfigIstioSmartDnsArgs']]:
        return pulumi.get(self, "smart_dns")

    @smart_dns.setter
    def smart_dns(self, value: Optional[pulumi.Input['MeshConfigIstioSmartDnsArgs']]):
        pulumi.set(self, "smart_dns", value)

    @property
    @pulumi.getter
    def tracing(self) -> Optional[pulumi.Input['MeshConfigIstioTracingArgs']]:
        return pulumi.get(self, "tracing")

    @tracing.setter
    def tracing(self, value: Optional[pulumi.Input['MeshConfigIstioTracingArgs']]):
        pulumi.set(self, "tracing", value)


@pulumi.input_type
class MeshConfigIstioSmartDnsArgs:
    def __init__(__self__, *,
                 istio_meta_dns_auto_allocate: Optional[pulumi.Input[bool]] = None,
                 istio_meta_dns_capture: Optional[pulumi.Input[bool]] = None):
        if istio_meta_dns_auto_allocate is not None:
            pulumi.set(__self__, "istio_meta_dns_auto_allocate", istio_meta_dns_auto_allocate)
        if istio_meta_dns_capture is not None:
            pulumi.set(__self__, "istio_meta_dns_capture", istio_meta_dns_capture)

    @property
    @pulumi.getter(name="istioMetaDnsAutoAllocate")
    def istio_meta_dns_auto_allocate(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "istio_meta_dns_auto_allocate")

    @istio_meta_dns_auto_allocate.setter
    def istio_meta_dns_auto_allocate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "istio_meta_dns_auto_allocate", value)

    @property
    @pulumi.getter(name="istioMetaDnsCapture")
    def istio_meta_dns_capture(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "istio_meta_dns_capture")

    @istio_meta_dns_capture.setter
    def istio_meta_dns_capture(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "istio_meta_dns_capture", value)


@pulumi.input_type
class MeshConfigIstioTracingArgs:
    def __init__(__self__, *,
                 apm: Optional[pulumi.Input['MeshConfigIstioTracingApmArgs']] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 sampling: Optional[pulumi.Input[float]] = None,
                 zipkin: Optional[pulumi.Input['MeshConfigIstioTracingZipkinArgs']] = None):
        if apm is not None:
            pulumi.set(__self__, "apm", apm)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if sampling is not None:
            pulumi.set(__self__, "sampling", sampling)
        if zipkin is not None:
            pulumi.set(__self__, "zipkin", zipkin)

    @property
    @pulumi.getter
    def apm(self) -> Optional[pulumi.Input['MeshConfigIstioTracingApmArgs']]:
        return pulumi.get(self, "apm")

    @apm.setter
    def apm(self, value: Optional[pulumi.Input['MeshConfigIstioTracingApmArgs']]):
        pulumi.set(self, "apm", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter
    def sampling(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "sampling")

    @sampling.setter
    def sampling(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "sampling", value)

    @property
    @pulumi.getter
    def zipkin(self) -> Optional[pulumi.Input['MeshConfigIstioTracingZipkinArgs']]:
        return pulumi.get(self, "zipkin")

    @zipkin.setter
    def zipkin(self, value: Optional[pulumi.Input['MeshConfigIstioTracingZipkinArgs']]):
        pulumi.set(self, "zipkin", value)


@pulumi.input_type
class MeshConfigIstioTracingApmArgs:
    def __init__(__self__, *,
                 enable: pulumi.Input[bool],
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "enable", enable)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class MeshConfigIstioTracingZipkinArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str]):
        pulumi.set(__self__, "address", address)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)


@pulumi.input_type
class MeshConfigPrometheusArgs:
    def __init__(__self__, *,
                 custom_prom: Optional[pulumi.Input['MeshConfigPrometheusCustomPromArgs']] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        if custom_prom is not None:
            pulumi.set(__self__, "custom_prom", custom_prom)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="customProm")
    def custom_prom(self) -> Optional[pulumi.Input['MeshConfigPrometheusCustomPromArgs']]:
        return pulumi.get(self, "custom_prom")

    @custom_prom.setter
    def custom_prom(self, value: Optional[pulumi.Input['MeshConfigPrometheusCustomPromArgs']]):
        pulumi.set(self, "custom_prom", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class MeshConfigPrometheusCustomPromArgs:
    def __init__(__self__, *,
                 auth_type: pulumi.Input[str],
                 url: pulumi.Input[str],
                 is_public_addr: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "auth_type", auth_type)
        pulumi.set(__self__, "url", url)
        if is_public_addr is not None:
            pulumi.set(__self__, "is_public_addr", is_public_addr)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "auth_type")

    @auth_type.setter
    def auth_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "auth_type", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="isPublicAddr")
    def is_public_addr(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "is_public_addr")

    @is_public_addr.setter
    def is_public_addr(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public_addr", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class MeshConfigSidecarResourcesArgs:
    def __init__(__self__, *,
                 limits: Optional[pulumi.Input[Sequence[pulumi.Input['MeshConfigSidecarResourcesLimitArgs']]]] = None,
                 requests: Optional[pulumi.Input[Sequence[pulumi.Input['MeshConfigSidecarResourcesRequestArgs']]]] = None):
        if limits is not None:
            pulumi.set(__self__, "limits", limits)
        if requests is not None:
            pulumi.set(__self__, "requests", requests)

    @property
    @pulumi.getter
    def limits(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MeshConfigSidecarResourcesLimitArgs']]]]:
        return pulumi.get(self, "limits")

    @limits.setter
    def limits(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MeshConfigSidecarResourcesLimitArgs']]]]):
        pulumi.set(self, "limits", value)

    @property
    @pulumi.getter
    def requests(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MeshConfigSidecarResourcesRequestArgs']]]]:
        return pulumi.get(self, "requests")

    @requests.setter
    def requests(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MeshConfigSidecarResourcesRequestArgs']]]]):
        pulumi.set(self, "requests", value)


@pulumi.input_type
class MeshConfigSidecarResourcesLimitArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 quantity: Optional[pulumi.Input[str]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)
        if quantity is not None:
            pulumi.set(__self__, "quantity", quantity)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def quantity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "quantity")

    @quantity.setter
    def quantity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "quantity", value)


@pulumi.input_type
class MeshConfigSidecarResourcesRequestArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 quantity: Optional[pulumi.Input[str]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)
        if quantity is not None:
            pulumi.set(__self__, "quantity", quantity)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def quantity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "quantity")

    @quantity.setter
    def quantity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "quantity", value)


@pulumi.input_type
class MeshConfigTracingArgs:
    def __init__(__self__, *,
                 apm: Optional[pulumi.Input['MeshConfigTracingApmArgs']] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 sampling: Optional[pulumi.Input[float]] = None,
                 zipkin: Optional[pulumi.Input['MeshConfigTracingZipkinArgs']] = None):
        if apm is not None:
            pulumi.set(__self__, "apm", apm)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if sampling is not None:
            pulumi.set(__self__, "sampling", sampling)
        if zipkin is not None:
            pulumi.set(__self__, "zipkin", zipkin)

    @property
    @pulumi.getter
    def apm(self) -> Optional[pulumi.Input['MeshConfigTracingApmArgs']]:
        return pulumi.get(self, "apm")

    @apm.setter
    def apm(self, value: Optional[pulumi.Input['MeshConfigTracingApmArgs']]):
        pulumi.set(self, "apm", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter
    def sampling(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "sampling")

    @sampling.setter
    def sampling(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "sampling", value)

    @property
    @pulumi.getter
    def zipkin(self) -> Optional[pulumi.Input['MeshConfigTracingZipkinArgs']]:
        return pulumi.get(self, "zipkin")

    @zipkin.setter
    def zipkin(self, value: Optional[pulumi.Input['MeshConfigTracingZipkinArgs']]):
        pulumi.set(self, "zipkin", value)


@pulumi.input_type
class MeshConfigTracingApmArgs:
    def __init__(__self__, *,
                 enable: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class MeshConfigTracingZipkinArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str]):
        pulumi.set(__self__, "address", address)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)


@pulumi.input_type
class MeshTagListArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str],
                 passthrough: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)
        if passthrough is not None:
            pulumi.set(__self__, "passthrough", passthrough)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def passthrough(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "passthrough")

    @passthrough.setter
    def passthrough(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "passthrough", value)


@pulumi.input_type
class PrometheusAttachmentPrometheusArgs:
    def __init__(__self__, *,
                 custom_prom: Optional[pulumi.Input['PrometheusAttachmentPrometheusCustomPromArgs']] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        if custom_prom is not None:
            pulumi.set(__self__, "custom_prom", custom_prom)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="customProm")
    def custom_prom(self) -> Optional[pulumi.Input['PrometheusAttachmentPrometheusCustomPromArgs']]:
        return pulumi.get(self, "custom_prom")

    @custom_prom.setter
    def custom_prom(self, value: Optional[pulumi.Input['PrometheusAttachmentPrometheusCustomPromArgs']]):
        pulumi.set(self, "custom_prom", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class PrometheusAttachmentPrometheusCustomPromArgs:
    def __init__(__self__, *,
                 auth_type: pulumi.Input[str],
                 url: pulumi.Input[str],
                 is_public_addr: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "auth_type", auth_type)
        pulumi.set(__self__, "url", url)
        if is_public_addr is not None:
            pulumi.set(__self__, "is_public_addr", is_public_addr)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "auth_type")

    @auth_type.setter
    def auth_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "auth_type", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="isPublicAddr")
    def is_public_addr(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "is_public_addr")

    @is_public_addr.setter
    def is_public_addr(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public_addr", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class TracingConfigApmArgs:
    def __init__(__self__, *,
                 enable: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class TracingConfigZipkinArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str]):
        pulumi.set(__self__, "address", address)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

