// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getLogstashInstanceLogs(args: GetLogstashInstanceLogsArgs, opts?: pulumi.InvokeOptions): Promise<GetLogstashInstanceLogsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Elasticsearch/getLogstashInstanceLogs:getLogstashInstanceLogs", {
        "endTime": args.endTime,
        "instanceId": args.instanceId,
        "logType": args.logType,
        "orderByType": args.orderByType,
        "resultOutputFile": args.resultOutputFile,
        "searchKey": args.searchKey,
        "startTime": args.startTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getLogstashInstanceLogs.
 */
export interface GetLogstashInstanceLogsArgs {
    endTime?: string;
    instanceId: string;
    logType: number;
    orderByType?: number;
    resultOutputFile?: string;
    searchKey?: string;
    startTime?: string;
}

/**
 * A collection of values returned by getLogstashInstanceLogs.
 */
export interface GetLogstashInstanceLogsResult {
    readonly endTime?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly instanceLogLists: outputs.Elasticsearch.GetLogstashInstanceLogsInstanceLogList[];
    readonly logType: number;
    readonly orderByType?: number;
    readonly resultOutputFile?: string;
    readonly searchKey?: string;
    readonly startTime?: string;
}
export function getLogstashInstanceLogsOutput(args: GetLogstashInstanceLogsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLogstashInstanceLogsResult> {
    return pulumi.output(args).apply((a: any) => getLogstashInstanceLogs(a, opts))
}

/**
 * A collection of arguments for invoking getLogstashInstanceLogs.
 */
export interface GetLogstashInstanceLogsOutputArgs {
    endTime?: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    logType: pulumi.Input<number>;
    orderByType?: pulumi.Input<number>;
    resultOutputFile?: pulumi.Input<string>;
    searchKey?: pulumi.Input<string>;
    startTime?: pulumi.Input<string>;
}
