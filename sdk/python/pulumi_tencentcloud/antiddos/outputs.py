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
    'CcBlackWhiteIpBlackWhiteIp',
    'CcPrecisionPolicyPolicyList',
    'DdosGeoIpBlockConfigDdosGeoIpBlockConfig',
    'DdosSpeedLimitConfigDdosSpeedLimitConfig',
    'DdosSpeedLimitConfigDdosSpeedLimitConfigDstPortScope',
    'DdosSpeedLimitConfigDdosSpeedLimitConfigSpeedValue',
    'DefaultAlarmThresholdDefaultAlarmConfig',
    'PacketFilterConfigPacketFilterConfig',
    'PortAclConfigAclConfig',
    'GetBasicDeviceStatusClbDataResult',
    'GetBasicDeviceStatusDataResult',
    'GetListListenerLayer4ListenerResult',
    'GetListListenerLayer4ListenerInstanceDetailResult',
    'GetListListenerLayer4ListenerInstanceDetailRuleResult',
    'GetListListenerLayer4ListenerRealServerResult',
    'GetListListenerLayer7ListenerResult',
    'GetListListenerLayer7ListenerInstanceDetailResult',
    'GetListListenerLayer7ListenerInstanceDetailRuleResult',
    'GetListListenerLayer7ListenerProxyTypeListResult',
    'GetListListenerLayer7ListenerRealServerResult',
    'GetOverviewDdosEventListEventListResult',
]

@pulumi.output_type
class CcBlackWhiteIpBlackWhiteIp(dict):
    def __init__(__self__, *,
                 ip: str,
                 mask: int):
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "mask", mask)

    @property
    @pulumi.getter
    def ip(self) -> str:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def mask(self) -> int:
        return pulumi.get(self, "mask")


@pulumi.output_type
class CcPrecisionPolicyPolicyList(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fieldName":
            suggest = "field_name"
        elif key == "fieldType":
            suggest = "field_type"
        elif key == "valueOperator":
            suggest = "value_operator"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CcPrecisionPolicyPolicyList. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CcPrecisionPolicyPolicyList.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CcPrecisionPolicyPolicyList.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 field_name: str,
                 field_type: str,
                 value: str,
                 value_operator: str):
        pulumi.set(__self__, "field_name", field_name)
        pulumi.set(__self__, "field_type", field_type)
        pulumi.set(__self__, "value", value)
        pulumi.set(__self__, "value_operator", value_operator)

    @property
    @pulumi.getter(name="fieldName")
    def field_name(self) -> str:
        return pulumi.get(self, "field_name")

    @property
    @pulumi.getter(name="fieldType")
    def field_type(self) -> str:
        return pulumi.get(self, "field_type")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="valueOperator")
    def value_operator(self) -> str:
        return pulumi.get(self, "value_operator")


@pulumi.output_type
class DdosGeoIpBlockConfigDdosGeoIpBlockConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "regionType":
            suggest = "region_type"
        elif key == "areaLists":
            suggest = "area_lists"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DdosGeoIpBlockConfigDdosGeoIpBlockConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DdosGeoIpBlockConfigDdosGeoIpBlockConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DdosGeoIpBlockConfigDdosGeoIpBlockConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: str,
                 region_type: str,
                 area_lists: Optional[Sequence[int]] = None):
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "region_type", region_type)
        if area_lists is not None:
            pulumi.set(__self__, "area_lists", area_lists)

    @property
    @pulumi.getter
    def action(self) -> str:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="regionType")
    def region_type(self) -> str:
        return pulumi.get(self, "region_type")

    @property
    @pulumi.getter(name="areaLists")
    def area_lists(self) -> Optional[Sequence[int]]:
        return pulumi.get(self, "area_lists")


