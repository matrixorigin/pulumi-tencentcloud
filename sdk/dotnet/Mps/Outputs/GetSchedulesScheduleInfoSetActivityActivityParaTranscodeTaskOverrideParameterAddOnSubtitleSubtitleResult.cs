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
    public sealed class GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitleSubtitleResult
    {
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitleSubtitleCosInputInfoResult> CosInputInfos;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitleSubtitleS3InputInfoResult> S3InputInfos;
        public readonly string Type;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitleSubtitleUrlInputInfoResult> UrlInputInfos;

        [OutputConstructor]
        private GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitleSubtitleResult(
            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitleSubtitleCosInputInfoResult> cosInputInfos,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitleSubtitleS3InputInfoResult> s3InputInfos,

            string type,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitleSubtitleUrlInputInfoResult> urlInputInfos)
        {
            CosInputInfos = cosInputInfos;
            S3InputInfos = s3InputInfos;
            Type = type;
            UrlInputInfos = urlInputInfos;
        }
    }
}