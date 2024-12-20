// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class GetAlarmAllNamespacesCommonNamespaceResult
    {
        public readonly ImmutableArray<Outputs.GetAlarmAllNamespacesCommonNamespaceDimensionResult> Dimensions;
        public readonly string Id;
        public readonly string MonitorType;
        public readonly string Name;

        [OutputConstructor]
        private GetAlarmAllNamespacesCommonNamespaceResult(
            ImmutableArray<Outputs.GetAlarmAllNamespacesCommonNamespaceDimensionResult> dimensions,

            string id,

            string monitorType,

            string name)
        {
            Dimensions = dimensions;
            Id = id;
            MonitorType = monitorType;
            Name = name;
        }
    }
}
