// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cls.Outputs
{

    [OutputType]
    public sealed class TopicExtendsAnonymousAccess
    {
        public readonly ImmutableArray<Outputs.TopicExtendsAnonymousAccessCondition> Conditions;
        public readonly ImmutableArray<string> Operations;

        [OutputConstructor]
        private TopicExtendsAnonymousAccess(
            ImmutableArray<Outputs.TopicExtendsAnonymousAccessCondition> conditions,

            ImmutableArray<string> operations)
        {
            Conditions = conditions;
            Operations = operations;
        }
    }
}
