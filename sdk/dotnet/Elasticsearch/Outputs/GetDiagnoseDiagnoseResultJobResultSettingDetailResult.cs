// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elasticsearch.Outputs
{

    [OutputType]
    public sealed class GetDiagnoseDiagnoseResultJobResultSettingDetailResult
    {
        public readonly string Advise;
        public readonly string Key;
        public readonly string Value;

        [OutputConstructor]
        private GetDiagnoseDiagnoseResultJobResultSettingDetailResult(
            string advise,

            string key,

            string value)
        {
            Advise = advise;
            Key = key;
            Value = value;
        }
    }
}
