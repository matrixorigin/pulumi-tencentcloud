// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class ScheduleActivityActivityParaSnapshotByTimeOffsetTask
    {
        public readonly int Definition;
        public readonly ImmutableArray<string> ExtTimeOffsetSets;
        public readonly Outputs.ScheduleActivityActivityParaSnapshotByTimeOffsetTaskObjectNumberFormat? ObjectNumberFormat;
        public readonly string? OutputObjectPath;
        public readonly Outputs.ScheduleActivityActivityParaSnapshotByTimeOffsetTaskOutputStorage? OutputStorage;
        public readonly ImmutableArray<Outputs.ScheduleActivityActivityParaSnapshotByTimeOffsetTaskWatermarkSet> WatermarkSets;

        [OutputConstructor]
        private ScheduleActivityActivityParaSnapshotByTimeOffsetTask(
            int definition,

            ImmutableArray<string> extTimeOffsetSets,

            Outputs.ScheduleActivityActivityParaSnapshotByTimeOffsetTaskObjectNumberFormat? objectNumberFormat,

            string? outputObjectPath,

            Outputs.ScheduleActivityActivityParaSnapshotByTimeOffsetTaskOutputStorage? outputStorage,

            ImmutableArray<Outputs.ScheduleActivityActivityParaSnapshotByTimeOffsetTaskWatermarkSet> watermarkSets)
        {
            Definition = definition;
            ExtTimeOffsetSets = extTimeOffsetSets;
            ObjectNumberFormat = objectNumberFormat;
            OutputObjectPath = outputObjectPath;
            OutputStorage = outputStorage;
            WatermarkSets = watermarkSets;
        }
    }
}
