// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getTasks(args: GetTasksArgs, opts?: pulumi.InvokeOptions): Promise<GetTasksResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Mps/getTasks:getTasks", {
        "limit": args.limit,
        "resultOutputFile": args.resultOutputFile,
        "scrollToken": args.scrollToken,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getTasks.
 */
export interface GetTasksArgs {
    limit?: number;
    resultOutputFile?: string;
    scrollToken?: string;
    status: string;
}

/**
 * A collection of values returned by getTasks.
 */
export interface GetTasksResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly limit?: number;
    readonly resultOutputFile?: string;
    readonly scrollToken: string;
    readonly status: string;
    readonly taskSets: outputs.Mps.GetTasksTaskSet[];
}
export function getTasksOutput(args: GetTasksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTasksResult> {
    return pulumi.output(args).apply((a: any) => getTasks(a, opts))
}

/**
 * A collection of arguments for invoking getTasks.
 */
export interface GetTasksOutputArgs {
    limit?: pulumi.Input<number>;
    resultOutputFile?: pulumi.Input<string>;
    scrollToken?: pulumi.Input<string>;
    status: pulumi.Input<string>;
}
