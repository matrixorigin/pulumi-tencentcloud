// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getProductEvent(args?: GetProductEventArgs, opts?: pulumi.InvokeOptions): Promise<GetProductEventResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Monitor/getProductEvent:getProductEvent", {
        "dimensions": args.dimensions,
        "endTime": args.endTime,
        "eventNames": args.eventNames,
        "instanceIds": args.instanceIds,
        "isAlarmConfig": args.isAlarmConfig,
        "productNames": args.productNames,
        "projectIds": args.projectIds,
        "regionLists": args.regionLists,
        "resultOutputFile": args.resultOutputFile,
        "startTime": args.startTime,
        "statuses": args.statuses,
        "types": args.types,
    }, opts);
}

/**
 * A collection of arguments for invoking getProductEvent.
 */
export interface GetProductEventArgs {
    dimensions?: inputs.Monitor.GetProductEventDimension[];
    endTime?: number;
    eventNames?: string[];
    instanceIds?: string[];
    isAlarmConfig?: number;
    productNames?: string[];
    projectIds?: string[];
    regionLists?: string[];
    resultOutputFile?: string;
    startTime?: number;
    statuses?: string[];
    types?: string[];
}

/**
 * A collection of values returned by getProductEvent.
 */
export interface GetProductEventResult {
    readonly dimensions?: outputs.Monitor.GetProductEventDimension[];
    readonly endTime?: number;
    readonly eventNames?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceIds?: string[];
    readonly isAlarmConfig?: number;
    readonly lists: outputs.Monitor.GetProductEventList[];
    readonly productNames?: string[];
    readonly projectIds?: string[];
    readonly regionLists?: string[];
    readonly resultOutputFile?: string;
    readonly startTime?: number;
    readonly statuses?: string[];
    readonly types?: string[];
}
export function getProductEventOutput(args?: GetProductEventOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProductEventResult> {
    return pulumi.output(args).apply((a: any) => getProductEvent(a, opts))
}

/**
 * A collection of arguments for invoking getProductEvent.
 */
export interface GetProductEventOutputArgs {
    dimensions?: pulumi.Input<pulumi.Input<inputs.Monitor.GetProductEventDimensionArgs>[]>;
    endTime?: pulumi.Input<number>;
    eventNames?: pulumi.Input<pulumi.Input<string>[]>;
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    isAlarmConfig?: pulumi.Input<number>;
    productNames?: pulumi.Input<pulumi.Input<string>[]>;
    projectIds?: pulumi.Input<pulumi.Input<string>[]>;
    regionLists?: pulumi.Input<pulumi.Input<string>[]>;
    resultOutputFile?: pulumi.Input<string>;
    startTime?: pulumi.Input<number>;
    statuses?: pulumi.Input<pulumi.Input<string>[]>;
    types?: pulumi.Input<pulumi.Input<string>[]>;
}