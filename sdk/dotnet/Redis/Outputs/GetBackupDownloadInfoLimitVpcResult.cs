// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Redis.Outputs
{

    [OutputType]
    public sealed class GetBackupDownloadInfoLimitVpcResult
    {
        public readonly string Region;
        public readonly ImmutableArray<string> VpcLists;

        [OutputConstructor]
        private GetBackupDownloadInfoLimitVpcResult(
            string region,

            ImmutableArray<string> vpcLists)
        {
            Region = region;
            VpcLists = vpcLists;
        }
    }
}