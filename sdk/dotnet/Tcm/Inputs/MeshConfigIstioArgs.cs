// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcm.Inputs
{

    public sealed class MeshConfigIstioArgs : global::Pulumi.ResourceArgs
    {
        [Input("disableHttpRetry")]
        public Input<bool>? DisableHttpRetry { get; set; }

        [Input("disablePolicyChecks")]
        public Input<bool>? DisablePolicyChecks { get; set; }

        [Input("enablePilotHttp")]
        public Input<bool>? EnablePilotHttp { get; set; }

        [Input("outboundTrafficPolicy", required: true)]
        public Input<string> OutboundTrafficPolicy { get; set; } = null!;

        [Input("smartDns")]
        public Input<Inputs.MeshConfigIstioSmartDnsArgs>? SmartDns { get; set; }

        [Input("tracing")]
        public Input<Inputs.MeshConfigIstioTracingArgs>? Tracing { get; set; }

        public MeshConfigIstioArgs()
        {
        }
        public static new MeshConfigIstioArgs Empty => new MeshConfigIstioArgs();
    }
}