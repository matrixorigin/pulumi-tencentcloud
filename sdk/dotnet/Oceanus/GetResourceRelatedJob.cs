// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Oceanus
{
    public static class GetResourceRelatedJob
    {
        public static Task<GetResourceRelatedJobResult> InvokeAsync(GetResourceRelatedJobArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourceRelatedJobResult>("tencentcloud:Oceanus/getResourceRelatedJob:getResourceRelatedJob", args ?? new GetResourceRelatedJobArgs(), options.WithDefaults());

        public static Output<GetResourceRelatedJobResult> Invoke(GetResourceRelatedJobInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceRelatedJobResult>("tencentcloud:Oceanus/getResourceRelatedJob:getResourceRelatedJob", args ?? new GetResourceRelatedJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceRelatedJobArgs : global::Pulumi.InvokeArgs
    {
        [Input("descByJobConfigCreateTime")]
        public int? DescByJobConfigCreateTime { get; set; }

        [Input("resourceConfigVersion")]
        public int? ResourceConfigVersion { get; set; }

        [Input("resourceId", required: true)]
        public string ResourceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("workSpaceId")]
        public string? WorkSpaceId { get; set; }

        public GetResourceRelatedJobArgs()
        {
        }
        public static new GetResourceRelatedJobArgs Empty => new GetResourceRelatedJobArgs();
    }

    public sealed class GetResourceRelatedJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("descByJobConfigCreateTime")]
        public Input<int>? DescByJobConfigCreateTime { get; set; }

        [Input("resourceConfigVersion")]
        public Input<int>? ResourceConfigVersion { get; set; }

        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("workSpaceId")]
        public Input<string>? WorkSpaceId { get; set; }

        public GetResourceRelatedJobInvokeArgs()
        {
        }
        public static new GetResourceRelatedJobInvokeArgs Empty => new GetResourceRelatedJobInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourceRelatedJobResult
    {
        public readonly int? DescByJobConfigCreateTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetResourceRelatedJobRefJobInfoResult> RefJobInfos;
        public readonly int? ResourceConfigVersion;
        public readonly string ResourceId;
        public readonly string? ResultOutputFile;
        public readonly string? WorkSpaceId;

        [OutputConstructor]
        private GetResourceRelatedJobResult(
            int? descByJobConfigCreateTime,

            string id,

            ImmutableArray<Outputs.GetResourceRelatedJobRefJobInfoResult> refJobInfos,

            int? resourceConfigVersion,

            string resourceId,

            string? resultOutputFile,

            string? workSpaceId)
        {
            DescByJobConfigCreateTime = descByJobConfigCreateTime;
            Id = id;
            RefJobInfos = refJobInfos;
            ResourceConfigVersion = resourceConfigVersion;
            ResourceId = resourceId;
            ResultOutputFile = resultOutputFile;
            WorkSpaceId = workSpaceId;
        }
    }
}
