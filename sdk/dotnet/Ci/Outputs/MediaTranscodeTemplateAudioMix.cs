// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ci.Outputs
{

    [OutputType]
    public sealed class MediaTranscodeTemplateAudioMix
    {
        public readonly string AudioSource;
        public readonly Outputs.MediaTranscodeTemplateAudioMixEffectConfig? EffectConfig;
        public readonly string? MixMode;
        public readonly string? Replace;

        [OutputConstructor]
        private MediaTranscodeTemplateAudioMix(
            string audioSource,

            Outputs.MediaTranscodeTemplateAudioMixEffectConfig? effectConfig,

            string? mixMode,

            string? replace)
        {
            AudioSource = audioSource;
            EffectConfig = effectConfig;
            MixMode = mixMode;
            Replace = replace;
        }
    }
}