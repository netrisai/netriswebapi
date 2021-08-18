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

package inventory

// InventoryAdd .
type InventoryAdd struct {
	AddAnycast         []interface{} `json:"add_anycast"`
	DelAnycast         []interface{} `json:"del_anycast"`
	DelPool            []interface{} `json:"del_pool"`
	Description        string        `json:"description"`
	ID                 int           `json:"id"`
	IPAddress          string        `json:"ip_address"`
	IPAddressLength    string        `json:"ip_address_length"`
	MacAddress         string        `json:"mac_address"`
	MngIPAddress       string        `json:"mng_ip_address"`
	MngIPAddressLength string        `json:"mng_ip_address_length"`
	MngIPAddressSubnet string        `json:"mng_ip_address_subnet"`
	Name               string        `json:"name"`
	NumberOfPorts      int           `json:"number_of_ports"`
	Pool               []interface{} `json:"pool"`
	OwnerTenantID      int           `json:"owner_tenant_id"`
	ProfileID          int           `json:"profile_id"`
	SiteID             int           `json:"site_id"`
	Status             string        `json:"status"`
	Type               string        `json:"type"`
	Ztp                bool          `json:"ztp"`
	NOS                string        `json:"nos"`
	NFVSwitch          string        `json:"nfv_switch"`
	OffloadPortID      int           `json:"offload_port_id"`
}

// Inventory .
type Inventory struct {
	Anycast          int                       `json:"anycast"`
	AnycastAddresses []InventoryAnycastAddress `json:"anycast_addresses"`
	ConductorVersion string                    `json:"conductor_version"`
	CreatedDate      int                       `json:"created_date"`
	CumulusVersion   string                    `json:"cumulus_version"`
	Description      string                    `json:"description"`
	DNS              string                    `json:"dns"`
	EgressAcls       int                       `json:"egress_acls"`
	GraphData        []InventoryGraphData      `json:"graph_data"`
	HaproxyVersion   interface{}               `json:"haproxy_version"`
	HardwareHealth   []InventoryHardwareHealth `json:"hardwareHealth"`
	Health           struct {
		Critical int `json:"critical"`
		Ok       int `json:"ok"`
		Unknown  int `json:"unknown"`
		Warning  int `json:"warning"`
	} `json:"health"`
	Heartbeat        int `json:"heartbeat"`
	ID               int `json:"id"`
	IngressAcls      int `json:"ingress_acls"`
	InstanceID       int `json:"instance_id"`
	InventoryProfile struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"inventoryProfile"`
	IPAddrID          int         `json:"ip_addr_id"`
	IPAddress         string      `json:"ip_address"`
	IsAnycastUsed     int         `json:"is_anycast_used"`
	LastSeen          int         `json:"last_seen"`
	MacAddress        string      `json:"mac_address"`
	Macs              int         `json:"macs"`
	MngIPAddrID       int         `json:"mng_ip_addr_id"`
	MngIPAddress      string      `json:"mng_ip_address"`
	ModifiedDate      int         `json:"modified_date"`
	Nos               string      `json:"nos"`
	NosName           string      `json:"nosName"`
	NTP               string      `json:"ntp"`
	NumberOfPorts     int         `json:"number_of_ports"`
	OffloadPortID     int         `json:"offload_port_id"`
	OffloaderPort     string      `json:"offloader_port"`
	OffloaderPortName string      `json:"offloader_port_name"`
	OpenstackVersion  interface{} `json:"openstack_version"`
	OwnerTenantID     int         `json:"owner_tenant_id"`
	OwnerTenantName   string      `json:"owner_tenant_name"`
	Password          interface{} `json:"password"`
	Platform          string      `json:"platform"`
	Port              string      `json:"port"`
	PortSwName        string      `json:"port_sw_name"`
	ProfileID         int         `json:"profile_id"`
	ProfileName       string      `json:"profile_name"`
	ProxmoxVersion    string      `json:"proxmox_version"`
	Rangecut          int         `json:"rangecut"`
	Routes            int         `json:"routes"`
	SiteID            int         `json:"site_id"`
	SiteName          string      `json:"site_name"`
	Status            string      `json:"status"`
	SubnetLength      int         `json:"subnet_length"`
	SwitchID          int         `json:"switch_id"`
	SwitchName        string      `json:"switch_name"`
	Timezone          string      `json:"timezone"`
	Type              string      `json:"type"`
	Uptime            string      `json:"uptime"`
	Username          interface{} `json:"username"`
	Vpn               string      `json:"vpn"`
	Ztp               int         `json:"ztp"`
}

// InventoryAnycastAddress .
type InventoryAnycastAddress struct {
	Address string `json:"address"`
	Pool    int    `json:"pool"`
	Used    int    `json:"used"`
}

// InventoryGraphData .
type InventoryGraphData struct {
	Interface string `json:"interface"`
	Switch    string `json:"switch"`
}

// InventoryHardwareHealth .
type InventoryHardwareHealth struct {
	CheckName        string                            `json:"check_name"`
	CriticalMessages []string                          `json:"critical_messages"`
	CurrentCriticals []string                          `json:"current_criticals"`
	Date             int                               `json:"date"`
	HwID             int                               `json:"hw_id"`
	HwName           string                            `json:"hw_name"`
	HwType           string                            `json:"hw_type"`
	LastCheckTime    int                               `json:"last_check_time"`
	Message          string                            `json:"message"`
	OkBriefMessage   string                            `json:"ok_brief_message"`
	OkMessages       []string                          `json:"ok_messages"`
	OksCount         int                               `json:"oks_count"`
	Port             string                            `json:"port"`
	PortID           int                               `json:"port_id"`
	PortStatus       string                            `json:"port_status"`
	Properties       []InventoryHardwareHealthProperty `json:"properties"`
	Recover          bool                              `json:"recover"`
	Severity         string                            `json:"severity"`
	Status           string                            `json:"status"`
	TenantID         int                               `json:"tenant_id"`
}

// InventoryHardwareHealthProperty .
type InventoryHardwareHealthProperty struct {
	Args       string `json:"Args"`
	BThreshold int    `json:"BThreshold"`
	CHigh      int    `json:"CHigh"`
	CLow       int    `json:"CLow"`
	CheckName  string `json:"CheckName"`
	Custom     bool   `json:"Custom"`
	Message    string `json:"Message"`
	PrevStatus string `json:"PrevStatus"`
	Property   string `json:"Property"`
	Status     string `json:"Status"`
	Unit       string `json:"Unit"`
	Use        bool   `json:"Use"`
	Value      int    `json:"Value"`
	WHigh      int    `json:"WHigh"`
	WLow       int    `json:"WLow"`
}

// InventoryUsedIPs .
type InventoryUsedIPs []string
