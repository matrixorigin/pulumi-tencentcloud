// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Oceanus.Inputs
{

    public sealed class JobConfigExpertModeConfigurationSlotSharingGroupGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("spec", required: true)]
        public Input<Inputs.JobConfigExpertModeConfigurationSlotSharingGroupSpecGetArgs> Spec { get; set; } = null!;

        public JobConfigExpertModeConfigurationSlotSharingGroupGetArgs()
        {
        }
        public static new JobConfigExpertModeConfigurationSlotSharingGroupGetArgs Empty => new JobConfigExpertModeConfigurationSlotSharingGroupGetArgs();
    }
}