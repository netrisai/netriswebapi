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

/*
Inventory Structure for GET requests
*/

type HW struct {
	Asn              int           `json:"asn"`
	AsnID            int           `json:"asnID"`
	AsnNumber        HWASNNumber   `json:"asnNumber"`
	ConductorVersion string        `json:"conductorVersion"`
	CreatedDate      int           `json:"createdDate"`
	CumulusVersion   interface{}   `json:"cumulusVersion"`
	Description      string        `json:"description"`
	HardwareHealth   []interface{} `json:"hardwareHealth"`
	Health           struct{}      `json:"health"`
	ID               int           `json:"id"`
	IngressAcls      int           `json:"ingressAcls"`
	LastSeen         int           `json:"lastSeen"`
	Links            []interface{} `json:"links"`
	MacAddrUbit      int           `json:"macAddrUbit"`
	MacAddress       string        `json:"macAddress"`
	Macs             int           `json:"macs"`
	MainAddress      string        `json:"mainAddress"`
	MainIP           HWMainIP      `json:"mainIP"`
	MaintenanceMode  bool          `json:"maintenanceMode"`
	MgmtAddress      string        `json:"mgmtAddress"`
	MgmtIP           HWMgmtIP      `json:"mgmtIP"`
	ModifiedDate     int           `json:"modifiedDate"`
	Name             string        `json:"name"`
	Nos              HWNOS         `json:"nos"`
	OpenstackVersion interface{}   `json:"openstackVersion"`
	Platform         string        `json:"platform"`
	PortCount        int           `json:"portCount"`
	Profile          IDName        `json:"profile"`
	ProxmoxVersion   interface{}   `json:"proxmoxVersion"`
	Rangecut         int           `json:"rangecut"`
	Routes           int           `json:"routes"`
	Site             IDName        `json:"site"`
	Status           string        `json:"status"`
	Tenant           IDName        `json:"tenant"`
	Timezone         string        `json:"timezone"`
	Type             string        `json:"type"`
	Uptime           string        `json:"uptime"`
}

type HWASNNumber struct {
	Asn int `json:"asn"`
	ID  int `json:"id"`
}

type HWMainIP struct {
	Address     string       `json:"address"`
	Description interface{}  `json:"description"`
	ID          int          `json:"id"`
	IPFamily    string       `json:"ipFamily"`
	Meta        HWMainIPMeta `json:"meta"`
	Name        string       `json:"name"`
	Readonly    string       `json:"readonly"`
	SubnetID    int          `json:"subnetID"`
	Type        string       `json:"type"`
}

type HWMainIPMeta struct {
	Tag string `json:"tag"`
}

type HWMgmtIP struct {
	Address     string       `json:"address"`
	Description interface{}  `json:"description"`
	ID          int          `json:"id"`
	IPFamily    string       `json:"ipFamily"`
	Meta        HWMgmtIPMeta `json:"meta"`
	Name        string       `json:"name"`
	Readonly    string       `json:"readonly"`
	SubnetID    int          `json:"subnetID"`
	Type        string       `json:"type"`
}

type HWMgmtIPMeta struct {
	Tag string `json:"tag"`
}

type HWNOS struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type IDName struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

/*
Inventory Switch Structure for POST requests
*/

type HWSwitchAdd struct {
	Asn         string   `json:"asn"`
	Description string   `json:"description"`
	Links       []HWLink `json:"links"`
	MacAddress  string   `json:"macAddress"`
	MainAddress string   `json:"mainAddress"`
	MgmtAddress string   `json:"mgmtAddress"`
	Name        string   `json:"name"`
	Nos         NOS      `json:"nos"`
	PortCount   int      `json:"portCount"`
	Profile     IDName   `json:"profile"`
	Site        IDName   `json:"site"`
	Tenant      IDName   `json:"tenant"`
	Type        string   `json:"type"`
}

