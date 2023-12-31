// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getDescribeUserType(args?: GetDescribeUserTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetDescribeUserTypeResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dlc/getDescribeUserType:getDescribeUserType", {
        "resultOutputFile": args.resultOutputFile,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDescribeUserType.
 */
export interface GetDescribeUserTypeArgs {
    resultOutputFile?: string;
    userId?: string;
}

/**
 * A collection of values returned by getDescribeUserType.
 */
export interface GetDescribeUserTypeResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly userId?: string;
    readonly userType: string;
}
export function getDescribeUserTypeOutput(args?: GetDescribeUserTypeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDescribeUserTypeResult> {
    return pulumi.output(args).apply((a: any) => getDescribeUserType(a, opts))
}

/**
 * A collection of arguments for invoking getDescribeUserType.
 */
export interface GetDescribeUserTypeOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
    userId?: pulumi.Input<string>;
}