@pulumi.output_type
class DdosSpeedLimitConfigDdosSpeedLimitConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "speedValues":
            suggest = "speed_values"
        elif key == "dstPortList":
            suggest = "dst_port_list"
        elif key == "dstPortScopes":
            suggest = "dst_port_scopes"
        elif key == "protocolList":
            suggest = "protocol_list"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DdosSpeedLimitConfigDdosSpeedLimitConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DdosSpeedLimitConfigDdosSpeedLimitConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DdosSpeedLimitConfigDdosSpeedLimitConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 mode: int,
                 speed_values: Sequence['outputs.DdosSpeedLimitConfigDdosSpeedLimitConfigSpeedValue'],
                 dst_port_list: Optional[str] = None,
                 dst_port_scopes: Optional[Sequence['outputs.DdosSpeedLimitConfigDdosSpeedLimitConfigDstPortScope']] = None,
                 protocol_list: Optional[str] = None):
        pulumi.set(__self__, "mode", mode)
        pulumi.set(__self__, "speed_values", speed_values)
        if dst_port_list is not None:
            pulumi.set(__self__, "dst_port_list", dst_port_list)
        if dst_port_scopes is not None:
            pulumi.set(__self__, "dst_port_scopes", dst_port_scopes)
        if protocol_list is not None:
            pulumi.set(__self__, "protocol_list", protocol_list)

    @property
    @pulumi.getter
    def mode(self) -> int:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="speedValues")
    def speed_values(self) -> Sequence['outputs.DdosSpeedLimitConfigDdosSpeedLimitConfigSpeedValue']:
        return pulumi.get(self, "speed_values")

    @property
    @pulumi.getter(name="dstPortList")
    def dst_port_list(self) -> Optional[str]:
        return pulumi.get(self, "dst_port_list")

    @property
    @pulumi.getter(name="dstPortScopes")
    def dst_port_scopes(self) -> Optional[Sequence['outputs.DdosSpeedLimitConfigDdosSpeedLimitConfigDstPortScope']]:
        return pulumi.get(self, "dst_port_scopes")

    @property
    @pulumi.getter(name="protocolList")
    def protocol_list(self) -> Optional[str]:
        return pulumi.get(self, "protocol_list")


@pulumi.output_type
class DdosSpeedLimitConfigDdosSpeedLimitConfigDstPortScope(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "beginPort":
            suggest = "begin_port"
        elif key == "endPort":
            suggest = "end_port"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DdosSpeedLimitConfigDdosSpeedLimitConfigDstPortScope. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DdosSpeedLimitConfigDdosSpeedLimitConfigDstPortScope.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DdosSpeedLimitConfigDdosSpeedLimitConfigDstPortScope.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 begin_port: int,
                 end_port: int):
        pulumi.set(__self__, "begin_port", begin_port)
        pulumi.set(__self__, "end_port", end_port)

    @property
    @pulumi.getter(name="beginPort")
    def begin_port(self) -> int:
        return pulumi.get(self, "begin_port")

    @property
    @pulumi.getter(name="endPort")
    def end_port(self) -> int:
        return pulumi.get(self, "end_port")


@pulumi.output_type
class DdosSpeedLimitConfigDdosSpeedLimitConfigSpeedValue(dict):
    def __init__(__self__, *,
                 type: int,
                 value: int):
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> int:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> int:
        return pulumi.get(self, "value")


@pulumi.output_type
class DefaultAlarmThresholdDefaultAlarmConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "alarmThreshold":
            suggest = "alarm_threshold"
        elif key == "alarmType":
            suggest = "alarm_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DefaultAlarmThresholdDefaultAlarmConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DefaultAlarmThresholdDefaultAlarmConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DefaultAlarmThresholdDefaultAlarmConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 alarm_threshold: Optional[int] = None,
                 alarm_type: Optional[int] = None):
        if alarm_threshold is not None:
            pulumi.set(__self__, "alarm_threshold", alarm_threshold)
        if alarm_type is not None:
            pulumi.set(__self__, "alarm_type", alarm_type)

    @property
    @pulumi.getter(name="alarmThreshold")
    def alarm_threshold(self) -> Optional[int]:
        return pulumi.get(self, "alarm_threshold")

    @property
    @pulumi.getter(name="alarmType")
    def alarm_type(self) -> Optional[int]:
        return pulumi.get(self, "alarm_type")


