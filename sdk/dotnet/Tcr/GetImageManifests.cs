// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcr
{
    public static class GetImageManifests
    {
        public static Task<GetImageManifestsResult> InvokeAsync(GetImageManifestsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetImageManifestsResult>("tencentcloud:Tcr/getImageManifests:getImageManifests", args ?? new GetImageManifestsArgs(), options.WithDefaults());

        public static Output<GetImageManifestsResult> Invoke(GetImageManifestsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetImageManifestsResult>("tencentcloud:Tcr/getImageManifests:getImageManifests", args ?? new GetImageManifestsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImageManifestsArgs : global::Pulumi.InvokeArgs
    {
        [Input("imageVersion", required: true)]
        public string ImageVersion { get; set; } = null!;

        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        [Input("registryId", required: true)]
        public string RegistryId { get; set; } = null!;

        [Input("repositoryName", required: true)]
        public string RepositoryName { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetImageManifestsArgs()
        {
        }
        public static new GetImageManifestsArgs Empty => new GetImageManifestsArgs();
    }

    public sealed class GetImageManifestsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("imageVersion", required: true)]
        public Input<string> ImageVersion { get; set; } = null!;

        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        [Input("repositoryName", required: true)]
        public Input<string> RepositoryName { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetImageManifestsInvokeArgs()
        {
        }
        public static new GetImageManifestsInvokeArgs Empty => new GetImageManifestsInvokeArgs();
    }


    [OutputType]
    public sealed class GetImageManifestsResult
    {
        public readonly string Config;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ImageVersion;
        public readonly string Manifest;
        public readonly string NamespaceName;
        public readonly string RegistryId;
        public readonly string RepositoryName;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetImageManifestsResult(
            string config,

            string id,

            string imageVersion,

            string manifest,

            string namespaceName,

            string registryId,

            string repositoryName,

            string? resultOutputFile)
        {
            Config = config;
            Id = id;
            ImageVersion = imageVersion;
            Manifest = manifest;
            NamespaceName = namespaceName;
            RegistryId = registryId;
            RepositoryName = repositoryName;
            ResultOutputFile = resultOutputFile;
        }
    }
}