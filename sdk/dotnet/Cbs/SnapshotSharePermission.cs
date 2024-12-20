// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cbs
{
    [TencentcloudResourceType("tencentcloud:Cbs/snapshotSharePermission:SnapshotSharePermission")]
    public partial class SnapshotSharePermission : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of account IDs with which a snapshot is shared. For the format of array-type parameters, see[API
        /// Introduction](https://cloud.tencent.com/document/api/213/568). You can find the account ID in[Account
        /// Information](https://console.cloud.tencent.com/developer).
        /// </summary>
        [Output("accountIds")]
        public Output<ImmutableArray<string>> AccountIds { get; private set; } = null!;

        /// <summary>
        /// The ID of the snapshot to be queried. You can obtain this by using
        /// [DescribeSnapshots](https://cloud.tencent.com/document/api/362/15647).
        /// </summary>
        [Output("snapshotId")]
        public Output<string> SnapshotId { get; private set; } = null!;


        /// <summary>
        /// Create a SnapshotSharePermission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SnapshotSharePermission(string name, SnapshotSharePermissionArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cbs/snapshotSharePermission:SnapshotSharePermission", name, args ?? new SnapshotSharePermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SnapshotSharePermission(string name, Input<string> id, SnapshotSharePermissionState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cbs/snapshotSharePermission:SnapshotSharePermission", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SnapshotSharePermission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SnapshotSharePermission Get(string name, Input<string> id, SnapshotSharePermissionState? state = null, CustomResourceOptions? options = null)
        {
            return new SnapshotSharePermission(name, id, state, options);
        }
    }

    public sealed class SnapshotSharePermissionArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountIds", required: true)]
        private InputList<string>? _accountIds;

        /// <summary>
        /// List of account IDs with which a snapshot is shared. For the format of array-type parameters, see[API
        /// Introduction](https://cloud.tencent.com/document/api/213/568). You can find the account ID in[Account
        /// Information](https://console.cloud.tencent.com/developer).
        /// </summary>
        public InputList<string> AccountIds
        {
            get => _accountIds ?? (_accountIds = new InputList<string>());
            set => _accountIds = value;
        }

        /// <summary>
        /// The ID of the snapshot to be queried. You can obtain this by using
        /// [DescribeSnapshots](https://cloud.tencent.com/document/api/362/15647).
        /// </summary>
        [Input("snapshotId", required: true)]
        public Input<string> SnapshotId { get; set; } = null!;

        public SnapshotSharePermissionArgs()
        {
        }
        public static new SnapshotSharePermissionArgs Empty => new SnapshotSharePermissionArgs();
    }

    public sealed class SnapshotSharePermissionState : global::Pulumi.ResourceArgs
    {
        [Input("accountIds")]
        private InputList<string>? _accountIds;

        /// <summary>
        /// List of account IDs with which a snapshot is shared. For the format of array-type parameters, see[API
        /// Introduction](https://cloud.tencent.com/document/api/213/568). You can find the account ID in[Account
        /// Information](https://console.cloud.tencent.com/developer).
        /// </summary>
        public InputList<string> AccountIds
        {
            get => _accountIds ?? (_accountIds = new InputList<string>());
            set => _accountIds = value;
        }

        /// <summary>
        /// The ID of the snapshot to be queried. You can obtain this by using
        /// [DescribeSnapshots](https://cloud.tencent.com/document/api/362/15647).
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public SnapshotSharePermissionState()
        {
        }
        public static new SnapshotSharePermissionState Empty => new SnapshotSharePermissionState();
    }
}
