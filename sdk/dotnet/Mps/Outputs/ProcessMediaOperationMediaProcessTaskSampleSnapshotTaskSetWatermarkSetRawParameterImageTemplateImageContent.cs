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
    public sealed class ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetRawParameterImageTemplateImageContent
    {
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetRawParameterImageTemplateImageContentCosInputInfo? CosInputInfo;
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetRawParameterImageTemplateImageContentS3InputInfo? S3InputInfo;
        public readonly string Type;
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetRawParameterImageTemplateImageContentUrlInputInfo? UrlInputInfo;

        [OutputConstructor]
        private ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetRawParameterImageTemplateImageContent(
            Outputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetRawParameterImageTemplateImageContentCosInputInfo? cosInputInfo,

            Outputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetRawParameterImageTemplateImageContentS3InputInfo? s3InputInfo,

            string type,

            Outputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetRawParameterImageTemplateImageContentUrlInputInfo? urlInputInfo)
        {
            CosInputInfo = cosInputInfo;
            S3InputInfo = s3InputInfo;
            Type = type;
            UrlInputInfo = urlInputInfo;
        }
    }
}