@pulumi.output_type
class PacketFilterConfigPacketFilterConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dportEnd":
            suggest = "dport_end"
        elif key == "dportStart":
            suggest = "dport_start"
        elif key == "pktlenMax":
            suggest = "pktlen_max"
        elif key == "pktlenMin":
            suggest = "pktlen_min"
        elif key == "sportEnd":
            suggest = "sport_end"
        elif key == "sportStart":
            suggest = "sport_start"
        elif key == "isNot":
            suggest = "is_not"
        elif key == "isNot2":
            suggest = "is_not2"
        elif key == "matchBegin":
            suggest = "match_begin"
        elif key == "matchBegin2":
            suggest = "match_begin2"
        elif key == "matchLogic":
            suggest = "match_logic"
        elif key == "matchType":
            suggest = "match_type"
        elif key == "matchType2":
            suggest = "match_type2"
        elif key == "pktLenGt":
            suggest = "pkt_len_gt"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PacketFilterConfigPacketFilterConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PacketFilterConfigPacketFilterConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PacketFilterConfigPacketFilterConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: str,
                 dport_end: int,
                 dport_start: int,
                 pktlen_max: int,
                 pktlen_min: int,
                 protocol: str,
                 sport_end: int,
                 sport_start: int,
                 depth: Optional[int] = None,
                 depth2: Optional[int] = None,
                 is_not: Optional[int] = None,
                 is_not2: Optional[int] = None,
                 match_begin: Optional[str] = None,
                 match_begin2: Optional[str] = None,
                 match_logic: Optional[str] = None,
                 match_type: Optional[str] = None,
                 match_type2: Optional[str] = None,
                 offset: Optional[int] = None,
                 offset2: Optional[int] = None,
                 pkt_len_gt: Optional[int] = None,
                 str: Optional[str] = None,
                 str2: Optional[str] = None):
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "dport_end", dport_end)
        pulumi.set(__self__, "dport_start", dport_start)
        pulumi.set(__self__, "pktlen_max", pktlen_max)
        pulumi.set(__self__, "pktlen_min", pktlen_min)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "sport_end", sport_end)
        pulumi.set(__self__, "sport_start", sport_start)
        if depth is not None:
            pulumi.set(__self__, "depth", depth)
        if depth2 is not None:
            pulumi.set(__self__, "depth2", depth2)
        if is_not is not None:
            pulumi.set(__self__, "is_not", is_not)
        if is_not2 is not None:
            pulumi.set(__self__, "is_not2", is_not2)
        if match_begin is not None:
            pulumi.set(__self__, "match_begin", match_begin)
        if match_begin2 is not None:
            pulumi.set(__self__, "match_begin2", match_begin2)
        if match_logic is not None:
            pulumi.set(__self__, "match_logic", match_logic)
        if match_type is not None:
            pulumi.set(__self__, "match_type", match_type)
        if match_type2 is not None:
            pulumi.set(__self__, "match_type2", match_type2)
        if offset is not None:
            pulumi.set(__self__, "offset", offset)
        if offset2 is not None:
            pulumi.set(__self__, "offset2", offset2)
        if pkt_len_gt is not None:
            pulumi.set(__self__, "pkt_len_gt", pkt_len_gt)
        if str is not None:
            pulumi.set(__self__, "str", str)
        if str2 is not None:
            pulumi.set(__self__, "str2", str2)

    @property
    @pulumi.getter
    def action(self) -> str:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="dportEnd")
    def dport_end(self) -> int:
        return pulumi.get(self, "dport_end")

    @property
    @pulumi.getter(name="dportStart")
    def dport_start(self) -> int:
        return pulumi.get(self, "dport_start")

    @property
    @pulumi.getter(name="pktlenMax")
    def pktlen_max(self) -> int:
        return pulumi.get(self, "pktlen_max")

    @property
    @pulumi.getter(name="pktlenMin")
    def pktlen_min(self) -> int:
        return pulumi.get(self, "pktlen_min")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="sportEnd")
    def sport_end(self) -> int:
        return pulumi.get(self, "sport_end")

    @property
    @pulumi.getter(name="sportStart")
    def sport_start(self) -> int:
        return pulumi.get(self, "sport_start")

    @property
    @pulumi.getter
    def depth(self) -> Optional[int]:
        return pulumi.get(self, "depth")

    @property
    @pulumi.getter
    def depth2(self) -> Optional[int]:
        return pulumi.get(self, "depth2")

    @property
    @pulumi.getter(name="isNot")
    def is_not(self) -> Optional[int]:
        return pulumi.get(self, "is_not")

    @property
    @pulumi.getter(name="isNot2")
    def is_not2(self) -> Optional[int]:
        return pulumi.get(self, "is_not2")

    @property
    @pulumi.getter(name="matchBegin")
    def match_begin(self) -> Optional[str]:
        return pulumi.get(self, "match_begin")

    @property
    @pulumi.getter(name="matchBegin2")
    def match_begin2(self) -> Optional[str]:
        return pulumi.get(self, "match_begin2")

    @property
    @pulumi.getter(name="matchLogic")
    def match_logic(self) -> Optional[str]:
        return pulumi.get(self, "match_logic")

    @property
    @pulumi.getter(name="matchType")
    def match_type(self) -> Optional[str]:
        return pulumi.get(self, "match_type")

    @property
    @pulumi.getter(name="matchType2")
    def match_type2(self) -> Optional[str]:
        return pulumi.get(self, "match_type2")

    @property
    @pulumi.getter
    def offset(self) -> Optional[int]:
        return pulumi.get(self, "offset")

    @property
    @pulumi.getter
    def offset2(self) -> Optional[int]:
        return pulumi.get(self, "offset2")

    @property
    @pulumi.getter(name="pktLenGt")
    def pkt_len_gt(self) -> Optional[int]:
        return pulumi.get(self, "pkt_len_gt")

    @property
    @pulumi.getter
    def str(self) -> Optional[str]:
        return pulumi.get(self, "str")

    @property
    @pulumi.getter
    def str2(self) -> Optional[str]:
        return pulumi.get(self, "str2")


