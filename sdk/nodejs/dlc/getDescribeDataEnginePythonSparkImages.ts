// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDescribeDataEnginePythonSparkImages(args: GetDescribeDataEnginePythonSparkImagesArgs, opts?: pulumi.InvokeOptions): Promise<GetDescribeDataEnginePythonSparkImagesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dlc/getDescribeDataEnginePythonSparkImages:getDescribeDataEnginePythonSparkImages", {
        "childImageVersionId": args.childImageVersionId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDescribeDataEnginePythonSparkImages.
 */
export interface GetDescribeDataEnginePythonSparkImagesArgs {
    childImageVersionId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDescribeDataEnginePythonSparkImages.
 */
export interface GetDescribeDataEnginePythonSparkImagesResult {
    readonly childImageVersionId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly pythonSparkImages: outputs.Dlc.GetDescribeDataEnginePythonSparkImagesPythonSparkImage[];
    readonly resultOutputFile?: string;
}
export function getDescribeDataEnginePythonSparkImagesOutput(args: GetDescribeDataEnginePythonSparkImagesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDescribeDataEnginePythonSparkImagesResult> {
    return pulumi.output(args).apply((a: any) => getDescribeDataEnginePythonSparkImages(a, opts))
}

/**
 * A collection of arguments for invoking getDescribeDataEnginePythonSparkImages.
 */
export interface GetDescribeDataEnginePythonSparkImagesOutputArgs {
    childImageVersionId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
