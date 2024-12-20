// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Oceanus.Outputs
{

    [OutputType]
    public sealed class GetClustersClusterSetSqlGatewayResourceRefResult
    {
        public readonly string ResourceId;
        public readonly int Type;
        public readonly int Version;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetClustersClusterSetSqlGatewayResourceRefResult(
            string resourceId,

            int type,

            int version,

            string workspaceId)
        {
            ResourceId = resourceId;
            Type = type;
            Version = version;
            WorkspaceId = workspaceId;
        }
    }
}
