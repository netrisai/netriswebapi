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
	"strconv"

	"github.com/netrisai/netriswebapi/http"
	v2address "github.com/netrisai/netriswebapi/http/addresses/v2"
)

type VNetClient struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *VNetClient {
	return &VNetClient{c}
}

func parseAPIVnets(APIResult *http.APIResponse) ([]*VNet, error) {
	var items []*VNet
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parseAPIVnet} %s", err)
	}
	return items, nil
}

func parseAPIVnet(APIResult *http.APIResponse) (*VNetDetailed, error) {
	var items *VNetDetailed
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parseAPIVnet} %s", err)
	}
	return items, nil
}

func (c *VNetClient) Get() ([]*VNet, error) {
	address := c.client.URL.String() + v2address.VNetBase
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}

	items, err := parseAPIVnets(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}
	return items, nil
}

func (c *VNetClient) GetByID(id int) (*VNetDetailed, error) {
	address := c.client.URL.String() + v2address.VNetBase + "/" + strconv.Itoa(id)
	fmt.Println(address)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetByID} %s", err)
	}

	vnet, err := parseAPIVnet(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetByID} %s", err)
	}
	return vnet, nil
}

// AddVNet .
func (c *VNetClient) Add(vnet *VNetAdd) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(vnet)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.VNetBase
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

// UpdateVNet .
func (c *VNetClient) Update(id int, vnet *VNetUpdate) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(vnet)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{UpdateVNet} %s", err)
	}
	address := c.client.URL.String() + v2address.VNetBase + "/" + strconv.Itoa(id)
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{UpdateVNet} %s", err)
	}

	return reply, nil
}

// DeleteVNet .
func (c *VNetClient) Delete(id int) (reply http.HTTPReply, err error) {
	address := c.client.URL.String() + v2address.VNetBase + "/" + strconv.Itoa(id)
	reply, err = c.client.Delete(address, nil)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
