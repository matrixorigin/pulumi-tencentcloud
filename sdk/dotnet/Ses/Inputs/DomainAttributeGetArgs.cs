// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ses.Inputs
{

    public sealed class DomainAttributeGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("expectedValue")]
        public Input<string>? ExpectedValue { get; set; }

        [Input("sendDomain")]
        public Input<string>? SendDomain { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public DomainAttributeGetArgs()
        {
        }
        public static new DomainAttributeGetArgs Empty => new DomainAttributeGetArgs();
    }
}