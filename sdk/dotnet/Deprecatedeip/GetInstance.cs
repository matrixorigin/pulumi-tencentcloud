// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Deprecatedeip
{
    public static class GetInstance
    {
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("tencentcloud:Deprecatedeip/getInstance:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("tencentcloud:Deprecatedeip/getInstance:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetInstanceFilterArgs>? _filters;
        public List<Inputs.GetInstanceFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetInstanceFilterArgs>());
            set => _filters = value;
        }

        [Input("includeArrears")]
        public bool? IncludeArrears { get; set; }

        [Input("includeBlocked")]
        public bool? IncludeBlocked { get; set; }

        public GetInstanceArgs()
        {
        }
        public static new GetInstanceArgs Empty => new GetInstanceArgs();
    }

    public sealed class GetInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetInstanceFilterInputArgs>? _filters;
        public InputList<Inputs.GetInstanceFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetInstanceFilterInputArgs>());
            set => _filters = value;
        }

        [Input("includeArrears")]
        public Input<bool>? IncludeArrears { get; set; }

        [Input("includeBlocked")]
        public Input<bool>? IncludeBlocked { get; set; }

        public GetInstanceInvokeArgs()
        {
        }
        public static new GetInstanceInvokeArgs Empty => new GetInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        public readonly ImmutableArray<Outputs.GetInstanceFilterResult> Filters;
        public readonly string Id;
        public readonly bool? IncludeArrears;
        public readonly bool? IncludeBlocked;
        public readonly string PublicIp;
        public readonly string Status;

        [OutputConstructor]
        private GetInstanceResult(
            ImmutableArray<Outputs.GetInstanceFilterResult> filters,

            string id,

            bool? includeArrears,

            bool? includeBlocked,

            string publicIp,

            string status)
        {
            Filters = filters;
            Id = id;
            IncludeArrears = includeArrears;
            IncludeBlocked = includeBlocked;
            PublicIp = publicIp;
            Status = status;
        }
    }
}