// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package unifi

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `Network` manages WAN/LAN/VLAN networks.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// 	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		vlanId := 10
// 		if param := cfg.GetFloat("vlanId"); param != 0 {
// 			vlanId = param
// 		}
// 		_, err := unifi.NewNetwork(ctx, "vlan", &unifi.NetworkArgs{
// 			Purpose:     pulumi.String("corporate"),
// 			Subnet:      pulumi.String("10.0.0.1/24"),
// 			VlanId:      pulumi.Float64(vlanId),
// 			DhcpStart:   pulumi.String("10.0.0.6"),
// 			DhcpStop:    pulumi.String("10.0.0.254"),
// 			DhcpEnabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = unifi.NewNetwork(ctx, "wan", &unifi.NetworkArgs{
// 			Purpose:         pulumi.String("wan"),
// 			WanNetworkgroup: pulumi.String("WAN"),
// 			WanType:         pulumi.String("pppoe"),
// 			WanIp:           pulumi.String("192.168.1.1"),
// 			WanEgressQos:    pulumi.Int(1),
// 			WanUsername:     pulumi.String("username"),
// 			XWanPassword:    pulumi.String("password"),
// 		})
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
// # import from provider configured site
//
// ```sh
//  $ pulumi import unifi:index/network:Network mynetwork 5dc28e5e9106d105bdc87217
// ```
//
// # import from another site
//
// ```sh
//  $ pulumi import unifi:index/network:Network mynetwork bfa2l6i7:5dc28e5e9106d105bdc87217
// ```
//
// # import network by name
//
// ```sh
//  $ pulumi import unifi:index/network:Network mynetwork name=LAN
// ```
type Network struct {
	pulumi.CustomResourceState

	// Specifies the IPv4 addresses for the DNS server to be returned from the DHCP server. Leave blank to disable this feature.
	DhcpDns pulumi.StringArrayOutput `pulumi:"dhcpDns"`
	// Specifies whether DHCP is enabled or not on this network.
	DhcpEnabled pulumi.BoolPtrOutput `pulumi:"dhcpEnabled"`
	// Specifies the lease time for DHCP addresses. Defaults to `86400`.
	DhcpLease pulumi.IntPtrOutput `pulumi:"dhcpLease"`
	// Specifies whether DHCP relay is enabled or not on this network.
	DhcpRelayEnabled pulumi.BoolPtrOutput `pulumi:"dhcpRelayEnabled"`
	// The IPv4 address where the DHCP range of addresses starts.
	DhcpStart pulumi.StringPtrOutput `pulumi:"dhcpStart"`
	// The IPv4 address where the DHCP range of addresses stops.
	DhcpStop pulumi.StringPtrOutput `pulumi:"dhcpStop"`
	// Toggles on the DHCP boot options. Should be set to true when you want to have dhcpd*boot*filename, and dhcpd*boot*server to take effect.
	DhcpdBootEnabled pulumi.BoolPtrOutput `pulumi:"dhcpdBootEnabled"`
	// Specifies the file to PXE boot from on the dhcpd*boot*server.
	DhcpdBootFilename pulumi.StringPtrOutput `pulumi:"dhcpdBootFilename"`
	// Specifies the IPv4 address of a TFTP server to network boot from.
	DhcpdBootServer pulumi.StringPtrOutput `pulumi:"dhcpdBootServer"`
	// The domain name of this network.
	DomainName pulumi.StringPtrOutput `pulumi:"domainName"`
	// Specifies whether IGMP snooping is enabled or not.
	IgmpSnooping pulumi.BoolPtrOutput `pulumi:"igmpSnooping"`
	// Specifies which type of IPv6 connection to use. Defaults to `none`.
	Ipv6InterfaceType pulumi.StringPtrOutput `pulumi:"ipv6InterfaceType"`
	// Specifies which WAN interface to use for IPv6 PD.
	Ipv6PdInterface pulumi.StringPtrOutput `pulumi:"ipv6PdInterface"`
	// Specifies the IPv6 Prefix ID.
	Ipv6PdPrefixid pulumi.StringPtrOutput `pulumi:"ipv6PdPrefixid"`
	// Specifies whether to enable router advertisements or not.
	Ipv6RaEnable pulumi.BoolPtrOutput `pulumi:"ipv6RaEnable"`
	// Specifies the static IPv6 subnet when ipv6*interface*type is 'static'.
	Ipv6StaticSubnet pulumi.StringPtrOutput `pulumi:"ipv6StaticSubnet"`
	// The name of the network.
	Name pulumi.StringOutput `pulumi:"name"`
	// The group of the network. Defaults to `LAN`.
	NetworkGroup pulumi.StringPtrOutput `pulumi:"networkGroup"`
	// The purpose of the network. Must be one of `corporate`, `guest`, `wan`, or `vlan-only`.
	Purpose pulumi.StringOutput `pulumi:"purpose"`
	// The name of the site to associate the network with.
	Site pulumi.StringOutput `pulumi:"site"`
	// The subnet of the network. Must be a valid CIDR address.
	Subnet pulumi.StringPtrOutput `pulumi:"subnet"`
	// The VLAN ID of the network.
	VlanId pulumi.IntPtrOutput `pulumi:"vlanId"`
	// DNS servers IPs of the WAN.
	WanDns pulumi.StringArrayOutput `pulumi:"wanDns"`
	// Specifies the WAN egress quality of service. Defaults to `0`.
	WanEgressQos pulumi.IntPtrOutput `pulumi:"wanEgressQos"`
	// The IPv4 gateway of the WAN.
	WanGateway pulumi.StringPtrOutput `pulumi:"wanGateway"`
	// The IPv4 address of the WAN.
	WanIp pulumi.StringPtrOutput `pulumi:"wanIp"`
	// The IPv4 netmask of the WAN.
	WanNetmask pulumi.StringPtrOutput `pulumi:"wanNetmask"`
	// Specifies the WAN network group. Must be one of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
	WanNetworkgroup pulumi.StringPtrOutput `pulumi:"wanNetworkgroup"`
	// Specifies the IPV4 WAN connection type. Must be one of either `disabled`, `static`, `dhcp`, or `pppoe`.
	WanType pulumi.StringPtrOutput `pulumi:"wanType"`
	// Specifies the IPV4 WAN username.
	WanUsername pulumi.StringPtrOutput `pulumi:"wanUsername"`
	// Specifies the IPV4 WAN password.
	XWanPassword pulumi.StringPtrOutput `pulumi:"xWanPassword"`
}

