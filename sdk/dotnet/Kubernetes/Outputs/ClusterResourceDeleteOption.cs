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
    public sealed class ClusterResourceDeleteOption
    {
        public readonly string DeleteMode;
        public readonly string ResourceType;
        public readonly bool? SkipDeletionProtection;

        [OutputConstructor]
        private ClusterResourceDeleteOption(
            string deleteMode,

            string resourceType,

            bool? skipDeletionProtection)
        {
            DeleteMode = deleteMode;
            ResourceType = resourceType;
            SkipDeletionProtection = skipDeletionProtection;
        }
    }
}
