// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Emr.Inputs
{

    public sealed class ClusterPlacementInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ClusterPlacementInfoArgs()
        {
        }
        public static new ClusterPlacementInfoArgs Empty => new ClusterPlacementInfoArgs();
    }
}
