// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssl
{
    [TencentcloudResourceType("tencentcloud:Ssl/revokeCertificateOperation:RevokeCertificateOperation")]
    public partial class RevokeCertificateOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Certificate ID.
        /// </summary>
        [Output("certificateId")]
        public Output<string> CertificateId { get; private set; } = null!;

        /// <summary>
        /// Reasons for revoking certificate.
        /// </summary>
        [Output("reason")]
        public Output<string?> Reason { get; private set; } = null!;


        /// <summary>
        /// Create a RevokeCertificateOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RevokeCertificateOperation(string name, RevokeCertificateOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssl/revokeCertificateOperation:RevokeCertificateOperation", name, args ?? new RevokeCertificateOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RevokeCertificateOperation(string name, Input<string> id, RevokeCertificateOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssl/revokeCertificateOperation:RevokeCertificateOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RevokeCertificateOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RevokeCertificateOperation Get(string name, Input<string> id, RevokeCertificateOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new RevokeCertificateOperation(name, id, state, options);
        }
    }

    public sealed class RevokeCertificateOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate ID.
        /// </summary>
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        /// <summary>
        /// Reasons for revoking certificate.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        public RevokeCertificateOperationArgs()
        {
        }
        public static new RevokeCertificateOperationArgs Empty => new RevokeCertificateOperationArgs();
    }

    public sealed class RevokeCertificateOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate ID.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// Reasons for revoking certificate.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        public RevokeCertificateOperationState()
        {
        }
        public static new RevokeCertificateOperationState Empty => new RevokeCertificateOperationState();
    }
}
