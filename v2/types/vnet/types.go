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

package vnet

/*
VNet Structure for GET requests
*/

type VNet struct {
	CreatedDate  int                `json:"createdDate"`
	Gateways     []VNetGateway      `json:"gateways"`
	ID           int                `json:"id"`
	Internal     int                `json:"internal"`
	IPFamily     string             `json:"ipFamily"`
	MacAddress   string             `json:"macAddress"`
	ModifiedDate int                `json:"modifiedDate"`
	Name         string             `json:"name"`
	NativeVlan   int                `json:"nativeVlan"`
	PortsCount   int                `json:"portsCount"`
	Provisioning bool               `json:"provisioning"`
	Ports        []VNetDetailedPort `json:"ports"`
	State        string             `json:"state"`
	Sites        []VNetDetailedSite `json:"sites"`
	Status       VNetStatus         `json:"status"`
	VlanAware    bool               `json:"vlanAware"`
	Vlans        string             `json:"vlans"`
	VxlanID      int                `json:"vxlanID"`
	Tenant       IDName             `json:"tenant"`
	TenantID     int                `json:"tenantID"`
	Vlan         int                `json:"vlan"`
	Tags         []string           `json:"tags"`
	PortTags     []VNetPortTag      `json:"portTags"`
	Vpc          IDName             `json:"vpc"`
	DhcpRelay    *VNetDhcpRelay     `json:"dhcpRelay"`
	Dhcpv6Relay  *VNetDhcpv6Relay   `json:"dhcpv6Relay"`
	IPv6ND       *VNetIPv6ND        `json:"ipv6ND"`
}

type IDName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

/*
VNetDhcpRelay configures DHCP Relay for a V-Net. Enabling DHCP Relay disables
DHCP configuration under Gateways.
*/
type VNetDhcpRelay struct {
	Enabled       bool    `json:"enabled"`
	Vpc           *IDName `json:"vpc"`
	PrimaryAddr   *string `json:"primaryAddr"`
	SecondaryAddr *string `json:"secondaryAddr"`
}

/*
VNetDhcpv6Relay configures DHCPv6 Relay for a V-Net.
*/
type VNetDhcpv6Relay struct {
	Enabled       bool    `json:"enabled"`
	Vpc           *IDName `json:"vpc"`
	PrimaryAddr   *string `json:"primaryAddr"`
	SecondaryAddr *string `json:"secondaryAddr"`
}

type VNetIPv6ND struct {
	RouterAdvertisement *VNetIPv6NDRouterAdvertisement `json:"routerAdvertisement"`
	PrefixAdvertisement *VNetIPv6NDPrefixAdvertisement `json:"prefixAdvertisement"`
	RDNSS               *VNetIPv6NDRDNSS               `json:"rdnss"`
}

type VNetIPv6NDSeconds struct {
	Seconds *int `json:"seconds"`
}

type VNetIPv6NDLifetime struct {
	Seconds  *int `json:"seconds"`
	Infinite bool `json:"infinite"`
}

type VNetIPv6NDRouterAdvertisement struct {
	Mode                  string             `json:"mode"`
	RouterLifetime        *VNetIPv6NDSeconds `json:"routerLifetime"`
	AdvertisementInterval *VNetIPv6NDSeconds `json:"advertisementInterval"`
	ManagedConfig         bool               `json:"managedConfig"`
	OtherConfig           bool               `json:"otherConfig"`
}

type VNetIPv6NDPrefixAdvertisement struct {
	Enabled           bool               `json:"enabled"`
	PreferredLifetime *VNetIPv6NDSeconds `json:"preferredLifetime"`
	ValidLifetime     *VNetIPv6NDSeconds `json:"validLifetime"`
	Autoconfig        bool               `json:"autoconfig"`
}

type VNetIPv6NDRDNSS struct {
	Enabled    bool                `json:"enabled"`
	DNSServers []string            `json:"dnsServers"`
	Lifetime   *VNetIPv6NDLifetime `json:"lifetime"`
}

type VNetGateway struct {
	DHCP           VNetGatewayDHCP `json:"dhcp"`
	DHCPEnabled    bool            `json:"dhcpEnabled"`
	DHCPLeaseCount int             `json:"dhcpLeaseCount"`
	IPFamily       string          `json:"ipFamily"`
	Prefix         string          `json:"prefix"`
	Vlan           string          `json:"vlan"`
}

type VNetPortTag struct {
	Name       string `json:"name"`
	AccessMode bool   `json:"accessMode"`
}

