// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CertificateArgs, CertificateState } from "./certificate";
export type Certificate = import("./certificate").Certificate;
export const Certificate: typeof import("./certificate").Certificate = null as any;
utilities.lazyLoad(exports, ["Certificate"], () => require("./certificate"));

export { CheckCertificateChainOperationArgs, CheckCertificateChainOperationState } from "./checkCertificateChainOperation";
export type CheckCertificateChainOperation = import("./checkCertificateChainOperation").CheckCertificateChainOperation;
export const CheckCertificateChainOperation: typeof import("./checkCertificateChainOperation").CheckCertificateChainOperation = null as any;
utilities.lazyLoad(exports, ["CheckCertificateChainOperation"], () => require("./checkCertificateChainOperation"));

export { CompleteCertificateOperationArgs, CompleteCertificateOperationState } from "./completeCertificateOperation";
export type CompleteCertificateOperation = import("./completeCertificateOperation").CompleteCertificateOperation;
export const CompleteCertificateOperation: typeof import("./completeCertificateOperation").CompleteCertificateOperation = null as any;
utilities.lazyLoad(exports, ["CompleteCertificateOperation"], () => require("./completeCertificateOperation"));

export { DeployCertificateInstanceOperationArgs, DeployCertificateInstanceOperationState } from "./deployCertificateInstanceOperation";
export type DeployCertificateInstanceOperation = import("./deployCertificateInstanceOperation").DeployCertificateInstanceOperation;
export const DeployCertificateInstanceOperation: typeof import("./deployCertificateInstanceOperation").DeployCertificateInstanceOperation = null as any;
utilities.lazyLoad(exports, ["DeployCertificateInstanceOperation"], () => require("./deployCertificateInstanceOperation"));

export { DeployCertificateRecordRetryOperationArgs, DeployCertificateRecordRetryOperationState } from "./deployCertificateRecordRetryOperation";
export type DeployCertificateRecordRetryOperation = import("./deployCertificateRecordRetryOperation").DeployCertificateRecordRetryOperation;
export const DeployCertificateRecordRetryOperation: typeof import("./deployCertificateRecordRetryOperation").DeployCertificateRecordRetryOperation = null as any;
utilities.lazyLoad(exports, ["DeployCertificateRecordRetryOperation"], () => require("./deployCertificateRecordRetryOperation"));

export { DeployCertificateRecordRollbackOperationArgs, DeployCertificateRecordRollbackOperationState } from "./deployCertificateRecordRollbackOperation";
export type DeployCertificateRecordRollbackOperation = import("./deployCertificateRecordRollbackOperation").DeployCertificateRecordRollbackOperation;
export const DeployCertificateRecordRollbackOperation: typeof import("./deployCertificateRecordRollbackOperation").DeployCertificateRecordRollbackOperation = null as any;
utilities.lazyLoad(exports, ["DeployCertificateRecordRollbackOperation"], () => require("./deployCertificateRecordRollbackOperation"));

export { DownloadCertificateOperationArgs, DownloadCertificateOperationState } from "./downloadCertificateOperation";
export type DownloadCertificateOperation = import("./downloadCertificateOperation").DownloadCertificateOperation;
export const DownloadCertificateOperation: typeof import("./downloadCertificateOperation").DownloadCertificateOperation = null as any;
utilities.lazyLoad(exports, ["DownloadCertificateOperation"], () => require("./downloadCertificateOperation"));

export { FreeCertificateArgs, FreeCertificateState } from "./freeCertificate";
export type FreeCertificate = import("./freeCertificate").FreeCertificate;
export const FreeCertificate: typeof import("./freeCertificate").FreeCertificate = null as any;
utilities.lazyLoad(exports, ["FreeCertificate"], () => require("./freeCertificate"));

export { GetCertificatesArgs, GetCertificatesResult, GetCertificatesOutputArgs } from "./getCertificates";
export const getCertificates: typeof import("./getCertificates").getCertificates = null as any;
export const getCertificatesOutput: typeof import("./getCertificates").getCertificatesOutput = null as any;
utilities.lazyLoad(exports, ["getCertificates","getCertificatesOutput"], () => require("./getCertificates"));

