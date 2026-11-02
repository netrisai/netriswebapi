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

import (
	"encoding/json"
	"fmt"
	nethttp "net/http"
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

func (c *Client) Get() ([]*DynamicVPCGroup, error) {
	address := c.client.URL.String() + v2address.DynamicVPCGroup
	reply, err := c.client.GetRequest(address)
	if err != nil {
		return nil, fmt.Errorf("{GetDynamicVPCGroup} %s", err)
	}
	data, err := parse(reply)
	if err != nil {
		return nil, fmt.Errorf("{GetDynamicVPCGroup} %s", err)
	}
	var items []*DynamicVPCGroup
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("{GetDynamicVPCGroup} %s", err)
	}
	if items == nil {
		return nil, fmt.Errorf("{GetDynamicVPCGroup} invalid response data")
	}
	for _, item := range items {
		if item == nil || item.ID <= 0 {
			return nil, fmt.Errorf("{GetDynamicVPCGroup} invalid response data")
		}
	}
	return items, nil
}

// GetByID returns nil only when the API responds with HTTP 404.
func (c *Client) GetByID(id int) (*DynamicVPCGroup, error) {
	address := c.client.URL.String() + v2address.DynamicVPCGroup + "/" + strconv.Itoa(id)
	reply, err := c.client.GetRequest(address)
	if err != nil {
		return nil, fmt.Errorf("{GetDynamicVPCGroupByID} %s", err)
	}
	if reply.StatusCode == nethttp.StatusNotFound {
		return nil, nil
	}
	data, err := parse(reply)
	if err != nil {
		return nil, fmt.Errorf("{GetDynamicVPCGroupByID} %s", err)
	}
	var item *DynamicVPCGroup
	if err := json.Unmarshal(data, &item); err != nil {
		return nil, fmt.Errorf("{GetDynamicVPCGroupByID} %s", err)
	}
	if item == nil || item.ID <= 0 {
		return nil, fmt.Errorf("{GetDynamicVPCGroupByID} invalid response data")
	}
	return item, nil
}

func parse(reply http.HTTPReply) (json.RawMessage, error) {
	var result struct {
		Data      json.RawMessage `json:"data"`
		IsSuccess bool            `json:"isSuccess"`
		Message   string          `json:"message"`
	}
	if err := json.Unmarshal(reply.Data, &result); err != nil {
		return nil, fmt.Errorf("HTTP status %d: invalid API response: %s", reply.StatusCode, err)
	}
	if reply.StatusCode != nethttp.StatusOK || !result.IsSuccess {
		return nil, fmt.Errorf("HTTP status %d: %s", reply.StatusCode, result.Message)
	}
	return result.Data, nil
}

func (c *Client) Add(item *DynamicVPCGroupAdd) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(item)
	if err != nil {
		return reply, err
	}
	address := c.client.URL.String() + v2address.DynamicVPCGroup
	return c.client.Post(address, js)
}

func (c *Client) Update(id int, item *DynamicVPCGroupUpdate) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(item)
	if err != nil {
		return reply, fmt.Errorf("{UpdateDynamicVPCGroup} %s", err)
	}
	address := c.client.URL.String() + v2address.DynamicVPCGroup + "/" + strconv.Itoa(id)
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{UpdateDynamicVPCGroup} %s", err)
	}
	return reply, nil
}

func (c *Client) Delete(id int) (reply http.HTTPReply, err error) {
	address := c.client.URL.String() + v2address.DynamicVPCGroup + "/" + strconv.Itoa(id)
	return c.client.Delete(address, nil)
}
