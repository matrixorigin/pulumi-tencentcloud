// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kubernetes.Outputs
{

    [OutputType]
    public sealed class GetClusterInstancesInstanceSetInstanceAdvancedSettingGpuArgResult
    {
        public readonly ImmutableArray<Outputs.GetClusterInstancesInstanceSetInstanceAdvancedSettingGpuArgCudaResult> Cudas;
        public readonly ImmutableArray<Outputs.GetClusterInstancesInstanceSetInstanceAdvancedSettingGpuArgCudnnResult> Cudnns;
        public readonly ImmutableArray<Outputs.GetClusterInstancesInstanceSetInstanceAdvancedSettingGpuArgCustomDriverResult> CustomDrivers;
        public readonly ImmutableArray<Outputs.GetClusterInstancesInstanceSetInstanceAdvancedSettingGpuArgDriverResult> Drivers;
        public readonly bool MigEnable;

        [OutputConstructor]
        private GetClusterInstancesInstanceSetInstanceAdvancedSettingGpuArgResult(
            ImmutableArray<Outputs.GetClusterInstancesInstanceSetInstanceAdvancedSettingGpuArgCudaResult> cudas,

            ImmutableArray<Outputs.GetClusterInstancesInstanceSetInstanceAdvancedSettingGpuArgCudnnResult> cudnns,

            ImmutableArray<Outputs.GetClusterInstancesInstanceSetInstanceAdvancedSettingGpuArgCustomDriverResult> customDrivers,

            ImmutableArray<Outputs.GetClusterInstancesInstanceSetInstanceAdvancedSettingGpuArgDriverResult> drivers,

            bool migEnable)
        {
            Cudas = cudas;
            Cudnns = cudnns;
            CustomDrivers = customDrivers;
            Drivers = drivers;
            MigEnable = migEnable;
        }
    }
}