// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dlc.Inputs
{

    public sealed class DataEngineCrontabResumeSuspendStrategyArgs : global::Pulumi.ResourceArgs
    {
        [Input("resumeTime")]
        public Input<string>? ResumeTime { get; set; }

        [Input("suspendStrategy")]
        public Input<int>? SuspendStrategy { get; set; }

        [Input("suspendTime")]
        public Input<string>? SuspendTime { get; set; }

        public DataEngineCrontabResumeSuspendStrategyArgs()
        {
        }
        public static new DataEngineCrontabResumeSuspendStrategyArgs Empty => new DataEngineCrontabResumeSuspendStrategyArgs();
    }
}