// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kubernetes
{
    public static class GetCharts
    {
        public static Task<GetChartsResult> InvokeAsync(GetChartsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetChartsResult>("tencentcloud:Kubernetes/getCharts:getCharts", args ?? new GetChartsArgs(), options.WithDefaults());

        public static Output<GetChartsResult> Invoke(GetChartsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetChartsResult>("tencentcloud:Kubernetes/getCharts:getCharts", args ?? new GetChartsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetChartsArgs : global::Pulumi.InvokeArgs
    {
        [Input("arch")]
        public string? Arch { get; set; }

        [Input("clusterType")]
        public string? ClusterType { get; set; }

        [Input("kind")]
        public string? Kind { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetChartsArgs()
        {
        }
        public static new GetChartsArgs Empty => new GetChartsArgs();
    }

    public sealed class GetChartsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("arch")]
        public Input<string>? Arch { get; set; }

        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetChartsInvokeArgs()
        {
        }
        public static new GetChartsInvokeArgs Empty => new GetChartsInvokeArgs();
    }


    [OutputType]
    public sealed class GetChartsResult
    {
        public readonly string? Arch;
        public readonly ImmutableArray<Outputs.GetChartsChartListResult> ChartLists;
        public readonly string? ClusterType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Kind;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetChartsResult(
            string? arch,

            ImmutableArray<Outputs.GetChartsChartListResult> chartLists,

            string? clusterType,

            string id,

            string? kind,

            string? resultOutputFile)
        {
            Arch = arch;
            ChartLists = chartLists;
            ClusterType = clusterType;
            Id = id;
            Kind = kind;
            ResultOutputFile = resultOutputFile;
        }
    }
}