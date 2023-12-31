// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDescribeHostLiveInstanceList(args: GetDescribeHostLiveInstanceListArgs, opts?: pulumi.InvokeOptions): Promise<GetDescribeHostLiveInstanceListResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Ssl/getDescribeHostLiveInstanceList:getDescribeHostLiveInstanceList", {
        "certificateId": args.certificateId,
        "filters": args.filters,
        "isCache": args.isCache,
        "oldCertificateId": args.oldCertificateId,
        "resourceType": args.resourceType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDescribeHostLiveInstanceList.
 */
export interface GetDescribeHostLiveInstanceListArgs {
    certificateId: string;
    filters?: inputs.Ssl.GetDescribeHostLiveInstanceListFilter[];
    isCache?: number;
    oldCertificateId?: string;
    resourceType: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDescribeHostLiveInstanceList.
 */
export interface GetDescribeHostLiveInstanceListResult {
    readonly certificateId: string;
    readonly filters?: outputs.Ssl.GetDescribeHostLiveInstanceListFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceLists: outputs.Ssl.GetDescribeHostLiveInstanceListInstanceList[];
    readonly isCache?: number;
    readonly oldCertificateId?: string;
    readonly resourceType: string;
    readonly resultOutputFile?: string;
}
export function getDescribeHostLiveInstanceListOutput(args: GetDescribeHostLiveInstanceListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDescribeHostLiveInstanceListResult> {
    return pulumi.output(args).apply((a: any) => getDescribeHostLiveInstanceList(a, opts))
}

/**
 * A collection of arguments for invoking getDescribeHostLiveInstanceList.
 */
export interface GetDescribeHostLiveInstanceListOutputArgs {
    certificateId: pulumi.Input<string>;
    filters?: pulumi.Input<pulumi.Input<inputs.Ssl.GetDescribeHostLiveInstanceListFilterArgs>[]>;
    isCache?: pulumi.Input<number>;
    oldCertificateId?: pulumi.Input<string>;
    resourceType: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
