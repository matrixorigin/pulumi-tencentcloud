// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Oceanus.Inputs
{

    public sealed class JobConfigExpertModeConfigurationNodeConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("configurations")]
        private InputList<Inputs.JobConfigExpertModeConfigurationNodeConfigConfigurationArgs>? _configurations;
        public InputList<Inputs.JobConfigExpertModeConfigurationNodeConfigConfigurationArgs> Configurations
        {
            get => _configurations ?? (_configurations = new InputList<Inputs.JobConfigExpertModeConfigurationNodeConfigConfigurationArgs>());
            set => _configurations = value;
        }

        [Input("id", required: true)]
        public Input<int> Id { get; set; } = null!;

        [Input("parallelism")]
        public Input<int>? Parallelism { get; set; }

        [Input("slotSharingGroup")]
        public Input<string>? SlotSharingGroup { get; set; }

        [Input("stateTtl")]
        public Input<string>? StateTtl { get; set; }

        public JobConfigExpertModeConfigurationNodeConfigArgs()
        {
        }
        public static new JobConfigExpertModeConfigurationNodeConfigArgs Empty => new JobConfigExpertModeConfigurationNodeConfigArgs();
    }
}
