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

package bgp

// EBGP .
type EBGP struct {
	AllowasIn            int    `json:"allowas_in"`
	BgpPassword          string `json:"bgp_password"`
	BgpPrefixes          string `json:"bgp_prefixes"`
	BgpState             string `json:"bgp_state"`
	BgpUptime            string `json:"bgp_uptime"`
	CircuitName          string `json:"circuitName"`
	CircuitInternal      int    `json:"circuit_internal"`
	Community            string `json:"community"`
	CreatedDate          int    `json:"created_date"`
	DefaultOriginate     string `json:"default_originate"`
	Description          string `json:"description"`
	GraphPort            string `json:"graphPort"`
	ID                   int    `json:"id"`
	InboundRouteMapName  string `json:"inboundRouteMapName"`
	InboundRouteMap      string `json:"inbound_route_map"`
	Internal             int    `json:"internal"`
	IPVersion            string `json:"ip_version"`
	KubenetInfo          string `json:"kubenet_info"`
	LocalAsn             int    `json:"localAsn"`
	LocalIP              string `json:"local_ip"`
	LocalPreference      int    `json:"local_preference"`
	ModifiedDate         int    `json:"modified_date"`
	Multihop             int    `json:"multihop"`
	Name                 string `json:"name"`
	NeighborAddress      string `json:"neighbor_address"`
	NeighborAs           int    `json:"neighbor_as"`
	NfvPortName          string `json:"nfv_port_name"`
	OffloadPortID        int    `json:"offload_port_id"`
	OffloadVlanID        int    `json:"offload_vlan_id"`
	OffloaderInterface   string `json:"offloader_interface"`
	OffloaderSwName      string `json:"offloader_sw_name"`
	Originate            string `json:"originate"`
	OutboundRouteMapName string `json:"outboundRouteMapName"`
	OutboundRouteMap     string `json:"outbound_route_map"`
	PName                string `json:"p_name"`
	Port                 string `json:"port"`
	PortName             string `json:"portName"`
	Portname             string `json:"port_name"`
	PortStatus           string `json:"port_status"`
	PrefixLength         int    `json:"prefix_length"`
	PrefixLimit          int    `json:"prefix_limit"`
	PrefixListInbound    string `json:"prefix_list_inbound"`
	PrefixListOutbound   string `json:"prefix_list_outbound"`
	PrependInbound       int    `json:"prepend_inbound"`
	PrependOutbound      int    `json:"prepend_outbound"`
	RcircuitID           int    `json:"rcircuit_id"`
	RemoteIP             string `json:"remote_ip"`
	RouteID              int    `json:"route_id"`
	SiteID               int    `json:"site_id"`
	SiteName             string `json:"site_name"`
	Status               string `json:"status"`
	SwitchID             int    `json:"switch_id"`
	SwitchName           string `json:"switch_name"`
	SwitchPortID         int    `json:"switch_port_id"`
	TermSwName           string `json:"term_sw_name"`
	TermSwitchID         int    `json:"term_switch_id"`
	TerminateOnSwitch    string `json:"terminate_on_switch"`
	UpdateSource         string `json:"update_source"`
	Vlan                 int    `json:"vlan"`
	Weight               int    `json:"weight"`
}

// EBGPAdd .
type EBGPAdd struct {
	AllowasIn          int     `json:"allowas_in"`
	BgpPassword        string  `json:"bgp_password"`
	Community          string  `json:"community"`
	Description        string  `json:"description"`
	InboundRouteMap    int     `json:"inboundRouteMap"`
	Internal           string  `json:"internal"`
	IPVersion          string  `json:"ip_version"`
	LocalIP            string  `json:"local_ip"`
	LocalPreference    int     `json:"local_preference"`
	Multihop           int     `json:"multihop"`
	Name               string  `json:"name"`
	NeighborAddress    *string `json:"neighbor_address"`
	NeighborAs         string  `json:"neighbor_as"`
	NfvID              int     `json:"nfv_id"`
	NfvPortID          int     `json:"nfv_port_id"`
	Originate          string  `json:"originate"`
	OutboundRouteMap   int     `json:"outboundRouteMap"`
	PrefixLength       int     `json:"prefix_length"`
	PrefixLimit        string  `json:"prefix_limit"`
	PrefixListInbound  string  `json:"prefix_list_inbound"`
	PrefixListOutbound string  `json:"prefix_list_outbound"`
	PrependInbound     int     `json:"prepend_inbound"`
	PrependOutbound    int     `json:"prepend_outbound"`
	RcircuitID         int     `json:"rcircuit_id"`
	RemoteIP           string  `json:"remote_ip"`
	SiteID             int     `json:"site_id"`
	Status             string  `json:"status"`
	SwitchID           int     `json:"switch_id"`
	SwitchName         string  `json:"switch_name"`
	SwitchPortID       int     `json:"switch_port_id"`
	TermSwitchID       int     `json:"term_switch_id"`
	TermSwitchName     string  `json:"term_switch_name"`
	TerminateOnSwitch  string  `json:"terminate_on_switch"`
	UpdateSource       string  `json:"update_source"`
	Vlan               int     `json:"vlan"`
	Weight             int     `json:"weight"`
}

