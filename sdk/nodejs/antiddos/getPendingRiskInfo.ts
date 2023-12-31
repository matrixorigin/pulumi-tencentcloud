// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getPendingRiskInfo(args?: GetPendingRiskInfoArgs, opts?: pulumi.InvokeOptions): Promise<GetPendingRiskInfoResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Antiddos/getPendingRiskInfo:getPendingRiskInfo", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getPendingRiskInfo.
 */
export interface GetPendingRiskInfoArgs {
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getPendingRiskInfo.
 */
export interface GetPendingRiskInfoResult {
    readonly attackingCount: number;
    readonly blockingCount: number;
    readonly expiredCount: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isPaidUsr: boolean;
    readonly resultOutputFile?: string;
    readonly total: number;
}
export function getPendingRiskInfoOutput(args?: GetPendingRiskInfoOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPendingRiskInfoResult> {
    return pulumi.output(args).apply((a: any) => getPendingRiskInfo(a, opts))
}

/**
 * A collection of arguments for invoking getPendingRiskInfo.
 */
export interface GetPendingRiskInfoOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
}