// NewNetwork registers a new resource with the given unique name, arguments, and options.
func NewNetwork(ctx *pulumi.Context,
	name string, args *NetworkArgs, opts ...pulumi.ResourceOption) (*Network, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Purpose == nil {
		return nil, errors.New("invalid value for required argument 'Purpose'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Network
	err := ctx.RegisterResource("unifi:index/network:Network", name, args, &resource, opts...)
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
	err := ctx.ReadResource("unifi:index/network:Network", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Network resources.
type networkState struct {
	// Specifies the IPv4 addresses for the DNS server to be returned from the DHCP server. Leave blank to disable this feature.
	DhcpDns []string `pulumi:"dhcpDns"`
	// Specifies whether DHCP is enabled or not on this network.
	DhcpEnabled *bool `pulumi:"dhcpEnabled"`
	// Specifies the lease time for DHCP addresses. Defaults to `86400`.
	DhcpLease *int `pulumi:"dhcpLease"`
	// Specifies whether DHCP relay is enabled or not on this network.
	DhcpRelayEnabled *bool `pulumi:"dhcpRelayEnabled"`
	// The IPv4 address where the DHCP range of addresses starts.
	DhcpStart *string `pulumi:"dhcpStart"`
	// The IPv4 address where the DHCP range of addresses stops.
	DhcpStop *string `pulumi:"dhcpStop"`
	// Toggles on the DHCP boot options. Should be set to true when you want to have dhcpd*boot*filename, and dhcpd*boot*server to take effect.
	DhcpdBootEnabled *bool `pulumi:"dhcpdBootEnabled"`
	// Specifies the file to PXE boot from on the dhcpd*boot*server.
	DhcpdBootFilename *string `pulumi:"dhcpdBootFilename"`
	// Specifies the IPv4 address of a TFTP server to network boot from.
	DhcpdBootServer *string `pulumi:"dhcpdBootServer"`
	// The domain name of this network.
	DomainName *string `pulumi:"domainName"`
	// Specifies whether IGMP snooping is enabled or not.
	IgmpSnooping *bool `pulumi:"igmpSnooping"`
	// Specifies which type of IPv6 connection to use. Defaults to `none`.
	Ipv6InterfaceType *string `pulumi:"ipv6InterfaceType"`
	// Specifies which WAN interface to use for IPv6 PD.
	Ipv6PdInterface *string `pulumi:"ipv6PdInterface"`
	// Specifies the IPv6 Prefix ID.
	Ipv6PdPrefixid *string `pulumi:"ipv6PdPrefixid"`
	// Specifies whether to enable router advertisements or not.
	Ipv6RaEnable *bool `pulumi:"ipv6RaEnable"`
	// Specifies the static IPv6 subnet when ipv6*interface*type is 'static'.
	Ipv6StaticSubnet *string `pulumi:"ipv6StaticSubnet"`
	// The name of the network.
	Name *string `pulumi:"name"`
	// The group of the network. Defaults to `LAN`.
	NetworkGroup *string `pulumi:"networkGroup"`
	// The purpose of the network. Must be one of `corporate`, `guest`, `wan`, or `vlan-only`.
	Purpose *string `pulumi:"purpose"`
	// The name of the site to associate the network with.
	Site *string `pulumi:"site"`
	// The subnet of the network. Must be a valid CIDR address.
	Subnet *string `pulumi:"subnet"`
	// The VLAN ID of the network.
	VlanId *int `pulumi:"vlanId"`
	// DNS servers IPs of the WAN.
	WanDns []string `pulumi:"wanDns"`
	// Specifies the WAN egress quality of service. Defaults to `0`.
	WanEgressQos *int `pulumi:"wanEgressQos"`
	// The IPv4 gateway of the WAN.
	WanGateway *string `pulumi:"wanGateway"`
	// The IPv4 address of the WAN.
	WanIp *string `pulumi:"wanIp"`
	// The IPv4 netmask of the WAN.
	WanNetmask *string `pulumi:"wanNetmask"`
	// Specifies the WAN network group. Must be one of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
	WanNetworkgroup *string `pulumi:"wanNetworkgroup"`
	// Specifies the IPV4 WAN connection type. Must be one of either `disabled`, `static`, `dhcp`, or `pppoe`.
	WanType *string `pulumi:"wanType"`
	// Specifies the IPV4 WAN username.
	WanUsername *string `pulumi:"wanUsername"`
	// Specifies the IPV4 WAN password.
	XWanPassword *string `pulumi:"xWanPassword"`
}

type NetworkState struct {
	// Specifies the IPv4 addresses for the DNS server to be returned from the DHCP server. Leave blank to disable this feature.
	DhcpDns pulumi.StringArrayInput
	// Specifies whether DHCP is enabled or not on this network.
	DhcpEnabled pulumi.BoolPtrInput
	// Specifies the lease time for DHCP addresses. Defaults to `86400`.
	DhcpLease pulumi.IntPtrInput
	// Specifies whether DHCP relay is enabled or not on this network.
	DhcpRelayEnabled pulumi.BoolPtrInput
	// The IPv4 address where the DHCP range of addresses starts.
	DhcpStart pulumi.StringPtrInput
	// The IPv4 address where the DHCP range of addresses stops.
	DhcpStop pulumi.StringPtrInput
	// Toggles on the DHCP boot options. Should be set to true when you want to have dhcpd*boot*filename, and dhcpd*boot*server to take effect.
	DhcpdBootEnabled pulumi.BoolPtrInput
	// Specifies the file to PXE boot from on the dhcpd*boot*server.
	DhcpdBootFilename pulumi.StringPtrInput
	// Specifies the IPv4 address of a TFTP server to network boot from.
	DhcpdBootServer pulumi.StringPtrInput
	// The domain name of this network.
	DomainName pulumi.StringPtrInput
	// Specifies whether IGMP snooping is enabled or not.
	IgmpSnooping pulumi.BoolPtrInput
	// Specifies which type of IPv6 connection to use. Defaults to `none`.
	Ipv6InterfaceType pulumi.StringPtrInput
	// Specifies which WAN interface to use for IPv6 PD.
	Ipv6PdInterface pulumi.StringPtrInput
	// Specifies the IPv6 Prefix ID.
	Ipv6PdPrefixid pulumi.StringPtrInput
	// Specifies whether to enable router advertisements or not.
	Ipv6RaEnable pulumi.BoolPtrInput
	// Specifies the static IPv6 subnet when ipv6*interface*type is 'static'.
	Ipv6StaticSubnet pulumi.StringPtrInput
	// The name of the network.
	Name pulumi.StringPtrInput
	// The group of the network. Defaults to `LAN`.
	NetworkGroup pulumi.StringPtrInput
	// The purpose of the network. Must be one of `corporate`, `guest`, `wan`, or `vlan-only`.
	Purpose pulumi.StringPtrInput
	// The name of the site to associate the network with.
	Site pulumi.StringPtrInput
	// The subnet of the network. Must be a valid CIDR address.
	Subnet pulumi.StringPtrInput
	// The VLAN ID of the network.
	VlanId pulumi.IntPtrInput
	// DNS servers IPs of the WAN.
	WanDns pulumi.StringArrayInput
	// Specifies the WAN egress quality of service. Defaults to `0`.
	WanEgressQos pulumi.IntPtrInput
	// The IPv4 gateway of the WAN.
	WanGateway pulumi.StringPtrInput
	// The IPv4 address of the WAN.
	WanIp pulumi.StringPtrInput
	// The IPv4 netmask of the WAN.
	WanNetmask pulumi.StringPtrInput
	// Specifies the WAN network group. Must be one of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
	WanNetworkgroup pulumi.StringPtrInput
	// Specifies the IPV4 WAN connection type. Must be one of either `disabled`, `static`, `dhcp`, or `pppoe`.
	WanType pulumi.StringPtrInput
	// Specifies the IPV4 WAN username.
	WanUsername pulumi.StringPtrInput
	// Specifies the IPV4 WAN password.
	XWanPassword pulumi.StringPtrInput
}

func (NetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkState)(nil)).Elem()
}

type networkArgs struct {
	// Specifies the IPv4 addresses for the DNS server to be returned from the DHCP server. Leave blank to disable this feature.
	DhcpDns []string `pulumi:"dhcpDns"`
	// Specifies whether DHCP is enabled or not on this network.
	DhcpEnabled *bool `pulumi:"dhcpEnabled"`
	// Specifies the lease time for DHCP addresses. Defaults to `86400`.
	DhcpLease *int `pulumi:"dhcpLease"`
	// Specifies whether DHCP relay is enabled or not on this network.
	DhcpRelayEnabled *bool `pulumi:"dhcpRelayEnabled"`
	// The IPv4 address where the DHCP range of addresses starts.
	DhcpStart *string `pulumi:"dhcpStart"`
	// The IPv4 address where the DHCP range of addresses stops.
	DhcpStop *string `pulumi:"dhcpStop"`
	// Toggles on the DHCP boot options. Should be set to true when you want to have dhcpd*boot*filename, and dhcpd*boot*server to take effect.
	DhcpdBootEnabled *bool `pulumi:"dhcpdBootEnabled"`
	// Specifies the file to PXE boot from on the dhcpd*boot*server.
	DhcpdBootFilename *string `pulumi:"dhcpdBootFilename"`
	// Specifies the IPv4 address of a TFTP server to network boot from.
	DhcpdBootServer *string `pulumi:"dhcpdBootServer"`
	// The domain name of this network.
	DomainName *string `pulumi:"domainName"`
	// Specifies whether IGMP snooping is enabled or not.
	IgmpSnooping *bool `pulumi:"igmpSnooping"`
	// Specifies which type of IPv6 connection to use. Defaults to `none`.
	Ipv6InterfaceType *string `pulumi:"ipv6InterfaceType"`
	// Specifies which WAN interface to use for IPv6 PD.
	Ipv6PdInterface *string `pulumi:"ipv6PdInterface"`
	// Specifies the IPv6 Prefix ID.
	Ipv6PdPrefixid *string `pulumi:"ipv6PdPrefixid"`
	// Specifies whether to enable router advertisements or not.
	Ipv6RaEnable *bool `pulumi:"ipv6RaEnable"`
	// Specifies the static IPv6 subnet when ipv6*interface*type is 'static'.
	Ipv6StaticSubnet *string `pulumi:"ipv6StaticSubnet"`
	// The name of the network.
	Name *string `pulumi:"name"`
	// The group of the network. Defaults to `LAN`.
	NetworkGroup *string `pulumi:"networkGroup"`
	// The purpose of the network. Must be one of `corporate`, `guest`, `wan`, or `vlan-only`.
	Purpose string `pulumi:"purpose"`
	// The name of the site to associate the network with.
	Site *string `pulumi:"site"`
	// The subnet of the network. Must be a valid CIDR address.
	Subnet *string `pulumi:"subnet"`
	// The VLAN ID of the network.
	VlanId *int `pulumi:"vlanId"`
	// DNS servers IPs of the WAN.
	WanDns []string `pulumi:"wanDns"`
	// Specifies the WAN egress quality of service. Defaults to `0`.
	WanEgressQos *int `pulumi:"wanEgressQos"`
	// The IPv4 gateway of the WAN.
	WanGateway *string `pulumi:"wanGateway"`
	// The IPv4 address of the WAN.
	WanIp *string `pulumi:"wanIp"`
	// The IPv4 netmask of the WAN.
	WanNetmask *string `pulumi:"wanNetmask"`
	// Specifies the WAN network group. Must be one of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
	WanNetworkgroup *string `pulumi:"wanNetworkgroup"`
	// Specifies the IPV4 WAN connection type. Must be one of either `disabled`, `static`, `dhcp`, or `pppoe`.
	WanType *string `pulumi:"wanType"`
	// Specifies the IPV4 WAN username.
	WanUsername *string `pulumi:"wanUsername"`
	// Specifies the IPV4 WAN password.
	XWanPassword *string `pulumi:"xWanPassword"`
}

// The set of arguments for constructing a Network resource.
type NetworkArgs struct {
	// Specifies the IPv4 addresses for the DNS server to be returned from the DHCP server. Leave blank to disable this feature.
	DhcpDns pulumi.StringArrayInput
	// Specifies whether DHCP is enabled or not on this network.
	DhcpEnabled pulumi.BoolPtrInput
	// Specifies the lease time for DHCP addresses. Defaults to `86400`.
	DhcpLease pulumi.IntPtrInput
	// Specifies whether DHCP relay is enabled or not on this network.
	DhcpRelayEnabled pulumi.BoolPtrInput
	// The IPv4 address where the DHCP range of addresses starts.
	DhcpStart pulumi.StringPtrInput
	// The IPv4 address where the DHCP range of addresses stops.
	DhcpStop pulumi.StringPtrInput
	// Toggles on the DHCP boot options. Should be set to true when you want to have dhcpd*boot*filename, and dhcpd*boot*server to take effect.
	DhcpdBootEnabled pulumi.BoolPtrInput
	// Specifies the file to PXE boot from on the dhcpd*boot*server.
	DhcpdBootFilename pulumi.StringPtrInput
	// Specifies the IPv4 address of a TFTP server to network boot from.
	DhcpdBootServer pulumi.StringPtrInput
	// The domain name of this network.
	DomainName pulumi.StringPtrInput
	// Specifies whether IGMP snooping is enabled or not.
	IgmpSnooping pulumi.BoolPtrInput
	// Specifies which type of IPv6 connection to use. Defaults to `none`.
	Ipv6InterfaceType pulumi.StringPtrInput
	// Specifies which WAN interface to use for IPv6 PD.
	Ipv6PdInterface pulumi.StringPtrInput
	// Specifies the IPv6 Prefix ID.
	Ipv6PdPrefixid pulumi.StringPtrInput
	// Specifies whether to enable router advertisements or not.
	Ipv6RaEnable pulumi.BoolPtrInput
	// Specifies the static IPv6 subnet when ipv6*interface*type is 'static'.
	Ipv6StaticSubnet pulumi.StringPtrInput
	// The name of the network.
	Name pulumi.StringPtrInput
	// The group of the network. Defaults to `LAN`.
	NetworkGroup pulumi.StringPtrInput
	// The purpose of the network. Must be one of `corporate`, `guest`, `wan`, or `vlan-only`.
	Purpose pulumi.StringInput
	// The name of the site to associate the network with.
	Site pulumi.StringPtrInput
	// The subnet of the network. Must be a valid CIDR address.
	Subnet pulumi.StringPtrInput
	// The VLAN ID of the network.
	VlanId pulumi.IntPtrInput
	// DNS servers IPs of the WAN.
	WanDns pulumi.StringArrayInput
	// Specifies the WAN egress quality of service. Defaults to `0`.
	WanEgressQos pulumi.IntPtrInput
	// The IPv4 gateway of the WAN.
	WanGateway pulumi.StringPtrInput
	// The IPv4 address of the WAN.
	WanIp pulumi.StringPtrInput
	// The IPv4 netmask of the WAN.
	WanNetmask pulumi.StringPtrInput
	// Specifies the WAN network group. Must be one of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
	WanNetworkgroup pulumi.StringPtrInput
	// Specifies the IPV4 WAN connection type. Must be one of either `disabled`, `static`, `dhcp`, or `pppoe`.
	WanType pulumi.StringPtrInput
	// Specifies the IPV4 WAN username.
	WanUsername pulumi.StringPtrInput
	// Specifies the IPV4 WAN password.
	XWanPassword pulumi.StringPtrInput
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
	return reflect.TypeOf((**Network)(nil)).Elem()
}

func (i *Network) ToNetworkOutput() NetworkOutput {
	return i.ToNetworkOutputWithContext(context.Background())
}

func (i *Network) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkOutput)
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
	return reflect.TypeOf((*[]*Network)(nil)).Elem()
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
	return reflect.TypeOf((*map[string]*Network)(nil)).Elem()
}

