// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mongodb
{
    [TencentcloudResourceType("tencentcloud:Mongodb/instanceBackupDownloadTask:InstanceBackupDownloadTask")]
    public partial class InstanceBackupDownloadTask : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the backup file to be downloaded can be obtained through the DescribeDBBackups interface.
        /// </summary>
        [Output("backupName")]
        public Output<string> BackupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the node names of replica sets to download or a list of shard names for sharded clusters.For example, the
        /// replica set cmgo-p8vnipr5, example (fixed value): BackupSets.0=cmgo-p8vnipr5_0, the full amount of data can be
        /// downloaded.For example, the sharded cluster cmgo-p8vnipr5, for example:
        /// BackupSets.0=cmgo-p8vnipr5_0&amp;amp;amp;BackupSets.1=cmgo-p8vnipr5_1, that is, to download the data of shard 0 and 1. If
        /// the sharded cluster needs to be downloaded in full, please pass in the example. Full slice name.
        /// </summary>
        [Output("backupSets")]
        public Output<ImmutableArray<Outputs.InstanceBackupDownloadTaskBackupSet>> BackupSets { get; private set; } = null!;

        /// <summary>
        /// Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceBackupDownloadTask resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceBackupDownloadTask(string name, InstanceBackupDownloadTaskArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mongodb/instanceBackupDownloadTask:InstanceBackupDownloadTask", name, args ?? new InstanceBackupDownloadTaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceBackupDownloadTask(string name, Input<string> id, InstanceBackupDownloadTaskState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mongodb/instanceBackupDownloadTask:InstanceBackupDownloadTask", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceBackupDownloadTask resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceBackupDownloadTask Get(string name, Input<string> id, InstanceBackupDownloadTaskState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceBackupDownloadTask(name, id, state, options);
        }
    }

    public sealed class InstanceBackupDownloadTaskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the backup file to be downloaded can be obtained through the DescribeDBBackups interface.
        /// </summary>
        [Input("backupName", required: true)]
        public Input<string> BackupName { get; set; } = null!;

        [Input("backupSets", required: true)]
        private InputList<Inputs.InstanceBackupDownloadTaskBackupSetArgs>? _backupSets;

        /// <summary>
        /// Specifies the node names of replica sets to download or a list of shard names for sharded clusters.For example, the
        /// replica set cmgo-p8vnipr5, example (fixed value): BackupSets.0=cmgo-p8vnipr5_0, the full amount of data can be
        /// downloaded.For example, the sharded cluster cmgo-p8vnipr5, for example:
        /// BackupSets.0=cmgo-p8vnipr5_0&amp;amp;amp;BackupSets.1=cmgo-p8vnipr5_1, that is, to download the data of shard 0 and 1. If
        /// the sharded cluster needs to be downloaded in full, please pass in the example. Full slice name.
        /// </summary>
        public InputList<Inputs.InstanceBackupDownloadTaskBackupSetArgs> BackupSets
        {
            get => _backupSets ?? (_backupSets = new InputList<Inputs.InstanceBackupDownloadTaskBackupSetArgs>());
            set => _backupSets = value;
        }

        /// <summary>
        /// Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public InstanceBackupDownloadTaskArgs()
        {
        }
        public static new InstanceBackupDownloadTaskArgs Empty => new InstanceBackupDownloadTaskArgs();
    }

    public sealed class InstanceBackupDownloadTaskState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the backup file to be downloaded can be obtained through the DescribeDBBackups interface.
        /// </summary>
        [Input("backupName")]
        public Input<string>? BackupName { get; set; }

        [Input("backupSets")]
        private InputList<Inputs.InstanceBackupDownloadTaskBackupSetGetArgs>? _backupSets;

        /// <summary>
        /// Specifies the node names of replica sets to download or a list of shard names for sharded clusters.For example, the
        /// replica set cmgo-p8vnipr5, example (fixed value): BackupSets.0=cmgo-p8vnipr5_0, the full amount of data can be
        /// downloaded.For example, the sharded cluster cmgo-p8vnipr5, for example:
        /// BackupSets.0=cmgo-p8vnipr5_0&amp;amp;amp;BackupSets.1=cmgo-p8vnipr5_1, that is, to download the data of shard 0 and 1. If
        /// the sharded cluster needs to be downloaded in full, please pass in the example. Full slice name.
        /// </summary>
        public InputList<Inputs.InstanceBackupDownloadTaskBackupSetGetArgs> BackupSets
        {
            get => _backupSets ?? (_backupSets = new InputList<Inputs.InstanceBackupDownloadTaskBackupSetGetArgs>());
            set => _backupSets = value;
        }

        /// <summary>
        /// Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public InstanceBackupDownloadTaskState()
        {
        }
        public static new InstanceBackupDownloadTaskState Empty => new InstanceBackupDownloadTaskState();
    }
}
