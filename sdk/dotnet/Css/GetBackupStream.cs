// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Css
{
    public static class GetBackupStream
    {
        public static Task<GetBackupStreamResult> InvokeAsync(GetBackupStreamArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBackupStreamResult>("tencentcloud:Css/getBackupStream:getBackupStream", args ?? new GetBackupStreamArgs(), options.WithDefaults());

        public static Output<GetBackupStreamResult> Invoke(GetBackupStreamInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBackupStreamResult>("tencentcloud:Css/getBackupStream:getBackupStream", args ?? new GetBackupStreamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupStreamArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("streamName")]
        public string? StreamName { get; set; }

        public GetBackupStreamArgs()
        {
        }
        public static new GetBackupStreamArgs Empty => new GetBackupStreamArgs();
    }

    public sealed class GetBackupStreamInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("streamName")]
        public Input<string>? StreamName { get; set; }

        public GetBackupStreamInvokeArgs()
        {
        }
        public static new GetBackupStreamInvokeArgs Empty => new GetBackupStreamInvokeArgs();
    }


    [OutputType]
    public sealed class GetBackupStreamResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<Outputs.GetBackupStreamStreamInfoListResult> StreamInfoLists;
        public readonly string? StreamName;

        [OutputConstructor]
        private GetBackupStreamResult(
            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetBackupStreamStreamInfoListResult> streamInfoLists,

            string? streamName)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            StreamInfoLists = streamInfoLists;
            StreamName = streamName;
        }
    }
}
