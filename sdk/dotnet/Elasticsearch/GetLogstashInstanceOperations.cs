// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elasticsearch
{
    public static class GetLogstashInstanceOperations
    {
        public static Task<GetLogstashInstanceOperationsResult> InvokeAsync(GetLogstashInstanceOperationsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLogstashInstanceOperationsResult>("tencentcloud:Elasticsearch/getLogstashInstanceOperations:getLogstashInstanceOperations", args ?? new GetLogstashInstanceOperationsArgs(), options.WithDefaults());

        public static Output<GetLogstashInstanceOperationsResult> Invoke(GetLogstashInstanceOperationsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLogstashInstanceOperationsResult>("tencentcloud:Elasticsearch/getLogstashInstanceOperations:getLogstashInstanceOperations", args ?? new GetLogstashInstanceOperationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLogstashInstanceOperationsArgs : global::Pulumi.InvokeArgs
    {
        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        public GetLogstashInstanceOperationsArgs()
        {
        }
        public static new GetLogstashInstanceOperationsArgs Empty => new GetLogstashInstanceOperationsArgs();
    }

    public sealed class GetLogstashInstanceOperationsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public GetLogstashInstanceOperationsInvokeArgs()
        {
        }
        public static new GetLogstashInstanceOperationsInvokeArgs Empty => new GetLogstashInstanceOperationsInvokeArgs();
    }


    [OutputType]
    public sealed class GetLogstashInstanceOperationsResult
    {
        public readonly string EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly ImmutableArray<Outputs.GetLogstashInstanceOperationsOperationResult> Operations;
        public readonly string? ResultOutputFile;
        public readonly string StartTime;

        [OutputConstructor]
        private GetLogstashInstanceOperationsResult(
            string endTime,

            string id,

            string instanceId,

            ImmutableArray<Outputs.GetLogstashInstanceOperationsOperationResult> operations,

            string? resultOutputFile,

            string startTime)
        {
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            Operations = operations;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
        }
    }
}