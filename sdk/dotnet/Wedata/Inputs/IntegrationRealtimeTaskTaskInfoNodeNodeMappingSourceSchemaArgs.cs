// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Wedata.Inputs
{

    public sealed class IntegrationRealtimeTaskTaskInfoNodeNodeMappingSourceSchemaArgs : global::Pulumi.ResourceArgs
    {
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("properties")]
        private InputList<Inputs.IntegrationRealtimeTaskTaskInfoNodeNodeMappingSourceSchemaPropertyArgs>? _properties;
        public InputList<Inputs.IntegrationRealtimeTaskTaskInfoNodeNodeMappingSourceSchemaPropertyArgs> Properties
        {
            get => _properties ?? (_properties = new InputList<Inputs.IntegrationRealtimeTaskTaskInfoNodeNodeMappingSourceSchemaPropertyArgs>());
            set => _properties = value;
        }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("value")]
        public Input<string>? Value { get; set; }

        public IntegrationRealtimeTaskTaskInfoNodeNodeMappingSourceSchemaArgs()
        {
        }
        public static new IntegrationRealtimeTaskTaskInfoNodeNodeMappingSourceSchemaArgs Empty => new IntegrationRealtimeTaskTaskInfoNodeNodeMappingSourceSchemaArgs();
    }
}