// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdwpg.Inputs
{

    public sealed class InstanceResourceDiskSpecArgs : global::Pulumi.ResourceArgs
    {
        [Input("diskCount", required: true)]
        public Input<int> DiskCount { get; set; } = null!;

        [Input("diskSize", required: true)]
        public Input<int> DiskSize { get; set; } = null!;

        [Input("diskType", required: true)]
        public Input<string> DiskType { get; set; } = null!;

        public InstanceResourceDiskSpecArgs()
        {
        }
        public static new InstanceResourceDiskSpecArgs Empty => new InstanceResourceDiskSpecArgs();
    }
}