@pulumi.output_type
class PortAclConfigAclConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dPortEnd":
            suggest = "d_port_end"
        elif key == "dPortStart":
            suggest = "d_port_start"
        elif key == "forwardProtocol":
            suggest = "forward_protocol"
        elif key == "sPortEnd":
            suggest = "s_port_end"
        elif key == "sPortStart":
            suggest = "s_port_start"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PortAclConfigAclConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PortAclConfigAclConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PortAclConfigAclConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: str,
                 d_port_end: int,
                 d_port_start: int,
                 forward_protocol: str,
                 s_port_end: int,
                 s_port_start: int,
                 priority: Optional[int] = None):
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "d_port_end", d_port_end)
        pulumi.set(__self__, "d_port_start", d_port_start)
        pulumi.set(__self__, "forward_protocol", forward_protocol)
        pulumi.set(__self__, "s_port_end", s_port_end)
        pulumi.set(__self__, "s_port_start", s_port_start)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)

    @property
    @pulumi.getter
    def action(self) -> str:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="dPortEnd")
    def d_port_end(self) -> int:
        return pulumi.get(self, "d_port_end")

    @property
    @pulumi.getter(name="dPortStart")
    def d_port_start(self) -> int:
        return pulumi.get(self, "d_port_start")

    @property
    @pulumi.getter(name="forwardProtocol")
    def forward_protocol(self) -> str:
        return pulumi.get(self, "forward_protocol")

    @property
    @pulumi.getter(name="sPortEnd")
    def s_port_end(self) -> int:
        return pulumi.get(self, "s_port_end")

    @property
    @pulumi.getter(name="sPortStart")
    def s_port_start(self) -> int:
        return pulumi.get(self, "s_port_start")

    @property
    @pulumi.getter
    def priority(self) -> Optional[int]:
        return pulumi.get(self, "priority")


