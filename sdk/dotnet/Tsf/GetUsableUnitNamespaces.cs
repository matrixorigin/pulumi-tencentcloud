// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tsf
{
    public static class GetUsableUnitNamespaces
    {
        public static Task<GetUsableUnitNamespacesResult> InvokeAsync(GetUsableUnitNamespacesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUsableUnitNamespacesResult>("tencentcloud:Tsf/getUsableUnitNamespaces:getUsableUnitNamespaces", args ?? new GetUsableUnitNamespacesArgs(), options.WithDefaults());

        public static Output<GetUsableUnitNamespacesResult> Invoke(GetUsableUnitNamespacesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUsableUnitNamespacesResult>("tencentcloud:Tsf/getUsableUnitNamespaces:getUsableUnitNamespaces", args ?? new GetUsableUnitNamespacesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUsableUnitNamespacesArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("searchWord")]
        public string? SearchWord { get; set; }

        public GetUsableUnitNamespacesArgs()
        {
        }
        public static new GetUsableUnitNamespacesArgs Empty => new GetUsableUnitNamespacesArgs();
    }

    public sealed class GetUsableUnitNamespacesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("searchWord")]
        public Input<string>? SearchWord { get; set; }

        public GetUsableUnitNamespacesInvokeArgs()
        {
        }
        public static new GetUsableUnitNamespacesInvokeArgs Empty => new GetUsableUnitNamespacesInvokeArgs();
    }


    [OutputType]
    public sealed class GetUsableUnitNamespacesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<Outputs.GetUsableUnitNamespacesResultResult> Results;
        public readonly string? SearchWord;

        [OutputConstructor]
        private GetUsableUnitNamespacesResult(
            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetUsableUnitNamespacesResultResult> results,

            string? searchWord)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            Results = results;
            SearchWord = searchWord;
        }
    }
}