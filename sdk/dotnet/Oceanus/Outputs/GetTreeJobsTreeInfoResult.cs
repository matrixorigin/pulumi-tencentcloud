// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Oceanus.Outputs
{

    [OutputType]
    public sealed class GetTreeJobsTreeInfoResult
    {
        public readonly string Children;
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetTreeJobsTreeInfoJobSetResult> JobSets;
        public readonly string Name;
        public readonly string ParentId;

        [OutputConstructor]
        private GetTreeJobsTreeInfoResult(
            string children,

            string id,

            ImmutableArray<Outputs.GetTreeJobsTreeInfoJobSetResult> jobSets,

            string name,

            string parentId)
        {
            Children = children;
            Id = id;
            JobSets = jobSets;
            Name = name;
            ParentId = parentId;
        }
    }
}
