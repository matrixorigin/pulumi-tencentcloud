// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdc
{
    [TencentcloudResourceType("tencentcloud:Cdc/dedicatedCluster:DedicatedCluster")]
    public partial class DedicatedCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Dedicated Cluster Description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Dedicated Cluster Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Dedicated Cluster Site ID.
        /// </summary>
        [Output("siteId")]
        public Output<string> SiteId { get; private set; } = null!;

        /// <summary>
        /// Dedicated Cluster Zone.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a DedicatedCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DedicatedCluster(string name, DedicatedClusterArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cdc/dedicatedCluster:DedicatedCluster", name, args ?? new DedicatedClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DedicatedCluster(string name, Input<string> id, DedicatedClusterState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cdc/dedicatedCluster:DedicatedCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DedicatedCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DedicatedCluster Get(string name, Input<string> id, DedicatedClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new DedicatedCluster(name, id, state, options);
        }
    }

    public sealed class DedicatedClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dedicated Cluster Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Dedicated Cluster Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Dedicated Cluster Site ID.
        /// </summary>
        [Input("siteId", required: true)]
        public Input<string> SiteId { get; set; } = null!;

        /// <summary>
        /// Dedicated Cluster Zone.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public DedicatedClusterArgs()
        {
        }
        public static new DedicatedClusterArgs Empty => new DedicatedClusterArgs();
    }

    public sealed class DedicatedClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dedicated Cluster Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Dedicated Cluster Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Dedicated Cluster Site ID.
        /// </summary>
        [Input("siteId")]
        public Input<string>? SiteId { get; set; }

        /// <summary>
        /// Dedicated Cluster Zone.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DedicatedClusterState()
        {
        }
        public static new DedicatedClusterState Empty => new DedicatedClusterState();
    }
}