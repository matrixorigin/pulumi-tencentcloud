// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getSchedules(args?: GetSchedulesArgs, opts?: pulumi.InvokeOptions): Promise<GetSchedulesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Mps/getSchedules:getSchedules", {
        "resultOutputFile": args.resultOutputFile,
        "scheduleIds": args.scheduleIds,
        "status": args.status,
        "triggerType": args.triggerType,
    }, opts);
}

/**
 * A collection of arguments for invoking getSchedules.
 */
export interface GetSchedulesArgs {
    resultOutputFile?: string;
    scheduleIds?: number[];
    status?: string;
    triggerType?: string;
}

/**
 * A collection of values returned by getSchedules.
 */
export interface GetSchedulesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly scheduleIds?: number[];
    readonly scheduleInfoSets: outputs.Mps.GetSchedulesScheduleInfoSet[];
    readonly status?: string;
    readonly triggerType?: string;
}
export function getSchedulesOutput(args?: GetSchedulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSchedulesResult> {
    return pulumi.output(args).apply((a: any) => getSchedules(a, opts))
}

/**
 * A collection of arguments for invoking getSchedules.
 */
export interface GetSchedulesOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
    scheduleIds?: pulumi.Input<pulumi.Input<number>[]>;
    status?: pulumi.Input<string>;
    triggerType?: pulumi.Input<string>;
}
