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

package subnet

import (
	"encoding/json"
	"fmt"

	"github.com/netrisai/netriswebapi/http"
	v1address "github.com/netrisai/netriswebapi/http/addresses/v1"
)

type SubnetClient struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *SubnetClient {
	return &SubnetClient{c}
}

func parse(APIResult *http.APIResponse) ([]*Subnet, error) {
	var sites []*Subnet
	err := http.Decode(APIResult.Data, &sites)
	if err != nil {
		return sites, fmt.Errorf("{parse} %s", err)
	}
	return sites, nil
}

func (c *SubnetClient) Get() ([]*Subnet, error) {
	address := c.client.URL.String() + v1address.Subnets
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetSubnets} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetSubnets} %s", err)
	}
	return items, nil
}

func (c *SubnetClient) Add(subnet *SubnetAdd) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(subnet)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.Subnets
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *SubnetClient) Update(subnet *Subnet) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(subnet)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{Update} %s", err)
	}
	address := c.client.URL.String() + v1address.Subnets
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{Update} %s", err)
	}

	return reply, nil
}
