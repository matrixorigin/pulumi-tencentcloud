// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssl
{
    public static class GetDescribeHostVodInstanceList
    {
        public static Task<GetDescribeHostVodInstanceListResult> InvokeAsync(GetDescribeHostVodInstanceListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDescribeHostVodInstanceListResult>("tencentcloud:Ssl/getDescribeHostVodInstanceList:getDescribeHostVodInstanceList", args ?? new GetDescribeHostVodInstanceListArgs(), options.WithDefaults());

        public static Output<GetDescribeHostVodInstanceListResult> Invoke(GetDescribeHostVodInstanceListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDescribeHostVodInstanceListResult>("tencentcloud:Ssl/getDescribeHostVodInstanceList:getDescribeHostVodInstanceList", args ?? new GetDescribeHostVodInstanceListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeHostVodInstanceListArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateId", required: true)]
        public string CertificateId { get; set; } = null!;

        [Input("filters")]
        private List<Inputs.GetDescribeHostVodInstanceListFilterArgs>? _filters;
        public List<Inputs.GetDescribeHostVodInstanceListFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetDescribeHostVodInstanceListFilterArgs>());
            set => _filters = value;
        }

        [Input("isCache")]
        public int? IsCache { get; set; }

        [Input("oldCertificateId")]
        public string? OldCertificateId { get; set; }

        [Input("resourceType", required: true)]
        public string ResourceType { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDescribeHostVodInstanceListArgs()
        {
        }
        public static new GetDescribeHostVodInstanceListArgs Empty => new GetDescribeHostVodInstanceListArgs();
    }

    public sealed class GetDescribeHostVodInstanceListInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        [Input("filters")]
        private InputList<Inputs.GetDescribeHostVodInstanceListFilterInputArgs>? _filters;
        public InputList<Inputs.GetDescribeHostVodInstanceListFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetDescribeHostVodInstanceListFilterInputArgs>());
            set => _filters = value;
        }

        [Input("isCache")]
        public Input<int>? IsCache { get; set; }

        [Input("oldCertificateId")]
        public Input<string>? OldCertificateId { get; set; }

        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDescribeHostVodInstanceListInvokeArgs()
        {
        }
        public static new GetDescribeHostVodInstanceListInvokeArgs Empty => new GetDescribeHostVodInstanceListInvokeArgs();
    }


    [OutputType]
    public sealed class GetDescribeHostVodInstanceListResult
    {
        public readonly string CertificateId;
        public readonly ImmutableArray<Outputs.GetDescribeHostVodInstanceListFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetDescribeHostVodInstanceListInstanceListResult> InstanceLists;
        public readonly int? IsCache;
        public readonly string? OldCertificateId;
        public readonly string ResourceType;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDescribeHostVodInstanceListResult(
            string certificateId,

            ImmutableArray<Outputs.GetDescribeHostVodInstanceListFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetDescribeHostVodInstanceListInstanceListResult> instanceLists,

            int? isCache,

            string? oldCertificateId,

            string resourceType,

            string? resultOutputFile)
        {
            CertificateId = certificateId;
            Filters = filters;
            Id = id;
            InstanceLists = instanceLists;
            IsCache = isCache;
            OldCertificateId = oldCertificateId;
            ResourceType = resourceType;
            ResultOutputFile = resultOutputFile;
        }
    }
}
