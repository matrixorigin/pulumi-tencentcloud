// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getModifyInstanceBundle(args: GetModifyInstanceBundleArgs, opts?: pulumi.InvokeOptions): Promise<GetModifyInstanceBundleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Lighthouse/getModifyInstanceBundle:getModifyInstanceBundle", {
        "filters": args.filters,
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getModifyInstanceBundle.
 */
export interface GetModifyInstanceBundleArgs {
    filters?: inputs.Lighthouse.GetModifyInstanceBundleFilter[];
    instanceId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getModifyInstanceBundle.
 */
export interface GetModifyInstanceBundleResult {
    readonly filters?: outputs.Lighthouse.GetModifyInstanceBundleFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly modifyBundleSets: outputs.Lighthouse.GetModifyInstanceBundleModifyBundleSet[];
    readonly resultOutputFile?: string;
}
export function getModifyInstanceBundleOutput(args: GetModifyInstanceBundleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetModifyInstanceBundleResult> {
    return pulumi.output(args).apply((a: any) => getModifyInstanceBundle(a, opts))
}

/**
 * A collection of arguments for invoking getModifyInstanceBundle.
 */
export interface GetModifyInstanceBundleOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.Lighthouse.GetModifyInstanceBundleFilterArgs>[]>;
    instanceId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
