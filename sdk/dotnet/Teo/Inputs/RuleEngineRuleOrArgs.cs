// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Teo.Inputs
{

    public sealed class RuleEngineRuleOrArgs : global::Pulumi.ResourceArgs
    {
        [Input("ands", required: true)]
        private InputList<Inputs.RuleEngineRuleOrAndArgs>? _ands;
        public InputList<Inputs.RuleEngineRuleOrAndArgs> Ands
        {
            get => _ands ?? (_ands = new InputList<Inputs.RuleEngineRuleOrAndArgs>());
            set => _ands = value;
        }

        public RuleEngineRuleOrArgs()
        {
        }
        public static new RuleEngineRuleOrArgs Empty => new RuleEngineRuleOrArgs();
    }
}