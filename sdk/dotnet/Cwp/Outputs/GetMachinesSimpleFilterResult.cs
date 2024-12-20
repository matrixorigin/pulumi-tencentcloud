// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cwp.Outputs
{

    [OutputType]
    public sealed class GetMachinesSimpleFilterResult
    {
        public readonly bool? ExactMatch;
        public readonly string Name;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetMachinesSimpleFilterResult(
            bool? exactMatch,

            string name,

            ImmutableArray<string> values)
        {
            ExactMatch = exactMatch;
            Name = name;
            Values = values;
        }
    }
}
