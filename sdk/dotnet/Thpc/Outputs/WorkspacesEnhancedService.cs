// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Thpc.Outputs
{

    [OutputType]
    public sealed class WorkspacesEnhancedService
    {
        public readonly Outputs.WorkspacesEnhancedServiceAutomationService? AutomationService;
        public readonly Outputs.WorkspacesEnhancedServiceMonitorService? MonitorService;
        public readonly Outputs.WorkspacesEnhancedServiceSecurityService? SecurityService;

        [OutputConstructor]
        private WorkspacesEnhancedService(
            Outputs.WorkspacesEnhancedServiceAutomationService? automationService,

            Outputs.WorkspacesEnhancedServiceMonitorService? monitorService,

            Outputs.WorkspacesEnhancedServiceSecurityService? securityService)
        {
            AutomationService = automationService;
            MonitorService = monitorService;
            SecurityService = securityService;
        }
    }
}