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

package route

type Route struct {
	CreateDate       int    `json:"create_date"`
	Description      string `json:"description"`
	FilteredSwitches []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"filteredSwitches"`
	ID           int    `json:"id"`
	Internal     bool   `json:"internal"`
	ModifiedDate int    `json:"modified_date"`
	Name         string `json:"name"`
	NextHop      string `json:"next_hop"`
	Prefix       string `json:"prefix"`
	PrefixLength int    `json:"prefix_length"`
	SiteID       int    `json:"site_id"`
	SiteName     string `json:"site_name"`
	State        string `json:"state"`
	Switches     []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"switches"`
	Vpc IDName `json:"vpc"`
}

type RouteAdd struct {
	Description string  `json:"description"`
	NextHop     string  `json:"next_hop"`
	Prefix      string  `json:"prefix"`
	RouteID     int     `json:"route_id,omitempty"`
	SiteID      int     `json:"site_id"`
	StateStatus string  `json:"stateStatus"`
	Switches    []int   `json:"switches"`
	Vpc         *IDName `json:"vpc,omitempty"`
}

type IDName struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
