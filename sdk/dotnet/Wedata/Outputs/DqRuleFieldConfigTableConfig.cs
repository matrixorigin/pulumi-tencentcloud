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
    public sealed class DqRuleFieldConfigTableConfig
    {
        public readonly string? DatabaseId;
        public readonly string? DatabaseName;
        public readonly ImmutableArray<Outputs.DqRuleFieldConfigTableConfigFieldConfig> FieldConfigs;
        public readonly string? TableId;
        public readonly string? TableKey;
        public readonly string? TableName;

        [OutputConstructor]
        private DqRuleFieldConfigTableConfig(
            string? databaseId,

            string? databaseName,

            ImmutableArray<Outputs.DqRuleFieldConfigTableConfigFieldConfig> fieldConfigs,

            string? tableId,

            string? tableKey,

            string? tableName)
        {
            DatabaseId = databaseId;
            DatabaseName = databaseName;
            FieldConfigs = fieldConfigs;
            TableId = tableId;
            TableKey = tableKey;
            TableName = tableName;
        }
    }
}