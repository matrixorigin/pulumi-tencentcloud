// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cfs
{
    [TencentcloudResourceType("tencentcloud:Cfs/fileSystem:FileSystem")]
    public partial class FileSystem : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of a access group.
        /// </summary>
        [Output("accessGroupId")]
        public Output<string> AccessGroupId { get; private set; } = null!;

        /// <summary>
        /// The available zone that the file system locates at.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// File system capacity, in GiB (required for the Turbo series). For Standard Turbo, the minimum purchase required is
        /// 40,960 GiB (40 TiB) and the expansion increment is 20,480 GiB (20 TiB). For High-Performance Turbo, the minimum purchase
        /// required is 20,480 GiB (20 TiB) and the expansion increment is 10,240 GiB (10 TiB).
        /// </summary>
        [Output("capacity")]
        public Output<int> Capacity { get; private set; } = null!;

        /// <summary>
        /// CCN instance ID (required if the network type is CCN).
        /// </summary>
        [Output("ccnId")]
        public Output<string> CcnId { get; private set; } = null!;

        /// <summary>
        /// CCN IP range used by the CFS (required if the network type is CCN), which cannot conflict with other IP ranges bound in
        /// CCN.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// Create time of the file system.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Mount root-directory.
        /// </summary>
        [Output("fsId")]
        public Output<string> FsId { get; private set; } = null!;

        /// <summary>
        /// IP of mount point.
        /// </summary>
        [Output("mountIp")]
        public Output<string> MountIp { get; private set; } = null!;

        /// <summary>
        /// Name of a file system.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network type, Default `VPC`. Valid values: `VPC` and `CCN`. Select `VPC` for a Standard or High-Performance file system,
        /// and `CCN` for a Standard Turbo or High-Performance Turbo one.
        /// </summary>
        [Output("netInterface")]
        public Output<string?> NetInterface { get; private set; } = null!;

        /// <summary>
        /// File system protocol. Valid values: `NFS`, `CIFS`, `TURBO`. If this parameter is left empty, `NFS` is used by default.
        /// For the Turbo series, you must set this parameter to `TURBO`.
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// Storage type of the file system. Valid values: `SD` (Standard), `HP` (High-Performance), `TB` (Standard Turbo), and `TP`
        /// (High-Performance Turbo). Default value: `SD`.
        /// </summary>
        [Output("storageType")]
        public Output<string?> StorageType { get; private set; } = null!;

        /// <summary>
        /// ID of a subnet.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Instance tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// ID of a VPC network.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a FileSystem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FileSystem(string name, FileSystemArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfs/fileSystem:FileSystem", name, args ?? new FileSystemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FileSystem(string name, Input<string> id, FileSystemState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfs/fileSystem:FileSystem", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FileSystem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FileSystem Get(string name, Input<string> id, FileSystemState? state = null, CustomResourceOptions? options = null)
        {
            return new FileSystem(name, id, state, options);
        }
    }

    public sealed class FileSystemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of a access group.
        /// </summary>
        [Input("accessGroupId", required: true)]
        public Input<string> AccessGroupId { get; set; } = null!;

        /// <summary>
        /// The available zone that the file system locates at.
        /// </summary>
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        /// <summary>
        /// File system capacity, in GiB (required for the Turbo series). For Standard Turbo, the minimum purchase required is
        /// 40,960 GiB (40 TiB) and the expansion increment is 20,480 GiB (20 TiB). For High-Performance Turbo, the minimum purchase
        /// required is 20,480 GiB (20 TiB) and the expansion increment is 10,240 GiB (10 TiB).
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// CCN instance ID (required if the network type is CCN).
        /// </summary>
        [Input("ccnId")]
        public Input<string>? CcnId { get; set; }

        /// <summary>
        /// CCN IP range used by the CFS (required if the network type is CCN), which cannot conflict with other IP ranges bound in
        /// CCN.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// IP of mount point.
        /// </summary>
        [Input("mountIp")]
        public Input<string>? MountIp { get; set; }

        /// <summary>
        /// Name of a file system.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network type, Default `VPC`. Valid values: `VPC` and `CCN`. Select `VPC` for a Standard or High-Performance file system,
        /// and `CCN` for a Standard Turbo or High-Performance Turbo one.
        /// </summary>
        [Input("netInterface")]
        public Input<string>? NetInterface { get; set; }

        /// <summary>
        /// File system protocol. Valid values: `NFS`, `CIFS`, `TURBO`. If this parameter is left empty, `NFS` is used by default.
        /// For the Turbo series, you must set this parameter to `TURBO`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Storage type of the file system. Valid values: `SD` (Standard), `HP` (High-Performance), `TB` (Standard Turbo), and `TP`
        /// (High-Performance Turbo). Default value: `SD`.
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        /// <summary>
        /// ID of a subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Instance tags.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of a VPC network.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public FileSystemArgs()
        {
        }
        public static new FileSystemArgs Empty => new FileSystemArgs();
    }

    public sealed class FileSystemState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of a access group.
        /// </summary>
        [Input("accessGroupId")]
        public Input<string>? AccessGroupId { get; set; }

        /// <summary>
        /// The available zone that the file system locates at.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// File system capacity, in GiB (required for the Turbo series). For Standard Turbo, the minimum purchase required is
        /// 40,960 GiB (40 TiB) and the expansion increment is 20,480 GiB (20 TiB). For High-Performance Turbo, the minimum purchase
        /// required is 20,480 GiB (20 TiB) and the expansion increment is 10,240 GiB (10 TiB).
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// CCN instance ID (required if the network type is CCN).
        /// </summary>
        [Input("ccnId")]
        public Input<string>? CcnId { get; set; }

        /// <summary>
        /// CCN IP range used by the CFS (required if the network type is CCN), which cannot conflict with other IP ranges bound in
        /// CCN.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Create time of the file system.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Mount root-directory.
        /// </summary>
        [Input("fsId")]
        public Input<string>? FsId { get; set; }

        /// <summary>
        /// IP of mount point.
        /// </summary>
        [Input("mountIp")]
        public Input<string>? MountIp { get; set; }

        /// <summary>
        /// Name of a file system.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network type, Default `VPC`. Valid values: `VPC` and `CCN`. Select `VPC` for a Standard or High-Performance file system,
        /// and `CCN` for a Standard Turbo or High-Performance Turbo one.
        /// </summary>
        [Input("netInterface")]
        public Input<string>? NetInterface { get; set; }

        /// <summary>
        /// File system protocol. Valid values: `NFS`, `CIFS`, `TURBO`. If this parameter is left empty, `NFS` is used by default.
        /// For the Turbo series, you must set this parameter to `TURBO`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Storage type of the file system. Valid values: `SD` (Standard), `HP` (High-Performance), `TB` (Standard Turbo), and `TP`
        /// (High-Performance Turbo). Default value: `SD`.
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        /// <summary>
        /// ID of a subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Instance tags.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of a VPC network.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public FileSystemState()
        {
        }
        public static new FileSystemState Empty => new FileSystemState();
    }
}