@pulumi.output_type
class GetBasicDeviceStatusClbDataResult(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class GetBasicDeviceStatusDataResult(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class GetListListenerLayer4ListenerResult(dict):
    def __init__(__self__, *,
                 backend_port: int,
                 frontend_port: int,
                 instance_detail_rules: Sequence['outputs.GetListListenerLayer4ListenerInstanceDetailRuleResult'],
                 instance_details: Sequence['outputs.GetListListenerLayer4ListenerInstanceDetailResult'],
                 protocol: str,
                 real_servers: Sequence['outputs.GetListListenerLayer4ListenerRealServerResult']):
        pulumi.set(__self__, "backend_port", backend_port)
        pulumi.set(__self__, "frontend_port", frontend_port)
        pulumi.set(__self__, "instance_detail_rules", instance_detail_rules)
        pulumi.set(__self__, "instance_details", instance_details)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "real_servers", real_servers)

    @property
    @pulumi.getter(name="backendPort")
    def backend_port(self) -> int:
        return pulumi.get(self, "backend_port")

    @property
    @pulumi.getter(name="frontendPort")
    def frontend_port(self) -> int:
        return pulumi.get(self, "frontend_port")

    @property
    @pulumi.getter(name="instanceDetailRules")
    def instance_detail_rules(self) -> Sequence['outputs.GetListListenerLayer4ListenerInstanceDetailRuleResult']:
        return pulumi.get(self, "instance_detail_rules")

    @property
    @pulumi.getter(name="instanceDetails")
    def instance_details(self) -> Sequence['outputs.GetListListenerLayer4ListenerInstanceDetailResult']:
        return pulumi.get(self, "instance_details")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="realServers")
    def real_servers(self) -> Sequence['outputs.GetListListenerLayer4ListenerRealServerResult']:
        return pulumi.get(self, "real_servers")


@pulumi.output_type
class GetListListenerLayer4ListenerInstanceDetailResult(dict):
    def __init__(__self__, *,
                 eip_lists: Sequence[str],
                 instance_id: str):
        pulumi.set(__self__, "eip_lists", eip_lists)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="eipLists")
    def eip_lists(self) -> Sequence[str]:
        return pulumi.get(self, "eip_lists")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")


@pulumi.output_type
class GetListListenerLayer4ListenerInstanceDetailRuleResult(dict):
    def __init__(__self__, *,
                 cname: str,
                 eip_lists: Sequence[str],
                 instance_id: str):
        pulumi.set(__self__, "cname", cname)
        pulumi.set(__self__, "eip_lists", eip_lists)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter
    def cname(self) -> str:
        return pulumi.get(self, "cname")

    @property
    @pulumi.getter(name="eipLists")
    def eip_lists(self) -> Sequence[str]:
        return pulumi.get(self, "eip_lists")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")


@pulumi.output_type
class GetListListenerLayer4ListenerRealServerResult(dict):
    def __init__(__self__, *,
                 port: int,
                 real_server: str,
                 rs_type: int,
                 weight: int):
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "real_server", real_server)
        pulumi.set(__self__, "rs_type", rs_type)
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="realServer")
    def real_server(self) -> str:
        return pulumi.get(self, "real_server")

    @property
    @pulumi.getter(name="rsType")
    def rs_type(self) -> int:
        return pulumi.get(self, "rs_type")

    @property
    @pulumi.getter
    def weight(self) -> int:
        return pulumi.get(self, "weight")


@pulumi.output_type
class GetListListenerLayer7ListenerResult(dict):
    def __init__(__self__, *,
                 domain: str,
                 instance_detail_rules: Sequence['outputs.GetListListenerLayer7ListenerInstanceDetailRuleResult'],
                 instance_details: Sequence['outputs.GetListListenerLayer7ListenerInstanceDetailResult'],
                 protocol: str,
                 proxy_type_lists: Sequence['outputs.GetListListenerLayer7ListenerProxyTypeListResult'],
                 real_servers: Sequence['outputs.GetListListenerLayer7ListenerRealServerResult'],
                 vport: int):
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "instance_detail_rules", instance_detail_rules)
        pulumi.set(__self__, "instance_details", instance_details)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "proxy_type_lists", proxy_type_lists)
        pulumi.set(__self__, "real_servers", real_servers)
        pulumi.set(__self__, "vport", vport)

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="instanceDetailRules")
    def instance_detail_rules(self) -> Sequence['outputs.GetListListenerLayer7ListenerInstanceDetailRuleResult']:
        return pulumi.get(self, "instance_detail_rules")

    @property
    @pulumi.getter(name="instanceDetails")
    def instance_details(self) -> Sequence['outputs.GetListListenerLayer7ListenerInstanceDetailResult']:
        return pulumi.get(self, "instance_details")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="proxyTypeLists")
    def proxy_type_lists(self) -> Sequence['outputs.GetListListenerLayer7ListenerProxyTypeListResult']:
        return pulumi.get(self, "proxy_type_lists")

    @property
    @pulumi.getter(name="realServers")
    def real_servers(self) -> Sequence['outputs.GetListListenerLayer7ListenerRealServerResult']:
        return pulumi.get(self, "real_servers")

    @property
    @pulumi.getter
    def vport(self) -> int:
        return pulumi.get(self, "vport")


