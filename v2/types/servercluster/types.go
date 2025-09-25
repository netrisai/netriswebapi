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

package servercluster

type IDName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Servers struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Shared bool   `json:"shared"`
}

type GatewayPrefix struct {
	Prefix string `json:"prefix"`
}

type VNet struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	IPv4Gateways []GatewayPrefix `json:"ipv4Gateways"`
	IPv6Gateways []GatewayPrefix `json:"ipv6Gateways"`
}

type Allocation struct {
	ID     int    `json:"id"`
	Prefix string `json:"prefix"`
}

type Subnet struct {
	ID     int    `json:"id"`
	Prefix string `json:"prefix"`
}

type Resources struct {
	VNets       []VNet       `json:"vnets"`
	Allocations []Allocation `json:"allocations"`
	Subnets     []Subnet     `json:"subnets"`
}

type StatusInfo struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type ServerCluster struct {
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	State              string      `json:"state"`
	Status             StatusInfo  `json:"status"`
	Admin              IDName      `json:"admin"`
	Site               IDName      `json:"site"`
	VPC                IDName      `json:"vpc"`
	SrvClusterTemplate IDName      `json:"srvClusterTemplate"`
	Tags               []string    `json:"tags"`
	Servers            []Servers   `json:"servers"`
	Resources          Resources   `json:"resources"`
	ModifiedDate       int         `json:"modifiedDate"`
	CreatedDate        int         `json:"createdDate"`
}

type ServerClusterW struct {
	Name               string    `json:"name"`
	Admin              IDName    `json:"admin"`
	Site               IDName    `json:"site"`
	VPC                IDName    `json:"vpc"`
	SrvClusterTemplate IDName    `json:"srvClusterTemplate"`
	Tags               []string  `json:"tags"`
	Servers            []Servers `json:"servers"`
}

type ServerClusterU struct {
	Tags    []string  `json:"tags"`
	Servers []Servers `json:"servers"`
}
