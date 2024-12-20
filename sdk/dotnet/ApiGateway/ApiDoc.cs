// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.ApiGateway
{
    [TencentcloudResourceType("tencentcloud:ApiGateway/apiDoc:ApiDoc")]
    public partial class ApiDoc : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Api Document count.
        /// </summary>
        [Output("apiCount")]
        public Output<int> ApiCount { get; private set; } = null!;

        /// <summary>
        /// Api Document ID.
        /// </summary>
        [Output("apiDocId")]
        public Output<string> ApiDocId { get; private set; } = null!;

        /// <summary>
        /// Api Document name.
        /// </summary>
        [Output("apiDocName")]
        public Output<string> ApiDocName { get; private set; } = null!;

        /// <summary>
        /// API Document Build Status.
        /// </summary>
        [Output("apiDocStatus")]
        public Output<string> ApiDocStatus { get; private set; } = null!;

        /// <summary>
        /// API Document Access URI.
        /// </summary>
        [Output("apiDocUri")]
        public Output<string> ApiDocUri { get; private set; } = null!;

        /// <summary>
        /// List of APIs for generating documents.
        /// </summary>
        [Output("apiIds")]
        public Output<ImmutableArray<string>> ApiIds { get; private set; } = null!;

        /// <summary>
        /// List of names for generating documents.
        /// </summary>
        [Output("apiNames")]
        public Output<ImmutableArray<string>> ApiNames { get; private set; } = null!;

        /// <summary>
        /// Env name.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// Number of API document releases.
        /// </summary>
        [Output("releaseCount")]
        public Output<int> ReleaseCount { get; private set; } = null!;

        /// <summary>
        /// Service name.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// API Document service name.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// API Document Sharing Password.
        /// </summary>
        [Output("sharePassword")]
        public Output<string> SharePassword { get; private set; } = null!;

        /// <summary>
        /// API Document update time.
        /// </summary>
        [Output("updatedTime")]
        public Output<string> UpdatedTime { get; private set; } = null!;

        /// <summary>
        /// API Document Viewing Times.
        /// </summary>
        [Output("viewCount")]
        public Output<int> ViewCount { get; private set; } = null!;


        /// <summary>
        /// Create a ApiDoc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiDoc(string name, ApiDocArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/apiDoc:ApiDoc", name, args ?? new ApiDocArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiDoc(string name, Input<string> id, ApiDocState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/apiDoc:ApiDoc", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApiDoc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiDoc Get(string name, Input<string> id, ApiDocState? state = null, CustomResourceOptions? options = null)
        {
            return new ApiDoc(name, id, state, options);
        }
    }

    public sealed class ApiDocArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Api Document name.
        /// </summary>
        [Input("apiDocName", required: true)]
        public Input<string> ApiDocName { get; set; } = null!;

        [Input("apiIds", required: true)]
        private InputList<string>? _apiIds;

        /// <summary>
        /// List of APIs for generating documents.
        /// </summary>
        public InputList<string> ApiIds
        {
            get => _apiIds ?? (_apiIds = new InputList<string>());
            set => _apiIds = value;
        }

        /// <summary>
        /// Env name.
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        /// <summary>
        /// Service name.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public ApiDocArgs()
        {
        }
        public static new ApiDocArgs Empty => new ApiDocArgs();
    }

    public sealed class ApiDocState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Api Document count.
        /// </summary>
        [Input("apiCount")]
        public Input<int>? ApiCount { get; set; }

        /// <summary>
        /// Api Document ID.
        /// </summary>
        [Input("apiDocId")]
        public Input<string>? ApiDocId { get; set; }

        /// <summary>
        /// Api Document name.
        /// </summary>
        [Input("apiDocName")]
        public Input<string>? ApiDocName { get; set; }

        /// <summary>
        /// API Document Build Status.
        /// </summary>
        [Input("apiDocStatus")]
        public Input<string>? ApiDocStatus { get; set; }

        /// <summary>
        /// API Document Access URI.
        /// </summary>
        [Input("apiDocUri")]
        public Input<string>? ApiDocUri { get; set; }

        [Input("apiIds")]
        private InputList<string>? _apiIds;

        /// <summary>
        /// List of APIs for generating documents.
        /// </summary>
        public InputList<string> ApiIds
        {
            get => _apiIds ?? (_apiIds = new InputList<string>());
            set => _apiIds = value;
        }

        [Input("apiNames")]
        private InputList<string>? _apiNames;

        /// <summary>
        /// List of names for generating documents.
        /// </summary>
        public InputList<string> ApiNames
        {
            get => _apiNames ?? (_apiNames = new InputList<string>());
            set => _apiNames = value;
        }

        /// <summary>
        /// Env name.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// Number of API document releases.
        /// </summary>
        [Input("releaseCount")]
        public Input<int>? ReleaseCount { get; set; }

        /// <summary>
        /// Service name.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// API Document service name.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// API Document Sharing Password.
        /// </summary>
        [Input("sharePassword")]
        public Input<string>? SharePassword { get; set; }

        /// <summary>
        /// API Document update time.
        /// </summary>
        [Input("updatedTime")]
        public Input<string>? UpdatedTime { get; set; }

        /// <summary>
        /// API Document Viewing Times.
        /// </summary>
        [Input("viewCount")]
        public Input<int>? ViewCount { get; set; }

        public ApiDocState()
        {
        }
        public static new ApiDocState Empty => new ApiDocState();
    }
}
