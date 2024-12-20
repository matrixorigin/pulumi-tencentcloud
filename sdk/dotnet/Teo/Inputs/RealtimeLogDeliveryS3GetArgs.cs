// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Teo.Inputs
{

    public sealed class RealtimeLogDeliveryS3GetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessId", required: true)]
        public Input<string> AccessId { get; set; } = null!;

        [Input("accessKey", required: true)]
        private Input<string>? _accessKey;
        public Input<string>? AccessKey
        {
            get => _accessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("compressType")]
        public Input<string>? CompressType { get; set; }

        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public RealtimeLogDeliveryS3GetArgs()
        {
        }
        public static new RealtimeLogDeliveryS3GetArgs Empty => new RealtimeLogDeliveryS3GetArgs();
    }
}
