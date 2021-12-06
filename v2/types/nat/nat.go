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

package nat

import (
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

func parseList(APIResult *http.APIResponse) ([]*NAT, error) {
	var items []*NAT
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse NAT List} %s", err)
	}
	return items, nil
}

func parseSingle(APIResult *http.APIResponse) (*NAT, error) {
	var item *NAT
	err := http.Decode(APIResult.Data, &item)
	if err != nil {
		return item, fmt.Errorf("{parse NAT} %s", err)
	}
	return item, nil
}

func (c *Client) Get() ([]*NAT, error) {
	address := c.client.URL.String() + v2address.NAT
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetNAT} %s", err)
	}

	items, err := parseList(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetNAT} %s", err)
	}
	return items, nil
}

func (c *Client) GetByID(id int) (*NAT, error) {
	address := c.client.URL.String() + v2address.NAT + "/" + strconv.Itoa(id)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetNATByID} %s", err)
	}

	vnet, err := parseSingle(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetNATByID} %s", err)
	}
	return vnet, nil
}
