// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cynosdb
{
    [TencentcloudResourceType("tencentcloud:Cynosdb/securityGroup:SecurityGroup")]
    public partial class SecurityGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster id.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Instance group type. Available values: -`HA` - HA group; -`RO` - Read-only group; -`ALL` - HA and RO group.
        /// </summary>
        [Output("instanceGroupType")]
        public Output<string> InstanceGroupType { get; private set; } = null!;

        /// <summary>
        /// A list of security group IDs to be modified, an array of one or more security group IDs.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroup(string name, SecurityGroupArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/securityGroup:SecurityGroup", name, args ?? new SecurityGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroup(string name, Input<string> id, SecurityGroupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/securityGroup:SecurityGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroup Get(string name, Input<string> id, SecurityGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityGroup(name, id, state, options);
        }
    }

    public sealed class SecurityGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster id.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Instance group type. Available values: -`HA` - HA group; -`RO` - Read-only group; -`ALL` - HA and RO group.
        /// </summary>
        [Input("instanceGroupType", required: true)]
        public Input<string> InstanceGroupType { get; set; } = null!;

        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security group IDs to be modified, an array of one or more security group IDs.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        public SecurityGroupArgs()
        {
        }
        public static new SecurityGroupArgs Empty => new SecurityGroupArgs();
    }

    public sealed class SecurityGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster id.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Instance group type. Available values: -`HA` - HA group; -`RO` - Read-only group; -`ALL` - HA and RO group.
        /// </summary>
        [Input("instanceGroupType")]
        public Input<string>? InstanceGroupType { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security group IDs to be modified, an array of one or more security group IDs.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        public SecurityGroupState()
        {
        }
        public static new SecurityGroupState Empty => new SecurityGroupState();
    }
}
