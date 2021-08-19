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

package port

type Port struct {
	AdminDown        string               `json:"adminDown"`
	AutoNeg          string               `json:"autoNeg"`
	BGP              PortBGP              `json:"bgp"`
	Breakout         int64                `json:"breakout"`
	CreatedDate      int64                `json:"createdDate"`
	Description      string               `json:"description"`
	DesiredSpeed     string               `json:"desiredSpeed"`
	Duplex           string               `json:"duplex"`
	Extension        int64                `json:"extension"`
	ID               int64                `json:"id"`
	IfName           interface{}          `json:"ifName"`
	Lacp             string               `json:"lacp"`
	ModifiedDate     int64                `json:"modifiedDate"`
	Mtu              int64                `json:"mtu"`
	Name             string               `json:"name"`
	ParentPort       int64                `json:"parentPort"`
	Port             string               `json:"port"`
	ShortName        string               `json:"shortName"`
	Site             IDName               `json:"site"`
	SlavePorts       []interface{}        `json:"slavePorts"`
	SortIndexAsc     string               `json:"sortIndexAsc"`
	SortIndexDesc    string               `json:"sortIndexDesc"`
	Speed            string               `json:"speed"`
	StateInHierarchy PortStateInHierarchy `json:"stateInHierarchy"`
	Status           PortStatus           `json:"status"`
	Switch           PortSwitch           `json:"switch"`
	SwitchName       string               `json:"switchName"`
	Tenant           IDName               `json:"tenant"`
	Transceiver      string               `json:"transceiver"`
	Used             bool                 `json:"used"`
}

type IDName struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type PortBGP struct {
	Prefix string      `json:"prefix"`
	State  string      `json:"state"`
	Time   interface{} `json:"time"`
}

type PortStateInHierarchy struct {
	Aggregated       bool  `json:"aggregated"`
	Breakout         bool  `json:"breakout"`
	BreakoutChild    int64 `json:"breakoutChild"`
	Extended         int64 `json:"extended"`
	ExtensionsParent int64 `json:"extensionsParent"`
	LagMember        bool  `json:"lagMember"`
}

type PortStatus struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type PortSwitch struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
