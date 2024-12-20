// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Emr
{
    public static class GetCvmQuota
    {
        public static Task<GetCvmQuotaResult> InvokeAsync(GetCvmQuotaArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCvmQuotaResult>("tencentcloud:Emr/getCvmQuota:getCvmQuota", args ?? new GetCvmQuotaArgs(), options.WithDefaults());

        public static Output<GetCvmQuotaResult> Invoke(GetCvmQuotaInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCvmQuotaResult>("tencentcloud:Emr/getCvmQuota:getCvmQuota", args ?? new GetCvmQuotaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCvmQuotaArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("zoneId")]
        public int? ZoneId { get; set; }

        public GetCvmQuotaArgs()
        {
        }
        public static new GetCvmQuotaArgs Empty => new GetCvmQuotaArgs();
    }

    public sealed class GetCvmQuotaInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("zoneId")]
        public Input<int>? ZoneId { get; set; }

        public GetCvmQuotaInvokeArgs()
        {
        }
        public static new GetCvmQuotaInvokeArgs Empty => new GetCvmQuotaInvokeArgs();
    }


    [OutputType]
    public sealed class GetCvmQuotaResult
    {
        public readonly string ClusterId;
        public readonly ImmutableArray<Outputs.GetCvmQuotaEksQuotaSetResult> EksQuotaSets;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetCvmQuotaPostPaidQuotaSetResult> PostPaidQuotaSets;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<Outputs.GetCvmQuotaSpotPaidQuotaSetResult> SpotPaidQuotaSets;
        public readonly int? ZoneId;

        [OutputConstructor]
        private GetCvmQuotaResult(
            string clusterId,

            ImmutableArray<Outputs.GetCvmQuotaEksQuotaSetResult> eksQuotaSets,

            string id,

            ImmutableArray<Outputs.GetCvmQuotaPostPaidQuotaSetResult> postPaidQuotaSets,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetCvmQuotaSpotPaidQuotaSetResult> spotPaidQuotaSets,

            int? zoneId)
        {
            ClusterId = clusterId;
            EksQuotaSets = eksQuotaSets;
            Id = id;
            PostPaidQuotaSets = postPaidQuotaSets;
            ResultOutputFile = resultOutputFile;
            SpotPaidQuotaSets = spotPaidQuotaSets;
            ZoneId = zoneId;
        }
    }
}
