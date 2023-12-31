// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getEdgeFwSwitches(args?: GetEdgeFwSwitchesArgs, opts?: pulumi.InvokeOptions): Promise<GetEdgeFwSwitchesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cfw/getEdgeFwSwitches:getEdgeFwSwitches", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getEdgeFwSwitches.
 */
export interface GetEdgeFwSwitchesArgs {
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getEdgeFwSwitches.
 */
export interface GetEdgeFwSwitchesResult {
    readonly datas: outputs.Cfw.GetEdgeFwSwitchesData[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
}
export function getEdgeFwSwitchesOutput(args?: GetEdgeFwSwitchesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEdgeFwSwitchesResult> {
    return pulumi.output(args).apply((a: any) => getEdgeFwSwitches(a, opts))
}

/**
 * A collection of arguments for invoking getEdgeFwSwitches.
 */
export interface GetEdgeFwSwitchesOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
}
