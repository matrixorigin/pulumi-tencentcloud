// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Classic.Outputs
{

    [OutputType]
    public sealed class GetElasticPublicIpv6sAddressSetResult
    {
        public readonly string AddressId;
        public readonly string AddressIp;
        public readonly string AddressName;
        public readonly string AddressStatus;
        public readonly string AddressType;
        public readonly string AntiDdosPackageId;
        public readonly int Bandwidth;
        public readonly string BandwidthPackageId;
        public readonly bool CascadeRelease;
        public readonly string CreatedTime;
        public readonly string DeadlineDate;
        public readonly string DedicatedClusterId;
        public readonly string Egress;
        public readonly Outputs.GetElasticPublicIpv6sAddressSetEipAlgTypeResult EipAlgType;
        public readonly string InstanceId;
        public readonly string InstanceType;
        public readonly string InternetChargeType;
        public readonly string InternetServiceProvider;
        public readonly bool IsArrears;
        public readonly bool IsBlocked;
        public readonly bool IsEipDirectConnection;
        public readonly bool LocalBgp;
        public readonly string NetworkInterfaceId;
        public readonly string PrivateAddressIp;
        public readonly string RenewFlag;
        public readonly ImmutableArray<Outputs.GetElasticPublicIpv6sAddressSetTagSetResult> TagSets;
        public readonly string UnVpcId;

        [OutputConstructor]
        private GetElasticPublicIpv6sAddressSetResult(
            string addressId,

            string addressIp,

            string addressName,

            string addressStatus,

            string addressType,

            string antiDdosPackageId,

            int bandwidth,

            string bandwidthPackageId,

            bool cascadeRelease,

            string createdTime,

            string deadlineDate,

            string dedicatedClusterId,

            string egress,

            Outputs.GetElasticPublicIpv6sAddressSetEipAlgTypeResult eipAlgType,

            string instanceId,

            string instanceType,

            string internetChargeType,

            string internetServiceProvider,

            bool isArrears,

            bool isBlocked,

            bool isEipDirectConnection,

            bool localBgp,

            string networkInterfaceId,

            string privateAddressIp,

            string renewFlag,

            ImmutableArray<Outputs.GetElasticPublicIpv6sAddressSetTagSetResult> tagSets,

            string unVpcId)
        {
            AddressId = addressId;
            AddressIp = addressIp;
            AddressName = addressName;
            AddressStatus = addressStatus;
            AddressType = addressType;
            AntiDdosPackageId = antiDdosPackageId;
            Bandwidth = bandwidth;
            BandwidthPackageId = bandwidthPackageId;
            CascadeRelease = cascadeRelease;
            CreatedTime = createdTime;
            DeadlineDate = deadlineDate;
            DedicatedClusterId = dedicatedClusterId;
            Egress = egress;
            EipAlgType = eipAlgType;
            InstanceId = instanceId;
            InstanceType = instanceType;
            InternetChargeType = internetChargeType;
            InternetServiceProvider = internetServiceProvider;
            IsArrears = isArrears;
            IsBlocked = isBlocked;
            IsEipDirectConnection = isEipDirectConnection;
            LocalBgp = localBgp;
            NetworkInterfaceId = networkInterfaceId;
            PrivateAddressIp = privateAddressIp;
            RenewFlag = renewFlag;
            TagSets = tagSets;
            UnVpcId = unVpcId;
        }
    }
}