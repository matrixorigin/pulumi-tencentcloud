// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Thpc.Inputs
{

    public sealed class WorkspacesTagSpecificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("tags", required: true)]
        private InputList<Inputs.WorkspacesTagSpecificationTagArgs>? _tags;
        public InputList<Inputs.WorkspacesTagSpecificationTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.WorkspacesTagSpecificationTagArgs>());
            set => _tags = value;
        }

        public WorkspacesTagSpecificationArgs()
        {
        }
        public static new WorkspacesTagSpecificationArgs Empty => new WorkspacesTagSpecificationArgs();
    }
}