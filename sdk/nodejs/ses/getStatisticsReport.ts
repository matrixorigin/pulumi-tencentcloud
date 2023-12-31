// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getStatisticsReport(args: GetStatisticsReportArgs, opts?: pulumi.InvokeOptions): Promise<GetStatisticsReportResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Ses/getStatisticsReport:getStatisticsReport", {
        "domain": args.domain,
        "endDate": args.endDate,
        "receivingMailboxType": args.receivingMailboxType,
        "resultOutputFile": args.resultOutputFile,
        "startDate": args.startDate,
    }, opts);
}

/**
 * A collection of arguments for invoking getStatisticsReport.
 */
export interface GetStatisticsReportArgs {
    domain?: string;
    endDate: string;
    receivingMailboxType?: string;
    resultOutputFile?: string;
    startDate: string;
}

/**
 * A collection of values returned by getStatisticsReport.
 */
export interface GetStatisticsReportResult {
    readonly dailyVolumes: outputs.Ses.GetStatisticsReportDailyVolume[];
    readonly domain?: string;
    readonly endDate: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly overallVolumes: outputs.Ses.GetStatisticsReportOverallVolume[];
    readonly receivingMailboxType?: string;
    readonly resultOutputFile?: string;
    readonly startDate: string;
}
export function getStatisticsReportOutput(args: GetStatisticsReportOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStatisticsReportResult> {
    return pulumi.output(args).apply((a: any) => getStatisticsReport(a, opts))
}

/**
 * A collection of arguments for invoking getStatisticsReport.
 */
export interface GetStatisticsReportOutputArgs {
    domain?: pulumi.Input<string>;
    endDate: pulumi.Input<string>;
    receivingMailboxType?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    startDate: pulumi.Input<string>;
}
