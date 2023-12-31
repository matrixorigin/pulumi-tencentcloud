// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getImageSpriteTemplates(args?: GetImageSpriteTemplatesArgs, opts?: pulumi.InvokeOptions): Promise<GetImageSpriteTemplatesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Vod/getImageSpriteTemplates:getImageSpriteTemplates", {
        "definition": args.definition,
        "resultOutputFile": args.resultOutputFile,
        "subAppId": args.subAppId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getImageSpriteTemplates.
 */
export interface GetImageSpriteTemplatesArgs {
    definition?: string;
    resultOutputFile?: string;
    subAppId?: number;
    type?: string;
}

/**
 * A collection of values returned by getImageSpriteTemplates.
 */
export interface GetImageSpriteTemplatesResult {
    readonly definition?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly subAppId?: number;
    readonly templateLists: outputs.Vod.GetImageSpriteTemplatesTemplateList[];
    readonly type?: string;
}
export function getImageSpriteTemplatesOutput(args?: GetImageSpriteTemplatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImageSpriteTemplatesResult> {
    return pulumi.output(args).apply((a: any) => getImageSpriteTemplates(a, opts))
}

/**
 * A collection of arguments for invoking getImageSpriteTemplates.
 */
export interface GetImageSpriteTemplatesOutputArgs {
    definition?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    subAppId?: pulumi.Input<number>;
    type?: pulumi.Input<string>;
}
