// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getCheckSavepoint(args: GetCheckSavepointArgs, opts?: pulumi.InvokeOptions): Promise<GetCheckSavepointResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Oceanus/getCheckSavepoint:getCheckSavepoint", {
        "jobId": args.jobId,
        "recordType": args.recordType,
        "resultOutputFile": args.resultOutputFile,
        "savepointPath": args.savepointPath,
        "serialId": args.serialId,
        "workSpaceId": args.workSpaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCheckSavepoint.
 */
export interface GetCheckSavepointArgs {
    jobId: string;
    recordType: number;
    resultOutputFile?: string;
    savepointPath: string;
    serialId: string;
    workSpaceId: string;
}

/**
 * A collection of values returned by getCheckSavepoint.
 */
export interface GetCheckSavepointResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly jobId: string;
    readonly recordType: number;
    readonly resultOutputFile?: string;
    readonly savepointPath: string;
    readonly savepointStatus: number;
    readonly serialId: string;
    readonly workSpaceId: string;
}
export function getCheckSavepointOutput(args: GetCheckSavepointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCheckSavepointResult> {
    return pulumi.output(args).apply((a: any) => getCheckSavepoint(a, opts))
}

/**
 * A collection of arguments for invoking getCheckSavepoint.
 */
export interface GetCheckSavepointOutputArgs {
    jobId: pulumi.Input<string>;
    recordType: pulumi.Input<number>;
    resultOutputFile?: pulumi.Input<string>;
    savepointPath: pulumi.Input<string>;
    serialId: pulumi.Input<string>;
    workSpaceId: pulumi.Input<string>;
}
