// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Antiddos.Inputs
{

    public sealed class DdosSpeedLimitConfigDdosSpeedLimitConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dstPortList")]
        public Input<string>? DstPortList { get; set; }

        [Input("dstPortScopes")]
        private InputList<Inputs.DdosSpeedLimitConfigDdosSpeedLimitConfigDstPortScopeGetArgs>? _dstPortScopes;
        public InputList<Inputs.DdosSpeedLimitConfigDdosSpeedLimitConfigDstPortScopeGetArgs> DstPortScopes
        {
            get => _dstPortScopes ?? (_dstPortScopes = new InputList<Inputs.DdosSpeedLimitConfigDdosSpeedLimitConfigDstPortScopeGetArgs>());
            set => _dstPortScopes = value;
        }

        [Input("mode", required: true)]
        public Input<int> Mode { get; set; } = null!;

        [Input("protocolList")]
        public Input<string>? ProtocolList { get; set; }

        [Input("speedValues", required: true)]
        private InputList<Inputs.DdosSpeedLimitConfigDdosSpeedLimitConfigSpeedValueGetArgs>? _speedValues;
        public InputList<Inputs.DdosSpeedLimitConfigDdosSpeedLimitConfigSpeedValueGetArgs> SpeedValues
        {
            get => _speedValues ?? (_speedValues = new InputList<Inputs.DdosSpeedLimitConfigDdosSpeedLimitConfigSpeedValueGetArgs>());
            set => _speedValues = value;
        }

        public DdosSpeedLimitConfigDdosSpeedLimitConfigGetArgs()
        {
        }
        public static new DdosSpeedLimitConfigDdosSpeedLimitConfigGetArgs Empty => new DdosSpeedLimitConfigDdosSpeedLimitConfigGetArgs();
    }
}
