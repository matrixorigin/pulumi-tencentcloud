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
    public sealed class GetSchedulesScheduleInfoSetActivityActivityParaAdaptiveDynamicStreamingTaskResult
    {
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaAdaptiveDynamicStreamingTaskAddOnSubtitleResult> AddOnSubtitles;
        public readonly int Definition;
        public readonly string OutputObjectPath;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaAdaptiveDynamicStreamingTaskOutputStorageResult> OutputStorages;
        public readonly string SegmentObjectName;
        public readonly string SubStreamObjectName;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetResult> WatermarkSets;

        [OutputConstructor]
        private GetSchedulesScheduleInfoSetActivityActivityParaAdaptiveDynamicStreamingTaskResult(
            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaAdaptiveDynamicStreamingTaskAddOnSubtitleResult> addOnSubtitles,

            int definition,

            string outputObjectPath,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaAdaptiveDynamicStreamingTaskOutputStorageResult> outputStorages,

            string segmentObjectName,

            string subStreamObjectName,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetResult> watermarkSets)
        {
            AddOnSubtitles = addOnSubtitles;
            Definition = definition;
            OutputObjectPath = outputObjectPath;
            OutputStorages = outputStorages;
            SegmentObjectName = segmentObjectName;
            SubStreamObjectName = subStreamObjectName;
            WatermarkSets = watermarkSets;
        }
    }
}