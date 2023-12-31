// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getListEntitiesForPolicy(args: GetListEntitiesForPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetListEntitiesForPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cam/getListEntitiesForPolicy:getListEntitiesForPolicy", {
        "entityFilter": args.entityFilter,
        "policyId": args.policyId,
        "resultOutputFile": args.resultOutputFile,
        "rp": args.rp,
    }, opts);
}

/**
 * A collection of arguments for invoking getListEntitiesForPolicy.
 */
export interface GetListEntitiesForPolicyArgs {
    entityFilter?: string;
    policyId: number;
    resultOutputFile?: string;
    rp?: number;
}

/**
 * A collection of values returned by getListEntitiesForPolicy.
 */
export interface GetListEntitiesForPolicyResult {
    readonly entityFilter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.Cam.GetListEntitiesForPolicyList[];
    readonly policyId: number;
    readonly resultOutputFile?: string;
    readonly rp?: number;
}
export function getListEntitiesForPolicyOutput(args: GetListEntitiesForPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetListEntitiesForPolicyResult> {
    return pulumi.output(args).apply((a: any) => getListEntitiesForPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getListEntitiesForPolicy.
 */
export interface GetListEntitiesForPolicyOutputArgs {
    entityFilter?: pulumi.Input<string>;
    policyId: pulumi.Input<number>;
    resultOutputFile?: pulumi.Input<string>;
    rp?: pulumi.Input<number>;
}
