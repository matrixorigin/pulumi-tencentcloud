// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getInstanceTaskList(args?: GetInstanceTaskListArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceTaskListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Redis/getInstanceTaskList:getInstanceTaskList", {
        "beginTime": args.beginTime,
        "endTime": args.endTime,
        "instanceId": args.instanceId,
        "instanceName": args.instanceName,
        "operateUins": args.operateUins,
        "projectIds": args.projectIds,
        "resultOutputFile": args.resultOutputFile,
        "results": args.results,
        "taskStatuses": args.taskStatuses,
        "taskTypes": args.taskTypes,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceTaskList.
 */
export interface GetInstanceTaskListArgs {
    beginTime?: string;
    endTime?: string;
    instanceId?: string;
    instanceName?: string;
    operateUins?: string[];
    projectIds?: number[];
    resultOutputFile?: string;
    results?: number[];
    taskStatuses?: number[];
    taskTypes?: string[];
}

/**
 * A collection of values returned by getInstanceTaskList.
 */
export interface GetInstanceTaskListResult {
    readonly beginTime?: string;
    readonly endTime?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId?: string;
    readonly instanceName?: string;
    readonly operateUins?: string[];
    readonly projectIds?: number[];
    readonly resultOutputFile?: string;
    readonly results?: number[];
    readonly taskStatuses?: number[];
    readonly taskTypes?: string[];
    readonly tasks: outputs.Redis.GetInstanceTaskListTask[];
}
export function getInstanceTaskListOutput(args?: GetInstanceTaskListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceTaskListResult> {
    return pulumi.output(args).apply((a: any) => getInstanceTaskList(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceTaskList.
 */
export interface GetInstanceTaskListOutputArgs {
    beginTime?: pulumi.Input<string>;
    endTime?: pulumi.Input<string>;
    instanceId?: pulumi.Input<string>;
    instanceName?: pulumi.Input<string>;
    operateUins?: pulumi.Input<pulumi.Input<string>[]>;
    projectIds?: pulumi.Input<pulumi.Input<number>[]>;
    resultOutputFile?: pulumi.Input<string>;
    results?: pulumi.Input<pulumi.Input<number>[]>;
    taskStatuses?: pulumi.Input<pulumi.Input<number>[]>;
    taskTypes?: pulumi.Input<pulumi.Input<string>[]>;
}
