/*
Copyright 2021. Netris, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v2

const (
	Auth = "/api/auth"

	L4LB = "/api/v2/l4lb"

	VNetBase      = "/api/v2/vnet"
	VNetUnmanaged = "/api/v2/vnet-unmanaged"

	InventoryBase         = "/api/v2/hw"
	InventorySwitch       = "/api/v2/hw/switch"
	InventoryEquinixMetal = "/api/v2/hw/equinix-metal-server"
	InventoryController   = "/api/v2/hw/controller"
	InventorySoftgate     = "/api/v2/hw/softgate"
	InventoryUsedIPs      = "/api/v2/hw/usedips"
	InventoryProfiles     = "/api/v2/hw/profiles"
	InventorySubnets      = "/api/v2/hw/subnets"
	InventoryNOS          = "/api/v2/hw/nos"

	Links = "/api/v2/link"

	IPAMBase       = "/api/v2/ipam"
	IPAMSubnets    = "/api/v2/ipam/subnets"
	IPAMHosts      = "/api/v2/ipam/hosts"
	IPAMAllocation = "/api/v2/ipam/allocation"
	IPAMSubnet     = "/api/v2/ipam/subnet"

	Ports          = "/api/v2/ports"
	PortExtensions = "/api/v2/ports/extensions"
	PortLAG        = "/api/v2/ports/lag"
	PortFreeUP     = "/api/v2/ports/freeup"

	BGP               = "/api/v2/ebgp"
	BGPSites          = "/api/v2/ebgp/sites"
	BGPVNets          = "/api/v2/ebgp/v-nets"
	BGPRouteMaps      = "/api/v2/ebgp/routemaps"
	BGPOffloaders     = "/api/v2/ebgp/offloaders"
	BGPPorts          = "/api/v2/ebgp/ports"
	BGPSwitches       = "/api/v2/ebgp/switches"
	BGPUpdatedSources = "/api/v2/ebgp/updatesources"

	ROH = "/api/v2/roh"

	NAT = "/api/v2/nat"

	DHCP = "/api/v2/dhcp-option-set"

	VlanReservation = "/api/v2/reservation/vlan"

	Sites = "/api/v2/sites"
)
