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
    public sealed class GetClusterNativeNodePoolsNodePoolNativeLifecycleResult
    {
        public readonly string PostInit;
        public readonly string PreInit;

        [OutputConstructor]
        private GetClusterNativeNodePoolsNodePoolNativeLifecycleResult(
            string postInit,

            string preInit)
        {
            PostInit = postInit;
            PreInit = preInit;
        }
    }
}
