// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssl.Outputs
{

    [OutputType]
    public sealed class GetDescribeCertificateResultDvRevokeAuthDetailResult
    {
        public readonly string DvAuthDomain;
        public readonly string DvAuthKey;
        public readonly string DvAuthPath;
        public readonly string DvAuthSubDomain;
        public readonly string DvAuthValue;
        public readonly string DvAuthVerifyType;

        [OutputConstructor]
        private GetDescribeCertificateResultDvRevokeAuthDetailResult(
            string dvAuthDomain,

            string dvAuthKey,

            string dvAuthPath,

            string dvAuthSubDomain,

            string dvAuthValue,

            string dvAuthVerifyType)
        {
            DvAuthDomain = dvAuthDomain;
            DvAuthKey = dvAuthKey;
            DvAuthPath = dvAuthPath;
            DvAuthSubDomain = dvAuthSubDomain;
            DvAuthValue = dvAuthValue;
            DvAuthVerifyType = dvAuthVerifyType;
        }
    }
}