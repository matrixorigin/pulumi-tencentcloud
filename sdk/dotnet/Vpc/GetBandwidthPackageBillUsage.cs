// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vpc
{
    public static class GetBandwidthPackageBillUsage
    {
        public static Task<GetBandwidthPackageBillUsageResult> InvokeAsync(GetBandwidthPackageBillUsageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBandwidthPackageBillUsageResult>("tencentcloud:Vpc/getBandwidthPackageBillUsage:getBandwidthPackageBillUsage", args ?? new GetBandwidthPackageBillUsageArgs(), options.WithDefaults());

        public static Output<GetBandwidthPackageBillUsageResult> Invoke(GetBandwidthPackageBillUsageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBandwidthPackageBillUsageResult>("tencentcloud:Vpc/getBandwidthPackageBillUsage:getBandwidthPackageBillUsage", args ?? new GetBandwidthPackageBillUsageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBandwidthPackageBillUsageArgs : global::Pulumi.InvokeArgs
    {
        [Input("bandwidthPackageId", required: true)]
        public string BandwidthPackageId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetBandwidthPackageBillUsageArgs()
        {
        }
        public static new GetBandwidthPackageBillUsageArgs Empty => new GetBandwidthPackageBillUsageArgs();
    }

    public sealed class GetBandwidthPackageBillUsageInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("bandwidthPackageId", required: true)]
        public Input<string> BandwidthPackageId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetBandwidthPackageBillUsageInvokeArgs()
        {
        }
        public static new GetBandwidthPackageBillUsageInvokeArgs Empty => new GetBandwidthPackageBillUsageInvokeArgs();
    }


    [OutputType]
    public sealed class GetBandwidthPackageBillUsageResult
    {
        public readonly ImmutableArray<Outputs.GetBandwidthPackageBillUsageBandwidthPackageBillBandwidthSetResult> BandwidthPackageBillBandwidthSets;
        public readonly string BandwidthPackageId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetBandwidthPackageBillUsageResult(
            ImmutableArray<Outputs.GetBandwidthPackageBillUsageBandwidthPackageBillBandwidthSetResult> bandwidthPackageBillBandwidthSets,

            string bandwidthPackageId,

            string id,

            string? resultOutputFile)
        {
            BandwidthPackageBillBandwidthSets = bandwidthPackageBillBandwidthSets;
            BandwidthPackageId = bandwidthPackageId;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}