export { GetDescribeCertificateArgs, GetDescribeCertificateResult, GetDescribeCertificateOutputArgs } from "./getDescribeCertificate";
export const getDescribeCertificate: typeof import("./getDescribeCertificate").getDescribeCertificate = null as any;
export const getDescribeCertificateOutput: typeof import("./getDescribeCertificate").getDescribeCertificateOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeCertificate","getDescribeCertificateOutput"], () => require("./getDescribeCertificate"));

export { GetDescribeCompaniesArgs, GetDescribeCompaniesResult, GetDescribeCompaniesOutputArgs } from "./getDescribeCompanies";
export const getDescribeCompanies: typeof import("./getDescribeCompanies").getDescribeCompanies = null as any;
export const getDescribeCompaniesOutput: typeof import("./getDescribeCompanies").getDescribeCompaniesOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeCompanies","getDescribeCompaniesOutput"], () => require("./getDescribeCompanies"));

export { GetDescribeHostApiGatewayInstanceListArgs, GetDescribeHostApiGatewayInstanceListResult, GetDescribeHostApiGatewayInstanceListOutputArgs } from "./getDescribeHostApiGatewayInstanceList";
export const getDescribeHostApiGatewayInstanceList: typeof import("./getDescribeHostApiGatewayInstanceList").getDescribeHostApiGatewayInstanceList = null as any;
export const getDescribeHostApiGatewayInstanceListOutput: typeof import("./getDescribeHostApiGatewayInstanceList").getDescribeHostApiGatewayInstanceListOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostApiGatewayInstanceList","getDescribeHostApiGatewayInstanceListOutput"], () => require("./getDescribeHostApiGatewayInstanceList"));

export { GetDescribeHostCdnInstanceListArgs, GetDescribeHostCdnInstanceListResult, GetDescribeHostCdnInstanceListOutputArgs } from "./getDescribeHostCdnInstanceList";
export const getDescribeHostCdnInstanceList: typeof import("./getDescribeHostCdnInstanceList").getDescribeHostCdnInstanceList = null as any;
export const getDescribeHostCdnInstanceListOutput: typeof import("./getDescribeHostCdnInstanceList").getDescribeHostCdnInstanceListOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostCdnInstanceList","getDescribeHostCdnInstanceListOutput"], () => require("./getDescribeHostCdnInstanceList"));

export { GetDescribeHostClbInstanceListArgs, GetDescribeHostClbInstanceListResult, GetDescribeHostClbInstanceListOutputArgs } from "./getDescribeHostClbInstanceList";
export const getDescribeHostClbInstanceList: typeof import("./getDescribeHostClbInstanceList").getDescribeHostClbInstanceList = null as any;
export const getDescribeHostClbInstanceListOutput: typeof import("./getDescribeHostClbInstanceList").getDescribeHostClbInstanceListOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostClbInstanceList","getDescribeHostClbInstanceListOutput"], () => require("./getDescribeHostClbInstanceList"));

export { GetDescribeHostCosInstanceListArgs, GetDescribeHostCosInstanceListResult, GetDescribeHostCosInstanceListOutputArgs } from "./getDescribeHostCosInstanceList";
export const getDescribeHostCosInstanceList: typeof import("./getDescribeHostCosInstanceList").getDescribeHostCosInstanceList = null as any;
export const getDescribeHostCosInstanceListOutput: typeof import("./getDescribeHostCosInstanceList").getDescribeHostCosInstanceListOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostCosInstanceList","getDescribeHostCosInstanceListOutput"], () => require("./getDescribeHostCosInstanceList"));

export { GetDescribeHostDdosInstanceListArgs, GetDescribeHostDdosInstanceListResult, GetDescribeHostDdosInstanceListOutputArgs } from "./getDescribeHostDdosInstanceList";
export const getDescribeHostDdosInstanceList: typeof import("./getDescribeHostDdosInstanceList").getDescribeHostDdosInstanceList = null as any;
export const getDescribeHostDdosInstanceListOutput: typeof import("./getDescribeHostDdosInstanceList").getDescribeHostDdosInstanceListOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostDdosInstanceList","getDescribeHostDdosInstanceListOutput"], () => require("./getDescribeHostDdosInstanceList"));

