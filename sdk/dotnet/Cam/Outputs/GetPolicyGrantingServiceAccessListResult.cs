// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cam.Outputs
{

    [OutputType]
    public sealed class GetPolicyGrantingServiceAccessListResult
    {
        public readonly ImmutableArray<Outputs.GetPolicyGrantingServiceAccessListActionResult> Actions;
        public readonly ImmutableArray<Outputs.GetPolicyGrantingServiceAccessListPolicyResult> Policies;
        public readonly ImmutableArray<Outputs.GetPolicyGrantingServiceAccessListServiceResult> Services;

        [OutputConstructor]
        private GetPolicyGrantingServiceAccessListResult(
            ImmutableArray<Outputs.GetPolicyGrantingServiceAccessListActionResult> actions,

            ImmutableArray<Outputs.GetPolicyGrantingServiceAccessListPolicyResult> policies,

            ImmutableArray<Outputs.GetPolicyGrantingServiceAccessListServiceResult> services)
        {
            Actions = actions;
            Policies = policies;
            Services = services;
        }
    }
}