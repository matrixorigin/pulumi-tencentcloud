// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getReadonlyGroups(args?: GetReadonlyGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetReadonlyGroupsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Sqlserver/getReadonlyGroups:getReadonlyGroups", {
        "masterInstanceId": args.masterInstanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getReadonlyGroups.
 */
export interface GetReadonlyGroupsArgs {
    masterInstanceId?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getReadonlyGroups.
 */
export interface GetReadonlyGroupsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.Sqlserver.GetReadonlyGroupsList[];
    readonly masterInstanceId?: string;
    readonly resultOutputFile?: string;
}
export function getReadonlyGroupsOutput(args?: GetReadonlyGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReadonlyGroupsResult> {
    return pulumi.output(args).apply((a: any) => getReadonlyGroups(a, opts))
}

/**
 * A collection of arguments for invoking getReadonlyGroups.
 */
export interface GetReadonlyGroupsOutputArgs {
    masterInstanceId?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
