// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CloudResourceAttachmentArgs, CloudResourceAttachmentState } from "./cloudResourceAttachment";
export type CloudResourceAttachment = import("./cloudResourceAttachment").CloudResourceAttachment;
export const CloudResourceAttachment: typeof import("./cloudResourceAttachment").CloudResourceAttachment = null as any;
utilities.lazyLoad(exports, ["CloudResourceAttachment"], () => require("./cloudResourceAttachment"));

export { ExternalKeyArgs, ExternalKeyState } from "./externalKey";
export type ExternalKey = import("./externalKey").ExternalKey;
export const ExternalKey: typeof import("./externalKey").ExternalKey = null as any;
utilities.lazyLoad(exports, ["ExternalKey"], () => require("./externalKey"));

export { GetDescribeKeysArgs, GetDescribeKeysResult, GetDescribeKeysOutputArgs } from "./getDescribeKeys";
export const getDescribeKeys: typeof import("./getDescribeKeys").getDescribeKeys = null as any;
export const getDescribeKeysOutput: typeof import("./getDescribeKeys").getDescribeKeysOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeKeys","getDescribeKeysOutput"], () => require("./getDescribeKeys"));

export { GetGetParametersForImportArgs, GetGetParametersForImportResult, GetGetParametersForImportOutputArgs } from "./getGetParametersForImport";
export const getGetParametersForImport: typeof import("./getGetParametersForImport").getGetParametersForImport = null as any;
export const getGetParametersForImportOutput: typeof import("./getGetParametersForImport").getGetParametersForImportOutput = null as any;
utilities.lazyLoad(exports, ["getGetParametersForImport","getGetParametersForImportOutput"], () => require("./getGetParametersForImport"));

export { GetKeysArgs, GetKeysResult, GetKeysOutputArgs } from "./getKeys";
export const getKeys: typeof import("./getKeys").getKeys = null as any;
export const getKeysOutput: typeof import("./getKeys").getKeysOutput = null as any;
utilities.lazyLoad(exports, ["getKeys","getKeysOutput"], () => require("./getKeys"));

export { GetListAlgorithmsArgs, GetListAlgorithmsResult, GetListAlgorithmsOutputArgs } from "./getListAlgorithms";
export const getListAlgorithms: typeof import("./getListAlgorithms").getListAlgorithms = null as any;
export const getListAlgorithmsOutput: typeof import("./getListAlgorithms").getListAlgorithmsOutput = null as any;
utilities.lazyLoad(exports, ["getListAlgorithms","getListAlgorithmsOutput"], () => require("./getListAlgorithms"));

export { GetListKeysArgs, GetListKeysResult, GetListKeysOutputArgs } from "./getListKeys";
export const getListKeys: typeof import("./getListKeys").getListKeys = null as any;
export const getListKeysOutput: typeof import("./getListKeys").getListKeysOutput = null as any;
utilities.lazyLoad(exports, ["getListKeys","getListKeysOutput"], () => require("./getListKeys"));

export { GetPublicKeyArgs, GetPublicKeyResult, GetPublicKeyOutputArgs } from "./getPublicKey";
export const getPublicKey: typeof import("./getPublicKey").getPublicKey = null as any;
export const getPublicKeyOutput: typeof import("./getPublicKey").getPublicKeyOutput = null as any;
utilities.lazyLoad(exports, ["getPublicKey","getPublicKeyOutput"], () => require("./getPublicKey"));

export { GetWhiteBoxDecryptKeyArgs, GetWhiteBoxDecryptKeyResult, GetWhiteBoxDecryptKeyOutputArgs } from "./getWhiteBoxDecryptKey";
export const getWhiteBoxDecryptKey: typeof import("./getWhiteBoxDecryptKey").getWhiteBoxDecryptKey = null as any;
export const getWhiteBoxDecryptKeyOutput: typeof import("./getWhiteBoxDecryptKey").getWhiteBoxDecryptKeyOutput = null as any;
utilities.lazyLoad(exports, ["getWhiteBoxDecryptKey","getWhiteBoxDecryptKeyOutput"], () => require("./getWhiteBoxDecryptKey"));

