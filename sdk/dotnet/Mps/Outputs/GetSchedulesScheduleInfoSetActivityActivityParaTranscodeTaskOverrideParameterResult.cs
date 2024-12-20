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
    public sealed class GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterResult
    {
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitleResult> AddOnSubtitles;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAddonAudioStreamResult> AddonAudioStreams;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAudioTemplateResult> AudioTemplates;
        public readonly string Container;
        public readonly int RemoveAudio;
        public readonly int RemoveVideo;
        public readonly string StdExtInfo;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterSubtitleTemplateResult> SubtitleTemplates;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterTehdConfigResult> TehdConfigs;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterVideoTemplateResult> VideoTemplates;

        [OutputConstructor]
        private GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterResult(
            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitleResult> addOnSubtitles,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAddonAudioStreamResult> addonAudioStreams,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAudioTemplateResult> audioTemplates,

            string container,

            int removeAudio,

            int removeVideo,

            string stdExtInfo,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterSubtitleTemplateResult> subtitleTemplates,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterTehdConfigResult> tehdConfigs,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterVideoTemplateResult> videoTemplates)
        {
            AddOnSubtitles = addOnSubtitles;
            AddonAudioStreams = addonAudioStreams;
            AudioTemplates = audioTemplates;
            Container = container;
            RemoveAudio = removeAudio;
            RemoveVideo = removeVideo;
            StdExtInfo = stdExtInfo;
            SubtitleTemplates = subtitleTemplates;
            TehdConfigs = tehdConfigs;
            VideoTemplates = videoTemplates;
        }
    }
}
