// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDeliverLogDownList(args?: GetDeliverLogDownListArgs, opts?: pulumi.InvokeOptions): Promise<GetDeliverLogDownListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Css/getDeliverLogDownList:getDeliverLogDownList", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDeliverLogDownList.
 */
export interface GetDeliverLogDownListArgs {
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDeliverLogDownList.
 */
export interface GetDeliverLogDownListResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly logInfoLists: outputs.Css.GetDeliverLogDownListLogInfoList[];
    readonly resultOutputFile?: string;
}
export function getDeliverLogDownListOutput(args?: GetDeliverLogDownListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeliverLogDownListResult> {
    return pulumi.output(args).apply((a: any) => getDeliverLogDownList(a, opts))
}

/**
 * A collection of arguments for invoking getDeliverLogDownList.
 */
export interface GetDeliverLogDownListOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
}
