// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dnspod
{
    [TencentcloudResourceType("tencentcloud:Dnspod/downloadSnapshotOperation:DownloadSnapshotOperation")]
    public partial class DownloadSnapshotOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Snapshot download url.
        /// </summary>
        [Output("cosUrl")]
        public Output<string> CosUrl { get; private set; } = null!;

        /// <summary>
        /// Domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Snapshot ID.
        /// </summary>
        [Output("snapshotId")]
        public Output<string> SnapshotId { get; private set; } = null!;


        /// <summary>
        /// Create a DownloadSnapshotOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DownloadSnapshotOperation(string name, DownloadSnapshotOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dnspod/downloadSnapshotOperation:DownloadSnapshotOperation", name, args ?? new DownloadSnapshotOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DownloadSnapshotOperation(string name, Input<string> id, DownloadSnapshotOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dnspod/downloadSnapshotOperation:DownloadSnapshotOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DownloadSnapshotOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DownloadSnapshotOperation Get(string name, Input<string> id, DownloadSnapshotOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new DownloadSnapshotOperation(name, id, state, options);
        }
    }

    public sealed class DownloadSnapshotOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Snapshot ID.
        /// </summary>
        [Input("snapshotId", required: true)]
        public Input<string> SnapshotId { get; set; } = null!;

        public DownloadSnapshotOperationArgs()
        {
        }
        public static new DownloadSnapshotOperationArgs Empty => new DownloadSnapshotOperationArgs();
    }

    public sealed class DownloadSnapshotOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Snapshot download url.
        /// </summary>
        [Input("cosUrl")]
        public Input<string>? CosUrl { get; set; }

        /// <summary>
        /// Domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Snapshot ID.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public DownloadSnapshotOperationState()
        {
        }
        public static new DownloadSnapshotOperationState Empty => new DownloadSnapshotOperationState();
    }
}