// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class GrafanaEnvConfig extends pulumi.CustomResource {
    /**
     * Get an existing GrafanaEnvConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GrafanaEnvConfigState, opts?: pulumi.CustomResourceOptions): GrafanaEnvConfig {
        return new GrafanaEnvConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Monitor/grafanaEnvConfig:GrafanaEnvConfig';

    /**
     * Returns true if the given object is an instance of GrafanaEnvConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GrafanaEnvConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GrafanaEnvConfig.__pulumiType;
    }

    /**
     * Environment variables.
     */
    public readonly envs!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Grafana instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a GrafanaEnvConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GrafanaEnvConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GrafanaEnvConfigArgs | GrafanaEnvConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GrafanaEnvConfigState | undefined;
            resourceInputs["envs"] = state ? state.envs : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as GrafanaEnvConfigArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["envs"] = args ? args.envs : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GrafanaEnvConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GrafanaEnvConfig resources.
 */
export interface GrafanaEnvConfigState {
    /**
     * Environment variables.
     */
    envs?: pulumi.Input<{[key: string]: any}>;
    /**
     * Grafana instance ID.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GrafanaEnvConfig resource.
 */
export interface GrafanaEnvConfigArgs {
    /**
     * Environment variables.
     */
    envs?: pulumi.Input<{[key: string]: any}>;
    /**
     * Grafana instance ID.
     */
    instanceId: pulumi.Input<string>;
}