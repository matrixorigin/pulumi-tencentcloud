// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps
{
    [TencentcloudResourceType("tencentcloud:Mps/wordSample:WordSample")]
    public partial class WordSample : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Keyword. Length limit: 20 characters.
        /// </summary>
        [Output("keyword")]
        public Output<string> Keyword { get; private set; } = null!;

        /// <summary>
        /// Keyword tag. Array length limit: 20 tags. Each tag length limit: 128 characters.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Keyword usage. Valid values: 1.`Recognition.Ocr`: OCR-based content recognition. 2.`Recognition.Asr`: ASR-based content
        /// recognition. 3.`Review.Ocr`: OCR-based inappropriate information recognition. 4.`Review.Asr`: ASR-based inappropriate
        /// information recognition.
        /// </summary>
        [Output("usages")]
        public Output<ImmutableArray<string>> Usages { get; private set; } = null!;


        /// <summary>
        /// Create a WordSample resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WordSample(string name, WordSampleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mps/wordSample:WordSample", name, args ?? new WordSampleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WordSample(string name, Input<string> id, WordSampleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mps/wordSample:WordSample", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WordSample resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WordSample Get(string name, Input<string> id, WordSampleState? state = null, CustomResourceOptions? options = null)
        {
            return new WordSample(name, id, state, options);
        }
    }

    public sealed class WordSampleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Keyword. Length limit: 20 characters.
        /// </summary>
        [Input("keyword", required: true)]
        public Input<string> Keyword { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Keyword tag. Array length limit: 20 tags. Each tag length limit: 128 characters.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("usages", required: true)]
        private InputList<string>? _usages;

        /// <summary>
        /// Keyword usage. Valid values: 1.`Recognition.Ocr`: OCR-based content recognition. 2.`Recognition.Asr`: ASR-based content
        /// recognition. 3.`Review.Ocr`: OCR-based inappropriate information recognition. 4.`Review.Asr`: ASR-based inappropriate
        /// information recognition.
        /// </summary>
        public InputList<string> Usages
        {
            get => _usages ?? (_usages = new InputList<string>());
            set => _usages = value;
        }

        public WordSampleArgs()
        {
        }
        public static new WordSampleArgs Empty => new WordSampleArgs();
    }

    public sealed class WordSampleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Keyword. Length limit: 20 characters.
        /// </summary>
        [Input("keyword")]
        public Input<string>? Keyword { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Keyword tag. Array length limit: 20 tags. Each tag length limit: 128 characters.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("usages")]
        private InputList<string>? _usages;

        /// <summary>
        /// Keyword usage. Valid values: 1.`Recognition.Ocr`: OCR-based content recognition. 2.`Recognition.Asr`: ASR-based content
        /// recognition. 3.`Review.Ocr`: OCR-based inappropriate information recognition. 4.`Review.Asr`: ASR-based inappropriate
        /// information recognition.
        /// </summary>
        public InputList<string> Usages
        {
            get => _usages ?? (_usages = new InputList<string>());
            set => _usages = value;
        }

        public WordSampleState()
        {
        }
        public static new WordSampleState Empty => new WordSampleState();
    }
}
