// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dbbrain.Outputs
{

    [OutputType]
    public sealed class GetTopSpaceSchemaTimeSeriesTopSpaceSchemaTimeSeriesResult
    {
        public readonly ImmutableArray<Outputs.GetTopSpaceSchemaTimeSeriesTopSpaceSchemaTimeSeriesSeriesDataResult> SeriesDatas;
        public readonly string TableSchema;

        [OutputConstructor]
        private GetTopSpaceSchemaTimeSeriesTopSpaceSchemaTimeSeriesResult(
            ImmutableArray<Outputs.GetTopSpaceSchemaTimeSeriesTopSpaceSchemaTimeSeriesSeriesDataResult> seriesDatas,

            string tableSchema)
        {
            SeriesDatas = seriesDatas;
            TableSchema = tableSchema;
        }
    }
}