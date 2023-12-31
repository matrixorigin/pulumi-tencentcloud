// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDescribeUserInfo(args?: GetDescribeUserInfoArgs, opts?: pulumi.InvokeOptions): Promise<GetDescribeUserInfoResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dlc/getDescribeUserInfo:getDescribeUserInfo", {
        "filters": args.filters,
        "resultOutputFile": args.resultOutputFile,
        "sortBy": args.sortBy,
        "sorting": args.sorting,
        "type": args.type,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDescribeUserInfo.
 */
export interface GetDescribeUserInfoArgs {
    filters?: inputs.Dlc.GetDescribeUserInfoFilter[];
    resultOutputFile?: string;
    sortBy?: string;
    sorting?: string;
    type?: string;
    userId?: string;
}

/**
 * A collection of values returned by getDescribeUserInfo.
 */
export interface GetDescribeUserInfoResult {
    readonly filters?: outputs.Dlc.GetDescribeUserInfoFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly sortBy?: string;
    readonly sorting?: string;
    readonly type?: string;
    readonly userId?: string;
    readonly userInfos: outputs.Dlc.GetDescribeUserInfoUserInfo[];
}
export function getDescribeUserInfoOutput(args?: GetDescribeUserInfoOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDescribeUserInfoResult> {
    return pulumi.output(args).apply((a: any) => getDescribeUserInfo(a, opts))
}

/**
 * A collection of arguments for invoking getDescribeUserInfo.
 */
export interface GetDescribeUserInfoOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.Dlc.GetDescribeUserInfoFilterArgs>[]>;
    resultOutputFile?: pulumi.Input<string>;
    sortBy?: pulumi.Input<string>;
    sorting?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userId?: pulumi.Input<string>;
}
