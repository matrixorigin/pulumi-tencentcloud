// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Events.Inputs
{

    public sealed class AuditTrackFiltersResourceFieldArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionType", required: true)]
        public Input<string> ActionType { get; set; } = null!;

        [Input("eventNames", required: true)]
        private InputList<string>? _eventNames;
        public InputList<string> EventNames
        {
            get => _eventNames ?? (_eventNames = new InputList<string>());
            set => _eventNames = value;
        }

        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        public AuditTrackFiltersResourceFieldArgs()
        {
        }
        public static new AuditTrackFiltersResourceFieldArgs Empty => new AuditTrackFiltersResourceFieldArgs();
    }
}