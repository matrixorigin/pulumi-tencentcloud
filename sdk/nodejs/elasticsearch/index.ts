// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DiagnoseArgs, DiagnoseState } from "./diagnose";
export type Diagnose = import("./diagnose").Diagnose;
export const Diagnose: typeof import("./diagnose").Diagnose = null as any;
utilities.lazyLoad(exports, ["Diagnose"], () => require("./diagnose"));

export { DiagnoseInstanceArgs, DiagnoseInstanceState } from "./diagnoseInstance";
export type DiagnoseInstance = import("./diagnoseInstance").DiagnoseInstance;
export const DiagnoseInstance: typeof import("./diagnoseInstance").DiagnoseInstance = null as any;
utilities.lazyLoad(exports, ["DiagnoseInstance"], () => require("./diagnoseInstance"));

export { GetDescribeIndexListArgs, GetDescribeIndexListResult, GetDescribeIndexListOutputArgs } from "./getDescribeIndexList";
export const getDescribeIndexList: typeof import("./getDescribeIndexList").getDescribeIndexList = null as any;
export const getDescribeIndexListOutput: typeof import("./getDescribeIndexList").getDescribeIndexListOutput = null as any;
utilities.lazyLoad(exports, ["getDescribeIndexList","getDescribeIndexListOutput"], () => require("./getDescribeIndexList"));

export { GetDiagnoseArgs, GetDiagnoseResult, GetDiagnoseOutputArgs } from "./getDiagnose";
export const getDiagnose: typeof import("./getDiagnose").getDiagnose = null as any;
export const getDiagnoseOutput: typeof import("./getDiagnose").getDiagnoseOutput = null as any;
utilities.lazyLoad(exports, ["getDiagnose","getDiagnoseOutput"], () => require("./getDiagnose"));

export { GetInstanceLogsArgs, GetInstanceLogsResult, GetInstanceLogsOutputArgs } from "./getInstanceLogs";
export const getInstanceLogs: typeof import("./getInstanceLogs").getInstanceLogs = null as any;
export const getInstanceLogsOutput: typeof import("./getInstanceLogs").getInstanceLogsOutput = null as any;
utilities.lazyLoad(exports, ["getInstanceLogs","getInstanceLogsOutput"], () => require("./getInstanceLogs"));

export { GetInstanceOperationsArgs, GetInstanceOperationsResult, GetInstanceOperationsOutputArgs } from "./getInstanceOperations";
export const getInstanceOperations: typeof import("./getInstanceOperations").getInstanceOperations = null as any;
export const getInstanceOperationsOutput: typeof import("./getInstanceOperations").getInstanceOperationsOutput = null as any;
utilities.lazyLoad(exports, ["getInstanceOperations","getInstanceOperationsOutput"], () => require("./getInstanceOperations"));

export { GetInstancePluginListArgs, GetInstancePluginListResult, GetInstancePluginListOutputArgs } from "./getInstancePluginList";
export const getInstancePluginList: typeof import("./getInstancePluginList").getInstancePluginList = null as any;
export const getInstancePluginListOutput: typeof import("./getInstancePluginList").getInstancePluginListOutput = null as any;
utilities.lazyLoad(exports, ["getInstancePluginList","getInstancePluginListOutput"], () => require("./getInstancePluginList"));

export { GetInstancesArgs, GetInstancesResult, GetInstancesOutputArgs } from "./getInstances";
export const getInstances: typeof import("./getInstances").getInstances = null as any;
export const getInstancesOutput: typeof import("./getInstances").getInstancesOutput = null as any;
utilities.lazyLoad(exports, ["getInstances","getInstancesOutput"], () => require("./getInstances"));

export { GetLogstashInstanceLogsArgs, GetLogstashInstanceLogsResult, GetLogstashInstanceLogsOutputArgs } from "./getLogstashInstanceLogs";
export const getLogstashInstanceLogs: typeof import("./getLogstashInstanceLogs").getLogstashInstanceLogs = null as any;
export const getLogstashInstanceLogsOutput: typeof import("./getLogstashInstanceLogs").getLogstashInstanceLogsOutput = null as any;
utilities.lazyLoad(exports, ["getLogstashInstanceLogs","getLogstashInstanceLogsOutput"], () => require("./getLogstashInstanceLogs"));

