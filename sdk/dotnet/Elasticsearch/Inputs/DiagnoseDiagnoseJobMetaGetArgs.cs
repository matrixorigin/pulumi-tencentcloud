// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elasticsearch.Inputs
{

    public sealed class DiagnoseDiagnoseJobMetaGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("jobDescription")]
        public Input<string>? JobDescription { get; set; }

        [Input("jobName")]
        public Input<string>? JobName { get; set; }

        [Input("jobZhName")]
        public Input<string>? JobZhName { get; set; }

        public DiagnoseDiagnoseJobMetaGetArgs()
        {
        }
        public static new DiagnoseDiagnoseJobMetaGetArgs Empty => new DiagnoseDiagnoseJobMetaGetArgs();
    }
}
