// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetPublicIpv6sArgs, GetPublicIpv6sResult, GetPublicIpv6sOutputArgs } from "./getPublicIpv6s";
export const getPublicIpv6s: typeof import("./getPublicIpv6s").getPublicIpv6s = null as any;
export const getPublicIpv6sOutput: typeof import("./getPublicIpv6s").getPublicIpv6sOutput = null as any;
utilities.lazyLoad(exports, ["getPublicIpv6s","getPublicIpv6sOutput"], () => require("./getPublicIpv6s"));

export { PublicIpv6Args, PublicIpv6State } from "./publicIpv6";
export type PublicIpv6 = import("./publicIpv6").PublicIpv6;
export const PublicIpv6: typeof import("./publicIpv6").PublicIpv6 = null as any;
utilities.lazyLoad(exports, ["PublicIpv6"], () => require("./publicIpv6"));

export { PublicIpv6AttachmentArgs, PublicIpv6AttachmentState } from "./publicIpv6Attachment";
export type PublicIpv6Attachment = import("./publicIpv6Attachment").PublicIpv6Attachment;
export const PublicIpv6Attachment: typeof import("./publicIpv6Attachment").PublicIpv6Attachment = null as any;
utilities.lazyLoad(exports, ["PublicIpv6Attachment"], () => require("./publicIpv6Attachment"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Elastic/publicIpv6:PublicIpv6":
                return new PublicIpv6(name, <any>undefined, { urn })
            case "tencentcloud:Elastic/publicIpv6Attachment:PublicIpv6Attachment":
                return new PublicIpv6Attachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Elastic/publicIpv6", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elastic/publicIpv6Attachment", _module)