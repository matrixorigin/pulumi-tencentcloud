// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcm.Outputs
{

    [OutputType]
    public sealed class MeshConfigPrometheus
    {
        public readonly Outputs.MeshConfigPrometheusCustomProm? CustomProm;
        public readonly string? InstanceId;
        public readonly string? Region;
        public readonly string? SubnetId;
        public readonly string? VpcId;

        [OutputConstructor]
        private MeshConfigPrometheus(
            Outputs.MeshConfigPrometheusCustomProm? customProm,

            string? instanceId,

            string? region,

            string? subnetId,

            string? vpcId)
        {
            CustomProm = customProm;
            InstanceId = instanceId;
            Region = region;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}