// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ValidateTxtValueOperation extends pulumi.CustomResource {
    /**
     * Get an existing ValidateTxtValueOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ValidateTxtValueOperationState, opts?: pulumi.CustomResourceOptions): ValidateTxtValueOperation {
        return new ValidateTxtValueOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Subdomain/validateTxtValueOperation:ValidateTxtValueOperation';

    /**
     * Returns true if the given object is an instance of ValidateTxtValueOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ValidateTxtValueOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ValidateTxtValueOperation.__pulumiType;
    }

    /**
     * The domain name for which TXT records need to be added.
     */
    public /*out*/ readonly domain!: pulumi.Output<string>;
    /**
     * The subdomain to add Zone domain.
     */
    public readonly domainZone!: pulumi.Output<string>;
    /**
     * Record types need to be added.
     */
    public /*out*/ readonly recordType!: pulumi.Output<string>;
    /**
     * Host records that need to be added to TXT records.
     */
    public /*out*/ readonly subdomain!: pulumi.Output<string>;
    /**
     * The record value of the TXT record needs to be added.
     */
    public /*out*/ readonly value!: pulumi.Output<string>;

    /**
     * Create a ValidateTxtValueOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ValidateTxtValueOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ValidateTxtValueOperationArgs | ValidateTxtValueOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ValidateTxtValueOperationState | undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["domainZone"] = state ? state.domainZone : undefined;
            resourceInputs["recordType"] = state ? state.recordType : undefined;
            resourceInputs["subdomain"] = state ? state.subdomain : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ValidateTxtValueOperationArgs | undefined;
            if ((!args || args.domainZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainZone'");
            }
            resourceInputs["domainZone"] = args ? args.domainZone : undefined;
            resourceInputs["domain"] = undefined /*out*/;
            resourceInputs["recordType"] = undefined /*out*/;
            resourceInputs["subdomain"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ValidateTxtValueOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ValidateTxtValueOperation resources.
 */
export interface ValidateTxtValueOperationState {
    /**
     * The domain name for which TXT records need to be added.
     */
    domain?: pulumi.Input<string>;
    /**
     * The subdomain to add Zone domain.
     */
    domainZone?: pulumi.Input<string>;
    /**
     * Record types need to be added.
     */
    recordType?: pulumi.Input<string>;
    /**
     * Host records that need to be added to TXT records.
     */
    subdomain?: pulumi.Input<string>;
    /**
     * The record value of the TXT record needs to be added.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ValidateTxtValueOperation resource.
 */
export interface ValidateTxtValueOperationArgs {
    /**
     * The subdomain to add Zone domain.
     */
    domainZone: pulumi.Input<string>;
}