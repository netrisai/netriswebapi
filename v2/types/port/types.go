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

/*
Port Structure for GET requests
*/

type Port struct {
	AdminDown        string               `json:"adminDown"`
	AutoNeg          string               `json:"autoNeg"`
	BGP              PortBGP              `json:"bgp"`
	Breakout         string               `json:"breakout"`
	CreatedDate      int                  `json:"createdDate"`
	Description      string               `json:"description"`
	DesiredSpeed     string               `json:"desiredSpeed"`
	Duplex           string               `json:"duplex"`
	Extension        int                  `json:"extension"`
	ID               int                  `json:"id"`
	IfName           interface{}          `json:"ifName"`
	Lacp             string               `json:"lacp"`
	ModifiedDate     int                  `json:"modifiedDate"`
	Mtu              int                  `json:"mtu"`
	Name             string               `json:"name"`
	ParentPort       int                  `json:"parentPort"`
	Port             string               `json:"port"`
	Port_            string               `json:"_port"`
	PortIndex        string               `json:"portIndex"`
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
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PortBGP struct {
	Prefix string      `json:"prefix"`
	State  string      `json:"state"`
	Time   interface{} `json:"time"`
}

type PortStateInHierarchy struct {
	Aggregated       bool `json:"aggregated"`
	Breakout         bool `json:"breakout"`
	BreakoutChild    int  `json:"breakoutChild"`
	Extended         int  `json:"extended"`
	ExtensionsParent int  `json:"extensionsParent"`
	LagMember        bool `json:"lagMember"`
}

type PortStatus struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type PortSwitch struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

/*
Port Structure for PUT requests
*/

type PortUpdate struct {
	AdminDown   string               `json:"adminDown"`
	AutoNeg     string               `json:"autoNeg"`
	Breakout    string               `json:"breakout"`
	Description string               `json:"description"`
	Duplex      string               `json:"duplex"`
	Extension   PortUpdateExtenstion `json:"extension"`
	ID          int                  `json:"id,omitempty"`
	Mtu         int                  `json:"mtu"`
	Speed       string               `json:"speed"`
	Tenant      IDName               `json:"tenant"`
}

type PortUpdateExtenstion struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	VLANFrom int    `json:"vlanFrom"`
	VLANTo   int    `json:"vlanTo"`
}

/*
Port Extenstion structure for GET requests
*/

type PortExtension struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	VlanFrom int    `json:"vlanFrom"`
	VlanTo   int    `json:"vlanTo"`
}

/*
Port LAG structure for POST requests
*/

type PortLAG struct {
	AggregatedPort IDName           `json:"aggregatedPort"`
	Description    string           `json:"description"`
	Extension      PortLAGExtension `json:"extension"`
	Mtu            int              `json:"mtu"`
	Ports          []IDName         `json:"ports"`
	Tenant         IDName           `json:"tenant"`
}

type PortLAGExtension struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	VlanFrom int    `json:"vlanFrom"`
	VlanTo   int    `json:"vlanTo"`
}
