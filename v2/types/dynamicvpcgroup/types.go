/*
Copyright 2026. Netris, Inc.

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

package dynamicvpcgroup

import "encoding/json"

type IDName struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

type VPCRef struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	VNI     *int   `json:"vni"`
	VrfName string `json:"vrfName"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Member struct {
	VPC   VPCRef `json:"vpc"`
	Users []User `json:"users"`
}

type Server struct {
	Host     string `json:"host"`
	AuthPort int    `json:"authPort"`
	AcctPort int    `json:"acctPort"`
	Secret   string `json:"secret"`
	Priority int    `json:"priority"`
}

type RadiusClient struct {
	NASIdentifier string `json:"nasIdentifier"`
	NASIP         string `json:"nasIp"`
	Switch        IDName `json:"switch"`
}

type Usage struct {
	VNet          int `json:"vnet"`
	ServerCluster int `json:"serverCluster"`
	Switch        int `json:"switch"`
}

type DynamicVPCGroup struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	DisplayID     string         `json:"displayId"`
	ResourceGroup IDName         `json:"resourceGroup"`
	RadiusServer  string         `json:"radiusServer"`
	Tags          []string       `json:"tags"`
	Members       []Member       `json:"members"`
	SharedSecret  string         `json:"sharedSecret"`
	Servers       []Server       `json:"servers"`
	RadiusClients []RadiusClient `json:"radiusClients"`
	MemberCount   int            `json:"memberCount"`
	UserCount     int            `json:"userCount"`
	ServerCount   int            `json:"serverCount"`
	SwitchCount   int            `json:"switchCount"`
	UsedBy        Usage          `json:"usedBy"`
	CreatedDate   int64          `json:"createdDate"`
	ModifiedDate  int64          `json:"modifiedDate"`
}

type UserWrite struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

type MemberWrite struct {
	VPC   IDName      `json:"vpc"`
	Users []UserWrite `json:"users,omitempty"`
}

type ServerWrite struct {
	Host     string `json:"host"`
	AuthPort int    `json:"authPort,omitempty"`
	AcctPort int    `json:"acctPort,omitempty"`
	Secret   string `json:"secret,omitempty"`
}

type DynamicVPCGroupAdd struct {
	Name          string         `json:"name"`
	ResourceGroup IDName         `json:"resourceGroup"`
	RadiusServer  string         `json:"radiusServer"`
	Tags          []string       `json:"tags,omitempty"`
	Members       []MemberWrite  `json:"members"`
	Servers       *[]ServerWrite `json:"servers,omitempty"`
}

// Members replaces the full member/user list; nil Servers preserves servers.
type DynamicVPCGroupUpdate struct {
	Name    string         `json:"name"`
	Tags    []string       `json:"tags"`
	Members []MemberWrite  `json:"members"`
	Servers *[]ServerWrite `json:"servers,omitempty"`
}

// Nil Tags preserves tags; a non-nil empty slice clears them.
func (g DynamicVPCGroupUpdate) MarshalJSON() ([]byte, error) {
	type groupUpdate DynamicVPCGroupUpdate
	if g.Tags == nil {
		return json.Marshal(struct {
			groupUpdate
			Tags []string `json:"tags,omitempty"`
		}{groupUpdate: groupUpdate(g)})
	}
	return json.Marshal(groupUpdate(g))
}
