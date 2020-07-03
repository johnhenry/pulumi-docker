// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ContainerUploadGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        [Input("contentBase64")]
        public Input<string>? ContentBase64 { get; set; }

        /// <summary>
        /// If true, the file will be uploaded with user
        /// executable permission.
        /// Defaults to false.
        /// </summary>
        [Input("executable")]
        public Input<bool>? Executable { get; set; }

        /// <summary>
        /// path to a file in the container.
        /// </summary>
        [Input("file", required: true)]
        public Input<string> File { get; set; } = null!;

        /// <summary>
        /// A filename that references a file which will be uploaded as the object content. This allows for large file uploads that do not get stored in state.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// If using `source`, this will force an update if the file content has updated but the filename has not.
        /// </summary>
        [Input("sourceHash")]
        public Input<string>? SourceHash { get; set; }

        public ContainerUploadGetArgs()
        {
        }
    }
}
