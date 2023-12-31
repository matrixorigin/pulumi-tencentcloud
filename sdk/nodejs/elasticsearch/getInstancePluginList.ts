// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getInstancePluginList(args: GetInstancePluginListArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancePluginListResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Elasticsearch/getInstancePluginList:getInstancePluginList", {
        "instanceId": args.instanceId,
        "orderBy": args.orderBy,
        "orderByType": args.orderByType,
        "pluginType": args.pluginType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstancePluginList.
 */
export interface GetInstancePluginListArgs {
    instanceId: string;
    orderBy?: string;
    orderByType?: string;
    pluginType?: number;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstancePluginList.
 */
export interface GetInstancePluginListResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly orderBy?: string;
    readonly orderByType?: string;
    readonly pluginLists: outputs.Elasticsearch.GetInstancePluginListPluginList[];
    readonly pluginType?: number;
    readonly resultOutputFile?: string;
}
export function getInstancePluginListOutput(args: GetInstancePluginListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancePluginListResult> {
    return pulumi.output(args).apply((a: any) => getInstancePluginList(a, opts))
}

/**
 * A collection of arguments for invoking getInstancePluginList.
 */
export interface GetInstancePluginListOutputArgs {
    instanceId: pulumi.Input<string>;
    orderBy?: pulumi.Input<string>;
    orderByType?: pulumi.Input<string>;
    pluginType?: pulumi.Input<number>;
    resultOutputFile?: pulumi.Input<string>;
}
