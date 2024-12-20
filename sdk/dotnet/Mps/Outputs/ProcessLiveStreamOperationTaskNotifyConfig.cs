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
    public sealed class ProcessLiveStreamOperationTaskNotifyConfig
    {
        public readonly string? CmqModel;
        public readonly string? CmqRegion;
        public readonly string? NotifyType;
        public readonly string? NotifyUrl;
        public readonly string? QueueName;
        public readonly string? TopicName;

        [OutputConstructor]
        private ProcessLiveStreamOperationTaskNotifyConfig(
            string? cmqModel,

            string? cmqRegion,

            string? notifyType,

            string? notifyUrl,

            string? queueName,

            string? topicName)
        {
            CmqModel = cmqModel;
            CmqRegion = cmqRegion;
            NotifyType = notifyType;
            NotifyUrl = notifyUrl;
            QueueName = queueName;
            TopicName = topicName;
        }
    }
}
