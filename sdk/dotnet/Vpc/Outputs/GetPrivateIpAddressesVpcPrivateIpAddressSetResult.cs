// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vpc.Outputs
{

    [OutputType]
    public sealed class GetPrivateIpAddressesVpcPrivateIpAddressSetResult
    {
        public readonly string CidrBlock;
        public readonly string CreatedTime;
        public readonly string PrivateIpAddress;
        public readonly string PrivateIpAddressType;

        [OutputConstructor]
        private GetPrivateIpAddressesVpcPrivateIpAddressSetResult(
            string cidrBlock,

            string createdTime,

            string privateIpAddress,

            string privateIpAddressType)
        {
            CidrBlock = cidrBlock;
            CreatedTime = createdTime;
            PrivateIpAddress = privateIpAddress;
            PrivateIpAddressType = privateIpAddressType;
        }
    }
}