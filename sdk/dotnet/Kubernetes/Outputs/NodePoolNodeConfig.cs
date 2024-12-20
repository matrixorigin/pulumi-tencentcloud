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
    public sealed class NodePoolNodeConfig
    {
        public readonly ImmutableArray<Outputs.NodePoolNodeConfigDataDisk> DataDisks;
        public readonly int? DesiredPodNum;
        public readonly string? DockerGraphPath;
        public readonly ImmutableArray<string> ExtraArgs;
        public readonly Outputs.NodePoolNodeConfigGpuArgs? GpuArgs;
        public readonly bool? IsSchedule;
        public readonly string? MountTarget;
        public readonly string? PreStartUserScript;
        public readonly string? UserData;

        [OutputConstructor]
        private NodePoolNodeConfig(
            ImmutableArray<Outputs.NodePoolNodeConfigDataDisk> dataDisks,

            int? desiredPodNum,

            string? dockerGraphPath,

            ImmutableArray<string> extraArgs,

            Outputs.NodePoolNodeConfigGpuArgs? gpuArgs,

            bool? isSchedule,

            string? mountTarget,

            string? preStartUserScript,

            string? userData)
        {
            DataDisks = dataDisks;
            DesiredPodNum = desiredPodNum;
            DockerGraphPath = dockerGraphPath;
            ExtraArgs = extraArgs;
            GpuArgs = gpuArgs;
            IsSchedule = isSchedule;
            MountTarget = mountTarget;
            PreStartUserScript = preStartUserScript;
            UserData = userData;
        }
    }
}
