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
    public sealed class ContentReviewTemplateUserDefineConfigure
    {
        public readonly Outputs.ContentReviewTemplateUserDefineConfigureAsrReviewInfo? AsrReviewInfo;
        public readonly Outputs.ContentReviewTemplateUserDefineConfigureFaceReviewInfo? FaceReviewInfo;
        public readonly Outputs.ContentReviewTemplateUserDefineConfigureOcrReviewInfo? OcrReviewInfo;

        [OutputConstructor]
        private ContentReviewTemplateUserDefineConfigure(
            Outputs.ContentReviewTemplateUserDefineConfigureAsrReviewInfo? asrReviewInfo,

            Outputs.ContentReviewTemplateUserDefineConfigureFaceReviewInfo? faceReviewInfo,

            Outputs.ContentReviewTemplateUserDefineConfigureOcrReviewInfo? ocrReviewInfo)
        {
            AsrReviewInfo = asrReviewInfo;
            FaceReviewInfo = faceReviewInfo;
            OcrReviewInfo = ocrReviewInfo;
        }
    }
}