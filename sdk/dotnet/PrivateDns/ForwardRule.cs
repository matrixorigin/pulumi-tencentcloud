// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.PrivateDns
{
    [TencentcloudResourceType("tencentcloud:PrivateDns/forwardRule:ForwardRule")]
    public partial class ForwardRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Endpoint ID.
        /// </summary>
        [Output("endPointId")]
        public Output<string> EndPointId { get; private set; } = null!;

        /// <summary>
        /// Forwarding rule name.
        /// </summary>
        [Output("ruleName")]
        public Output<string> RuleName { get; private set; } = null!;

        /// <summary>
        /// Forwarding rule type. DOWN: From cloud to off-cloud; UP: From off-cloud to cloud.
        /// </summary>
        [Output("ruleType")]
        public Output<string> RuleType { get; private set; } = null!;

        /// <summary>
        /// Private domain ID, which can be viewed on the private domain list page.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a ForwardRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ForwardRule(string name, ForwardRuleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:PrivateDns/forwardRule:ForwardRule", name, args ?? new ForwardRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ForwardRule(string name, Input<string> id, ForwardRuleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:PrivateDns/forwardRule:ForwardRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ForwardRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ForwardRule Get(string name, Input<string> id, ForwardRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ForwardRule(name, id, state, options);
        }
    }

    public sealed class ForwardRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Endpoint ID.
        /// </summary>
        [Input("endPointId", required: true)]
        public Input<string> EndPointId { get; set; } = null!;

        /// <summary>
        /// Forwarding rule name.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        /// <summary>
        /// Forwarding rule type. DOWN: From cloud to off-cloud; UP: From off-cloud to cloud.
        /// </summary>
        [Input("ruleType", required: true)]
        public Input<string> RuleType { get; set; } = null!;

        /// <summary>
        /// Private domain ID, which can be viewed on the private domain list page.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public ForwardRuleArgs()
        {
        }
        public static new ForwardRuleArgs Empty => new ForwardRuleArgs();
    }

    public sealed class ForwardRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Endpoint ID.
        /// </summary>
        [Input("endPointId")]
        public Input<string>? EndPointId { get; set; }

        /// <summary>
        /// Forwarding rule name.
        /// </summary>
        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        /// <summary>
        /// Forwarding rule type. DOWN: From cloud to off-cloud; UP: From off-cloud to cloud.
        /// </summary>
        [Input("ruleType")]
        public Input<string>? RuleType { get; set; }

        /// <summary>
        /// Private domain ID, which can be viewed on the private domain list page.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ForwardRuleState()
        {
        }
        public static new ForwardRuleState Empty => new ForwardRuleState();
    }
}
