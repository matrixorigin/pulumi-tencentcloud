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
    public sealed class GetSgSnapshotFileContentOriginalDataAddressTemplateResult
    {
        public readonly string AddressGroupId;
        public readonly string AddressId;

        [OutputConstructor]
        private GetSgSnapshotFileContentOriginalDataAddressTemplateResult(
            string addressGroupId,

            string addressId)
        {
            AddressGroupId = addressGroupId;
            AddressId = addressId;
        }
    }
}