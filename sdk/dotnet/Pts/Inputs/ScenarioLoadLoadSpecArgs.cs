// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Pts.Inputs
{

    public sealed class ScenarioLoadLoadSpecArgs : global::Pulumi.ResourceArgs
    {
        [Input("concurrency")]
        public Input<Inputs.ScenarioLoadLoadSpecConcurrencyArgs>? Concurrency { get; set; }

        [Input("requestsPerSecond")]
        public Input<Inputs.ScenarioLoadLoadSpecRequestsPerSecondArgs>? RequestsPerSecond { get; set; }

        [Input("scriptOrigin")]
        public Input<Inputs.ScenarioLoadLoadSpecScriptOriginArgs>? ScriptOrigin { get; set; }

        public ScenarioLoadLoadSpecArgs()
        {
        }
        public static new ScenarioLoadLoadSpecArgs Empty => new ScenarioLoadLoadSpecArgs();
    }
}