func (i NetworkMap) ToNetworkMapOutput() NetworkMapOutput {
	return i.ToNetworkMapOutputWithContext(context.Background())
}

func (i NetworkMap) ToNetworkMapOutputWithContext(ctx context.Context) NetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkMapOutput)
}

type NetworkOutput struct{ *pulumi.OutputState }

func (NetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil)).Elem()
}

func (o NetworkOutput) ToNetworkOutput() NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return o
}

// Specifies the IPv4 addresses for the DNS server to be returned from the DHCP server. Leave blank to disable this feature.
func (o NetworkOutput) DhcpDns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Network) pulumi.StringArrayOutput { return v.DhcpDns }).(pulumi.StringArrayOutput)
}

// Specifies whether DHCP is enabled or not on this network.
func (o NetworkOutput) DhcpEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.BoolPtrOutput { return v.DhcpEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies the lease time for DHCP addresses. Defaults to `86400`.
func (o NetworkOutput) DhcpLease() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.IntPtrOutput { return v.DhcpLease }).(pulumi.IntPtrOutput)
}

// Specifies whether DHCP relay is enabled or not on this network.
func (o NetworkOutput) DhcpRelayEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.BoolPtrOutput { return v.DhcpRelayEnabled }).(pulumi.BoolPtrOutput)
}

