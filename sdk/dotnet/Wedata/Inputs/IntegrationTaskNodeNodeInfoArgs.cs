// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Wedata.Inputs
{

    public sealed class IntegrationTaskNodeNodeInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        [Input("configs")]
        private InputList<Inputs.IntegrationTaskNodeNodeInfoConfigArgs>? _configs;
        public InputList<Inputs.IntegrationTaskNodeNodeInfoConfigArgs> Configs
        {
            get => _configs ?? (_configs = new InputList<Inputs.IntegrationTaskNodeNodeInfoConfigArgs>());
            set => _configs = value;
        }

        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("creatorUin")]
        public Input<string>? CreatorUin { get; set; }

        [Input("datasourceId")]
        public Input<string>? DatasourceId { get; set; }

        [Input("extConfigs")]
        private InputList<Inputs.IntegrationTaskNodeNodeInfoExtConfigArgs>? _extConfigs;
        public InputList<Inputs.IntegrationTaskNodeNodeInfoExtConfigArgs> ExtConfigs
        {
            get => _extConfigs ?? (_extConfigs = new InputList<Inputs.IntegrationTaskNodeNodeInfoExtConfigArgs>());
            set => _extConfigs = value;
        }

        [Input("nodeMapping")]
        public Input<Inputs.IntegrationTaskNodeNodeInfoNodeMappingArgs>? NodeMapping { get; set; }

        [Input("operatorUin")]
        public Input<string>? OperatorUin { get; set; }

        [Input("ownerUin")]
        public Input<string>? OwnerUin { get; set; }

        [Input("schemas")]
        private InputList<Inputs.IntegrationTaskNodeNodeInfoSchemaArgs>? _schemas;
        public InputList<Inputs.IntegrationTaskNodeNodeInfoSchemaArgs> Schemas
        {
            get => _schemas ?? (_schemas = new InputList<Inputs.IntegrationTaskNodeNodeInfoSchemaArgs>());
            set => _schemas = value;
        }

        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public IntegrationTaskNodeNodeInfoArgs()
        {
        }
        public static new IntegrationTaskNodeNodeInfoArgs Empty => new IntegrationTaskNodeNodeInfoArgs();
    }
}
