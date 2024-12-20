// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cfw
{
    [TencentcloudResourceType("tencentcloud:Cfw/blockIgnore:BlockIgnore")]
    public partial class BlockIgnore : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Remarks information, length cannot exceed 50.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Rule direction, 0 outbound, 1 inbound, 3 intranet.
        /// </summary>
        [Output("direction")]
        public Output<string> Direction { get; private set; } = null!;

        /// <summary>
        /// Rule domain name, one of IP and Domain is required.
        /// </summary>
        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        /// <summary>
        /// Rule end time, format: 2006-01-02 15:04:05, must be greater than the current time.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Rule IP address, one of IP and Domain is required.
        /// </summary>
        [Output("ip")]
        public Output<string?> Ip { get; private set; } = null!;

        /// <summary>
        /// Rule type, 1 block, 2 ignore, domain block is not supported.
        /// </summary>
        [Output("ruleType")]
        public Output<int> RuleType { get; private set; } = null!;

        /// <summary>
        /// Rule start time.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;


        /// <summary>
        /// Create a BlockIgnore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BlockIgnore(string name, BlockIgnoreArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfw/blockIgnore:BlockIgnore", name, args ?? new BlockIgnoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BlockIgnore(string name, Input<string> id, BlockIgnoreState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfw/blockIgnore:BlockIgnore", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BlockIgnore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BlockIgnore Get(string name, Input<string> id, BlockIgnoreState? state = null, CustomResourceOptions? options = null)
        {
            return new BlockIgnore(name, id, state, options);
        }
    }

    public sealed class BlockIgnoreArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Remarks information, length cannot exceed 50.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Rule direction, 0 outbound, 1 inbound, 3 intranet.
        /// </summary>
        [Input("direction", required: true)]
        public Input<string> Direction { get; set; } = null!;

        /// <summary>
        /// Rule domain name, one of IP and Domain is required.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Rule end time, format: 2006-01-02 15:04:05, must be greater than the current time.
        /// </summary>
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        /// <summary>
        /// Rule IP address, one of IP and Domain is required.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Rule type, 1 block, 2 ignore, domain block is not supported.
        /// </summary>
        [Input("ruleType", required: true)]
        public Input<int> RuleType { get; set; } = null!;

        /// <summary>
        /// Rule start time.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public BlockIgnoreArgs()
        {
        }
        public static new BlockIgnoreArgs Empty => new BlockIgnoreArgs();
    }

    public sealed class BlockIgnoreState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Remarks information, length cannot exceed 50.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Rule direction, 0 outbound, 1 inbound, 3 intranet.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// Rule domain name, one of IP and Domain is required.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Rule end time, format: 2006-01-02 15:04:05, must be greater than the current time.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Rule IP address, one of IP and Domain is required.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Rule type, 1 block, 2 ignore, domain block is not supported.
        /// </summary>
        [Input("ruleType")]
        public Input<int>? RuleType { get; set; }

        /// <summary>
        /// Rule start time.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public BlockIgnoreState()
        {
        }
        public static new BlockIgnoreState Empty => new BlockIgnoreState();
    }
}
