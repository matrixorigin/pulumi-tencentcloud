// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clickhouse
{
    public static class GetInstanceShards
    {
        public static Task<GetInstanceShardsResult> InvokeAsync(GetInstanceShardsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceShardsResult>("tencentcloud:Clickhouse/getInstanceShards:getInstanceShards", args ?? new GetInstanceShardsArgs(), options.WithDefaults());

        public static Output<GetInstanceShardsResult> Invoke(GetInstanceShardsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceShardsResult>("tencentcloud:Clickhouse/getInstanceShards:getInstanceShards", args ?? new GetInstanceShardsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceShardsArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstanceShardsArgs()
        {
        }
        public static new GetInstanceShardsArgs Empty => new GetInstanceShardsArgs();
    }

    public sealed class GetInstanceShardsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstanceShardsInvokeArgs()
        {
        }
        public static new GetInstanceShardsInvokeArgs Empty => new GetInstanceShardsInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceShardsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string InstanceShardsList;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstanceShardsResult(
            string id,

            string instanceId,

            string instanceShardsList,

            string? resultOutputFile)
        {
            Id = id;
            InstanceId = instanceId;
            InstanceShardsList = instanceShardsList;
            ResultOutputFile = resultOutputFile;
        }
    }
}