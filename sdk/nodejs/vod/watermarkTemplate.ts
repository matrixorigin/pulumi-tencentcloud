// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class WatermarkTemplate extends pulumi.CustomResource {
    /**
     * Get an existing WatermarkTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WatermarkTemplateState, opts?: pulumi.CustomResourceOptions): WatermarkTemplate {
        return new WatermarkTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Vod/watermarkTemplate:WatermarkTemplate';

    /**
     * Returns true if the given object is an instance of WatermarkTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WatermarkTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WatermarkTemplate.__pulumiType;
    }

    /**
     * Template description. Length limit: 256 characters.
     */
    public readonly comment!: pulumi.Output<string>;
    /**
     * Origin position. Valid values: TopLeft: the origin of coordinates is in the top-left corner of the video, and the origin
     * of the watermark is in the top-left corner of the image or text; TopRight: the origin of coordinates is in the top-right
     * corner of the video, and the origin of the watermark is in the top-right corner of the image or text; BottomLeft: the
     * origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left
     * corner of the image or text; BottomRight: the origin of coordinates is in the bottom-right corner of the video, and the
     * origin of the watermark is in the bottom-right corner of the image or text.Default value: TopLeft.
     */
    public readonly coordinateOrigin!: pulumi.Output<string>;
    /**
     * Image watermarking template. This field is required when `Type` is `image` and is invalid when `Type` is `text`.
     */
    public readonly imageTemplate!: pulumi.Output<outputs.Vod.WatermarkTemplateImageTemplate>;
    /**
     * Watermarking template name. Length limit: 64 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD
     * service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default
     * application or a newly created one), they must fill in this field with the application ID.
     */
    public readonly subAppId!: pulumi.Output<number>;
    /**
     * SVG watermarking template. This field is required when `Type` is `svg` and is invalid when `Type` is `image` or `text`.
     */
    public readonly svgTemplate!: pulumi.Output<outputs.Vod.WatermarkTemplateSvgTemplate>;
    /**
     * Text watermarking template. This field is required when `Type` is `text` and is invalid when `Type` is `image`.
     */
    public readonly textTemplate!: pulumi.Output<outputs.Vod.WatermarkTemplateTextTemplate>;
    /**
     * Watermarking type. Valid values: image: image watermark; text: text watermark; svg: SVG watermark.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The horizontal position of the origin of the watermark relative to the origin of coordinates of the video. % and px
     * formats are supported: If the string ends in %, the `XPos` of the watermark will be the specified percentage of the
     * video width; for example, `10%` means that `XPos` is 10% of the video width; If the string ends in px, the `XPos` of the
     * watermark will be the specified px; for example, `100px` means that `XPos` is 100 px.Default value: 0 px.
     */
    public readonly xPos!: pulumi.Output<string>;
    /**
     * The vertical position of the origin of the watermark relative to the origin of coordinates of the video. % and px
     * formats are supported: If the string ends in %, the `YPos` of the watermark will be the specified percentage of the
     * video height; for example, `10%` means that `YPos` is 10% of the video height; If the string ends in px, the `YPos` of
     * the watermark will be the specified px; for example, `100px` means that `YPos` is 100 px.Default value: 0 px.
     */
    public readonly yPos!: pulumi.Output<string>;

    /**
     * Create a WatermarkTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WatermarkTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WatermarkTemplateArgs | WatermarkTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WatermarkTemplateState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["coordinateOrigin"] = state ? state.coordinateOrigin : undefined;
            resourceInputs["imageTemplate"] = state ? state.imageTemplate : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["subAppId"] = state ? state.subAppId : undefined;
            resourceInputs["svgTemplate"] = state ? state.svgTemplate : undefined;
            resourceInputs["textTemplate"] = state ? state.textTemplate : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["xPos"] = state ? state.xPos : undefined;
            resourceInputs["yPos"] = state ? state.yPos : undefined;
        } else {
            const args = argsOrState as WatermarkTemplateArgs | undefined;
            if ((!args || args.subAppId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subAppId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["coordinateOrigin"] = args ? args.coordinateOrigin : undefined;
            resourceInputs["imageTemplate"] = args ? args.imageTemplate : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["subAppId"] = args ? args.subAppId : undefined;
            resourceInputs["svgTemplate"] = args ? args.svgTemplate : undefined;
            resourceInputs["textTemplate"] = args ? args.textTemplate : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["xPos"] = args ? args.xPos : undefined;
            resourceInputs["yPos"] = args ? args.yPos : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WatermarkTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WatermarkTemplate resources.
 */
