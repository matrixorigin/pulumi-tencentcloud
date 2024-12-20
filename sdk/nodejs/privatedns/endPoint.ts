// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class EndPoint extends pulumi.CustomResource {
    /**
     * Get an existing EndPoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndPointState, opts?: pulumi.CustomResourceOptions): EndPoint {
        return new EndPoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:PrivateDns/endPoint:EndPoint';

    /**
     * Returns true if the given object is an instance of EndPoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndPoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndPoint.__pulumiType;
    }

    /**
     * Endpoint name.
     */
    public readonly endPointName!: pulumi.Output<string>;
    /**
     * Endpoint region, which should be consistent with the region of the endpoint service.
     */
    public readonly endPointRegion!: pulumi.Output<string>;
    /**
     * Endpoint service ID (namely, VPC endpoint service ID).
     */
    public readonly endPointServiceId!: pulumi.Output<string>;
    /**
     * Vip list of endpoint.
     */
    public /*out*/ readonly endPointVipSets!: pulumi.Output<string[]>;
    /**
     * Number of endpoint IP addresses.
     */
    public readonly ipNum!: pulumi.Output<number>;

    /**
     * Create a EndPoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndPointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndPointArgs | EndPointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndPointState | undefined;
            resourceInputs["endPointName"] = state ? state.endPointName : undefined;
            resourceInputs["endPointRegion"] = state ? state.endPointRegion : undefined;
            resourceInputs["endPointServiceId"] = state ? state.endPointServiceId : undefined;
            resourceInputs["endPointVipSets"] = state ? state.endPointVipSets : undefined;
            resourceInputs["ipNum"] = state ? state.ipNum : undefined;
        } else {
            const args = argsOrState as EndPointArgs | undefined;
            if ((!args || args.endPointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endPointName'");
            }
            if ((!args || args.endPointRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endPointRegion'");
            }
            if ((!args || args.endPointServiceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endPointServiceId'");
            }
            resourceInputs["endPointName"] = args ? args.endPointName : undefined;
            resourceInputs["endPointRegion"] = args ? args.endPointRegion : undefined;
            resourceInputs["endPointServiceId"] = args ? args.endPointServiceId : undefined;
            resourceInputs["ipNum"] = args ? args.ipNum : undefined;
            resourceInputs["endPointVipSets"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EndPoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndPoint resources.
 */
export interface EndPointState {
    /**
     * Endpoint name.
     */
    endPointName?: pulumi.Input<string>;
    /**
     * Endpoint region, which should be consistent with the region of the endpoint service.
     */
    endPointRegion?: pulumi.Input<string>;
    /**
     * Endpoint service ID (namely, VPC endpoint service ID).
     */
    endPointServiceId?: pulumi.Input<string>;
    /**
     * Vip list of endpoint.
     */
    endPointVipSets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of endpoint IP addresses.
     */
    ipNum?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a EndPoint resource.
 */
export interface EndPointArgs {
    /**
     * Endpoint name.
     */
    endPointName: pulumi.Input<string>;
    /**
     * Endpoint region, which should be consistent with the region of the endpoint service.
     */
    endPointRegion: pulumi.Input<string>;
    /**
     * Endpoint service ID (namely, VPC endpoint service ID).
     */
    endPointServiceId: pulumi.Input<string>;
    /**
     * Number of endpoint IP addresses.
     */
    ipNum?: pulumi.Input<number>;
}
