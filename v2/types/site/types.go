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

type Site struct {
	AclPolicy             string                 `json:"aclPolicy"`
	ID                    int                    `json:"id,omitempty"`
	Name                  string                 `json:"name"`
	PublicAsn             int                    `json:"publicAsn"`
	RohAsn                int                    `json:"rohAsn"`
	RohProfile            *RohProfile            `json:"rohProfile,omitempty"`
	SiteMesh              IDName                 `json:"siteMesh"`
	SwitchFabric          string                 `json:"switchFabric"`
	SwitchFabricProviders *SwitchFabricProviders `json:"switchFabricProviders,omitempty"`
	VlanRange             string                 `json:"vlanRange"`
	VlanRangeAutoAssign		string								 `json:"vlanRangeAutoAssign,omitempty"`
	VMAsn                 int                    `json:"vmAsn"`
}

type RohProfile struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type IDName struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type SwitchFabricProviders struct {
	EquinixMetal  *EquinixMetal  `json:"equinixMetal,omitempty"`
	PhoenixNapBmc *PhoenixNapBmc `json:"phoenixNapBmc,omitempty"`
}

type EquinixMetal struct {
	Location      string `json:"location,omitempty"`
	ProjectAPIKey string `json:"projectApiKey,omitempty"`
	ProjectID     string `json:"projectId,omitempty"`
}

type PhoenixNapBmc struct {
	ClientID     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	Location     string `json:"location,omitempty"`
}
