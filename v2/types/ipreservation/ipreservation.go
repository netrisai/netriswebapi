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

func parseList(APIResult *http.APIResponse) ([]*IPReservation, error) {
	var items []*IPReservation
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse IPReservation List} %s", err)
	}
	return items, nil
}

func parseItem(APIResult *http.APIResponse) (*IPReservation, error) {
	var item *IPReservation
	err := http.Decode(APIResult.Data, &item)
	if err != nil {
		return item, fmt.Errorf("{parse IPReservation List} %s", err)
	}
	return item, nil
}

func (c *Client) Get() ([]*IPReservation, error) {
	address := c.client.URL.String() + v2address.IPReservation
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetIPReservation} %s", err)
	}

	items, err := parseList(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetIPReservation} %s", err)
	}
	return items, nil
}

func (c *Client) GetBySite(id int) ([]*IPReservation, error) {
	address := c.client.URL.String() + v2address.IPReservation + "?filterBySite=" + strconv.Itoa(id)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetIPReservationBySite} %s", err)
	}

	items, err := parseList(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetIPReservationBySite} %s", err)
	}
	return items, nil
}

func (c *Client) GetByID(id int) (*IPReservation, error) {
	address := c.client.URL.String() + v2address.IPReservation + "/" + strconv.Itoa(id)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetIPReservationByID} %s", err)
	}

	item, err := parseItem(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetIPReservationByID} %s", err)
	}
	return item, nil
}

func (c *Client) GetByIPAddress(ip string) (*IPReservation, error) {
	address := c.client.URL.String() + v2address.IPReservation + "/" + ip
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetIPReservationByIPAddress} %s", err)
	}

	item, err := parseItem(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetIPReservationByIPAddress} %s", err)
	}
	return item, nil
}

func (c *Client) Add(item *IPReservation) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(item)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.IPReservation
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *Client) Delete(id int) (reply http.HTTPReply, err error) {
	address := c.client.URL.String() + v2address.IPReservation + "/" + strconv.Itoa(id)
	reply, err = c.client.Delete(address, nil)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
