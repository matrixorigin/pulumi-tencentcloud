// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getValidateStatus(args: GetValidateStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetValidateStatusResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Subdomain/getValidateStatus:getValidateStatus", {
        "domainZone": args.domainZone,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getValidateStatus.
 */
export interface GetValidateStatusArgs {
    domainZone: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getValidateStatus.
 */
export interface GetValidateStatusResult {
    readonly domainZone: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly status: number;
}
export function getValidateStatusOutput(args: GetValidateStatusOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetValidateStatusResult> {
    return pulumi.output(args).apply((a: any) => getValidateStatus(a, opts))
}

/**
 * A collection of arguments for invoking getValidateStatus.
 */
export interface GetValidateStatusOutputArgs {
    domainZone: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}