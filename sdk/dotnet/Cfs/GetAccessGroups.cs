// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cfs
{
    public static class GetAccessGroups
    {
        public static Task<GetAccessGroupsResult> InvokeAsync(GetAccessGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessGroupsResult>("tencentcloud:Cfs/getAccessGroups:getAccessGroups", args ?? new GetAccessGroupsArgs(), options.WithDefaults());

        public static Output<GetAccessGroupsResult> Invoke(GetAccessGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessGroupsResult>("tencentcloud:Cfs/getAccessGroups:getAccessGroups", args ?? new GetAccessGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessGroupsArgs : global::Pulumi.InvokeArgs
    {
        [Input("accessGroupId")]
        public string? AccessGroupId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetAccessGroupsArgs()
        {
        }
        public static new GetAccessGroupsArgs Empty => new GetAccessGroupsArgs();
    }

    public sealed class GetAccessGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("accessGroupId")]
        public Input<string>? AccessGroupId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetAccessGroupsInvokeArgs()
        {
        }
        public static new GetAccessGroupsInvokeArgs Empty => new GetAccessGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccessGroupsResult
    {
        public readonly string? AccessGroupId;
        public readonly ImmutableArray<Outputs.GetAccessGroupsAccessGroupListResult> AccessGroupLists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetAccessGroupsResult(
            string? accessGroupId,

            ImmutableArray<Outputs.GetAccessGroupsAccessGroupListResult> accessGroupLists,

            string id,

            string? name,

            string? resultOutputFile)
        {
            AccessGroupId = accessGroupId;
            AccessGroupLists = accessGroupLists;
            Id = id;
            Name = name;
            ResultOutputFile = resultOutputFile;
        }
    }
}