// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Waf
{
    public static class GetAttackLogList
    {
        public static Task<GetAttackLogListResult> InvokeAsync(GetAttackLogListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAttackLogListResult>("tencentcloud:Waf/getAttackLogList:getAttackLogList", args ?? new GetAttackLogListArgs(), options.WithDefaults());

        public static Output<GetAttackLogListResult> Invoke(GetAttackLogListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAttackLogListResult>("tencentcloud:Waf/getAttackLogList:getAttackLogList", args ?? new GetAttackLogListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAttackLogListArgs : global::Pulumi.InvokeArgs
    {
        [Input("domain", required: true)]
        public string Domain { get; set; } = null!;

        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        [Input("page")]
        public int? Page { get; set; }

        [Input("queryCount")]
        public int? QueryCount { get; set; }

        [Input("queryString", required: true)]
        public string QueryString { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("sort")]
        public string? Sort { get; set; }

        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        public GetAttackLogListArgs()
        {
        }
        public static new GetAttackLogListArgs Empty => new GetAttackLogListArgs();
    }

    public sealed class GetAttackLogListInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        [Input("page")]
        public Input<int>? Page { get; set; }

        [Input("queryCount")]
        public Input<int>? QueryCount { get; set; }

        [Input("queryString", required: true)]
        public Input<string> QueryString { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("sort")]
        public Input<string>? Sort { get; set; }

        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public GetAttackLogListInvokeArgs()
        {
        }
        public static new GetAttackLogListInvokeArgs Empty => new GetAttackLogListInvokeArgs();
    }


    [OutputType]
    public sealed class GetAttackLogListResult
    {
        public readonly ImmutableArray<Outputs.GetAttackLogListDataResult> Datas;
        public readonly string Domain;
        public readonly string EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? Page;
        public readonly int? QueryCount;
        public readonly string QueryString;
        public readonly string? ResultOutputFile;
        public readonly string? Sort;
        public readonly string StartTime;

        [OutputConstructor]
        private GetAttackLogListResult(
            ImmutableArray<Outputs.GetAttackLogListDataResult> datas,

            string domain,

            string endTime,

            string id,

            int? page,

            int? queryCount,

            string queryString,

            string? resultOutputFile,

            string? sort,

            string startTime)
        {
            Datas = datas;
            Domain = domain;
            EndTime = endTime;
            Id = id;
            Page = page;
            QueryCount = queryCount;
            QueryString = queryString;
            ResultOutputFile = resultOutputFile;
            Sort = sort;
            StartTime = startTime;
        }
    }
}