type VNetGatewayDHCP struct {
	End       string `json:"end"`
	OptionSet IDName `json:"optionSet"`
	Start     string `json:"start"`
}

type VNetStatus struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

/*
VNet Structure for GET by id request
*/

type VNetDetailed struct {
	CreatedDate  int                       `json:"createdDate"`
	Gateways     []VNetDetailedGateway     `json:"gateways"`
	GuestTenants []VNetDetailedGuestTenant `json:"guestTenants"`
	ID           int                       `json:"id"`
	Internal     int                       `json:"internal"`
	IPFamily     string                    `json:"ipFamily"`
	MacAddress   string                    `json:"macAddress"`
	ModifiedDate int                       `json:"modifiedDate"`
	Name         string                    `json:"name"`
	NativeVlan   int                       `json:"nativeVlan"`
	Ports        []VNetDetailedPort        `json:"ports"`
	PortsCount   int                       `json:"portsCount"`
	Provisioning bool                      `json:"provisioning"`
	Sites        []VNetDetailedSite        `json:"sites"`
	State        string                    `json:"state"`
	Status       VNetStatus                `json:"status"`
	Tenant       VNetDetailedTenant        `json:"tenant"`
	VlanAware    bool                      `json:"vlanAware"`
	Vlans        string                    `json:"vlans"`
	VxlanID      int                       `json:"vxlanID"`
	Vlan         int                       `json:"vlan"`
	Tags         []string                  `json:"tags"`
	PortTags     []VNetPortTag             `json:"portTags"`
	Vpc          IDName                    `json:"vpc"`
	DhcpRelay    *VNetDhcpRelay            `json:"dhcpRelay"`
	Dhcpv6Relay  *VNetDhcpv6Relay          `json:"dhcpv6Relay"`
	IPv6ND       *VNetIPv6ND               `json:"ipv6ND"`
}

type VNetDetailedGateway struct {
	DHCP           *VNetGatewayDHCP `json:"dhcp"`
	DHCPEnabled    bool             `json:"dhcpEnabled"`
	DHCPLeaseCount int              `json:"dhcpLeaseCount"`
	Prefix         string           `json:"prefix"`
	Vlan           string           `json:"vlan"`
}

type VNetDetailedGuestTenant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type VNetDetailedPortInfo struct {
	IfName string `json:"ifName"`
	Port   string `json:"port"`
}

type VNetDetailedSite struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type VNetDetailedTenant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type VNetDetailedPort struct {
	Access           bool                             `json:"access"`
	AccessMode       bool                             `json:"accessMode"`
	AdminDown        string                           `json:"adminDown"`
	AutoNeg          string                           `json:"autoNeg"`
	Breakout         string                           `json:"breakout"`
	CreatedDate      int                              `json:"createdDate"`
	Description      string                           `json:"description"`
	DesiredSpeed     interface{}                      `json:"desiredSpeed"`
	Duplex           string                           `json:"duplex"`
	Extension        int                              `json:"extension"`
	ID               int                              `json:"id"`
	IfName           interface{}                      `json:"ifName"`
	Lacp             string                           `json:"lacp"`
	MacCount         int                              `json:"macCount"`
	ModifiedDate     int                              `json:"modifiedDate"`
	Mtu              int                              `json:"mtu"`
	Name             string                           `json:"name"`
	ParentPort       int                              `json:"parentPort"`
	Port             string                           `json:"port"`
	Site             VNetDetailedPortSite             `json:"site"`
	SlavePorts       []interface{}                    `json:"slavePorts"`
	Speed            string                           `json:"speed"`
	State            VNetDetailedPortState            `json:"state"`
	StateInHierarchy VNetDetailedPortStateInHierarchy `json:"stateInHierarchy"`
	Status           VNetStatus                       `json:"status"`
	Switch           VNetDetailedPortSwitch           `json:"switch"`
	SwitchName       string                           `json:"switchName"`
	Tenant           VNetDetailedPortTenant           `json:"tenant"`
	Transceiver      string                           `json:"transceiver"`
	Untagged         bool                             `json:"untagged"`
	Used             bool                             `json:"used"`
	Vlan             string                           `json:"vlan"`
	Info             VNetDetailedPortInfo             `json:"info"`
}

type VNetDetailedPortSite struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type VNetDetailedPortState struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}

