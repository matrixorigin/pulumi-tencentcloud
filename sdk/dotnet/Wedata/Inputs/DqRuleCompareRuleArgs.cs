// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Wedata.Inputs
{

    public sealed class DqRuleCompareRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("cycleStep")]
        public Input<int>? CycleStep { get; set; }

        [Input("items")]
        private InputList<Inputs.DqRuleCompareRuleItemArgs>? _items;
        public InputList<Inputs.DqRuleCompareRuleItemArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.DqRuleCompareRuleItemArgs>());
            set => _items = value;
        }

        public DqRuleCompareRuleArgs()
        {
        }
        public static new DqRuleCompareRuleArgs Empty => new DqRuleCompareRuleArgs();
    }
}
