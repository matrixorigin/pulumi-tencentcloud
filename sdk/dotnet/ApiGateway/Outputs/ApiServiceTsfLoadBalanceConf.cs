// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.ApiGateway.Outputs
{

    [OutputType]
    public sealed class ApiServiceTsfLoadBalanceConf
    {
        public readonly bool? IsLoadBalance;
        public readonly string? Method;
        public readonly bool? SessionStickRequired;
        public readonly int? SessionStickTimeout;

        [OutputConstructor]
        private ApiServiceTsfLoadBalanceConf(
            bool? isLoadBalance,

            string? method,

            bool? sessionStickRequired,

            int? sessionStickTimeout)
        {
            IsLoadBalance = isLoadBalance;
            Method = method;
            SessionStickRequired = sessionStickRequired;
            SessionStickTimeout = sessionStickTimeout;
        }
    }
}