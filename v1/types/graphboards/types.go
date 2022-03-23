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

type board struct {
	ActiveDashboardID int         `json:"activeDashboardID"`
	Dashboards        []dashboard `json:"dashboards"`
}

type Board struct {
	ActiveDashboardID int         `json:"activeDashboardID"`
	Dashboards        []Dashboard `json:"dashboards"`
}

type dashboard struct {
	Data       string `json:"data"`
	GridSize   int    `json:"grid_size"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Public     int    `json:"public"`
	TenantID   int    `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
}

type Dashboard struct {
	Data       []DashboardData `json:"data"`
	GridSize   int             `json:"grid_size"`
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	Public     int             `json:"public"`
	TenantID   int             `json:"tenant_id"`
	TenantName string          `json:"tenant_name"`
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

func (d dashboard) getData() []DashboardData {
	var data []DashboardData
	_ = json.Unmarshal([]byte(d.Data), &data)
	return data
}

func (d dashboard) ConvertToDashboard() Dashboard {
	return Dashboard{
		Data:       d.getData(),
		GridSize:   d.GridSize,
		ID:         d.ID,
		Name:       d.Name,
		Public:     d.Public,
		TenantID:   d.TenantID,
		TenantName: d.TenantName,
	}
}

func (d *board) ConvertToBoard() *Board {
	dashboards := []Dashboard{}
	for _, dboard := range d.Dashboards {
		dashboards = append(dashboards, dboard.ConvertToDashboard())
	}
	return &Board{
		ActiveDashboardID: d.ActiveDashboardID,
		Dashboards:        dashboards,
	}
}
