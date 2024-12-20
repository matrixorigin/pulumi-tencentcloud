// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Waf
{
    [TencentcloudResourceType("tencentcloud:Waf/protectionMode:ProtectionMode")]
    public partial class ProtectionMode : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf, default is sparta-waf.
        /// </summary>
        [Output("edition")]
        public Output<string?> Edition { get; private set; } = null!;

        /// <summary>
        /// Protection status:10: Rule observation; AI off mode, 11: Rule observation; AI observation mode, 12: Rule observation; AI
        /// interception mode20: Rule interception; AI off mode, 21: Rule interception; AI observation mode, 22: Rule interception;
        /// AI interception mode.
        /// </summary>
        [Output("mode")]
        public Output<int> Mode { get; private set; } = null!;

        /// <summary>
        /// 0 is to modify the rule engine status, 1 is to modify the AI status.
        /// </summary>
        [Output("type")]
        public Output<int?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ProtectionMode resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProtectionMode(string name, ProtectionModeArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Waf/protectionMode:ProtectionMode", name, args ?? new ProtectionModeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProtectionMode(string name, Input<string> id, ProtectionModeState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Waf/protectionMode:ProtectionMode", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProtectionMode resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProtectionMode Get(string name, Input<string> id, ProtectionModeState? state = null, CustomResourceOptions? options = null)
        {
            return new ProtectionMode(name, id, state, options);
        }
    }

    public sealed class ProtectionModeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf, default is sparta-waf.
        /// </summary>
        [Input("edition")]
        public Input<string>? Edition { get; set; }

        /// <summary>
        /// Protection status:10: Rule observation; AI off mode, 11: Rule observation; AI observation mode, 12: Rule observation; AI
        /// interception mode20: Rule interception; AI off mode, 21: Rule interception; AI observation mode, 22: Rule interception;
        /// AI interception mode.
        /// </summary>
        [Input("mode", required: true)]
        public Input<int> Mode { get; set; } = null!;

        /// <summary>
        /// 0 is to modify the rule engine status, 1 is to modify the AI status.
        /// </summary>
        [Input("type")]
        public Input<int>? Type { get; set; }

        public ProtectionModeArgs()
        {
        }
        public static new ProtectionModeArgs Empty => new ProtectionModeArgs();
    }

    public sealed class ProtectionModeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf, default is sparta-waf.
        /// </summary>
        [Input("edition")]
        public Input<string>? Edition { get; set; }

        /// <summary>
        /// Protection status:10: Rule observation; AI off mode, 11: Rule observation; AI observation mode, 12: Rule observation; AI
        /// interception mode20: Rule interception; AI off mode, 21: Rule interception; AI observation mode, 22: Rule interception;
        /// AI interception mode.
        /// </summary>
        [Input("mode")]
        public Input<int>? Mode { get; set; }

        /// <summary>
        /// 0 is to modify the rule engine status, 1 is to modify the AI status.
        /// </summary>
        [Input("type")]
        public Input<int>? Type { get; set; }

        public ProtectionModeState()
        {
        }
        public static new ProtectionModeState Empty => new ProtectionModeState();
    }
}
