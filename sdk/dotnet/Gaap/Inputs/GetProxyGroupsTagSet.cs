// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Gaap.Inputs
{

    public sealed class GetProxyGroupsTagSetArgs : global::Pulumi.InvokeArgs
    {
        [Input("tagKey", required: true)]
        public string TagKey { get; set; } = null!;

        [Input("tagValue", required: true)]
        public string TagValue { get; set; } = null!;

        public GetProxyGroupsTagSetArgs()
        {
        }
        public static new GetProxyGroupsTagSetArgs Empty => new GetProxyGroupsTagSetArgs();
    }
}
