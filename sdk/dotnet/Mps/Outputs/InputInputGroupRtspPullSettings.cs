// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class InputInputGroupRtspPullSettings
    {
        public readonly ImmutableArray<Outputs.InputInputGroupRtspPullSettingsSourceAddress> SourceAddresses;

        [OutputConstructor]
        private InputInputGroupRtspPullSettings(ImmutableArray<Outputs.InputInputGroupRtspPullSettingsSourceAddress> sourceAddresses)
        {
            SourceAddresses = sourceAddresses;
        }
    }
}