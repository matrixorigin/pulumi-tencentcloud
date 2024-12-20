// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Deprecatedroute
{
    [TencentcloudResourceType("tencentcloud:Deprecatedroute/entry:Entry")]
    public partial class Entry : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The RouteEntry's target network segment.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The route entry's next hub. CVM instance ID or VPC router interface ID.
        /// </summary>
        [Output("nextHub")]
        public Output<string> NextHub { get; private set; } = null!;

        /// <summary>
        /// The next hop type. Valid values:
        /// `public_gateway`,`vpn_gateway`,`sslvpn_gateway`,`dc_gateway`,`peering_connection`,`nat_gateway`,`havip`,`local_gateway`
        /// and `instance`. `instance` points to CVM Instance.
        /// </summary>
        [Output("nextType")]
        public Output<string> NextType { get; private set; } = null!;

        /// <summary>
        /// The ID of the route table.
        /// </summary>
        [Output("routeTableId")]
        public Output<string> RouteTableId { get; private set; } = null!;

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Entry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Entry(string name, EntryArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Deprecatedroute/entry:Entry", name, args ?? new EntryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Entry(string name, Input<string> id, EntryState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Deprecatedroute/entry:Entry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Entry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Entry Get(string name, Input<string> id, EntryState? state = null, CustomResourceOptions? options = null)
        {
            return new Entry(name, id, state, options);
        }
    }

    public sealed class EntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The RouteEntry's target network segment.
        /// </summary>
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// The route entry's next hub. CVM instance ID or VPC router interface ID.
        /// </summary>
        [Input("nextHub", required: true)]
        public Input<string> NextHub { get; set; } = null!;

        /// <summary>
        /// The next hop type. Valid values:
        /// `public_gateway`,`vpn_gateway`,`sslvpn_gateway`,`dc_gateway`,`peering_connection`,`nat_gateway`,`havip`,`local_gateway`
        /// and `instance`. `instance` points to CVM Instance.
        /// </summary>
        [Input("nextType", required: true)]
        public Input<string> NextType { get; set; } = null!;

        /// <summary>
        /// The ID of the route table.
        /// </summary>
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public EntryArgs()
        {
        }
        public static new EntryArgs Empty => new EntryArgs();
    }

    public sealed class EntryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The RouteEntry's target network segment.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The route entry's next hub. CVM instance ID or VPC router interface ID.
        /// </summary>
        [Input("nextHub")]
        public Input<string>? NextHub { get; set; }

        /// <summary>
        /// The next hop type. Valid values:
        /// `public_gateway`,`vpn_gateway`,`sslvpn_gateway`,`dc_gateway`,`peering_connection`,`nat_gateway`,`havip`,`local_gateway`
        /// and `instance`. `instance` points to CVM Instance.
        /// </summary>
        [Input("nextType")]
        public Input<string>? NextType { get; set; }

        /// <summary>
        /// The ID of the route table.
        /// </summary>
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public EntryState()
        {
        }
        public static new EntryState Empty => new EntryState();
    }
}
