// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcr.Inputs
{

    public sealed class InstanceReplicationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("regionId")]
        public Input<int>? RegionId { get; set; }

        [Input("regionName")]
        public Input<string>? RegionName { get; set; }

        [Input("synTag")]
        public Input<bool>? SynTag { get; set; }

        public InstanceReplicationGetArgs()
        {
        }
        public static new InstanceReplicationGetArgs Empty => new InstanceReplicationGetArgs();
    }
}
