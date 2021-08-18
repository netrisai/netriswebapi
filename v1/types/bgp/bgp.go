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

package bgp

import (
	"encoding/json"
	"fmt"

	"github.com/netrisai/netriswebapi/http"
	v1address "github.com/netrisai/netriswebapi/http/addresses/v1"
	"github.com/netrisai/netriswebapi/v1/types/site"
)

type BGPClient struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *BGPClient {
	return &BGPClient{c}
}

func (c *BGPClient) Get() ([]*EBGP, error) {
	sites, err := site.New(c.client).Get()
	if err != nil {
		return nil, err
	}
	siteList := ""
	for _, s := range sites {
		siteList += fmt.Sprintf("selectedSites[]=%d&", s.ID)
	}
	address := c.client.URL.String() + v1address.BGP + "?" + siteList
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

func (c *BGPClient) GetSites() ([]*EBGPSite, error) {
	address := c.client.URL.String() + v1address.BGPSites
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}

	items, err := parseSites(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}
	return items, nil
}

func (c *BGPClient) GetVNets() ([]*EBGPVNet, error) {
	address := c.client.URL.String() + v1address.BGPVNets
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}

	items, err := parseVNets(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}
	return items, nil
}

func (c *BGPClient) GetRouteMaps() ([]*EBGPRouteMap, error) {
	address := c.client.URL.String() + v1address.BGPRouteMaps
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}

	items, err := parseRouteMaps(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}
	return items, nil
}

func (c *BGPClient) GetOffloaders(siteID int) ([]*EBGPOffloader, error) {
	address := c.client.URL.String() + fmt.Sprintf("%s?id=%d", v1address.BGPOffloaders, siteID)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}

	items, err := parseOffloaders(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}
	return items, nil
}

func (c *BGPClient) GetPorts(siteID int) ([]*EBGPPort, error) {
	address := c.client.URL.String() + fmt.Sprintf("%s?id=%d", v1address.BGPPorts, siteID)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}

	items, err := parsePorts(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}
	return items, nil
}

func (c *BGPClient) GetSwitches(siteID int) ([]*EBGPSwitch, error) {
	address := c.client.URL.String() + fmt.Sprintf("%s?id=%d", v1address.BGPSwitches, siteID)
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}

	items, err := parseSwitches(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}
	return items, nil
}

func (c *BGPClient) GetUpdateSources() ([]*EBGPUpdatedSource, error) {
	address := c.client.URL.String() + v1address.BGPUpdatedSources
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}

	items, err := parseUpdateSources(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{Get} %s", err)
	}
	return items, nil
}

func (c *BGPClient) Add(bgp *EBGPAdd) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(bgp)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.BGP
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *BGPClient) Update(bgp *EBGPUpdate) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(bgp)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{Update} %s", err)
	}
	address := c.client.URL.String() + v1address.BGP
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{Update} %s", err)
	}

	return reply, nil
}

func (c *BGPClient) Delete(id int) (reply http.HTTPReply, err error) {
	lb := struct {
		ID int `json:"id"`
	}{id}
	js, err := json.Marshal(lb)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.BGP
	reply, err = c.client.Delete(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
