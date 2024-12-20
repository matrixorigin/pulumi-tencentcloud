// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Css
{
    public static class GetPadTemplates
    {
        public static Task<GetPadTemplatesResult> InvokeAsync(GetPadTemplatesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPadTemplatesResult>("tencentcloud:Css/getPadTemplates:getPadTemplates", args ?? new GetPadTemplatesArgs(), options.WithDefaults());

        public static Output<GetPadTemplatesResult> Invoke(GetPadTemplatesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPadTemplatesResult>("tencentcloud:Css/getPadTemplates:getPadTemplates", args ?? new GetPadTemplatesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPadTemplatesArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetPadTemplatesArgs()
        {
        }
        public static new GetPadTemplatesArgs Empty => new GetPadTemplatesArgs();
    }

    public sealed class GetPadTemplatesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetPadTemplatesInvokeArgs()
        {
        }
        public static new GetPadTemplatesInvokeArgs Empty => new GetPadTemplatesInvokeArgs();
    }


    [OutputType]
    public sealed class GetPadTemplatesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<Outputs.GetPadTemplatesTemplateResult> Templates;

        [OutputConstructor]
        private GetPadTemplatesResult(
            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetPadTemplatesTemplateResult> templates)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            Templates = templates;
        }
    }
}
