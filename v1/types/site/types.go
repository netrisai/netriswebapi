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

package site

// Site .
type Site struct {
	ACLPolicy           string `json:"acl_policy"`
	ASN                 int    `json:"asn"`
	HypervisorASN       int    `json:"hypervisor_asn"`
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Nat                 string `json:"nat"`
	PhysicalInstanceAsn int    `json:"physical_instance_asn"`
	RoutingProfileID    int    `json:"rp_id"`
	RoutingProfilName   string `json:"rp_name"`
	RoutingProfilTag    string `json:"rp_tag"`
	SiteToSiteVPN       string `json:"siteToSiteVpn"`
	SpineASN            int    `json:"spine_asn"`
	TorASN              int    `json:"tor_asn"`
	VirtualInstanceASN  int    `json:"virtual_instance_asn"`
	VPN                 string `json:"vpn"`
}

// SiteAdd .
type SiteAdd struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	BorderSwitchASN     string `json:"borderSwitchAsn"`
	SpineSwitchASN      string `json:"spineSwitchAsn"`
	TorSwitchASN        string `json:"torSwitchAsn"`
	HypervisorHostASN   string `json:"hypervisorHostAsn"`
	PhysicalInstanceASN string `json:"physicalInstanceASN"`
	VirtualInstanceASN  string `json:"virtualInstanceASN"`
	VPN                 string `json:"vpn"`
	RoutingProfileID    int    `json:"rp_id"`
	Nat                 string `json:"nat"`
	ACLPolicy           string `json:"acl_policy"`
}
