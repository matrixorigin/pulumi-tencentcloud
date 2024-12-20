// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Sqlserver
{
    [TencentcloudResourceType("tencentcloud:Sqlserver/readonlyInstance:ReadonlyInstance")]
    public partial class ReadonlyInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
        /// </summary>
        [Output("autoVoucher")]
        public Output<int?> AutoVoucher { get; private set; } = null!;

        /// <summary>
        /// Availability zone.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// Pay type of the SQL Server instance. Available values `PREPAID`, `POSTPAID_BY_HOUR`.
        /// </summary>
        [Output("chargeType")]
        public Output<string?> ChargeType { get; private set; } = null!;

        /// <summary>
        /// Create time of the SQL Server instance.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Indicate that the master instance upgrade or not. `true` for upgrading the master SQL Server instance to cluster type by
        /// force. Default is false. Note: this is not supported with `DUAL`(ha_type), `2017`(engine_version) master SQL Server
        /// instance, for it will cause ha_type of the master SQL Server instance change.
        /// </summary>
        [Output("forceUpgrade")]
        public Output<bool?> ForceUpgrade { get; private set; } = null!;

        /// <summary>
        /// Indicates the master instance ID of recovery instances.
        /// </summary>
        [Output("masterInstanceId")]
        public Output<string> MasterInstanceId { get; private set; } = null!;

        /// <summary>
        /// Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloud_sqlserver_specinfos`
        /// provides.
        /// </summary>
        [Output("memory")]
        public Output<int> Memory { get; private set; } = null!;

        /// <summary>
        /// Name of the SQL Server instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Purchase instance period in month. The value does not exceed 48.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// ID of the readonly group that this instance belongs to. When `readonly_group_type` set value `3`, it must be set with
        /// valid value.
        /// </summary>
        [Output("readonlyGroupId")]
        public Output<string> ReadonlyGroupId { get; private set; } = null!;

        /// <summary>
        /// Required when `readonly_group_type`=2, the name of the newly created read-only group.
        /// </summary>
        [Output("readonlyGroupName")]
        public Output<string> ReadonlyGroupName { get; private set; } = null!;

        /// <summary>
        /// Type of readonly group. Valid values: `1`, `3`. `1` for one auto-assigned readonly instance per one readonly group, `2`
        /// for creating new readonly group, `3` for all exist readonly instances stay in the exist readonly group. For now, only
        /// `1` and `3` are supported.
        /// </summary>
        [Output("readonlyGroupType")]
        public Output<int> ReadonlyGroupType { get; private set; } = null!;

        /// <summary>
        /// Required when `readonly_group_type`=2, whether the newly created read-only group has delay elimination enabled,
        /// 1-enabled, 0-disabled. When the delay between the read-only copy and the primary instance exceeds the threshold, it is
        /// automatically removed.
        /// </summary>
        [Output("readonlyGroupsIsOfflineDelay")]
        public Output<int> ReadonlyGroupsIsOfflineDelay { get; private set; } = null!;

        /// <summary>
        /// Required when `readonly_group_type`=2 and `readonly_groups_is_offline_delay`=1, the threshold for delayed elimination of
        /// newly created read-only groups.
        /// </summary>
        [Output("readonlyGroupsMaxDelayTime")]
        public Output<int> ReadonlyGroupsMaxDelayTime { get; private set; } = null!;

        /// <summary>
        /// When `readonly_group_type`=2 and `readonly_groups_is_offline_delay`=1, it is required. After the newly created read-only
        /// group is delayed and removed, at least the number of read-only copies should be retained.
        /// </summary>
        [Output("readonlyGroupsMinInGroup")]
        public Output<int> ReadonlyGroupsMinInGroup { get; private set; } = null!;

        /// <summary>
        /// Readonly flag. `RO` (read-only instance), `MASTER` (primary instance with read-only instances). If it is left empty, it
        /// refers to an instance which is not read-only and has no RO group.
        /// </summary>
        [Output("roFlag")]
        public Output<string> RoFlag { get; private set; } = null!;

        /// <summary>
        /// Security group bound to the instance.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// Status of the SQL Server instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for
        /// recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly,
        /// 12 for rebooting.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storage_min` and
        /// `storage_max` which data source `tencentcloud_sqlserver_specinfos` provides.
        /// </summary>
        [Output("storage")]
        public Output<int> Storage { get; private set; } = null!;

        /// <summary>
        /// ID of subnet.
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        /// <summary>
        /// The tags of the SQL Server.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// IP for private access.
        /// </summary>
        [Output("vip")]
        public Output<string> Vip { get; private set; } = null!;

        /// <summary>
        /// An array of voucher IDs, currently only one can be used for a single order.
        /// </summary>
        [Output("voucherIds")]
        public Output<ImmutableArray<string>> VoucherIds { get; private set; } = null!;

        /// <summary>
        /// ID of VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;

        /// <summary>
        /// Port for private access.
        /// </summary>
        [Output("vport")]
        public Output<int> Vport { get; private set; } = null!;

        /// <summary>
        /// The way to execute the allocation. Supported values include: 0 - execute immediately, 1 - execute in maintenance window.
        /// </summary>
        [Output("waitSwitch")]
        public Output<int?> WaitSwitch { get; private set; } = null!;


        /// <summary>
        /// Create a ReadonlyInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReadonlyInstance(string name, ReadonlyInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/readonlyInstance:ReadonlyInstance", name, args ?? new ReadonlyInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReadonlyInstance(string name, Input<string> id, ReadonlyInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/readonlyInstance:ReadonlyInstance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/matrixorigin",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReadonlyInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReadonlyInstance Get(string name, Input<string> id, ReadonlyInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new ReadonlyInstance(name, id, state, options);
        }
    }

    public sealed class ReadonlyInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
        /// </summary>
        [Input("autoVoucher")]
        public Input<int>? AutoVoucher { get; set; }

        /// <summary>
        /// Availability zone.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Pay type of the SQL Server instance. Available values `PREPAID`, `POSTPAID_BY_HOUR`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Indicate that the master instance upgrade or not. `true` for upgrading the master SQL Server instance to cluster type by
        /// force. Default is false. Note: this is not supported with `DUAL`(ha_type), `2017`(engine_version) master SQL Server
        /// instance, for it will cause ha_type of the master SQL Server instance change.
        /// </summary>
        [Input("forceUpgrade")]
        public Input<bool>? ForceUpgrade { get; set; }

        /// <summary>
        /// Indicates the master instance ID of recovery instances.
        /// </summary>
        [Input("masterInstanceId", required: true)]
        public Input<string> MasterInstanceId { get; set; } = null!;

        /// <summary>
        /// Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloud_sqlserver_specinfos`
        /// provides.
        /// </summary>
        [Input("memory", required: true)]
        public Input<int> Memory { get; set; } = null!;

        /// <summary>
        /// Name of the SQL Server instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Purchase instance period in month. The value does not exceed 48.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// ID of the readonly group that this instance belongs to. When `readonly_group_type` set value `3`, it must be set with
        /// valid value.
        /// </summary>
        [Input("readonlyGroupId")]
        public Input<string>? ReadonlyGroupId { get; set; }

        /// <summary>
        /// Required when `readonly_group_type`=2, the name of the newly created read-only group.
        /// </summary>
        [Input("readonlyGroupName")]
        public Input<string>? ReadonlyGroupName { get; set; }

        /// <summary>
        /// Type of readonly group. Valid values: `1`, `3`. `1` for one auto-assigned readonly instance per one readonly group, `2`
        /// for creating new readonly group, `3` for all exist readonly instances stay in the exist readonly group. For now, only
        /// `1` and `3` are supported.
        /// </summary>
        [Input("readonlyGroupType", required: true)]
        public Input<int> ReadonlyGroupType { get; set; } = null!;

        /// <summary>
        /// Required when `readonly_group_type`=2, whether the newly created read-only group has delay elimination enabled,
        /// 1-enabled, 0-disabled. When the delay between the read-only copy and the primary instance exceeds the threshold, it is
        /// automatically removed.
        /// </summary>
        [Input("readonlyGroupsIsOfflineDelay")]
        public Input<int>? ReadonlyGroupsIsOfflineDelay { get; set; }

        /// <summary>
        /// Required when `readonly_group_type`=2 and `readonly_groups_is_offline_delay`=1, the threshold for delayed elimination of
        /// newly created read-only groups.
        /// </summary>
        [Input("readonlyGroupsMaxDelayTime")]
        public Input<int>? ReadonlyGroupsMaxDelayTime { get; set; }

        /// <summary>
        /// When `readonly_group_type`=2 and `readonly_groups_is_offline_delay`=1, it is required. After the newly created read-only
        /// group is delayed and removed, at least the number of read-only copies should be retained.
        /// </summary>
        [Input("readonlyGroupsMinInGroup")]
        public Input<int>? ReadonlyGroupsMinInGroup { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// Security group bound to the instance.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storage_min` and
        /// `storage_max` which data source `tencentcloud_sqlserver_specinfos` provides.
        /// </summary>
        [Input("storage", required: true)]
        public Input<int> Storage { get; set; } = null!;

        /// <summary>
        /// ID of subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of the SQL Server.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("voucherIds")]
        private InputList<string>? _voucherIds;

        /// <summary>
        /// An array of voucher IDs, currently only one can be used for a single order.
        /// </summary>
        public InputList<string> VoucherIds
        {
            get => _voucherIds ?? (_voucherIds = new InputList<string>());
            set => _voucherIds = value;
        }

        /// <summary>
        /// ID of VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The way to execute the allocation. Supported values include: 0 - execute immediately, 1 - execute in maintenance window.
        /// </summary>
        [Input("waitSwitch")]
        public Input<int>? WaitSwitch { get; set; }

        public ReadonlyInstanceArgs()
        {
        }
        public static new ReadonlyInstanceArgs Empty => new ReadonlyInstanceArgs();
    }

    public sealed class ReadonlyInstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
        /// </summary>
        [Input("autoVoucher")]
        public Input<int>? AutoVoucher { get; set; }

        /// <summary>
        /// Availability zone.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Pay type of the SQL Server instance. Available values `PREPAID`, `POSTPAID_BY_HOUR`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Create time of the SQL Server instance.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Indicate that the master instance upgrade or not. `true` for upgrading the master SQL Server instance to cluster type by
        /// force. Default is false. Note: this is not supported with `DUAL`(ha_type), `2017`(engine_version) master SQL Server
        /// instance, for it will cause ha_type of the master SQL Server instance change.
        /// </summary>
        [Input("forceUpgrade")]
        public Input<bool>? ForceUpgrade { get; set; }

        /// <summary>
        /// Indicates the master instance ID of recovery instances.
        /// </summary>
        [Input("masterInstanceId")]
        public Input<string>? MasterInstanceId { get; set; }

        /// <summary>
        /// Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloud_sqlserver_specinfos`
        /// provides.
        /// </summary>
        [Input("memory")]
        public Input<int>? Memory { get; set; }

        /// <summary>
        /// Name of the SQL Server instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Purchase instance period in month. The value does not exceed 48.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// ID of the readonly group that this instance belongs to. When `readonly_group_type` set value `3`, it must be set with
        /// valid value.
        /// </summary>
        [Input("readonlyGroupId")]
        public Input<string>? ReadonlyGroupId { get; set; }

        /// <summary>
        /// Required when `readonly_group_type`=2, the name of the newly created read-only group.
        /// </summary>
        [Input("readonlyGroupName")]
        public Input<string>? ReadonlyGroupName { get; set; }

        /// <summary>
        /// Type of readonly group. Valid values: `1`, `3`. `1` for one auto-assigned readonly instance per one readonly group, `2`
        /// for creating new readonly group, `3` for all exist readonly instances stay in the exist readonly group. For now, only
        /// `1` and `3` are supported.
        /// </summary>
        [Input("readonlyGroupType")]
        public Input<int>? ReadonlyGroupType { get; set; }

        /// <summary>
        /// Required when `readonly_group_type`=2, whether the newly created read-only group has delay elimination enabled,
        /// 1-enabled, 0-disabled. When the delay between the read-only copy and the primary instance exceeds the threshold, it is
        /// automatically removed.
        /// </summary>
        [Input("readonlyGroupsIsOfflineDelay")]
        public Input<int>? ReadonlyGroupsIsOfflineDelay { get; set; }

        /// <summary>
        /// Required when `readonly_group_type`=2 and `readonly_groups_is_offline_delay`=1, the threshold for delayed elimination of
        /// newly created read-only groups.
        /// </summary>
        [Input("readonlyGroupsMaxDelayTime")]
        public Input<int>? ReadonlyGroupsMaxDelayTime { get; set; }

        /// <summary>
        /// When `readonly_group_type`=2 and `readonly_groups_is_offline_delay`=1, it is required. After the newly created read-only
        /// group is delayed and removed, at least the number of read-only copies should be retained.
        /// </summary>
        [Input("readonlyGroupsMinInGroup")]
        public Input<int>? ReadonlyGroupsMinInGroup { get; set; }

        /// <summary>
        /// Readonly flag. `RO` (read-only instance), `MASTER` (primary instance with read-only instances). If it is left empty, it
        /// refers to an instance which is not read-only and has no RO group.
        /// </summary>
        [Input("roFlag")]
        public Input<string>? RoFlag { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// Security group bound to the instance.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Status of the SQL Server instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for
        /// recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly,
        /// 12 for rebooting.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storage_min` and
        /// `storage_max` which data source `tencentcloud_sqlserver_specinfos` provides.
        /// </summary>
        [Input("storage")]
        public Input<int>? Storage { get; set; }

        /// <summary>
        /// ID of subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of the SQL Server.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// IP for private access.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        [Input("voucherIds")]
        private InputList<string>? _voucherIds;

        /// <summary>
        /// An array of voucher IDs, currently only one can be used for a single order.
        /// </summary>
        public InputList<string> VoucherIds
        {
            get => _voucherIds ?? (_voucherIds = new InputList<string>());
            set => _voucherIds = value;
        }

        /// <summary>
        /// ID of VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Port for private access.
        /// </summary>
        [Input("vport")]
        public Input<int>? Vport { get; set; }

        /// <summary>
        /// The way to execute the allocation. Supported values include: 0 - execute immediately, 1 - execute in maintenance window.
        /// </summary>
        [Input("waitSwitch")]
        public Input<int>? WaitSwitch { get; set; }

        public ReadonlyInstanceState()
        {
        }
        public static new ReadonlyInstanceState Empty => new ReadonlyInstanceState();
    }
}
