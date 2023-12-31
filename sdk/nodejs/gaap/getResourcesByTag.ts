// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getResourcesByTag(args: GetResourcesByTagArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcesByTagResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Gaap/getResourcesByTag:getResourcesByTag", {
        "resourceType": args.resourceType,
        "resultOutputFile": args.resultOutputFile,
        "tagKey": args.tagKey,
        "tagValue": args.tagValue,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourcesByTag.
 */
export interface GetResourcesByTagArgs {
    resourceType?: string;
    resultOutputFile?: string;
    tagKey: string;
    tagValue: string;
}

/**
 * A collection of values returned by getResourcesByTag.
 */
export interface GetResourcesByTagResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resourceSets: outputs.Gaap.GetResourcesByTagResourceSet[];
    readonly resourceType?: string;
    readonly resultOutputFile?: string;
    readonly tagKey: string;
    readonly tagValue: string;
}
export function getResourcesByTagOutput(args: GetResourcesByTagOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourcesByTagResult> {
    return pulumi.output(args).apply((a: any) => getResourcesByTag(a, opts))
}

/**
 * A collection of arguments for invoking getResourcesByTag.
 */
export interface GetResourcesByTagOutputArgs {
    resourceType?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    tagKey: pulumi.Input<string>;
    tagValue: pulumi.Input<string>;
}
