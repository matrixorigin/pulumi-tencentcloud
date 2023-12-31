// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getBlackEmailAddress(args: GetBlackEmailAddressArgs, opts?: pulumi.InvokeOptions): Promise<GetBlackEmailAddressResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Ses/getBlackEmailAddress:getBlackEmailAddress", {
        "emailAddress": args.emailAddress,
        "endDate": args.endDate,
        "resultOutputFile": args.resultOutputFile,
        "startDate": args.startDate,
        "taskId": args.taskId,
    }, opts);
}

/**
 * A collection of arguments for invoking getBlackEmailAddress.
 */
export interface GetBlackEmailAddressArgs {
    emailAddress?: string;
    endDate: string;
    resultOutputFile?: string;
    startDate: string;
    taskId?: string;
}

/**
 * A collection of values returned by getBlackEmailAddress.
 */
export interface GetBlackEmailAddressResult {
    readonly blackLists: outputs.Ses.GetBlackEmailAddressBlackList[];
    readonly emailAddress?: string;
    readonly endDate: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly startDate: string;
    readonly taskId?: string;
}
export function getBlackEmailAddressOutput(args: GetBlackEmailAddressOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBlackEmailAddressResult> {
    return pulumi.output(args).apply((a: any) => getBlackEmailAddress(a, opts))
}

/**
 * A collection of arguments for invoking getBlackEmailAddress.
 */
export interface GetBlackEmailAddressOutputArgs {
    emailAddress?: pulumi.Input<string>;
    endDate: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    startDate: pulumi.Input<string>;
    taskId?: pulumi.Input<string>;
}
