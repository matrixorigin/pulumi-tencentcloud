// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class DbDiagReportTask extends pulumi.CustomResource {
    /**
     * Get an existing DbDiagReportTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DbDiagReportTaskState, opts?: pulumi.CustomResourceOptions): DbDiagReportTask {
        return new DbDiagReportTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dbbrain/dbDiagReportTask:DbDiagReportTask';

    /**
     * Returns true if the given object is an instance of DbDiagReportTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DbDiagReportTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DbDiagReportTask.__pulumiType;
    }

    /**
     * An array of contact group IDs to receive mail from.
     */
    public readonly contactGroups!: pulumi.Output<number[] | undefined>;
    /**
     * An array of contact IDs to receive emails from.
     */
    public readonly contactPeople!: pulumi.Output<number[] | undefined>;
    /**
     * End time, such as 2020-11-09T14:00:00+08:00.
     */
    public readonly endTime!: pulumi.Output<string>;
    /**
     * instance id.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for
     * MySQL.
     */
    public readonly product!: pulumi.Output<string>;
    /**
     * Whether to send mail: 0 - no, 1 - yes.
     */
    public readonly sendMailFlag!: pulumi.Output<number>;
    /**
     * Start time, such as 2020-11-08T14:00:00+08:00.
     */
    public readonly startTime!: pulumi.Output<string>;

    /**
     * Create a DbDiagReportTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DbDiagReportTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DbDiagReportTaskArgs | DbDiagReportTaskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DbDiagReportTaskState | undefined;
            resourceInputs["contactGroups"] = state ? state.contactGroups : undefined;
            resourceInputs["contactPeople"] = state ? state.contactPeople : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["product"] = state ? state.product : undefined;
            resourceInputs["sendMailFlag"] = state ? state.sendMailFlag : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
        } else {
            const args = argsOrState as DbDiagReportTaskArgs | undefined;
            if ((!args || args.endTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endTime'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.product === undefined) && !opts.urn) {
                throw new Error("Missing required property 'product'");
            }
            if ((!args || args.sendMailFlag === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sendMailFlag'");
            }
            if ((!args || args.startTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startTime'");
            }
            resourceInputs["contactGroups"] = args ? args.contactGroups : undefined;
            resourceInputs["contactPeople"] = args ? args.contactPeople : undefined;
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["product"] = args ? args.product : undefined;
            resourceInputs["sendMailFlag"] = args ? args.sendMailFlag : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DbDiagReportTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DbDiagReportTask resources.
 */
export interface DbDiagReportTaskState {
    /**
     * An array of contact group IDs to receive mail from.
     */
    contactGroups?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * An array of contact IDs to receive emails from.
     */
    contactPeople?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * End time, such as 2020-11-09T14:00:00+08:00.
     */
    endTime?: pulumi.Input<string>;
    /**
     * instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for
     * MySQL.
     */
    product?: pulumi.Input<string>;
    /**
     * Whether to send mail: 0 - no, 1 - yes.
     */
    sendMailFlag?: pulumi.Input<number>;
    /**
     * Start time, such as 2020-11-08T14:00:00+08:00.
     */
    startTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DbDiagReportTask resource.
 */
export interface DbDiagReportTaskArgs {
    /**
     * An array of contact group IDs to receive mail from.
     */
    contactGroups?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * An array of contact IDs to receive emails from.
     */
    contactPeople?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * End time, such as 2020-11-09T14:00:00+08:00.
     */
    endTime: pulumi.Input<string>;
    /**
     * instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for
     * MySQL.
     */
    product: pulumi.Input<string>;
    /**
     * Whether to send mail: 0 - no, 1 - yes.
     */
    sendMailFlag: pulumi.Input<number>;
    /**
     * Start time, such as 2020-11-08T14:00:00+08:00.
     */
    startTime: pulumi.Input<string>;
}