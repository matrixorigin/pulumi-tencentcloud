// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elasticsearch.Outputs
{

    [OutputType]
    public sealed class GetLogstashInstanceLogsInstanceLogListResult
    {
        public readonly string Ip;
        public readonly string Level;
        public readonly string Message;
        public readonly string NodeId;
        public readonly string Time;

        [OutputConstructor]
        private GetLogstashInstanceLogsInstanceLogListResult(
            string ip,

            string level,

            string message,

            string nodeId,

            string time)
        {
            Ip = ip;
            Level = level;
            Message = message;
            NodeId = nodeId;
            Time = time;
        }
    }
}