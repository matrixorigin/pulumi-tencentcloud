// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentArgs : global::Pulumi.ResourceArgs
    {
        [Input("cosInputInfo")]
        public Input<Inputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentCosInputInfoArgs>? CosInputInfo { get; set; }

        [Input("s3InputInfo")]
        public Input<Inputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentS3InputInfoArgs>? S3InputInfo { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("urlInputInfo")]
        public Input<Inputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentUrlInputInfoArgs>? UrlInputInfo { get; set; }

        public ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentArgs()
        {
        }
        public static new ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentArgs Empty => new ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentArgs();
    }
}
