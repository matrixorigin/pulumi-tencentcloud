// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Reserved
{
    public static class GetInstances
    {
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("tencentcloud:Reserved/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithDefaults());

        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("tencentcloud:Reserved/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancesArgs : global::Pulumi.InvokeArgs
    {
        [Input("availabilityZone")]
        public string? AvailabilityZone { get; set; }

        [Input("instanceType")]
        public string? InstanceType { get; set; }

        [Input("reservedInstanceId")]
        public string? ReservedInstanceId { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstancesArgs()
        {
        }
        public static new GetInstancesArgs Empty => new GetInstancesArgs();
    }

    public sealed class GetInstancesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("reservedInstanceId")]
        public Input<string>? ReservedInstanceId { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstancesInvokeArgs()
        {
        }
        public static new GetInstancesInvokeArgs Empty => new GetInstancesInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        public readonly string? AvailabilityZone;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceType;
        public readonly string? ReservedInstanceId;
        public readonly ImmutableArray<Outputs.GetInstancesReservedInstanceListResult> ReservedInstanceLists;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstancesResult(
            string? availabilityZone,

            string id,

            string? instanceType,

            string? reservedInstanceId,

            ImmutableArray<Outputs.GetInstancesReservedInstanceListResult> reservedInstanceLists,

            string? resultOutputFile)
        {
            AvailabilityZone = availabilityZone;
            Id = id;
            InstanceType = instanceType;
            ReservedInstanceId = reservedInstanceId;
            ReservedInstanceLists = reservedInstanceLists;
            ResultOutputFile = resultOutputFile;
        }
    }
}