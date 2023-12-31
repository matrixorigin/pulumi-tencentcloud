// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getInstanceDetail(args?: GetInstanceDetailArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceDetailResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Clb/getInstanceDetail:getInstanceDetail", {
        "fields": args.fields,
        "filters": args.filters,
        "resultOutputFile": args.resultOutputFile,
        "targetType": args.targetType,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceDetail.
 */
export interface GetInstanceDetailArgs {
    fields?: string[];
    filters?: inputs.Clb.GetInstanceDetailFilter[];
    resultOutputFile?: string;
    targetType?: string;
}

/**
 * A collection of values returned by getInstanceDetail.
 */
export interface GetInstanceDetailResult {
    readonly fields?: string[];
    readonly filters?: outputs.Clb.GetInstanceDetailFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly loadBalancerDetailSets: outputs.Clb.GetInstanceDetailLoadBalancerDetailSet[];
    readonly resultOutputFile?: string;
    readonly targetType?: string;
}
export function getInstanceDetailOutput(args?: GetInstanceDetailOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceDetailResult> {
    return pulumi.output(args).apply((a: any) => getInstanceDetail(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceDetail.
 */
export interface GetInstanceDetailOutputArgs {
    fields?: pulumi.Input<pulumi.Input<string>[]>;
    filters?: pulumi.Input<pulumi.Input<inputs.Clb.GetInstanceDetailFilterArgs>[]>;
    resultOutputFile?: pulumi.Input<string>;
    targetType?: pulumi.Input<string>;
}
