// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Pts
{
    public static class GetScenarioWithJobs
    {
        public static Task<GetScenarioWithJobsResult> InvokeAsync(GetScenarioWithJobsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetScenarioWithJobsResult>("tencentcloud:Pts/getScenarioWithJobs:getScenarioWithJobs", args ?? new GetScenarioWithJobsArgs(), options.WithDefaults());

        public static Output<GetScenarioWithJobsResult> Invoke(GetScenarioWithJobsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetScenarioWithJobsResult>("tencentcloud:Pts/getScenarioWithJobs:getScenarioWithJobs", args ?? new GetScenarioWithJobsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetScenarioWithJobsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ascend")]
        public bool? Ascend { get; set; }

        [Input("ignoreDataset")]
        public bool? IgnoreDataset { get; set; }

        [Input("ignoreScript")]
        public bool? IgnoreScript { get; set; }

        [Input("orderBy")]
        public string? OrderBy { get; set; }

        [Input("owner")]
        public string? Owner { get; set; }

        [Input("projectIds", required: true)]
        private List<string>? _projectIds;
        public List<string> ProjectIds
        {
            get => _projectIds ?? (_projectIds = new List<string>());
            set => _projectIds = value;
        }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("scenarioIds")]
        private List<string>? _scenarioIds;
        public List<string> ScenarioIds
        {
            get => _scenarioIds ?? (_scenarioIds = new List<string>());
            set => _scenarioIds = value;
        }

        [Input("scenarioName")]
        public string? ScenarioName { get; set; }

        [Input("scenarioType")]
        public string? ScenarioType { get; set; }

        public GetScenarioWithJobsArgs()
        {
        }
        public static new GetScenarioWithJobsArgs Empty => new GetScenarioWithJobsArgs();
    }

    public sealed class GetScenarioWithJobsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ascend")]
        public Input<bool>? Ascend { get; set; }

        [Input("ignoreDataset")]
        public Input<bool>? IgnoreDataset { get; set; }

        [Input("ignoreScript")]
        public Input<bool>? IgnoreScript { get; set; }

        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("projectIds", required: true)]
        private InputList<string>? _projectIds;
        public InputList<string> ProjectIds
        {
            get => _projectIds ?? (_projectIds = new InputList<string>());
            set => _projectIds = value;
        }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("scenarioIds")]
        private InputList<string>? _scenarioIds;
        public InputList<string> ScenarioIds
        {
            get => _scenarioIds ?? (_scenarioIds = new InputList<string>());
            set => _scenarioIds = value;
        }

        [Input("scenarioName")]
        public Input<string>? ScenarioName { get; set; }

        [Input("scenarioType")]
        public Input<string>? ScenarioType { get; set; }

        public GetScenarioWithJobsInvokeArgs()
        {
        }
        public static new GetScenarioWithJobsInvokeArgs Empty => new GetScenarioWithJobsInvokeArgs();
    }


    [OutputType]
    public sealed class GetScenarioWithJobsResult
    {
        public readonly bool? Ascend;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IgnoreDataset;
        public readonly bool? IgnoreScript;
        public readonly string? OrderBy;
        public readonly string? Owner;
        public readonly ImmutableArray<string> ProjectIds;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<string> ScenarioIds;
        public readonly string? ScenarioName;
        public readonly string? ScenarioType;
        public readonly ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetResult> ScenarioWithJobsSets;

        [OutputConstructor]
        private GetScenarioWithJobsResult(
            bool? ascend,

            string id,

            bool? ignoreDataset,

            bool? ignoreScript,

            string? orderBy,

            string? owner,

            ImmutableArray<string> projectIds,

            string? resultOutputFile,

            ImmutableArray<string> scenarioIds,

            string? scenarioName,

            string? scenarioType,

            ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetResult> scenarioWithJobsSets)
        {
            Ascend = ascend;
            Id = id;
            IgnoreDataset = ignoreDataset;
            IgnoreScript = ignoreScript;
            OrderBy = orderBy;
            Owner = owner;
            ProjectIds = projectIds;
            ResultOutputFile = resultOutputFile;
            ScenarioIds = scenarioIds;
            ScenarioName = scenarioName;
            ScenarioType = scenarioType;
            ScenarioWithJobsSets = scenarioWithJobsSets;
        }
    }
}