// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cls.Inputs
{

    public sealed class ScheduledSqlDstResourceGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("bizType")]
        public Input<int>? BizType { get; set; }

        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("topicId", required: true)]
        public Input<string> TopicId { get; set; } = null!;

        public ScheduledSqlDstResourceGetArgs()
        {
        }
        public static new ScheduledSqlDstResourceGetArgs Empty => new ScheduledSqlDstResourceGetArgs();
    }
}