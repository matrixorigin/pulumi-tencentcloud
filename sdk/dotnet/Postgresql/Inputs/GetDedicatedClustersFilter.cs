// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Postgresql.Inputs
{

    public sealed class GetDedicatedClustersFilterArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        [Input("values")]
        private List<string>? _values;
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetDedicatedClustersFilterArgs()
        {
        }
        public static new GetDedicatedClustersFilterArgs Empty => new GetDedicatedClustersFilterArgs();
    }
}
