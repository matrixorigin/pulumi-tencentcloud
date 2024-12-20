// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Gaap
{
    public static class GetAccessRegions
    {
        public static Task<GetAccessRegionsResult> InvokeAsync(GetAccessRegionsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessRegionsResult>("tencentcloud:Gaap/getAccessRegions:getAccessRegions", args ?? new GetAccessRegionsArgs(), options.WithDefaults());

        public static Output<GetAccessRegionsResult> Invoke(GetAccessRegionsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessRegionsResult>("tencentcloud:Gaap/getAccessRegions:getAccessRegions", args ?? new GetAccessRegionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessRegionsArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetAccessRegionsArgs()
        {
        }
        public static new GetAccessRegionsArgs Empty => new GetAccessRegionsArgs();
    }

    public sealed class GetAccessRegionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetAccessRegionsInvokeArgs()
        {
        }
        public static new GetAccessRegionsInvokeArgs Empty => new GetAccessRegionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccessRegionsResult
    {
        public readonly ImmutableArray<Outputs.GetAccessRegionsAccessRegionSetResult> AccessRegionSets;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetAccessRegionsResult(
            ImmutableArray<Outputs.GetAccessRegionsAccessRegionSetResult> accessRegionSets,

            string id,

            string? resultOutputFile)
        {
            AccessRegionSets = accessRegionSets;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
