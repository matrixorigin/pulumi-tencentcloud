// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Waf.Inputs
{

    public sealed class IpAccessControlV2JobDateTimeCronArgs : global::Pulumi.ResourceArgs
    {
        [Input("days")]
        private InputList<int>? _days;
        public InputList<int> Days
        {
            get => _days ?? (_days = new InputList<int>());
            set => _days = value;
        }

        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        [Input("wDays")]
        private InputList<int>? _wDays;
        public InputList<int> WDays
        {
            get => _wDays ?? (_wDays = new InputList<int>());
            set => _wDays = value;
        }

        public IpAccessControlV2JobDateTimeCronArgs()
        {
        }
        public static new IpAccessControlV2JobDateTimeCronArgs Empty => new IpAccessControlV2JobDateTimeCronArgs();
    }
}