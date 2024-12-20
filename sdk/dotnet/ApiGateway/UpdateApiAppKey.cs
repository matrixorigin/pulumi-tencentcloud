// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.ApiGateway
{
    [TencentcloudResourceType("tencentcloud:ApiGateway/updateApiAppKey:UpdateApiAppKey")]
    public partial class UpdateApiAppKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Application unique ID.
        /// </summary>
        [Output("apiAppId")]
        public Output<string> ApiAppId { get; private set; } = null!;

        /// <summary>
        /// Key of the application.
        /// </summary>
        [Output("apiAppKey")]
        public Output<string> ApiAppKey { get; private set; } = null!;


        /// <summary>
        /// Create a UpdateApiAppKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UpdateApiAppKey(string name, UpdateApiAppKeyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/updateApiAppKey:UpdateApiAppKey", name, args ?? new UpdateApiAppKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UpdateApiAppKey(string name, Input<string> id, UpdateApiAppKeyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/updateApiAppKey:UpdateApiAppKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UpdateApiAppKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UpdateApiAppKey Get(string name, Input<string> id, UpdateApiAppKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new UpdateApiAppKey(name, id, state, options);
        }
    }

    public sealed class UpdateApiAppKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application unique ID.
        /// </summary>
        [Input("apiAppId", required: true)]
        public Input<string> ApiAppId { get; set; } = null!;

        /// <summary>
        /// Key of the application.
        /// </summary>
        [Input("apiAppKey", required: true)]
        public Input<string> ApiAppKey { get; set; } = null!;

        public UpdateApiAppKeyArgs()
        {
        }
        public static new UpdateApiAppKeyArgs Empty => new UpdateApiAppKeyArgs();
    }

    public sealed class UpdateApiAppKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application unique ID.
        /// </summary>
        [Input("apiAppId")]
        public Input<string>? ApiAppId { get; set; }

        /// <summary>
        /// Key of the application.
        /// </summary>
        [Input("apiAppKey")]
        public Input<string>? ApiAppKey { get; set; }

        public UpdateApiAppKeyState()
        {
        }
        public static new UpdateApiAppKeyState Empty => new UpdateApiAppKeyState();
    }
}
