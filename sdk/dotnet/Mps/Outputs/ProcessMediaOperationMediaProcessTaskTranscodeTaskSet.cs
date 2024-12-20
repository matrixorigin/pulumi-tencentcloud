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
    public sealed class ProcessMediaOperationMediaProcessTaskTranscodeTaskSet
    {
        public readonly int Definition;
        public readonly double? EndTimeOffset;
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameter? HeadTailParameter;
        public readonly ImmutableArray<Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetMosaicSet> MosaicSets;
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetObjectNumberFormat? ObjectNumberFormat;
        public readonly string? OutputObjectPath;
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOutputStorage? OutputStorage;
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOverrideParameter? OverrideParameter;
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetRawParameter? RawParameter;
        public readonly string? SegmentObjectName;
        public readonly double? StartTimeOffset;
        public readonly ImmutableArray<Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetWatermarkSet> WatermarkSets;

        [OutputConstructor]
        private ProcessMediaOperationMediaProcessTaskTranscodeTaskSet(
            int definition,

            double? endTimeOffset,

            Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameter? headTailParameter,

            ImmutableArray<Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetMosaicSet> mosaicSets,

            Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetObjectNumberFormat? objectNumberFormat,

            string? outputObjectPath,

            Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOutputStorage? outputStorage,

            Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOverrideParameter? overrideParameter,

            Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetRawParameter? rawParameter,

            string? segmentObjectName,

            double? startTimeOffset,

            ImmutableArray<Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetWatermarkSet> watermarkSets)
        {
            Definition = definition;
            EndTimeOffset = endTimeOffset;
            HeadTailParameter = headTailParameter;
            MosaicSets = mosaicSets;
            ObjectNumberFormat = objectNumberFormat;
            OutputObjectPath = outputObjectPath;
            OutputStorage = outputStorage;
            OverrideParameter = overrideParameter;
            RawParameter = rawParameter;
            SegmentObjectName = segmentObjectName;
            StartTimeOffset = startTimeOffset;
            WatermarkSets = watermarkSets;
        }
    }
}
