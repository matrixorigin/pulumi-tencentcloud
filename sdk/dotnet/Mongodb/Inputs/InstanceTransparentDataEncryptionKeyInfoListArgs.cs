// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mongodb.Inputs
{

    public sealed class InstanceTransparentDataEncryptionKeyInfoListArgs : global::Pulumi.ResourceArgs
    {
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        [Input("keyOrigin")]
        public Input<string>? KeyOrigin { get; set; }

        [Input("keyUsage")]
        public Input<string>? KeyUsage { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public InstanceTransparentDataEncryptionKeyInfoListArgs()
        {
        }
        public static new InstanceTransparentDataEncryptionKeyInfoListArgs Empty => new InstanceTransparentDataEncryptionKeyInfoListArgs();
    }
}