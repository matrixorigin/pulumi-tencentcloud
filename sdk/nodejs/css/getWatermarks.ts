// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getWatermarks(args?: GetWatermarksArgs, opts?: pulumi.InvokeOptions): Promise<GetWatermarksResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Css/getWatermarks:getWatermarks", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getWatermarks.
 */
export interface GetWatermarksArgs {
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getWatermarks.
 */
export interface GetWatermarksResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly watermarkLists: outputs.Css.GetWatermarksWatermarkList[];
}
export function getWatermarksOutput(args?: GetWatermarksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWatermarksResult> {
    return pulumi.output(args).apply((a: any) => getWatermarks(a, opts))
}

/**
 * A collection of arguments for invoking getWatermarks.
 */
export interface GetWatermarksOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
}