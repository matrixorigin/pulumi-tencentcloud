// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Oceanus.Inputs
{

    public sealed class ResourceResourceLocArgs : global::Pulumi.ResourceArgs
    {
        [Input("param", required: true)]
        public Input<Inputs.ResourceResourceLocParamArgs> Param { get; set; } = null!;

        [Input("storageType", required: true)]
        public Input<int> StorageType { get; set; } = null!;

        public ResourceResourceLocArgs()
        {
        }
        public static new ResourceResourceLocArgs Empty => new ResourceResourceLocArgs();
    }
}