// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Css
{
    public static class GetPullStreamTaskStatus
    {
        public static Task<GetPullStreamTaskStatusResult> InvokeAsync(GetPullStreamTaskStatusArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPullStreamTaskStatusResult>("tencentcloud:Css/getPullStreamTaskStatus:getPullStreamTaskStatus", args ?? new GetPullStreamTaskStatusArgs(), options.WithDefaults());

        public static Output<GetPullStreamTaskStatusResult> Invoke(GetPullStreamTaskStatusInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPullStreamTaskStatusResult>("tencentcloud:Css/getPullStreamTaskStatus:getPullStreamTaskStatus", args ?? new GetPullStreamTaskStatusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPullStreamTaskStatusArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("taskId", required: true)]
        public string TaskId { get; set; } = null!;

        public GetPullStreamTaskStatusArgs()
        {
        }
        public static new GetPullStreamTaskStatusArgs Empty => new GetPullStreamTaskStatusArgs();
    }

    public sealed class GetPullStreamTaskStatusInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("taskId", required: true)]
        public Input<string> TaskId { get; set; } = null!;

        public GetPullStreamTaskStatusInvokeArgs()
        {
        }
        public static new GetPullStreamTaskStatusInvokeArgs Empty => new GetPullStreamTaskStatusInvokeArgs();
    }


    [OutputType]
    public sealed class GetPullStreamTaskStatusResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly string TaskId;
        public readonly ImmutableArray<Outputs.GetPullStreamTaskStatusTaskStatusInfoResult> TaskStatusInfos;

        [OutputConstructor]
        private GetPullStreamTaskStatusResult(
            string id,

            string? resultOutputFile,

            string taskId,

            ImmutableArray<Outputs.GetPullStreamTaskStatusTaskStatusInfoResult> taskStatusInfos)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            TaskId = taskId;
            TaskStatusInfos = taskStatusInfos;
        }
    }
}