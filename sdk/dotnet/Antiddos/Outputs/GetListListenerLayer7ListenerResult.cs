// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Antiddos.Outputs
{

    [OutputType]
    public sealed class GetListListenerLayer7ListenerResult
    {
        public readonly string Domain;
        public readonly ImmutableArray<Outputs.GetListListenerLayer7ListenerInstanceDetailRuleResult> InstanceDetailRules;
        public readonly ImmutableArray<Outputs.GetListListenerLayer7ListenerInstanceDetailResult> InstanceDetails;
        public readonly string Protocol;
        public readonly ImmutableArray<Outputs.GetListListenerLayer7ListenerProxyTypeListResult> ProxyTypeLists;
        public readonly ImmutableArray<Outputs.GetListListenerLayer7ListenerRealServerResult> RealServers;
        public readonly int Vport;

        [OutputConstructor]
        private GetListListenerLayer7ListenerResult(
            string domain,

            ImmutableArray<Outputs.GetListListenerLayer7ListenerInstanceDetailRuleResult> instanceDetailRules,

            ImmutableArray<Outputs.GetListListenerLayer7ListenerInstanceDetailResult> instanceDetails,

            string protocol,

            ImmutableArray<Outputs.GetListListenerLayer7ListenerProxyTypeListResult> proxyTypeLists,

            ImmutableArray<Outputs.GetListListenerLayer7ListenerRealServerResult> realServers,

            int vport)
        {
            Domain = domain;
            InstanceDetailRules = instanceDetailRules;
            InstanceDetails = instanceDetails;
            Protocol = protocol;
            ProxyTypeLists = proxyTypeLists;
            RealServers = realServers;
            Vport = vport;
        }
    }
}
