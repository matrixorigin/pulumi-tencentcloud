// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Organization
{
    public static class GetOrgFinancialByMember
    {
        public static Task<GetOrgFinancialByMemberResult> InvokeAsync(GetOrgFinancialByMemberArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrgFinancialByMemberResult>("tencentcloud:Organization/getOrgFinancialByMember:getOrgFinancialByMember", args ?? new GetOrgFinancialByMemberArgs(), options.WithDefaults());

        public static Output<GetOrgFinancialByMemberResult> Invoke(GetOrgFinancialByMemberInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrgFinancialByMemberResult>("tencentcloud:Organization/getOrgFinancialByMember:getOrgFinancialByMember", args ?? new GetOrgFinancialByMemberInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrgFinancialByMemberArgs : global::Pulumi.InvokeArgs
    {
        [Input("endMonth")]
        public string? EndMonth { get; set; }

        [Input("memberUins")]
        private List<int>? _memberUins;
        public List<int> MemberUins
        {
            get => _memberUins ?? (_memberUins = new List<int>());
            set => _memberUins = value;
        }

        [Input("month", required: true)]
        public string Month { get; set; } = null!;

        [Input("productCodes")]
        private List<string>? _productCodes;
        public List<string> ProductCodes
        {
            get => _productCodes ?? (_productCodes = new List<string>());
            set => _productCodes = value;
        }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetOrgFinancialByMemberArgs()
        {
        }
        public static new GetOrgFinancialByMemberArgs Empty => new GetOrgFinancialByMemberArgs();
    }

    public sealed class GetOrgFinancialByMemberInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("endMonth")]
        public Input<string>? EndMonth { get; set; }

        [Input("memberUins")]
        private InputList<int>? _memberUins;
        public InputList<int> MemberUins
        {
            get => _memberUins ?? (_memberUins = new InputList<int>());
            set => _memberUins = value;
        }

        [Input("month", required: true)]
        public Input<string> Month { get; set; } = null!;

        [Input("productCodes")]
        private InputList<string>? _productCodes;
        public InputList<string> ProductCodes
        {
            get => _productCodes ?? (_productCodes = new InputList<string>());
            set => _productCodes = value;
        }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetOrgFinancialByMemberInvokeArgs()
        {
        }
        public static new GetOrgFinancialByMemberInvokeArgs Empty => new GetOrgFinancialByMemberInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrgFinancialByMemberResult
    {
        public readonly string? EndMonth;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetOrgFinancialByMemberItemResult> Items;
        public readonly ImmutableArray<int> MemberUins;
        public readonly string Month;
        public readonly ImmutableArray<string> ProductCodes;
        public readonly string? ResultOutputFile;
        public readonly double TotalCost;

        [OutputConstructor]
        private GetOrgFinancialByMemberResult(
            string? endMonth,

            string id,

            ImmutableArray<Outputs.GetOrgFinancialByMemberItemResult> items,

            ImmutableArray<int> memberUins,

            string month,

            ImmutableArray<string> productCodes,

            string? resultOutputFile,

            double totalCost)
        {
            EndMonth = endMonth;
            Id = id;
            Items = items;
            MemberUins = memberUins;
            Month = month;
            ProductCodes = productCodes;
            ResultOutputFile = resultOutputFile;
            TotalCost = totalCost;
        }
    }
}