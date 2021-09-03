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

package vnet

import (
	"encoding/json"
	"fmt"

	"github.com/netrisai/netriswebapi/http"
	v1address "github.com/netrisai/netriswebapi/http/addresses/v1"
)

type VNetClient struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *VNetClient {
	return &VNetClient{c}
}

func parse(APIResult *http.APIResponse) ([]*VNet, error) {
	var sites []*VNet
	err := http.Decode(APIResult.Data, &sites)
	if err != nil {
		return sites, fmt.Errorf("{parse} %s", err)
	}
	return sites, nil
}

func parseInfo(APIResult *http.APIResponse) ([]*VNetInfo, error) {
	var sites []*VNetInfo
	err := http.Decode(APIResult.Data, &sites)
	if err != nil {
		return sites, fmt.Errorf("{parseInfo} %s", err)
	}
	return sites, nil
}

func (c *VNetClient) Get() ([]*VNet, error) {
	address := c.client.URL.String() + v1address.VNet
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetVNet} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetVNet} %s", err)
	}
	return items, nil
}

func (c *VNetClient) GetByID(id int) ([]*VNetInfo, error) {
	address := c.client.URL.String() + v1address.VNetInfo + fmt.Sprintf("?ids=%d", id)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetVNetByID} %s", err)
	}

	vnet, err := parseInfo(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetVNetByID} %s", err)
	}
	return vnet, nil
}

func (c *VNetClient) Add(vnet *VNetAdd) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(vnet)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.VNet
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *VNetClient) Update(vnet *VNetUpdate) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(vnet)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{Update} %s", err)
	}
	address := c.client.URL.String() + v1address.VNet
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{Update} %s", err)
	}

	return reply, nil
}

func (c *VNetClient) Validate(vnet *VNetUpdate) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(vnet)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{Update} %s", err)
	}
	address := c.client.URL.String() + v1address.VNetValidate
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{Update} %s", err)
	}

	return reply, nil
}

func (c *VNetClient) Delete(id int, circuitTenants []int) (reply http.HTTPReply, err error) {
	lb := struct {
		ID             int   `json:"id"`
		CircuitTenants []int `json:"circuitTenants"`
	}{ID: id, CircuitTenants: circuitTenants}
	js, err := json.Marshal(lb)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.VNet
	reply, err = c.client.Delete(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
