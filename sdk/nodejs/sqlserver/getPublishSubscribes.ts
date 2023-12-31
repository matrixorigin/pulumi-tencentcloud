// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getPublishSubscribes(args: GetPublishSubscribesArgs, opts?: pulumi.InvokeOptions): Promise<GetPublishSubscribesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Sqlserver/getPublishSubscribes:getPublishSubscribes", {
        "instanceId": args.instanceId,
        "pubOrSubInstanceId": args.pubOrSubInstanceId,
        "pubOrSubInstanceIp": args.pubOrSubInstanceIp,
        "publishDatabase": args.publishDatabase,
        "publishSubscribeId": args.publishSubscribeId,
        "publishSubscribeName": args.publishSubscribeName,
        "resultOutputFile": args.resultOutputFile,
        "subscribeDatabase": args.subscribeDatabase,
    }, opts);
}

/**
 * A collection of arguments for invoking getPublishSubscribes.
 */
export interface GetPublishSubscribesArgs {
    instanceId: string;
    pubOrSubInstanceId?: string;
    pubOrSubInstanceIp?: string;
    publishDatabase?: string;
    publishSubscribeId?: number;
    publishSubscribeName?: string;
    resultOutputFile?: string;
    subscribeDatabase?: string;
}

/**
 * A collection of values returned by getPublishSubscribes.
 */
export interface GetPublishSubscribesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly pubOrSubInstanceId?: string;
    readonly pubOrSubInstanceIp?: string;
    readonly publishDatabase?: string;
    readonly publishSubscribeId?: number;
    readonly publishSubscribeLists: outputs.Sqlserver.GetPublishSubscribesPublishSubscribeList[];
    readonly publishSubscribeName?: string;
    readonly resultOutputFile?: string;
    readonly subscribeDatabase?: string;
}
export function getPublishSubscribesOutput(args: GetPublishSubscribesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPublishSubscribesResult> {
    return pulumi.output(args).apply((a: any) => getPublishSubscribes(a, opts))
}

/**
 * A collection of arguments for invoking getPublishSubscribes.
 */
export interface GetPublishSubscribesOutputArgs {
    instanceId: pulumi.Input<string>;
    pubOrSubInstanceId?: pulumi.Input<string>;
    pubOrSubInstanceIp?: pulumi.Input<string>;
    publishDatabase?: pulumi.Input<string>;
    publishSubscribeId?: pulumi.Input<number>;
    publishSubscribeName?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    subscribeDatabase?: pulumi.Input<string>;
}
