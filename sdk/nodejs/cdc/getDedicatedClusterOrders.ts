// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDedicatedClusterOrders(args?: GetDedicatedClusterOrdersArgs, opts?: pulumi.InvokeOptions): Promise<GetDedicatedClusterOrdersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cdc/getDedicatedClusterOrders:getDedicatedClusterOrders", {
        "actionType": args.actionType,
        "dedicatedClusterIds": args.dedicatedClusterIds,
        "resultOutputFile": args.resultOutputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getDedicatedClusterOrders.
 */
export interface GetDedicatedClusterOrdersArgs {
    actionType?: string;
    dedicatedClusterIds?: string[];
    resultOutputFile?: string;
    status?: string;
}

/**
 * A collection of values returned by getDedicatedClusterOrders.
 */
export interface GetDedicatedClusterOrdersResult {
    readonly actionType?: string;
    readonly dedicatedClusterIds?: string[];
    readonly dedicatedClusterOrderSets: outputs.Cdc.GetDedicatedClusterOrdersDedicatedClusterOrderSet[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly status?: string;
}
export function getDedicatedClusterOrdersOutput(args?: GetDedicatedClusterOrdersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDedicatedClusterOrdersResult> {
    return pulumi.output(args).apply((a: any) => getDedicatedClusterOrders(a, opts))
}

/**
 * A collection of arguments for invoking getDedicatedClusterOrders.
 */
export interface GetDedicatedClusterOrdersOutputArgs {
    actionType?: pulumi.Input<string>;
    dedicatedClusterIds?: pulumi.Input<pulumi.Input<string>[]>;
    resultOutputFile?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}
