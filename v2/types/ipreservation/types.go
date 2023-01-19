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

package ipreservation

type IPReservation struct {
	Consumer Consumer               `json:"consumer"`
	Host     Host                   `json:"host"`
	ID       int                    `json:"id"`
	Meta     map[string]interface{} `json:"meta"`
}

type Consumer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Host struct {
	Address string `json:"address"`
	ID      int    `json:"id"`
	Subnet  Subnet `json:"subnet"`
}

type Subnet struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Prefix  string `json:"prefix"`
	Purpose string `json:"purpose"`
}
