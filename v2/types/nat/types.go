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

package nat

type NAT struct {
	Action struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"action"`
	Comment     string `json:"comment"`
	CreatedDate int    `json:"createdDate"`
	Creator     struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"creator"`
	DestinationAddress string `json:"destinationAddress"`
	DestinationPort    string `json:"destinationPort"`
	DnatToIP           string `json:"dnatToIP"`
	DnatToPort         int    `json:"dnatToPort"`
	Editor             struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"editor"`
	ID           int    `json:"id"`
	ModifiedDate int    `json:"modifiedDate"`
	Name         string `json:"name"`
	Protocol     struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"protocol"`
	Site struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"site"`
	SnatToIP      string `json:"snatToIP"`
	SnatToPool    string `json:"snatToPool"`
	SourceAddress string `json:"sourceAddress"`
	SourcePort    string `json:"sourcePort"`
	State         struct {
		Label  string `json:"label"`
		Status string `json:"status"`
		Value  string `json:"value"`
	} `json:"state"`
}
