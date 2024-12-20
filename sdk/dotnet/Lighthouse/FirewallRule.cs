// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Lighthouse
{
    [TencentcloudResourceType("tencentcloud:Lighthouse/firewallRule:FirewallRule")]
    public partial class FirewallRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Firewall rule list.
        /// </summary>
        [Output("firewallRules")]
        public Output<ImmutableArray<Outputs.FirewallRuleFirewallRule>> FirewallRules { get; private set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallRule(string name, FirewallRuleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Lighthouse/firewallRule:FirewallRule", name, args ?? new FirewallRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallRule(string name, Input<string> id, FirewallRuleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Lighthouse/firewallRule:FirewallRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallRule Get(string name, Input<string> id, FirewallRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallRule(name, id, state, options);
        }
    }

    public sealed class FirewallRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("firewallRules", required: true)]
        private InputList<Inputs.FirewallRuleFirewallRuleArgs>? _firewallRules;

        /// <summary>
        /// Firewall rule list.
        /// </summary>
        public InputList<Inputs.FirewallRuleFirewallRuleArgs> FirewallRules
        {
            get => _firewallRules ?? (_firewallRules = new InputList<Inputs.FirewallRuleFirewallRuleArgs>());
            set => _firewallRules = value;
        }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public FirewallRuleArgs()
        {
        }
        public static new FirewallRuleArgs Empty => new FirewallRuleArgs();
    }

    public sealed class FirewallRuleState : global::Pulumi.ResourceArgs
    {
        [Input("firewallRules")]
        private InputList<Inputs.FirewallRuleFirewallRuleGetArgs>? _firewallRules;

        /// <summary>
        /// Firewall rule list.
        /// </summary>
        public InputList<Inputs.FirewallRuleFirewallRuleGetArgs> FirewallRules
        {
            get => _firewallRules ?? (_firewallRules = new InputList<Inputs.FirewallRuleFirewallRuleGetArgs>());
            set => _firewallRules = value;
        }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public FirewallRuleState()
        {
        }
        public static new FirewallRuleState Empty => new FirewallRuleState();
    }
}
