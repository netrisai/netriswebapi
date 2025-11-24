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

package nvlinkledger

type NVLinkledger struct {
	SrvCluster  IDName      `json:"srvCluster"`
	NmxcID      string      `json:"nmxcID"`
	PartitionID int         `json:"partitionID"`
	GPUUids     []string    `json:"gpuUids"`
	Servers     []IDName    `json:"servers"`
	Custom      interface{} `json:"custom"`
}

type IDName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type NVLinkledgerW struct {
	SrvClusterID int         `json:"srvClusterID"`
	NmxcID       string      `json:"nmxcID"`
	PartitionID  int         `json:"partitionID"`
	GPUUids      []string    `json:"gpuUids"`
	Custom       interface{} `json:"custom"`
}

type NVLinkledgerU struct {
	GPUUids []string    `json:"gpuUids"`
	Custom  interface{} `json:"custom"`
}
