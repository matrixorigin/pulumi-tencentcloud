// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kubernetes
{
    [TencentcloudResourceType("tencentcloud:Kubernetes/serverlessNodePool:ServerlessNodePool")]
    public partial class ServerlessNodePool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// cluster id of serverless node pool.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// labels of serverless node.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, object>?> Labels { get; private set; } = null!;

        /// <summary>
        /// life state of serverless node pool.
        /// </summary>
        [Output("lifeState")]
        public Output<string> LifeState { get; private set; } = null!;

        /// <summary>
        /// serverless node pool name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// security groups of serverless node pool.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// node list of serverless node pool.
        /// </summary>
        [Output("serverlessNodes")]
        public Output<ImmutableArray<Outputs.ServerlessNodePoolServerlessNode>> ServerlessNodes { get; private set; } = null!;

        /// <summary>
        /// taints of serverless node.
        /// </summary>
        [Output("taints")]
        public Output<ImmutableArray<Outputs.ServerlessNodePoolTaint>> Taints { get; private set; } = null!;


        /// <summary>
        /// Create a ServerlessNodePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerlessNodePool(string name, ServerlessNodePoolArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/serverlessNodePool:ServerlessNodePool", name, args ?? new ServerlessNodePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerlessNodePool(string name, Input<string> id, ServerlessNodePoolState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/serverlessNodePool:ServerlessNodePool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerlessNodePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerlessNodePool Get(string name, Input<string> id, ServerlessNodePoolState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerlessNodePool(name, id, state, options);
        }
    }

    public sealed class ServerlessNodePoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// cluster id of serverless node pool.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// labels of serverless node.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// serverless node pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// security groups of serverless node pool.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("serverlessNodes", required: true)]
        private InputList<Inputs.ServerlessNodePoolServerlessNodeArgs>? _serverlessNodes;

        /// <summary>
        /// node list of serverless node pool.
        /// </summary>
        public InputList<Inputs.ServerlessNodePoolServerlessNodeArgs> ServerlessNodes
        {
            get => _serverlessNodes ?? (_serverlessNodes = new InputList<Inputs.ServerlessNodePoolServerlessNodeArgs>());
            set => _serverlessNodes = value;
        }

        [Input("taints")]
        private InputList<Inputs.ServerlessNodePoolTaintArgs>? _taints;

        /// <summary>
        /// taints of serverless node.
        /// </summary>
        public InputList<Inputs.ServerlessNodePoolTaintArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.ServerlessNodePoolTaintArgs>());
            set => _taints = value;
        }

        public ServerlessNodePoolArgs()
        {
        }
        public static new ServerlessNodePoolArgs Empty => new ServerlessNodePoolArgs();
    }

    public sealed class ServerlessNodePoolState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// cluster id of serverless node pool.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// labels of serverless node.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// life state of serverless node pool.
        /// </summary>
        [Input("lifeState")]
        public Input<string>? LifeState { get; set; }

        /// <summary>
        /// serverless node pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// security groups of serverless node pool.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("serverlessNodes")]
        private InputList<Inputs.ServerlessNodePoolServerlessNodeGetArgs>? _serverlessNodes;

        /// <summary>
        /// node list of serverless node pool.
        /// </summary>
        public InputList<Inputs.ServerlessNodePoolServerlessNodeGetArgs> ServerlessNodes
        {
            get => _serverlessNodes ?? (_serverlessNodes = new InputList<Inputs.ServerlessNodePoolServerlessNodeGetArgs>());
            set => _serverlessNodes = value;
        }

        [Input("taints")]
        private InputList<Inputs.ServerlessNodePoolTaintGetArgs>? _taints;

        /// <summary>
        /// taints of serverless node.
        /// </summary>
        public InputList<Inputs.ServerlessNodePoolTaintGetArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.ServerlessNodePoolTaintGetArgs>());
            set => _taints = value;
        }

        public ServerlessNodePoolState()
        {
        }
        public static new ServerlessNodePoolState Empty => new ServerlessNodePoolState();
    }
}
