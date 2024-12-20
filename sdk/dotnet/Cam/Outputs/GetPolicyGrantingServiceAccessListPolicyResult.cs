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
    public sealed class GetPolicyGrantingServiceAccessListPolicyResult
    {
        public readonly string PolicyDescription;
        public readonly string PolicyId;
        public readonly string PolicyName;
        public readonly string PolicyType;

        [OutputConstructor]
        private GetPolicyGrantingServiceAccessListPolicyResult(
            string policyDescription,

            string policyId,

            string policyName,

            string policyType)
        {
            PolicyDescription = policyDescription;
            PolicyId = policyId;
            PolicyName = policyName;
            PolicyType = policyType;
        }
    }
}