type HWLink struct {
	Local  IDName `json:"local"`
	Remote IDName `json:"remote"`
}

/*
Inventory Switch Structure for PUT requests
*/

type HWSwitchUpdate struct {
	Asn             string   `json:"asn"`
	Description     string   `json:"description"`
	Links           []HWLink `json:"links"`
	MacAddress      string   `json:"macAddress"`
	MainAddress     string   `json:"mainAddress"`
	MaintenanceMode bool     `json:"maintenanceMode"`
	MgmtAddress     string   `json:"mgmtAddress"`
	Name            string   `json:"name"`
	Nos             NOS      `json:"nos"`
	Profile         IDName   `json:"profile"`
}

/*
Inventory Controller Structure for POST requests
*/

type HWController struct {
	Description string `json:"description"`
	MainAddress string `json:"mainAddress"`
	Name        string `json:"name"`
	Site        IDName `json:"site"`
	Tenant      IDName `json:"tenant"`
}

/*
Inventory Controller Structure for PUT requests
*/

type HWControllerUpdate struct {
	Description     string `json:"description"`
	MainAddress     string `json:"mainAddress"`
	MaintenanceMode bool   `json:"maintenanceMode"`
	Name            string `json:"name"`
}

/*
Inventory Softgate Structure for POST requests
*/

type HWSoftgate struct {
	Description string   `json:"description"`
	Links       []HWLink `json:"links"`
	MainAddress string   `json:"mainAddress"`
	MgmtAddress string   `json:"mgmtAddress"`
	Name        string   `json:"name"`
	Profile     IDName   `json:"profile"`
	Site        IDName   `json:"site"`
	Tenant      IDName   `json:"tenant"`
}

/*
Inventory Softgate Structure for PUT requests
*/

type HWSoftgateUpdate struct {
	Description string   `json:"description"`
	Links       []HWLink `json:"links"`
	MainAddress string   `json:"mainAddress"`
	MgmtAddress string   `json:"mgmtAddress"`
	Name        string   `json:"name"`
	Profile     IDName   `json:"profile"`
	Site        IDName   `json:"site"`
	Tenant      IDName   `json:"tenant"`
}

/*
Inventory additional Structures for GET requests
*/

type UsedIP struct {
	Address      string `json:"address"`
	ConsumerID   int    `json:"consumer_id"`
	ConsumerType string `json:"consumer_type"`
	HostID       int    `json:"host_id"`
	ID           int    `json:"id"`
	Meta         string `json:"meta"`
	Name         string `json:"name"`
	SubnetID     int    `json:"subnet_id"`
}

type Profile struct {
	CreatedDate  int           `json:"created_date"`
	CustomRules  []interface{} `json:"customRules"`
	Description  string        `json:"description"`
	DNSServers   string        `json:"dns_servers"`
	ID           int           `json:"id"`
	Ipv4SSH      string        `json:"ipv4_ssh"`
	Ipv6SSH      string        `json:"ipv6_ssh"`
	ModifiedDate int           `json:"modified_date"`
	Name         string        `json:"name"`
	NTPServers   string        `json:"ntp_servers"`
	Timezone     string        `json:"timezone"`
}

type HWSubnet struct {
	AllocationID   int      `json:"allocationID"`
	DefaultGateway string   `json:"defaultGateway"`
	Description    string   `json:"description"`
	ID             int      `json:"id"`
	IPFamily       string   `json:"ipFamily"`
	Name           string   `json:"name"`
	ParentID       int      `json:"parentID"`
	Prefix         string   `json:"prefix"`
	Purpose        string   `json:"purpose"`
	Readonly       string   `json:"readonly"`
	Sites          []IDName `json:"sites"`
	Subnet         []Subnet `json:"subnet"`
	Tenant         []IDName `json:"tenant"`
	Type           string   `json:"type"`
}

type Subnet struct {
	Length int    `json:"length"`
	Prefix string `json:"prefix"`
}

type NOS struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}
