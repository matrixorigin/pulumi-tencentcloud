// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ccn.Inputs
{

    public sealed class RouteTableBroadcastPoliciesPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        [Input("broadcastConditions", required: true)]
        private InputList<Inputs.RouteTableBroadcastPoliciesPolicyBroadcastConditionGetArgs>? _broadcastConditions;
        public InputList<Inputs.RouteTableBroadcastPoliciesPolicyBroadcastConditionGetArgs> BroadcastConditions
        {
            get => _broadcastConditions ?? (_broadcastConditions = new InputList<Inputs.RouteTableBroadcastPoliciesPolicyBroadcastConditionGetArgs>());
            set => _broadcastConditions = value;
        }

        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("routeConditions", required: true)]
        private InputList<Inputs.RouteTableBroadcastPoliciesPolicyRouteConditionGetArgs>? _routeConditions;
        public InputList<Inputs.RouteTableBroadcastPoliciesPolicyRouteConditionGetArgs> RouteConditions
        {
            get => _routeConditions ?? (_routeConditions = new InputList<Inputs.RouteTableBroadcastPoliciesPolicyRouteConditionGetArgs>());
            set => _routeConditions = value;
        }

        public RouteTableBroadcastPoliciesPolicyGetArgs()
        {
        }
        public static new RouteTableBroadcastPoliciesPolicyGetArgs Empty => new RouteTableBroadcastPoliciesPolicyGetArgs();
    }
}
