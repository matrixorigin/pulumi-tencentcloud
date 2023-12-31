// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getMonitorReport(args: GetMonitorReportArgs, opts?: pulumi.InvokeOptions): Promise<GetMonitorReportResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Css/getMonitorReport:getMonitorReport", {
        "monitorId": args.monitorId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getMonitorReport.
 */
export interface GetMonitorReportArgs {
    monitorId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getMonitorReport.
 */
export interface GetMonitorReportResult {
    readonly diagnoseResults: outputs.Css.GetMonitorReportDiagnoseResult[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly monitorId: string;
    readonly mpsResults: outputs.Css.GetMonitorReportMpsResult[];
    readonly resultOutputFile?: string;
}
export function getMonitorReportOutput(args: GetMonitorReportOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMonitorReportResult> {
    return pulumi.output(args).apply((a: any) => getMonitorReport(a, opts))
}

/**
 * A collection of arguments for invoking getMonitorReport.
 */
export interface GetMonitorReportOutputArgs {
    monitorId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
