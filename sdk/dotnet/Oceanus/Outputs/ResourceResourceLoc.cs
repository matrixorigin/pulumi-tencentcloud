// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Oceanus.Outputs
{

    [OutputType]
    public sealed class ResourceResourceLoc
    {
        public readonly Outputs.ResourceResourceLocParam Param;
        public readonly int StorageType;

        [OutputConstructor]
        private ResourceResourceLoc(
            Outputs.ResourceResourceLocParam param,

            int storageType)
        {
            Param = param;
            StorageType = storageType;
        }
    }
}
