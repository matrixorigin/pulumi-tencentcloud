// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kubernetes
{
    [TencentcloudResourceType("tencentcloud:Kubernetes/healthCheckPolicy:HealthCheckPolicy")]
    public partial class HealthCheckPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Health Check Policy Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Health check policy rule list.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.HealthCheckPolicyRule>> Rules { get; private set; } = null!;


        /// <summary>
        /// Create a HealthCheckPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HealthCheckPolicy(string name, HealthCheckPolicyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/healthCheckPolicy:HealthCheckPolicy", name, args ?? new HealthCheckPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HealthCheckPolicy(string name, Input<string> id, HealthCheckPolicyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/healthCheckPolicy:HealthCheckPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HealthCheckPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HealthCheckPolicy Get(string name, Input<string> id, HealthCheckPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new HealthCheckPolicy(name, id, state, options);
        }
    }

    public sealed class HealthCheckPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Health Check Policy Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rules", required: true)]
        private InputList<Inputs.HealthCheckPolicyRuleArgs>? _rules;

        /// <summary>
        /// Health check policy rule list.
        /// </summary>
        public InputList<Inputs.HealthCheckPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.HealthCheckPolicyRuleArgs>());
            set => _rules = value;
        }

        public HealthCheckPolicyArgs()
        {
        }
        public static new HealthCheckPolicyArgs Empty => new HealthCheckPolicyArgs();
    }

    public sealed class HealthCheckPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Health Check Policy Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rules")]
        private InputList<Inputs.HealthCheckPolicyRuleGetArgs>? _rules;

        /// <summary>
        /// Health check policy rule list.
        /// </summary>
        public InputList<Inputs.HealthCheckPolicyRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.HealthCheckPolicyRuleGetArgs>());
            set => _rules = value;
        }

        public HealthCheckPolicyState()
        {
        }
        public static new HealthCheckPolicyState Empty => new HealthCheckPolicyState();
    }
}