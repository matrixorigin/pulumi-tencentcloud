// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cvm.Inputs
{

    public sealed class LaunchTemplateVirtualPrivateCloudGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("asVpcGateway")]
        public Input<bool>? AsVpcGateway { get; set; }

        [Input("ipv6AddressCount")]
        public Input<int>? Ipv6AddressCount { get; set; }

        [Input("privateIpAddresses")]
        private InputList<string>? _privateIpAddresses;
        public InputList<string> PrivateIpAddresses
        {
            get => _privateIpAddresses ?? (_privateIpAddresses = new InputList<string>());
            set => _privateIpAddresses = value;
        }

        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public LaunchTemplateVirtualPrivateCloudGetArgs()
        {
        }
        public static new LaunchTemplateVirtualPrivateCloudGetArgs Empty => new LaunchTemplateVirtualPrivateCloudGetArgs();
    }
}