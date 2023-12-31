// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getInstance(args?: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Deprecatedeip/getInstance:getInstance", {
        "filters": args.filters,
        "includeArrears": args.includeArrears,
        "includeBlocked": args.includeBlocked,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    filters?: inputs.Deprecatedeip.GetInstanceFilter[];
    includeArrears?: boolean;
    includeBlocked?: boolean;
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    readonly filters?: outputs.Deprecatedeip.GetInstanceFilter[];
    readonly id: string;
    readonly includeArrears?: boolean;
    readonly includeBlocked?: boolean;
    readonly publicIp: string;
    readonly status: string;
}
export function getInstanceOutput(args?: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply((a: any) => getInstance(a, opts))
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.Deprecatedeip.GetInstanceFilterArgs>[]>;
    includeArrears?: pulumi.Input<boolean>;
    includeBlocked?: pulumi.Input<boolean>;
}