export interface WatermarkTemplateState {
    /**
     * Template description. Length limit: 256 characters.
     */
    comment?: pulumi.Input<string>;
    /**
     * Origin position. Valid values: TopLeft: the origin of coordinates is in the top-left corner of the video, and the origin
     * of the watermark is in the top-left corner of the image or text; TopRight: the origin of coordinates is in the top-right
     * corner of the video, and the origin of the watermark is in the top-right corner of the image or text; BottomLeft: the
     * origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left
     * corner of the image or text; BottomRight: the origin of coordinates is in the bottom-right corner of the video, and the
     * origin of the watermark is in the bottom-right corner of the image or text.Default value: TopLeft.
     */
    coordinateOrigin?: pulumi.Input<string>;
    /**
     * Image watermarking template. This field is required when `Type` is `image` and is invalid when `Type` is `text`.
     */
    imageTemplate?: pulumi.Input<inputs.Vod.WatermarkTemplateImageTemplate>;
    /**
     * Watermarking template name. Length limit: 64 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD
     * service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default
     * application or a newly created one), they must fill in this field with the application ID.
     */
    subAppId?: pulumi.Input<number>;
    /**
     * SVG watermarking template. This field is required when `Type` is `svg` and is invalid when `Type` is `image` or `text`.
     */
    svgTemplate?: pulumi.Input<inputs.Vod.WatermarkTemplateSvgTemplate>;
    /**
     * Text watermarking template. This field is required when `Type` is `text` and is invalid when `Type` is `image`.
     */
    textTemplate?: pulumi.Input<inputs.Vod.WatermarkTemplateTextTemplate>;
    /**
     * Watermarking type. Valid values: image: image watermark; text: text watermark; svg: SVG watermark.
     */
    type?: pulumi.Input<string>;
    /**
     * The horizontal position of the origin of the watermark relative to the origin of coordinates of the video. % and px
     * formats are supported: If the string ends in %, the `XPos` of the watermark will be the specified percentage of the
     * video width; for example, `10%` means that `XPos` is 10% of the video width; If the string ends in px, the `XPos` of the
     * watermark will be the specified px; for example, `100px` means that `XPos` is 100 px.Default value: 0 px.
     */
    xPos?: pulumi.Input<string>;
    /**
     * The vertical position of the origin of the watermark relative to the origin of coordinates of the video. % and px
     * formats are supported: If the string ends in %, the `YPos` of the watermark will be the specified percentage of the
     * video height; for example, `10%` means that `YPos` is 10% of the video height; If the string ends in px, the `YPos` of
     * the watermark will be the specified px; for example, `100px` means that `YPos` is 100 px.Default value: 0 px.
     */
    yPos?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WatermarkTemplate resource.
 */
export interface WatermarkTemplateArgs {
    /**
     * Template description. Length limit: 256 characters.
     */
    comment?: pulumi.Input<string>;
    /**
     * Origin position. Valid values: TopLeft: the origin of coordinates is in the top-left corner of the video, and the origin
     * of the watermark is in the top-left corner of the image or text; TopRight: the origin of coordinates is in the top-right
     * corner of the video, and the origin of the watermark is in the top-right corner of the image or text; BottomLeft: the
     * origin of coordinates is in the bottom-left corner of the video, and the origin of the watermark is in the bottom-left
     * corner of the image or text; BottomRight: the origin of coordinates is in the bottom-right corner of the video, and the
     * origin of the watermark is in the bottom-right corner of the image or text.Default value: TopLeft.
     */
    coordinateOrigin?: pulumi.Input<string>;
    /**
     * Image watermarking template. This field is required when `Type` is `image` and is invalid when `Type` is `text`.
     */
    imageTemplate?: pulumi.Input<inputs.Vod.WatermarkTemplateImageTemplate>;
    /**
     * Watermarking template name. Length limit: 64 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD
     * service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default
     * application or a newly created one), they must fill in this field with the application ID.
     */
    subAppId: pulumi.Input<number>;
    /**
     * SVG watermarking template. This field is required when `Type` is `svg` and is invalid when `Type` is `image` or `text`.
     */
    svgTemplate?: pulumi.Input<inputs.Vod.WatermarkTemplateSvgTemplate>;
    /**
     * Text watermarking template. This field is required when `Type` is `text` and is invalid when `Type` is `image`.
     */
    textTemplate?: pulumi.Input<inputs.Vod.WatermarkTemplateTextTemplate>;
    /**
     * Watermarking type. Valid values: image: image watermark; text: text watermark; svg: SVG watermark.
     */
    type: pulumi.Input<string>;
    /**
     * The horizontal position of the origin of the watermark relative to the origin of coordinates of the video. % and px
     * formats are supported: If the string ends in %, the `XPos` of the watermark will be the specified percentage of the
     * video width; for example, `10%` means that `XPos` is 10% of the video width; If the string ends in px, the `XPos` of the
     * watermark will be the specified px; for example, `100px` means that `XPos` is 100 px.Default value: 0 px.
     */
    xPos?: pulumi.Input<string>;
    /**
     * The vertical position of the origin of the watermark relative to the origin of coordinates of the video. % and px
     * formats are supported: If the string ends in %, the `YPos` of the watermark will be the specified percentage of the
     * video height; for example, `10%` means that `YPos` is 10% of the video height; If the string ends in px, the `YPos` of
     * the watermark will be the specified px; for example, `100px` means that `YPos` is 100 px.Default value: 0 px.
     */
    yPos?: pulumi.Input<string>;
}