// EBGPUpdate .
type EBGPUpdate struct {
	AllowasIn          int     `json:"allowas_in"`
	BgpPassword        string  `json:"bgp_password"`
	Community          string  `json:"community"`
	Description        string  `json:"description"`
	ID                 int     `json:"id"`
	InboundRouteMap    int     `json:"inboundRouteMap"`
	Internal           string  `json:"internal"`
	IPVersion          string  `json:"ip_version"`
	LocalIP            string  `json:"local_ip"`
	LocalPreference    int     `json:"local_preference"`
	Multihop           int     `json:"multihop"`
	Name               string  `json:"name"`
	NeighborAddress    *string `json:"neighbor_address"`
	NeighborAs         string  `json:"neighbor_as"`
	NfvID              int     `json:"nfv_id"`
	NfvPortID          int     `json:"nfv_port_id"`
	Originate          string  `json:"originate"`
	OutboundRouteMap   int     `json:"outboundRouteMap"`
	PrefixLength       int     `json:"prefix_length"`
	PrefixLimit        string  `json:"prefix_limit"`
	PrefixListInbound  string  `json:"prefix_list_inbound"`
	PrefixListOutbound string  `json:"prefix_list_outbound"`
	PrependInbound     int     `json:"prepend_inbound"`
	PrependOutbound    int     `json:"prepend_outbound"`
	RcircuitID         int     `json:"rcircuit_id"`
	RemoteIP           string  `json:"remote_ip"`
	SiteID             int     `json:"site_id"`
	Status             string  `json:"status"`
	SwitchID           int     `json:"switch_id"`
	SwitchName         string  `json:"switch_name"`
	SwitchPortID       int     `json:"switch_port_id"`
	TermSwitchID       int     `json:"term_switch_id"`
	TermSwitchName     string  `json:"term_switch_name"`
	TerminateOnSwitch  string  `json:"terminate_on_switch"`
	UpdateSource       string  `json:"update_source"`
	Vlan               int     `json:"vlan"`
	Weight             int     `json:"weight"`
}

// EBGPSite .
type EBGPSite struct {
	ACLPolicy           string `json:"acl_policy"`
	ASN                 int    `json:"asn"`
	HypervisorASN       int    `json:"hypervisor_asn"`
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Nat                 string `json:"nat"`
	PhysicalInstanceASN int    `json:"physical_instance_asn"`
	RpID                int    `json:"rp_id"`
	RpName              string `json:"rp_name"`
	RpTag               string `json:"rp_tag"`
	SpineAsn            int    `json:"spine_asn"`
	TorASN              int    `json:"tor_asn"`
	VirtualInstanceASN  int    `json:"virtual_instance_asn"`
	VPN                 string `json:"vpn"`
}

// EBGPVNet .
type EBGPVNet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// EBGPRouteMap .
type EBGPRouteMap struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// EBGPOffloader .
type EBGPOffloader struct {
	IPAddress     string `json:"ip_address"`
	Location      string `json:"location"`
	OffloadPortID int    `json:"offload_port_id"`
	SwitchID      int    `json:"switch_id"`
}

// EBGPPort .
type EBGPPort struct {
	OneG              int         `json:"1g"`
	AdminDown         interface{} `json:"admin_down"`
	BgpPrefixes       string      `json:"bgp_prefixes"`
	BgpState          string      `json:"bgp_state"`
	BgpUptime         string      `json:"bgp_uptime"`
	InstanceID        int         `json:"instance_id"`
	Lacp              string      `json:"lacp"`
	ModifiedDate      string      `json:"modified_date"`
	MonthlyFee        int         `json:"monthly_fee"`
	MonthlyFeePartial int         `json:"monthly_fee_partial"`
	Mtu               int         `json:"mtu"`
	Name              string      `json:"name"`
	OrderNumber       string      `json:"order_number"`
	ParentPort        int         `json:"parent_port"`
	Port              string      `json:"port"`
	PortName          string      `json:"portName"`
	Portname          string      `json:"portName_"`
	PortExtension     int         `json:"port_extension"`
	PortID            int         `json:"port_id"`
	PortIndex         int         `json:"port_index"`
	Power             int         `json:"power"`
	Present           int         `json:"present"`
	Price             int         `json:"price"`
	Price10           int         `json:"price_10"`
	Price20           int         `json:"price_20"`
	Price30           int         `json:"price_30"`
	Price40           int         `json:"price_40"`
	ProviderID        int         `json:"provider_id"`
	Speed             string      `json:"speed"`
	Status            string      `json:"status"`
	SwitchName        string      `json:"switchName"`
	SwitchID          int         `json:"switch_id"`
	Switchname        string      `json:"switch_name"`
	SwitchType        string      `json:"switch_type"`
	TenantID          int         `json:"tenant_id"`
	Transceiver       string      `json:"transceiver"`
	Type              string      `json:"type"`
	VlanFrom          int         `json:"vlan_from"`
	VlanTo            int         `json:"vlan_to"`
}

// EBGPSwitch .
type EBGPSwitch struct {
	Location string `json:"location"`
	PortID   int    `json:"port_id"`
	SwitchID int    `json:"switch_id"`
	Type     string `json:"type"`
}

// EBGPUpdatedSource .
type EBGPUpdatedSource struct {
	IPAddress string `json:"ip_address"`
	Location  string `json:"location"`
	SwitchID  int    `json:"switch_id"`
}
