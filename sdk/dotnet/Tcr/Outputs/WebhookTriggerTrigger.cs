// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcr.Outputs
{

    [OutputType]
    public sealed class WebhookTriggerTrigger
    {
        public readonly string Condition;
        public readonly string? Description;
        public readonly bool Enabled;
        public readonly ImmutableArray<string> EventTypes;
        public readonly int? Id;
        public readonly string Name;
        public readonly int? NamespaceId;
        public readonly ImmutableArray<Outputs.WebhookTriggerTriggerTarget> Targets;

        [OutputConstructor]
        private WebhookTriggerTrigger(
            string condition,

            string? description,

            bool enabled,

            ImmutableArray<string> eventTypes,

            int? id,

            string name,

            int? namespaceId,

            ImmutableArray<Outputs.WebhookTriggerTriggerTarget> targets)
        {
            Condition = condition;
            Description = description;
            Enabled = enabled;
            EventTypes = eventTypes;
            Id = id;
            Name = name;
            NamespaceId = namespaceId;
            Targets = targets;
        }
    }
}