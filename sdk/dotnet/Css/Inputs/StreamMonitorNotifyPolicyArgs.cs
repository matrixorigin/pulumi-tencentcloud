// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Css.Inputs
{

    public sealed class StreamMonitorNotifyPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("callbackUrl")]
        public Input<string>? CallbackUrl { get; set; }

        [Input("notifyPolicyType")]
        public Input<int>? NotifyPolicyType { get; set; }

        public StreamMonitorNotifyPolicyArgs()
        {
        }
        public static new StreamMonitorNotifyPolicyArgs Empty => new StreamMonitorNotifyPolicyArgs();
    }
}
