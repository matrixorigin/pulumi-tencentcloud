// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cynosdb
{
    public static class GetAccountAllGrantPrivileges
    {
        public static Task<GetAccountAllGrantPrivilegesResult> InvokeAsync(GetAccountAllGrantPrivilegesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountAllGrantPrivilegesResult>("tencentcloud:Cynosdb/getAccountAllGrantPrivileges:getAccountAllGrantPrivileges", args ?? new GetAccountAllGrantPrivilegesArgs(), options.WithDefaults());

        public static Output<GetAccountAllGrantPrivilegesResult> Invoke(GetAccountAllGrantPrivilegesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountAllGrantPrivilegesResult>("tencentcloud:Cynosdb/getAccountAllGrantPrivileges:getAccountAllGrantPrivileges", args ?? new GetAccountAllGrantPrivilegesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountAllGrantPrivilegesArgs : global::Pulumi.InvokeArgs
    {
        [Input("account", required: true)]
        public Inputs.GetAccountAllGrantPrivilegesAccountArgs Account { get; set; } = null!;

        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetAccountAllGrantPrivilegesArgs()
        {
        }
        public static new GetAccountAllGrantPrivilegesArgs Empty => new GetAccountAllGrantPrivilegesArgs();
    }

    public sealed class GetAccountAllGrantPrivilegesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("account", required: true)]
        public Input<Inputs.GetAccountAllGrantPrivilegesAccountInputArgs> Account { get; set; } = null!;

        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetAccountAllGrantPrivilegesInvokeArgs()
        {
        }
        public static new GetAccountAllGrantPrivilegesInvokeArgs Empty => new GetAccountAllGrantPrivilegesInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccountAllGrantPrivilegesResult
    {
        public readonly Outputs.GetAccountAllGrantPrivilegesAccountResult Account;
        public readonly string ClusterId;
        public readonly ImmutableArray<Outputs.GetAccountAllGrantPrivilegesDatabasePrivilegeResult> DatabasePrivileges;
        public readonly ImmutableArray<string> GlobalPrivileges;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> PrivilegeStatements;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<Outputs.GetAccountAllGrantPrivilegesTablePrivilegeResult> TablePrivileges;

        [OutputConstructor]
        private GetAccountAllGrantPrivilegesResult(
            Outputs.GetAccountAllGrantPrivilegesAccountResult account,

            string clusterId,

            ImmutableArray<Outputs.GetAccountAllGrantPrivilegesDatabasePrivilegeResult> databasePrivileges,

            ImmutableArray<string> globalPrivileges,

            string id,

            ImmutableArray<string> privilegeStatements,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetAccountAllGrantPrivilegesTablePrivilegeResult> tablePrivileges)
        {
            Account = account;
            ClusterId = clusterId;
            DatabasePrivileges = databasePrivileges;
            GlobalPrivileges = globalPrivileges;
            Id = id;
            PrivilegeStatements = privilegeStatements;
            ResultOutputFile = resultOutputFile;
            TablePrivileges = tablePrivileges;
        }
    }
}