// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Postgresql.Inputs
{

    public sealed class InstanceDbNodeSetGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dedicatedClusterId")]
        public Input<string>? DedicatedClusterId { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public InstanceDbNodeSetGetArgs()
        {
        }
        public static new InstanceDbNodeSetGetArgs Empty => new InstanceDbNodeSetGetArgs();
    }
}
