// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class AiRecognitionTemplateAsrWordsConfigure
    {
        public readonly ImmutableArray<string> LabelSets;
        public readonly string Switch;

        [OutputConstructor]
        private AiRecognitionTemplateAsrWordsConfigure(
            ImmutableArray<string> labelSets,

            string @switch)
        {
            LabelSets = labelSets;
            Switch = @switch;
        }
    }
}