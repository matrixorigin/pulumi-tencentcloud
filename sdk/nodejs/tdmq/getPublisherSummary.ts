// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getPublisherSummary(args: GetPublisherSummaryArgs, opts?: pulumi.InvokeOptions): Promise<GetPublisherSummaryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Tdmq/getPublisherSummary:getPublisherSummary", {
        "clusterId": args.clusterId,
        "namespace": args.namespace,
        "resultOutputFile": args.resultOutputFile,
        "topic": args.topic,
    }, opts);
}

/**
 * A collection of arguments for invoking getPublisherSummary.
 */
export interface GetPublisherSummaryArgs {
    clusterId: string;
    namespace: string;
    resultOutputFile?: string;
    topic: string;
}

/**
 * A collection of values returned by getPublisherSummary.
 */
export interface GetPublisherSummaryResult {
    readonly clusterId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly msgRateIn: number;
    readonly msgThroughputIn: number;
    readonly namespace: string;
    readonly publisherCount: number;
    readonly resultOutputFile?: string;
    readonly storageSize: number;
    readonly topic: string;
}
export function getPublisherSummaryOutput(args: GetPublisherSummaryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPublisherSummaryResult> {
    return pulumi.output(args).apply((a: any) => getPublisherSummary(a, opts))
}

/**
 * A collection of arguments for invoking getPublisherSummary.
 */
export interface GetPublisherSummaryOutputArgs {
    clusterId: pulumi.Input<string>;
    namespace: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    topic: pulumi.Input<string>;
}
