// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getSendTasks(args?: GetSendTasksArgs, opts?: pulumi.InvokeOptions): Promise<GetSendTasksResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Ses/getSendTasks:getSendTasks", {
        "receiverId": args.receiverId,
        "resultOutputFile": args.resultOutputFile,
        "status": args.status,
        "taskType": args.taskType,
    }, opts);
}

/**
 * A collection of arguments for invoking getSendTasks.
 */
export interface GetSendTasksArgs {
    receiverId?: number;
    resultOutputFile?: string;
    status?: number;
    taskType?: number;
}

/**
 * A collection of values returned by getSendTasks.
 */
export interface GetSendTasksResult {
    readonly datas: outputs.Ses.GetSendTasksData[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly receiverId?: number;
    readonly resultOutputFile?: string;
    readonly status?: number;
    readonly taskType?: number;
}
export function getSendTasksOutput(args?: GetSendTasksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSendTasksResult> {
    return pulumi.output(args).apply((a: any) => getSendTasks(a, opts))
}

/**
 * A collection of arguments for invoking getSendTasks.
 */
export interface GetSendTasksOutputArgs {
    receiverId?: pulumi.Input<number>;
    resultOutputFile?: pulumi.Input<string>;
    status?: pulumi.Input<number>;
    taskType?: pulumi.Input<number>;
}