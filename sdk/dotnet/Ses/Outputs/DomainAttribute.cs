// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ses.Outputs
{

    [OutputType]
    public sealed class DomainAttribute
    {
        public readonly string? ExpectedValue;
        public readonly string? SendDomain;
        public readonly string? Type;

        [OutputConstructor]
        private DomainAttribute(
            string? expectedValue,

            string? sendDomain,

            string? type)
        {
            ExpectedValue = expectedValue;
            SendDomain = sendDomain;
            Type = type;
        }
    }
}
