// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Csip.Outputs
{

    [OutputType]
    public sealed class RiskCenterTaskAdvanceCfg
    {
        public readonly ImmutableArray<Outputs.RiskCenterTaskAdvanceCfgCfgRisk> CfgRisks;
        public readonly ImmutableArray<Outputs.RiskCenterTaskAdvanceCfgPortRisk> PortRisks;
        public readonly ImmutableArray<Outputs.RiskCenterTaskAdvanceCfgVulRisk> VulRisks;
        public readonly ImmutableArray<Outputs.RiskCenterTaskAdvanceCfgWeakPwdRisk> WeakPwdRisks;

        [OutputConstructor]
        private RiskCenterTaskAdvanceCfg(
            ImmutableArray<Outputs.RiskCenterTaskAdvanceCfgCfgRisk> cfgRisks,

            ImmutableArray<Outputs.RiskCenterTaskAdvanceCfgPortRisk> portRisks,

            ImmutableArray<Outputs.RiskCenterTaskAdvanceCfgVulRisk> vulRisks,

            ImmutableArray<Outputs.RiskCenterTaskAdvanceCfgWeakPwdRisk> weakPwdRisks)
        {
            CfgRisks = cfgRisks;
            PortRisks = portRisks;
            VulRisks = vulRisks;
            WeakPwdRisks = weakPwdRisks;
        }
    }
}
