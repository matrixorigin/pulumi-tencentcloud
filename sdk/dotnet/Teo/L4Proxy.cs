// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Teo
{
    [TencentcloudResourceType("tencentcloud:Teo/l4Proxy:L4Proxy")]
    public partial class L4Proxy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether to enable network optimization in the Chinese mainland. The default value off is used if left empty.
        /// This configuration can only be enabled in certain acceleration zones and security protection configurations. For
        /// details, see [Creating an L4 Proxy
        /// Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable;
        /// `off`: Disable.
        /// </summary>
        [Output("accelerateMainland")]
        public Output<string?> AccelerateMainland { get; private set; } = null!;

        /// <summary>
        /// Acceleration zone of the Layer 4 proxy instance. `mainland`: Availability zone in the Chinese mainland; `overseas`:
        /// Global availability zone (excluding the Chinese mainland); `global`: Global availability zone.
        /// </summary>
        [Output("area")]
        public Output<string?> Area { get; private set; } = null!;

        /// <summary>
        /// Layer 3/Layer 4 DDoS protection. The default protection option of the platform will be used if it is left empty. For
        /// details, see [Exclusive DDoS Protection
        /// Usage](https://intl.cloud.tencent.com/document/product/1552/95994?from_cn_redirect=1).
        /// </summary>
        [Output("ddosProtectionConfig")]
        public Output<Outputs.L4ProxyDdosProtectionConfig?> DdosProtectionConfig { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable IPv6 access. The default value off is used if left empty. This configuration can only be
        /// enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy
        /// Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable;
        /// `off`: Disable.
        /// </summary>
        [Output("ipv6")]
        public Output<string?> Ipv6 { get; private set; } = null!;

        /// <summary>
        /// Layer 4 proxy instance name. You can enter 1-50 characters. Valid characters are a-z, 0-9, and hyphens (-). However,
        /// hyphens (-) cannot be used individually or consecutively and should not be placed at the beginning or end of the name.
        /// Modifications are not allowed after creation.
        /// </summary>
        [Output("proxyName")]
        public Output<string> ProxyName { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable the fixed IP address. The default value off is used if left empty. This configuration can
        /// only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4
        /// Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`:
        /// Enable; `off`: Disable.
        /// </summary>
        [Output("staticIp")]
        public Output<string?> StaticIp { get; private set; } = null!;

        /// <summary>
        /// Site ID.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a L4Proxy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public L4Proxy(string name, L4ProxyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Teo/l4Proxy:L4Proxy", name, args ?? new L4ProxyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private L4Proxy(string name, Input<string> id, L4ProxyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Teo/l4Proxy:L4Proxy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing L4Proxy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static L4Proxy Get(string name, Input<string> id, L4ProxyState? state = null, CustomResourceOptions? options = null)
        {
            return new L4Proxy(name, id, state, options);
        }
    }

    public sealed class L4ProxyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to enable network optimization in the Chinese mainland. The default value off is used if left empty.
        /// This configuration can only be enabled in certain acceleration zones and security protection configurations. For
        /// details, see [Creating an L4 Proxy
        /// Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable;
        /// `off`: Disable.
        /// </summary>
        [Input("accelerateMainland")]
        public Input<string>? AccelerateMainland { get; set; }

        /// <summary>
        /// Acceleration zone of the Layer 4 proxy instance. `mainland`: Availability zone in the Chinese mainland; `overseas`:
        /// Global availability zone (excluding the Chinese mainland); `global`: Global availability zone.
        /// </summary>
        [Input("area")]
        public Input<string>? Area { get; set; }

        /// <summary>
        /// Layer 3/Layer 4 DDoS protection. The default protection option of the platform will be used if it is left empty. For
        /// details, see [Exclusive DDoS Protection
        /// Usage](https://intl.cloud.tencent.com/document/product/1552/95994?from_cn_redirect=1).
        /// </summary>
        [Input("ddosProtectionConfig")]
        public Input<Inputs.L4ProxyDdosProtectionConfigArgs>? DdosProtectionConfig { get; set; }

        /// <summary>
        /// Specifies whether to enable IPv6 access. The default value off is used if left empty. This configuration can only be
        /// enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy
        /// Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable;
        /// `off`: Disable.
        /// </summary>
        [Input("ipv6")]
        public Input<string>? Ipv6 { get; set; }

        /// <summary>
        /// Layer 4 proxy instance name. You can enter 1-50 characters. Valid characters are a-z, 0-9, and hyphens (-). However,
        /// hyphens (-) cannot be used individually or consecutively and should not be placed at the beginning or end of the name.
        /// Modifications are not allowed after creation.
        /// </summary>
        [Input("proxyName", required: true)]
        public Input<string> ProxyName { get; set; } = null!;

        /// <summary>
        /// Specifies whether to enable the fixed IP address. The default value off is used if left empty. This configuration can
        /// only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4
        /// Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`:
        /// Enable; `off`: Disable.
        /// </summary>
        [Input("staticIp")]
        public Input<string>? StaticIp { get; set; }

        /// <summary>
        /// Site ID.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public L4ProxyArgs()
        {
        }
        public static new L4ProxyArgs Empty => new L4ProxyArgs();
    }

    public sealed class L4ProxyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to enable network optimization in the Chinese mainland. The default value off is used if left empty.
        /// This configuration can only be enabled in certain acceleration zones and security protection configurations. For
        /// details, see [Creating an L4 Proxy
        /// Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable;
        /// `off`: Disable.
        /// </summary>
        [Input("accelerateMainland")]
        public Input<string>? AccelerateMainland { get; set; }

        /// <summary>
        /// Acceleration zone of the Layer 4 proxy instance. `mainland`: Availability zone in the Chinese mainland; `overseas`:
        /// Global availability zone (excluding the Chinese mainland); `global`: Global availability zone.
        /// </summary>
        [Input("area")]
        public Input<string>? Area { get; set; }

        /// <summary>
        /// Layer 3/Layer 4 DDoS protection. The default protection option of the platform will be used if it is left empty. For
        /// details, see [Exclusive DDoS Protection
        /// Usage](https://intl.cloud.tencent.com/document/product/1552/95994?from_cn_redirect=1).
        /// </summary>
        [Input("ddosProtectionConfig")]
        public Input<Inputs.L4ProxyDdosProtectionConfigGetArgs>? DdosProtectionConfig { get; set; }

        /// <summary>
        /// Specifies whether to enable IPv6 access. The default value off is used if left empty. This configuration can only be
        /// enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy
        /// Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable;
        /// `off`: Disable.
        /// </summary>
        [Input("ipv6")]
        public Input<string>? Ipv6 { get; set; }

        /// <summary>
        /// Layer 4 proxy instance name. You can enter 1-50 characters. Valid characters are a-z, 0-9, and hyphens (-). However,
        /// hyphens (-) cannot be used individually or consecutively and should not be placed at the beginning or end of the name.
        /// Modifications are not allowed after creation.
        /// </summary>
        [Input("proxyName")]
        public Input<string>? ProxyName { get; set; }

        /// <summary>
        /// Specifies whether to enable the fixed IP address. The default value off is used if left empty. This configuration can
        /// only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4
        /// Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`:
        /// Enable; `off`: Disable.
        /// </summary>
        [Input("staticIp")]
        public Input<string>? StaticIp { get; set; }

        /// <summary>
        /// Site ID.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public L4ProxyState()
        {
        }
        public static new L4ProxyState Empty => new L4ProxyState();
    }
}