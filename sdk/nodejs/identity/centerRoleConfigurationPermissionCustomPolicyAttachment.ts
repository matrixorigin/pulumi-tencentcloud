// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class CenterRoleConfigurationPermissionCustomPolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing CenterRoleConfigurationPermissionCustomPolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CenterRoleConfigurationPermissionCustomPolicyAttachmentState, opts?: pulumi.CustomResourceOptions): CenterRoleConfigurationPermissionCustomPolicyAttachment {
        return new CenterRoleConfigurationPermissionCustomPolicyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Identity/centerRoleConfigurationPermissionCustomPolicyAttachment:CenterRoleConfigurationPermissionCustomPolicyAttachment';

    /**
     * Returns true if the given object is an instance of CenterRoleConfigurationPermissionCustomPolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CenterRoleConfigurationPermissionCustomPolicyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CenterRoleConfigurationPermissionCustomPolicyAttachment.__pulumiType;
    }

    /**
     * Role policy add time.
     */
    public /*out*/ readonly addTime!: pulumi.Output<string>;
    /**
     * Permission configuration ID.
     */
    public readonly roleConfigurationId!: pulumi.Output<string>;
    /**
     * Role policy document.
     */
    public readonly rolePolicyDocument!: pulumi.Output<string>;
    /**
     * Role policy name.
     */
    public readonly rolePolicyName!: pulumi.Output<string>;
    /**
     * Role policy type.
     */
    public /*out*/ readonly rolePolicyType!: pulumi.Output<string>;
    /**
     * Space ID.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a CenterRoleConfigurationPermissionCustomPolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CenterRoleConfigurationPermissionCustomPolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CenterRoleConfigurationPermissionCustomPolicyAttachmentArgs | CenterRoleConfigurationPermissionCustomPolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CenterRoleConfigurationPermissionCustomPolicyAttachmentState | undefined;
            resourceInputs["addTime"] = state ? state.addTime : undefined;
            resourceInputs["roleConfigurationId"] = state ? state.roleConfigurationId : undefined;
            resourceInputs["rolePolicyDocument"] = state ? state.rolePolicyDocument : undefined;
            resourceInputs["rolePolicyName"] = state ? state.rolePolicyName : undefined;
            resourceInputs["rolePolicyType"] = state ? state.rolePolicyType : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as CenterRoleConfigurationPermissionCustomPolicyAttachmentArgs | undefined;
            if ((!args || args.roleConfigurationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleConfigurationId'");
            }
            if ((!args || args.rolePolicyDocument === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rolePolicyDocument'");
            }
            if ((!args || args.rolePolicyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rolePolicyName'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["roleConfigurationId"] = args ? args.roleConfigurationId : undefined;
            resourceInputs["rolePolicyDocument"] = args ? args.rolePolicyDocument : undefined;
            resourceInputs["rolePolicyName"] = args ? args.rolePolicyName : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["addTime"] = undefined /*out*/;
            resourceInputs["rolePolicyType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CenterRoleConfigurationPermissionCustomPolicyAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CenterRoleConfigurationPermissionCustomPolicyAttachment resources.
 */
export interface CenterRoleConfigurationPermissionCustomPolicyAttachmentState {
    /**
     * Role policy add time.
     */
    addTime?: pulumi.Input<string>;
    /**
     * Permission configuration ID.
     */
    roleConfigurationId?: pulumi.Input<string>;
    /**
     * Role policy document.
     */
    rolePolicyDocument?: pulumi.Input<string>;
    /**
     * Role policy name.
     */
    rolePolicyName?: pulumi.Input<string>;
    /**
     * Role policy type.
     */
    rolePolicyType?: pulumi.Input<string>;
    /**
     * Space ID.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CenterRoleConfigurationPermissionCustomPolicyAttachment resource.
 */
export interface CenterRoleConfigurationPermissionCustomPolicyAttachmentArgs {
    /**
     * Permission configuration ID.
     */
    roleConfigurationId: pulumi.Input<string>;
    /**
     * Role policy document.
     */
    rolePolicyDocument: pulumi.Input<string>;
    /**
     * Role policy name.
     */
    rolePolicyName: pulumi.Input<string>;
    /**
     * Space ID.
     */
    zoneId: pulumi.Input<string>;
}