// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getCrossRegionZone(args: GetCrossRegionZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetCrossRegionZoneResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Sqlserver/getCrossRegionZone:getCrossRegionZone", {
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getCrossRegionZone.
 */
export interface GetCrossRegionZoneArgs {
    instanceId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getCrossRegionZone.
 */
export interface GetCrossRegionZoneResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly region: string;
    readonly resultOutputFile?: string;
    readonly zone: string;
}
export function getCrossRegionZoneOutput(args: GetCrossRegionZoneOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCrossRegionZoneResult> {
    return pulumi.output(args).apply((a: any) => getCrossRegionZone(a, opts))
}

/**
 * A collection of arguments for invoking getCrossRegionZone.
 */
export interface GetCrossRegionZoneOutputArgs {
    instanceId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
