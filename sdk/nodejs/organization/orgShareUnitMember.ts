// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class OrgShareUnitMember extends pulumi.CustomResource {
    /**
     * Get an existing OrgShareUnitMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgShareUnitMemberState, opts?: pulumi.CustomResourceOptions): OrgShareUnitMember {
        return new OrgShareUnitMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Organization/orgShareUnitMember:OrgShareUnitMember';

    /**
     * Returns true if the given object is an instance of OrgShareUnitMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgShareUnitMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgShareUnitMember.__pulumiType;
    }

    /**
     * Shared unit region.
     */
    public readonly area!: pulumi.Output<string>;
    /**
     * Shared member list. Up to 10 items.
     */
    public readonly members!: pulumi.Output<outputs.Organization.OrgShareUnitMemberMember[]>;
    /**
     * Shared unit ID.
     */
    public readonly unitId!: pulumi.Output<string>;

    /**
     * Create a OrgShareUnitMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgShareUnitMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgShareUnitMemberArgs | OrgShareUnitMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgShareUnitMemberState | undefined;
            resourceInputs["area"] = state ? state.area : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["unitId"] = state ? state.unitId : undefined;
        } else {
            const args = argsOrState as OrgShareUnitMemberArgs | undefined;
            if ((!args || args.area === undefined) && !opts.urn) {
                throw new Error("Missing required property 'area'");
            }
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.unitId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'unitId'");
            }
            resourceInputs["area"] = args ? args.area : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["unitId"] = args ? args.unitId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrgShareUnitMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgShareUnitMember resources.
 */
export interface OrgShareUnitMemberState {
    /**
     * Shared unit region.
     */
    area?: pulumi.Input<string>;
    /**
     * Shared member list. Up to 10 items.
     */
    members?: pulumi.Input<pulumi.Input<inputs.Organization.OrgShareUnitMemberMember>[]>;
    /**
     * Shared unit ID.
     */
    unitId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrgShareUnitMember resource.
 */
export interface OrgShareUnitMemberArgs {
    /**
     * Shared unit region.
     */
    area: pulumi.Input<string>;
    /**
     * Shared member list. Up to 10 items.
     */
    members: pulumi.Input<pulumi.Input<inputs.Organization.OrgShareUnitMemberMember>[]>;
    /**
     * Shared unit ID.
     */
    unitId: pulumi.Input<string>;
}
