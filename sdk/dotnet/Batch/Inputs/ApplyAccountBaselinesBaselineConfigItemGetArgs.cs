// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Batch.Inputs
{

    public sealed class ApplyAccountBaselinesBaselineConfigItemGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("configuration")]
        public Input<string>? Configuration { get; set; }

        [Input("identifier")]
        public Input<string>? Identifier { get; set; }

        public ApplyAccountBaselinesBaselineConfigItemGetArgs()
        {
        }
        public static new ApplyAccountBaselinesBaselineConfigItemGetArgs Empty => new ApplyAccountBaselinesBaselineConfigItemGetArgs();
    }
}
