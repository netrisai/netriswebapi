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

import (
	"encoding/json"
	"fmt"
)

// VNet .
type VNet struct {
	CreateDate   int           `json:"create_date"`
	ForSearch    string        `json:"forSearch"`
	Gateways     []VNetGateway `json:"gateways"`
	ID           string        `json:"id"`
	Internal     int           `json:"internal"`
	MacAddress   string        `json:"mac_address"`
	MembersCount int           `json:"membersCount"`
	ModifiedDate int           `json:"modified_date"`
	Name         string        `json:"name"`
	Owner        int           `json:"owner"`
	Provisioning int           `json:"provisioning"`
	Sites        string        `json:"sites"`
	SitesID      []int         `json:"sites_id"`
	State        string        `json:"state"`
	Tenants      string        `json:"tenants"`
	TenantsID    []int         `json:"tenants_id"`
	VaMode       int           `json:"va_mode"`
	VaNativeVLAN int           `json:"va_native_vlan"`
	VaVLANs      string        `json:"va_vlans"`
	VxLANID      int           `json:"vxlan_id"`
}

// VNetUpdate .
type VNetUpdate struct {
	Gateways     []VNetGateway `json:"gateways"`
	ID           int           `json:"id"`
	Members      string        `json:"members"`
	Name         string        `json:"name"`
	Owner        int           `json:"owner"`
	Provisioning int           `json:"provisioning"`
	Sites        []int         `json:"sites"`
	State        string        `json:"state"`
	Tenants      []int         `json:"tenants"`
	VaMode       bool          `json:"va_mode"`
	VaNativeVLAN string        `json:"va_native_vlan"`
	VaVLANs      string        `json:"va_vlans"`
}

// VNetInfo .
type VNetInfo struct {
	CircuitTenants          []int            `json:"circuitTenants"`
	CreateDate              int              `json:"create_date"`
	Gateways                []VNetGateway    `json:"gateways"`
	ID                      int              `json:"id"`
	MacAddress              string           `json:"mac_address"`
	McastHwAddressLastOctet int              `json:"mcast_hw_address_last_octet"`
	Members                 []VNetInfoMember `json:"members"`
	MembersCount            int              `json:"membersCount"`
	ModifiedDate            int              `json:"modified_date"`
	Name                    string           `json:"name"`
	Owner                   int              `json:"owner"`
	OwnerName               string           `json:"owner_name"`
	Provisioning            int              `json:"provisioning"`
	SiteName                string           `json:"site_name"`
	Sites                   string           `json:"sites"`
	SitesID                 []int            `json:"sites_id"`
	SitesName               []string         `json:"sites_name"`
	State                   string           `json:"state"`
	Tenants                 string           `json:"tenants"`
	TenantsID               []int            `json:"tenants_id"`
	TenantsName             []string         `json:"tenants_name"`
	VaMode                  int              `json:"va_mode"`
	VaNativeVlan            int              `json:"va_native_vlan"`
	VaVlans                 string           `json:"va_vlans"`
	VxlanID                 int              `json:"vxlan_id"`
}

// VNetInfoMember .
type VNetInfoMember struct {
	AdminDown          int         `json:"admin_down"`
	ChildPort          int         `json:"childPort"`
	ID                 int         `json:"id"`
	Lacp               string      `json:"lacp"`
	LastCheckTime      int         `json:"last_check_time"`
	MacCount           int         `json:"mac_count"`
	MemberInternal     int         `json:"member_internal"`
	MemberProvisioning int         `json:"member_provisioning"`
	MemberState        string      `json:"member_state"`
	Message            string      `json:"message"`
	Mtu                int         `json:"mtu"`
	Name               string      `json:"name"`
	ParentPort         int         `json:"parent_port"`
	Port               string      `json:"port"`
	PortAddDate        int         `json:"port_add_date"`
	PortID             int         `json:"port_id"`
	PortModDate        int         `json:"port_mod_date"`
	PortStatus         string      `json:"port_status"`
	PortTenantName     string      `json:"port_tenant_name"`
	Portname           string      `json:"portname"`
	SiteID             int         `json:"site_id"`
	SiteName           string      `json:"site_name"`
	SlavePortIds       string      `json:"slave_port_ids"`
	SlavePorts         interface{} `json:"slave_ports"`
	Speed              string      `json:"speed"`
	SwitchID           int         `json:"switch_id"`
	SwitchIPAddress    string      `json:"switch_ip_address"`
	SwitchName         string      `json:"switch_name"`
	TenantID           int         `json:"tenant_id"`
	VlanID             int         `json:"vlan_id"`
	VlansList          interface{} `json:"vlansList"`
}

// VNETSite .
type VNETSite struct {
	SiteID   int    `json:"site_id"`
	SiteName string `json:"site_name"`
}

// VNetGateway .
type VNetGateway struct {
	Gateway  string `json:"gateway"`
	GwLength int    `json:"gw_length"`
	ID       int    `json:"id,omitempty"`
	VaVLANID int    `json:"va_vlan_id,omitempty"`
	Nos      string `json:"nos,omitempty"`
	Version  string `json:"version,omitempty"`
}

// VNetAdd .
type VNetAdd struct {
	Gateways     []VNetGateway `json:"gateways"`
	ID           int           `json:"id"`
	Members      string        `json:"members"`
	Name         string        `json:"name"`
	Owner        int           `json:"owner"`
	Provisioning int           `json:"provisioning"`
	Sites        []int         `json:"sites"`
	State        string        `json:"state"`
	Tenants      []int         `json:"tenants"`
	VaMode       bool          `json:"va_mode"`
	VaNativeVLAN int           `json:"va_native_vlan"`
	VaVLANs      string        `json:"va_vlans"`
}

// VNetAddGateway .
type VNetAddGateway struct {
	Gateway  string      `json:"gateway"`
	GwLength string      `json:"gw_length"`
	Nos      interface{} `json:"nos"`
	Version  string      `json:"version"`
}

// VNetMember .
type VNetMember struct {
	ChildPort      int    `json:"childPort"`
	LACP           string `json:"lacp"`
	MemberState    string `json:"member_state"`
	ParentPort     int    `json:"parentPort"`
	PortIsUntagged bool   `json:"portIsUntagged"`
	PortID         int    `json:"port_id"`
	PortName       string `json:"port_name"`
	TenantID       int    `json:"tenant_id"`
	VLANID         int    `json:"vlan_id"`
}

// VNetMembers .
type VNetMembers struct {
	members []VNetMember
}

// VNetAddReply .
type VNetAddReply struct {
	CircuitID int `json:"circuitID"`
}

// Add .
func (m *VNetMembers) Add(member VNetMember) {
	m.members = append(m.members, member)
}

func (m *VNetMembers) String() string {
	js, err := json.Marshal(m.members)
	if err != nil {
		fmt.Println(err)
	}
	return string(js)
}
