// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mongodb.Inputs
{

    public sealed class InstanceAccountAuthRoleGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("mask", required: true)]
        public Input<int> Mask { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public InstanceAccountAuthRoleGetArgs()
        {
        }
        public static new InstanceAccountAuthRoleGetArgs Empty => new InstanceAccountAuthRoleGetArgs();
    }
}