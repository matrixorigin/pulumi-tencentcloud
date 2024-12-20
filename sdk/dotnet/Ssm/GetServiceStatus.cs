// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssm
{
    public static class GetServiceStatus
    {
        public static Task<GetServiceStatusResult> InvokeAsync(GetServiceStatusArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceStatusResult>("tencentcloud:Ssm/getServiceStatus:getServiceStatus", args ?? new GetServiceStatusArgs(), options.WithDefaults());

        public static Output<GetServiceStatusResult> Invoke(GetServiceStatusInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceStatusResult>("tencentcloud:Ssm/getServiceStatus:getServiceStatus", args ?? new GetServiceStatusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceStatusArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetServiceStatusArgs()
        {
        }
        public static new GetServiceStatusArgs Empty => new GetServiceStatusArgs();
    }

    public sealed class GetServiceStatusInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetServiceStatusInvokeArgs()
        {
        }
        public static new GetServiceStatusInvokeArgs Empty => new GetServiceStatusInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceStatusResult
    {
        public readonly bool AccessKeyEscrowEnabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InvalidType;
        public readonly string? ResultOutputFile;
        public readonly bool ServiceEnabled;

        [OutputConstructor]
        private GetServiceStatusResult(
            bool accessKeyEscrowEnabled,

            string id,

            int invalidType,

            string? resultOutputFile,

            bool serviceEnabled)
        {
            AccessKeyEscrowEnabled = accessKeyEscrowEnabled;
            Id = id;
            InvalidType = invalidType;
            ResultOutputFile = resultOutputFile;
            ServiceEnabled = serviceEnabled;
        }
    }
}
