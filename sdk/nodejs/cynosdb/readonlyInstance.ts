// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ReadonlyInstance extends pulumi.CustomResource {
    /**
     * Get an existing ReadonlyInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReadonlyInstanceState, opts?: pulumi.CustomResourceOptions): ReadonlyInstance {
        return new ReadonlyInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cynosdb/readonlyInstance:ReadonlyInstance';

    /**
     * Returns true if the given object is an instance of ReadonlyInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReadonlyInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReadonlyInstance.__pulumiType;
    }

    /**
     * Cluster ID which the readonly instance belongs to.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted
     * instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
     */
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    /**
     * The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster.
     * Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade
     * from console.
     */
    public readonly instanceCpuCore!: pulumi.Output<number | undefined>;
    /**
     * Duration time for maintenance, unit in second. `3600` by default.
     */
    public readonly instanceMaintainDuration!: pulumi.Output<number | undefined>;
    /**
     * Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
     */
    public readonly instanceMaintainStartTime!: pulumi.Output<number | undefined>;
    /**
     * Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
     */
    public readonly instanceMaintainWeekdays!: pulumi.Output<string[]>;
    /**
     * Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of
     * this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
     */
    public readonly instanceMemorySize!: pulumi.Output<number | undefined>;
    /**
     * Name of instance.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * Status of the instance.
     */
    public /*out*/ readonly instanceStatus!: pulumi.Output<string>;
    /**
     * Storage size of the instance, unit in GB.
     */
    public /*out*/ readonly instanceStorageSize!: pulumi.Output<number>;
    /**
     * ID of the subnet within this VPC.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * ID of the VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a ReadonlyInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReadonlyInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReadonlyInstanceArgs | ReadonlyInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReadonlyInstanceState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["forceDelete"] = state ? state.forceDelete : undefined;
            resourceInputs["instanceCpuCore"] = state ? state.instanceCpuCore : undefined;
            resourceInputs["instanceMaintainDuration"] = state ? state.instanceMaintainDuration : undefined;
            resourceInputs["instanceMaintainStartTime"] = state ? state.instanceMaintainStartTime : undefined;
            resourceInputs["instanceMaintainWeekdays"] = state ? state.instanceMaintainWeekdays : undefined;
            resourceInputs["instanceMemorySize"] = state ? state.instanceMemorySize : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceStatus"] = state ? state.instanceStatus : undefined;
            resourceInputs["instanceStorageSize"] = state ? state.instanceStorageSize : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as ReadonlyInstanceArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["forceDelete"] = args ? args.forceDelete : undefined;
            resourceInputs["instanceCpuCore"] = args ? args.instanceCpuCore : undefined;
            resourceInputs["instanceMaintainDuration"] = args ? args.instanceMaintainDuration : undefined;
            resourceInputs["instanceMaintainStartTime"] = args ? args.instanceMaintainStartTime : undefined;
            resourceInputs["instanceMaintainWeekdays"] = args ? args.instanceMaintainWeekdays : undefined;
            resourceInputs["instanceMemorySize"] = args ? args.instanceMemorySize : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["instanceStatus"] = undefined /*out*/;
            resourceInputs["instanceStorageSize"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReadonlyInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReadonlyInstance resources.
 */
export interface ReadonlyInstanceState {
    /**
     * Cluster ID which the readonly instance belongs to.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted
     * instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster.
     * Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade
     * from console.
     */
    instanceCpuCore?: pulumi.Input<number>;
    /**
     * Duration time for maintenance, unit in second. `3600` by default.
     */
    instanceMaintainDuration?: pulumi.Input<number>;
    /**
     * Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
     */
    instanceMaintainStartTime?: pulumi.Input<number>;
    /**
     * Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
     */
    instanceMaintainWeekdays?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of
     * this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
     */
    instanceMemorySize?: pulumi.Input<number>;
    /**
     * Name of instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Status of the instance.
     */
    instanceStatus?: pulumi.Input<string>;
    /**
     * Storage size of the instance, unit in GB.
     */
    instanceStorageSize?: pulumi.Input<number>;
    /**
     * ID of the subnet within this VPC.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * ID of the VPC.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReadonlyInstance resource.
 */
export interface ReadonlyInstanceArgs {
    /**
     * Cluster ID which the readonly instance belongs to.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted
     * instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster.
     * Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade
     * from console.
     */
    instanceCpuCore?: pulumi.Input<number>;
    /**
     * Duration time for maintenance, unit in second. `3600` by default.
     */
    instanceMaintainDuration?: pulumi.Input<number>;
    /**
     * Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
     */
    instanceMaintainStartTime?: pulumi.Input<number>;
    /**
     * Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
     */
    instanceMaintainWeekdays?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of
     * this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
     */
    instanceMemorySize?: pulumi.Input<number>;
    /**
     * Name of instance.
     */
    instanceName: pulumi.Input<string>;
    /**
     * ID of the subnet within this VPC.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * ID of the VPC.
     */
    vpcId?: pulumi.Input<string>;
}