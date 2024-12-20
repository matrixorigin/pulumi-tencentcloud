// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dasb
{
    [TencentcloudResourceType("tencentcloud:Dasb/cmdTemplate:CmdTemplate")]
    public partial class CmdTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Command list, n separated, maximum length 32768 bytes.
        /// </summary>
        [Output("cmdList")]
        public Output<string> CmdList { get; private set; } = null!;

        /// <summary>
        /// Template name, maximum length 32 characters, cannot contain blank characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a CmdTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CmdTemplate(string name, CmdTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dasb/cmdTemplate:CmdTemplate", name, args ?? new CmdTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CmdTemplate(string name, Input<string> id, CmdTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dasb/cmdTemplate:CmdTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CmdTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CmdTemplate Get(string name, Input<string> id, CmdTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new CmdTemplate(name, id, state, options);
        }
    }

    public sealed class CmdTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Command list, n separated, maximum length 32768 bytes.
        /// </summary>
        [Input("cmdList", required: true)]
        public Input<string> CmdList { get; set; } = null!;

        /// <summary>
        /// Template name, maximum length 32 characters, cannot contain blank characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public CmdTemplateArgs()
        {
        }
        public static new CmdTemplateArgs Empty => new CmdTemplateArgs();
    }

    public sealed class CmdTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Command list, n separated, maximum length 32768 bytes.
        /// </summary>
        [Input("cmdList")]
        public Input<string>? CmdList { get; set; }

        /// <summary>
        /// Template name, maximum length 32 characters, cannot contain blank characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public CmdTemplateState()
        {
        }
        public static new CmdTemplateState Empty => new CmdTemplateState();
    }
}
