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
    public sealed class GetDescribeManagersManagerResult
    {
        public readonly int CertCount;
        public readonly string CreateTime;
        public readonly int DomainCount;
        public readonly string ExpireTime;
        public readonly string ManagerDepartment;
        public readonly string ManagerFirstName;
        public readonly int ManagerId;
        public readonly string ManagerLastName;
        public readonly string ManagerMail;
        public readonly string ManagerPhone;
        public readonly string ManagerPosition;
        public readonly string Status;
        public readonly string SubmitAuditTime;
        public readonly string VerifyTime;

        [OutputConstructor]
        private GetDescribeManagersManagerResult(
            int certCount,

            string createTime,

            int domainCount,

            string expireTime,

            string managerDepartment,

            string managerFirstName,

            int managerId,

            string managerLastName,

            string managerMail,

            string managerPhone,

            string managerPosition,

            string status,

            string submitAuditTime,

            string verifyTime)
        {
            CertCount = certCount;
            CreateTime = createTime;
            DomainCount = domainCount;
            ExpireTime = expireTime;
            ManagerDepartment = managerDepartment;
            ManagerFirstName = managerFirstName;
            ManagerId = managerId;
            ManagerLastName = managerLastName;
            ManagerMail = managerMail;
            ManagerPhone = managerPhone;
            ManagerPosition = managerPosition;
            Status = status;
            SubmitAuditTime = submitAuditTime;
            VerifyTime = verifyTime;
        }
    }
}