// The IPv4 address where the DHCP range of addresses starts.
func (o NetworkOutput) DhcpStart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.DhcpStart }).(pulumi.StringPtrOutput)
}

// The IPv4 address where the DHCP range of addresses stops.
func (o NetworkOutput) DhcpStop() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.DhcpStop }).(pulumi.StringPtrOutput)
}

// Toggles on the DHCP boot options. Should be set to true when you want to have dhcpd*boot*filename, and dhcpd*boot*server to take effect.
func (o NetworkOutput) DhcpdBootEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.BoolPtrOutput { return v.DhcpdBootEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies the file to PXE boot from on the dhcpd*boot*server.
func (o NetworkOutput) DhcpdBootFilename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.DhcpdBootFilename }).(pulumi.StringPtrOutput)
}

// Specifies the IPv4 address of a TFTP server to network boot from.
func (o NetworkOutput) DhcpdBootServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.DhcpdBootServer }).(pulumi.StringPtrOutput)
}

// The domain name of this network.
func (o NetworkOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.DomainName }).(pulumi.StringPtrOutput)
}

// Specifies whether IGMP snooping is enabled or not.
func (o NetworkOutput) IgmpSnooping() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.BoolPtrOutput { return v.IgmpSnooping }).(pulumi.BoolPtrOutput)
}

