// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcr.Inputs
{

    public sealed class ServiceAccountPermissionArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<string>? _actions;
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        public ServiceAccountPermissionArgs()
        {
        }
        public static new ServiceAccountPermissionArgs Empty => new ServiceAccountPermissionArgs();
    }
}