// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class EditMediaOperationOutputConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("container")]
        public Input<string>? Container { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public EditMediaOperationOutputConfigGetArgs()
        {
        }
        public static new EditMediaOperationOutputConfigGetArgs Empty => new EditMediaOperationOutputConfigGetArgs();
    }
}