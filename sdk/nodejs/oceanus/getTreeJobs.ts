// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getTreeJobs(args?: GetTreeJobsArgs, opts?: pulumi.InvokeOptions): Promise<GetTreeJobsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Oceanus/getTreeJobs:getTreeJobs", {
        "filters": args.filters,
        "resultOutputFile": args.resultOutputFile,
        "workSpaceId": args.workSpaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTreeJobs.
 */
export interface GetTreeJobsArgs {
    filters?: inputs.Oceanus.GetTreeJobsFilter[];
    resultOutputFile?: string;
    workSpaceId?: string;
}

/**
 * A collection of values returned by getTreeJobs.
 */
export interface GetTreeJobsResult {
    readonly filters?: outputs.Oceanus.GetTreeJobsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly treeInfos: outputs.Oceanus.GetTreeJobsTreeInfo[];
    readonly workSpaceId?: string;
}
export function getTreeJobsOutput(args?: GetTreeJobsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTreeJobsResult> {
    return pulumi.output(args).apply((a: any) => getTreeJobs(a, opts))
}

/**
 * A collection of arguments for invoking getTreeJobs.
 */
export interface GetTreeJobsOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.Oceanus.GetTreeJobsFilterArgs>[]>;
    resultOutputFile?: pulumi.Input<string>;
    workSpaceId?: pulumi.Input<string>;
}
