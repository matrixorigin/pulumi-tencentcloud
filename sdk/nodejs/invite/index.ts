// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { OrganizationMemberOperationArgs, OrganizationMemberOperationState } from "./organizationMemberOperation";
export type OrganizationMemberOperation = import("./organizationMemberOperation").OrganizationMemberOperation;
export const OrganizationMemberOperation: typeof import("./organizationMemberOperation").OrganizationMemberOperation = null as any;
utilities.lazyLoad(exports, ["OrganizationMemberOperation"], () => require("./organizationMemberOperation"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Invite/organizationMemberOperation:OrganizationMemberOperation":
                return new OrganizationMemberOperation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Invite/organizationMemberOperation", _module)
