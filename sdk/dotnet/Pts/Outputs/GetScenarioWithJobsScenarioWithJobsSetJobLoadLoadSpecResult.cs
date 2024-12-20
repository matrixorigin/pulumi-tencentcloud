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
    public sealed class GetScenarioWithJobsScenarioWithJobsSetJobLoadLoadSpecResult
    {
        public readonly ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetJobLoadLoadSpecConcurrencyResult> Concurrencies;
        public readonly ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetJobLoadLoadSpecRequestsPerSecondResult> RequestsPerSeconds;
        public readonly ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetJobLoadLoadSpecScriptOriginResult> ScriptOrigins;

        [OutputConstructor]
        private GetScenarioWithJobsScenarioWithJobsSetJobLoadLoadSpecResult(
            ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetJobLoadLoadSpecConcurrencyResult> concurrencies,

            ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetJobLoadLoadSpecRequestsPerSecondResult> requestsPerSeconds,

            ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetJobLoadLoadSpecScriptOriginResult> scriptOrigins)
        {
            Concurrencies = concurrencies;
            RequestsPerSeconds = requestsPerSeconds;
            ScriptOrigins = scriptOrigins;
        }
    }
}
