// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Lighthouse.Inputs
{

    public sealed class RenewInstanceInstanceChargePrepaidArgs : global::Pulumi.ResourceArgs
    {
        [Input("period", required: true)]
        public Input<int> Period { get; set; } = null!;

        [Input("renewFlag")]
        public Input<string>? RenewFlag { get; set; }

        public RenewInstanceInstanceChargePrepaidArgs()
        {
        }
        public static new RenewInstanceInstanceChargePrepaidArgs Empty => new RenewInstanceInstanceChargePrepaidArgs();
    }
}