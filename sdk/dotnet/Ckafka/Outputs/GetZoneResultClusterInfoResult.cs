// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class GetZoneResultClusterInfoResult
    {
        public readonly int AvailableBandWidth;
        public readonly int AvailableDiskSize;
        public readonly int ClusterId;
        public readonly string ClusterName;
        public readonly int MaxBandWidth;
        public readonly int MaxDiskSize;
        public readonly int ZoneId;
        public readonly ImmutableArray<int> ZoneIds;

        [OutputConstructor]
        private GetZoneResultClusterInfoResult(
            int availableBandWidth,

            int availableDiskSize,

            int clusterId,

            string clusterName,

            int maxBandWidth,

            int maxDiskSize,

            int zoneId,

            ImmutableArray<int> zoneIds)
        {
            AvailableBandWidth = availableBandWidth;
            AvailableDiskSize = availableDiskSize;
            ClusterId = clusterId;
            ClusterName = clusterName;
            MaxBandWidth = maxBandWidth;
            MaxDiskSize = maxDiskSize;
            ZoneId = zoneId;
            ZoneIds = zoneIds;
        }
    }
}