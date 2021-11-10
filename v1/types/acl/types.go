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

package acl

type ACL struct {
	AclDstGrpName    string      `json:"acl_dst_grp_name"`
	AclDstGrpPorts   string      `json:"acl_dst_grp_ports"`
	AclSrcGrpName    string      `json:"acl_src_grp_name"`
	AclSrcGrpPorts   string      `json:"acl_src_grp_ports"`
	Action           string      `json:"action"`
	AdminRoleApprove int         `json:"admin_role_approve"`
	Approval         string      `json:"approval"`
	Approve          string      `json:"approve"`
	Comment          string      `json:"comment"`
	CreateDate       int         `json:"create_date"`
	Creator          string      `json:"creator"`
	CreatorID        int         `json:"creator_id"`
	Destination      string      `json:"destination"`
	DstLength        int         `json:"dst_length"`
	DstPortFrom      int         `json:"dst_port_from"`
	DstPortGroup     int         `json:"dst_port_group"`
	DstPortTo        int         `json:"dst_port_to"`
	DstPrefix        string      `json:"dst_prefix"`
	EditDate         int         `json:"edit_date"`
	Editor           string      `json:"editor"`
	EditorID         int         `json:"editor_id"`
	Established      int         `json:"established"`
	ID               int         `json:"id"`
	IPVersion        string      `json:"ip_version"`
	Name             string      `json:"name"`
	Protocol         string      `json:"protocol"`
	Reverse          string      `json:"reverse"`
	Sites            []int       `json:"sites"`
	Source           string      `json:"source"`
	SrcLength        int         `json:"src_length"`
	SrcPortFrom      int         `json:"src_port_from"`
	SrcPortGroup     int         `json:"src_port_group"`
	SrcPortTo        int         `json:"src_port_to"`
	SrcPrefix        string      `json:"src_prefix"`
	Status           string      `json:"status"`
	Tenants          string      `json:"tenants"`
	TenantsID        string      `json:"tenantsID"`
	ValidUntil       interface{} `json:"valid_until"`
}
