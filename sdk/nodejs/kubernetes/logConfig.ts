// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class LogConfig extends pulumi.CustomResource {
    /**
     * Get an existing LogConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogConfigState, opts?: pulumi.CustomResourceOptions): LogConfig {
        return new LogConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Kubernetes/logConfig:LogConfig';

    /**
     * Returns true if the given object is an instance of LogConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogConfig.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The current cluster type supports tke and eks, default is tke.
     */
    public readonly clusterType!: pulumi.Output<string | undefined>;
    /**
     * JSON expression of log collection configuration.
     */
    public readonly logConfig!: pulumi.Output<string>;
    /**
     * Log config name.
     */
    public readonly logConfigName!: pulumi.Output<string>;
    /**
     * CLS log set ID.
     */
    public readonly logsetId!: pulumi.Output<string | undefined>;

    /**
     * Create a LogConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogConfigArgs | LogConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogConfigState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["clusterType"] = state ? state.clusterType : undefined;
            resourceInputs["logConfig"] = state ? state.logConfig : undefined;
            resourceInputs["logConfigName"] = state ? state.logConfigName : undefined;
            resourceInputs["logsetId"] = state ? state.logsetId : undefined;
        } else {
            const args = argsOrState as LogConfigArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.logConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logConfig'");
            }
            if ((!args || args.logConfigName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logConfigName'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["clusterType"] = args ? args.clusterType : undefined;
            resourceInputs["logConfig"] = args ? args.logConfig : undefined;
            resourceInputs["logConfigName"] = args ? args.logConfigName : undefined;
            resourceInputs["logsetId"] = args ? args.logsetId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogConfig resources.
 */
export interface LogConfigState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The current cluster type supports tke and eks, default is tke.
     */
    clusterType?: pulumi.Input<string>;
    /**
     * JSON expression of log collection configuration.
     */
    logConfig?: pulumi.Input<string>;
    /**
     * Log config name.
     */
    logConfigName?: pulumi.Input<string>;
    /**
     * CLS log set ID.
     */
    logsetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogConfig resource.
 */
export interface LogConfigArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The current cluster type supports tke and eks, default is tke.
     */
    clusterType?: pulumi.Input<string>;
    /**
     * JSON expression of log collection configuration.
     */
    logConfig: pulumi.Input<string>;
    /**
     * Log config name.
     */
    logConfigName: pulumi.Input<string>;
    /**
     * CLS log set ID.
     */
    logsetId?: pulumi.Input<string>;
}