// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cam
{
    public static class GetSecretLastUsedTime
    {
        public static Task<GetSecretLastUsedTimeResult> InvokeAsync(GetSecretLastUsedTimeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecretLastUsedTimeResult>("tencentcloud:Cam/getSecretLastUsedTime:getSecretLastUsedTime", args ?? new GetSecretLastUsedTimeArgs(), options.WithDefaults());

        public static Output<GetSecretLastUsedTimeResult> Invoke(GetSecretLastUsedTimeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretLastUsedTimeResult>("tencentcloud:Cam/getSecretLastUsedTime:getSecretLastUsedTime", args ?? new GetSecretLastUsedTimeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretLastUsedTimeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("secretIdLists", required: true)]
        private List<string>? _secretIdLists;
        public List<string> SecretIdLists
        {
            get => _secretIdLists ?? (_secretIdLists = new List<string>());
            set => _secretIdLists = value;
        }

        public GetSecretLastUsedTimeArgs()
        {
        }
        public static new GetSecretLastUsedTimeArgs Empty => new GetSecretLastUsedTimeArgs();
    }

    public sealed class GetSecretLastUsedTimeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("secretIdLists", required: true)]
        private InputList<string>? _secretIdLists;
        public InputList<string> SecretIdLists
        {
            get => _secretIdLists ?? (_secretIdLists = new InputList<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableArray.Create<string>());
                _secretIdLists = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        public GetSecretLastUsedTimeInvokeArgs()
        {
        }
        public static new GetSecretLastUsedTimeInvokeArgs Empty => new GetSecretLastUsedTimeInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecretLastUsedTimeResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<Outputs.GetSecretLastUsedTimeSecretIdLastUsedRowResult> SecretIdLastUsedRows;
        public readonly ImmutableArray<string> SecretIdLists;

        [OutputConstructor]
        private GetSecretLastUsedTimeResult(
            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetSecretLastUsedTimeSecretIdLastUsedRowResult> secretIdLastUsedRows,

            ImmutableArray<string> secretIdLists)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            SecretIdLastUsedRows = secretIdLastUsedRows;
            SecretIdLists = secretIdLists;
        }
    }
}