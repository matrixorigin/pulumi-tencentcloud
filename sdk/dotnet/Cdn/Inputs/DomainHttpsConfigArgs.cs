// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdn.Inputs
{

    public sealed class DomainHttpsConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("clientCertificateConfig")]
        public Input<Inputs.DomainHttpsConfigClientCertificateConfigArgs>? ClientCertificateConfig { get; set; }

        [Input("forceRedirect")]
        public Input<Inputs.DomainHttpsConfigForceRedirectArgs>? ForceRedirect { get; set; }

        [Input("http2Switch")]
        public Input<string>? Http2Switch { get; set; }

        [Input("httpsSwitch", required: true)]
        public Input<string> HttpsSwitch { get; set; } = null!;

        [Input("ocspStaplingSwitch")]
        public Input<string>? OcspStaplingSwitch { get; set; }

        [Input("serverCertificateConfig")]
        public Input<Inputs.DomainHttpsConfigServerCertificateConfigArgs>? ServerCertificateConfig { get; set; }

        [Input("spdySwitch")]
        public Input<string>? SpdySwitch { get; set; }

        [Input("tlsVersions")]
        private InputList<string>? _tlsVersions;
        public InputList<string> TlsVersions
        {
            get => _tlsVersions ?? (_tlsVersions = new InputList<string>());
            set => _tlsVersions = value;
        }

        [Input("verifyClient")]
        public Input<string>? VerifyClient { get; set; }

        public DomainHttpsConfigArgs()
        {
        }
        public static new DomainHttpsConfigArgs Empty => new DomainHttpsConfigArgs();
    }
}