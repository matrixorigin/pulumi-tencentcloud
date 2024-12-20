// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Redis
{
    [TencentcloudResourceType("tencentcloud:Redis/upgradeProxyVersionOperation:UpgradeProxyVersionOperation")]
    public partial class UpgradeProxyVersionOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Current proxy version.
        /// </summary>
        [Output("currentProxyVersion")]
        public Output<string> CurrentProxyVersion { get; private set; } = null!;

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Switch mode:1 - Upgrade now0 - Maintenance window upgrade.
        /// </summary>
        [Output("instanceTypeUpgradeNow")]
        public Output<int> InstanceTypeUpgradeNow { get; private set; } = null!;

        /// <summary>
        /// Upgradeable redis proxy version.
        /// </summary>
        [Output("upgradeProxyVersion")]
        public Output<string> UpgradeProxyVersion { get; private set; } = null!;


        /// <summary>
        /// Create a UpgradeProxyVersionOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UpgradeProxyVersionOperation(string name, UpgradeProxyVersionOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/upgradeProxyVersionOperation:UpgradeProxyVersionOperation", name, args ?? new UpgradeProxyVersionOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UpgradeProxyVersionOperation(string name, Input<string> id, UpgradeProxyVersionOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/upgradeProxyVersionOperation:UpgradeProxyVersionOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UpgradeProxyVersionOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UpgradeProxyVersionOperation Get(string name, Input<string> id, UpgradeProxyVersionOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new UpgradeProxyVersionOperation(name, id, state, options);
        }
    }

    public sealed class UpgradeProxyVersionOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Current proxy version.
        /// </summary>
        [Input("currentProxyVersion", required: true)]
        public Input<string> CurrentProxyVersion { get; set; } = null!;

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Switch mode:1 - Upgrade now0 - Maintenance window upgrade.
        /// </summary>
        [Input("instanceTypeUpgradeNow", required: true)]
        public Input<int> InstanceTypeUpgradeNow { get; set; } = null!;

        /// <summary>
        /// Upgradeable redis proxy version.
        /// </summary>
        [Input("upgradeProxyVersion", required: true)]
        public Input<string> UpgradeProxyVersion { get; set; } = null!;

        public UpgradeProxyVersionOperationArgs()
        {
        }
        public static new UpgradeProxyVersionOperationArgs Empty => new UpgradeProxyVersionOperationArgs();
    }

    public sealed class UpgradeProxyVersionOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Current proxy version.
        /// </summary>
        [Input("currentProxyVersion")]
        public Input<string>? CurrentProxyVersion { get; set; }

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Switch mode:1 - Upgrade now0 - Maintenance window upgrade.
        /// </summary>
        [Input("instanceTypeUpgradeNow")]
        public Input<int>? InstanceTypeUpgradeNow { get; set; }

        /// <summary>
        /// Upgradeable redis proxy version.
        /// </summary>
        [Input("upgradeProxyVersion")]
        public Input<string>? UpgradeProxyVersion { get; set; }

        public UpgradeProxyVersionOperationState()
        {
        }
        public static new UpgradeProxyVersionOperationState Empty => new UpgradeProxyVersionOperationState();
    }
}
