// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu
{
    [TencentcloudResourceType("tencentcloud:Dayu/ddosPolicyAttachment:DdosPolicyAttachment")]
    public partial class DdosPolicyAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the policy.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// ID of the attached resource.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// Type of the resource that the DDoS policy works for. Valid values are `bgpip`, `bgp`, `bgp-multip`, `net`.
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;


        /// <summary>
        /// Create a DdosPolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DdosPolicyAttachment(string name, DdosPolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dayu/ddosPolicyAttachment:DdosPolicyAttachment", name, args ?? new DdosPolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DdosPolicyAttachment(string name, Input<string> id, DdosPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dayu/ddosPolicyAttachment:DdosPolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DdosPolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DdosPolicyAttachment Get(string name, Input<string> id, DdosPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new DdosPolicyAttachment(name, id, state, options);
        }
    }

    public sealed class DdosPolicyAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the policy.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        /// <summary>
        /// ID of the attached resource.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// Type of the resource that the DDoS policy works for. Valid values are `bgpip`, `bgp`, `bgp-multip`, `net`.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        public DdosPolicyAttachmentArgs()
        {
        }
        public static new DdosPolicyAttachmentArgs Empty => new DdosPolicyAttachmentArgs();
    }

    public sealed class DdosPolicyAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the policy.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// ID of the attached resource.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// Type of the resource that the DDoS policy works for. Valid values are `bgpip`, `bgp`, `bgp-multip`, `net`.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        public DdosPolicyAttachmentState()
        {
        }
        public static new DdosPolicyAttachmentState Empty => new DdosPolicyAttachmentState();
    }
}