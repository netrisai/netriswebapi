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

package route

import (
	"encoding/json"
	"fmt"

	"github.com/netrisai/netriswebapi/http"
	v1address "github.com/netrisai/netriswebapi/http/addresses/v1"
	"github.com/netrisai/netriswebapi/v1/types/site"
)

type RouteClient struct {
	client *http.HTTPCred
}

func New(c *http.HTTPCred) *RouteClient {
	return &RouteClient{c}
}

func parse(APIResult *http.APIResponse) ([]*Route, error) {
	var sites []*Route
	err := http.Decode(APIResult.Data, &sites)
	if err != nil {
		return sites, fmt.Errorf("{parse} %s", err)
	}
	return sites, nil
}

func (c *RouteClient) Get() ([]*Route, error) {
	sites, err := site.New(c.client).Get()
	if err != nil {
		return nil, err
	}
	siteList := ""
	for _, s := range sites {
		siteList += fmt.Sprintf("filterBySites[]=%d&", s.ID)
	}
	address := c.client.URL.String() + v1address.Routes + "?" + siteList
	APIResult, err := c.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("{GetRoutes} %s", err)
	}

	items, err := parse(APIResult)
	if err != nil {
		return nil, fmt.Errorf("{GetRoutes} %s", err)
	}
	return items, nil
}

func (c *RouteClient) Add(route *RouteAdd) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(route)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.Routes
	reply, err = c.client.Post(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (c *RouteClient) Update(route *RouteAdd) (reply http.HTTPReply, err error) {
	js, err := json.Marshal(route)
	if err != nil {
		return http.HTTPReply{}, fmt.Errorf("{Update} %s", err)
	}
	address := c.client.URL.String() + v1address.Routes
	reply, err = c.client.Put(address, js)
	if err != nil {
		return reply, fmt.Errorf("{Update} %s", err)
	}

	return reply, nil
}

func (c *RouteClient) Delete(id int) (reply http.HTTPReply, err error) {
	lb := struct {
		ID int `json:"id"`
	}{id}
	js, err := json.Marshal(lb)
	if err != nil {
		return reply, err
	}

	address := c.client.URL.String() + v1address.Routes
	reply, err = c.client.Delete(address, js)
	if err != nil {
		return reply, err
	}

	return reply, nil
}
