// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class ProcessLiveStreamOperationTaskNotifyConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("cmqModel")]
        public Input<string>? CmqModel { get; set; }

        [Input("cmqRegion")]
        public Input<string>? CmqRegion { get; set; }

        [Input("notifyType")]
        public Input<string>? NotifyType { get; set; }

        [Input("notifyUrl")]
        public Input<string>? NotifyUrl { get; set; }

        [Input("queueName")]
        public Input<string>? QueueName { get; set; }

        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public ProcessLiveStreamOperationTaskNotifyConfigArgs()
        {
        }
        public static new ProcessLiveStreamOperationTaskNotifyConfigArgs Empty => new ProcessLiveStreamOperationTaskNotifyConfigArgs();
    }
}
