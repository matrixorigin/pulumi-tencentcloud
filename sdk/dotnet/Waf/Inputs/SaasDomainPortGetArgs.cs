// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Waf.Inputs
{

    public sealed class SaasDomainPortGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("nginxServerId")]
        public Input<string>? NginxServerId { get; set; }

        [Input("port", required: true)]
        public Input<string> Port { get; set; } = null!;

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("upstreamPort", required: true)]
        public Input<string> UpstreamPort { get; set; } = null!;

        [Input("upstreamProtocol", required: true)]
        public Input<string> UpstreamProtocol { get; set; } = null!;

        public SaasDomainPortGetArgs()
        {
        }
        public static new SaasDomainPortGetArgs Empty => new SaasDomainPortGetArgs();
    }
}
