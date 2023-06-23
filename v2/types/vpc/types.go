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

package vpc

type VPC struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	AdminTenant  AdminTenant   `json:"adminTenant"`
	GuestTenant  []GuestTenant `json:"guestTenant"`
	Tags         []string      `json:"tags"`
	IsDefault    bool          `json:"isDefault"`
	ModifiedDate int           `json:"modifiedDate"`
	CreatedDate  int           `json:"createdDate"`
}

type GuestTenant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AdminTenant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type VPCw struct {
	Name        string        `json:"name"`
	AdminTenant AdminTenant   `json:"adminTenant"`
	GuestTenant []GuestTenant `json:"guestTenant"`
	Tags        []string      `json:"tags"`
}
