// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdwdoris.Inputs
{

    public sealed class InstanceChargePropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        [Input("renewFlag")]
        public Input<int>? RenewFlag { get; set; }

        [Input("timeSpan")]
        public Input<int>? TimeSpan { get; set; }

        [Input("timeUnit")]
        public Input<string>? TimeUnit { get; set; }

        public InstanceChargePropertiesArgs()
        {
        }
        public static new InstanceChargePropertiesArgs Empty => new InstanceChargePropertiesArgs();
    }
}
