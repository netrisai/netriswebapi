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

package user

import "github.com/netrisai/netriswebapi/v1/types/userrole"

type User struct {
	AuthSchemeID int          `json:"auth_scheme_id"`
	Company      string       `json:"company"`
	CreateDate   int          `json:"create_date"`
	Email        string       `json:"email"`
	EmailCc      string       `json:"email_cc"`
	Fullname     string       `json:"fullname"`
	ID           int          `json:"id"`
	ModifiedDate int          `json:"modified_date"`
	Name         string       `json:"name"`
	PermID       int          `json:"perm_id"`
	PermName     string       `json:"perm_name"`
	Phone        string       `json:"phone"`
	Position     string       `json:"position"`
	RoleID       int          `json:"role_id"`
	Rolename     string       `json:"rolename"`
	Tenants      []UserTenant `json:"tenants"`
}

type UserTenant struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	TenantWrite bool   `json:"tenant_write"`
}

type UserAdd struct {
	ID              int               `json:"id,omitempty"`
	Company         string            `json:"company"`
	Email           string            `json:"email"`
	EmailCc         string            `json:"email_cc"`
	Fullname        string            `json:"fullname"`
	Name            string            `json:"name"`
	PermissionGroup PermissionGroup   `json:"permissionGroup,omitempty"`
	Phonenumber     string            `json:"phonenumber"`
	Position        string            `json:"position"`
	Tenants         []userrole.Tenant `json:"tenants"`
	UserRole        userrole.UserRole `json:"userRole"`
}

type PermissionGroup struct {
	ID   int `json:"id"`
	Name int `json:"name"`
}
