// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clb.Outputs
{

    [OutputType]
    public sealed class GetTargetHealthLoadBalancerResult
    {
        public readonly ImmutableArray<Outputs.GetTargetHealthLoadBalancerListenerResult> Listeners;
        public readonly string LoadBalancerId;
        public readonly string LoadBalancerName;

        [OutputConstructor]
        private GetTargetHealthLoadBalancerResult(
            ImmutableArray<Outputs.GetTargetHealthLoadBalancerListenerResult> listeners,

            string loadBalancerId,

            string loadBalancerName)
        {
            Listeners = listeners;
            LoadBalancerId = loadBalancerId;
            LoadBalancerName = loadBalancerName;
        }
    }
}