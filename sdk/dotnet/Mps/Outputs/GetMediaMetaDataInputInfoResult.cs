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
    public sealed class GetMediaMetaDataInputInfoResult
    {
        public readonly Outputs.GetMediaMetaDataInputInfoCosInputInfoResult? CosInputInfo;
        public readonly Outputs.GetMediaMetaDataInputInfoS3InputInfoResult? S3InputInfo;
        public readonly string Type;
        public readonly Outputs.GetMediaMetaDataInputInfoUrlInputInfoResult? UrlInputInfo;

        [OutputConstructor]
        private GetMediaMetaDataInputInfoResult(
            Outputs.GetMediaMetaDataInputInfoCosInputInfoResult? cosInputInfo,

            Outputs.GetMediaMetaDataInputInfoS3InputInfoResult? s3InputInfo,

            string type,

            Outputs.GetMediaMetaDataInputInfoUrlInputInfoResult? urlInputInfo)
        {
            CosInputInfo = cosInputInfo;
            S3InputInfo = s3InputInfo;
            Type = type;
            UrlInputInfo = urlInputInfo;
        }
    }
}
