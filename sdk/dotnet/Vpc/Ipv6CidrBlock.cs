// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vpc
{
    [TencentcloudResourceType("tencentcloud:Vpc/ipv6CidrBlock:Ipv6CidrBlock")]
    public partial class Ipv6CidrBlock : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ipv6 cidr block.
        /// </summary>
        [Output("ipv6CidrBlock")]
        public Output<string> VpcIpv6CidrBlock { get; private set; } = null!;

        /// <summary>
        /// `VPC` instance `ID`, in the form of `vpc-f49l6u0z`.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Ipv6CidrBlock resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipv6CidrBlock(string name, Ipv6CidrBlockArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/ipv6CidrBlock:Ipv6CidrBlock", name, args ?? new Ipv6CidrBlockArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipv6CidrBlock(string name, Input<string> id, Ipv6CidrBlockState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/ipv6CidrBlock:Ipv6CidrBlock", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipv6CidrBlock resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipv6CidrBlock Get(string name, Input<string> id, Ipv6CidrBlockState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipv6CidrBlock(name, id, state, options);
        }
    }

    public sealed class Ipv6CidrBlockArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `VPC` instance `ID`, in the form of `vpc-f49l6u0z`.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public Ipv6CidrBlockArgs()
        {
        }
        public static new Ipv6CidrBlockArgs Empty => new Ipv6CidrBlockArgs();
    }

    public sealed class Ipv6CidrBlockState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ipv6 cidr block.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? VpcIpv6CidrBlock { get; set; }

        /// <summary>
        /// `VPC` instance `ID`, in the form of `vpc-f49l6u0z`.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public Ipv6CidrBlockState()
        {
        }
        public static new Ipv6CidrBlockState Empty => new Ipv6CidrBlockState();
    }
}