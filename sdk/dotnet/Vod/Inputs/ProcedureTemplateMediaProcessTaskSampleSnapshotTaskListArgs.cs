// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod.Inputs
{

    public sealed class ProcedureTemplateMediaProcessTaskSampleSnapshotTaskListArgs : global::Pulumi.ResourceArgs
    {
        [Input("definition", required: true)]
        public Input<string> Definition { get; set; } = null!;

        [Input("watermarkLists")]
        private InputList<Inputs.ProcedureTemplateMediaProcessTaskSampleSnapshotTaskListWatermarkListArgs>? _watermarkLists;
        public InputList<Inputs.ProcedureTemplateMediaProcessTaskSampleSnapshotTaskListWatermarkListArgs> WatermarkLists
        {
            get => _watermarkLists ?? (_watermarkLists = new InputList<Inputs.ProcedureTemplateMediaProcessTaskSampleSnapshotTaskListWatermarkListArgs>());
            set => _watermarkLists = value;
        }

        public ProcedureTemplateMediaProcessTaskSampleSnapshotTaskListArgs()
        {
        }
        public static new ProcedureTemplateMediaProcessTaskSampleSnapshotTaskListArgs Empty => new ProcedureTemplateMediaProcessTaskSampleSnapshotTaskListArgs();
    }
}