// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getDescribeEngineUsageInfo(args: GetDescribeEngineUsageInfoArgs, opts?: pulumi.InvokeOptions): Promise<GetDescribeEngineUsageInfoResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dlc/getDescribeEngineUsageInfo:getDescribeEngineUsageInfo", {
        "dataEngineId": args.dataEngineId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDescribeEngineUsageInfo.
 */
export interface GetDescribeEngineUsageInfoArgs {
    dataEngineId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDescribeEngineUsageInfo.
 */
export interface GetDescribeEngineUsageInfoResult {
    readonly available: number;
    readonly dataEngineId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly used: number;
}
export function getDescribeEngineUsageInfoOutput(args: GetDescribeEngineUsageInfoOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDescribeEngineUsageInfoResult> {
    return pulumi.output(args).apply((a: any) => getDescribeEngineUsageInfo(a, opts))
}

/**
 * A collection of arguments for invoking getDescribeEngineUsageInfo.
 */
export interface GetDescribeEngineUsageInfoOutputArgs {
    dataEngineId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}