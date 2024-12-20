// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elastic
{
    [TencentcloudResourceType("tencentcloud:Elastic/publicIpv6Attachment:PublicIpv6Attachment")]
    public partial class PublicIpv6Attachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Elastic IPv6 unique ID, EIPv6 unique ID is like eipv6-11112222.
        /// </summary>
        [Output("ipv6AddressId")]
        public Output<string> Ipv6AddressId { get; private set; } = null!;

        /// <summary>
        /// Whether to keep the Elastic Network Interface bound when unbinding.
        /// </summary>
        [Output("keepBindWithEni")]
        public Output<bool?> KeepBindWithEni { get; private set; } = null!;

        /// <summary>
        /// Elastic Network Interface ID to bind. Elastic Network Interface ID is like eni-11112222. NetworkInterfaceId and
        /// InstanceId cannot be specified simultaneously. The Elastic Network Interface ID can be queried by logging in to the
        /// console, or obtained through the networkInterfaceId in the return value of the DescribeNetworkInterfaces interface.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string?> NetworkInterfaceId { get; private set; } = null!;

        /// <summary>
        /// The intranet IPv6 to bind. If NetworkInterfaceId is specified, PrivateIPv6Address must also be specified, which means
        /// that the EIP is bound to the specified private network IP of the specified Elastic Network Interface. Also ensure that
        /// the specified PrivateIPv6Address is an intranet IPv6 on the specified NetworkInterfaceId. The intranet IPv6 of the
        /// specified Elastic Network Interface can be queried by logging in to the console, or obtained through the
        /// Ipv6AddressSet.Address in the return value of the DescribeNetworkInterfaces interface.
        /// </summary>
        [Output("privateIpv6Address")]
        public Output<string?> PrivateIpv6Address { get; private set; } = null!;


        /// <summary>
        /// Create a PublicIpv6Attachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PublicIpv6Attachment(string name, PublicIpv6AttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Elastic/publicIpv6Attachment:PublicIpv6Attachment", name, args ?? new PublicIpv6AttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PublicIpv6Attachment(string name, Input<string> id, PublicIpv6AttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Elastic/publicIpv6Attachment:PublicIpv6Attachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PublicIpv6Attachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PublicIpv6Attachment Get(string name, Input<string> id, PublicIpv6AttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new PublicIpv6Attachment(name, id, state, options);
        }
    }

    public sealed class PublicIpv6AttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Elastic IPv6 unique ID, EIPv6 unique ID is like eipv6-11112222.
        /// </summary>
        [Input("ipv6AddressId", required: true)]
        public Input<string> Ipv6AddressId { get; set; } = null!;

        /// <summary>
        /// Whether to keep the Elastic Network Interface bound when unbinding.
        /// </summary>
        [Input("keepBindWithEni")]
        public Input<bool>? KeepBindWithEni { get; set; }

        /// <summary>
        /// Elastic Network Interface ID to bind. Elastic Network Interface ID is like eni-11112222. NetworkInterfaceId and
        /// InstanceId cannot be specified simultaneously. The Elastic Network Interface ID can be queried by logging in to the
        /// console, or obtained through the networkInterfaceId in the return value of the DescribeNetworkInterfaces interface.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The intranet IPv6 to bind. If NetworkInterfaceId is specified, PrivateIPv6Address must also be specified, which means
        /// that the EIP is bound to the specified private network IP of the specified Elastic Network Interface. Also ensure that
        /// the specified PrivateIPv6Address is an intranet IPv6 on the specified NetworkInterfaceId. The intranet IPv6 of the
        /// specified Elastic Network Interface can be queried by logging in to the console, or obtained through the
        /// Ipv6AddressSet.Address in the return value of the DescribeNetworkInterfaces interface.
        /// </summary>
        [Input("privateIpv6Address")]
        public Input<string>? PrivateIpv6Address { get; set; }

        public PublicIpv6AttachmentArgs()
        {
        }
        public static new PublicIpv6AttachmentArgs Empty => new PublicIpv6AttachmentArgs();
    }

    public sealed class PublicIpv6AttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Elastic IPv6 unique ID, EIPv6 unique ID is like eipv6-11112222.
        /// </summary>
        [Input("ipv6AddressId")]
        public Input<string>? Ipv6AddressId { get; set; }

        /// <summary>
        /// Whether to keep the Elastic Network Interface bound when unbinding.
        /// </summary>
        [Input("keepBindWithEni")]
        public Input<bool>? KeepBindWithEni { get; set; }

        /// <summary>
        /// Elastic Network Interface ID to bind. Elastic Network Interface ID is like eni-11112222. NetworkInterfaceId and
        /// InstanceId cannot be specified simultaneously. The Elastic Network Interface ID can be queried by logging in to the
        /// console, or obtained through the networkInterfaceId in the return value of the DescribeNetworkInterfaces interface.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The intranet IPv6 to bind. If NetworkInterfaceId is specified, PrivateIPv6Address must also be specified, which means
        /// that the EIP is bound to the specified private network IP of the specified Elastic Network Interface. Also ensure that
        /// the specified PrivateIPv6Address is an intranet IPv6 on the specified NetworkInterfaceId. The intranet IPv6 of the
        /// specified Elastic Network Interface can be queried by logging in to the console, or obtained through the
        /// Ipv6AddressSet.Address in the return value of the DescribeNetworkInterfaces interface.
        /// </summary>
        [Input("privateIpv6Address")]
        public Input<string>? PrivateIpv6Address { get; set; }

        public PublicIpv6AttachmentState()
        {
        }
        public static new PublicIpv6AttachmentState Empty => new PublicIpv6AttachmentState();
    }
}
