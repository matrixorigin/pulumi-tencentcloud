// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOutputStorageArgs : global::Pulumi.ResourceArgs
    {
        [Input("cosOutputStorage")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOutputStorageCosOutputStorageArgs>? CosOutputStorage { get; set; }

        [Input("s3OutputStorage")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOutputStorageS3OutputStorageArgs>? S3OutputStorage { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOutputStorageArgs()
        {
        }
        public static new ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOutputStorageArgs Empty => new ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOutputStorageArgs();
    }
}