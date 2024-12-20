// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf.Inputs
{

    public sealed class CustomDomainEndpointsConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        [Input("pathMatch", required: true)]
        public Input<string> PathMatch { get; set; } = null!;

        [Input("pathRewrites")]
        private InputList<Inputs.CustomDomainEndpointsConfigPathRewriteArgs>? _pathRewrites;
        public InputList<Inputs.CustomDomainEndpointsConfigPathRewriteArgs> PathRewrites
        {
            get => _pathRewrites ?? (_pathRewrites = new InputList<Inputs.CustomDomainEndpointsConfigPathRewriteArgs>());
            set => _pathRewrites = value;
        }

        [Input("qualifier", required: true)]
        public Input<string> Qualifier { get; set; } = null!;

        public CustomDomainEndpointsConfigArgs()
        {
        }
        public static new CustomDomainEndpointsConfigArgs Empty => new CustomDomainEndpointsConfigArgs();
    }
}
