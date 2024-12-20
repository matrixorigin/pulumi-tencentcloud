// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssl
{
    public static class GetDescribeManagers
    {
        public static Task<GetDescribeManagersResult> InvokeAsync(GetDescribeManagersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDescribeManagersResult>("tencentcloud:Ssl/getDescribeManagers:getDescribeManagers", args ?? new GetDescribeManagersArgs(), options.WithDefaults());

        public static Output<GetDescribeManagersResult> Invoke(GetDescribeManagersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDescribeManagersResult>("tencentcloud:Ssl/getDescribeManagers:getDescribeManagers", args ?? new GetDescribeManagersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeManagersArgs : global::Pulumi.InvokeArgs
    {
        [Input("companyId", required: true)]
        public int CompanyId { get; set; }

        [Input("managerMail")]
        public string? ManagerMail { get; set; }

        [Input("managerName")]
        public string? ManagerName { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("searchKey")]
        public string? SearchKey { get; set; }

        [Input("status")]
        public string? Status { get; set; }

        public GetDescribeManagersArgs()
        {
        }
        public static new GetDescribeManagersArgs Empty => new GetDescribeManagersArgs();
    }

    public sealed class GetDescribeManagersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("companyId", required: true)]
        public Input<int> CompanyId { get; set; } = null!;

        [Input("managerMail")]
        public Input<string>? ManagerMail { get; set; }

        [Input("managerName")]
        public Input<string>? ManagerName { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("searchKey")]
        public Input<string>? SearchKey { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetDescribeManagersInvokeArgs()
        {
        }
        public static new GetDescribeManagersInvokeArgs Empty => new GetDescribeManagersInvokeArgs();
    }


    [OutputType]
    public sealed class GetDescribeManagersResult
    {
        public readonly int CompanyId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ManagerMail;
        public readonly string? ManagerName;
        public readonly ImmutableArray<Outputs.GetDescribeManagersManagerResult> Managers;
        public readonly string? ResultOutputFile;
        public readonly string? SearchKey;
        public readonly string? Status;

        [OutputConstructor]
        private GetDescribeManagersResult(
            int companyId,

            string id,

            string? managerMail,

            string? managerName,

            ImmutableArray<Outputs.GetDescribeManagersManagerResult> managers,

            string? resultOutputFile,

            string? searchKey,

            string? status)
        {
            CompanyId = companyId;
            Id = id;
            ManagerMail = managerMail;
            ManagerName = managerName;
            Managers = managers;
            ResultOutputFile = resultOutputFile;
            SearchKey = searchKey;
            Status = status;
        }
    }
}