export { GetDescribeHostDeployRecordArgs, GetDescribeHostDeployRecordResult, GetDescribeHostDeployRecordOutputArgs } from "./getDescribeHostDeployRecord";
export const getDescribeHostDeployRecord: typeof import("./getDescribeHostDeployRecord").getDescribeHostDeployRecord = null as any;
export const getDescribeHostDeployRecordOutput: typeof import("./getDescribeHostDeployRecord").getDescribeHostDeployRecordOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostDeployRecord","getDescribeHostDeployRecordOutput"], () => require("./getDescribeHostDeployRecord"));

export { GetDescribeHostDeployRecordDetailArgs, GetDescribeHostDeployRecordDetailResult, GetDescribeHostDeployRecordDetailOutputArgs } from "./getDescribeHostDeployRecordDetail";
export const getDescribeHostDeployRecordDetail: typeof import("./getDescribeHostDeployRecordDetail").getDescribeHostDeployRecordDetail = null as any;
export const getDescribeHostDeployRecordDetailOutput: typeof import("./getDescribeHostDeployRecordDetail").getDescribeHostDeployRecordDetailOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostDeployRecordDetail","getDescribeHostDeployRecordDetailOutput"], () => require("./getDescribeHostDeployRecordDetail"));

export { GetDescribeHostLighthouseInstanceListArgs, GetDescribeHostLighthouseInstanceListResult, GetDescribeHostLighthouseInstanceListOutputArgs } from "./getDescribeHostLighthouseInstanceList";
export const getDescribeHostLighthouseInstanceList: typeof import("./getDescribeHostLighthouseInstanceList").getDescribeHostLighthouseInstanceList = null as any;
export const getDescribeHostLighthouseInstanceListOutput: typeof import("./getDescribeHostLighthouseInstanceList").getDescribeHostLighthouseInstanceListOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostLighthouseInstanceList","getDescribeHostLighthouseInstanceListOutput"], () => require("./getDescribeHostLighthouseInstanceList"));

export { GetDescribeHostLiveInstanceListArgs, GetDescribeHostLiveInstanceListResult, GetDescribeHostLiveInstanceListOutputArgs } from "./getDescribeHostLiveInstanceList";
export const getDescribeHostLiveInstanceList: typeof import("./getDescribeHostLiveInstanceList").getDescribeHostLiveInstanceList = null as any;
export const getDescribeHostLiveInstanceListOutput: typeof import("./getDescribeHostLiveInstanceList").getDescribeHostLiveInstanceListOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostLiveInstanceList","getDescribeHostLiveInstanceListOutput"], () => require("./getDescribeHostLiveInstanceList"));

export { GetDescribeHostTeoInstanceListArgs, GetDescribeHostTeoInstanceListResult, GetDescribeHostTeoInstanceListOutputArgs } from "./getDescribeHostTeoInstanceList";
export const getDescribeHostTeoInstanceList: typeof import("./getDescribeHostTeoInstanceList").getDescribeHostTeoInstanceList = null as any;
export const getDescribeHostTeoInstanceListOutput: typeof import("./getDescribeHostTeoInstanceList").getDescribeHostTeoInstanceListOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostTeoInstanceList","getDescribeHostTeoInstanceListOutput"], () => require("./getDescribeHostTeoInstanceList"));

export { GetDescribeHostTkeInstanceListArgs, GetDescribeHostTkeInstanceListResult, GetDescribeHostTkeInstanceListOutputArgs } from "./getDescribeHostTkeInstanceList";
export const getDescribeHostTkeInstanceList: typeof import("./getDescribeHostTkeInstanceList").getDescribeHostTkeInstanceList = null as any;
export const getDescribeHostTkeInstanceListOutput: typeof import("./getDescribeHostTkeInstanceList").getDescribeHostTkeInstanceListOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostTkeInstanceList","getDescribeHostTkeInstanceListOutput"], () => require("./getDescribeHostTkeInstanceList"));

export { GetDescribeHostUpdateRecordArgs, GetDescribeHostUpdateRecordResult, GetDescribeHostUpdateRecordOutputArgs } from "./getDescribeHostUpdateRecord";
export const getDescribeHostUpdateRecord: typeof import("./getDescribeHostUpdateRecord").getDescribeHostUpdateRecord = null as any;
export const getDescribeHostUpdateRecordOutput: typeof import("./getDescribeHostUpdateRecord").getDescribeHostUpdateRecordOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostUpdateRecord","getDescribeHostUpdateRecordOutput"], () => require("./getDescribeHostUpdateRecord"));

