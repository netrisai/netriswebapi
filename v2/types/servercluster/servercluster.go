/*
Copyright 2024. Netris, Inc.

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

package servercluster

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

func parse(APIResult *http.APIResponse) ([]*ServerCluster, error) {
	var items []*ServerCluster
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}

func parseSingle(APIResult *http.APIResponse) (*ServerCluster, error) {
	var items *ServerCluster
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parse} %s", err)
	}
	return items, nil
}

func (c *Client) Get() ([]*ServerCluster, error) {
	address := c.client.URL.String() + v2address.ServerCluster
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetServerCluster} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetServerCluster} %s", err)
	}
	return items, nil
}

func (c *Client) GetByID(id int) (*ServerCluster, error) {
	address := c.client.URL.String() + v2address.ServerCluster + "/" + strconv.Itoa(id)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetServerClusterByID} %s", err)
	}

	item, err := parseSingle(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetServerClusterByID} %s", err)
	}
	return item, nil
}

func (c *Client) Add(servercluster *ServerClusterW) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(servercluster)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.ServerCluster
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *Client) Update(id int, servercluster *ServerClusterU) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(servercluster)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{UpdateServerCluster} %s", err)
	}
	address := c.client.URL.String() + v2address.ServerCluster + "/" + strconv.Itoa(id)
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{UpdateServerCluster} %s", err)
	}

	return reply, nil
}

func (c *Client) Delete(id int) (reply http.HTTPReply, err error) {
	address := c.client.URL.String() + v2address.ServerCluster + "/" + strconv.Itoa(id)
	reply, err = c.client.Delete(address, nil)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
