// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cynosdb
{
    [TencentcloudResourceType("tencentcloud:Cynosdb/clusterSlaveZone:ClusterSlaveZone")]
    public partial class ClusterSlaveZone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Slave zone.
        /// </summary>
        [Output("slaveZone")]
        public Output<string> SlaveZone { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterSlaveZone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterSlaveZone(string name, ClusterSlaveZoneArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/clusterSlaveZone:ClusterSlaveZone", name, args ?? new ClusterSlaveZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterSlaveZone(string name, Input<string> id, ClusterSlaveZoneState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/clusterSlaveZone:ClusterSlaveZone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClusterSlaveZone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterSlaveZone Get(string name, Input<string> id, ClusterSlaveZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterSlaveZone(name, id, state, options);
        }
    }

    public sealed class ClusterSlaveZoneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Slave zone.
        /// </summary>
        [Input("slaveZone", required: true)]
        public Input<string> SlaveZone { get; set; } = null!;

        public ClusterSlaveZoneArgs()
        {
        }
        public static new ClusterSlaveZoneArgs Empty => new ClusterSlaveZoneArgs();
    }

    public sealed class ClusterSlaveZoneState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Slave zone.
        /// </summary>
        [Input("slaveZone")]
        public Input<string>? SlaveZone { get; set; }

        public ClusterSlaveZoneState()
        {
        }
        public static new ClusterSlaveZoneState Empty => new ClusterSlaveZoneState();
    }
}
