// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getScenarioWithJobs(args: GetScenarioWithJobsArgs, opts?: pulumi.InvokeOptions): Promise<GetScenarioWithJobsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Pts/getScenarioWithJobs:getScenarioWithJobs", {
        "ascend": args.ascend,
        "ignoreDataset": args.ignoreDataset,
        "ignoreScript": args.ignoreScript,
        "orderBy": args.orderBy,
        "owner": args.owner,
        "projectIds": args.projectIds,
        "resultOutputFile": args.resultOutputFile,
        "scenarioIds": args.scenarioIds,
        "scenarioName": args.scenarioName,
        "scenarioType": args.scenarioType,
    }, opts);
}

/**
 * A collection of arguments for invoking getScenarioWithJobs.
 */
export interface GetScenarioWithJobsArgs {
    ascend?: boolean;
    ignoreDataset?: boolean;
    ignoreScript?: boolean;
    orderBy?: string;
    owner?: string;
    projectIds: string[];
    resultOutputFile?: string;
    scenarioIds?: string[];
    scenarioName?: string;
    scenarioType?: string;
}

/**
 * A collection of values returned by getScenarioWithJobs.
 */
export interface GetScenarioWithJobsResult {
    readonly ascend?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ignoreDataset?: boolean;
    readonly ignoreScript?: boolean;
    readonly orderBy?: string;
    readonly owner?: string;
    readonly projectIds: string[];
    readonly resultOutputFile?: string;
    readonly scenarioIds?: string[];
    readonly scenarioName?: string;
    readonly scenarioType?: string;
    readonly scenarioWithJobsSets: outputs.Pts.GetScenarioWithJobsScenarioWithJobsSet[];
}
export function getScenarioWithJobsOutput(args: GetScenarioWithJobsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetScenarioWithJobsResult> {
    return pulumi.output(args).apply((a: any) => getScenarioWithJobs(a, opts))
}

/**
 * A collection of arguments for invoking getScenarioWithJobs.
 */
export interface GetScenarioWithJobsOutputArgs {
    ascend?: pulumi.Input<boolean>;
    ignoreDataset?: pulumi.Input<boolean>;
    ignoreScript?: pulumi.Input<boolean>;
    orderBy?: pulumi.Input<string>;
    owner?: pulumi.Input<string>;
    projectIds: pulumi.Input<pulumi.Input<string>[]>;
    resultOutputFile?: pulumi.Input<string>;
    scenarioIds?: pulumi.Input<pulumi.Input<string>[]>;
    scenarioName?: pulumi.Input<string>;
    scenarioType?: pulumi.Input<string>;
}
