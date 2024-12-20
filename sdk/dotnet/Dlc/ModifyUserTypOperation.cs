// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dlc
{
    [TencentcloudResourceType("tencentcloud:Dlc/modifyUserTypOperation:ModifyUserTypOperation")]
    public partial class ModifyUserTypOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// User id (uin), if left blank, it defaults to the caller's sub-uin.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;

        /// <summary>
        /// User type, only support: ADMIN: ddministrator/COMMON: ordinary user.
        /// </summary>
        [Output("userType")]
        public Output<string> UserType { get; private set; } = null!;


        /// <summary>
        /// Create a ModifyUserTypOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ModifyUserTypOperation(string name, ModifyUserTypOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/modifyUserTypOperation:ModifyUserTypOperation", name, args ?? new ModifyUserTypOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ModifyUserTypOperation(string name, Input<string> id, ModifyUserTypOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/modifyUserTypOperation:ModifyUserTypOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ModifyUserTypOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ModifyUserTypOperation Get(string name, Input<string> id, ModifyUserTypOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new ModifyUserTypOperation(name, id, state, options);
        }
    }

    public sealed class ModifyUserTypOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User id (uin), if left blank, it defaults to the caller's sub-uin.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        /// <summary>
        /// User type, only support: ADMIN: ddministrator/COMMON: ordinary user.
        /// </summary>
        [Input("userType", required: true)]
        public Input<string> UserType { get; set; } = null!;

        public ModifyUserTypOperationArgs()
        {
        }
        public static new ModifyUserTypOperationArgs Empty => new ModifyUserTypOperationArgs();
    }

    public sealed class ModifyUserTypOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User id (uin), if left blank, it defaults to the caller's sub-uin.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// User type, only support: ADMIN: ddministrator/COMMON: ordinary user.
        /// </summary>
        [Input("userType")]
        public Input<string>? UserType { get; set; }

        public ModifyUserTypOperationState()
        {
        }
        public static new ModifyUserTypOperationState Empty => new ModifyUserTypOperationState();
    }
}
