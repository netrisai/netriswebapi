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

package subnet

// Subnet .
type Subnet struct {
	Children    []SubnetChild  `json:"children"`
	Description string         `json:"description"`
	ID          string         `json:"id"`
	IPVersion   string         `json:"ip_version"`
	Len         int            `json:"len"`
	Length      int            `json:"length"`
	Name        string         `json:"name"`
	NetpopID    int            `json:"netpop_id"`
	NetpopName  string         `json:"netpop_name"`
	Netpops     []SubnetNetpop `json:"netpops"`
	Prefix      string         `json:"prefix"`
	Purpose     string         `json:"purpose"`
	SiteID      string         `json:"siteID"`
	SiteName    string         `json:"site_name"`
	TENANTID    int            `json:"tenantID"`
	TenantID    int            `json:"tenant_id"`
	TenantName  string         `json:"tenant_name"`
	Type        string         `json:"type"`
}

// SubnetChild .
type SubnetChild struct {
	Anycast         int            `json:"anycast"`
	Children        []SubnetChild  `json:"children"`
	Description     string         `json:"description"`
	ID              string         `json:"id"`
	IPVersion       string         `json:"ip_version"`
	Len             int            `json:"len"`
	Length          int            `json:"length"`
	Name            string         `json:"name"`
	NetpopID        int            `json:"netpop_id"`
	NetpopName      string         `json:"netpop_name"`
	Netpops         []SubnetNetpop `json:"netpops"`
	OutOfAllocation bool           `json:"outOfAllocation"`
	ParentSite      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"parentSite"`
	Prefix     string `json:"prefix"`
	Purpose    string `json:"purpose"`
	SiteID     string `json:"siteID"`
	SiteName   string `json:"site_name"`
	TENANTID   int    `json:"tenantID"`
	TenantID   int    `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
	Type       string `json:"type"`
}

// SubnetNetpop .
type SubnetNetpop struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SubnetAdd .
type SubnetAdd struct {
	CIDR           string `json:"cidr"`
	ID             int    `json:"id"`
	IPVersion      string `json:"ipVersion"`
	Name           string `json:"name"`
	Purpose        string `json:"purpose"`
	SelectedSite   int    `json:"selectedSite"`
	SelectedSites  []Site `json:"selectedSites"`
	SelectedType   string `json:"selectedType"`
	SelectedTenant int    `json:"selectedTenant"`
}

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
