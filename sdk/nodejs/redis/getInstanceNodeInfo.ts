// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getInstanceNodeInfo(args: GetInstanceNodeInfoArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceNodeInfoResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Redis/getInstanceNodeInfo:getInstanceNodeInfo", {
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceNodeInfo.
 */
export interface GetInstanceNodeInfoArgs {
    instanceId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstanceNodeInfo.
 */
export interface GetInstanceNodeInfoResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly proxies: outputs.Redis.GetInstanceNodeInfoProxy[];
    readonly proxyCount: number;
    readonly redis: outputs.Redis.GetInstanceNodeInfoRedi[];
    readonly redisCount: number;
    readonly resultOutputFile?: string;
}
export function getInstanceNodeInfoOutput(args: GetInstanceNodeInfoOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceNodeInfoResult> {
    return pulumi.output(args).apply((a: any) => getInstanceNodeInfo(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceNodeInfo.
 */
export interface GetInstanceNodeInfoOutputArgs {
    instanceId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
