// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Sqlserver.Outputs
{

    [OutputType]
    public sealed class GetZoneConfigZoneListResult
    {
        public readonly string AvailabilityZone;
        public readonly ImmutableArray<Outputs.GetZoneConfigZoneListSpecinfoListResult> SpecinfoLists;
        public readonly int ZoneId;

        [OutputConstructor]
        private GetZoneConfigZoneListResult(
            string availabilityZone,

            ImmutableArray<Outputs.GetZoneConfigZoneListSpecinfoListResult> specinfoLists,

            int zoneId)
        {
            AvailabilityZone = availabilityZone;
            SpecinfoLists = specinfoLists;
            ZoneId = zoneId;
        }
    }
}