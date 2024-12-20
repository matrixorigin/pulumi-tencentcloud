// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cls.Outputs
{

    [OutputType]
    public sealed class AlarmNoticeNoticeReceiver
    {
        public readonly string? EndTime;
        public readonly int? Index;
        public readonly string? NoticeContentId;
        public readonly ImmutableArray<string> ReceiverChannels;
        public readonly ImmutableArray<int> ReceiverIds;
        public readonly string ReceiverType;
        public readonly string? StartTime;

        [OutputConstructor]
        private AlarmNoticeNoticeReceiver(
            string? endTime,

            int? index,

            string? noticeContentId,

            ImmutableArray<string> receiverChannels,

            ImmutableArray<int> receiverIds,

            string receiverType,

            string? startTime)
        {
            EndTime = endTime;
            Index = index;
            NoticeContentId = noticeContentId;
            ReceiverChannels = receiverChannels;
            ReceiverIds = receiverIds;
            ReceiverType = receiverType;
            StartTime = startTime;
        }
    }
}
