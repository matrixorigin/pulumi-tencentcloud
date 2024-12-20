// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elasticsearch
{
    [TencentcloudResourceType("tencentcloud:Elasticsearch/restartLogstashInstanceOperation:RestartLogstashInstanceOperation")]
    public partial class RestartLogstashInstanceOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Restart type, 0 full restart, 1 rolling restart.
        /// </summary>
        [Output("type")]
        public Output<int> Type { get; private set; } = null!;


        /// <summary>
        /// Create a RestartLogstashInstanceOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RestartLogstashInstanceOperation(string name, RestartLogstashInstanceOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Elasticsearch/restartLogstashInstanceOperation:RestartLogstashInstanceOperation", name, args ?? new RestartLogstashInstanceOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RestartLogstashInstanceOperation(string name, Input<string> id, RestartLogstashInstanceOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Elasticsearch/restartLogstashInstanceOperation:RestartLogstashInstanceOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RestartLogstashInstanceOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RestartLogstashInstanceOperation Get(string name, Input<string> id, RestartLogstashInstanceOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new RestartLogstashInstanceOperation(name, id, state, options);
        }
    }

    public sealed class RestartLogstashInstanceOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Restart type, 0 full restart, 1 rolling restart.
        /// </summary>
        [Input("type", required: true)]
        public Input<int> Type { get; set; } = null!;

        public RestartLogstashInstanceOperationArgs()
        {
        }
        public static new RestartLogstashInstanceOperationArgs Empty => new RestartLogstashInstanceOperationArgs();
    }

    public sealed class RestartLogstashInstanceOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Restart type, 0 full restart, 1 rolling restart.
        /// </summary>
        [Input("type")]
        public Input<int>? Type { get; set; }

        public RestartLogstashInstanceOperationState()
        {
        }
        public static new RestartLogstashInstanceOperationState Empty => new RestartLogstashInstanceOperationState();
    }
}
