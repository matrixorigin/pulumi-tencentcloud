// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kubernetes
{
    [TencentcloudResourceType("tencentcloud:Kubernetes/authAttachment:AuthAttachment")]
    public partial class AuthAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creating ClientId of the identity provider.
        /// </summary>
        [Output("autoCreateClientIds")]
        public Output<ImmutableArray<string>> AutoCreateClientIds { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, the rbac rule will be created automatically which allow anonymous user to access
        /// '/.well-known/openid-configuration' and '/openid/v1/jwks'.
        /// </summary>
        [Output("autoCreateDiscoveryAnonymousAuth")]
        public Output<bool?> AutoCreateDiscoveryAnonymousAuth { get; private set; } = null!;

        /// <summary>
        /// Creating an identity provider.
        /// </summary>
        [Output("autoCreateOidcConfig")]
        public Output<bool> AutoCreateOidcConfig { get; private set; } = null!;

        /// <summary>
        /// Creating the PodIdentityWebhook component. if `auto_create_oidc_config` is true, this field must set true.
        /// </summary>
        [Output("autoInstallPodIdentityWebhookAddon")]
        public Output<bool> AutoInstallPodIdentityWebhookAddon { get; private set; } = null!;

        /// <summary>
        /// ID of clusters.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Specify service-account-issuer. If use_tke_default is set to `true`, please do not set this field.
        /// </summary>
        [Output("issuer")]
        public Output<string?> Issuer { get; private set; } = null!;

        /// <summary>
        /// Specify service-account-jwks-uri. If use_tke_default is set to `true`, please do not set this field.
        /// </summary>
        [Output("jwksUri")]
        public Output<string?> JwksUri { get; private set; } = null!;

        /// <summary>
        /// The default issuer of tke. If use_tke_default is set to `true`, this parameter will be set to the default value.
        /// </summary>
        [Output("tkeDefaultIssuer")]
        public Output<string> TkeDefaultIssuer { get; private set; } = null!;

        /// <summary>
        /// The default jwks_uri of tke. If use_tke_default is set to `true`, this parameter will be set to the default value.
        /// </summary>
        [Output("tkeDefaultJwksUri")]
        public Output<string> TkeDefaultJwksUri { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, the issuer and jwks_uri will be generated automatically by tke, please do not set issuer and jwks_uri.
        /// </summary>
        [Output("useTkeDefault")]
        public Output<bool?> UseTkeDefault { get; private set; } = null!;


        /// <summary>
        /// Create a AuthAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthAttachment(string name, AuthAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/authAttachment:AuthAttachment", name, args ?? new AuthAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthAttachment(string name, Input<string> id, AuthAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/authAttachment:AuthAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthAttachment Get(string name, Input<string> id, AuthAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthAttachment(name, id, state, options);
        }
    }

    public sealed class AuthAttachmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoCreateClientIds")]
        private InputList<string>? _autoCreateClientIds;

        /// <summary>
        /// Creating ClientId of the identity provider.
        /// </summary>
        public InputList<string> AutoCreateClientIds
        {
            get => _autoCreateClientIds ?? (_autoCreateClientIds = new InputList<string>());
            set => _autoCreateClientIds = value;
        }

        /// <summary>
        /// If set to `true`, the rbac rule will be created automatically which allow anonymous user to access
        /// '/.well-known/openid-configuration' and '/openid/v1/jwks'.
        /// </summary>
        [Input("autoCreateDiscoveryAnonymousAuth")]
        public Input<bool>? AutoCreateDiscoveryAnonymousAuth { get; set; }

        /// <summary>
        /// Creating an identity provider.
        /// </summary>
        [Input("autoCreateOidcConfig")]
        public Input<bool>? AutoCreateOidcConfig { get; set; }

        /// <summary>
        /// Creating the PodIdentityWebhook component. if `auto_create_oidc_config` is true, this field must set true.
        /// </summary>
        [Input("autoInstallPodIdentityWebhookAddon")]
        public Input<bool>? AutoInstallPodIdentityWebhookAddon { get; set; }

        /// <summary>
        /// ID of clusters.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Specify service-account-issuer. If use_tke_default is set to `true`, please do not set this field.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// Specify service-account-jwks-uri. If use_tke_default is set to `true`, please do not set this field.
        /// </summary>
        [Input("jwksUri")]
        public Input<string>? JwksUri { get; set; }

        /// <summary>
        /// If set to `true`, the issuer and jwks_uri will be generated automatically by tke, please do not set issuer and jwks_uri.
        /// </summary>
        [Input("useTkeDefault")]
        public Input<bool>? UseTkeDefault { get; set; }

        public AuthAttachmentArgs()
        {
        }
        public static new AuthAttachmentArgs Empty => new AuthAttachmentArgs();
    }

    public sealed class AuthAttachmentState : global::Pulumi.ResourceArgs
    {
        [Input("autoCreateClientIds")]
        private InputList<string>? _autoCreateClientIds;

        /// <summary>
        /// Creating ClientId of the identity provider.
        /// </summary>
        public InputList<string> AutoCreateClientIds
        {
            get => _autoCreateClientIds ?? (_autoCreateClientIds = new InputList<string>());
            set => _autoCreateClientIds = value;
        }

        /// <summary>
        /// If set to `true`, the rbac rule will be created automatically which allow anonymous user to access
        /// '/.well-known/openid-configuration' and '/openid/v1/jwks'.
        /// </summary>
        [Input("autoCreateDiscoveryAnonymousAuth")]
        public Input<bool>? AutoCreateDiscoveryAnonymousAuth { get; set; }

        /// <summary>
        /// Creating an identity provider.
        /// </summary>
        [Input("autoCreateOidcConfig")]
        public Input<bool>? AutoCreateOidcConfig { get; set; }

        /// <summary>
        /// Creating the PodIdentityWebhook component. if `auto_create_oidc_config` is true, this field must set true.
        /// </summary>
        [Input("autoInstallPodIdentityWebhookAddon")]
        public Input<bool>? AutoInstallPodIdentityWebhookAddon { get; set; }

        /// <summary>
        /// ID of clusters.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Specify service-account-issuer. If use_tke_default is set to `true`, please do not set this field.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// Specify service-account-jwks-uri. If use_tke_default is set to `true`, please do not set this field.
        /// </summary>
        [Input("jwksUri")]
        public Input<string>? JwksUri { get; set; }

        /// <summary>
        /// The default issuer of tke. If use_tke_default is set to `true`, this parameter will be set to the default value.
        /// </summary>
        [Input("tkeDefaultIssuer")]
        public Input<string>? TkeDefaultIssuer { get; set; }

        /// <summary>
        /// The default jwks_uri of tke. If use_tke_default is set to `true`, this parameter will be set to the default value.
        /// </summary>
        [Input("tkeDefaultJwksUri")]
        public Input<string>? TkeDefaultJwksUri { get; set; }

        /// <summary>
        /// If set to `true`, the issuer and jwks_uri will be generated automatically by tke, please do not set issuer and jwks_uri.
        /// </summary>
        [Input("useTkeDefault")]
        public Input<bool>? UseTkeDefault { get; set; }

        public AuthAttachmentState()
        {
        }
        public static new AuthAttachmentState Empty => new AuthAttachmentState();
    }
}