export { GetDescribeHostUpdateRecordDetailArgs, GetDescribeHostUpdateRecordDetailResult, GetDescribeHostUpdateRecordDetailOutputArgs } from "./getDescribeHostUpdateRecordDetail";
export const getDescribeHostUpdateRecordDetail: typeof import("./getDescribeHostUpdateRecordDetail").getDescribeHostUpdateRecordDetail = null as any;
export const getDescribeHostUpdateRecordDetailOutput: typeof import("./getDescribeHostUpdateRecordDetail").getDescribeHostUpdateRecordDetailOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostUpdateRecordDetail","getDescribeHostUpdateRecordDetailOutput"], () => require("./getDescribeHostUpdateRecordDetail"));

export { GetDescribeHostVodInstanceListArgs, GetDescribeHostVodInstanceListResult, GetDescribeHostVodInstanceListOutputArgs } from "./getDescribeHostVodInstanceList";
export const getDescribeHostVodInstanceList: typeof import("./getDescribeHostVodInstanceList").getDescribeHostVodInstanceList = null as any;
export const getDescribeHostVodInstanceListOutput: typeof import("./getDescribeHostVodInstanceList").getDescribeHostVodInstanceListOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostVodInstanceList","getDescribeHostVodInstanceListOutput"], () => require("./getDescribeHostVodInstanceList"));

export { GetDescribeHostWafInstanceListArgs, GetDescribeHostWafInstanceListResult, GetDescribeHostWafInstanceListOutputArgs } from "./getDescribeHostWafInstanceList";
export const getDescribeHostWafInstanceList: typeof import("./getDescribeHostWafInstanceList").getDescribeHostWafInstanceList = null as any;
export const getDescribeHostWafInstanceListOutput: typeof import("./getDescribeHostWafInstanceList").getDescribeHostWafInstanceListOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeHostWafInstanceList","getDescribeHostWafInstanceListOutput"], () => require("./getDescribeHostWafInstanceList"));

export { GetDescribeManagerDetailArgs, GetDescribeManagerDetailResult, GetDescribeManagerDetailOutputArgs } from "./getDescribeManagerDetail";
export const getDescribeManagerDetail: typeof import("./getDescribeManagerDetail").getDescribeManagerDetail = null as any;
export const getDescribeManagerDetailOutput: typeof import("./getDescribeManagerDetail").getDescribeManagerDetailOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeManagerDetail","getDescribeManagerDetailOutput"], () => require("./getDescribeManagerDetail"));

export { GetDescribeManagersArgs, GetDescribeManagersResult, GetDescribeManagersOutputArgs } from "./getDescribeManagers";
export const getDescribeManagers: typeof import("./getDescribeManagers").getDescribeManagers = null as any;
export const getDescribeManagersOutput: typeof import("./getDescribeManagers").getDescribeManagersOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeManagers","getDescribeManagersOutput"], () => require("./getDescribeManagers"));

export { PayCertificateArgs, PayCertificateState } from "./payCertificate";
export type PayCertificate = import("./payCertificate").PayCertificate;
export const PayCertificate: typeof import("./payCertificate").PayCertificate = null as any;
utilities.lazyLoad(exports, ["PayCertificate"], () => require("./payCertificate"));

export { ReplaceCertificateOperationArgs, ReplaceCertificateOperationState } from "./replaceCertificateOperation";
export type ReplaceCertificateOperation = import("./replaceCertificateOperation").ReplaceCertificateOperation;
export const ReplaceCertificateOperation: typeof import("./replaceCertificateOperation").ReplaceCertificateOperation = null as any;
utilities.lazyLoad(exports, ["ReplaceCertificateOperation"], () => require("./replaceCertificateOperation"));

export { RevokeCertificateOperationArgs, RevokeCertificateOperationState } from "./revokeCertificateOperation";
export type RevokeCertificateOperation = import("./revokeCertificateOperation").RevokeCertificateOperation;
export const RevokeCertificateOperation: typeof import("./revokeCertificateOperation").RevokeCertificateOperation = null as any;
utilities.lazyLoad(exports, ["RevokeCertificateOperation"], () => require("./revokeCertificateOperation"));

