// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobState, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Oceanus/job:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * When ClusterType=2, it is required to specify the ID of the exclusive cluster to which the job is submitted.
     */
    public readonly clusterId!: pulumi.Output<string | undefined>;
    /**
     * The type of the cluster. 1 indicates shared cluster, and 2 indicates exclusive cluster.
     */
    public readonly clusterType!: pulumi.Output<number>;
    /**
     * Set the memory specification of each CU, in GB. It supports 2, 4, 8, and 16 (which needs to apply for the whitelist
     * before use). The default is 4, that is, 1 CU corresponds to 4 GB of running memory.
     */
    public readonly cuMem!: pulumi.Output<number | undefined>;
    /**
     * The Flink version that the job runs.
     */
    public readonly flinkVersion!: pulumi.Output<string | undefined>;
    /**
     * The folder ID to which the job name belongs. The root directory is root.
     */
    public readonly folderId!: pulumi.Output<string | undefined>;
    /**
     * The type of the job. 1 indicates SQL job, and 2 indicates JAR job.
     */
    public readonly jobType!: pulumi.Output<number>;
    /**
     * The name of the job. It can be composed of Chinese, English, numbers, hyphens (-), underscores (_), and periods (.), and
     * the length cannot exceed 50 characters. Note that the job name cannot be the same as an existing job.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The remark information of the job. It can be set arbitrarily.
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * The workspace SerialId.
     */
    public readonly workSpaceId!: pulumi.Output<string | undefined>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobArgs | JobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as JobState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["clusterType"] = state ? state.clusterType : undefined;
            resourceInputs["cuMem"] = state ? state.cuMem : undefined;
            resourceInputs["flinkVersion"] = state ? state.flinkVersion : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["jobType"] = state ? state.jobType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["workSpaceId"] = state ? state.workSpaceId : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            if ((!args || args.clusterType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterType'");
            }
            if ((!args || args.jobType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobType'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["clusterType"] = args ? args.clusterType : undefined;
            resourceInputs["cuMem"] = args ? args.cuMem : undefined;
            resourceInputs["flinkVersion"] = args ? args.flinkVersion : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["jobType"] = args ? args.jobType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["workSpaceId"] = args ? args.workSpaceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Job.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Job resources.
 */
export interface JobState {
    /**
     * When ClusterType=2, it is required to specify the ID of the exclusive cluster to which the job is submitted.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The type of the cluster. 1 indicates shared cluster, and 2 indicates exclusive cluster.
     */
    clusterType?: pulumi.Input<number>;
    /**
     * Set the memory specification of each CU, in GB. It supports 2, 4, 8, and 16 (which needs to apply for the whitelist
     * before use). The default is 4, that is, 1 CU corresponds to 4 GB of running memory.
     */
    cuMem?: pulumi.Input<number>;
    /**
     * The Flink version that the job runs.
     */
    flinkVersion?: pulumi.Input<string>;
    /**
     * The folder ID to which the job name belongs. The root directory is root.
     */
    folderId?: pulumi.Input<string>;
    /**
     * The type of the job. 1 indicates SQL job, and 2 indicates JAR job.
     */
    jobType?: pulumi.Input<number>;
    /**
     * The name of the job. It can be composed of Chinese, English, numbers, hyphens (-), underscores (_), and periods (.), and
     * the length cannot exceed 50 characters. Note that the job name cannot be the same as an existing job.
     */
    name?: pulumi.Input<string>;
    /**
     * The remark information of the job. It can be set arbitrarily.
     */
    remark?: pulumi.Input<string>;
    /**
     * The workspace SerialId.
     */
    workSpaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * When ClusterType=2, it is required to specify the ID of the exclusive cluster to which the job is submitted.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The type of the cluster. 1 indicates shared cluster, and 2 indicates exclusive cluster.
     */
    clusterType: pulumi.Input<number>;
    /**
     * Set the memory specification of each CU, in GB. It supports 2, 4, 8, and 16 (which needs to apply for the whitelist
     * before use). The default is 4, that is, 1 CU corresponds to 4 GB of running memory.
     */
    cuMem?: pulumi.Input<number>;
    /**
     * The Flink version that the job runs.
     */
    flinkVersion?: pulumi.Input<string>;
    /**
     * The folder ID to which the job name belongs. The root directory is root.
     */
    folderId?: pulumi.Input<string>;
    /**
     * The type of the job. 1 indicates SQL job, and 2 indicates JAR job.
     */
    jobType: pulumi.Input<number>;
    /**
     * The name of the job. It can be composed of Chinese, English, numbers, hyphens (-), underscores (_), and periods (.), and
     * the length cannot exceed 50 characters. Note that the job name cannot be the same as an existing job.
     */
    name?: pulumi.Input<string>;
    /**
     * The remark information of the job. It can be set arbitrarily.
     */
    remark?: pulumi.Input<string>;
    /**
     * The workspace SerialId.
     */
    workSpaceId?: pulumi.Input<string>;
}