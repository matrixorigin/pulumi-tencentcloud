// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor
{
    public static class GetAlarmBasicAlarms
    {
        public static Task<GetAlarmBasicAlarmsResult> InvokeAsync(GetAlarmBasicAlarmsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAlarmBasicAlarmsResult>("tencentcloud:Monitor/getAlarmBasicAlarms:getAlarmBasicAlarms", args ?? new GetAlarmBasicAlarmsArgs(), options.WithDefaults());

        public static Output<GetAlarmBasicAlarmsResult> Invoke(GetAlarmBasicAlarmsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlarmBasicAlarmsResult>("tencentcloud:Monitor/getAlarmBasicAlarms:getAlarmBasicAlarms", args ?? new GetAlarmBasicAlarmsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAlarmBasicAlarmsArgs : global::Pulumi.InvokeArgs
    {
        [Input("alarmStatuses")]
        private List<int>? _alarmStatuses;
        public List<int> AlarmStatuses
        {
            get => _alarmStatuses ?? (_alarmStatuses = new List<int>());
            set => _alarmStatuses = value;
        }

        [Input("endTime")]
        public int? EndTime { get; set; }

        [Input("instanceGroupIds")]
        private List<int>? _instanceGroupIds;
        public List<int> InstanceGroupIds
        {
            get => _instanceGroupIds ?? (_instanceGroupIds = new List<int>());
            set => _instanceGroupIds = value;
        }

        [Input("metricNames")]
        private List<string>? _metricNames;
        public List<string> MetricNames
        {
            get => _metricNames ?? (_metricNames = new List<string>());
            set => _metricNames = value;
        }

        [Input("module", required: true)]
        public string Module { get; set; } = null!;

        [Input("objLike")]
        public string? ObjLike { get; set; }

        [Input("occurTimeOrder")]
        public string? OccurTimeOrder { get; set; }

        [Input("projectIds")]
        private List<int>? _projectIds;
        public List<int> ProjectIds
        {
            get => _projectIds ?? (_projectIds = new List<int>());
            set => _projectIds = value;
        }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("startTime")]
        public int? StartTime { get; set; }

        [Input("viewNames")]
        private List<string>? _viewNames;
        public List<string> ViewNames
        {
            get => _viewNames ?? (_viewNames = new List<string>());
            set => _viewNames = value;
        }

        public GetAlarmBasicAlarmsArgs()
        {
        }
        public static new GetAlarmBasicAlarmsArgs Empty => new GetAlarmBasicAlarmsArgs();
    }

    public sealed class GetAlarmBasicAlarmsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("alarmStatuses")]
        private InputList<int>? _alarmStatuses;
        public InputList<int> AlarmStatuses
        {
            get => _alarmStatuses ?? (_alarmStatuses = new InputList<int>());
            set => _alarmStatuses = value;
        }

        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        [Input("instanceGroupIds")]
        private InputList<int>? _instanceGroupIds;
        public InputList<int> InstanceGroupIds
        {
            get => _instanceGroupIds ?? (_instanceGroupIds = new InputList<int>());
            set => _instanceGroupIds = value;
        }

        [Input("metricNames")]
        private InputList<string>? _metricNames;
        public InputList<string> MetricNames
        {
            get => _metricNames ?? (_metricNames = new InputList<string>());
            set => _metricNames = value;
        }

        [Input("module", required: true)]
        public Input<string> Module { get; set; } = null!;

        [Input("objLike")]
        public Input<string>? ObjLike { get; set; }

        [Input("occurTimeOrder")]
        public Input<string>? OccurTimeOrder { get; set; }

        [Input("projectIds")]
        private InputList<int>? _projectIds;
        public InputList<int> ProjectIds
        {
            get => _projectIds ?? (_projectIds = new InputList<int>());
            set => _projectIds = value;
        }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        [Input("viewNames")]
        private InputList<string>? _viewNames;
        public InputList<string> ViewNames
        {
            get => _viewNames ?? (_viewNames = new InputList<string>());
            set => _viewNames = value;
        }

        public GetAlarmBasicAlarmsInvokeArgs()
        {
        }
        public static new GetAlarmBasicAlarmsInvokeArgs Empty => new GetAlarmBasicAlarmsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAlarmBasicAlarmsResult
    {
        public readonly ImmutableArray<int> AlarmStatuses;
        public readonly ImmutableArray<Outputs.GetAlarmBasicAlarmsAlarmResult> Alarms;
        public readonly int? EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<int> InstanceGroupIds;
        public readonly ImmutableArray<string> MetricNames;
        public readonly string Module;
        public readonly string? ObjLike;
        public readonly string? OccurTimeOrder;
        public readonly ImmutableArray<int> ProjectIds;
        public readonly string? ResultOutputFile;
        public readonly int? StartTime;
        public readonly ImmutableArray<string> ViewNames;
        public readonly string Warning;

        [OutputConstructor]
        private GetAlarmBasicAlarmsResult(
            ImmutableArray<int> alarmStatuses,

            ImmutableArray<Outputs.GetAlarmBasicAlarmsAlarmResult> alarms,

            int? endTime,

            string id,

            ImmutableArray<int> instanceGroupIds,

            ImmutableArray<string> metricNames,

            string module,

            string? objLike,

            string? occurTimeOrder,

            ImmutableArray<int> projectIds,

            string? resultOutputFile,

            int? startTime,

            ImmutableArray<string> viewNames,

            string warning)
        {
            AlarmStatuses = alarmStatuses;
            Alarms = alarms;
            EndTime = endTime;
            Id = id;
            InstanceGroupIds = instanceGroupIds;
            MetricNames = metricNames;
            Module = module;
            ObjLike = objLike;
            OccurTimeOrder = occurTimeOrder;
            ProjectIds = projectIds;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
            ViewNames = viewNames;
            Warning = warning;
        }
    }
}