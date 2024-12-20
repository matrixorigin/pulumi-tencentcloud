// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Eip
{
    [TencentcloudResourceType("tencentcloud:Eip/normalAddressReturn:NormalAddressReturn")]
    public partial class NormalAddressReturn : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The IP address of the EIP, example: 101.35.139.183.
        /// </summary>
        [Output("addressIps")]
        public Output<ImmutableArray<string>> AddressIps { get; private set; } = null!;


        /// <summary>
        /// Create a NormalAddressReturn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NormalAddressReturn(string name, NormalAddressReturnArgs? args = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Eip/normalAddressReturn:NormalAddressReturn", name, args ?? new NormalAddressReturnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NormalAddressReturn(string name, Input<string> id, NormalAddressReturnState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Eip/normalAddressReturn:NormalAddressReturn", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NormalAddressReturn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NormalAddressReturn Get(string name, Input<string> id, NormalAddressReturnState? state = null, CustomResourceOptions? options = null)
        {
            return new NormalAddressReturn(name, id, state, options);
        }
    }

    public sealed class NormalAddressReturnArgs : global::Pulumi.ResourceArgs
    {
        [Input("addressIps")]
        private InputList<string>? _addressIps;

        /// <summary>
        /// The IP address of the EIP, example: 101.35.139.183.
        /// </summary>
        public InputList<string> AddressIps
        {
            get => _addressIps ?? (_addressIps = new InputList<string>());
            set => _addressIps = value;
        }

        public NormalAddressReturnArgs()
        {
        }
        public static new NormalAddressReturnArgs Empty => new NormalAddressReturnArgs();
    }

    public sealed class NormalAddressReturnState : global::Pulumi.ResourceArgs
    {
        [Input("addressIps")]
        private InputList<string>? _addressIps;

        /// <summary>
        /// The IP address of the EIP, example: 101.35.139.183.
        /// </summary>
        public InputList<string> AddressIps
        {
            get => _addressIps ?? (_addressIps = new InputList<string>());
            set => _addressIps = value;
        }

        public NormalAddressReturnState()
        {
        }
        public static new NormalAddressReturnState Empty => new NormalAddressReturnState();
    }
}