export { GetLogstashInstanceOperationsArgs, GetLogstashInstanceOperationsResult, GetLogstashInstanceOperationsOutputArgs } from "./getLogstashInstanceOperations";
export const getLogstashInstanceOperations: typeof import("./getLogstashInstanceOperations").getLogstashInstanceOperations = null as any;
export const getLogstashInstanceOperationsOutput: typeof import("./getLogstashInstanceOperations").getLogstashInstanceOperationsOutput = null as any;
utilities.lazyLoad(exports, ["getLogstashInstanceOperations","getLogstashInstanceOperationsOutput"], () => require("./getLogstashInstanceOperations"));

export { GetViewsArgs, GetViewsResult, GetViewsOutputArgs } from "./getViews";
export const getViews: typeof import("./getViews").getViews = null as any;
export const getViewsOutput: typeof import("./getViews").getViewsOutput = null as any;
utilities.lazyLoad(exports, ["getViews","getViewsOutput"], () => require("./getViews"));

export { IndexArgs, IndexState } from "./index_";
export type Index = import("./index_").Index;
export const Index: typeof import("./index_").Index = null as any;
utilities.lazyLoad(exports, ["Index"], () => require("./index_"));

export { InstanceArgs, InstanceState } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));

export { LogstashArgs, LogstashState } from "./logstash";
export type Logstash = import("./logstash").Logstash;
export const Logstash: typeof import("./logstash").Logstash = null as any;
utilities.lazyLoad(exports, ["Logstash"], () => require("./logstash"));

export { LogstashPipelineArgs, LogstashPipelineState } from "./logstashPipeline";
export type LogstashPipeline = import("./logstashPipeline").LogstashPipeline;
export const LogstashPipeline: typeof import("./logstashPipeline").LogstashPipeline = null as any;
utilities.lazyLoad(exports, ["LogstashPipeline"], () => require("./logstashPipeline"));

export { RestartInstanceOperationArgs, RestartInstanceOperationState } from "./restartInstanceOperation";
export type RestartInstanceOperation = import("./restartInstanceOperation").RestartInstanceOperation;
export const RestartInstanceOperation: typeof import("./restartInstanceOperation").RestartInstanceOperation = null as any;
utilities.lazyLoad(exports, ["RestartInstanceOperation"], () => require("./restartInstanceOperation"));

export { RestartKibanaOperationArgs, RestartKibanaOperationState } from "./restartKibanaOperation";
export type RestartKibanaOperation = import("./restartKibanaOperation").RestartKibanaOperation;
export const RestartKibanaOperation: typeof import("./restartKibanaOperation").RestartKibanaOperation = null as any;
utilities.lazyLoad(exports, ["RestartKibanaOperation"], () => require("./restartKibanaOperation"));

export { RestartLogstashInstanceOperationArgs, RestartLogstashInstanceOperationState } from "./restartLogstashInstanceOperation";
export type RestartLogstashInstanceOperation = import("./restartLogstashInstanceOperation").RestartLogstashInstanceOperation;
export const RestartLogstashInstanceOperation: typeof import("./restartLogstashInstanceOperation").RestartLogstashInstanceOperation = null as any;
utilities.lazyLoad(exports, ["RestartLogstashInstanceOperation"], () => require("./restartLogstashInstanceOperation"));

export { RestartNodesOperationArgs, RestartNodesOperationState } from "./restartNodesOperation";
export type RestartNodesOperation = import("./restartNodesOperation").RestartNodesOperation;
export const RestartNodesOperation: typeof import("./restartNodesOperation").RestartNodesOperation = null as any;
utilities.lazyLoad(exports, ["RestartNodesOperation"], () => require("./restartNodesOperation"));

export { SecurityGroupArgs, SecurityGroupState } from "./securityGroup";
export type SecurityGroup = import("./securityGroup").SecurityGroup;
export const SecurityGroup: typeof import("./securityGroup").SecurityGroup = null as any;
utilities.lazyLoad(exports, ["SecurityGroup"], () => require("./securityGroup"));

