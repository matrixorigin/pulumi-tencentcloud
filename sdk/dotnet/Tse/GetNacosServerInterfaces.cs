// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tse
{
    public static class GetNacosServerInterfaces
    {
        public static Task<GetNacosServerInterfacesResult> InvokeAsync(GetNacosServerInterfacesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNacosServerInterfacesResult>("tencentcloud:Tse/getNacosServerInterfaces:getNacosServerInterfaces", args ?? new GetNacosServerInterfacesArgs(), options.WithDefaults());

        public static Output<GetNacosServerInterfacesResult> Invoke(GetNacosServerInterfacesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNacosServerInterfacesResult>("tencentcloud:Tse/getNacosServerInterfaces:getNacosServerInterfaces", args ?? new GetNacosServerInterfacesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNacosServerInterfacesArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetNacosServerInterfacesArgs()
        {
        }
        public static new GetNacosServerInterfacesArgs Empty => new GetNacosServerInterfacesArgs();
    }

    public sealed class GetNacosServerInterfacesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetNacosServerInterfacesInvokeArgs()
        {
        }
        public static new GetNacosServerInterfacesInvokeArgs Empty => new GetNacosServerInterfacesInvokeArgs();
    }


    [OutputType]
    public sealed class GetNacosServerInterfacesResult
    {
        public readonly ImmutableArray<Outputs.GetNacosServerInterfacesContentResult> Contents;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetNacosServerInterfacesResult(
            ImmutableArray<Outputs.GetNacosServerInterfacesContentResult> contents,

            string id,

            string? instanceId,

            string? resultOutputFile)
        {
            Contents = contents;
            Id = id;
            InstanceId = instanceId;
            ResultOutputFile = resultOutputFile;
        }
    }
}