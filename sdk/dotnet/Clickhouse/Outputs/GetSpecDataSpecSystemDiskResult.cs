// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clickhouse.Outputs
{

    [OutputType]
    public sealed class GetSpecDataSpecSystemDiskResult
    {
        public readonly int DiskCount;
        public readonly string DiskDesc;
        public readonly string DiskType;
        public readonly int MaxDiskSize;
        public readonly int MinDiskSize;

        [OutputConstructor]
        private GetSpecDataSpecSystemDiskResult(
            int diskCount,

            string diskDesc,

            string diskType,

            int maxDiskSize,

            int minDiskSize)
        {
            DiskCount = diskCount;
            DiskDesc = diskDesc;
            DiskType = diskType;
            MaxDiskSize = maxDiskSize;
            MinDiskSize = minDiskSize;
        }
    }
}