// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dts.Inputs
{

    public sealed class MigrateJobMigrateOptionConsistencyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        public MigrateJobMigrateOptionConsistencyGetArgs()
        {
        }
        public static new MigrateJobMigrateOptionConsistencyGetArgs Empty => new MigrateJobMigrateOptionConsistencyGetArgs();
    }
}