// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getInstances(args: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Tdcpg/getInstances:getInstances", {
        "clusterId": args.clusterId,
        "instanceId": args.instanceId,
        "instanceName": args.instanceName,
        "instanceType": args.instanceType,
        "resultOutputFile": args.resultOutputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    clusterId: string;
    instanceId?: string;
    instanceName?: string;
    instanceType?: string;
    resultOutputFile?: string;
    status?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    readonly clusterId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId?: string;
    readonly instanceName?: string;
    readonly instanceType?: string;
    readonly lists: outputs.Tdcpg.GetInstancesList[];
    readonly resultOutputFile?: string;
    readonly status?: string;
}
export function getInstancesOutput(args: GetInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesResult> {
    return pulumi.output(args).apply((a: any) => getInstances(a, opts))
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    clusterId: pulumi.Input<string>;
    instanceId?: pulumi.Input<string>;
    instanceName?: pulumi.Input<string>;
    instanceType?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}
