// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getServiceEnvironmentList(args: GetServiceEnvironmentListArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceEnvironmentListResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:ApiGateway/getServiceEnvironmentList:getServiceEnvironmentList", {
        "resultOutputFile": args.resultOutputFile,
        "serviceId": args.serviceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceEnvironmentList.
 */
export interface GetServiceEnvironmentListArgs {
    resultOutputFile?: string;
    serviceId: string;
}

/**
 * A collection of values returned by getServiceEnvironmentList.
 */
export interface GetServiceEnvironmentListResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly results: outputs.ApiGateway.GetServiceEnvironmentListResult[];
    readonly serviceId: string;
}
export function getServiceEnvironmentListOutput(args: GetServiceEnvironmentListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceEnvironmentListResult> {
    return pulumi.output(args).apply((a: any) => getServiceEnvironmentList(a, opts))
}

/**
 * A collection of arguments for invoking getServiceEnvironmentList.
 */
export interface GetServiceEnvironmentListOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
    serviceId: pulumi.Input<string>;
}
