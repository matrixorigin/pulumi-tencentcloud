// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elasticsearch.Outputs
{

    [OutputType]
    public sealed class GetDescribeIndexListIndexMetaFieldIndexPolicyFieldResult
    {
        public readonly string ColdAction;
        public readonly string ColdEnable;
        public readonly string ColdMinAge;
        public readonly string FrozenEnable;
        public readonly string FrozenMinAge;
        public readonly string WarmEnable;
        public readonly string WarmMinAge;

        [OutputConstructor]
        private GetDescribeIndexListIndexMetaFieldIndexPolicyFieldResult(
            string coldAction,

            string coldEnable,

            string coldMinAge,

            string frozenEnable,

            string frozenMinAge,

            string warmEnable,

            string warmMinAge)
        {
            ColdAction = coldAction;
            ColdEnable = coldEnable;
            ColdMinAge = coldMinAge;
            FrozenEnable = frozenEnable;
            FrozenMinAge = frozenMinAge;
            WarmEnable = warmEnable;
            WarmMinAge = warmMinAge;
        }
    }
}