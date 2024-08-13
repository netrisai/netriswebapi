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

type ServerCluster struct {
	ID                 int      `json:"id"`
	Name               string   `json:"name"`
	Admin              IDName   `json:"admin"`
	Site               IDName   `json:"site"`
	VPC                IDName   `json:"vpc"`
	SrvClusterTemplate IDName   `json:"srvClusterTemplate"`
	Tags               []string `json:"tags"`
	Servers            []IDName `json:"servers"`
	ModifiedDate       int      `json:"modifiedDate"`
	CreatedDate        int      `json:"createdDate"`
}

type ServerClusterW struct {
	Name               string   `json:"name"`
	Admin              IDName   `json:"admin"`
	Site               IDName   `json:"site"`
	VPC                IDName   `json:"vpc"`
	SrvClusterTemplate IDName   `json:"srvClusterTemplate"`
	Tags               []string `json:"tags"`
	Servers            []IDName `json:"servers"`
}

type ServerClusterU struct {
	Tags    []string `json:"tags"`
	Servers []IDName `json:"servers"`
}
