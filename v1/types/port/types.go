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

package port

// Port .
type Port struct {
	OneG                     int         `json:"1g"`
	AdminDown                int         `json:"admin_down"`
	ID                       int         `json:"id"`
	Lacp                     string      `json:"lacp"`
	MemberState              string      `json:"member_state"`
	MonitoredPort            int         `json:"monitored_port"`
	Mtu                      int         `json:"mtu"`
	Name                     string      `json:"name"`
	Number                   int         `json:"number"`
	ParentPort               int         `json:"parent_port"`
	Port                     string      `json:"port"`
	PortName                 string      `json:"portName"`
	PortNameFull             string      `json:"port_name"`
	PortNameForShowAndExport string      `json:"portNameForShowAndExport"`
	PortExtension            int         `json:"port_extension"`
	PortID                   int         `json:"port_id"`
	PortType                 string      `json:"port_type"`
	Power                    interface{} `json:"power"`
	Present                  int         `json:"present"`
	Price                    int         `json:"price"`
	ProviderID               int         `json:"provider_id"`
	SiteID                   int         `json:"site_id"`
	SiteName                 string      `json:"site_name"`
	SlavePortName            string      `json:"slavePortName"`
	SlavePorts               []Port      `json:"slavePorts"`
	SlavePortID              int         `json:"slave_port_id"`
	SlavePortIds             string      `json:"slave_port_ids"`
	Speed                    string      `json:"speed"`
	Status                   string      `json:"status"`
	Switch                   string      `json:"switch"`
	SwitchID                 int         `json:"switch_id"`
	SwitchName               string      `json:"switch_name"`
	SwitchType               string      `json:"switch_type"`
	Tenant                   string      `json:"tenant"`
	TenantID                 int         `json:"tenant_id"`
	Transceiver              string      `json:"transceiver"`
	Type                     string      `json:"type"`
	VlanFrom                 int         `json:"vlan_from"`
	VlanTo                   int         `json:"vlan_to"`
}
