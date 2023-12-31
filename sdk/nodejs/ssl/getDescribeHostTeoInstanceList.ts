// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDescribeHostTeoInstanceList(args: GetDescribeHostTeoInstanceListArgs, opts?: pulumi.InvokeOptions): Promise<GetDescribeHostTeoInstanceListResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Ssl/getDescribeHostTeoInstanceList:getDescribeHostTeoInstanceList", {
        "certificateId": args.certificateId,
        "filters": args.filters,
        "isCache": args.isCache,
        "oldCertificateId": args.oldCertificateId,
        "resourceType": args.resourceType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDescribeHostTeoInstanceList.
 */
export interface GetDescribeHostTeoInstanceListArgs {
    certificateId: string;
    filters?: inputs.Ssl.GetDescribeHostTeoInstanceListFilter[];
    isCache?: number;
    oldCertificateId?: string;
    resourceType: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDescribeHostTeoInstanceList.
 */
export interface GetDescribeHostTeoInstanceListResult {
    readonly certificateId: string;
    readonly filters?: outputs.Ssl.GetDescribeHostTeoInstanceListFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceLists: outputs.Ssl.GetDescribeHostTeoInstanceListInstanceList[];
    readonly isCache?: number;
    readonly oldCertificateId?: string;
    readonly resourceType: string;
    readonly resultOutputFile?: string;
}
export function getDescribeHostTeoInstanceListOutput(args: GetDescribeHostTeoInstanceListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDescribeHostTeoInstanceListResult> {
    return pulumi.output(args).apply((a: any) => getDescribeHostTeoInstanceList(a, opts))
}

/**
 * A collection of arguments for invoking getDescribeHostTeoInstanceList.
 */
export interface GetDescribeHostTeoInstanceListOutputArgs {
    certificateId: pulumi.Input<string>;
    filters?: pulumi.Input<pulumi.Input<inputs.Ssl.GetDescribeHostTeoInstanceListFilterArgs>[]>;
    isCache?: pulumi.Input<number>;
    oldCertificateId?: pulumi.Input<string>;
    resourceType: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