@pulumi.output_type
class GetListListenerLayer7ListenerInstanceDetailResult(dict):
    def __init__(__self__, *,
                 eip_lists: Sequence[str],
                 instance_id: str):
        pulumi.set(__self__, "eip_lists", eip_lists)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="eipLists")
    def eip_lists(self) -> Sequence[str]:
        return pulumi.get(self, "eip_lists")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")


@pulumi.output_type
class GetListListenerLayer7ListenerInstanceDetailRuleResult(dict):
    def __init__(__self__, *,
                 cname: str,
                 eip_lists: Sequence[str],
                 instance_id: str):
        pulumi.set(__self__, "cname", cname)
        pulumi.set(__self__, "eip_lists", eip_lists)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter
    def cname(self) -> str:
        return pulumi.get(self, "cname")

    @property
    @pulumi.getter(name="eipLists")
    def eip_lists(self) -> Sequence[str]:
        return pulumi.get(self, "eip_lists")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")


@pulumi.output_type
class GetListListenerLayer7ListenerProxyTypeListResult(dict):
    def __init__(__self__, *,
                 proxy_ports: Sequence[int],
                 proxy_type: str):
        pulumi.set(__self__, "proxy_ports", proxy_ports)
        pulumi.set(__self__, "proxy_type", proxy_type)

    @property
    @pulumi.getter(name="proxyPorts")
    def proxy_ports(self) -> Sequence[int]:
        return pulumi.get(self, "proxy_ports")

    @property
    @pulumi.getter(name="proxyType")
    def proxy_type(self) -> str:
        return pulumi.get(self, "proxy_type")


@pulumi.output_type
class GetListListenerLayer7ListenerRealServerResult(dict):
    def __init__(__self__, *,
                 port: int,
                 real_server: str,
                 rs_type: int,
                 weight: int):
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "real_server", real_server)
        pulumi.set(__self__, "rs_type", rs_type)
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="realServer")
    def real_server(self) -> str:
        return pulumi.get(self, "real_server")

    @property
    @pulumi.getter(name="rsType")
    def rs_type(self) -> int:
        return pulumi.get(self, "rs_type")

    @property
    @pulumi.getter
    def weight(self) -> int:
        return pulumi.get(self, "weight")


@pulumi.output_type
class GetOverviewDdosEventListEventListResult(dict):
    def __init__(__self__, *,
                 attack_status: int,
                 attack_type: str,
                 business: str,
                 end_time: str,
                 id: str,
                 instance_id: str,
                 instance_name: str,
                 mbps: int,
                 pps: int,
                 start_time: str,
                 vip: str):
        pulumi.set(__self__, "attack_status", attack_status)
        pulumi.set(__self__, "attack_type", attack_type)
        pulumi.set(__self__, "business", business)
        pulumi.set(__self__, "end_time", end_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "mbps", mbps)
        pulumi.set(__self__, "pps", pps)
        pulumi.set(__self__, "start_time", start_time)
        pulumi.set(__self__, "vip", vip)

    @property
    @pulumi.getter(name="attackStatus")
    def attack_status(self) -> int:
        return pulumi.get(self, "attack_status")

    @property
    @pulumi.getter(name="attackType")
    def attack_type(self) -> str:
        return pulumi.get(self, "attack_type")

    @property
    @pulumi.getter
    def business(self) -> str:
        return pulumi.get(self, "business")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter
    def mbps(self) -> int:
        return pulumi.get(self, "mbps")

    @property
    @pulumi.getter
    def pps(self) -> int:
        return pulumi.get(self, "pps")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def vip(self) -> str:
        return pulumi.get(self, "vip")


