// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class AccountPrivileges extends pulumi.CustomResource {
    /**
     * Get an existing AccountPrivileges resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountPrivilegesState, opts?: pulumi.CustomResourceOptions): AccountPrivileges {
        return new AccountPrivileges(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cynosdb/accountPrivileges:AccountPrivileges';

    /**
     * Returns true if the given object is an instance of AccountPrivileges.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountPrivileges {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountPrivileges.__pulumiType;
    }

    /**
     * Account.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Array of database permissions.
     */
    public readonly databasePrivileges!: pulumi.Output<outputs.Cynosdb.AccountPrivilegesDatabasePrivilege[] | undefined>;
    /**
     * Array of global permissions.
     */
    public readonly globalPrivileges!: pulumi.Output<string[]>;
    /**
     * Host, default `%`.
     */
    public readonly host!: pulumi.Output<string | undefined>;
    /**
     * array of table permissions.
     */
    public readonly tablePrivileges!: pulumi.Output<outputs.Cynosdb.AccountPrivilegesTablePrivilege[] | undefined>;

    /**
     * Create a AccountPrivileges resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountPrivilegesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountPrivilegesArgs | AccountPrivilegesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountPrivilegesState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["databasePrivileges"] = state ? state.databasePrivileges : undefined;
            resourceInputs["globalPrivileges"] = state ? state.globalPrivileges : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["tablePrivileges"] = state ? state.tablePrivileges : undefined;
        } else {
            const args = argsOrState as AccountPrivilegesArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.globalPrivileges === undefined) && !opts.urn) {
                throw new Error("Missing required property 'globalPrivileges'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["databasePrivileges"] = args ? args.databasePrivileges : undefined;
            resourceInputs["globalPrivileges"] = args ? args.globalPrivileges : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["tablePrivileges"] = args ? args.tablePrivileges : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccountPrivileges.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountPrivileges resources.
 */
export interface AccountPrivilegesState {
    /**
     * Account.
     */
    accountName?: pulumi.Input<string>;
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Array of database permissions.
     */
    databasePrivileges?: pulumi.Input<pulumi.Input<inputs.Cynosdb.AccountPrivilegesDatabasePrivilege>[]>;
    /**
     * Array of global permissions.
     */
    globalPrivileges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Host, default `%`.
     */
    host?: pulumi.Input<string>;
    /**
     * array of table permissions.
     */
    tablePrivileges?: pulumi.Input<pulumi.Input<inputs.Cynosdb.AccountPrivilegesTablePrivilege>[]>;
}

/**
 * The set of arguments for constructing a AccountPrivileges resource.
 */
export interface AccountPrivilegesArgs {
    /**
     * Account.
     */
    accountName: pulumi.Input<string>;
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Array of database permissions.
     */
    databasePrivileges?: pulumi.Input<pulumi.Input<inputs.Cynosdb.AccountPrivilegesDatabasePrivilege>[]>;
    /**
     * Array of global permissions.
     */
    globalPrivileges: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Host, default `%`.
     */
    host?: pulumi.Input<string>;
    /**
     * array of table permissions.
     */
    tablePrivileges?: pulumi.Input<pulumi.Input<inputs.Cynosdb.AccountPrivilegesTablePrivilege>[]>;
}