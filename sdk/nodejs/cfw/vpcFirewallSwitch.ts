// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class VpcFirewallSwitch extends pulumi.CustomResource {
    /**
     * Get an existing VpcFirewallSwitch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcFirewallSwitchState, opts?: pulumi.CustomResourceOptions): VpcFirewallSwitch {
        return new VpcFirewallSwitch(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cfw/vpcFirewallSwitch:VpcFirewallSwitch';

    /**
     * Returns true if the given object is an instance of VpcFirewallSwitch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcFirewallSwitch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcFirewallSwitch.__pulumiType;
    }

    /**
     * Turn the switch on or off. 0: turn off the switch; 1: Turn on the switch.
     */
    public readonly enable!: pulumi.Output<number>;
    /**
     * Firewall switch ID.
     */
    public readonly switchId!: pulumi.Output<string>;
    /**
     * Firewall instance id.
     */
    public readonly vpcInsId!: pulumi.Output<string>;

    /**
     * Create a VpcFirewallSwitch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcFirewallSwitchArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcFirewallSwitchArgs | VpcFirewallSwitchState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcFirewallSwitchState | undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
            resourceInputs["switchId"] = state ? state.switchId : undefined;
            resourceInputs["vpcInsId"] = state ? state.vpcInsId : undefined;
        } else {
            const args = argsOrState as VpcFirewallSwitchArgs | undefined;
            if ((!args || args.enable === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enable'");
            }
            if ((!args || args.switchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'switchId'");
            }
            if ((!args || args.vpcInsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcInsId'");
            }
            resourceInputs["enable"] = args ? args.enable : undefined;
            resourceInputs["switchId"] = args ? args.switchId : undefined;
            resourceInputs["vpcInsId"] = args ? args.vpcInsId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcFirewallSwitch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcFirewallSwitch resources.
 */
export interface VpcFirewallSwitchState {
    /**
     * Turn the switch on or off. 0: turn off the switch; 1: Turn on the switch.
     */
    enable?: pulumi.Input<number>;
    /**
     * Firewall switch ID.
     */
    switchId?: pulumi.Input<string>;
    /**
     * Firewall instance id.
     */
    vpcInsId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcFirewallSwitch resource.
 */
export interface VpcFirewallSwitchArgs {
    /**
     * Turn the switch on or off. 0: turn off the switch; 1: Turn on the switch.
     */
    enable: pulumi.Input<number>;
    /**
     * Firewall switch ID.
     */
    switchId: pulumi.Input<string>;
    /**
     * Firewall instance id.
     */
    vpcInsId: pulumi.Input<string>;
}
