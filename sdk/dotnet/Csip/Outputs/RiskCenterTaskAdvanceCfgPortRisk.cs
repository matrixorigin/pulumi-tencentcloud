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
    public sealed class RiskCenterTaskAdvanceCfgPortRisk
    {
        public readonly int CheckType;
        public readonly string Detail;
        public readonly int Enable;
        public readonly string PortSets;

        [OutputConstructor]
        private RiskCenterTaskAdvanceCfgPortRisk(
            int checkType,

            string detail,

            int enable,

            string portSets)
        {
            CheckType = checkType;
            Detail = detail;
            Enable = enable;
            PortSets = portSets;
        }
    }
}
