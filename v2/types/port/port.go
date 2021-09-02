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

package port

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/netrisai/netriswebapi/http"
	v2address "github.com/netrisai/netriswebapi/http/addresses/v2"
)

type PortClient struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *PortClient {
	return &PortClient{c}
}

func parse(APIResult *http.APIResponse) ([]*Port, error) {
	var items []*Port
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parseAPIVnet} %s", err)
	}
	return items, nil
}

func (c *PortClient) Get() ([]*Port, error) {
	address := c.client.URL.String() + v2address.Ports
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}
	return items, nil
}

func (c *PortClient) GetByID(id int) (*Port, error) {
	address := c.client.URL.String() + v2address.Ports + "/" + strconv.Itoa(id)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetByID} %s", err)
	}

	var port *Port
	err = http.Decode(APIResult.Data, &port)
	if err != nil {
		return port, fmt.Errorf("{GetByID} %s", err)
	}
	return port, nil
}

func (c *PortClient) UpdateList(ports []*PortUpdate) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(ports)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{UpdatePort} %s", err)
	}
	address := c.client.URL.String() + v2address.Ports
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{UpdatePort} %s", err)
	}

	return reply, nil
}

func (c *PortClient) GetExtenstion() ([]*PortExtension, error) {
	address := c.client.URL.String() + v2address.PortExtensions
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetExtenstion} %s", err)
	}

	var items []*PortExtension
	err = http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{GetExtenstion} %s", err)
	}
	return items, nil
}

func (c *PortClient) DeleteExtension(id int) (reply http.HTTPReply, err error) {
	address := c.client.URL.String() + v2address.PortExtensions + "/" + strconv.Itoa(id)
	reply, err = c.client.Delete(address, nil)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *PortClient) AddToLAG(port *PortLAG) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(port)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.PortLAG
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *PortClient) FreeUP(ports []*IDName) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(ports)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{FreeUP} %s", err)
	}
	address := c.client.URL.String() + v2address.PortFreeUP
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{FreeUP} %s", err)
	}

	return reply, nil
}
