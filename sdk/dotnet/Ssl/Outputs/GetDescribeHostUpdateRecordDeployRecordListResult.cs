// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssl.Outputs
{

    [OutputType]
    public sealed class GetDescribeHostUpdateRecordDeployRecordListResult
    {
        public readonly string CertId;
        public readonly string CreateTime;
        public readonly int Id;
        public readonly string OldCertId;
        public readonly ImmutableArray<string> Regions;
        public readonly ImmutableArray<string> ResourceTypes;
        public readonly int Status;
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetDescribeHostUpdateRecordDeployRecordListResult(
            string certId,

            string createTime,

            int id,

            string oldCertId,

            ImmutableArray<string> regions,

            ImmutableArray<string> resourceTypes,

            int status,

            string updateTime)
        {
            CertId = certId;
            CreateTime = createTime;
            Id = id;
            OldCertId = oldCertId;
            Regions = regions;
            ResourceTypes = resourceTypes;
            Status = status;
            UpdateTime = updateTime;
        }
    }
}