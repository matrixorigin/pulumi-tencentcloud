// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcm.Outputs
{

    [OutputType]
    public sealed class MeshConfigIstioSmartDns
    {
        public readonly bool? IstioMetaDnsAutoAllocate;
        public readonly bool? IstioMetaDnsCapture;

        [OutputConstructor]
        private MeshConfigIstioSmartDns(
            bool? istioMetaDnsAutoAllocate,

            bool? istioMetaDnsCapture)
        {
            IstioMetaDnsAutoAllocate = istioMetaDnsAutoAllocate;
            IstioMetaDnsCapture = istioMetaDnsCapture;
        }
    }
}