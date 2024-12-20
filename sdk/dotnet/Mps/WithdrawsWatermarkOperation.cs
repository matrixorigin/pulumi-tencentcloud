// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps
{
    [TencentcloudResourceType("tencentcloud:Mps/withdrawsWatermarkOperation:WithdrawsWatermarkOperation")]
    public partial class WithdrawsWatermarkOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Input information of file for metadata getting.
        /// </summary>
        [Output("inputInfo")]
        public Output<Outputs.WithdrawsWatermarkOperationInputInfo> InputInfo { get; private set; } = null!;

        /// <summary>
        /// The source context which is used to pass through the user request information. The task flow status change callback will
        /// return the value of this field.
        /// </summary>
        [Output("sessionContext")]
        public Output<string?> SessionContext { get; private set; } = null!;

        /// <summary>
        /// Event notification information of a task. If this parameter is left empty, no event notifications will be obtained.
        /// </summary>
        [Output("taskNotifyConfig")]
        public Output<Outputs.WithdrawsWatermarkOperationTaskNotifyConfig?> TaskNotifyConfig { get; private set; } = null!;


        /// <summary>
        /// Create a WithdrawsWatermarkOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WithdrawsWatermarkOperation(string name, WithdrawsWatermarkOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mps/withdrawsWatermarkOperation:WithdrawsWatermarkOperation", name, args ?? new WithdrawsWatermarkOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WithdrawsWatermarkOperation(string name, Input<string> id, WithdrawsWatermarkOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mps/withdrawsWatermarkOperation:WithdrawsWatermarkOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WithdrawsWatermarkOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WithdrawsWatermarkOperation Get(string name, Input<string> id, WithdrawsWatermarkOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new WithdrawsWatermarkOperation(name, id, state, options);
        }
    }

    public sealed class WithdrawsWatermarkOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Input information of file for metadata getting.
        /// </summary>
        [Input("inputInfo", required: true)]
        public Input<Inputs.WithdrawsWatermarkOperationInputInfoArgs> InputInfo { get; set; } = null!;

        /// <summary>
        /// The source context which is used to pass through the user request information. The task flow status change callback will
        /// return the value of this field.
        /// </summary>
        [Input("sessionContext")]
        public Input<string>? SessionContext { get; set; }

        /// <summary>
        /// Event notification information of a task. If this parameter is left empty, no event notifications will be obtained.
        /// </summary>
        [Input("taskNotifyConfig")]
        public Input<Inputs.WithdrawsWatermarkOperationTaskNotifyConfigArgs>? TaskNotifyConfig { get; set; }

        public WithdrawsWatermarkOperationArgs()
        {
        }
        public static new WithdrawsWatermarkOperationArgs Empty => new WithdrawsWatermarkOperationArgs();
    }

    public sealed class WithdrawsWatermarkOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Input information of file for metadata getting.
        /// </summary>
        [Input("inputInfo")]
        public Input<Inputs.WithdrawsWatermarkOperationInputInfoGetArgs>? InputInfo { get; set; }

        /// <summary>
        /// The source context which is used to pass through the user request information. The task flow status change callback will
        /// return the value of this field.
        /// </summary>
        [Input("sessionContext")]
        public Input<string>? SessionContext { get; set; }

        /// <summary>
        /// Event notification information of a task. If this parameter is left empty, no event notifications will be obtained.
        /// </summary>
        [Input("taskNotifyConfig")]
        public Input<Inputs.WithdrawsWatermarkOperationTaskNotifyConfigGetArgs>? TaskNotifyConfig { get; set; }

        public WithdrawsWatermarkOperationState()
        {
        }
        public static new WithdrawsWatermarkOperationState Empty => new WithdrawsWatermarkOperationState();
    }
}
