// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cynosdb/getInstances:getInstances", {
        "clusterId": args.clusterId,
        "dbType": args.dbType,
        "instanceId": args.instanceId,
        "instanceName": args.instanceName,
        "projectId": args.projectId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    clusterId?: string;
    dbType?: string;
    instanceId?: string;
    instanceName?: string;
    projectId?: number;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    readonly clusterId?: string;
    readonly dbType?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId?: string;
    readonly instanceLists: outputs.Cynosdb.GetInstancesInstanceList[];
    readonly instanceName?: string;
    readonly projectId?: number;
    readonly resultOutputFile?: string;
}
export function getInstancesOutput(args?: GetInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesResult> {
    return pulumi.output(args).apply((a: any) => getInstances(a, opts))
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    clusterId?: pulumi.Input<string>;
    dbType?: pulumi.Input<string>;
    instanceId?: pulumi.Input<string>;
    instanceName?: pulumi.Input<string>;
    projectId?: pulumi.Input<number>;
    resultOutputFile?: pulumi.Input<string>;
}
