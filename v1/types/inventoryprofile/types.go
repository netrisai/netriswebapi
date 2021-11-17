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

package inventoryprofile

type Profile struct {
	CreatedDate  int          `json:"created_date"`
	CustomRules  []CustomRule `json:"customRules"`
	Description  string       `json:"description"`
	DNSServers   string       `json:"dns_servers"`
	ID           int          `json:"id"`
	Ipv4SSH      string       `json:"ipv4_ssh"`
	Ipv6SSH      string       `json:"ipv6_ssh"`
	ModifiedDate int          `json:"modified_date"`
	Name         string       `json:"name"`
	NTPServers   string       `json:"ntp_servers"`
	Timezone     string       `json:"timezone"`
}

type CustomRule struct {
	Deleted   bool   `json:"deleted,omitempty"`
	DstPort   string `json:"dstPort"`
	ID        int    `json:"id"`
	Protocol  string `json:"protocol"`
	SrcPort   string `json:"srcPort"`
	SrcSubnet string `json:"srcSubnet"`
}

type ProfileW struct {
	CustomRules []CustomRule `json:"customRules"`
	Description string       `json:"description"`
	DNSServers  string       `json:"dns_servers"`
	ID          int64        `json:"id"`
	Ipv4List    string       `json:"ipv4_list"`
	Ipv6List    string       `json:"ipv6_list"`
	Name        string       `json:"name"`
	NTPServers  string       `json:"ntp_servers"`
	Timezone    Timezone     `json:"timezone"`
}

type Timezone struct {
	Label  string `json:"label"`
	Offset string `json:"offset"`
	TzCode string `json:"tzCode"`
}
