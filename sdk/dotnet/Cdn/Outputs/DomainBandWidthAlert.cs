// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdn.Outputs
{

    [OutputType]
    public sealed class DomainBandWidthAlert
    {
        public readonly int? AlertPercentage;
        public readonly string? AlertSwitch;
        public readonly int? BpsThreshold;
        public readonly string? CounterMeasure;
        public readonly string? LastTriggerTime;
        public readonly string? LastTriggerTimeOverseas;
        public readonly string? Metric;
        public readonly Outputs.DomainBandWidthAlertStatisticItem? StatisticItem;
        public readonly string Switch;

        [OutputConstructor]
        private DomainBandWidthAlert(
            int? alertPercentage,

            string? alertSwitch,

            int? bpsThreshold,

            string? counterMeasure,

            string? lastTriggerTime,

            string? lastTriggerTimeOverseas,

            string? metric,

            Outputs.DomainBandWidthAlertStatisticItem? statisticItem,

            string @switch)
        {
            AlertPercentage = alertPercentage;
            AlertSwitch = alertSwitch;
            BpsThreshold = bpsThreshold;
            CounterMeasure = counterMeasure;
            LastTriggerTime = lastTriggerTime;
            LastTriggerTimeOverseas = lastTriggerTimeOverseas;
            Metric = metric;
            StatisticItem = statisticItem;
            Switch = @switch;
        }
    }
}