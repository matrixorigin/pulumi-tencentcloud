// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class StopLogstashPipelineOperation extends pulumi.CustomResource {
    /**
     * Get an existing StopLogstashPipelineOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StopLogstashPipelineOperationState, opts?: pulumi.CustomResourceOptions): StopLogstashPipelineOperation {
        return new StopLogstashPipelineOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Elasticsearch/stopLogstashPipelineOperation:StopLogstashPipelineOperation';

    /**
     * Returns true if the given object is an instance of StopLogstashPipelineOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StopLogstashPipelineOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StopLogstashPipelineOperation.__pulumiType;
    }

    /**
     * Instance id.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Pipeline id.
     */
    public readonly pipelineId!: pulumi.Output<string>;

    /**
     * Create a StopLogstashPipelineOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StopLogstashPipelineOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StopLogstashPipelineOperationArgs | StopLogstashPipelineOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StopLogstashPipelineOperationState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["pipelineId"] = state ? state.pipelineId : undefined;
        } else {
            const args = argsOrState as StopLogstashPipelineOperationArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.pipelineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pipelineId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["pipelineId"] = args ? args.pipelineId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StopLogstashPipelineOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StopLogstashPipelineOperation resources.
 */
export interface StopLogstashPipelineOperationState {
    /**
     * Instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Pipeline id.
     */
    pipelineId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StopLogstashPipelineOperation resource.
 */
export interface StopLogstashPipelineOperationArgs {
    /**
     * Instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Pipeline id.
     */
    pipelineId: pulumi.Input<string>;
}