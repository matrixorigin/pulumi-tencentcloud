// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf.Inputs
{

    public sealed class FunctionAliasRoutingConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalVersionMatches")]
        private InputList<Inputs.FunctionAliasRoutingConfigAdditionalVersionMatchGetArgs>? _additionalVersionMatches;
        public InputList<Inputs.FunctionAliasRoutingConfigAdditionalVersionMatchGetArgs> AdditionalVersionMatches
        {
            get => _additionalVersionMatches ?? (_additionalVersionMatches = new InputList<Inputs.FunctionAliasRoutingConfigAdditionalVersionMatchGetArgs>());
            set => _additionalVersionMatches = value;
        }

        [Input("additionalVersionWeights")]
        private InputList<Inputs.FunctionAliasRoutingConfigAdditionalVersionWeightGetArgs>? _additionalVersionWeights;
        public InputList<Inputs.FunctionAliasRoutingConfigAdditionalVersionWeightGetArgs> AdditionalVersionWeights
        {
            get => _additionalVersionWeights ?? (_additionalVersionWeights = new InputList<Inputs.FunctionAliasRoutingConfigAdditionalVersionWeightGetArgs>());
            set => _additionalVersionWeights = value;
        }

        public FunctionAliasRoutingConfigGetArgs()
        {
        }
        public static new FunctionAliasRoutingConfigGetArgs Empty => new FunctionAliasRoutingConfigGetArgs();
    }
}