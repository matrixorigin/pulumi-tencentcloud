// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Css
{
    public static class GetMonitorReport
    {
        public static Task<GetMonitorReportResult> InvokeAsync(GetMonitorReportArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMonitorReportResult>("tencentcloud:Css/getMonitorReport:getMonitorReport", args ?? new GetMonitorReportArgs(), options.WithDefaults());

        public static Output<GetMonitorReportResult> Invoke(GetMonitorReportInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitorReportResult>("tencentcloud:Css/getMonitorReport:getMonitorReport", args ?? new GetMonitorReportInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMonitorReportArgs : global::Pulumi.InvokeArgs
    {
        [Input("monitorId", required: true)]
        public string MonitorId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetMonitorReportArgs()
        {
        }
        public static new GetMonitorReportArgs Empty => new GetMonitorReportArgs();
    }

    public sealed class GetMonitorReportInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("monitorId", required: true)]
        public Input<string> MonitorId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetMonitorReportInvokeArgs()
        {
        }
        public static new GetMonitorReportInvokeArgs Empty => new GetMonitorReportInvokeArgs();
    }


    [OutputType]
    public sealed class GetMonitorReportResult
    {
        public readonly ImmutableArray<Outputs.GetMonitorReportDiagnoseResultResult> DiagnoseResults;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string MonitorId;
        public readonly ImmutableArray<Outputs.GetMonitorReportMpsResultResult> MpsResults;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetMonitorReportResult(
            ImmutableArray<Outputs.GetMonitorReportDiagnoseResultResult> diagnoseResults,

            string id,

            string monitorId,

            ImmutableArray<Outputs.GetMonitorReportMpsResultResult> mpsResults,

            string? resultOutputFile)
        {
            DiagnoseResults = diagnoseResults;
            Id = id;
            MonitorId = monitorId;
            MpsResults = mpsResults;
            ResultOutputFile = resultOutputFile;
        }
    }
}