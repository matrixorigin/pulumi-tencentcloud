// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getProxyGroupStatistics(args: GetProxyGroupStatisticsArgs, opts?: pulumi.InvokeOptions): Promise<GetProxyGroupStatisticsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Gaap/getProxyGroupStatistics:getProxyGroupStatistics", {
        "endTime": args.endTime,
        "granularity": args.granularity,
        "groupId": args.groupId,
        "metricNames": args.metricNames,
        "resultOutputFile": args.resultOutputFile,
        "startTime": args.startTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getProxyGroupStatistics.
 */
export interface GetProxyGroupStatisticsArgs {
    endTime: string;
    granularity: number;
    groupId: string;
    metricNames: string[];
    resultOutputFile?: string;
    startTime: string;
}

/**
 * A collection of values returned by getProxyGroupStatistics.
 */
export interface GetProxyGroupStatisticsResult {
    readonly endTime: string;
    readonly granularity: number;
    readonly groupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly metricNames: string[];
    readonly resultOutputFile?: string;
    readonly startTime: string;
    readonly statisticsDatas: outputs.Gaap.GetProxyGroupStatisticsStatisticsData[];
}
export function getProxyGroupStatisticsOutput(args: GetProxyGroupStatisticsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProxyGroupStatisticsResult> {
    return pulumi.output(args).apply((a: any) => getProxyGroupStatistics(a, opts))
}

/**
 * A collection of arguments for invoking getProxyGroupStatistics.
 */
export interface GetProxyGroupStatisticsOutputArgs {
    endTime: pulumi.Input<string>;
    granularity: pulumi.Input<number>;
    groupId: pulumi.Input<string>;
    metricNames: pulumi.Input<pulumi.Input<string>[]>;
    resultOutputFile?: pulumi.Input<string>;
    startTime: pulumi.Input<string>;
}