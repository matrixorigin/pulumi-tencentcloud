// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssl
{
    [TencentcloudResourceType("tencentcloud:Ssl/updateCertificateRecordRetryOperation:UpdateCertificateRecordRetryOperation")]
    public partial class UpdateCertificateRecordRetryOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Deployment record details ID to be retried.
        /// </summary>
        [Output("deployRecordDetailId")]
        public Output<int?> DeployRecordDetailId { get; private set; } = null!;

        /// <summary>
        /// Deployment record ID to be retried.
        /// </summary>
        [Output("deployRecordId")]
        public Output<int?> DeployRecordId { get; private set; } = null!;


        /// <summary>
        /// Create a UpdateCertificateRecordRetryOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UpdateCertificateRecordRetryOperation(string name, UpdateCertificateRecordRetryOperationArgs? args = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssl/updateCertificateRecordRetryOperation:UpdateCertificateRecordRetryOperation", name, args ?? new UpdateCertificateRecordRetryOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UpdateCertificateRecordRetryOperation(string name, Input<string> id, UpdateCertificateRecordRetryOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssl/updateCertificateRecordRetryOperation:UpdateCertificateRecordRetryOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UpdateCertificateRecordRetryOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UpdateCertificateRecordRetryOperation Get(string name, Input<string> id, UpdateCertificateRecordRetryOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new UpdateCertificateRecordRetryOperation(name, id, state, options);
        }
    }

    public sealed class UpdateCertificateRecordRetryOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Deployment record details ID to be retried.
        /// </summary>
        [Input("deployRecordDetailId")]
        public Input<int>? DeployRecordDetailId { get; set; }

        /// <summary>
        /// Deployment record ID to be retried.
        /// </summary>
        [Input("deployRecordId")]
        public Input<int>? DeployRecordId { get; set; }

        public UpdateCertificateRecordRetryOperationArgs()
        {
        }
        public static new UpdateCertificateRecordRetryOperationArgs Empty => new UpdateCertificateRecordRetryOperationArgs();
    }

    public sealed class UpdateCertificateRecordRetryOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Deployment record details ID to be retried.
        /// </summary>
        [Input("deployRecordDetailId")]
        public Input<int>? DeployRecordDetailId { get; set; }

        /// <summary>
        /// Deployment record ID to be retried.
        /// </summary>
        [Input("deployRecordId")]
        public Input<int>? DeployRecordId { get; set; }

        public UpdateCertificateRecordRetryOperationState()
        {
        }
        public static new UpdateCertificateRecordRetryOperationState Empty => new UpdateCertificateRecordRetryOperationState();
    }
}
