// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tsf
{
    [TencentcloudResourceType("tencentcloud:Tsf/instancesAttachment:InstancesAttachment")]
    public partial class InstancesAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Image feature ID list.
        /// </summary>
        [Output("featureIdLists")]
        public Output<ImmutableArray<string>> FeatureIdLists { get; private set; } = null!;

        /// <summary>
        /// Operating system image ID.
        /// </summary>
        [Output("imageId")]
        public Output<string?> ImageId { get; private set; } = null!;

        /// <summary>
        /// Additional instance parameter information.
        /// </summary>
        [Output("instanceAdvancedSettings")]
        public Output<Outputs.InstancesAttachmentInstanceAdvancedSettings?> InstanceAdvancedSettings { get; private set; } = null!;

        /// <summary>
        /// Cloud server ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Cloud server import mode, required for virtual machine clusters, not required for container clusters. R: Reinstall TSF
        /// system image, M: Manual installation of agent.
        /// </summary>
        [Output("instanceImportMode")]
        public Output<string?> InstanceImportMode { get; private set; } = null!;

        /// <summary>
        /// Associated key for system reinstallation.
        /// </summary>
        [Output("keyId")]
        public Output<string?> KeyId { get; private set; } = null!;

        /// <summary>
        /// Image customization type.
        /// </summary>
        [Output("osCustomizeType")]
        public Output<string?> OsCustomizeType { get; private set; } = null!;

        /// <summary>
        /// Operating system name.
        /// </summary>
        [Output("osName")]
        public Output<string?> OsName { get; private set; } = null!;

        /// <summary>
        /// Reset system password.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Security group.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Security group setting.
        /// </summary>
        [Output("sgId")]
        public Output<string?> SgId { get; private set; } = null!;


        /// <summary>
        /// Create a InstancesAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstancesAttachment(string name, InstancesAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/instancesAttachment:InstancesAttachment", name, args ?? new InstancesAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstancesAttachment(string name, Input<string> id, InstancesAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/instancesAttachment:InstancesAttachment", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/matrixorigin",
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing InstancesAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstancesAttachment Get(string name, Input<string> id, InstancesAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new InstancesAttachment(name, id, state, options);
        }
    }

    public sealed class InstancesAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("featureIdLists")]
        private InputList<string>? _featureIdLists;

        /// <summary>
        /// Image feature ID list.
        /// </summary>
        public InputList<string> FeatureIdLists
        {
            get => _featureIdLists ?? (_featureIdLists = new InputList<string>());
            set => _featureIdLists = value;
        }

        /// <summary>
        /// Operating system image ID.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// Additional instance parameter information.
        /// </summary>
        [Input("instanceAdvancedSettings")]
        public Input<Inputs.InstancesAttachmentInstanceAdvancedSettingsArgs>? InstanceAdvancedSettings { get; set; }

        /// <summary>
        /// Cloud server ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Cloud server import mode, required for virtual machine clusters, not required for container clusters. R: Reinstall TSF
        /// system image, M: Manual installation of agent.
        /// </summary>
        [Input("instanceImportMode")]
        public Input<string>? InstanceImportMode { get; set; }

        /// <summary>
        /// Associated key for system reinstallation.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// Image customization type.
        /// </summary>
        [Input("osCustomizeType")]
        public Input<string>? OsCustomizeType { get; set; }

        /// <summary>
        /// Operating system name.
        /// </summary>
        [Input("osName")]
        public Input<string>? OsName { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Reset system password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security group.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Security group setting.
        /// </summary>
        [Input("sgId")]
        public Input<string>? SgId { get; set; }

        public InstancesAttachmentArgs()
        {
        }
        public static new InstancesAttachmentArgs Empty => new InstancesAttachmentArgs();
    }

    public sealed class InstancesAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("featureIdLists")]
        private InputList<string>? _featureIdLists;

        /// <summary>
        /// Image feature ID list.
        /// </summary>
        public InputList<string> FeatureIdLists
        {
            get => _featureIdLists ?? (_featureIdLists = new InputList<string>());
            set => _featureIdLists = value;
        }

        /// <summary>
        /// Operating system image ID.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// Additional instance parameter information.
        /// </summary>
        [Input("instanceAdvancedSettings")]
        public Input<Inputs.InstancesAttachmentInstanceAdvancedSettingsGetArgs>? InstanceAdvancedSettings { get; set; }

        /// <summary>
        /// Cloud server ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Cloud server import mode, required for virtual machine clusters, not required for container clusters. R: Reinstall TSF
        /// system image, M: Manual installation of agent.
        /// </summary>
        [Input("instanceImportMode")]
        public Input<string>? InstanceImportMode { get; set; }

        /// <summary>
        /// Associated key for system reinstallation.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// Image customization type.
        /// </summary>
        [Input("osCustomizeType")]
        public Input<string>? OsCustomizeType { get; set; }

        /// <summary>
        /// Operating system name.
        /// </summary>
        [Input("osName")]
        public Input<string>? OsName { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Reset system password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security group.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Security group setting.
        /// </summary>
        [Input("sgId")]
        public Input<string>? SgId { get; set; }

        public InstancesAttachmentState()
        {
        }
        public static new InstancesAttachmentState Empty => new InstancesAttachmentState();
    }
}