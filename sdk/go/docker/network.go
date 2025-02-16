// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"context"
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
//
// ## Import
//
// Docker networks can be imported using the long id, e.g. for a network with the short id `p73jelnrme5f`
//
// ```sh
//  $ pulumi import docker:index/network:Network foo $(docker network inspect -f {{.ID}} p73)
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

type NetworkInput interface {
	pulumi.Input

	ToNetworkOutput() NetworkOutput
	ToNetworkOutputWithContext(ctx context.Context) NetworkOutput
}

func (*Network) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil))
}

func (i *Network) ToNetworkOutput() NetworkOutput {
	return i.ToNetworkOutputWithContext(context.Background())
}

func (i *Network) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkOutput)
}

func (i *Network) ToNetworkPtrOutput() NetworkPtrOutput {
	return i.ToNetworkPtrOutputWithContext(context.Background())
}

func (i *Network) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPtrOutput)
}

type NetworkPtrInput interface {
	pulumi.Input

	ToNetworkPtrOutput() NetworkPtrOutput
	ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput
}

type networkPtrType NetworkArgs

func (*networkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil))
}

func (i *networkPtrType) ToNetworkPtrOutput() NetworkPtrOutput {
	return i.ToNetworkPtrOutputWithContext(context.Background())
}

func (i *networkPtrType) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPtrOutput)
}

// NetworkArrayInput is an input type that accepts NetworkArray and NetworkArrayOutput values.
// You can construct a concrete instance of `NetworkArrayInput` via:
//
//          NetworkArray{ NetworkArgs{...} }
type NetworkArrayInput interface {
	pulumi.Input

	ToNetworkArrayOutput() NetworkArrayOutput
	ToNetworkArrayOutputWithContext(context.Context) NetworkArrayOutput
}

type NetworkArray []NetworkInput

func (NetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Network)(nil))
}

func (i NetworkArray) ToNetworkArrayOutput() NetworkArrayOutput {
	return i.ToNetworkArrayOutputWithContext(context.Background())
}

func (i NetworkArray) ToNetworkArrayOutputWithContext(ctx context.Context) NetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkArrayOutput)
}

// NetworkMapInput is an input type that accepts NetworkMap and NetworkMapOutput values.
// You can construct a concrete instance of `NetworkMapInput` via:
//
//          NetworkMap{ "key": NetworkArgs{...} }
type NetworkMapInput interface {
	pulumi.Input

	ToNetworkMapOutput() NetworkMapOutput
	ToNetworkMapOutputWithContext(context.Context) NetworkMapOutput
}

type NetworkMap map[string]NetworkInput

func (NetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Network)(nil))
}

func (i NetworkMap) ToNetworkMapOutput() NetworkMapOutput {
	return i.ToNetworkMapOutputWithContext(context.Background())
}

func (i NetworkMap) ToNetworkMapOutputWithContext(ctx context.Context) NetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkMapOutput)
}

type NetworkOutput struct {
	*pulumi.OutputState
}

func (NetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil))
}

func (o NetworkOutput) ToNetworkOutput() NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkPtrOutput() NetworkPtrOutput {
	return o.ToNetworkPtrOutputWithContext(context.Background())
}

func (o NetworkOutput) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return o.ApplyT(func(v Network) *Network {
		return &v
	}).(NetworkPtrOutput)
}

type NetworkPtrOutput struct {
	*pulumi.OutputState
}

func (NetworkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil))
}

func (o NetworkPtrOutput) ToNetworkPtrOutput() NetworkPtrOutput {
	return o
}

func (o NetworkPtrOutput) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return o
}

type NetworkArrayOutput struct{ *pulumi.OutputState }

func (NetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Network)(nil))
}

func (o NetworkArrayOutput) ToNetworkArrayOutput() NetworkArrayOutput {
	return o
}

func (o NetworkArrayOutput) ToNetworkArrayOutputWithContext(ctx context.Context) NetworkArrayOutput {
	return o
}

func (o NetworkArrayOutput) Index(i pulumi.IntInput) NetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Network {
		return vs[0].([]Network)[vs[1].(int)]
	}).(NetworkOutput)
}

type NetworkMapOutput struct{ *pulumi.OutputState }

func (NetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Network)(nil))
}

func (o NetworkMapOutput) ToNetworkMapOutput() NetworkMapOutput {
	return o
}

func (o NetworkMapOutput) ToNetworkMapOutputWithContext(ctx context.Context) NetworkMapOutput {
	return o
}

func (o NetworkMapOutput) MapIndex(k pulumi.StringInput) NetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Network {
		return vs[0].(map[string]Network)[vs[1].(string)]
	}).(NetworkOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkOutput{})
	pulumi.RegisterOutputType(NetworkPtrOutput{})
	pulumi.RegisterOutputType(NetworkArrayOutput{})
	pulumi.RegisterOutputType(NetworkMapOutput{})
}
