// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Postgresql
{
    public static class GetXlogs
    {
        public static Task<GetXlogsResult> InvokeAsync(GetXlogsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetXlogsResult>("tencentcloud:Postgresql/getXlogs:getXlogs", args ?? new GetXlogsArgs(), options.WithDefaults());

        public static Output<GetXlogsResult> Invoke(GetXlogsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetXlogsResult>("tencentcloud:Postgresql/getXlogs:getXlogs", args ?? new GetXlogsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetXlogsArgs : global::Pulumi.InvokeArgs
    {
        [Input("endTime")]
        public string? EndTime { get; set; }

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("startTime")]
        public string? StartTime { get; set; }

        public GetXlogsArgs()
        {
        }
        public static new GetXlogsArgs Empty => new GetXlogsArgs();
    }

    public sealed class GetXlogsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public GetXlogsInvokeArgs()
        {
        }
        public static new GetXlogsInvokeArgs Empty => new GetXlogsInvokeArgs();
    }


    [OutputType]
    public sealed class GetXlogsResult
    {
        public readonly string? EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly ImmutableArray<Outputs.GetXlogsListResult> Lists;
        public readonly string? ResultOutputFile;
        public readonly string? StartTime;

        [OutputConstructor]
        private GetXlogsResult(
            string? endTime,

            string id,

            string instanceId,

            ImmutableArray<Outputs.GetXlogsListResult> lists,

            string? resultOutputFile,

            string? startTime)
        {
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            Lists = lists;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
        }
    }
}