// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getChcHosts(args?: GetChcHostsArgs, opts?: pulumi.InvokeOptions): Promise<GetChcHostsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cvm/getChcHosts:getChcHosts", {
        "chcIds": args.chcIds,
        "filters": args.filters,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getChcHosts.
 */
export interface GetChcHostsArgs {
    chcIds?: string[];
    filters?: inputs.Cvm.GetChcHostsFilter[];
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getChcHosts.
 */
export interface GetChcHostsResult {
    readonly chcHostSets: outputs.Cvm.GetChcHostsChcHostSet[];
    readonly chcIds?: string[];
    readonly filters?: outputs.Cvm.GetChcHostsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
}
export function getChcHostsOutput(args?: GetChcHostsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetChcHostsResult> {
    return pulumi.output(args).apply((a: any) => getChcHosts(a, opts))
}

/**
 * A collection of arguments for invoking getChcHosts.
 */
export interface GetChcHostsOutputArgs {
    chcIds?: pulumi.Input<pulumi.Input<string>[]>;
    filters?: pulumi.Input<pulumi.Input<inputs.Cvm.GetChcHostsFilterArgs>[]>;
    resultOutputFile?: pulumi.Input<string>;
}
