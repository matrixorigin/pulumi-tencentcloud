// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Redis
{
    [TencentcloudResourceType("tencentcloud:Redis/replicateAttachment:ReplicateAttachment")]
    public partial class ReplicateAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of group.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// All instance ids of the replication group.
        /// </summary>
        [Output("instanceIds")]
        public Output<ImmutableArray<string>> InstanceIds { get; private set; } = null!;

        /// <summary>
        /// The ID of master instance.
        /// </summary>
        [Output("masterInstanceId")]
        public Output<string> MasterInstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicateAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicateAttachment(string name, ReplicateAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/replicateAttachment:ReplicateAttachment", name, args ?? new ReplicateAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplicateAttachment(string name, Input<string> id, ReplicateAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/replicateAttachment:ReplicateAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReplicateAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicateAttachment Get(string name, Input<string> id, ReplicateAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplicateAttachment(name, id, state, options);
        }
    }

    public sealed class ReplicateAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of group.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("instanceIds", required: true)]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// All instance ids of the replication group.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// The ID of master instance.
        /// </summary>
        [Input("masterInstanceId", required: true)]
        public Input<string> MasterInstanceId { get; set; } = null!;

        public ReplicateAttachmentArgs()
        {
        }
        public static new ReplicateAttachmentArgs Empty => new ReplicateAttachmentArgs();
    }

    public sealed class ReplicateAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// All instance ids of the replication group.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// The ID of master instance.
        /// </summary>
        [Input("masterInstanceId")]
        public Input<string>? MasterInstanceId { get; set; }

        public ReplicateAttachmentState()
        {
        }
        public static new ReplicateAttachmentState Empty => new ReplicateAttachmentState();
    }
}
