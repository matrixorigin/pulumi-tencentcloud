// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cfw.Inputs
{

    public sealed class NatInstanceNewModeItemsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("eips", required: true)]
        private InputList<string>? _eips;
        public InputList<string> Eips
        {
            get => _eips ?? (_eips = new InputList<string>());
            set => _eips = value;
        }

        [Input("vpcLists", required: true)]
        private InputList<string>? _vpcLists;
        public InputList<string> VpcLists
        {
            get => _vpcLists ?? (_vpcLists = new InputList<string>());
            set => _vpcLists = value;
        }

        public NatInstanceNewModeItemsGetArgs()
        {
        }
        public static new NatInstanceNewModeItemsGetArgs Empty => new NatInstanceNewModeItemsGetArgs();
    }
}
