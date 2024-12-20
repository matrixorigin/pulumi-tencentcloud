// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Postgresql
{
    public static class GetDedicatedClusters
    {
        public static Task<GetDedicatedClustersResult> InvokeAsync(GetDedicatedClustersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDedicatedClustersResult>("tencentcloud:Postgresql/getDedicatedClusters:getDedicatedClusters", args ?? new GetDedicatedClustersArgs(), options.WithDefaults());

        public static Output<GetDedicatedClustersResult> Invoke(GetDedicatedClustersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDedicatedClustersResult>("tencentcloud:Postgresql/getDedicatedClusters:getDedicatedClusters", args ?? new GetDedicatedClustersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDedicatedClustersArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetDedicatedClustersFilterArgs>? _filters;
        public List<Inputs.GetDedicatedClustersFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetDedicatedClustersFilterArgs>());
            set => _filters = value;
        }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDedicatedClustersArgs()
        {
        }
        public static new GetDedicatedClustersArgs Empty => new GetDedicatedClustersArgs();
    }

    public sealed class GetDedicatedClustersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetDedicatedClustersFilterInputArgs>? _filters;
        public InputList<Inputs.GetDedicatedClustersFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetDedicatedClustersFilterInputArgs>());
            set => _filters = value;
        }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDedicatedClustersInvokeArgs()
        {
        }
        public static new GetDedicatedClustersInvokeArgs Empty => new GetDedicatedClustersInvokeArgs();
    }


    [OutputType]
    public sealed class GetDedicatedClustersResult
    {
        public readonly ImmutableArray<Outputs.GetDedicatedClustersDedicatedClusterSetResult> DedicatedClusterSets;
        public readonly ImmutableArray<Outputs.GetDedicatedClustersFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDedicatedClustersResult(
            ImmutableArray<Outputs.GetDedicatedClustersDedicatedClusterSetResult> dedicatedClusterSets,

            ImmutableArray<Outputs.GetDedicatedClustersFilterResult> filters,

            string id,

            string? resultOutputFile)
        {
            DedicatedClusterSets = dedicatedClusterSets;
            Filters = filters;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
