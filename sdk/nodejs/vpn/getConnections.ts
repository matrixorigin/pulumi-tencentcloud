// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getConnections(args?: GetConnectionsArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectionsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Vpn/getConnections:getConnections", {
        "customerGatewayId": args.customerGatewayId,
        "id": args.id,
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
        "tags": args.tags,
        "vpcId": args.vpcId,
        "vpnGatewayId": args.vpnGatewayId,
    }, opts);
}

/**
 * A collection of arguments for invoking getConnections.
 */
export interface GetConnectionsArgs {
    customerGatewayId?: string;
    id?: string;
    name?: string;
    resultOutputFile?: string;
    tags?: {[key: string]: any};
    vpcId?: string;
    vpnGatewayId?: string;
}

/**
 * A collection of values returned by getConnections.
 */
export interface GetConnectionsResult {
    readonly connectionLists: outputs.Vpn.GetConnectionsConnectionList[];
    readonly customerGatewayId?: string;
    readonly id?: string;
    readonly name?: string;
    readonly resultOutputFile?: string;
    readonly tags?: {[key: string]: any};
    readonly vpcId?: string;
    readonly vpnGatewayId?: string;
}
export function getConnectionsOutput(args?: GetConnectionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConnectionsResult> {
    return pulumi.output(args).apply((a: any) => getConnections(a, opts))
}

/**
 * A collection of arguments for invoking getConnections.
 */
export interface GetConnectionsOutputArgs {
    customerGatewayId?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: any}>;
    vpcId?: pulumi.Input<string>;
    vpnGatewayId?: pulumi.Input<string>;
}