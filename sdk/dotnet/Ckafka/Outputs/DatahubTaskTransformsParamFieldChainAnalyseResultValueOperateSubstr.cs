// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateSubstr
    {
        public readonly int End;
        public readonly int Start;

        [OutputConstructor]
        private DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateSubstr(
            int end,

            int start)
        {
            End = end;
            Start = start;
        }
    }
}