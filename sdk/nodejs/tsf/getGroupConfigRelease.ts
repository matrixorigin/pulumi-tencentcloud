// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getGroupConfigRelease(args: GetGroupConfigReleaseArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupConfigReleaseResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Tsf/getGroupConfigRelease:getGroupConfigRelease", {
        "groupId": args.groupId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupConfigRelease.
 */
export interface GetGroupConfigReleaseArgs {
    groupId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getGroupConfigRelease.
 */
export interface GetGroupConfigReleaseResult {
    readonly groupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly results: outputs.Tsf.GetGroupConfigReleaseResult[];
}
export function getGroupConfigReleaseOutput(args: GetGroupConfigReleaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupConfigReleaseResult> {
    return pulumi.output(args).apply((a: any) => getGroupConfigRelease(a, opts))
}

/**
 * A collection of arguments for invoking getGroupConfigRelease.
 */
export interface GetGroupConfigReleaseOutputArgs {
    groupId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
