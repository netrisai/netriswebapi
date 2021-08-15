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
	Asn              int64         `json:"asn"`
	AsnID            int64         `json:"asnID"`
	AsnNumber        HWASNNumber   `json:"asnNumber"`
	ConductorVersion string        `json:"conductorVersion"`
	CreatedDate      int64         `json:"createdDate"`
	CumulusVersion   interface{}   `json:"cumulusVersion"`
	Description      string        `json:"description"`
	HardwareHealth   []interface{} `json:"hardwareHealth"`
	Health           struct{}      `json:"health"`
	ID               int64         `json:"id"`
	IngressAcls      int64         `json:"ingressAcls"`
	LastSeen         int64         `json:"lastSeen"`
	Links            []interface{} `json:"links"`
	MacAddrUbit      int64         `json:"macAddrUbit"`
	MacAddress       string        `json:"macAddress"`
	Macs             int64         `json:"macs"`
	MainAddress      string        `json:"mainAddress"`
	MainIP           HWMainIP      `json:"mainIP"`
	MaintenanceMode  bool          `json:"maintenanceMode"`
	MgmtAddress      string        `json:"mgmtAddress"`
	MgmtIP           HWMgmtIP      `json:"mgmtIP"`
	ModifiedDate     int64         `json:"modifiedDate"`
	Name             string        `json:"name"`
	Nos              HWNOS         `json:"nos"`
	OpenstackVersion interface{}   `json:"openstackVersion"`
	Platform         string        `json:"platform"`
	PortCount        int64         `json:"portCount"`
	Profile          HWProfile     `json:"profile"`
	ProxmoxVersion   interface{}   `json:"proxmoxVersion"`
	Rangecut         int64         `json:"rangecut"`
	Routes           int64         `json:"routes"`
	Site             HWSite        `json:"site"`
	Status           string        `json:"status"`
	Tenant           HWTenant      `json:"tenant"`
	Timezone         string        `json:"timezone"`
	Type             string        `json:"type"`
	Uptime           string        `json:"uptime"`
}

type HWASNNumber struct {
	Asn int64 `json:"asn"`
	ID  int64 `json:"id"`
}

type HWMainIP struct {
	Address     string       `json:"address"`
	Description interface{}  `json:"description"`
	ID          int64        `json:"id"`
	IPFamily    string       `json:"ipFamily"`
	Meta        HWMainIPMeta `json:"meta"`
	Name        string       `json:"name"`
	Readonly    string       `json:"readonly"`
	SubnetID    int64        `json:"subnetID"`
	Type        string       `json:"type"`
}

type HWMainIPMeta struct {
	Tag string `json:"tag"`
}

type HWMgmtIP struct {
	Address     string       `json:"address"`
	Description interface{}  `json:"description"`
	ID          int64        `json:"id"`
	IPFamily    string       `json:"ipFamily"`
	Meta        HWMgmtIPMeta `json:"meta"`
	Name        string       `json:"name"`
	Readonly    string       `json:"readonly"`
	SubnetID    int64        `json:"subnetID"`
	Type        string       `json:"type"`
}

type HWMgmtIPMeta struct {
	Tag string `json:"tag"`
}

type HWNOS struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type HWProfile struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type HWSite struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type HWTenant struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
