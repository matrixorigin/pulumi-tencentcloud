// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssm
{
    public static class GetRotationHistory
    {
        public static Task<GetRotationHistoryResult> InvokeAsync(GetRotationHistoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRotationHistoryResult>("tencentcloud:Ssm/getRotationHistory:getRotationHistory", args ?? new GetRotationHistoryArgs(), options.WithDefaults());

        public static Output<GetRotationHistoryResult> Invoke(GetRotationHistoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRotationHistoryResult>("tencentcloud:Ssm/getRotationHistory:getRotationHistory", args ?? new GetRotationHistoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRotationHistoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("secretName", required: true)]
        public string SecretName { get; set; } = null!;

        public GetRotationHistoryArgs()
        {
        }
        public static new GetRotationHistoryArgs Empty => new GetRotationHistoryArgs();
    }

    public sealed class GetRotationHistoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("secretName", required: true)]
        public Input<string> SecretName { get; set; } = null!;

        public GetRotationHistoryInvokeArgs()
        {
        }
        public static new GetRotationHistoryInvokeArgs Empty => new GetRotationHistoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetRotationHistoryResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly string SecretName;
        public readonly ImmutableArray<string> VersionIds;

        [OutputConstructor]
        private GetRotationHistoryResult(
            string id,

            string? resultOutputFile,

            string secretName,

            ImmutableArray<string> versionIds)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            SecretName = secretName;
            VersionIds = versionIds;
        }
    }
}
