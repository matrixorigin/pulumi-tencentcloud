// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getCheckDataEngineImageCanBeRollback(args: GetCheckDataEngineImageCanBeRollbackArgs, opts?: pulumi.InvokeOptions): Promise<GetCheckDataEngineImageCanBeRollbackResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dlc/getCheckDataEngineImageCanBeRollback:getCheckDataEngineImageCanBeRollback", {
        "dataEngineId": args.dataEngineId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getCheckDataEngineImageCanBeRollback.
 */
export interface GetCheckDataEngineImageCanBeRollbackArgs {
    dataEngineId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getCheckDataEngineImageCanBeRollback.
 */
export interface GetCheckDataEngineImageCanBeRollbackResult {
    readonly dataEngineId: string;
    readonly fromRecordId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isRollback: boolean;
    readonly resultOutputFile?: string;
    readonly toRecordId: string;
}
export function getCheckDataEngineImageCanBeRollbackOutput(args: GetCheckDataEngineImageCanBeRollbackOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCheckDataEngineImageCanBeRollbackResult> {
    return pulumi.output(args).apply((a: any) => getCheckDataEngineImageCanBeRollback(a, opts))
}

/**
 * A collection of arguments for invoking getCheckDataEngineImageCanBeRollback.
 */
export interface GetCheckDataEngineImageCanBeRollbackOutputArgs {
    dataEngineId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}