// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class EmbedTokenApply extends pulumi.CustomResource {
    /**
     * Get an existing EmbedTokenApply resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EmbedTokenApplyState, opts?: pulumi.CustomResourceOptions): EmbedTokenApply {
        return new EmbedTokenApply(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Bi/embedTokenApply:EmbedTokenApply';

    /**
     * Returns true if the given object is an instance of EmbedTokenApply.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EmbedTokenApply {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EmbedTokenApply.__pulumiType;
    }

    /**
     * Create the generated token.
     */
    public /*out*/ readonly biToken!: pulumi.Output<string>;
    /**
     * Create time.
     */
    public /*out*/ readonly createAt!: pulumi.Output<string>;
    /**
     * Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
     */
    public readonly expireTime!: pulumi.Output<string | undefined>;
    /**
     * Sharing page id, this is empty value 0 when embedding the board.
     */
    public readonly pageId!: pulumi.Output<number | undefined>;
    /**
     * Share project id.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * Page means embedding the page, and panel means embedding the entire board.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
     */
    public readonly ticketNum!: pulumi.Output<number | undefined>;
    /**
     * Upadte time.
     */
    public /*out*/ readonly udpateAt!: pulumi.Output<string>;
    /**
     * User enterprise ID (for multi-user only).
     */
    public readonly userCorpId!: pulumi.Output<string | undefined>;
    /**
     * UserId (for multi-user only).
     */
    public readonly userId!: pulumi.Output<string | undefined>;

    /**
     * Create a EmbedTokenApply resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EmbedTokenApplyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EmbedTokenApplyArgs | EmbedTokenApplyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EmbedTokenApplyState | undefined;
            resourceInputs["biToken"] = state ? state.biToken : undefined;
            resourceInputs["createAt"] = state ? state.createAt : undefined;
            resourceInputs["expireTime"] = state ? state.expireTime : undefined;
            resourceInputs["pageId"] = state ? state.pageId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["ticketNum"] = state ? state.ticketNum : undefined;
            resourceInputs["udpateAt"] = state ? state.udpateAt : undefined;
            resourceInputs["userCorpId"] = state ? state.userCorpId : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as EmbedTokenApplyArgs | undefined;
            resourceInputs["expireTime"] = args ? args.expireTime : undefined;
            resourceInputs["pageId"] = args ? args.pageId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["ticketNum"] = args ? args.ticketNum : undefined;
            resourceInputs["userCorpId"] = args ? args.userCorpId : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["biToken"] = undefined /*out*/;
            resourceInputs["createAt"] = undefined /*out*/;
            resourceInputs["udpateAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EmbedTokenApply.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EmbedTokenApply resources.
 */
export interface EmbedTokenApplyState {
    /**
     * Create the generated token.
     */
    biToken?: pulumi.Input<string>;
    /**
     * Create time.
     */
    createAt?: pulumi.Input<string>;
    /**
     * Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * Sharing page id, this is empty value 0 when embedding the board.
     */
    pageId?: pulumi.Input<number>;
    /**
     * Share project id.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Page means embedding the page, and panel means embedding the entire board.
     */
    scope?: pulumi.Input<string>;
    /**
     * Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
     */
    ticketNum?: pulumi.Input<number>;
    /**
     * Upadte time.
     */
    udpateAt?: pulumi.Input<string>;
    /**
     * User enterprise ID (for multi-user only).
     */
    userCorpId?: pulumi.Input<string>;
    /**
     * UserId (for multi-user only).
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EmbedTokenApply resource.
 */
export interface EmbedTokenApplyArgs {
    /**
     * Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * Sharing page id, this is empty value 0 when embedding the board.
     */
    pageId?: pulumi.Input<number>;
    /**
     * Share project id.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Page means embedding the page, and panel means embedding the entire board.
     */
    scope?: pulumi.Input<string>;
    /**
     * Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
     */
    ticketNum?: pulumi.Input<number>;
    /**
     * User enterprise ID (for multi-user only).
     */
    userCorpId?: pulumi.Input<string>;
    /**
     * UserId (for multi-user only).
     */
    userId?: pulumi.Input<string>;
}
