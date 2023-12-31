// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getSyncJobs(args?: GetSyncJobsArgs, opts?: pulumi.InvokeOptions): Promise<GetSyncJobsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dts/getSyncJobs:getSyncJobs", {
        "jobId": args.jobId,
        "jobName": args.jobName,
        "jobType": args.jobType,
        "order": args.order,
        "orderSeq": args.orderSeq,
        "payMode": args.payMode,
        "resultOutputFile": args.resultOutputFile,
        "runMode": args.runMode,
        "statuses": args.statuses,
        "tagFilters": args.tagFilters,
    }, opts);
}

/**
 * A collection of arguments for invoking getSyncJobs.
 */
export interface GetSyncJobsArgs {
    jobId?: string;
    jobName?: string;
    jobType?: string;
    order?: string;
    orderSeq?: string;
    payMode?: string;
    resultOutputFile?: string;
    runMode?: string;
    statuses?: string[];
    tagFilters?: inputs.Dts.GetSyncJobsTagFilter[];
}

/**
 * A collection of values returned by getSyncJobs.
 */
export interface GetSyncJobsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly jobId?: string;
    readonly jobName?: string;
    readonly jobType?: string;
    readonly lists: outputs.Dts.GetSyncJobsList[];
    readonly order?: string;
    readonly orderSeq?: string;
    readonly payMode?: string;
    readonly resultOutputFile?: string;
    readonly runMode?: string;
    readonly statuses?: string[];
    readonly tagFilters?: outputs.Dts.GetSyncJobsTagFilter[];
}
export function getSyncJobsOutput(args?: GetSyncJobsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSyncJobsResult> {
    return pulumi.output(args).apply((a: any) => getSyncJobs(a, opts))
}

/**
 * A collection of arguments for invoking getSyncJobs.
 */
export interface GetSyncJobsOutputArgs {
    jobId?: pulumi.Input<string>;
    jobName?: pulumi.Input<string>;
    jobType?: pulumi.Input<string>;
    order?: pulumi.Input<string>;
    orderSeq?: pulumi.Input<string>;
    payMode?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    runMode?: pulumi.Input<string>;
    statuses?: pulumi.Input<pulumi.Input<string>[]>;
    tagFilters?: pulumi.Input<pulumi.Input<inputs.Dts.GetSyncJobsTagFilterArgs>[]>;
}
