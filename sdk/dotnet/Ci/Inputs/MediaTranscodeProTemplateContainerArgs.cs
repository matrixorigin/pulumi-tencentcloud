// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ci.Inputs
{

    public sealed class MediaTranscodeProTemplateContainerArgs : global::Pulumi.ResourceArgs
    {
        [Input("clipConfig")]
        public Input<Inputs.MediaTranscodeProTemplateContainerClipConfigArgs>? ClipConfig { get; set; }

        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        public MediaTranscodeProTemplateContainerArgs()
        {
        }
        public static new MediaTranscodeProTemplateContainerArgs Empty => new MediaTranscodeProTemplateContainerArgs();
    }
}