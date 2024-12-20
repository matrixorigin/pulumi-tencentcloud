// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssl
{
    public static class GetDescribeHostCosInstanceList
    {
        public static Task<GetDescribeHostCosInstanceListResult> InvokeAsync(GetDescribeHostCosInstanceListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDescribeHostCosInstanceListResult>("tencentcloud:Ssl/getDescribeHostCosInstanceList:getDescribeHostCosInstanceList", args ?? new GetDescribeHostCosInstanceListArgs(), options.WithDefaults());

        public static Output<GetDescribeHostCosInstanceListResult> Invoke(GetDescribeHostCosInstanceListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDescribeHostCosInstanceListResult>("tencentcloud:Ssl/getDescribeHostCosInstanceList:getDescribeHostCosInstanceList", args ?? new GetDescribeHostCosInstanceListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeHostCosInstanceListArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateId", required: true)]
        public string CertificateId { get; set; } = null!;

        [Input("filters")]
        private List<Inputs.GetDescribeHostCosInstanceListFilterArgs>? _filters;
        public List<Inputs.GetDescribeHostCosInstanceListFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetDescribeHostCosInstanceListFilterArgs>());
            set => _filters = value;
        }

        [Input("isCache")]
        public int? IsCache { get; set; }

        [Input("resourceType", required: true)]
        public string ResourceType { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDescribeHostCosInstanceListArgs()
        {
        }
        public static new GetDescribeHostCosInstanceListArgs Empty => new GetDescribeHostCosInstanceListArgs();
    }

    public sealed class GetDescribeHostCosInstanceListInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        [Input("filters")]
        private InputList<Inputs.GetDescribeHostCosInstanceListFilterInputArgs>? _filters;
        public InputList<Inputs.GetDescribeHostCosInstanceListFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetDescribeHostCosInstanceListFilterInputArgs>());
            set => _filters = value;
        }

        [Input("isCache")]
        public Input<int>? IsCache { get; set; }

        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDescribeHostCosInstanceListInvokeArgs()
        {
        }
        public static new GetDescribeHostCosInstanceListInvokeArgs Empty => new GetDescribeHostCosInstanceListInvokeArgs();
    }


    [OutputType]
    public sealed class GetDescribeHostCosInstanceListResult
    {
        public readonly string AsyncCacheTime;
        public readonly int AsyncOffset;
        public readonly int AsyncTotalNum;
        public readonly string CertificateId;
        public readonly ImmutableArray<Outputs.GetDescribeHostCosInstanceListFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetDescribeHostCosInstanceListInstanceListResult> InstanceLists;
        public readonly int? IsCache;
        public readonly string ResourceType;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDescribeHostCosInstanceListResult(
            string asyncCacheTime,

            int asyncOffset,

            int asyncTotalNum,

            string certificateId,

            ImmutableArray<Outputs.GetDescribeHostCosInstanceListFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetDescribeHostCosInstanceListInstanceListResult> instanceLists,

            int? isCache,

            string resourceType,

            string? resultOutputFile)
        {
            AsyncCacheTime = asyncCacheTime;
            AsyncOffset = asyncOffset;
            AsyncTotalNum = asyncTotalNum;
            CertificateId = certificateId;
            Filters = filters;
            Id = id;
            InstanceLists = instanceLists;
            IsCache = isCache;
            ResourceType = resourceType;
            ResultOutputFile = resultOutputFile;
        }
    }
}
