// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Emr.Inputs
{

    public sealed class ClusterResourceSpecTaskResourceSpecArgs : global::Pulumi.ResourceArgs
    {
        [Input("cpu")]
        public Input<int>? Cpu { get; set; }

        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        [Input("memSize")]
        public Input<int>? MemSize { get; set; }

        [Input("multiDisks")]
        private InputList<Inputs.ClusterResourceSpecTaskResourceSpecMultiDiskArgs>? _multiDisks;
        public InputList<Inputs.ClusterResourceSpecTaskResourceSpecMultiDiskArgs> MultiDisks
        {
            get => _multiDisks ?? (_multiDisks = new InputList<Inputs.ClusterResourceSpecTaskResourceSpecMultiDiskArgs>());
            set => _multiDisks = value;
        }

        [Input("rootSize")]
        public Input<int>? RootSize { get; set; }

        [Input("spec")]
        public Input<string>? Spec { get; set; }

        [Input("storageType")]
        public Input<int>? StorageType { get; set; }

        public ClusterResourceSpecTaskResourceSpecArgs()
        {
        }
        public static new ClusterResourceSpecTaskResourceSpecArgs Empty => new ClusterResourceSpecTaskResourceSpecArgs();
    }
}
