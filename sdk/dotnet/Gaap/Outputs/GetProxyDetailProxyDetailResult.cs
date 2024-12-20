// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Gaap.Outputs
{

    [OutputType]
    public sealed class GetProxyDetailProxyDetailResult
    {
        public readonly string AccessRegion;
        public readonly ImmutableArray<Outputs.GetProxyDetailProxyDetailAccessRegionInfoResult> AccessRegionInfos;
        public readonly string BanStatus;
        public readonly int Bandwidth;
        public readonly int BillingType;
        public readonly ImmutableArray<int> ClientIpMethods;
        public readonly int Concurrent;
        public readonly int CreateTime;
        public readonly string Domain;
        public readonly int FeatureBitmap;
        public readonly string ForwardIp;
        public readonly string GroupId;
        public readonly int Http3Supported;
        public readonly int InBanBlacklist;
        public readonly string InstanceId;
        public readonly string Ip;
        public readonly string IpAddressVersion;
        public readonly ImmutableArray<Outputs.GetProxyDetailProxyDetailIpListResult> IpLists;
        public readonly int IsSupportTlsChoice;
        public readonly int ModifyConfigTime;
        public readonly string NetworkType;
        public readonly string PackageType;
        public readonly string PolicyId;
        public readonly int ProjectId;
        public readonly string ProxyId;
        public readonly string ProxyName;
        public readonly int ProxyType;
        public readonly string RealServerRegion;
        public readonly ImmutableArray<Outputs.GetProxyDetailProxyDetailRealServerRegionInfoResult> RealServerRegionInfos;
        public readonly ImmutableArray<string> RelatedGlobalDomains;
        public readonly int Scalarable;
        public readonly string Status;
        public readonly ImmutableArray<string> SupportProtocols;
        public readonly int SupportSecurity;
        public readonly ImmutableArray<Outputs.GetProxyDetailProxyDetailTagSetResult> TagSets;
        public readonly string Version;

        [OutputConstructor]
        private GetProxyDetailProxyDetailResult(
            string accessRegion,

            ImmutableArray<Outputs.GetProxyDetailProxyDetailAccessRegionInfoResult> accessRegionInfos,

            string banStatus,

            int bandwidth,

            int billingType,

            ImmutableArray<int> clientIpMethods,

            int concurrent,

            int createTime,

            string domain,

            int featureBitmap,

            string forwardIp,

            string groupId,

            int http3Supported,

            int inBanBlacklist,

            string instanceId,

            string ip,

            string ipAddressVersion,

            ImmutableArray<Outputs.GetProxyDetailProxyDetailIpListResult> ipLists,

            int isSupportTlsChoice,

            int modifyConfigTime,

            string networkType,

            string packageType,

            string policyId,

            int projectId,

            string proxyId,

            string proxyName,

            int proxyType,

            string realServerRegion,

            ImmutableArray<Outputs.GetProxyDetailProxyDetailRealServerRegionInfoResult> realServerRegionInfos,

            ImmutableArray<string> relatedGlobalDomains,

            int scalarable,

            string status,

            ImmutableArray<string> supportProtocols,

            int supportSecurity,

            ImmutableArray<Outputs.GetProxyDetailProxyDetailTagSetResult> tagSets,

            string version)
        {
            AccessRegion = accessRegion;
            AccessRegionInfos = accessRegionInfos;
            BanStatus = banStatus;
            Bandwidth = bandwidth;
            BillingType = billingType;
            ClientIpMethods = clientIpMethods;
            Concurrent = concurrent;
            CreateTime = createTime;
            Domain = domain;
            FeatureBitmap = featureBitmap;
            ForwardIp = forwardIp;
            GroupId = groupId;
            Http3Supported = http3Supported;
            InBanBlacklist = inBanBlacklist;
            InstanceId = instanceId;
            Ip = ip;
            IpAddressVersion = ipAddressVersion;
            IpLists = ipLists;
            IsSupportTlsChoice = isSupportTlsChoice;
            ModifyConfigTime = modifyConfigTime;
            NetworkType = networkType;
            PackageType = packageType;
            PolicyId = policyId;
            ProjectId = projectId;
            ProxyId = proxyId;
            ProxyName = proxyName;
            ProxyType = proxyType;
            RealServerRegion = realServerRegion;
            RealServerRegionInfos = realServerRegionInfos;
            RelatedGlobalDomains = relatedGlobalDomains;
            Scalarable = scalarable;
            Status = status;
            SupportProtocols = supportProtocols;
            SupportSecurity = supportSecurity;
            TagSets = tagSets;
            Version = version;
        }
    }
}
