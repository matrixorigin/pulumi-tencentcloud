// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Csip.Inputs
{

    public sealed class RiskCenterTaskAdvanceCfgWeakPwdRiskGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("checkItemId", required: true)]
        public Input<int> CheckItemId { get; set; } = null!;

        [Input("enable", required: true)]
        public Input<int> Enable { get; set; } = null!;

        public RiskCenterTaskAdvanceCfgWeakPwdRiskGetArgs()
        {
        }
        public static new RiskCenterTaskAdvanceCfgWeakPwdRiskGetArgs Empty => new RiskCenterTaskAdvanceCfgWeakPwdRiskGetArgs();
    }
}
