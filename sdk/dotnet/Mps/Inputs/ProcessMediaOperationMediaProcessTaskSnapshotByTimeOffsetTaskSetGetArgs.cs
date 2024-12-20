// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        [Input("extTimeOffsetSets")]
        private InputList<string>? _extTimeOffsetSets;
        public InputList<string> ExtTimeOffsetSets
        {
            get => _extTimeOffsetSets ?? (_extTimeOffsetSets = new InputList<string>());
            set => _extTimeOffsetSets = value;
        }

        [Input("objectNumberFormat")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetObjectNumberFormatGetArgs>? ObjectNumberFormat { get; set; }

        [Input("outputObjectPath")]
        public Input<string>? OutputObjectPath { get; set; }

        [Input("outputStorage")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetOutputStorageGetArgs>? OutputStorage { get; set; }

        [Input("timeOffsetSets")]
        private InputList<double>? _timeOffsetSets;
        public InputList<double> TimeOffsetSets
        {
            get => _timeOffsetSets ?? (_timeOffsetSets = new InputList<double>());
            set => _timeOffsetSets = value;
        }

        [Input("watermarkSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetWatermarkSetGetArgs>? _watermarkSets;
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetWatermarkSetGetArgs> WatermarkSets
        {
            get => _watermarkSets ?? (_watermarkSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetWatermarkSetGetArgs>());
            set => _watermarkSets = value;
        }

        public ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetGetArgs()
        {
        }
        public static new ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetGetArgs Empty => new ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetGetArgs();
    }
}