// Specifies which type of IPv6 connection to use. Defaults to `none`.
func (o NetworkOutput) Ipv6InterfaceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.Ipv6InterfaceType }).(pulumi.StringPtrOutput)
}

// Specifies which WAN interface to use for IPv6 PD.
func (o NetworkOutput) Ipv6PdInterface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.Ipv6PdInterface }).(pulumi.StringPtrOutput)
}

// Specifies the IPv6 Prefix ID.
func (o NetworkOutput) Ipv6PdPrefixid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.Ipv6PdPrefixid }).(pulumi.StringPtrOutput)
}

// Specifies whether to enable router advertisements or not.
func (o NetworkOutput) Ipv6RaEnable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.BoolPtrOutput { return v.Ipv6RaEnable }).(pulumi.BoolPtrOutput)
}

// Specifies the static IPv6 subnet when ipv6*interface*type is 'static'.
func (o NetworkOutput) Ipv6StaticSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.Ipv6StaticSubnet }).(pulumi.StringPtrOutput)
}

// The name of the network.
func (o NetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The group of the network. Defaults to `LAN`.
func (o NetworkOutput) NetworkGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.NetworkGroup }).(pulumi.StringPtrOutput)
}

// The purpose of the network. Must be one of `corporate`, `guest`, `wan`, or `vlan-only`.
func (o NetworkOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Purpose }).(pulumi.StringOutput)
}

