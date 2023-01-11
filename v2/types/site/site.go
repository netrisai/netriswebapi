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

package site

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

func parse(APIResult *http.APIResponse) ([]*Site, error) {
	var sites []*Site
	err := http.Decode(APIResult.Data, &sites)
	if err != nil {
		return sites, fmt.Errorf("{parse} %s", err)
	}
	return sites, nil
}

func parseSingle(APIResult *http.APIResponse) (*Site, error) {
	var site *Site
	err := http.Decode(APIResult.Data, &site)
	if err != nil {
		return site, fmt.Errorf("{parse} %s", err)
	}
	return site, nil
}

func (c *Client) Get() ([]*Site, error) {
	address := c.client.URL.String() + v2address.Sites
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetSites} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetSites} %s", err)
	}
	return items, nil
}

func (c *Client) GetByID(id int) (*Site, error) {
	address := c.client.URL.String() + v2address.Sites + "/" + strconv.Itoa(id)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetSitesByID} %s", err)
	}

	item, err := parseSingle(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetSitesByID} %s", err)
	}
	return item, nil
}

// Add .
func (c *Client) Add(site *Site) (reply http.HTTPReply, err error) {
	if site.SwitchFabric == "" {
		site.SwitchFabric = "netris"
	}
	js, err := json.Marshal(site)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.Sites
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

// Update .
func (c *Client) Update(id int, site *Site) (reply http.HTTPReply, err error) {
	if site.SwitchFabric == "" {
		site.SwitchFabric = "netris"
	}
	js, err := json.Marshal(site)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.Sites + "/" + strconv.Itoa(id)
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *Client) Delete(id int) (reply http.HTTPReply, err error) {
	address := c.client.URL.String() + v2address.Sites + "/" + strconv.Itoa(id)
	reply, err = c.client.Delete(address, nil)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
