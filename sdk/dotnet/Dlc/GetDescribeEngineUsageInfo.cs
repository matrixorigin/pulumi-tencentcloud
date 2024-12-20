// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dlc
{
    public static class GetDescribeEngineUsageInfo
    {
        public static Task<GetDescribeEngineUsageInfoResult> InvokeAsync(GetDescribeEngineUsageInfoArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDescribeEngineUsageInfoResult>("tencentcloud:Dlc/getDescribeEngineUsageInfo:getDescribeEngineUsageInfo", args ?? new GetDescribeEngineUsageInfoArgs(), options.WithDefaults());

        public static Output<GetDescribeEngineUsageInfoResult> Invoke(GetDescribeEngineUsageInfoInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDescribeEngineUsageInfoResult>("tencentcloud:Dlc/getDescribeEngineUsageInfo:getDescribeEngineUsageInfo", args ?? new GetDescribeEngineUsageInfoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeEngineUsageInfoArgs : global::Pulumi.InvokeArgs
    {
        [Input("dataEngineId", required: true)]
        public string DataEngineId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDescribeEngineUsageInfoArgs()
        {
        }
        public static new GetDescribeEngineUsageInfoArgs Empty => new GetDescribeEngineUsageInfoArgs();
    }

    public sealed class GetDescribeEngineUsageInfoInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("dataEngineId", required: true)]
        public Input<string> DataEngineId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDescribeEngineUsageInfoInvokeArgs()
        {
        }
        public static new GetDescribeEngineUsageInfoInvokeArgs Empty => new GetDescribeEngineUsageInfoInvokeArgs();
    }


    [OutputType]
    public sealed class GetDescribeEngineUsageInfoResult
    {
        public readonly int Available;
        public readonly string DataEngineId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly int Used;

        [OutputConstructor]
        private GetDescribeEngineUsageInfoResult(
            int available,

            string dataEngineId,

            string id,

            string? resultOutputFile,

            int used)
        {
            Available = available;
            DataEngineId = dataEngineId;
            Id = id;
            ResultOutputFile = resultOutputFile;
            Used = used;
        }
    }
}
