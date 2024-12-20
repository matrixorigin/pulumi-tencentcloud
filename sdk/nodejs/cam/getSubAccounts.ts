// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getSubAccounts(args: GetSubAccountsArgs, opts?: pulumi.InvokeOptions): Promise<GetSubAccountsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cam/getSubAccounts:getSubAccounts", {
        "filterSubAccountUins": args.filterSubAccountUins,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubAccounts.
 */
export interface GetSubAccountsArgs {
    filterSubAccountUins: number[];
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getSubAccounts.
 */
export interface GetSubAccountsResult {
    readonly filterSubAccountUins: number[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly subAccounts: outputs.Cam.GetSubAccountsSubAccount[];
}
export function getSubAccountsOutput(args: GetSubAccountsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubAccountsResult> {
    return pulumi.output(args).apply((a: any) => getSubAccounts(a, opts))
}

/**
 * A collection of arguments for invoking getSubAccounts.
 */
export interface GetSubAccountsOutputArgs {
    filterSubAccountUins: pulumi.Input<pulumi.Input<number>[]>;
    resultOutputFile?: pulumi.Input<string>;
}
