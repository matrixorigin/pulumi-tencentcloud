// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Identity
{
    [TencentcloudResourceType("tencentcloud:Identity/centerRoleConfigurationPermissionCustomPolicyAttachment:CenterRoleConfigurationPermissionCustomPolicyAttachment")]
    public partial class CenterRoleConfigurationPermissionCustomPolicyAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Role policy add time.
        /// </summary>
        [Output("addTime")]
        public Output<string> AddTime { get; private set; } = null!;

        /// <summary>
        /// Permission configuration ID.
        /// </summary>
        [Output("roleConfigurationId")]
        public Output<string> RoleConfigurationId { get; private set; } = null!;

        /// <summary>
        /// Role policy document.
        /// </summary>
        [Output("rolePolicyDocument")]
        public Output<string> RolePolicyDocument { get; private set; } = null!;

        /// <summary>
        /// Role policy name.
        /// </summary>
        [Output("rolePolicyName")]
        public Output<string> RolePolicyName { get; private set; } = null!;

        /// <summary>
        /// Role policy type.
        /// </summary>
        [Output("rolePolicyType")]
        public Output<string> RolePolicyType { get; private set; } = null!;

        /// <summary>
        /// Space ID.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a CenterRoleConfigurationPermissionCustomPolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CenterRoleConfigurationPermissionCustomPolicyAttachment(string name, CenterRoleConfigurationPermissionCustomPolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Identity/centerRoleConfigurationPermissionCustomPolicyAttachment:CenterRoleConfigurationPermissionCustomPolicyAttachment", name, args ?? new CenterRoleConfigurationPermissionCustomPolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CenterRoleConfigurationPermissionCustomPolicyAttachment(string name, Input<string> id, CenterRoleConfigurationPermissionCustomPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Identity/centerRoleConfigurationPermissionCustomPolicyAttachment:CenterRoleConfigurationPermissionCustomPolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CenterRoleConfigurationPermissionCustomPolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CenterRoleConfigurationPermissionCustomPolicyAttachment Get(string name, Input<string> id, CenterRoleConfigurationPermissionCustomPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new CenterRoleConfigurationPermissionCustomPolicyAttachment(name, id, state, options);
        }
    }

    public sealed class CenterRoleConfigurationPermissionCustomPolicyAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Permission configuration ID.
        /// </summary>
        [Input("roleConfigurationId", required: true)]
        public Input<string> RoleConfigurationId { get; set; } = null!;

        /// <summary>
        /// Role policy document.
        /// </summary>
        [Input("rolePolicyDocument", required: true)]
        public Input<string> RolePolicyDocument { get; set; } = null!;

        /// <summary>
        /// Role policy name.
        /// </summary>
        [Input("rolePolicyName", required: true)]
        public Input<string> RolePolicyName { get; set; } = null!;

        /// <summary>
        /// Space ID.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public CenterRoleConfigurationPermissionCustomPolicyAttachmentArgs()
        {
        }
        public static new CenterRoleConfigurationPermissionCustomPolicyAttachmentArgs Empty => new CenterRoleConfigurationPermissionCustomPolicyAttachmentArgs();
    }

    public sealed class CenterRoleConfigurationPermissionCustomPolicyAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Role policy add time.
        /// </summary>
        [Input("addTime")]
        public Input<string>? AddTime { get; set; }

        /// <summary>
        /// Permission configuration ID.
        /// </summary>
        [Input("roleConfigurationId")]
        public Input<string>? RoleConfigurationId { get; set; }

        /// <summary>
        /// Role policy document.
        /// </summary>
        [Input("rolePolicyDocument")]
        public Input<string>? RolePolicyDocument { get; set; }

        /// <summary>
        /// Role policy name.
        /// </summary>
        [Input("rolePolicyName")]
        public Input<string>? RolePolicyName { get; set; }

        /// <summary>
        /// Role policy type.
        /// </summary>
        [Input("rolePolicyType")]
        public Input<string>? RolePolicyType { get; set; }

        /// <summary>
        /// Space ID.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public CenterRoleConfigurationPermissionCustomPolicyAttachmentState()
        {
        }
        public static new CenterRoleConfigurationPermissionCustomPolicyAttachmentState Empty => new CenterRoleConfigurationPermissionCustomPolicyAttachmentState();
    }
}