// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class GlobalDomain extends pulumi.CustomResource {
    /**
     * Get an existing GlobalDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GlobalDomainState, opts?: pulumi.CustomResourceOptions): GlobalDomain {
        return new GlobalDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Gaap/globalDomain:GlobalDomain';

    /**
     * Returns true if the given object is an instance of GlobalDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GlobalDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GlobalDomain.__pulumiType;
    }

    /**
     * alias.
     */
    public readonly alias!: pulumi.Output<string | undefined>;
    /**
     * Domain name default entry.
     */
    public readonly defaultValue!: pulumi.Output<string>;
    /**
     * Domain Name Project ID.
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * Global domain statue. Available values: open and close, default is open.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Instance tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a GlobalDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GlobalDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GlobalDomainArgs | GlobalDomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GlobalDomainState | undefined;
            resourceInputs["alias"] = state ? state.alias : undefined;
            resourceInputs["defaultValue"] = state ? state.defaultValue : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as GlobalDomainArgs | undefined;
            if ((!args || args.defaultValue === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultValue'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["defaultValue"] = args ? args.defaultValue : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GlobalDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GlobalDomain resources.
 */
export interface GlobalDomainState {
    /**
     * alias.
     */
    alias?: pulumi.Input<string>;
    /**
     * Domain name default entry.
     */
    defaultValue?: pulumi.Input<string>;
    /**
     * Domain Name Project ID.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Global domain statue. Available values: open and close, default is open.
     */
    status?: pulumi.Input<string>;
    /**
     * Instance tags.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a GlobalDomain resource.
 */
export interface GlobalDomainArgs {
    /**
     * alias.
     */
    alias?: pulumi.Input<string>;
    /**
     * Domain name default entry.
     */
    defaultValue: pulumi.Input<string>;
    /**
     * Domain Name Project ID.
     */
    projectId: pulumi.Input<number>;
    /**
     * Global domain statue. Available values: open and close, default is open.
     */
    status?: pulumi.Input<string>;
    /**
     * Instance tags.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}