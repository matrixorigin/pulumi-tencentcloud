// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ses.Outputs
{

    [OutputType]
    public sealed class GetReceiversDataResult
    {
        public readonly int Count;
        public readonly string CreateTime;
        public readonly string Desc;
        public readonly int ReceiverId;
        public readonly string ReceiversName;
        public readonly int ReceiversStatus;

        [OutputConstructor]
        private GetReceiversDataResult(
            int count,

            string createTime,

            string desc,

            int receiverId,

            string receiversName,

            int receiversStatus)
        {
            Count = count;
            CreateTime = createTime;
            Desc = desc;
            ReceiverId = receiverId;
            ReceiversName = receiversName;
            ReceiversStatus = receiversStatus;
        }
    }
}
