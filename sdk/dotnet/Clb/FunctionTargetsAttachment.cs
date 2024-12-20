// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clb
{
    [TencentcloudResourceType("tencentcloud:Clb/functionTargetsAttachment:FunctionTargetsAttachment")]
    public partial class FunctionTargetsAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The domain name of the target forwarding rule. If the LocationId parameter has been entered, this parameter will not
        /// take effect.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// List of cloud functions to be bound.
        /// </summary>
        [Output("functionTargets")]
        public Output<Outputs.FunctionTargetsAttachmentFunctionTargets> FunctionTargets { get; private set; } = null!;

        /// <summary>
        /// Load Balancer Listener ID.
        /// </summary>
        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        /// <summary>
        /// Load Balancer Instance ID.
        /// </summary>
        [Output("loadBalancerId")]
        public Output<string> LoadBalancerId { get; private set; } = null!;

        /// <summary>
        /// The ID of the target forwarding rule. When binding the cloud function to a layer-7 forwarding rule, this parameter or
        /// the Domain+Url parameter must be entered.
        /// </summary>
        [Output("locationId")]
        public Output<string> LocationId { get; private set; } = null!;

        /// <summary>
        /// The URL of the target forwarding rule. If the LocationId parameter has been entered, this parameter will not take
        /// effect.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a FunctionTargetsAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FunctionTargetsAttachment(string name, FunctionTargetsAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Clb/functionTargetsAttachment:FunctionTargetsAttachment", name, args ?? new FunctionTargetsAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FunctionTargetsAttachment(string name, Input<string> id, FunctionTargetsAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Clb/functionTargetsAttachment:FunctionTargetsAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FunctionTargetsAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FunctionTargetsAttachment Get(string name, Input<string> id, FunctionTargetsAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new FunctionTargetsAttachment(name, id, state, options);
        }
    }

    public sealed class FunctionTargetsAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain name of the target forwarding rule. If the LocationId parameter has been entered, this parameter will not
        /// take effect.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// List of cloud functions to be bound.
        /// </summary>
        [Input("functionTargets", required: true)]
        public Input<Inputs.FunctionTargetsAttachmentFunctionTargetsArgs> FunctionTargets { get; set; } = null!;

        /// <summary>
        /// Load Balancer Listener ID.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        /// <summary>
        /// Load Balancer Instance ID.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        /// <summary>
        /// The ID of the target forwarding rule. When binding the cloud function to a layer-7 forwarding rule, this parameter or
        /// the Domain+Url parameter must be entered.
        /// </summary>
        [Input("locationId")]
        public Input<string>? LocationId { get; set; }

        /// <summary>
        /// The URL of the target forwarding rule. If the LocationId parameter has been entered, this parameter will not take
        /// effect.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public FunctionTargetsAttachmentArgs()
        {
        }
        public static new FunctionTargetsAttachmentArgs Empty => new FunctionTargetsAttachmentArgs();
    }

    public sealed class FunctionTargetsAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain name of the target forwarding rule. If the LocationId parameter has been entered, this parameter will not
        /// take effect.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// List of cloud functions to be bound.
        /// </summary>
        [Input("functionTargets")]
        public Input<Inputs.FunctionTargetsAttachmentFunctionTargetsGetArgs>? FunctionTargets { get; set; }

        /// <summary>
        /// Load Balancer Listener ID.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// Load Balancer Instance ID.
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        /// <summary>
        /// The ID of the target forwarding rule. When binding the cloud function to a layer-7 forwarding rule, this parameter or
        /// the Domain+Url parameter must be entered.
        /// </summary>
        [Input("locationId")]
        public Input<string>? LocationId { get; set; }

        /// <summary>
        /// The URL of the target forwarding rule. If the LocationId parameter has been entered, this parameter will not take
        /// effect.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public FunctionTargetsAttachmentState()
        {
        }
        public static new FunctionTargetsAttachmentState Empty => new FunctionTargetsAttachmentState();
    }
}
