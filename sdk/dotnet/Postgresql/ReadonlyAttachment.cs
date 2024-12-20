// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Postgresql
{
    [TencentcloudResourceType("tencentcloud:Postgresql/readonlyAttachment:ReadonlyAttachment")]
    public partial class ReadonlyAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Read only instance ID.
        /// </summary>
        [Output("dbInstanceId")]
        public Output<string> DbInstanceId { get; private set; } = null!;

        /// <summary>
        /// Read only group ID.
        /// </summary>
        [Output("readOnlyGroupId")]
        public Output<string> ReadOnlyGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a ReadonlyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReadonlyAttachment(string name, ReadonlyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/readonlyAttachment:ReadonlyAttachment", name, args ?? new ReadonlyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReadonlyAttachment(string name, Input<string> id, ReadonlyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/readonlyAttachment:ReadonlyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReadonlyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReadonlyAttachment Get(string name, Input<string> id, ReadonlyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ReadonlyAttachment(name, id, state, options);
        }
    }

    public sealed class ReadonlyAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Read only instance ID.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        /// <summary>
        /// Read only group ID.
        /// </summary>
        [Input("readOnlyGroupId", required: true)]
        public Input<string> ReadOnlyGroupId { get; set; } = null!;

        public ReadonlyAttachmentArgs()
        {
        }
        public static new ReadonlyAttachmentArgs Empty => new ReadonlyAttachmentArgs();
    }

    public sealed class ReadonlyAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Read only instance ID.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        /// <summary>
        /// Read only group ID.
        /// </summary>
        [Input("readOnlyGroupId")]
        public Input<string>? ReadOnlyGroupId { get; set; }

        public ReadonlyAttachmentState()
        {
        }
        public static new ReadonlyAttachmentState Empty => new ReadonlyAttachmentState();
    }
}
