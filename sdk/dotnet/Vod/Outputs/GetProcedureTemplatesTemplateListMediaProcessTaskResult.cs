// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod.Outputs
{

    [OutputType]
    public sealed class GetProcedureTemplatesTemplateListMediaProcessTaskResult
    {
        public readonly ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskAdaptiveDynamicStreamingTaskListResult> AdaptiveDynamicStreamingTaskLists;
        public readonly ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskAnimatedGraphicTaskListResult> AnimatedGraphicTaskLists;
        public readonly ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskCoverBySnapshotTaskListResult> CoverBySnapshotTaskLists;
        public readonly ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskImageSpriteTaskListResult> ImageSpriteTaskLists;
        public readonly ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskSampleSnapshotTaskListResult> SampleSnapshotTaskLists;
        public readonly ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskSnapshotByTimeOffsetTaskListResult> SnapshotByTimeOffsetTaskLists;
        public readonly ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskTranscodeTaskListResult> TranscodeTaskLists;

        [OutputConstructor]
        private GetProcedureTemplatesTemplateListMediaProcessTaskResult(
            ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskAdaptiveDynamicStreamingTaskListResult> adaptiveDynamicStreamingTaskLists,

            ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskAnimatedGraphicTaskListResult> animatedGraphicTaskLists,

            ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskCoverBySnapshotTaskListResult> coverBySnapshotTaskLists,

            ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskImageSpriteTaskListResult> imageSpriteTaskLists,

            ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskSampleSnapshotTaskListResult> sampleSnapshotTaskLists,

            ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskSnapshotByTimeOffsetTaskListResult> snapshotByTimeOffsetTaskLists,

            ImmutableArray<Outputs.GetProcedureTemplatesTemplateListMediaProcessTaskTranscodeTaskListResult> transcodeTaskLists)
        {
            AdaptiveDynamicStreamingTaskLists = adaptiveDynamicStreamingTaskLists;
            AnimatedGraphicTaskLists = animatedGraphicTaskLists;
            CoverBySnapshotTaskLists = coverBySnapshotTaskLists;
            ImageSpriteTaskLists = imageSpriteTaskLists;
            SampleSnapshotTaskLists = sampleSnapshotTaskLists;
            SnapshotByTimeOffsetTaskLists = snapshotByTimeOffsetTaskLists;
            TranscodeTaskLists = transcodeTaskLists;
        }
    }
}