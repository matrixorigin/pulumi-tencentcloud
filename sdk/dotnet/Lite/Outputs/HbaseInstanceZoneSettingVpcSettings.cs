// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Lite.Outputs
{

    [OutputType]
    public sealed class HbaseInstanceZoneSettingVpcSettings
    {
        public readonly string SubnetId;
        public readonly string VpcId;

        [OutputConstructor]
        private HbaseInstanceZoneSettingVpcSettings(
            string subnetId,

            string vpcId)
        {
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}
