// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.As.Outputs
{

    [OutputType]
    public sealed class LoadBalancerForwardLoadBalancer
    {
        public readonly string ListenerId;
        public readonly string LoadBalancerId;
        public readonly string? LocationId;
        public readonly string? Region;
        public readonly ImmutableArray<Outputs.LoadBalancerForwardLoadBalancerTargetAttribute> TargetAttributes;

        [OutputConstructor]
        private LoadBalancerForwardLoadBalancer(
            string listenerId,

            string loadBalancerId,

            string? locationId,

            string? region,

            ImmutableArray<Outputs.LoadBalancerForwardLoadBalancerTargetAttribute> targetAttributes)
        {
            ListenerId = listenerId;
            LoadBalancerId = loadBalancerId;
            LocationId = locationId;
            Region = region;
            TargetAttributes = targetAttributes;
        }
    }
}