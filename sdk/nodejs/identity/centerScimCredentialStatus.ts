// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class CenterScimCredentialStatus extends pulumi.CustomResource {
    /**
     * Get an existing CenterScimCredentialStatus resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CenterScimCredentialStatusState, opts?: pulumi.CustomResourceOptions): CenterScimCredentialStatus {
        return new CenterScimCredentialStatus(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Identity/centerScimCredentialStatus:CenterScimCredentialStatus';

    /**
     * Returns true if the given object is an instance of CenterScimCredentialStatus.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CenterScimCredentialStatus {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CenterScimCredentialStatus.__pulumiType;
    }

    /**
     * SCIM key ID. scimcred-prefix and followed by 12 random digits/lowercase letters.
     */
    public readonly credentialId!: pulumi.Output<string>;
    /**
     * SCIM key status. Enabled-enabled. Disabled-disabled.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Space ID. z-prefix starts with 12 random digits/lowercase letters.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a CenterScimCredentialStatus resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CenterScimCredentialStatusArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CenterScimCredentialStatusArgs | CenterScimCredentialStatusState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CenterScimCredentialStatusState | undefined;
            resourceInputs["credentialId"] = state ? state.credentialId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as CenterScimCredentialStatusArgs | undefined;
            if ((!args || args.credentialId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'credentialId'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["credentialId"] = args ? args.credentialId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CenterScimCredentialStatus.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CenterScimCredentialStatus resources.
 */
export interface CenterScimCredentialStatusState {
    /**
     * SCIM key ID. scimcred-prefix and followed by 12 random digits/lowercase letters.
     */
    credentialId?: pulumi.Input<string>;
    /**
     * SCIM key status. Enabled-enabled. Disabled-disabled.
     */
    status?: pulumi.Input<string>;
    /**
     * Space ID. z-prefix starts with 12 random digits/lowercase letters.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CenterScimCredentialStatus resource.
 */
export interface CenterScimCredentialStatusArgs {
    /**
     * SCIM key ID. scimcred-prefix and followed by 12 random digits/lowercase letters.
     */
    credentialId: pulumi.Input<string>;
    /**
     * SCIM key status. Enabled-enabled. Disabled-disabled.
     */
    status: pulumi.Input<string>;
    /**
     * Space ID. z-prefix starts with 12 random digits/lowercase letters.
     */
    zoneId: pulumi.Input<string>;
}