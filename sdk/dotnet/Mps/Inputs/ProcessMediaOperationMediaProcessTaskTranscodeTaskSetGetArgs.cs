// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class ProcessMediaOperationMediaProcessTaskTranscodeTaskSetGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        [Input("endTimeOffset")]
        public Input<double>? EndTimeOffset { get; set; }

        [Input("headTailParameter")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameterGetArgs>? HeadTailParameter { get; set; }

        [Input("mosaicSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetMosaicSetGetArgs>? _mosaicSets;
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetMosaicSetGetArgs> MosaicSets
        {
            get => _mosaicSets ?? (_mosaicSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetMosaicSetGetArgs>());
            set => _mosaicSets = value;
        }

        [Input("objectNumberFormat")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetObjectNumberFormatGetArgs>? ObjectNumberFormat { get; set; }

        [Input("outputObjectPath")]
        public Input<string>? OutputObjectPath { get; set; }

        [Input("outputStorage")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOutputStorageGetArgs>? OutputStorage { get; set; }

        [Input("overrideParameter")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOverrideParameterGetArgs>? OverrideParameter { get; set; }

        [Input("rawParameter")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetRawParameterGetArgs>? RawParameter { get; set; }

        [Input("segmentObjectName")]
        public Input<string>? SegmentObjectName { get; set; }

        [Input("startTimeOffset")]
        public Input<double>? StartTimeOffset { get; set; }

        [Input("watermarkSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetWatermarkSetGetArgs>? _watermarkSets;
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetWatermarkSetGetArgs> WatermarkSets
        {
            get => _watermarkSets ?? (_watermarkSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetWatermarkSetGetArgs>());
            set => _watermarkSets = value;
        }

        public ProcessMediaOperationMediaProcessTaskTranscodeTaskSetGetArgs()
        {
        }
        public static new ProcessMediaOperationMediaProcessTaskTranscodeTaskSetGetArgs Empty => new ProcessMediaOperationMediaProcessTaskTranscodeTaskSetGetArgs();
    }
}