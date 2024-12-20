// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Postgresql
{
    [TencentcloudResourceType("tencentcloud:Postgresql/baseBackup:BaseBackup")]
    public partial class BaseBackup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Base backup ID.
        /// </summary>
        [Output("baseBackupId")]
        public Output<string> BaseBackupId { get; private set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("dbInstanceId")]
        public Output<string> DbInstanceId { get; private set; } = null!;

        /// <summary>
        /// New expiration time.
        /// </summary>
        [Output("newExpireTime")]
        public Output<string> NewExpireTime { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a BaseBackup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BaseBackup(string name, BaseBackupArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/baseBackup:BaseBackup", name, args ?? new BaseBackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BaseBackup(string name, Input<string> id, BaseBackupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/baseBackup:BaseBackup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BaseBackup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BaseBackup Get(string name, Input<string> id, BaseBackupState? state = null, CustomResourceOptions? options = null)
        {
            return new BaseBackup(name, id, state, options);
        }
    }

    public sealed class BaseBackupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        /// <summary>
        /// New expiration time.
        /// </summary>
        [Input("newExpireTime")]
        public Input<string>? NewExpireTime { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public BaseBackupArgs()
        {
        }
        public static new BaseBackupArgs Empty => new BaseBackupArgs();
    }

    public sealed class BaseBackupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Base backup ID.
        /// </summary>
        [Input("baseBackupId")]
        public Input<string>? BaseBackupId { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        /// <summary>
        /// New expiration time.
        /// </summary>
        [Input("newExpireTime")]
        public Input<string>? NewExpireTime { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public BaseBackupState()
        {
        }
        public static new BaseBackupState Empty => new BaseBackupState();
    }
}
