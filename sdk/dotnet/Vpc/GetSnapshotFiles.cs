// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vpc
{
    public static class GetSnapshotFiles
    {
        public static Task<GetSnapshotFilesResult> InvokeAsync(GetSnapshotFilesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotFilesResult>("tencentcloud:Vpc/getSnapshotFiles:getSnapshotFiles", args ?? new GetSnapshotFilesArgs(), options.WithDefaults());

        public static Output<GetSnapshotFilesResult> Invoke(GetSnapshotFilesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotFilesResult>("tencentcloud:Vpc/getSnapshotFiles:getSnapshotFiles", args ?? new GetSnapshotFilesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotFilesArgs : global::Pulumi.InvokeArgs
    {
        [Input("businessType", required: true)]
        public string BusinessType { get; set; } = null!;

        [Input("endDate", required: true)]
        public string EndDate { get; set; } = null!;

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("startDate", required: true)]
        public string StartDate { get; set; } = null!;

        public GetSnapshotFilesArgs()
        {
        }
        public static new GetSnapshotFilesArgs Empty => new GetSnapshotFilesArgs();
    }

    public sealed class GetSnapshotFilesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("businessType", required: true)]
        public Input<string> BusinessType { get; set; } = null!;

        [Input("endDate", required: true)]
        public Input<string> EndDate { get; set; } = null!;

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        public GetSnapshotFilesInvokeArgs()
        {
        }
        public static new GetSnapshotFilesInvokeArgs Empty => new GetSnapshotFilesInvokeArgs();
    }


    [OutputType]
    public sealed class GetSnapshotFilesResult
    {
        public readonly string BusinessType;
        public readonly string EndDate;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<Outputs.GetSnapshotFilesSnapshotFileSetResult> SnapshotFileSets;
        public readonly string StartDate;

        [OutputConstructor]
        private GetSnapshotFilesResult(
            string businessType,

            string endDate,

            string id,

            string instanceId,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetSnapshotFilesSnapshotFileSetResult> snapshotFileSets,

            string startDate)
        {
            BusinessType = businessType;
            EndDate = endDate;
            Id = id;
            InstanceId = instanceId;
            ResultOutputFile = resultOutputFile;
            SnapshotFileSets = snapshotFileSets;
            StartDate = startDate;
        }
    }
}