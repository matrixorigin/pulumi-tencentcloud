// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ModifyDiagDbInstanceOperation extends pulumi.CustomResource {
    /**
     * Get an existing ModifyDiagDbInstanceOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ModifyDiagDbInstanceOperationState, opts?: pulumi.CustomResourceOptions): ModifyDiagDbInstanceOperation {
        return new ModifyDiagDbInstanceOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dbbrain/modifyDiagDbInstanceOperation:ModifyDiagDbInstanceOperation';

    /**
     * Returns true if the given object is an instance of ModifyDiagDbInstanceOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ModifyDiagDbInstanceOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ModifyDiagDbInstanceOperation.__pulumiType;
    }

    /**
     * Instance configuration, including inspection, overview switch, etc.
     */
    public readonly instanceConfs!: pulumi.Output<outputs.Dbbrain.ModifyDiagDbInstanceOperationInstanceConfs>;
    /**
     * Specifies the ID of the instance whose inspection status is changed.
     */
    public readonly instanceIds!: pulumi.Output<string[] | undefined>;
    /**
     * Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for
     * MySQL.
     */
    public readonly product!: pulumi.Output<string>;
    /**
     * Effective instance region, the value is All, which means all regions.
     */
    public readonly regions!: pulumi.Output<string | undefined>;

    /**
     * Create a ModifyDiagDbInstanceOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModifyDiagDbInstanceOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ModifyDiagDbInstanceOperationArgs | ModifyDiagDbInstanceOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ModifyDiagDbInstanceOperationState | undefined;
            resourceInputs["instanceConfs"] = state ? state.instanceConfs : undefined;
            resourceInputs["instanceIds"] = state ? state.instanceIds : undefined;
            resourceInputs["product"] = state ? state.product : undefined;
            resourceInputs["regions"] = state ? state.regions : undefined;
        } else {
            const args = argsOrState as ModifyDiagDbInstanceOperationArgs | undefined;
            if ((!args || args.instanceConfs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceConfs'");
            }
            if ((!args || args.product === undefined) && !opts.urn) {
                throw new Error("Missing required property 'product'");
            }
            resourceInputs["instanceConfs"] = args ? args.instanceConfs : undefined;
            resourceInputs["instanceIds"] = args ? args.instanceIds : undefined;
            resourceInputs["product"] = args ? args.product : undefined;
            resourceInputs["regions"] = args ? args.regions : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ModifyDiagDbInstanceOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ModifyDiagDbInstanceOperation resources.
 */
export interface ModifyDiagDbInstanceOperationState {
    /**
     * Instance configuration, including inspection, overview switch, etc.
     */
    instanceConfs?: pulumi.Input<inputs.Dbbrain.ModifyDiagDbInstanceOperationInstanceConfs>;
    /**
     * Specifies the ID of the instance whose inspection status is changed.
     */
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for
     * MySQL.
     */
    product?: pulumi.Input<string>;
    /**
     * Effective instance region, the value is All, which means all regions.
     */
    regions?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ModifyDiagDbInstanceOperation resource.
 */
export interface ModifyDiagDbInstanceOperationArgs {
    /**
     * Instance configuration, including inspection, overview switch, etc.
     */
    instanceConfs: pulumi.Input<inputs.Dbbrain.ModifyDiagDbInstanceOperationInstanceConfs>;
    /**
     * Specifies the ID of the instance whose inspection status is changed.
     */
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for
     * MySQL.
     */
    product: pulumi.Input<string>;
    /**
     * Effective instance region, the value is All, which means all regions.
     */
    regions?: pulumi.Input<string>;
}