// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class MigrateJob extends pulumi.CustomResource {
    /**
     * Get an existing MigrateJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MigrateJobState, opts?: pulumi.CustomResourceOptions): MigrateJob {
        return new MigrateJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dts/migrateJob:MigrateJob';

    /**
     * Returns true if the given object is an instance of MigrateJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MigrateJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MigrateJob.__pulumiType;
    }

    /**
     * AutoRetryTimeRangeMinutes.
     */
    public readonly autoRetryTimeRangeMinutes!: pulumi.Output<number | undefined>;
    /**
     * DstInfo.
     */
    public readonly dstInfo!: pulumi.Output<outputs.Dts.MigrateJobDstInfo>;
    /**
     * ExpectRunTime.
     */
    public readonly expectRunTime!: pulumi.Output<string>;
    /**
     * Migration job configuration options, used to describe how the task performs migration.
     */
    public readonly migrateOption!: pulumi.Output<outputs.Dts.MigrateJobMigrateOption>;
    /**
     * Run Mode. eg:immediate,timed.
     */
    public readonly runMode!: pulumi.Output<string>;
    /**
     * Migrate service Id from `tencentcloud_dts_migrate_service`.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * SrcInfo.
     */
    public readonly srcInfo!: pulumi.Output<outputs.Dts.MigrateJobSrcInfo>;
    /**
     * Migrate job status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a MigrateJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MigrateJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MigrateJobArgs | MigrateJobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MigrateJobState | undefined;
            resourceInputs["autoRetryTimeRangeMinutes"] = state ? state.autoRetryTimeRangeMinutes : undefined;
            resourceInputs["dstInfo"] = state ? state.dstInfo : undefined;
            resourceInputs["expectRunTime"] = state ? state.expectRunTime : undefined;
            resourceInputs["migrateOption"] = state ? state.migrateOption : undefined;
            resourceInputs["runMode"] = state ? state.runMode : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["srcInfo"] = state ? state.srcInfo : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as MigrateJobArgs | undefined;
            if ((!args || args.dstInfo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstInfo'");
            }
            if ((!args || args.migrateOption === undefined) && !opts.urn) {
                throw new Error("Missing required property 'migrateOption'");
            }
            if ((!args || args.runMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runMode'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            if ((!args || args.srcInfo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcInfo'");
            }
            resourceInputs["autoRetryTimeRangeMinutes"] = args ? args.autoRetryTimeRangeMinutes : undefined;
            resourceInputs["dstInfo"] = args ? args.dstInfo : undefined;
            resourceInputs["expectRunTime"] = args ? args.expectRunTime : undefined;
            resourceInputs["migrateOption"] = args ? args.migrateOption : undefined;
            resourceInputs["runMode"] = args ? args.runMode : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["srcInfo"] = args ? args.srcInfo : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MigrateJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MigrateJob resources.
 */
export interface MigrateJobState {
    /**
     * AutoRetryTimeRangeMinutes.
     */
    autoRetryTimeRangeMinutes?: pulumi.Input<number>;
    /**
     * DstInfo.
     */
    dstInfo?: pulumi.Input<inputs.Dts.MigrateJobDstInfo>;
    /**
     * ExpectRunTime.
     */
    expectRunTime?: pulumi.Input<string>;
    /**
     * Migration job configuration options, used to describe how the task performs migration.
     */
    migrateOption?: pulumi.Input<inputs.Dts.MigrateJobMigrateOption>;
    /**
     * Run Mode. eg:immediate,timed.
     */
    runMode?: pulumi.Input<string>;
    /**
     * Migrate service Id from `tencentcloud_dts_migrate_service`.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * SrcInfo.
     */
    srcInfo?: pulumi.Input<inputs.Dts.MigrateJobSrcInfo>;
    /**
     * Migrate job status.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MigrateJob resource.
 */
export interface MigrateJobArgs {
    /**
     * AutoRetryTimeRangeMinutes.
     */
    autoRetryTimeRangeMinutes?: pulumi.Input<number>;
    /**
     * DstInfo.
     */
    dstInfo: pulumi.Input<inputs.Dts.MigrateJobDstInfo>;
    /**
     * ExpectRunTime.
     */
    expectRunTime?: pulumi.Input<string>;
    /**
     * Migration job configuration options, used to describe how the task performs migration.
     */
    migrateOption: pulumi.Input<inputs.Dts.MigrateJobMigrateOption>;
    /**
     * Run Mode. eg:immediate,timed.
     */
    runMode: pulumi.Input<string>;
    /**
     * Migrate service Id from `tencentcloud_dts_migrate_service`.
     */
    serviceId: pulumi.Input<string>;
    /**
     * SrcInfo.
     */
    srcInfo: pulumi.Input<inputs.Dts.MigrateJobSrcInfo>;
}