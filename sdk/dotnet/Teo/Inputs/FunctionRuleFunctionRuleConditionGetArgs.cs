// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Teo.Inputs
{

    public sealed class FunctionRuleFunctionRuleConditionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("ruleConditions", required: true)]
        private InputList<Inputs.FunctionRuleFunctionRuleConditionRuleConditionGetArgs>? _ruleConditions;
        public InputList<Inputs.FunctionRuleFunctionRuleConditionRuleConditionGetArgs> RuleConditions
        {
            get => _ruleConditions ?? (_ruleConditions = new InputList<Inputs.FunctionRuleFunctionRuleConditionRuleConditionGetArgs>());
            set => _ruleConditions = value;
        }

        public FunctionRuleFunctionRuleConditionGetArgs()
        {
        }
        public static new FunctionRuleFunctionRuleConditionGetArgs Empty => new FunctionRuleFunctionRuleConditionGetArgs();
    }
}
