// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Thpc.Outputs
{

    [OutputType]
    public sealed class WorkspacesSpaceChargePrepaid
    {
        public readonly int? Period;
        public readonly string? RenewFlag;

        [OutputConstructor]
        private WorkspacesSpaceChargePrepaid(
            int? period,

            string? renewFlag)
        {
            Period = period;
            RenewFlag = renewFlag;
        }
    }
}