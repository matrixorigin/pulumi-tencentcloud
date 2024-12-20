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
    public sealed class NativeNodePoolNativeManagement
    {
        public readonly ImmutableArray<string> Hosts;
        public readonly ImmutableArray<string> KernelArgs;
        public readonly ImmutableArray<string> Nameservers;

        [OutputConstructor]
        private NativeNodePoolNativeManagement(
            ImmutableArray<string> hosts,

            ImmutableArray<string> kernelArgs,

            ImmutableArray<string> nameservers)
        {
            Hosts = hosts;
            KernelArgs = kernelArgs;
            Nameservers = nameservers;
        }
    }
}
