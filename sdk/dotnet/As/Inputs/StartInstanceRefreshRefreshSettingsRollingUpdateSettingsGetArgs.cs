// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.As.Inputs
{

    public sealed class StartInstanceRefreshRefreshSettingsRollingUpdateSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("batchNumber", required: true)]
        public Input<int> BatchNumber { get; set; } = null!;

        [Input("batchPause")]
        public Input<string>? BatchPause { get; set; }

        [Input("maxSurge")]
        public Input<int>? MaxSurge { get; set; }

        public StartInstanceRefreshRefreshSettingsRollingUpdateSettingsGetArgs()
        {
        }
        public static new StartInstanceRefreshRefreshSettingsRollingUpdateSettingsGetArgs Empty => new StartInstanceRefreshRefreshSettingsRollingUpdateSettingsGetArgs();
    }
}
