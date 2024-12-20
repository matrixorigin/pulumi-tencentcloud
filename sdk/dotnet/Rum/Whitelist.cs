// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Rum
{
    [TencentcloudResourceType("tencentcloud:Rum/whitelist:Whitelist")]
    public partial class Whitelist : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Business identifier.
        /// </summary>
        [Output("aid")]
        public Output<string?> Aid { get; private set; } = null!;

        /// <summary>
        /// Creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Creator ID.
        /// </summary>
        [Output("createUser")]
        public Output<string> CreateUser { get; private set; } = null!;

        /// <summary>
        /// Instance ID, such as taw-123.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Remarks.
        /// </summary>
        [Output("remark")]
        public Output<string> Remark { get; private set; } = null!;

        /// <summary>
        /// End time.
        /// </summary>
        [Output("ttl")]
        public Output<string> Ttl { get; private set; } = null!;

        /// <summary>
        /// uin: business identifier.
        /// </summary>
        [Output("whitelistUin")]
        public Output<string> WhitelistUin { get; private set; } = null!;

        /// <summary>
        /// Auto-Increment allowlist ID.
        /// </summary>
        [Output("wid")]
        public Output<string> Wid { get; private set; } = null!;


        /// <summary>
        /// Create a Whitelist resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Whitelist(string name, WhitelistArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Rum/whitelist:Whitelist", name, args ?? new WhitelistArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Whitelist(string name, Input<string> id, WhitelistState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Rum/whitelist:Whitelist", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Whitelist resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Whitelist Get(string name, Input<string> id, WhitelistState? state = null, CustomResourceOptions? options = null)
        {
            return new Whitelist(name, id, state, options);
        }
    }

    public sealed class WhitelistArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Business identifier.
        /// </summary>
        [Input("aid")]
        public Input<string>? Aid { get; set; }

        /// <summary>
        /// Instance ID, such as taw-123.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Remarks.
        /// </summary>
        [Input("remark", required: true)]
        public Input<string> Remark { get; set; } = null!;

        /// <summary>
        /// uin: business identifier.
        /// </summary>
        [Input("whitelistUin", required: true)]
        public Input<string> WhitelistUin { get; set; } = null!;

        public WhitelistArgs()
        {
        }
        public static new WhitelistArgs Empty => new WhitelistArgs();
    }

    public sealed class WhitelistState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Business identifier.
        /// </summary>
        [Input("aid")]
        public Input<string>? Aid { get; set; }

        /// <summary>
        /// Creation time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Creator ID.
        /// </summary>
        [Input("createUser")]
        public Input<string>? CreateUser { get; set; }

        /// <summary>
        /// Instance ID, such as taw-123.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Remarks.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// End time.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// uin: business identifier.
        /// </summary>
        [Input("whitelistUin")]
        public Input<string>? WhitelistUin { get; set; }

        /// <summary>
        /// Auto-Increment allowlist ID.
        /// </summary>
        [Input("wid")]
        public Input<string>? Wid { get; set; }

        public WhitelistState()
        {
        }
        public static new WhitelistState Empty => new WhitelistState();
    }
}
