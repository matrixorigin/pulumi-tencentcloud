// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class FlowInputGroupSrtSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("latency")]
        public Input<int>? Latency { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("passphrase")]
        public Input<string>? Passphrase { get; set; }

        [Input("pbKeyLen")]
        public Input<int>? PbKeyLen { get; set; }

        [Input("peerIdleTimeout")]
        public Input<int>? PeerIdleTimeout { get; set; }

        [Input("peerLatency")]
        public Input<int>? PeerLatency { get; set; }

        [Input("recvLatency")]
        public Input<int>? RecvLatency { get; set; }

        [Input("sourceAddresses")]
        private InputList<Inputs.FlowInputGroupSrtSettingsSourceAddressGetArgs>? _sourceAddresses;
        public InputList<Inputs.FlowInputGroupSrtSettingsSourceAddressGetArgs> SourceAddresses
        {
            get => _sourceAddresses ?? (_sourceAddresses = new InputList<Inputs.FlowInputGroupSrtSettingsSourceAddressGetArgs>());
            set => _sourceAddresses = value;
        }

        [Input("streamId")]
        public Input<string>? StreamId { get; set; }

        public FlowInputGroupSrtSettingsGetArgs()
        {
        }
        public static new FlowInputGroupSrtSettingsGetArgs Empty => new FlowInputGroupSrtSettingsGetArgs();
    }
}