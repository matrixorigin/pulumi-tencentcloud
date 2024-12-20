// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dc
{
    [TencentcloudResourceType("tencentcloud:Dc/gatewayCcnRoute:GatewayCcnRoute")]
    public partial class GatewayCcnRoute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// As path list of the BGP.
        /// </summary>
        [Output("asPaths")]
        public Output<ImmutableArray<string>> AsPaths { get; private set; } = null!;

        /// <summary>
        /// A network address segment of IDC.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// ID of the DCG.
        /// </summary>
        [Output("dcgId")]
        public Output<string> DcgId { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayCcnRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayCcnRoute(string name, GatewayCcnRouteArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dc/gatewayCcnRoute:GatewayCcnRoute", name, args ?? new GatewayCcnRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayCcnRoute(string name, Input<string> id, GatewayCcnRouteState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dc/gatewayCcnRoute:GatewayCcnRoute", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayCcnRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayCcnRoute Get(string name, Input<string> id, GatewayCcnRouteState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayCcnRoute(name, id, state, options);
        }
    }

    public sealed class GatewayCcnRouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A network address segment of IDC.
        /// </summary>
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// ID of the DCG.
        /// </summary>
        [Input("dcgId", required: true)]
        public Input<string> DcgId { get; set; } = null!;

        public GatewayCcnRouteArgs()
        {
        }
        public static new GatewayCcnRouteArgs Empty => new GatewayCcnRouteArgs();
    }

    public sealed class GatewayCcnRouteState : global::Pulumi.ResourceArgs
    {
        [Input("asPaths")]
        private InputList<string>? _asPaths;

        /// <summary>
        /// As path list of the BGP.
        /// </summary>
        public InputList<string> AsPaths
        {
            get => _asPaths ?? (_asPaths = new InputList<string>());
            set => _asPaths = value;
        }

        /// <summary>
        /// A network address segment of IDC.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// ID of the DCG.
        /// </summary>
        [Input("dcgId")]
        public Input<string>? DcgId { get; set; }

        public GatewayCcnRouteState()
        {
        }
        public static new GatewayCcnRouteState Empty => new GatewayCcnRouteState();
    }
}
