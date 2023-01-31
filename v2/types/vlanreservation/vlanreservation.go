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

package vlanreservation

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/netrisai/netriswebapi/http"
	v2address "github.com/netrisai/netriswebapi/http/addresses/v2"
)

type Client struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *Client {
	return &Client{c}
}

func parseList(APIResult *http.APIResponse) ([]*VlanReservation, error) {
	var items []*VlanReservation
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse VlanReservation List} %s", err)
	}
	return items, nil
}

func (c *Client) Get() ([]*VlanReservation, error) {
	address := c.client.URL.String() + v2address.VlanReservation
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetVlanReservation} %s", err)
	}

	items, err := parseList(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetVlanReservation} %s", err)
	}
	return items, nil
}

func (c *Client) GetByVlanId(id int) ([]*VlanReservation, error) {
	address := c.client.URL.String() + v2address.VlanReservation + "/" + strconv.Itoa(id)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetVlanReservationByVlanId} %s", err)
	}

	items, err := parseList(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetVlanReservationByVlanId} %s", err)
	}
	return items, nil
}

func (c *Client) GetByVlanIdAndSite(vlanId, siteId int) ([]*VlanReservation, error) {
	address := c.client.URL.String() + v2address.VlanReservation + "/" + strconv.Itoa(vlanId) + "?filterBySite=" + strconv.Itoa(siteId)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetVlanReservationBySite} %s", err)
	}

	items, err := parseList(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetVlanReservationBySite} %s", err)
	}
	return items, nil
}

func (c *Client) GetBySite(id int) ([]*VlanReservation, error) {
	address := c.client.URL.String() + v2address.VlanReservation + "?filterBySite=" + strconv.Itoa(id)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetVlanReservationBySite} %s", err)
	}

	items, err := parseList(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetVlanReservationBySite} %s", err)
	}
	return items, nil
}

func (c *Client) Post(vlan *VlanReservation) (reply http.HTTPReply, err error) {
	if vlan.Meta == nil {
		vlan.Meta = make(map[string]interface{})
	}
	js, err := json.Marshal(vlan)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.VlanReservation
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *Client) Delete(vlanId, siteId int) (reply http.HTTPReply, err error) {
	address := c.client.URL.String() + v2address.VlanReservation + "/" + strconv.Itoa(vlanId) + "/" + strconv.Itoa(siteId)
	reply, err = c.client.Delete(address, nil)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