export { StartLogstashPipelineOperationArgs, StartLogstashPipelineOperationState } from "./startLogstashPipelineOperation";
export type StartLogstashPipelineOperation = import("./startLogstashPipelineOperation").StartLogstashPipelineOperation;
export const StartLogstashPipelineOperation: typeof import("./startLogstashPipelineOperation").StartLogstashPipelineOperation = null as any;
utilities.lazyLoad(exports, ["StartLogstashPipelineOperation"], () => require("./startLogstashPipelineOperation"));

export { StopLogstashPipelineOperationArgs, StopLogstashPipelineOperationState } from "./stopLogstashPipelineOperation";
export type StopLogstashPipelineOperation = import("./stopLogstashPipelineOperation").StopLogstashPipelineOperation;
export const StopLogstashPipelineOperation: typeof import("./stopLogstashPipelineOperation").StopLogstashPipelineOperation = null as any;
utilities.lazyLoad(exports, ["StopLogstashPipelineOperation"], () => require("./stopLogstashPipelineOperation"));

export { UpdatePluginsOperationArgs, UpdatePluginsOperationState } from "./updatePluginsOperation";
export type UpdatePluginsOperation = import("./updatePluginsOperation").UpdatePluginsOperation;
export const UpdatePluginsOperation: typeof import("./updatePluginsOperation").UpdatePluginsOperation = null as any;
utilities.lazyLoad(exports, ["UpdatePluginsOperation"], () => require("./updatePluginsOperation"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Elasticsearch/diagnose:Diagnose":
                return new Diagnose(name, <any>undefined, { urn })
            case "tencentcloud:Elasticsearch/diagnoseInstance:DiagnoseInstance":
                return new DiagnoseInstance(name, <any>undefined, { urn })
            case "tencentcloud:Elasticsearch/index:Index":
                return new Index(name, <any>undefined, { urn })
            case "tencentcloud:Elasticsearch/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "tencentcloud:Elasticsearch/logstash:Logstash":
                return new Logstash(name, <any>undefined, { urn })
            case "tencentcloud:Elasticsearch/logstashPipeline:LogstashPipeline":
                return new LogstashPipeline(name, <any>undefined, { urn })
            case "tencentcloud:Elasticsearch/restartInstanceOperation:RestartInstanceOperation":
                return new RestartInstanceOperation(name, <any>undefined, { urn })
            case "tencentcloud:Elasticsearch/restartKibanaOperation:RestartKibanaOperation":
                return new RestartKibanaOperation(name, <any>undefined, { urn })
            case "tencentcloud:Elasticsearch/restartLogstashInstanceOperation:RestartLogstashInstanceOperation":
                return new RestartLogstashInstanceOperation(name, <any>undefined, { urn })
            case "tencentcloud:Elasticsearch/restartNodesOperation:RestartNodesOperation":
                return new RestartNodesOperation(name, <any>undefined, { urn })
            case "tencentcloud:Elasticsearch/securityGroup:SecurityGroup":
                return new SecurityGroup(name, <any>undefined, { urn })
            case "tencentcloud:Elasticsearch/startLogstashPipelineOperation:StartLogstashPipelineOperation":
                return new StartLogstashPipelineOperation(name, <any>undefined, { urn })
            case "tencentcloud:Elasticsearch/stopLogstashPipelineOperation:StopLogstashPipelineOperation":
                return new StopLogstashPipelineOperation(name, <any>undefined, { urn })
            case "tencentcloud:Elasticsearch/updatePluginsOperation:UpdatePluginsOperation":
                return new UpdatePluginsOperation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/diagnose", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/diagnoseInstance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/index", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/instance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/logstash", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/logstashPipeline", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/restartInstanceOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/restartKibanaOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/restartLogstashInstanceOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/restartNodesOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/securityGroup", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/startLogstashPipelineOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/stopLogstashPipelineOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Elasticsearch/updatePluginsOperation", _module)
