// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Teo.Inputs
{

    public sealed class OriginGroupReferenceGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        public OriginGroupReferenceGetArgs()
        {
        }
        public static new OriginGroupReferenceGetArgs Empty => new OriginGroupReferenceGetArgs();
    }
}
