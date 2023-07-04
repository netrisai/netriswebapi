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

package ipam

/*
IPAM Structure for GET requests
*/

type IPAM struct {
	AllocationID   int      `json:"allocationID"`
	Children       []*IPAM  `json:"children"`
	DefaultGateway string   `json:"defaultGateway"`
	Description    string   `json:"description"`
	FullName       string   `json:"fullName"`
	HasHosts       bool     `json:"hasHosts"`
	ID             int      `json:"id"`
	IPFamily       string   `json:"ipFamily"`
	Name           string   `json:"name"`
	ParentID       int      `json:"parentID"`
	Prefix         string   `json:"prefix"`
	Purpose        string   `json:"purpose"`
	Readonly       string   `json:"readonly"`
	Sites          []IDName `json:"sites"`
	Subnet         struct {
		Length int    `json:"length"`
		Prefix string `json:"prefix"`
	} `json:"subnet"`
	Tenant      IDName      `json:"tenant"`
	Type        string      `json:"type"`
	Tags        []string    `json:"tags"`
	Utilization Utilization `json:"utilization"`
	Vpc         IDName      `json:"vpc"`
}

type IDName struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Utilization struct {
	Percent int    `json:"percent"`
	Type    string `json:"type"`
}

type Host struct {
	Address     string      `json:"address"`
	Description string      `json:"description"`
	ID          int         `json:"id"`
	IPFamily    string      `json:"ipFamily"`
	Meta        HostAnycast `json:"meta"`
	Name        string      `json:"name"`
	Readonly    string      `json:"readonly"`
	SubnetID    int         `json:"subnetID"`
	Type        string      `json:"type"`
}

type HostAnycast struct {
	Anycast bool `json:"anycast"`
}

/*
IPAM Allocation Structure for POST requests
*/

type Allocation struct {
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Prefix      string  `json:"prefix"`
	Tenant      IDName  `json:"tenant"`
	Vpc         *IDName `json:"vpc,omitempty"`
}

/*
IPAM Subnet Structure for POST requests
*/

type Subnet struct {
	DefaultGateway string   `json:"defaultGateway"`
	Description    string   `json:"description"`
	Name           string   `json:"name"`
	Prefix         string   `json:"prefix"`
	Purpose        string   `json:"purpose"`
	Sites          []IDName `json:"sites"`
	Tenant         IDName   `json:"tenant"`
	Tags           []string `json:"tags"`
	Vpc            *IDName  `json:"vpc,omitempty"`
}
