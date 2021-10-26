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

package userrole

import "github.com/netrisai/netriswebapi/v1/types/permission"

type UserRole struct {
	CreateDate   int      `json:"create_date"`
	Description  string   `json:"description"`
	ID           int      `json:"id"`
	ModifiedDate int      `json:"modified_date"`
	Name         string   `json:"name"`
	PermID       int      `json:"perm_id"`
	PermName     string   `json:"perm_name"`
	Tenants      []Tenant `json:"tenants"`
}

type Tenant struct {
	T2r         int    `json:"t2r,omitempty"`
	ID          int    `json:"id,omitempty"`
	TenantID    int    `json:"tenant_id"`
	TenantName  string `json:"tenant_name"`
	TenantRead  string `json:"tenant_read,omitempty"`
	TenantWrite string `json:"tenant_write,omitempty"`
}

type UserRoleAdd struct {
	Description     string                     `json:"description"`
	ID              int                        `json:"id,omitempty"`
	Name            string                     `json:"name"`
	PermissionGroup permission.PermissionGroup `json:"permissionGroup"`
	Tenants         []Tenant                   `json:"tenants"`
}
