// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class RefreshNatDcRoute extends pulumi.CustomResource {
    /**
     * Get an existing RefreshNatDcRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RefreshNatDcRouteState, opts?: pulumi.CustomResourceOptions): RefreshNatDcRoute {
        return new RefreshNatDcRoute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Nat/refreshNatDcRoute:RefreshNatDcRoute';

    /**
     * Returns true if the given object is an instance of RefreshNatDcRoute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RefreshNatDcRoute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RefreshNatDcRoute.__pulumiType;
    }

    /**
     * Whether to pre-refresh, valid values: True:yes, False:no.
     */
    public readonly dryRun!: pulumi.Output<boolean>;
    /**
     * Unique identifier of Nat Gateway.
     */
    public readonly natGatewayId!: pulumi.Output<string>;
    /**
     * Unique identifier of Vpc.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a RefreshNatDcRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RefreshNatDcRouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RefreshNatDcRouteArgs | RefreshNatDcRouteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RefreshNatDcRouteState | undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["natGatewayId"] = state ? state.natGatewayId : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as RefreshNatDcRouteArgs | undefined;
            if ((!args || args.dryRun === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dryRun'");
            }
            if ((!args || args.natGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'natGatewayId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["natGatewayId"] = args ? args.natGatewayId : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RefreshNatDcRoute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RefreshNatDcRoute resources.
 */
export interface RefreshNatDcRouteState {
    /**
     * Whether to pre-refresh, valid values: True:yes, False:no.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Unique identifier of Nat Gateway.
     */
    natGatewayId?: pulumi.Input<string>;
    /**
     * Unique identifier of Vpc.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RefreshNatDcRoute resource.
 */
export interface RefreshNatDcRouteArgs {
    /**
     * Whether to pre-refresh, valid values: True:yes, False:no.
     */
    dryRun: pulumi.Input<boolean>;
    /**
     * Unique identifier of Nat Gateway.
     */
    natGatewayId: pulumi.Input<string>;
    /**
     * Unique identifier of Vpc.
     */
    vpcId: pulumi.Input<string>;
}