// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clb
{
    public static class GetInstanceTraffic
    {
        public static Task<GetInstanceTrafficResult> InvokeAsync(GetInstanceTrafficArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceTrafficResult>("tencentcloud:Clb/getInstanceTraffic:getInstanceTraffic", args ?? new GetInstanceTrafficArgs(), options.WithDefaults());

        public static Output<GetInstanceTrafficResult> Invoke(GetInstanceTrafficInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceTrafficResult>("tencentcloud:Clb/getInstanceTraffic:getInstanceTraffic", args ?? new GetInstanceTrafficInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceTrafficArgs : global::Pulumi.InvokeArgs
    {
        [Input("loadBalancerRegion")]
        public string? LoadBalancerRegion { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstanceTrafficArgs()
        {
        }
        public static new GetInstanceTrafficArgs Empty => new GetInstanceTrafficArgs();
    }

    public sealed class GetInstanceTrafficInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("loadBalancerRegion")]
        public Input<string>? LoadBalancerRegion { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstanceTrafficInvokeArgs()
        {
        }
        public static new GetInstanceTrafficInvokeArgs Empty => new GetInstanceTrafficInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceTrafficResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? LoadBalancerRegion;
        public readonly ImmutableArray<Outputs.GetInstanceTrafficLoadBalancerTrafficResult> LoadBalancerTraffics;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstanceTrafficResult(
            string id,

            string? loadBalancerRegion,

            ImmutableArray<Outputs.GetInstanceTrafficLoadBalancerTrafficResult> loadBalancerTraffics,

            string? resultOutputFile)
        {
            Id = id;
            LoadBalancerRegion = loadBalancerRegion;
            LoadBalancerTraffics = loadBalancerTraffics;
            ResultOutputFile = resultOutputFile;
        }
    }
}