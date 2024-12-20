// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getInstanceNodes(args: GetInstanceNodesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceNodesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Clickhouse/getInstanceNodes:getInstanceNodes", {
        "displayPolicy": args.displayPolicy,
        "forceAll": args.forceAll,
        "instanceId": args.instanceId,
        "nodeRole": args.nodeRole,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceNodes.
 */
export interface GetInstanceNodesArgs {
    displayPolicy?: string;
    forceAll?: boolean;
    instanceId: string;
    nodeRole?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstanceNodes.
 */
export interface GetInstanceNodesResult {
    readonly displayPolicy?: string;
    readonly forceAll?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly instanceNodesLists: outputs.Clickhouse.GetInstanceNodesInstanceNodesList[];
    readonly nodeRole?: string;
    readonly resultOutputFile?: string;
}
export function getInstanceNodesOutput(args: GetInstanceNodesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceNodesResult> {
    return pulumi.output(args).apply((a: any) => getInstanceNodes(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceNodes.
 */
export interface GetInstanceNodesOutputArgs {
    displayPolicy?: pulumi.Input<string>;
    forceAll?: pulumi.Input<boolean>;
    instanceId: pulumi.Input<string>;
    nodeRole?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
