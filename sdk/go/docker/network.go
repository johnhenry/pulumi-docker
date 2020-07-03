// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Docker Network. This can be used alongside
// [docker\_container](https://www.terraform.io/docs/providers/docker/r/container.html)
// to create virtual networks within the docker environment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-docker/sdk/v2/go/docker"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := docker.NewNetwork(ctx, "privateNetwork", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Network struct {
	pulumi.CustomResourceState

	// Enable manual container attachment to the network.
	// Defaults to `false`.
	Attachable pulumi.BoolPtrOutput `pulumi:"attachable"`
	// Requests daemon to check for networks
	// with same name.
	CheckDuplicate pulumi.BoolPtrOutput `pulumi:"checkDuplicate"`
	// Name of the network driver to use. Defaults to
	// `bridge` driver.
	Driver pulumi.StringOutput `pulumi:"driver"`
	// Create swarm routing-mesh network.
	// Defaults to `false`.
	Ingress pulumi.BoolPtrOutput `pulumi:"ingress"`
	// Restrict external access to the network.
	// Defaults to `false`.
	Internal pulumi.BoolOutput `pulumi:"internal"`
	// See IPAM config below for
	// details.
	IpamConfigs NetworkIpamConfigArrayOutput `pulumi:"ipamConfigs"`
	// Driver used by the custom IP scheme of the
	// network.
	IpamDriver pulumi.StringPtrOutput `pulumi:"ipamDriver"`
	// Enable IPv6 networking.
	// Defaults to `false`.
	Ipv6 pulumi.BoolPtrOutput `pulumi:"ipv6"`
	// See Labels below for details.
	Labels NetworkLabelArrayOutput `pulumi:"labels"`
	// The name of the Docker network.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network specific options to be used by
	// the drivers.
	Options pulumi.MapOutput    `pulumi:"options"`
	Scope   pulumi.StringOutput `pulumi:"scope"`
}

// NewNetwork registers a new resource with the given unique name, arguments, and options.
func NewNetwork(ctx *pulumi.Context,
	name string, args *NetworkArgs, opts ...pulumi.ResourceOption) (*Network, error) {
	if args == nil {
		args = &NetworkArgs{}
	}
	var resource Network
	err := ctx.RegisterResource("docker:index/network:Network", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetwork gets an existing Network resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkState, opts ...pulumi.ResourceOption) (*Network, error) {
	var resource Network
	err := ctx.ReadResource("docker:index/network:Network", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Network resources.
type networkState struct {
	// Enable manual container attachment to the network.
	// Defaults to `false`.
	Attachable *bool `pulumi:"attachable"`
	// Requests daemon to check for networks
	// with same name.
	CheckDuplicate *bool `pulumi:"checkDuplicate"`
	// Name of the network driver to use. Defaults to
	// `bridge` driver.
	Driver *string `pulumi:"driver"`
	// Create swarm routing-mesh network.
	// Defaults to `false`.
	Ingress *bool `pulumi:"ingress"`
	// Restrict external access to the network.
	// Defaults to `false`.
	Internal *bool `pulumi:"internal"`
	// See IPAM config below for
	// details.
	IpamConfigs []NetworkIpamConfig `pulumi:"ipamConfigs"`
	// Driver used by the custom IP scheme of the
	// network.
	IpamDriver *string `pulumi:"ipamDriver"`
	// Enable IPv6 networking.
	// Defaults to `false`.
	Ipv6 *bool `pulumi:"ipv6"`
	// See Labels below for details.
	Labels []NetworkLabel `pulumi:"labels"`
	// The name of the Docker network.
	Name *string `pulumi:"name"`
	// Network specific options to be used by
	// the drivers.
	Options map[string]interface{} `pulumi:"options"`
	Scope   *string                `pulumi:"scope"`
}

type NetworkState struct {
	// Enable manual container attachment to the network.
	// Defaults to `false`.
	Attachable pulumi.BoolPtrInput
	// Requests daemon to check for networks
	// with same name.
	CheckDuplicate pulumi.BoolPtrInput
	// Name of the network driver to use. Defaults to
	// `bridge` driver.
	Driver pulumi.StringPtrInput
	// Create swarm routing-mesh network.
	// Defaults to `false`.
	Ingress pulumi.BoolPtrInput
	// Restrict external access to the network.
	// Defaults to `false`.
	Internal pulumi.BoolPtrInput
	// See IPAM config below for
	// details.
	IpamConfigs NetworkIpamConfigArrayInput
	// Driver used by the custom IP scheme of the
	// network.
	IpamDriver pulumi.StringPtrInput
	// Enable IPv6 networking.
	// Defaults to `false`.
	Ipv6 pulumi.BoolPtrInput
	// See Labels below for details.
	Labels NetworkLabelArrayInput
	// The name of the Docker network.
	Name pulumi.StringPtrInput
	// Network specific options to be used by
	// the drivers.
	Options pulumi.MapInput
	Scope   pulumi.StringPtrInput
}

func (NetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkState)(nil)).Elem()
}

type networkArgs struct {
	// Enable manual container attachment to the network.
	// Defaults to `false`.
	Attachable *bool `pulumi:"attachable"`
	// Requests daemon to check for networks
	// with same name.
	CheckDuplicate *bool `pulumi:"checkDuplicate"`
	// Name of the network driver to use. Defaults to
	// `bridge` driver.
	Driver *string `pulumi:"driver"`
	// Create swarm routing-mesh network.
	// Defaults to `false`.
	Ingress *bool `pulumi:"ingress"`
	// Restrict external access to the network.
	// Defaults to `false`.
	Internal *bool `pulumi:"internal"`
	// See IPAM config below for
	// details.
	IpamConfigs []NetworkIpamConfig `pulumi:"ipamConfigs"`
	// Driver used by the custom IP scheme of the
	// network.
	IpamDriver *string `pulumi:"ipamDriver"`
	// Enable IPv6 networking.
	// Defaults to `false`.
	Ipv6 *bool `pulumi:"ipv6"`
	// See Labels below for details.
	Labels []NetworkLabel `pulumi:"labels"`
	// The name of the Docker network.
	Name *string `pulumi:"name"`
	// Network specific options to be used by
	// the drivers.
	Options map[string]interface{} `pulumi:"options"`
}

// The set of arguments for constructing a Network resource.
type NetworkArgs struct {
	// Enable manual container attachment to the network.
	// Defaults to `false`.
	Attachable pulumi.BoolPtrInput
	// Requests daemon to check for networks
	// with same name.
	CheckDuplicate pulumi.BoolPtrInput
	// Name of the network driver to use. Defaults to
	// `bridge` driver.
	Driver pulumi.StringPtrInput
	// Create swarm routing-mesh network.
	// Defaults to `false`.
	Ingress pulumi.BoolPtrInput
	// Restrict external access to the network.
	// Defaults to `false`.
	Internal pulumi.BoolPtrInput
	// See IPAM config below for
	// details.
	IpamConfigs NetworkIpamConfigArrayInput
	// Driver used by the custom IP scheme of the
	// network.
	IpamDriver pulumi.StringPtrInput
	// Enable IPv6 networking.
	// Defaults to `false`.
	Ipv6 pulumi.BoolPtrInput
	// See Labels below for details.
	Labels NetworkLabelArrayInput
	// The name of the Docker network.
	Name pulumi.StringPtrInput
	// Network specific options to be used by
	// the drivers.
	Options pulumi.MapInput
}

func (NetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkArgs)(nil)).Elem()
}
