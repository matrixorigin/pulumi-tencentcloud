// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Postgresql.Inputs
{

    public sealed class AccountPrivilegesOperationModifyPrivilegeSetDatabasePrivilegeObjectGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("objectName", required: true)]
        public Input<string> ObjectName { get; set; } = null!;

        [Input("objectType", required: true)]
        public Input<string> ObjectType { get; set; } = null!;

        [Input("schemaName")]
        public Input<string>? SchemaName { get; set; }

        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        public AccountPrivilegesOperationModifyPrivilegeSetDatabasePrivilegeObjectGetArgs()
        {
        }
        public static new AccountPrivilegesOperationModifyPrivilegeSetDatabasePrivilegeObjectGetArgs Empty => new AccountPrivilegesOperationModifyPrivilegeSetDatabasePrivilegeObjectGetArgs();
    }
}