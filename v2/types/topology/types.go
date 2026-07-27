/*
Copyright 2026. Netris, Inc.

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

package topology

type UFMImportLink struct {
	SourceGUID                     string `json:"source_guid"`
	DestinationGUID                string `json:"destination_guid"`
	SourcePortName                 string `json:"source_port_name,omitempty"`
	DestinationPortName            string `json:"destination_port_name,omitempty"`
	SourcePortNodeDescription      string `json:"source_port_node_description,omitempty"`
	DestinationPortNodeDescription string `json:"destination_port_node_description,omitempty"`
	Name                           string `json:"name,omitempty"`
}

type UFMImportLinkRaw map[string]interface{}

type UFMImportResult struct {
	Imported      bool   `json:"imported"`
	UfmID         string `json:"ufmId"`
	Links         int    `json:"links"`
	Nodes         int    `json:"nodes"`
	NetrisMatched int    `json:"netrisMatched"`
}

type UFMFabric struct {
	Type     string `json:"type"`
	FabricID string `json:"fabricId"`
}

type UFMEndpoint struct {
	ID   string `json:"id"`
	Node string `json:"node"`
}

type UFMTierNode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UFMNetworkPath struct {
	NetworkPath []string                 `json:"network_path"`
	Hops        int                      `json:"hops"`
	Tiers       map[string][]UFMTierNode `json:"tiers"`
}

type UFMServerPathProps struct {
	GUIDs []string `json:"guids"`
}

type UFMPathResult struct {
	Fabric     UFMFabric          `json:"fabric"`
	ObservedAt string             `json:"observed_at"`
	From       UFMEndpoint        `json:"from,omitempty"`
	To         UFMEndpoint        `json:"to,omitempty"`
	Reachable  *bool              `json:"reachable,omitempty"`
	ID         string             `json:"id"`
	Name       string             `json:"name"`
	Type       string             `json:"type"`
	Props      UFMServerPathProps `json:"props"`
	PathCount  int                `json:"path_count"`
	Paths      []UFMNetworkPath   `json:"paths"`
}

type UFMPathResponse struct {
	From        string          `json:"from,omitempty"`
	To          string          `json:"to,omitempty"`
	Server      string          `json:"server,omitempty"`
	ResultCount int             `json:"result_count"`
	Results     []UFMPathResult `json:"results"`
	Note        string          `json:"note,omitempty"`
}
