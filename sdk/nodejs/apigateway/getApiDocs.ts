// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getApiDocs(args?: GetApiDocsArgs, opts?: pulumi.InvokeOptions): Promise<GetApiDocsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:ApiGateway/getApiDocs:getApiDocs", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getApiDocs.
 */
export interface GetApiDocsArgs {
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getApiDocs.
 */
export interface GetApiDocsResult {
    readonly apiDocLists: outputs.ApiGateway.GetApiDocsApiDocList[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
}
export function getApiDocsOutput(args?: GetApiDocsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApiDocsResult> {
    return pulumi.output(args).apply((a: any) => getApiDocs(a, opts))
}

/**
 * A collection of arguments for invoking getApiDocs.
 */
export interface GetApiDocsOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
}
