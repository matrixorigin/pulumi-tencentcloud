// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor.Inputs
{

    public sealed class TmpAlertGroupCustomReceiverAllowedTimeRangeGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("end")]
        public Input<string>? End { get; set; }

        [Input("start")]
        public Input<string>? Start { get; set; }

        public TmpAlertGroupCustomReceiverAllowedTimeRangeGetArgs()
        {
        }
        public static new TmpAlertGroupCustomReceiverAllowedTimeRangeGetArgs Empty => new TmpAlertGroupCustomReceiverAllowedTimeRangeGetArgs();
    }
}
