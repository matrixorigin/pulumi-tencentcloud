// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Pts
{
    [TencentcloudResourceType("tencentcloud:Pts/cronJobAbort:CronJobAbort")]
    public partial class CronJobAbort : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cron job ID.
        /// </summary>
        [Output("cronJobId")]
        public Output<string> CronJobId { get; private set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a CronJobAbort resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CronJobAbort(string name, CronJobAbortArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Pts/cronJobAbort:CronJobAbort", name, args ?? new CronJobAbortArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CronJobAbort(string name, Input<string> id, CronJobAbortState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Pts/cronJobAbort:CronJobAbort", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CronJobAbort resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CronJobAbort Get(string name, Input<string> id, CronJobAbortState? state = null, CustomResourceOptions? options = null)
        {
            return new CronJobAbort(name, id, state, options);
        }
    }

    public sealed class CronJobAbortArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cron job ID.
        /// </summary>
        [Input("cronJobId", required: true)]
        public Input<string> CronJobId { get; set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public CronJobAbortArgs()
        {
        }
        public static new CronJobAbortArgs Empty => new CronJobAbortArgs();
    }

    public sealed class CronJobAbortState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cron job ID.
        /// </summary>
        [Input("cronJobId")]
        public Input<string>? CronJobId { get; set; }

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public CronJobAbortState()
        {
        }
        public static new CronJobAbortState Empty => new CronJobAbortState();
    }
}
