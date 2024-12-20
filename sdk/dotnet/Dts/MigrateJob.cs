// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dts
{
    [TencentcloudResourceType("tencentcloud:Dts/migrateJob:MigrateJob")]
    public partial class MigrateJob : global::Pulumi.CustomResource
    {
        /// <summary>
        /// AutoRetryTimeRangeMinutes.
        /// </summary>
        [Output("autoRetryTimeRangeMinutes")]
        public Output<int?> AutoRetryTimeRangeMinutes { get; private set; } = null!;

        /// <summary>
        /// DstInfo.
        /// </summary>
        [Output("dstInfo")]
        public Output<Outputs.MigrateJobDstInfo> DstInfo { get; private set; } = null!;

        /// <summary>
        /// ExpectRunTime.
        /// </summary>
        [Output("expectRunTime")]
        public Output<string> ExpectRunTime { get; private set; } = null!;

        /// <summary>
        /// Migration job configuration options, used to describe how the task performs migration.
        /// </summary>
        [Output("migrateOption")]
        public Output<Outputs.MigrateJobMigrateOption> MigrateOption { get; private set; } = null!;

        /// <summary>
        /// Run Mode. eg:immediate,timed.
        /// </summary>
        [Output("runMode")]
        public Output<string> RunMode { get; private set; } = null!;

        /// <summary>
        /// Migrate service Id from `tencentcloud_dts_migrate_service`.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// SrcInfo.
        /// </summary>
        [Output("srcInfo")]
        public Output<Outputs.MigrateJobSrcInfo> SrcInfo { get; private set; } = null!;

        /// <summary>
        /// Migrate job status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a MigrateJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MigrateJob(string name, MigrateJobArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dts/migrateJob:MigrateJob", name, args ?? new MigrateJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MigrateJob(string name, Input<string> id, MigrateJobState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dts/migrateJob:MigrateJob", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/matrixorigin",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MigrateJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MigrateJob Get(string name, Input<string> id, MigrateJobState? state = null, CustomResourceOptions? options = null)
        {
            return new MigrateJob(name, id, state, options);
        }
    }

    public sealed class MigrateJobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AutoRetryTimeRangeMinutes.
        /// </summary>
        [Input("autoRetryTimeRangeMinutes")]
        public Input<int>? AutoRetryTimeRangeMinutes { get; set; }

        /// <summary>
        /// DstInfo.
        /// </summary>
        [Input("dstInfo", required: true)]
        public Input<Inputs.MigrateJobDstInfoArgs> DstInfo { get; set; } = null!;

        /// <summary>
        /// ExpectRunTime.
        /// </summary>
        [Input("expectRunTime")]
        public Input<string>? ExpectRunTime { get; set; }

        /// <summary>
        /// Migration job configuration options, used to describe how the task performs migration.
        /// </summary>
        [Input("migrateOption", required: true)]
        public Input<Inputs.MigrateJobMigrateOptionArgs> MigrateOption { get; set; } = null!;

        /// <summary>
        /// Run Mode. eg:immediate,timed.
        /// </summary>
        [Input("runMode", required: true)]
        public Input<string> RunMode { get; set; } = null!;

        /// <summary>
        /// Migrate service Id from `tencentcloud_dts_migrate_service`.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        /// <summary>
        /// SrcInfo.
        /// </summary>
        [Input("srcInfo", required: true)]
        public Input<Inputs.MigrateJobSrcInfoArgs> SrcInfo { get; set; } = null!;

        public MigrateJobArgs()
        {
        }
        public static new MigrateJobArgs Empty => new MigrateJobArgs();
    }

    public sealed class MigrateJobState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AutoRetryTimeRangeMinutes.
        /// </summary>
        [Input("autoRetryTimeRangeMinutes")]
        public Input<int>? AutoRetryTimeRangeMinutes { get; set; }

        /// <summary>
        /// DstInfo.
        /// </summary>
        [Input("dstInfo")]
        public Input<Inputs.MigrateJobDstInfoGetArgs>? DstInfo { get; set; }

        /// <summary>
        /// ExpectRunTime.
        /// </summary>
        [Input("expectRunTime")]
        public Input<string>? ExpectRunTime { get; set; }

        /// <summary>
        /// Migration job configuration options, used to describe how the task performs migration.
        /// </summary>
        [Input("migrateOption")]
        public Input<Inputs.MigrateJobMigrateOptionGetArgs>? MigrateOption { get; set; }

        /// <summary>
        /// Run Mode. eg:immediate,timed.
        /// </summary>
        [Input("runMode")]
        public Input<string>? RunMode { get; set; }

        /// <summary>
        /// Migrate service Id from `tencentcloud_dts_migrate_service`.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// SrcInfo.
        /// </summary>
        [Input("srcInfo")]
        public Input<Inputs.MigrateJobSrcInfoGetArgs>? SrcInfo { get; set; }

        /// <summary>
        /// Migrate job status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public MigrateJobState()
        {
        }
        public static new MigrateJobState Empty => new MigrateJobState();
    }
}
