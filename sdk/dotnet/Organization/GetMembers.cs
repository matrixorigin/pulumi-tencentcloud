// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Organization
{
    public static class GetMembers
    {
        public static Task<GetMembersResult> InvokeAsync(GetMembersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMembersResult>("tencentcloud:Organization/getMembers:getMembers", args ?? new GetMembersArgs(), options.WithDefaults());

        public static Output<GetMembersResult> Invoke(GetMembersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMembersResult>("tencentcloud:Organization/getMembers:getMembers", args ?? new GetMembersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMembersArgs : global::Pulumi.InvokeArgs
    {
        [Input("authName")]
        public string? AuthName { get; set; }

        [Input("lang")]
        public string? Lang { get; set; }

        [Input("product")]
        public string? Product { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("searchKey")]
        public string? SearchKey { get; set; }

        public GetMembersArgs()
        {
        }
        public static new GetMembersArgs Empty => new GetMembersArgs();
    }

    public sealed class GetMembersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("authName")]
        public Input<string>? AuthName { get; set; }

        [Input("lang")]
        public Input<string>? Lang { get; set; }

        [Input("product")]
        public Input<string>? Product { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("searchKey")]
        public Input<string>? SearchKey { get; set; }

        public GetMembersInvokeArgs()
        {
        }
        public static new GetMembersInvokeArgs Empty => new GetMembersInvokeArgs();
    }


    [OutputType]
    public sealed class GetMembersResult
    {
        public readonly string? AuthName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetMembersItemResult> Items;
        public readonly string? Lang;
        public readonly string? Product;
        public readonly string? ResultOutputFile;
        public readonly string? SearchKey;

        [OutputConstructor]
        private GetMembersResult(
            string? authName,

            string id,

            ImmutableArray<Outputs.GetMembersItemResult> items,

            string? lang,

            string? product,

            string? resultOutputFile,

            string? searchKey)
        {
            AuthName = authName;
            Id = id;
            Items = items;
            Lang = lang;
            Product = product;
            ResultOutputFile = resultOutputFile;
            SearchKey = searchKey;
        }
    }
}
