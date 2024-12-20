// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Gaap
{
    public static class GetRegionAndPrice
    {
        public static Task<GetRegionAndPriceResult> InvokeAsync(GetRegionAndPriceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegionAndPriceResult>("tencentcloud:Gaap/getRegionAndPrice:getRegionAndPrice", args ?? new GetRegionAndPriceArgs(), options.WithDefaults());

        public static Output<GetRegionAndPriceResult> Invoke(GetRegionAndPriceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegionAndPriceResult>("tencentcloud:Gaap/getRegionAndPrice:getRegionAndPrice", args ?? new GetRegionAndPriceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionAndPriceArgs : global::Pulumi.InvokeArgs
    {
        [Input("ipAddressVersion")]
        public string? IpAddressVersion { get; set; }

        [Input("packageType")]
        public string? PackageType { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetRegionAndPriceArgs()
        {
        }
        public static new GetRegionAndPriceArgs Empty => new GetRegionAndPriceArgs();
    }

    public sealed class GetRegionAndPriceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ipAddressVersion")]
        public Input<string>? IpAddressVersion { get; set; }

        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetRegionAndPriceInvokeArgs()
        {
        }
        public static new GetRegionAndPriceInvokeArgs Empty => new GetRegionAndPriceInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegionAndPriceResult
    {
        public readonly ImmutableArray<Outputs.GetRegionAndPriceBandwidthUnitPriceResult> BandwidthUnitPrices;
        public readonly string Currency;
        public readonly ImmutableArray<Outputs.GetRegionAndPriceDestRegionSetResult> DestRegionSets;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IpAddressVersion;
        public readonly string? PackageType;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetRegionAndPriceResult(
            ImmutableArray<Outputs.GetRegionAndPriceBandwidthUnitPriceResult> bandwidthUnitPrices,

            string currency,

            ImmutableArray<Outputs.GetRegionAndPriceDestRegionSetResult> destRegionSets,

            string id,

            string? ipAddressVersion,

            string? packageType,

            string? resultOutputFile)
        {
            BandwidthUnitPrices = bandwidthUnitPrices;
            Currency = currency;
            DestRegionSets = destRegionSets;
            Id = id;
            IpAddressVersion = ipAddressVersion;
            PackageType = packageType;
            ResultOutputFile = resultOutputFile;
        }
    }
}
