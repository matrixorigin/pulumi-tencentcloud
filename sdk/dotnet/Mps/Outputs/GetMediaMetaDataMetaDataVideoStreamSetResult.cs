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
    public sealed class GetMediaMetaDataMetaDataVideoStreamSetResult
    {
        public readonly int Bitrate;
        public readonly string Codec;
        public readonly string ColorPrimaries;
        public readonly string ColorSpace;
        public readonly string ColorTransfer;
        public readonly int Fps;
        public readonly string HdrType;
        public readonly int Height;
        public readonly int Width;

        [OutputConstructor]
        private GetMediaMetaDataMetaDataVideoStreamSetResult(
            int bitrate,

            string codec,

            string colorPrimaries,

            string colorSpace,

            string colorTransfer,

            int fps,

            string hdrType,

            int height,

            int width)
        {
            Bitrate = bitrate;
            Codec = codec;
            ColorPrimaries = colorPrimaries;
            ColorSpace = colorSpace;
            ColorTransfer = colorTransfer;
            Fps = fps;
            HdrType = hdrType;
            Height = height;
            Width = width;
        }
    }
}