// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getInfo(args?: GetInfoArgs, opts?: pulumi.InvokeOptions): Promise<GetInfoResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:User/getInfo:getInfo", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInfo.
 */
export interface GetInfoArgs {
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInfo.
 */
export interface GetInfoResult {
    readonly appId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly ownerUin: string;
    readonly resultOutputFile?: string;
    readonly uin: string;
}
export function getInfoOutput(args?: GetInfoOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInfoResult> {
    return pulumi.output(args).apply((a: any) => getInfo(a, opts))
}

/**
 * A collection of arguments for invoking getInfo.
 */
export interface GetInfoOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
}