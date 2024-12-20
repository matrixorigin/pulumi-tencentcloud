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
    public sealed class GetSchedulesScheduleInfoSetActivityActivityParaSampleSnapshotTaskResult
    {
        public readonly int Definition;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaSampleSnapshotTaskObjectNumberFormatResult> ObjectNumberFormats;
        public readonly string OutputObjectPath;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaSampleSnapshotTaskOutputStorageResult> OutputStorages;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaSampleSnapshotTaskWatermarkSetResult> WatermarkSets;

        [OutputConstructor]
        private GetSchedulesScheduleInfoSetActivityActivityParaSampleSnapshotTaskResult(
            int definition,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaSampleSnapshotTaskObjectNumberFormatResult> objectNumberFormats,

            string outputObjectPath,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaSampleSnapshotTaskOutputStorageResult> outputStorages,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaSampleSnapshotTaskWatermarkSetResult> watermarkSets)
        {
            Definition = definition;
            ObjectNumberFormats = objectNumberFormats;
            OutputObjectPath = outputObjectPath;
            OutputStorages = outputStorages;
            WatermarkSets = watermarkSets;
        }
    }
}
