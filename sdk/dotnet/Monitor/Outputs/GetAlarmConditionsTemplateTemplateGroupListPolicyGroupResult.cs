// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class GetAlarmConditionsTemplateTemplateGroupListPolicyGroupResult
    {
        public readonly bool CanSetDefault;
        public readonly bool Enable;
        public readonly int GroupId;
        public readonly string GroupName;
        public readonly int InsertTime;
        public readonly int IsDefault;
        public readonly int IsUnionRule;
        public readonly int LastEditUin;
        public readonly int NoShieldedInstanceCount;
        public readonly int ParentGroupId;
        public readonly int ProjectId;
        public readonly ImmutableArray<Outputs.GetAlarmConditionsTemplateTemplateGroupListPolicyGroupReceiverInfoResult> ReceiverInfos;
        public readonly string Remark;
        public readonly int TotalInstanceCount;
        public readonly int UpdateTime;
        public readonly string ViewName;

        [OutputConstructor]
        private GetAlarmConditionsTemplateTemplateGroupListPolicyGroupResult(
            bool canSetDefault,

            bool enable,

            int groupId,

            string groupName,

            int insertTime,

            int isDefault,

            int isUnionRule,

            int lastEditUin,

            int noShieldedInstanceCount,

            int parentGroupId,

            int projectId,

            ImmutableArray<Outputs.GetAlarmConditionsTemplateTemplateGroupListPolicyGroupReceiverInfoResult> receiverInfos,

            string remark,

            int totalInstanceCount,

            int updateTime,

            string viewName)
        {
            CanSetDefault = canSetDefault;
            Enable = enable;
            GroupId = groupId;
            GroupName = groupName;
            InsertTime = insertTime;
            IsDefault = isDefault;
            IsUnionRule = isUnionRule;
            LastEditUin = lastEditUin;
            NoShieldedInstanceCount = noShieldedInstanceCount;
            ParentGroupId = parentGroupId;
            ProjectId = projectId;
            ReceiverInfos = receiverInfos;
            Remark = remark;
            TotalInstanceCount = totalInstanceCount;
            UpdateTime = updateTime;
            ViewName = viewName;
        }
    }
}