// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker
{
    /// <summary>
    /// Creates and destroys a volume in Docker. This can be used alongside
    /// [docker\_container](https://www.terraform.io/docs/providers/docker/r/container.html)
    /// to prepare volumes that can be shared across containers.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-docker/blob/master/website/docs/r/volume.html.markdown.
    /// </summary>
    public partial class Volume : Pulumi.CustomResource
    {
        /// <summary>
        /// Driver type for the volume (defaults to local).
        /// </summary>
        [Output("driver")]
        public Output<string> Driver { get; private set; } = null!;

        /// <summary>
        /// Options specific to the driver.
        /// </summary>
        [Output("driverOpts")]
        public Output<ImmutableDictionary<string, object>?> DriverOpts { get; private set; } = null!;

        /// <summary>
        /// User-defined key/value metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.VolumeLabels>> Labels { get; private set; } = null!;

        [Output("mountpoint")]
        public Output<string> Mountpoint { get; private set; } = null!;

        /// <summary>
        /// The name of the Docker volume (generated if not
        /// provided).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Volume resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Volume(string name, VolumeArgs? args = null, CustomResourceOptions? options = null)
            : base("docker:index/volume:Volume", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Volume(string name, Input<string> id, VolumeState? state = null, CustomResourceOptions? options = null)
            : base("docker:index/volume:Volume", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Volume resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Volume Get(string name, Input<string> id, VolumeState? state = null, CustomResourceOptions? options = null)
        {
            return new Volume(name, id, state, options);
        }
    }

    public sealed class VolumeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Driver type for the volume (defaults to local).
        /// </summary>
        [Input("driver")]
        public Input<string>? Driver { get; set; }

        [Input("driverOpts")]
        private InputMap<object>? _driverOpts;

        /// <summary>
        /// Options specific to the driver.
        /// </summary>
        public InputMap<object> DriverOpts
        {
            get => _driverOpts ?? (_driverOpts = new InputMap<object>());
            set => _driverOpts = value;
        }

        [Input("labels")]
        private InputList<Inputs.VolumeLabelsArgs>? _labels;

        /// <summary>
        /// User-defined key/value metadata.
        /// </summary>
        public InputList<Inputs.VolumeLabelsArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.VolumeLabelsArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the Docker volume (generated if not
        /// provided).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VolumeArgs()
        {
        }
    }

    public sealed class VolumeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Driver type for the volume (defaults to local).
        /// </summary>
        [Input("driver")]
        public Input<string>? Driver { get; set; }

        [Input("driverOpts")]
        private InputMap<object>? _driverOpts;

        /// <summary>
        /// Options specific to the driver.
        /// </summary>
        public InputMap<object> DriverOpts
        {
            get => _driverOpts ?? (_driverOpts = new InputMap<object>());
            set => _driverOpts = value;
        }

        [Input("labels")]
        private InputList<Inputs.VolumeLabelsGetArgs>? _labels;

        /// <summary>
        /// User-defined key/value metadata.
        /// </summary>
        public InputList<Inputs.VolumeLabelsGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.VolumeLabelsGetArgs>());
            set => _labels = value;
        }

        [Input("mountpoint")]
        public Input<string>? Mountpoint { get; set; }

        /// <summary>
        /// The name of the Docker volume (generated if not
        /// provided).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VolumeState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class VolumeLabelsArgs : Pulumi.ResourceArgs
    {
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public VolumeLabelsArgs()
        {
        }
    }

    public sealed class VolumeLabelsGetArgs : Pulumi.ResourceArgs
    {
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public VolumeLabelsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class VolumeLabels
    {
        public readonly string Label;
        public readonly string Value;

        [OutputConstructor]
        private VolumeLabels(
            string label,
            string value)
        {
            Label = label;
            Value = value;
        }
    }
    }
}
