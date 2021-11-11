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

package roh

import "github.com/netrisai/netriswebapi/v2/types/port"

/*
ROH Structure for GET requests
*/

type ROH struct {
	Addresses       []Address      `json:"addresses"`
	BgpType         string         `json:"bgpType"`
	Compute         interface{}    `json:"compute"`
	CreatedDate     int            `json:"createdDate"`
	ID              int            `json:"id"`
	InboundPrefixes []interface{}  `json:"inboundPrefixes"`
	Inherit         bool           `json:"inherit"`
	Interface       interface{}    `json:"interface"`
	LegacyMode      bool           `json:"legacyMode"`
	LinkLocals      []LinkLocal    `json:"linkLocals"`
	ModifiedDate    int            `json:"modifiedDate"`
	Name            string         `json:"name"`
	Ports           []port.Port    `json:"ports"`
	RoutingProfile  RoutingProfile `json:"routingProfile"`
	Site            Site           `json:"site"`
	Tenant          IDName         `json:"tenant"`
	Type            string         `json:"type"`
	UnicastAddress  string         `json:"unicastAddress"`
	Zone            interface{}    `json:"zone"`
}

type RoutingProfile struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type Site struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	RpID int    `json:"rpID"`
}

type IDName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LinkLocal struct {
	ID         int    `json:"id"`
	InstanceID int    `json:"instanceID"`
	PortID     int    `json:"portID"`
	Prefix     string `json:"prefix"`
}

type Address struct {
	Anycast bool   `json:"anycast"`
	ID      int    `json:"id,omitempty"`
	Prefix  string `json:"prefix"`
}

/*
ROH Structure for POST and PUT requests
*/

type ROHw struct {
	Addresses       []Address       `json:"addresses"`
	InboundPrefixes []InboundPrefix `json:"inboundPrefixes"`
	LegacyMode      bool            `json:"legacyMode"`
	Name            string          `json:"name"`
	Ports           []IDName        `json:"ports"`
	RoutingProfile  string          `json:"routingProfile"`
	Site            IDName          `json:"site"`
	Tenant          IDName          `json:"tenant"`
	Type            string          `json:"type"`
	UnicastAddress  string          `json:"unicastAddress"`
}

type InboundPrefix struct {
	Action    string `json:"action"`
	Condition string `json:"condition"`
	Subnet    string `json:"subnet"`
}
