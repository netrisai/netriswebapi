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
	ID   int    `json:"id"`
	Name string `json:"name"`
}

/*
VNet Structure for POST requests
*/

type HWAdd struct {
	Asn         string      `json:"asn"`
	Description string      `json:"description"`
	Links       []HWAddLink `json:"links"`
	MacAddress  string      `json:"macAddress"`
	MainAddress string      `json:"mainAddress"`
	MgmtAddress string      `json:"mgmtAddress"`
	Name        string      `json:"name"`
	Nos         string      `json:"nos"`
	PortCount   int         `json:"portCount"`
	Profile     IDName      `json:"profile"`
	Site        IDName      `json:"site"`
	Tenant      IDName      `json:"tenant"`
}

type HWAddLink struct {
	Local  IDName `json:"local"`
	Remote IDName `json:"remote"`
}
