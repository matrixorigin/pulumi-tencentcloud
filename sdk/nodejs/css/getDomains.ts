// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDomains(args?: GetDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Css/getDomains:getDomains", {
        "domainPrefix": args.domainPrefix,
        "domainStatus": args.domainStatus,
        "domainType": args.domainType,
        "isDelayLive": args.isDelayLive,
        "playType": args.playType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsArgs {
    domainPrefix?: string;
    domainStatus?: number;
    domainType?: number;
    isDelayLive?: number;
    playType?: number;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDomains.
 */
export interface GetDomainsResult {
    readonly domainLists: outputs.Css.GetDomainsDomainList[];
    readonly domainPrefix?: string;
    readonly domainStatus?: number;
    readonly domainType?: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isDelayLive?: number;
    readonly playType?: number;
    readonly resultOutputFile?: string;
}
export function getDomainsOutput(args?: GetDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainsResult> {
    return pulumi.output(args).apply((a: any) => getDomains(a, opts))
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsOutputArgs {
    domainPrefix?: pulumi.Input<string>;
    domainStatus?: pulumi.Input<number>;
    domainType?: pulumi.Input<number>;
    isDelayLive?: pulumi.Input<number>;
    playType?: pulumi.Input<number>;
    resultOutputFile?: pulumi.Input<string>;
}