// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getCluster(args?: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Tsf/getCluster:getCluster", {
        "clusterIdLists": args.clusterIdLists,
        "clusterType": args.clusterType,
        "disableProgramAuthCheck": args.disableProgramAuthCheck,
        "resultOutputFile": args.resultOutputFile,
        "searchWord": args.searchWord,
    }, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    clusterIdLists?: string[];
    clusterType?: string;
    disableProgramAuthCheck?: boolean;
    resultOutputFile?: string;
    searchWord?: string;
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    readonly clusterIdLists?: string[];
    readonly clusterType?: string;
    readonly disableProgramAuthCheck?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly results: outputs.Tsf.GetClusterResult[];
    readonly searchWord?: string;
}
export function getClusterOutput(args?: GetClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterResult> {
    return pulumi.output(args).apply((a: any) => getCluster(a, opts))
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterOutputArgs {
    clusterIdLists?: pulumi.Input<pulumi.Input<string>[]>;
    clusterType?: pulumi.Input<string>;
    disableProgramAuthCheck?: pulumi.Input<boolean>;
    resultOutputFile?: pulumi.Input<string>;
    searchWord?: pulumi.Input<string>;
}