// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Oceanus
{
    public static class GetWorkSpaces
    {
        public static Task<GetWorkSpacesResult> InvokeAsync(GetWorkSpacesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkSpacesResult>("tencentcloud:Oceanus/getWorkSpaces:getWorkSpaces", args ?? new GetWorkSpacesArgs(), options.WithDefaults());

        public static Output<GetWorkSpacesResult> Invoke(GetWorkSpacesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkSpacesResult>("tencentcloud:Oceanus/getWorkSpaces:getWorkSpaces", args ?? new GetWorkSpacesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkSpacesArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetWorkSpacesFilterArgs>? _filters;
        public List<Inputs.GetWorkSpacesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetWorkSpacesFilterArgs>());
            set => _filters = value;
        }

        [Input("orderType")]
        public int? OrderType { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetWorkSpacesArgs()
        {
        }
        public static new GetWorkSpacesArgs Empty => new GetWorkSpacesArgs();
    }

    public sealed class GetWorkSpacesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetWorkSpacesFilterInputArgs>? _filters;
        public InputList<Inputs.GetWorkSpacesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetWorkSpacesFilterInputArgs>());
            set => _filters = value;
        }

        [Input("orderType")]
        public Input<int>? OrderType { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetWorkSpacesInvokeArgs()
        {
        }
        public static new GetWorkSpacesInvokeArgs Empty => new GetWorkSpacesInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkSpacesResult
    {
        public readonly ImmutableArray<Outputs.GetWorkSpacesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? OrderType;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<Outputs.GetWorkSpacesWorkSpaceSetItemResult> WorkSpaceSetItems;

        [OutputConstructor]
        private GetWorkSpacesResult(
            ImmutableArray<Outputs.GetWorkSpacesFilterResult> filters,

            string id,

            int? orderType,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetWorkSpacesWorkSpaceSetItemResult> workSpaceSetItems)
        {
            Filters = filters;
            Id = id;
            OrderType = orderType;
            ResultOutputFile = resultOutputFile;
            WorkSpaceSetItems = workSpaceSetItems;
        }
    }
}
