// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ApplyAccountBaselinesArgs, ApplyAccountBaselinesState } from "./applyAccountBaselines";
export type ApplyAccountBaselines = import("./applyAccountBaselines").ApplyAccountBaselines;
export const ApplyAccountBaselines: typeof import("./applyAccountBaselines").ApplyAccountBaselines = null as any;
utilities.lazyLoad(exports, ["ApplyAccountBaselines"], () => require("./applyAccountBaselines"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Batch/applyAccountBaselines:ApplyAccountBaselines":
                return new ApplyAccountBaselines(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Batch/applyAccountBaselines", _module)