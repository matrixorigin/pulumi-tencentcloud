// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class GetTopicSubscribeGroupGroupsInfoResult
    {
        public readonly string ErrorCode;
        public readonly string Group;
        public readonly ImmutableArray<Outputs.GetTopicSubscribeGroupGroupsInfoMemberResult> Members;
        public readonly string Protocol;
        public readonly string ProtocolType;
        public readonly string State;

        [OutputConstructor]
        private GetTopicSubscribeGroupGroupsInfoResult(
            string errorCode,

            string group,

            ImmutableArray<Outputs.GetTopicSubscribeGroupGroupsInfoMemberResult> members,

            string protocol,

            string protocolType,

            string state)
        {
            ErrorCode = errorCode;
            Group = group;
            Members = members;
            Protocol = protocol;
            ProtocolType = protocolType;
            State = state;
        }
    }
}