// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Thpc.Inputs
{

    public sealed class WorkspacesSpaceChargePrepaidGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("renewFlag")]
        public Input<string>? RenewFlag { get; set; }

        public WorkspacesSpaceChargePrepaidGetArgs()
        {
        }
        public static new WorkspacesSpaceChargePrepaidGetArgs Empty => new WorkspacesSpaceChargePrepaidGetArgs();
    }
}