// The name of the site to associate the network with.
func (o NetworkOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

// The subnet of the network. Must be a valid CIDR address.
func (o NetworkOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.Subnet }).(pulumi.StringPtrOutput)
}

// The VLAN ID of the network.
func (o NetworkOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.IntPtrOutput { return v.VlanId }).(pulumi.IntPtrOutput)
}

// DNS servers IPs of the WAN.
func (o NetworkOutput) WanDns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Network) pulumi.StringArrayOutput { return v.WanDns }).(pulumi.StringArrayOutput)
}

// Specifies the WAN egress quality of service. Defaults to `0`.
func (o NetworkOutput) WanEgressQos() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.IntPtrOutput { return v.WanEgressQos }).(pulumi.IntPtrOutput)
}

// The IPv4 gateway of the WAN.
func (o NetworkOutput) WanGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.WanGateway }).(pulumi.StringPtrOutput)
}

// The IPv4 address of the WAN.
func (o NetworkOutput) WanIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.WanIp }).(pulumi.StringPtrOutput)
}

// The IPv4 netmask of the WAN.
func (o NetworkOutput) WanNetmask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.WanNetmask }).(pulumi.StringPtrOutput)
}

// Specifies the WAN network group. Must be one of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
func (o NetworkOutput) WanNetworkgroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.WanNetworkgroup }).(pulumi.StringPtrOutput)
}

