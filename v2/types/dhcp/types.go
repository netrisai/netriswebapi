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

package dhcp

type DHCPOptionSet struct {
	AdditionalOptions []AdditionalOption `json:"additionalOptions"`
	CreatedDate       int                `json:"createdDate"`
	Description       string             `json:"description"`
	DNSServers        []string           `json:"dnsServers"`
	DomainSearch      string             `json:"domainSearch"`
	ID                int                `json:"id"`
	IsDefault         bool               `json:"isDefault"`
	LeaseTime         int                `json:"leaseTime"`
	ModifiedDate      int                `json:"modifiedDate"`
	Name              string             `json:"name"`
	NTPServers        []string           `json:"ntpServers"`
}

type AdditionalOption struct {
	Code  interface{} `json:"code"`
	Type  string      `json:"type"`
	Value string      `json:"value"`
}

type DHCPw struct {
	AdditionalOptions []AdditionalOption `json:"additionalOptions"`
	Description       string             `json:"description"`
	DNSServers        []string           `json:"dnsServers"`
	DomainSearch      string             `json:"domainSearch"`
	LeaseTime         int                `json:"leaseTime"`
	Name              string             `json:"name"`
	NTPServers        []string           `json:"ntpServers"`
}
