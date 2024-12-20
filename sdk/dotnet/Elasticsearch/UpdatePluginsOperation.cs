// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elasticsearch
{
    [TencentcloudResourceType("tencentcloud:Elasticsearch/updatePluginsOperation:UpdatePluginsOperation")]
    public partial class UpdatePluginsOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to force a restart. Default is false.
        /// </summary>
        [Output("forceRestart")]
        public Output<bool?> ForceRestart { get; private set; } = null!;

        /// <summary>
        /// Whether to reinstall, default value false.
        /// </summary>
        [Output("forceUpdate")]
        public Output<bool?> ForceUpdate { get; private set; } = null!;

        /// <summary>
        /// List of plugins that need to be installed.
        /// </summary>
        [Output("installPluginLists")]
        public Output<ImmutableArray<string>> InstallPluginLists { get; private set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Plugin type. 0: system plugin.
        /// </summary>
        [Output("pluginType")]
        public Output<int?> PluginType { get; private set; } = null!;

        /// <summary>
        /// List of plugins that need to be uninstalled.
        /// </summary>
        [Output("removePluginLists")]
        public Output<ImmutableArray<string>> RemovePluginLists { get; private set; } = null!;


        /// <summary>
        /// Create a UpdatePluginsOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UpdatePluginsOperation(string name, UpdatePluginsOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Elasticsearch/updatePluginsOperation:UpdatePluginsOperation", name, args ?? new UpdatePluginsOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UpdatePluginsOperation(string name, Input<string> id, UpdatePluginsOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Elasticsearch/updatePluginsOperation:UpdatePluginsOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UpdatePluginsOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UpdatePluginsOperation Get(string name, Input<string> id, UpdatePluginsOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new UpdatePluginsOperation(name, id, state, options);
        }
    }

    public sealed class UpdatePluginsOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to force a restart. Default is false.
        /// </summary>
        [Input("forceRestart")]
        public Input<bool>? ForceRestart { get; set; }

        /// <summary>
        /// Whether to reinstall, default value false.
        /// </summary>
        [Input("forceUpdate")]
        public Input<bool>? ForceUpdate { get; set; }

        [Input("installPluginLists")]
        private InputList<string>? _installPluginLists;

        /// <summary>
        /// List of plugins that need to be installed.
        /// </summary>
        public InputList<string> InstallPluginLists
        {
            get => _installPluginLists ?? (_installPluginLists = new InputList<string>());
            set => _installPluginLists = value;
        }

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Plugin type. 0: system plugin.
        /// </summary>
        [Input("pluginType")]
        public Input<int>? PluginType { get; set; }

        [Input("removePluginLists")]
        private InputList<string>? _removePluginLists;

        /// <summary>
        /// List of plugins that need to be uninstalled.
        /// </summary>
        public InputList<string> RemovePluginLists
        {
            get => _removePluginLists ?? (_removePluginLists = new InputList<string>());
            set => _removePluginLists = value;
        }

        public UpdatePluginsOperationArgs()
        {
        }
        public static new UpdatePluginsOperationArgs Empty => new UpdatePluginsOperationArgs();
    }

    public sealed class UpdatePluginsOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to force a restart. Default is false.
        /// </summary>
        [Input("forceRestart")]
        public Input<bool>? ForceRestart { get; set; }

        /// <summary>
        /// Whether to reinstall, default value false.
        /// </summary>
        [Input("forceUpdate")]
        public Input<bool>? ForceUpdate { get; set; }

        [Input("installPluginLists")]
        private InputList<string>? _installPluginLists;

        /// <summary>
        /// List of plugins that need to be installed.
        /// </summary>
        public InputList<string> InstallPluginLists
        {
            get => _installPluginLists ?? (_installPluginLists = new InputList<string>());
            set => _installPluginLists = value;
        }

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Plugin type. 0: system plugin.
        /// </summary>
        [Input("pluginType")]
        public Input<int>? PluginType { get; set; }

        [Input("removePluginLists")]
        private InputList<string>? _removePluginLists;

        /// <summary>
        /// List of plugins that need to be uninstalled.
        /// </summary>
        public InputList<string> RemovePluginLists
        {
            get => _removePluginLists ?? (_removePluginLists = new InputList<string>());
            set => _removePluginLists = value;
        }

        public UpdatePluginsOperationState()
        {
        }
        public static new UpdatePluginsOperationState Empty => new UpdatePluginsOperationState();
    }
}
