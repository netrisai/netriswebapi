/*
Copyright 2023. Netris, Inc.

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

package serverclustertemplate

type ServerClusterTemplate struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	CreatedDate  int    `json:"createdDate"`
	ModifiedDate int    `json:"modifiedDate"`
	Vnets        []Vnet `json:"vnets"`
}

type Vnet struct {
	ID          int      `json:"id"`
	Postfix     string   `json:"postfix"`
	Type        string   `json:"type"`
	Vlan        string   `json:"vlan"`
	VlanID      string   `json:"vlanID"`
	ServerNics  []string `json:"serverNics"`
	IPv4Gateway string   `json:"ipv4Gateway,omitempty"`
}
