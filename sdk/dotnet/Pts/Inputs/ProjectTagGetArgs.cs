// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Pts.Inputs
{

    public sealed class ProjectTagGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("tagKey")]
        public Input<string>? TagKey { get; set; }

        [Input("tagValue")]
        public Input<string>? TagValue { get; set; }

        public ProjectTagGetArgs()
        {
        }
        public static new ProjectTagGetArgs Empty => new ProjectTagGetArgs();
    }
}