// Specifies the IPV4 WAN connection type. Must be one of either `disabled`, `static`, `dhcp`, or `pppoe`.
func (o NetworkOutput) WanType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.WanType }).(pulumi.StringPtrOutput)
}

// Specifies the IPV4 WAN username.
func (o NetworkOutput) WanUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.WanUsername }).(pulumi.StringPtrOutput)
}

// Specifies the IPV4 WAN password.
func (o NetworkOutput) XWanPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.XWanPassword }).(pulumi.StringPtrOutput)
}

type NetworkArrayOutput struct{ *pulumi.OutputState }

func (NetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Network)(nil)).Elem()
}

func (o NetworkArrayOutput) ToNetworkArrayOutput() NetworkArrayOutput {
	return o
}

func (o NetworkArrayOutput) ToNetworkArrayOutputWithContext(ctx context.Context) NetworkArrayOutput {
	return o
}

func (o NetworkArrayOutput) Index(i pulumi.IntInput) NetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Network {
		return vs[0].([]*Network)[vs[1].(int)]
	}).(NetworkOutput)
}

type NetworkMapOutput struct{ *pulumi.OutputState }

func (NetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Network)(nil)).Elem()
}

func (o NetworkMapOutput) ToNetworkMapOutput() NetworkMapOutput {
	return o
}

func (o NetworkMapOutput) ToNetworkMapOutputWithContext(ctx context.Context) NetworkMapOutput {
	return o
}

func (o NetworkMapOutput) MapIndex(k pulumi.StringInput) NetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Network {
		return vs[0].(map[string]*Network)[vs[1].(string)]
	}).(NetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInput)(nil)).Elem(), &Network{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkArrayInput)(nil)).Elem(), NetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkMapInput)(nil)).Elem(), NetworkMap{})
	pulumi.RegisterOutputType(NetworkOutput{})
	pulumi.RegisterOutputType(NetworkArrayOutput{})
	pulumi.RegisterOutputType(NetworkMapOutput{})
}
