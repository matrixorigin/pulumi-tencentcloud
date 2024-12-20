// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class GetAlarmAllNamespacesCommonNamespaceDimensionResult
    {
        public readonly bool CanFilterHistory;
        public readonly bool CanFilterPolicy;
        public readonly bool CanGroupBy;
        public readonly bool IsMultiple;
        public readonly bool IsMutable;
        public readonly bool IsRequired;
        public readonly bool IsVisible;
        public readonly string Key;
        public readonly bool MustGroupBy;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetAlarmAllNamespacesCommonNamespaceDimensionOperatorResult> Operators;
        public readonly string ShowValueReplace;

        [OutputConstructor]
        private GetAlarmAllNamespacesCommonNamespaceDimensionResult(
            bool canFilterHistory,

            bool canFilterPolicy,

            bool canGroupBy,

            bool isMultiple,

            bool isMutable,

            bool isRequired,

            bool isVisible,

            string key,

            bool mustGroupBy,

            string name,

            ImmutableArray<Outputs.GetAlarmAllNamespacesCommonNamespaceDimensionOperatorResult> operators,

            string showValueReplace)
        {
            CanFilterHistory = canFilterHistory;
            CanFilterPolicy = canFilterPolicy;
            CanGroupBy = canGroupBy;
            IsMultiple = isMultiple;
            IsMutable = isMutable;
            IsRequired = isRequired;
            IsVisible = isVisible;
            Key = key;
            MustGroupBy = mustGroupBy;
            Name = name;
            Operators = operators;
            ShowValueReplace = showValueReplace;
        }
    }
}
