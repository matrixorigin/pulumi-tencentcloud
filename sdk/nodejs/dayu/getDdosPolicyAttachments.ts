// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDdosPolicyAttachments(args: GetDdosPolicyAttachmentsArgs, opts?: pulumi.InvokeOptions): Promise<GetDdosPolicyAttachmentsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dayu/getDdosPolicyAttachments:getDdosPolicyAttachments", {
        "policyId": args.policyId,
        "resourceId": args.resourceId,
        "resourceType": args.resourceType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDdosPolicyAttachments.
 */
export interface GetDdosPolicyAttachmentsArgs {
    policyId?: string;
    resourceId?: string;
    resourceType: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDdosPolicyAttachments.
 */
export interface GetDdosPolicyAttachmentsResult {
    readonly dayuDdosPolicyAttachmentLists: outputs.Dayu.GetDdosPolicyAttachmentsDayuDdosPolicyAttachmentList[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly policyId?: string;
    readonly resourceId?: string;
    readonly resourceType: string;
    readonly resultOutputFile?: string;
}
export function getDdosPolicyAttachmentsOutput(args: GetDdosPolicyAttachmentsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDdosPolicyAttachmentsResult> {
    return pulumi.output(args).apply((a: any) => getDdosPolicyAttachments(a, opts))
}

/**
 * A collection of arguments for invoking getDdosPolicyAttachments.
 */
export interface GetDdosPolicyAttachmentsOutputArgs {
    policyId?: pulumi.Input<string>;
    resourceId?: pulumi.Input<string>;
    resourceType: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
