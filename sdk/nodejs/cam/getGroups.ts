// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getGroups(args?: GetGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cam/getGroups:getGroups", {
        "groupId": args.groupId,
        "name": args.name,
        "remark": args.remark,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroups.
 */
export interface GetGroupsArgs {
    groupId?: string;
    name?: string;
    remark?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getGroups.
 */
export interface GetGroupsResult {
    readonly groupId?: string;
    readonly groupLists: outputs.Cam.GetGroupsGroupList[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    readonly remark?: string;
    readonly resultOutputFile?: string;
}
export function getGroupsOutput(args?: GetGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupsResult> {
    return pulumi.output(args).apply((a: any) => getGroups(a, opts))
}

/**
 * A collection of arguments for invoking getGroups.
 */
export interface GetGroupsOutputArgs {
    groupId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    remark?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
