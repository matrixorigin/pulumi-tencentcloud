// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ccn.Outputs
{

    [OutputType]
    public sealed class RouteTableBroadcastPoliciesPolicy
    {
        public readonly string Action;
        public readonly ImmutableArray<Outputs.RouteTableBroadcastPoliciesPolicyBroadcastCondition> BroadcastConditions;
        public readonly string Description;
        public readonly ImmutableArray<Outputs.RouteTableBroadcastPoliciesPolicyRouteCondition> RouteConditions;

        [OutputConstructor]
        private RouteTableBroadcastPoliciesPolicy(
            string action,

            ImmutableArray<Outputs.RouteTableBroadcastPoliciesPolicyBroadcastCondition> broadcastConditions,

            string description,

            ImmutableArray<Outputs.RouteTableBroadcastPoliciesPolicyRouteCondition> routeConditions)
        {
            Action = action;
            BroadcastConditions = broadcastConditions;
            Description = description;
            RouteConditions = routeConditions;
        }
    }
}