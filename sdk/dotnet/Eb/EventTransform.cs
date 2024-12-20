// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Eb
{
    [TencentcloudResourceType("tencentcloud:Eb/eventTransform:EventTransform")]
    public partial class EventTransform : global::Pulumi.CustomResource
    {
        /// <summary>
        /// event bus Id.
        /// </summary>
        [Output("eventBusId")]
        public Output<string> EventBusId { get; private set; } = null!;

        /// <summary>
        /// ruleId.
        /// </summary>
        [Output("ruleId")]
        public Output<string> RuleId { get; private set; } = null!;

        /// <summary>
        /// A list of transformation rules, currently only one.
        /// </summary>
        [Output("transformations")]
        public Output<ImmutableArray<Outputs.EventTransformTransformation>> Transformations { get; private set; } = null!;


        /// <summary>
        /// Create a EventTransform resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventTransform(string name, EventTransformArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Eb/eventTransform:EventTransform", name, args ?? new EventTransformArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventTransform(string name, Input<string> id, EventTransformState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Eb/eventTransform:EventTransform", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventTransform resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventTransform Get(string name, Input<string> id, EventTransformState? state = null, CustomResourceOptions? options = null)
        {
            return new EventTransform(name, id, state, options);
        }
    }

    public sealed class EventTransformArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// event bus Id.
        /// </summary>
        [Input("eventBusId", required: true)]
        public Input<string> EventBusId { get; set; } = null!;

        /// <summary>
        /// ruleId.
        /// </summary>
        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        [Input("transformations", required: true)]
        private InputList<Inputs.EventTransformTransformationArgs>? _transformations;

        /// <summary>
        /// A list of transformation rules, currently only one.
        /// </summary>
        public InputList<Inputs.EventTransformTransformationArgs> Transformations
        {
            get => _transformations ?? (_transformations = new InputList<Inputs.EventTransformTransformationArgs>());
            set => _transformations = value;
        }

        public EventTransformArgs()
        {
        }
        public static new EventTransformArgs Empty => new EventTransformArgs();
    }

    public sealed class EventTransformState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// event bus Id.
        /// </summary>
        [Input("eventBusId")]
        public Input<string>? EventBusId { get; set; }

        /// <summary>
        /// ruleId.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        [Input("transformations")]
        private InputList<Inputs.EventTransformTransformationGetArgs>? _transformations;

        /// <summary>
        /// A list of transformation rules, currently only one.
        /// </summary>
        public InputList<Inputs.EventTransformTransformationGetArgs> Transformations
        {
            get => _transformations ?? (_transformations = new InputList<Inputs.EventTransformTransformationGetArgs>());
            set => _transformations = value;
        }

        public EventTransformState()
        {
        }
        public static new EventTransformState Empty => new EventTransformState();
    }
}
