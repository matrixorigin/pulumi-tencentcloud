// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.ApiGateway.Outputs
{

    [OutputType]
    public sealed class GetApiAppApiResultServiceConfigResult
    {
        public readonly string Method;
        public readonly string Path;
        public readonly string Product;
        public readonly string UniqVpcId;
        public readonly string UpstreamId;
        public readonly string Url;

        [OutputConstructor]
        private GetApiAppApiResultServiceConfigResult(
            string method,

            string path,

            string product,

            string uniqVpcId,

            string upstreamId,

            string url)
        {
            Method = method;
            Path = path;
            Product = product;
            UniqVpcId = uniqVpcId;
            UpstreamId = upstreamId;
            Url = url;
        }
    }
}