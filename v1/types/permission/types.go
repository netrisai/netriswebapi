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

package permission

type PermissionGroup struct {
	CreateDate   int    `json:"create_date"`
	Description  string `json:"description"`
	ExternalAcl  string `json:"external_acl"`
	Hidden       string `json:"hidden"`
	ID           int    `json:"id"`
	ModifiedDate int    `json:"modified_date"`
	Name         string `json:"name"`
	Readonly     string `json:"readonly"`
}

type PermissionGroupAdd struct {
	Description string   `json:"description"`
	ExternalACL bool     `json:"externalACL"`
	Hidden      []string `json:"hidden"`
	ID          int64    `json:"id,omitempty"`
	Name        string   `json:"name"`
	ReadOnly    []string `json:"readOnly"`
}
