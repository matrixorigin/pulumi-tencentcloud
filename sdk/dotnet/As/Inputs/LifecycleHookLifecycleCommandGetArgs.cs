// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.As.Inputs
{

    public sealed class LifecycleHookLifecycleCommandGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("commandId", required: true)]
        public Input<string> CommandId { get; set; } = null!;

        [Input("parameters")]
        public Input<string>? Parameters { get; set; }

        public LifecycleHookLifecycleCommandGetArgs()
        {
        }
        public static new LifecycleHookLifecycleCommandGetArgs Empty => new LifecycleHookLifecycleCommandGetArgs();
    }
}