export { UpdateCertificateInstanceOperationArgs, UpdateCertificateInstanceOperationState } from "./updateCertificateInstanceOperation";
export type UpdateCertificateInstanceOperation = import("./updateCertificateInstanceOperation").UpdateCertificateInstanceOperation;
export const UpdateCertificateInstanceOperation: typeof import("./updateCertificateInstanceOperation").UpdateCertificateInstanceOperation = null as any;
utilities.lazyLoad(exports, ["UpdateCertificateInstanceOperation"], () => require("./updateCertificateInstanceOperation"));

export { UpdateCertificateRecordRetryOperationArgs, UpdateCertificateRecordRetryOperationState } from "./updateCertificateRecordRetryOperation";
export type UpdateCertificateRecordRetryOperation = import("./updateCertificateRecordRetryOperation").UpdateCertificateRecordRetryOperation;
export const UpdateCertificateRecordRetryOperation: typeof import("./updateCertificateRecordRetryOperation").UpdateCertificateRecordRetryOperation = null as any;
utilities.lazyLoad(exports, ["UpdateCertificateRecordRetryOperation"], () => require("./updateCertificateRecordRetryOperation"));

export { UpdateCertificateRecordRollbackOperationArgs, UpdateCertificateRecordRollbackOperationState } from "./updateCertificateRecordRollbackOperation";
export type UpdateCertificateRecordRollbackOperation = import("./updateCertificateRecordRollbackOperation").UpdateCertificateRecordRollbackOperation;
export const UpdateCertificateRecordRollbackOperation: typeof import("./updateCertificateRecordRollbackOperation").UpdateCertificateRecordRollbackOperation = null as any;
utilities.lazyLoad(exports, ["UpdateCertificateRecordRollbackOperation"], () => require("./updateCertificateRecordRollbackOperation"));

export { UploadRevokeLetterOperationArgs, UploadRevokeLetterOperationState } from "./uploadRevokeLetterOperation";
export type UploadRevokeLetterOperation = import("./uploadRevokeLetterOperation").UploadRevokeLetterOperation;
export const UploadRevokeLetterOperation: typeof import("./uploadRevokeLetterOperation").UploadRevokeLetterOperation = null as any;
utilities.lazyLoad(exports, ["UploadRevokeLetterOperation"], () => require("./uploadRevokeLetterOperation"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Ssl/certificate:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/checkCertificateChainOperation:CheckCertificateChainOperation":
                return new CheckCertificateChainOperation(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/completeCertificateOperation:CompleteCertificateOperation":
                return new CompleteCertificateOperation(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/deployCertificateInstanceOperation:DeployCertificateInstanceOperation":
                return new DeployCertificateInstanceOperation(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/deployCertificateRecordRetryOperation:DeployCertificateRecordRetryOperation":
                return new DeployCertificateRecordRetryOperation(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/deployCertificateRecordRollbackOperation:DeployCertificateRecordRollbackOperation":
                return new DeployCertificateRecordRollbackOperation(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/downloadCertificateOperation:DownloadCertificateOperation":
                return new DownloadCertificateOperation(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/freeCertificate:FreeCertificate":
                return new FreeCertificate(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/payCertificate:PayCertificate":
                return new PayCertificate(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/replaceCertificateOperation:ReplaceCertificateOperation":
                return new ReplaceCertificateOperation(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/revokeCertificateOperation:RevokeCertificateOperation":
                return new RevokeCertificateOperation(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/updateCertificateInstanceOperation:UpdateCertificateInstanceOperation":
                return new UpdateCertificateInstanceOperation(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/updateCertificateRecordRetryOperation:UpdateCertificateRecordRetryOperation":
                return new UpdateCertificateRecordRetryOperation(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/updateCertificateRecordRollbackOperation:UpdateCertificateRecordRollbackOperation":
                return new UpdateCertificateRecordRollbackOperation(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/uploadRevokeLetterOperation:UploadRevokeLetterOperation":
                return new UploadRevokeLetterOperation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/certificate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/checkCertificateChainOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/completeCertificateOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/deployCertificateInstanceOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/deployCertificateRecordRetryOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/deployCertificateRecordRollbackOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/downloadCertificateOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/freeCertificate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/payCertificate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/replaceCertificateOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/revokeCertificateOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/updateCertificateInstanceOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/updateCertificateRecordRetryOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/updateCertificateRecordRollbackOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/uploadRevokeLetterOperation", _module)