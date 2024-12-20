// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetArgs : global::Pulumi.ResourceArgs
    {
        [Input("addOnSubtitles")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetAddOnSubtitleArgs>? _addOnSubtitles;
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetAddOnSubtitleArgs> AddOnSubtitles
        {
            get => _addOnSubtitles ?? (_addOnSubtitles = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetAddOnSubtitleArgs>());
            set => _addOnSubtitles = value;
        }

        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        [Input("outputObjectPath")]
        public Input<string>? OutputObjectPath { get; set; }

        [Input("outputStorage")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetOutputStorageArgs>? OutputStorage { get; set; }

        [Input("segmentObjectName")]
        public Input<string>? SegmentObjectName { get; set; }

        [Input("subStreamObjectName")]
        public Input<string>? SubStreamObjectName { get; set; }

        [Input("watermarkSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetArgs>? _watermarkSets;
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetArgs> WatermarkSets
        {
            get => _watermarkSets ?? (_watermarkSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetArgs>());
            set => _watermarkSets = value;
        }

        public ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetArgs()
        {
        }
        public static new ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetArgs Empty => new ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetArgs();
    }
}
