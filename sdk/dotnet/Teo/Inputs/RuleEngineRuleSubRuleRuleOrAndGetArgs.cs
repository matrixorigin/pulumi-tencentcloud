// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Teo.Inputs
{

    public sealed class RuleEngineRuleSubRuleRuleOrAndGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("ignoreCase")]
        public Input<bool>? IgnoreCase { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        [Input("values")]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public RuleEngineRuleSubRuleRuleOrAndGetArgs()
        {
        }
        public static new RuleEngineRuleSubRuleRuleOrAndGetArgs Empty => new RuleEngineRuleSubRuleRuleOrAndGetArgs();
    }
}