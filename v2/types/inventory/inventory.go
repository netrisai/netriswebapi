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

package inventory

import (
	"fmt"
	"strconv"

	"github.com/netrisai/netriswebapi/http"
	v2address "github.com/netrisai/netriswebapi/http/addresses/v2"
)

type InventoryClient struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *InventoryClient {
	return &InventoryClient{c}
}

func parseInventories(APIResult *http.APIResponse) ([]*HW, error) {
	var items []*HW
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parseInventories} %s", err)
	}
	return items, nil
}

func parseInventory(APIResult *http.APIResponse) (*HW, error) {
	var items *HW
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parseInventories} %s", err)
	}
	return items, nil
}

func (c *InventoryClient) Get() ([]*HW, error) {
	address := c.client.URL.String() + v2address.InventoryBase
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}

	items, err := parseInventories(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}
	return items, nil
}

func (c *InventoryClient) GetByID(id int) (*HW, error) {
	address := c.client.URL.String() + v2address.InventoryBase + "/" + strconv.Itoa(id)
	fmt.Println(address)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetByID} %s", err)
	}

	hw, err := parseInventory(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetByID} %s", err)
	}
	return hw, nil
}
