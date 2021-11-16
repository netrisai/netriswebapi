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

package portgroup

type PortGroup struct {
	Approve         interface{} `json:"approve"`
	ApproveAction   bool        `json:"approveAction"`
	ApproveState    bool        `json:"approveState"`
	CreatedDate     int         `json:"created_date"`
	ID              int         `json:"id"`
	ModifiedDate    int         `json:"modified_date"`
	Name            string      `json:"name"`
	ParentPortGrpID int         `json:"parent_port_grp_id"`
	PortFrom        []int       `json:"port_from"`
	PortGrpID       int         `json:"port_grp_id"`
	PortTo          []int       `json:"port_to"`
	Ports           []string    `json:"ports"`
	Status          string      `json:"status"`
	TenantName      []string    `json:"tenant_name"`
}
