// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Wedata.Outputs
{

    [OutputType]
    public sealed class GetRuleTemplatesDataResult
    {
        public readonly int CitationCount;
        public readonly int CompareType;
        public readonly string Description;
        public readonly bool MultiSourceFlag;
        public readonly string Name;
        public readonly int QualityDim;
        public readonly int RuleTemplateId;
        public readonly string SourceContent;
        public readonly ImmutableArray<int> SourceEngineTypes;
        public readonly int SourceObjectDataType;
        public readonly int SourceObjectType;
        public readonly string SqlExpression;
        public readonly int SubQualityDim;
        public readonly int Type;
        public readonly string UpdateTime;
        public readonly int UserId;
        public readonly string UserName;
        public readonly bool WhereFlag;

        [OutputConstructor]
        private GetRuleTemplatesDataResult(
            int citationCount,

            int compareType,

            string description,

            bool multiSourceFlag,

            string name,

            int qualityDim,

            int ruleTemplateId,

            string sourceContent,

            ImmutableArray<int> sourceEngineTypes,

            int sourceObjectDataType,

            int sourceObjectType,

            string sqlExpression,

            int subQualityDim,

            int type,

            string updateTime,

            int userId,

            string userName,

            bool whereFlag)
        {
            CitationCount = citationCount;
            CompareType = compareType;
            Description = description;
            MultiSourceFlag = multiSourceFlag;
            Name = name;
            QualityDim = qualityDim;
            RuleTemplateId = ruleTemplateId;
            SourceContent = sourceContent;
            SourceEngineTypes = sourceEngineTypes;
            SourceObjectDataType = sourceObjectDataType;
            SourceObjectType = sourceObjectType;
            SqlExpression = sqlExpression;
            SubQualityDim = subQualityDim;
            Type = type;
            UpdateTime = updateTime;
            UserId = userId;
            UserName = userName;
            WhereFlag = whereFlag;
        }
    }
}