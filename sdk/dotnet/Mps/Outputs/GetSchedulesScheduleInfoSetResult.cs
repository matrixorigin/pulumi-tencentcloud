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
    public sealed class GetSchedulesScheduleInfoSetResult
    {
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityResult> Activities;
        public readonly string CreateTime;
        public readonly string OutputDir;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetOutputStorageResult> OutputStorages;
        public readonly string ResourceId;
        public readonly int ScheduleId;
        public readonly string ScheduleName;
        public readonly string Status;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetTaskNotifyConfigResult> TaskNotifyConfigs;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetTriggerResult> Triggers;
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetSchedulesScheduleInfoSetResult(
            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityResult> activities,

            string createTime,

            string outputDir,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetOutputStorageResult> outputStorages,

            string resourceId,

            int scheduleId,

            string scheduleName,

            string status,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetTaskNotifyConfigResult> taskNotifyConfigs,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetTriggerResult> triggers,

            string updateTime)
        {
            Activities = activities;
            CreateTime = createTime;
            OutputDir = outputDir;
            OutputStorages = outputStorages;
            ResourceId = resourceId;
            ScheduleId = scheduleId;
            ScheduleName = scheduleName;
            Status = status;
            TaskNotifyConfigs = taskNotifyConfigs;
            Triggers = triggers;
            UpdateTime = updateTime;
        }
    }
}