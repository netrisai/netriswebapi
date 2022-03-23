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

package graphboards

import "encoding/json"

type Board struct {
	ActiveDashboardID int         `json:"activeDashboardID"`
	Dashboards        []Dashboard `json:"dashboards"`
}

type Dashboard struct {
	Data       string `json:"data"`
	GridSize   int    `json:"grid_size"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Public     int    `json:"public"`
	TenantID   int    `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
}

type DashboardData struct {
	Activegraphtype string `json:"activegraphtype"`
	GraphSize       int    `json:"graph_size"`
	MFunction       string `json:"mFunction"`
	MMembers        []struct {
		ID        string `json:"id"`
		Interface string `json:"interface"`
		Name      string `json:"name"`
		Switch    string `json:"switch"`
		VlanID    string `json:"vlan_id"`
	} `json:"mMembers"`
	MTime  []string `json:"mTime"`
	MTitle string   `json:"mTitle"`
	MUUID  string   `json:"mUUID"`
}

func (d Dashboard) GetData() []DashboardData {
	var data []DashboardData
	_ = json.Unmarshal([]byte(d.Data), &data)
	return data
}