type VNetDetailedPortStateInHierarchy struct {
	Aggregated       bool `json:"aggregated"`
	Breakout         bool `json:"breakout"`
	BreakoutChild    int  `json:"breakoutChild"`
	Extended         int  `json:"extended"`
	ExtensionsParent int  `json:"extensionsParent"`
	LagMember        bool `json:"lagMember"`
}

type VNetDetailedPortSwitch struct {
	ID          int    `json:"id"`
	MainAddress string `json:"mainAddress"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type VNetDetailedPortTenant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

/*
VNet Structure for POST requests
*/

type VNetAdd struct {
	Name         string           `json:"name"`
	Sites        []VNetAddSite    `json:"sites"`
	Gateways     []VNetAddGateway `json:"gateways"`
	GuestTenants []VNetAddTenant  `json:"guestTenants"`
	IPFamily     string           `json:"ipFamily"`
	NativeVlan   int              `json:"nativeVlan"`
	Ports        []VNetAddPort    `json:"ports"`
	Provisioning bool             `json:"provisioning"`
	State        string           `json:"state"`
	Tenant       VNetAddTenant    `json:"tenant"`
	VlanAware    bool             `json:"vlanAware"`
	Vlans        string           `json:"vlans"`
	Vlan         interface{}      `json:"vlan"`
	Tags         []string         `json:"tags"`
	PortTags     []VNetPortTag    `json:"portTags"`
	VxlanID      int              `json:"vxlanID"`
	Vpc          *IDName          `json:"vpc,omitempty"`
	DhcpRelay    *VNetDhcpRelay   `json:"dhcpRelay,omitempty"`
	Dhcpv6Relay  *VNetDhcpv6Relay `json:"dhcpv6Relay,omitempty"`
	IPv6ND       *VNetIPv6ND      `json:"ipv6ND,omitempty"`
}

type VNetAddGateway struct {
	DHCP           *VNetGatewayDHCP `json:"dhcp"`
	DHCPEnabled    bool             `json:"dhcpEnabled"`
	DHCPLeaseCount int              `json:"dhcpLeaseCount"`
	Prefix         string           `json:"prefix"`
	Vlan           string           `json:"vlan"`
}

type VNetAddTenant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type VNetAddPort struct {
	Name       string `json:"name"`
	Access     bool   `json:"access"`
	AccessMode bool   `json:"accessMode"`
	ID         int    `json:"id"`
	Lacp       string `json:"lacp"`
	State      string `json:"state"`
	Vlan       string `json:"vlan"`
}

type VNetAddSite struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

/*
VNet Structure for PUT requests
*/

type VNetUpdate struct {
	Gateways     []VNetUpdateGateway     `json:"gateways"`
	GuestTenants []VNetUpdateGuestTenant `json:"guestTenants"`
	IPFamily     string                  `json:"ipFamily"`
	Name         string                  `json:"name"`
	NativeVlan   int                     `json:"nativeVlan"`
	Ports        []VNetUpdatePort        `json:"ports"`
	Provisioning bool                    `json:"provisioning"`
	Sites        []VNetUpdateSite        `json:"sites"`
	State        string                  `json:"state"`
	Vlans        string                  `json:"vlans"`
	Vlan         interface{}             `json:"vlan"`
	Tags         []string                `json:"tags"`
	PortTags     []VNetPortTag           `json:"portTags"`
	VxlanID      int                     `json:"vxlanID"`
	DhcpRelay    *VNetDhcpRelay          `json:"dhcpRelay,omitempty"`
	Dhcpv6Relay  *VNetDhcpv6Relay        `json:"dhcpv6Relay,omitempty"`
	IPv6ND       *VNetIPv6ND             `json:"ipv6ND,omitempty"`
}

type VNetUpdateGateway struct {
	DHCP           *VNetGatewayDHCP `json:"dhcp"`
	DHCPEnabled    bool             `json:"dhcpEnabled"`
	DHCPLeaseCount int              `json:"dhcpLeaseCount"`
	Prefix         string           `json:"prefix"`
	Vlan           string           `json:"vlan"`
}

type VNetUpdateGuestTenant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type VNetUpdatePort struct {
	Access     bool   `json:"access"`
	AccessMode bool   `json:"accessMode"`
	Name       string `json:"name"`
	ID         int    `json:"id"`
	Lacp       string `json:"lacp"`
	State      string `json:"state"`
	Vlan       string `json:"vlan"`
	Maccount   int    `json:"macCount"`
	Untagged   bool   `json:"untagged"`
}

type VNetUpdateSite struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
