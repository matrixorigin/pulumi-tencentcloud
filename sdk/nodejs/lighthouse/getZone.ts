// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getZone(args?: GetZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Lighthouse/getZone:getZone", {
        "order": args.order,
        "orderField": args.orderField,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneArgs {
    order?: string;
    orderField?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getZone.
 */
export interface GetZoneResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly order?: string;
    readonly orderField?: string;
    readonly resultOutputFile?: string;
    readonly zoneInfoSets: outputs.Lighthouse.GetZoneZoneInfoSet[];
}
export function getZoneOutput(args?: GetZoneOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZoneResult> {
    return pulumi.output(args).apply((a: any) => getZone(a, opts))
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneOutputArgs {
    order?: pulumi.Input<string>;
    orderField?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}