// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class FlowInputGroupResilientStreamGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("bufferTime")]
        public Input<int>? BufferTime { get; set; }

        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        public FlowInputGroupResilientStreamGetArgs()
        {
        }
        public static new FlowInputGroupResilientStreamGetArgs Empty => new FlowInputGroupResilientStreamGetArgs();
    }
}