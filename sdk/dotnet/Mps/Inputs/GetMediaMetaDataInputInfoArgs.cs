// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class GetMediaMetaDataInputInfoInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("cosInputInfo")]
        public Input<Inputs.GetMediaMetaDataInputInfoCosInputInfoInputArgs>? CosInputInfo { get; set; }

        [Input("s3InputInfo")]
        public Input<Inputs.GetMediaMetaDataInputInfoS3InputInfoInputArgs>? S3InputInfo { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("urlInputInfo")]
        public Input<Inputs.GetMediaMetaDataInputInfoUrlInputInfoInputArgs>? UrlInputInfo { get; set; }

        public GetMediaMetaDataInputInfoInputArgs()
        {
        }
        public static new GetMediaMetaDataInputInfoInputArgs Empty => new GetMediaMetaDataInputInfoInputArgs();
    }
}
