// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class GetSchedulesScheduleInfoSetOutputStorageResult
    {
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetOutputStorageCosOutputStorageResult> CosOutputStorages;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetOutputStorageS3OutputStorageResult> S3OutputStorages;
        public readonly string Type;

        [OutputConstructor]
        private GetSchedulesScheduleInfoSetOutputStorageResult(
            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetOutputStorageCosOutputStorageResult> cosOutputStorages,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetOutputStorageS3OutputStorageResult> s3OutputStorages,

            string type)
        {
            CosOutputStorages = cosOutputStorages;
            S3OutputStorages = s3OutputStorages;
            Type = type;
        }
    }
}