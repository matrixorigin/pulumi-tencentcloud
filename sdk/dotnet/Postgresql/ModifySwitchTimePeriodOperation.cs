// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Postgresql
{
    [TencentcloudResourceType("tencentcloud:Postgresql/modifySwitchTimePeriodOperation:ModifySwitchTimePeriodOperation")]
    public partial class ModifySwitchTimePeriodOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the instance waiting for a switch.
        /// </summary>
        [Output("dbInstanceId")]
        public Output<string> DbInstanceId { get; private set; } = null!;

        /// <summary>
        /// Valid value: `0` (switch immediately).
        /// </summary>
        [Output("switchTag")]
        public Output<int> SwitchTag { get; private set; } = null!;


        /// <summary>
        /// Create a ModifySwitchTimePeriodOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ModifySwitchTimePeriodOperation(string name, ModifySwitchTimePeriodOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/modifySwitchTimePeriodOperation:ModifySwitchTimePeriodOperation", name, args ?? new ModifySwitchTimePeriodOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ModifySwitchTimePeriodOperation(string name, Input<string> id, ModifySwitchTimePeriodOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/modifySwitchTimePeriodOperation:ModifySwitchTimePeriodOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ModifySwitchTimePeriodOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ModifySwitchTimePeriodOperation Get(string name, Input<string> id, ModifySwitchTimePeriodOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new ModifySwitchTimePeriodOperation(name, id, state, options);
        }
    }

    public sealed class ModifySwitchTimePeriodOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the instance waiting for a switch.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        /// <summary>
        /// Valid value: `0` (switch immediately).
        /// </summary>
        [Input("switchTag", required: true)]
        public Input<int> SwitchTag { get; set; } = null!;

        public ModifySwitchTimePeriodOperationArgs()
        {
        }
        public static new ModifySwitchTimePeriodOperationArgs Empty => new ModifySwitchTimePeriodOperationArgs();
    }

    public sealed class ModifySwitchTimePeriodOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the instance waiting for a switch.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        /// <summary>
        /// Valid value: `0` (switch immediately).
        /// </summary>
        [Input("switchTag")]
        public Input<int>? SwitchTag { get; set; }

        public ModifySwitchTimePeriodOperationState()
        {
        }
        public static new ModifySwitchTimePeriodOperationState Empty => new ModifySwitchTimePeriodOperationState();
    }
}
