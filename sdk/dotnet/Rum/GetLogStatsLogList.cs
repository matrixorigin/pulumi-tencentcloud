// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Rum
{
    public static class GetLogStatsLogList
    {
        public static Task<GetLogStatsLogListResult> InvokeAsync(GetLogStatsLogListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLogStatsLogListResult>("tencentcloud:Rum/getLogStatsLogList:getLogStatsLogList", args ?? new GetLogStatsLogListArgs(), options.WithDefaults());

        public static Output<GetLogStatsLogListResult> Invoke(GetLogStatsLogListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLogStatsLogListResult>("tencentcloud:Rum/getLogStatsLogList:getLogStatsLogList", args ?? new GetLogStatsLogListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLogStatsLogListArgs : global::Pulumi.InvokeArgs
    {
        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        [Input("projectId", required: true)]
        public int ProjectId { get; set; }

        [Input("query", required: true)]
        public string Query { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        public GetLogStatsLogListArgs()
        {
        }
        public static new GetLogStatsLogListArgs Empty => new GetLogStatsLogListArgs();
    }

    public sealed class GetLogStatsLogListInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        [Input("query", required: true)]
        public Input<string> Query { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public GetLogStatsLogListInvokeArgs()
        {
        }
        public static new GetLogStatsLogListInvokeArgs Empty => new GetLogStatsLogListInvokeArgs();
    }


    [OutputType]
    public sealed class GetLogStatsLogListResult
    {
        public readonly string EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int ProjectId;
        public readonly string Query;
        public readonly string Result;
        public readonly string? ResultOutputFile;
        public readonly string StartTime;

        [OutputConstructor]
        private GetLogStatsLogListResult(
            string endTime,

            string id,

            int projectId,

            string query,

            string result,

            string? resultOutputFile,

            string startTime)
        {
            EndTime = endTime;
            Id = id;
            ProjectId = projectId;
            Query = query;
            Result = result;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
        }
    }
}