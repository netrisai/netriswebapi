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

package ipam

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/netrisai/netriswebapi/http"
	v2address "github.com/netrisai/netriswebapi/http/addresses/v2"
)

type IPAMClient struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *IPAMClient {
	return &IPAMClient{c}
}

func parse(APIResult *http.APIResponse) ([]*IPAM, error) {
	var items []*IPAM
	err := http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{parseInventories} %s", err)
	}
	return items, nil
}

func (c *IPAMClient) Get() ([]*IPAM, error) {
	address := c.client.URL.String() + v2address.IPAMBase
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetIPAM} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetIPAM} %s", err)
	}
	return items, nil
}

func (c *IPAMClient) GetBySites(sites []int) ([]*IPAM, error) {
	siteList := ""
	for _, s := range sites {
		siteList += fmt.Sprintf("filterBySites[]=%d&", s)
	}

	address := c.client.URL.String() + v2address.IPAMBase + "?useSites=true&" + siteList
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetIPAMBySites} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetIPAMBySites} %s", err)
	}
	return items, nil
}

func (c *IPAMClient) GetSubnets() ([]*IPAM, error) {
	address := c.client.URL.String() + v2address.IPAMSubnets
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

func (c *IPAMClient) GetHosts(id int) ([]*Host, error) {
	address := c.client.URL.String() + v2address.IPAMHosts + "/" + strconv.Itoa(id)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetHosts} %s", err)
	}

	var items []*Host
	err = http.Decode(APIResult.Data, &items)
	if err != nil {
		return items, fmt.Errorf("{GetHosts} %s", err)
	}
	return items, nil
}

func (c *IPAMClient) AddAllocation(allocation *Allocation) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(allocation)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.IPAMAllocation
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *IPAMClient) UpdateAllocation(id int, allocation *Allocation) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(allocation)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{UpdateAllocation} %s", err)
	}
	address := c.client.URL.String() + v2address.IPAMAllocation + "/" + strconv.Itoa(id)
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{UpdateAllocation} %s", err)
	}

	return reply, nil
}

func (c *IPAMClient) AddSubnet(subnet *Subnet) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(subnet)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v2address.IPAMSubnet
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *IPAMClient) UpdateSubnet(id int, subnet *Subnet) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(subnet)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{UpdateSubnetn} %s", err)
	}
	address := c.client.URL.String() + v2address.IPAMSubnet + "/" + strconv.Itoa(id)
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{UpdateSubnetn} %s", err)
	}

	return reply, nil
}

func (c *IPAMClient) Delete(kind string, id int) (reply http.HTTPReply, err error) {
	if !(kind == "allocation" || kind == "subnet") {
		return reply, fmt.Errorf("Invalid hardware type. Available values: allocation, subnet")
	}

	address := c.client.URL.String() + v2address.IPAMBase + "/" + kind + "/" + strconv.Itoa(id)
	reply, err = c.client.Delete(address, nil)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *IPAMClient) DeleteByPrefix(kind, prefix string) (reply http.HTTPReply, err error) {
	if !(kind == "allocation" || kind == "subnet") {
		return reply, fmt.Errorf("Invalid hardware type. Available values: allocation, subnet")
	}

	address := c.client.URL.String() + v2address.IPAMBase + "/" + kind + "/" + prefix
	reply, err = c.client.Delete(address, nil)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
