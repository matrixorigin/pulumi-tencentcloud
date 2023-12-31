// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getUsedIpAddress(args: GetUsedIpAddressArgs, opts?: pulumi.InvokeOptions): Promise<GetUsedIpAddressResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Vpc/getUsedIpAddress:getUsedIpAddress", {
        "ipAddresses": args.ipAddresses,
        "resultOutputFile": args.resultOutputFile,
        "subnetId": args.subnetId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsedIpAddress.
 */
export interface GetUsedIpAddressArgs {
    ipAddresses?: string[];
    resultOutputFile?: string;
    subnetId?: string;
    vpcId: string;
}

/**
 * A collection of values returned by getUsedIpAddress.
 */
export interface GetUsedIpAddressResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipAddressStates: outputs.Vpc.GetUsedIpAddressIpAddressState[];
    readonly ipAddresses?: string[];
    readonly resultOutputFile?: string;
    readonly subnetId?: string;
    readonly vpcId: string;
}
export function getUsedIpAddressOutput(args: GetUsedIpAddressOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUsedIpAddressResult> {
    return pulumi.output(args).apply((a: any) => getUsedIpAddress(a, opts))
}

/**
 * A collection of arguments for invoking getUsedIpAddress.
 */
export interface GetUsedIpAddressOutputArgs {
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    resultOutputFile?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
    vpcId: pulumi.Input<string>;
}
