// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Wedata.Inputs
{

    public sealed class BaselineAlarmRuleDtoGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("alarmLevelType")]
        public Input<string>? AlarmLevelType { get; set; }

        [Input("alarmRuleId")]
        public Input<string>? AlarmRuleId { get; set; }

        public BaselineAlarmRuleDtoGetArgs()
        {
        }
        public static new BaselineAlarmRuleDtoGetArgs Empty => new BaselineAlarmRuleDtoGetArgs();
    }
}