export { GetWhiteBoxDeviceFingerprintsArgs, GetWhiteBoxDeviceFingerprintsResult, GetWhiteBoxDeviceFingerprintsOutputArgs } from "./getWhiteBoxDeviceFingerprints";
export const getWhiteBoxDeviceFingerprints: typeof import("./getWhiteBoxDeviceFingerprints").getWhiteBoxDeviceFingerprints = null as any;
export const getWhiteBoxDeviceFingerprintsOutput: typeof import("./getWhiteBoxDeviceFingerprints").getWhiteBoxDeviceFingerprintsOutput = null as any;
utilities.lazyLoad(exports, ["getWhiteBoxDeviceFingerprints","getWhiteBoxDeviceFingerprintsOutput"], () => require("./getWhiteBoxDeviceFingerprints"));

export { GetWhiteBoxKeyDetailsArgs, GetWhiteBoxKeyDetailsResult, GetWhiteBoxKeyDetailsOutputArgs } from "./getWhiteBoxKeyDetails";
export const getWhiteBoxKeyDetails: typeof import("./getWhiteBoxKeyDetails").getWhiteBoxKeyDetails = null as any;
export const getWhiteBoxKeyDetailsOutput: typeof import("./getWhiteBoxKeyDetails").getWhiteBoxKeyDetailsOutput = null as any;
utilities.lazyLoad(exports, ["getWhiteBoxKeyDetails","getWhiteBoxKeyDetailsOutput"], () => require("./getWhiteBoxKeyDetails"));

export { KeyArgs, KeyState } from "./key";
export type Key = import("./key").Key;
export const Key: typeof import("./key").Key = null as any;
utilities.lazyLoad(exports, ["Key"], () => require("./key"));

export { OverwriteWhiteBoxDeviceFingerprintsArgs, OverwriteWhiteBoxDeviceFingerprintsState } from "./overwriteWhiteBoxDeviceFingerprints";
export type OverwriteWhiteBoxDeviceFingerprints = import("./overwriteWhiteBoxDeviceFingerprints").OverwriteWhiteBoxDeviceFingerprints;
export const OverwriteWhiteBoxDeviceFingerprints: typeof import("./overwriteWhiteBoxDeviceFingerprints").OverwriteWhiteBoxDeviceFingerprints = null as any;
utilities.lazyLoad(exports, ["OverwriteWhiteBoxDeviceFingerprints"], () => require("./overwriteWhiteBoxDeviceFingerprints"));

export { WhiteBoxKeyArgs, WhiteBoxKeyState } from "./whiteBoxKey";
export type WhiteBoxKey = import("./whiteBoxKey").WhiteBoxKey;
export const WhiteBoxKey: typeof import("./whiteBoxKey").WhiteBoxKey = null as any;
utilities.lazyLoad(exports, ["WhiteBoxKey"], () => require("./whiteBoxKey"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Kms/cloudResourceAttachment:CloudResourceAttachment":
                return new CloudResourceAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Kms/externalKey:ExternalKey":
                return new ExternalKey(name, <any>undefined, { urn })
            case "tencentcloud:Kms/key:Key":
                return new Key(name, <any>undefined, { urn })
            case "tencentcloud:Kms/overwriteWhiteBoxDeviceFingerprints:OverwriteWhiteBoxDeviceFingerprints":
                return new OverwriteWhiteBoxDeviceFingerprints(name, <any>undefined, { urn })
            case "tencentcloud:Kms/whiteBoxKey:WhiteBoxKey":
                return new WhiteBoxKey(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Kms/cloudResourceAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Kms/externalKey", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Kms/key", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Kms/overwriteWhiteBoxDeviceFingerprints", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Kms/whiteBoxKey", _module)