// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getClusterNodePools(args: GetClusterNodePoolsArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterNodePoolsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Kubernetes/getClusterNodePools:getClusterNodePools", {
        "clusterId": args.clusterId,
        "filters": args.filters,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusterNodePools.
 */
export interface GetClusterNodePoolsArgs {
    clusterId: string;
    filters?: inputs.Kubernetes.GetClusterNodePoolsFilter[];
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getClusterNodePools.
 */
export interface GetClusterNodePoolsResult {
    readonly clusterId: string;
    readonly filters?: outputs.Kubernetes.GetClusterNodePoolsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nodePoolSets: outputs.Kubernetes.GetClusterNodePoolsNodePoolSet[];
    readonly resultOutputFile?: string;
}
export function getClusterNodePoolsOutput(args: GetClusterNodePoolsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterNodePoolsResult> {
    return pulumi.output(args).apply((a: any) => getClusterNodePools(a, opts))
}

/**
 * A collection of arguments for invoking getClusterNodePools.
 */
export interface GetClusterNodePoolsOutputArgs {
    clusterId: pulumi.Input<string>;
    filters?: pulumi.Input<pulumi.Input<inputs.Kubernetes.GetClusterNodePoolsFilterArgs>[]>;
    resultOutputFile?: pulumi.Input<string>;
}
