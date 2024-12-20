// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vpc
{
    public static class GetSubnets
    {
        public static Task<GetSubnetsResult> InvokeAsync(GetSubnetsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSubnetsResult>("tencentcloud:Vpc/getSubnets:getSubnets", args ?? new GetSubnetsArgs(), options.WithDefaults());

        public static Output<GetSubnetsResult> Invoke(GetSubnetsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSubnetsResult>("tencentcloud:Vpc/getSubnets:getSubnets", args ?? new GetSubnetsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSubnetsArgs : global::Pulumi.InvokeArgs
    {
        [Input("availabilityZone")]
        public string? AvailabilityZone { get; set; }

        [Input("cdcId")]
        public string? CdcId { get; set; }

        [Input("cidrBlock")]
        public string? CidrBlock { get; set; }

        [Input("isDefault")]
        public bool? IsDefault { get; set; }

        [Input("isRemoteVpcSnat")]
        public bool? IsRemoteVpcSnat { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("subnetId")]
        public string? SubnetId { get; set; }

        [Input("tagKey")]
        public string? TagKey { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetSubnetsArgs()
        {
        }
        public static new GetSubnetsArgs Empty => new GetSubnetsArgs();
    }

    public sealed class GetSubnetsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("cdcId")]
        public Input<string>? CdcId { get; set; }

        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        [Input("isRemoteVpcSnat")]
        public Input<bool>? IsRemoteVpcSnat { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tagKey")]
        public Input<string>? TagKey { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetSubnetsInvokeArgs()
        {
        }
        public static new GetSubnetsInvokeArgs Empty => new GetSubnetsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSubnetsResult
    {
        public readonly string? AvailabilityZone;
        public readonly string? CdcId;
        public readonly string? CidrBlock;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetSubnetsInstanceListResult> InstanceLists;
        public readonly bool? IsDefault;
        public readonly bool? IsRemoteVpcSnat;
        public readonly string? Name;
        public readonly string? ResultOutputFile;
        public readonly string? SubnetId;
        public readonly string? TagKey;
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly string? VpcId;

        [OutputConstructor]
        private GetSubnetsResult(
            string? availabilityZone,

            string? cdcId,

            string? cidrBlock,

            string id,

            ImmutableArray<Outputs.GetSubnetsInstanceListResult> instanceLists,

            bool? isDefault,

            bool? isRemoteVpcSnat,

            string? name,

            string? resultOutputFile,

            string? subnetId,

            string? tagKey,

            ImmutableDictionary<string, object>? tags,

            string? vpcId)
        {
            AvailabilityZone = availabilityZone;
            CdcId = cdcId;
            CidrBlock = cidrBlock;
            Id = id;
            InstanceLists = instanceLists;
            IsDefault = isDefault;
            IsRemoteVpcSnat = isRemoteVpcSnat;
            Name = name;
            ResultOutputFile = resultOutputFile;
            SubnetId = subnetId;
            TagKey = tagKey;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}
