// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tem.Inputs
{

    public sealed class GatewayIngressTlArgs : global::Pulumi.ResourceArgs
    {
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        [Input("hosts", required: true)]
        private InputList<string>? _hosts;
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        [Input("secretName")]
        public Input<string>? SecretName { get; set; }

        public GatewayIngressTlArgs()
        {
        }
        public static new GatewayIngressTlArgs Empty => new GatewayIngressTlArgs();
    }
}