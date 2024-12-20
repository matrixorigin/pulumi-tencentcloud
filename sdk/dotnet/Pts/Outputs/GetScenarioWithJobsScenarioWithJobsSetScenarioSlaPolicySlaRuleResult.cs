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
    public sealed class GetScenarioWithJobsScenarioWithJobsSetScenarioSlaPolicySlaRuleResult
    {
        public readonly bool AbortFlag;
        public readonly string Aggregation;
        public readonly string Condition;
        public readonly string For;
        public readonly ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetScenarioSlaPolicySlaRuleLabelFilterResult> LabelFilters;
        public readonly string Metric;
        public readonly double Value;

        [OutputConstructor]
        private GetScenarioWithJobsScenarioWithJobsSetScenarioSlaPolicySlaRuleResult(
            bool abortFlag,

            string aggregation,

            string condition,

            string @for,

            ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetScenarioSlaPolicySlaRuleLabelFilterResult> labelFilters,

            string metric,

            double value)
        {
            AbortFlag = abortFlag;
            Aggregation = aggregation;
            Condition = condition;
            For = @for;
            LabelFilters = labelFilters;
            Metric = metric;
            Value = value;
        }
    }
}
