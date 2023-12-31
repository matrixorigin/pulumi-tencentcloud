// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getGroupMemberships(args?: GetGroupMembershipsArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupMembershipsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cam/getGroupMemberships:getGroupMemberships", {
        "groupId": args.groupId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupMemberships.
 */
export interface GetGroupMembershipsArgs {
    groupId?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getGroupMemberships.
 */
export interface GetGroupMembershipsResult {
    readonly groupId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly membershipLists: outputs.Cam.GetGroupMembershipsMembershipList[];
    readonly resultOutputFile?: string;
}
export function getGroupMembershipsOutput(args?: GetGroupMembershipsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupMembershipsResult> {
    return pulumi.output(args).apply((a: any) => getGroupMemberships(a, opts))
}

/**
 * A collection of arguments for invoking getGroupMemberships.
 */
export interface GetGroupMembershipsOutputArgs {
    groupId?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
