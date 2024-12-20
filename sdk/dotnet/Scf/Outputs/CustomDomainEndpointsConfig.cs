// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf.Outputs
{

    [OutputType]
    public sealed class CustomDomainEndpointsConfig
    {
        public readonly string FunctionName;
        public readonly string Namespace;
        public readonly string PathMatch;
        public readonly ImmutableArray<Outputs.CustomDomainEndpointsConfigPathRewrite> PathRewrites;
        public readonly string Qualifier;

        [OutputConstructor]
        private CustomDomainEndpointsConfig(
            string functionName,

            string @namespace,

            string pathMatch,

            ImmutableArray<Outputs.CustomDomainEndpointsConfigPathRewrite> pathRewrites,

            string qualifier)
        {
            FunctionName = functionName;
            Namespace = @namespace;
            PathMatch = pathMatch;
            PathRewrites = pathRewrites;
            Qualifier = qualifier;
        }
    }
}
