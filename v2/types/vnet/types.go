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
	MacAddress   string             `json:"macAddress"`
	ModifiedDate int                `json:"modifiedDate"`
	Name         string             `json:"name"`
	NativeVlan   int                `json:"nativeVlan"`
	PortsCount   int                `json:"portsCount"`
	Provisioning bool               `json:"provisioning"`
	State        string             `json:"state"`
	Sites        []VNetDetailedSite `json:"sites"`
	Status       VNetStatus         `json:"status"`
	VlanAware    bool               `json:"vlanAware"`
	Vlans        string             `json:"vlans"`
	VxlanID      int                `json:"vxlanID"`
	Tenant       IDName             `json:"tenant"`
	TenantID     int                `json:"tenantID"`
	Vlan         int                `json:"vlan"`
}

type IDName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type VNetGateway struct {
	DHCP           VNetGatewayDHCP `json:"dhcp"`
	DHCPEnabled    bool            `json:"dhcpEnabled"`
	DHCPLeaseCount int             `json:"dhcpLeaseCount"`
	IPFamily       string          `json:"ipFamily"`
	Prefix         string          `json:"prefix"`
	Vlan           string          `json:"vlan"`
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
}

type VNetDetailedGateway struct {
	DHCP           VNetGatewayDHCP `json:"dhcp"`
	DHCPEnabled    bool            `json:"dhcpEnabled"`
	DHCPLeaseCount int             `json:"dhcpLeaseCount"`
	Prefix         string          `json:"prefix"`
	Vlan           string          `json:"vlan"`
}

type VNetDetailedGuestTenant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
	NativeVlan   int              `json:"nativeVlan"`
	Ports        []VNetAddPort    `json:"ports"`
	Provisioning bool             `json:"provisioning"`
	State        string           `json:"state"`
	Tenant       VNetAddTenant    `json:"tenant"`
	VlanAware    bool             `json:"vlanAware"`
	Vlans        string           `json:"vlans"`
	Vlan         interface{}      `json:"vlan"`
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
	Name   string `json:"name"`
	Access bool   `json:"access"`
	ID     int    `json:"id"`
	Lacp   string `json:"lacp"`
	State  string `json:"state"`
	Vlan   string `json:"vlan"`
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
	Name         string                  `json:"name"`
	NativeVlan   int                     `json:"nativeVlan"`
	Ports        []VNetUpdatePort        `json:"ports"`
	Provisioning bool                    `json:"provisioning"`
	Sites        []VNetUpdateSite        `json:"sites"`
	State        string                  `json:"state"`
	Vlans        string                  `json:"vlans"`
	Vlan         interface{}             `json:"vlan"`
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
	Access   bool   `json:"access"`
	Name     string `json:"name"`
	ID       int    `json:"id"`
	Lacp     string `json:"lacp"`
	State    string `json:"state"`
	Vlan     string `json:"vlan"`
	Maccount int    `json:"macCount"`
}

type VNetUpdateSite struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
