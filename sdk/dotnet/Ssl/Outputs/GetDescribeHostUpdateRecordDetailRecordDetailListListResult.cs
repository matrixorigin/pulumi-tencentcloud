// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssl.Outputs
{

    [OutputType]
    public sealed class GetDescribeHostUpdateRecordDetailRecordDetailListListResult
    {
        public readonly string Bucket;
        public readonly string CertId;
        public readonly string CreateTime;
        public readonly ImmutableArray<string> Domains;
        public readonly string EnvId;
        public readonly string ErrorMsg;
        public readonly int Id;
        public readonly string InstanceId;
        public readonly string InstanceName;
        public readonly string ListenerId;
        public readonly string ListenerName;
        public readonly string Namespace;
        public readonly string OldCertId;
        public readonly int Port;
        public readonly string Protocol;
        public readonly string Region;
        public readonly string ResourceType;
        public readonly string SecretName;
        public readonly int SniSwitch;
        public readonly int Status;
        public readonly string TCBType;
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetDescribeHostUpdateRecordDetailRecordDetailListListResult(
            string bucket,

            string certId,

            string createTime,

            ImmutableArray<string> domains,

            string envId,

            string errorMsg,

            int id,

            string instanceId,

            string instanceName,

            string listenerId,

            string listenerName,

            string @namespace,

            string oldCertId,

            int port,

            string protocol,

            string region,

            string resourceType,

            string secretName,

            int sniSwitch,

            int status,

            string tCBType,

            string updateTime)
        {
            Bucket = bucket;
            CertId = certId;
            CreateTime = createTime;
            Domains = domains;
            EnvId = envId;
            ErrorMsg = errorMsg;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            ListenerId = listenerId;
            ListenerName = listenerName;
            Namespace = @namespace;
            OldCertId = oldCertId;
            Port = port;
            Protocol = protocol;
            Region = region;
            ResourceType = resourceType;
            SecretName = secretName;
            SniSwitch = sniSwitch;
            Status = status;
            TCBType = tCBType;
            UpdateTime = updateTime;
        }
    }
}