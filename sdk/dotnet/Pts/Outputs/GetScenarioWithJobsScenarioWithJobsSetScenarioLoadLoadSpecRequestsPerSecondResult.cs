// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Pts.Outputs
{

    [OutputType]
    public sealed class GetScenarioWithJobsScenarioWithJobsSetScenarioLoadLoadSpecRequestsPerSecondResult
    {
        public readonly int DurationSeconds;
        public readonly int GracefulStopSeconds;
        public readonly int MaxRequestsPerSecond;
        public readonly int Resources;
        public readonly int StartRequestsPerSecond;
        public readonly int TargetRequestsPerSecond;
        public readonly int TargetVirtualUsers;

        [OutputConstructor]
        private GetScenarioWithJobsScenarioWithJobsSetScenarioLoadLoadSpecRequestsPerSecondResult(
            int durationSeconds,

            int gracefulStopSeconds,

            int maxRequestsPerSecond,

            int resources,

            int startRequestsPerSecond,

            int targetRequestsPerSecond,

            int targetVirtualUsers)
        {
            DurationSeconds = durationSeconds;
            GracefulStopSeconds = gracefulStopSeconds;
            MaxRequestsPerSecond = maxRequestsPerSecond;
            Resources = resources;
            StartRequestsPerSecond = startRequestsPerSecond;
            TargetRequestsPerSecond = targetRequestsPerSecond;
            TargetVirtualUsers = targetVirtualUsers;
        }
    }
}